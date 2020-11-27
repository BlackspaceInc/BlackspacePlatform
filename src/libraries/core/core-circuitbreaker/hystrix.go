package core_circuitbreaker

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/afex/hystrix-go/hystrix"
	"github.com/eapache/go-resiliency/retrier"
	"github.com/sony/gobreaker"
	"github.com/spf13/viper"
	"go.uber.org/zap"

	core_messaging "github.com/BlackspaceInc/BlackspacePlatform/src/libraries/core/core-messaging/rabbitmq"
)


type FallbackFunc func(error) error

func init() {
	log.SetOutput(ioutil.Discard)
}

// A CircuitBreaker represents an active circuit breaker object that in which operations can
// be performed within.
type CircuitBreaker struct {
	// Logger holds a reference to the zap logger loging handle
	Logger *zap.Logger
	// Retries specifies the number of retries to perform within a circuit breaker
	Retries int
	// AMQP Client witholds a connection to some underlying message queue connection
	AMQPClient *core_messaging.RabbitMQClient
	// HTTPClient is an http client through which http operations will be performed
	HTTPClient *http.Client
	// Circuit Breaker
	Breaker *gobreaker.CircuitBreaker
}

// NewCircuitBreaker returns a new reference to a circuit breaker
func NewCircuitBreaker(logger *zap.Logger, retries int, amqpClient *core_messaging.RabbitMQClient, client *http.Client) *CircuitBreaker {
	cbreaker := &CircuitBreaker{
		Logger:     logger,
		Retries:    retries,
		AMQPClient: amqpClient,
		HTTPClient: client,
		Breaker: cb,
	}

	cbreaker.Logger.Info("configured circuit breaker")

	// Publish presence on RabbitMQ continually
	cbreaker.publishDiscoveryToken()
	return cbreaker
}

// publishDiscoveryToken publishes a discovery token on the queue of interest
func (cb *CircuitBreaker) publishDiscoveryToken() {
	ip, err := ResolveIPFromHostsFile()
	if err != nil {
		ip = GetIPWithPrefix("10.0.")
	}
	token := DiscoveryToken{
		State:   "UP",
		Address: ip,
	}
	bytes, _ := json.Marshal(token)
	go func() {
		for {
			ctx := context.TODO()
			_ = cb.AMQPClient.PublishOnQueue(ctx, bytes, "discovery")
			time.Sleep(time.Second * 30)
		}
	}()
}

// PerformHTTPRequest performs a HTTP call inside a circuit breaker.
func (cb *CircuitBreaker) PerformHTTPRequest(ctx context.Context, breakerName string, url string, method string, fallback func(error) error) ([]byte, error) {

	if fallback == nil {
		fallback = func(err error) error {
			cb.Logger.Error(fmt.Sprintf("In fallback function for breaker %v, error: %v", breakerName), zap.Error(err))
			circuit, _, _ := hystrix.GetCircuit(breakerName)
			cb.Logger.Error(fmt.Sprintf("Circuit state is: %v", circuit.IsOpen()))
			return err
		}
	}

	output := make(chan []byte, 1)

	_, err := cb.Breaker.Execute(func() (interface{}, error){
		// construct new http request
		req, _ := http.NewRequest(method, url, nil)
		// TODO: add tracing to request from context
		// attempt http call with retries configured
		if err := cb.callRequestWithRetries(req, output); err != nil {
			return nil, err
		}

		return nil, nil
	})

	if err != nil {
		cb.Logger.Error(fmt.Sprintf("Got error on channel in breaker %v. Msg: %v", breakerName, err.Error()))
		return nil, err
	}

	select {
	case out := <-output:
		cb.Logger.Debug(fmt.Sprintf("Call in breaker %v successful", breakerName))
		return out, nil
	}
}

// PerformRequest performs the supplied http.Request within a circuit breaker.
func (cb *CircuitBreaker) PerformRequest(
	ctx context.Context,
	breakerName string,
	req *http.Request,
	fallback func(error) error) ([]byte, error) {

	if fallback == nil {
		fallback = func(err error) error {
			cb.Logger.Error(fmt.Sprintf("In fallback function for breaker %v, error: %v", breakerName, err.Error()))
			return err
		}
	}

	// attempt http request through circuit breaker
	output := make(chan []byte, 1)

	_, err := cb.Breaker.Execute(func() (interface{}, error){
		// TODO: add tracing
		// attempt http call with retries configured
		if err := cb.callRequestWithRetries(req, output); err != nil {
			return nil, err
		}

		return nil, nil
	})

	if err != nil {
		cb.Logger.Error(fmt.Sprintf("Got error on channel in breaker %v. Msg: %v", breakerName, err.Error()))
		return nil, err
	}

	select {
	case out := <-output:
		cb.Logger.Debug(fmt.Sprintf("Call in breaker %v successful", breakerName))
		return out, nil
	}
}

// PerformServiceClientRequestCircuitBreaker performs the supplied service level request within a circuit breaker.
func (cb *CircuitBreaker) PerformServiceRequest(ctx context.Context, breakerName string, fn func() (interface{}, error), fallback func(error) error) (interface{}, error) {
	output := make(chan interface{}, 1)

	if fallback == nil {
		fallback = func(err error) error {
			cb.Logger.Error(fmt.Sprintf("In fallback function for breaker %v, error: %v", breakerName, err.Error()))
			return err
		}
	}

	_, err := cb.Breaker.Execute(func() (interface{}, error){
		// TODO: add tracing
		if err := cb.callFunctionWithRetries(ctx, fn, output); err != nil {
			return nil, err
		}
		return nil, nil
	})

	if err != nil {
		cb.Logger.Error(fmt.Sprintf("Got error on channel in breaker %v. Msg: %v", breakerName, err.Error()))
		return nil, err
	}

	select {
	case out := <-output:
		cb.Logger.Debug(fmt.Sprintf("Call in breaker %v successful", breakerName))
		return out, nil
	}
}

// PerformCustomServiceClientRequestCircuitBreaker performs the supplied service level request within a circuit breaker.
func (cb *CircuitBreaker) PerformCustomServiceRequest(ctx context.Context, breakerName string, fn func() error, fallback func(error) error) error {

	if fallback == nil {
		fallback = func(err error) error {
			cb.Logger.Error(fmt.Sprintf("In fallback function for breaker %v, error: %v", breakerName, err.Error()))
			return err
		}
	}

	_, err := cb.Breaker.Execute(func() (interface{}, error){
		// TODO: add tracing
		if err := cb.CallCustomFuncWithRetries(ctx, fn); err != nil {
			return nil, err
		}
		return nil, nil
	})

	if err != nil {
		cb.Logger.Error(fmt.Sprintf("Got error on channel in breaker %v. Msg: %v", breakerName, err.Error()))
		return err
	}

	return nil
}

func (cb *CircuitBreaker) callRequestWithRetries(req *http.Request, output chan []byte) error {
	r := retrier.New(retrier.ConstantBackoff(cb.Retries, 100*time.Millisecond), nil)
	attempt := 0
	err := r.Run(func() error {
		attempt++
		resp, err := cb.HTTPClient.Do(req)
		if err == nil && resp.StatusCode < 299 {
			responseBody, err := ioutil.ReadAll(resp.Body)
			if err == nil {
				output <- responseBody
				return nil
			}
			return err
		} else if err == nil {
			err = fmt.Errorf("Status was %v", resp.StatusCode)
		}

		cb.Logger.Error(fmt.Sprintf("Retrier failed, attempt %v", attempt))
		return err
	})
	return err
}

// callFuncWithRetries calls a given function in a retry loop
func (cb *CircuitBreaker) callFunctionWithRetries(ctx context.Context, fn func() (interface{}, error), output chan interface{}) error {
	// TODO - make backoff configurable instead of constant
	r := retrier.New(retrier.ConstantBackoff(cb.Retries, 50*time.Millisecond), nil)
	attempt := 0
	err := r.Run(func() error {
		attempt++
		res, err := fn()
		cb.Logger.Info("result", zap.Any("returned result", res))
		if err == nil {
			output <- res
		}

		cb.Logger.Error(fmt.Sprintf("Retrier failed, attempt %v", attempt))
		return err
	})

	return err
}

// CallCustomFuncWithRetries calls a given function in a retry loop
func (cb *CircuitBreaker) CallCustomFuncWithRetries(ctx context.Context, fn func() error) error {
	// TODO - make backoff configurable instead of constant
	r := retrier.New(retrier.ConstantBackoff(cb.Retries, 50*time.Millisecond), nil)
	attempt := 0
	err := r.Run(func() error {
		attempt++
		err := fn()
		cb.Logger.Error(fmt.Sprintf("Retrier failed, attempt %v", attempt))
		return err
	})

	return err
}

// Deregister publishes a Deregister token to message bus
func (cb *CircuitBreaker) Deregister() {
	ip, err := ResolveIPFromHostsFile()
	if err != nil {
		ip = GetIPWithPrefix("10.0.")
	}
	token := DiscoveryToken{
		State:   "DOWN",
		Address: ip,
	}
	bytes, _ := json.Marshal(token)
	_ = cb.AMQPClient.PublishOnQueue(context.TODO(), bytes, "discovery")
	cb.Logger.Info("Sent deregistration token over bus")
}

func resolveProperty(command string, prop string) int {
	if viper.IsSet("hystrix.command." + command + "." + prop) {
		return viper.GetInt("hystrix.command." + command + "." + prop)
	}
	return getDefaultHystrixConfigPropertyValue(prop)
}

func getDefaultHystrixConfigPropertyValue(prop string) int {
	switch prop {
	case "Timeout":
		return 1000 // hystrix.DefaultTimeout
	case "MaxConcurrentRequests":
		return 200 // hystrix.DefaultMaxConcurrent
	case "RequestVolumeThreshold":
		return hystrix.DefaultVolumeThreshold
	case "SleepWindow":
		return hystrix.DefaultSleepWindow
	case "ErrorPercentThreshold":
		return hystrix.DefaultErrorPercentThreshold
	}
	panic("Got unknown hystrix property: " + prop + ". Panicing!")
}

// DiscoveryToken defines a struct for transmitting the state of a circuit breaker stream producer.
type DiscoveryToken struct {
	State   string `json:"state"` // UP, RUNNING, DOWN ??
	Address string `json:"address"`
}
