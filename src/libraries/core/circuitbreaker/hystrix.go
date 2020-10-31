package circuitbreaker

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/afex/hystrix-go/hystrix"
	"github.com/eapache/go-resiliency/retrier"
	"github.com/spf13/viper"
	"go.uber.org/zap"

	"github.com/BlackspaceInc/common/messaging/rabbitmq"
	"github.com/BlackspaceInc/common/tracing"
	"github.com/BlackspaceInc/common/util"
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
	// AMQPClient witholds a reference to the rabbitMQ client
	AMQPClient *rabbitmq.RabbitMQClient
	// HTTPClient is an http client through which http operations will be performed
	HTTPClient *http.Client
}

// NewCircuitBreaker returns a new reference to a circuit breaker
func NewCircuitBreaker(logger *zap.Logger, retries int, amqpClient *rabbitmq.RabbitMQClient, client *http.Client) *CircuitBreaker {
	return &CircuitBreaker{
		Logger:     logger,
		Retries:    retries,
		AMQPClient: amqpClient,
		HTTPClient: client,
	}
}

// ConfigureHystrix sets up hystrix circuit breakers.
func (cb *CircuitBreaker) ConfigureHystrix(commands []string) {
	// assign to global logger
	for _, command := range commands {
		hystrix.ConfigureCommand(command, hystrix.CommandConfig{
			Timeout:                resolveProperty(command, "Timeout"),
			MaxConcurrentRequests:  resolveProperty(command, "MaxConcurrentRequests"),
			ErrorPercentThreshold:  resolveProperty(command, "ErrorPercentThreshold"),
			RequestVolumeThreshold: resolveProperty(command, "RequestVolumeThreshold"),
			SleepWindow:            resolveProperty(command, "SleepWindow"),
		})

		cb.Logger.Info(fmt.Sprintf("Circuit %v settings: %v\n", command, hystrix.GetCircuitSettings()[command]))
	}

	// define a server able to display dashboard metrics
	hystrixStreamHandler := hystrix.NewStreamHandler()

	// start watching in-memory circuit breakers for metrics
	hystrixStreamHandler.Start()

	// associate all requests at the defined port below to the dashboard metrics
	go http.ListenAndServe(net.JoinHostPort("", "8181"), hystrixStreamHandler)
	cb.Logger.Info("Launched hystrixStreamHandler at 8181")

	// Publish presence on RabbitMQ continually
	// cb.publishDiscoveryToken()
}

// publishDiscoveryToken publishes a discovery token on the queue of interest
func (cb *CircuitBreaker) publishDiscoveryToken() {
	ip, err := util.ResolveIPFromHostsFile()
	if err != nil {
		ip = util.GetIPWithPrefix("10.0.")
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
	errors := hystrix.Go(breakerName, func() error {
		// construct new http request
		req, _ := http.NewRequest(method, url, nil)
		// add tracing to request from context
		tracing.AddTracingToReqFromContext(ctx, req)
		// attempt http call with retries configured
		err := cb.callRequestWithRetries(req, output)
		return err
	}, fallback)

	select {
	case out := <-output:
		cb.Logger.Debug(fmt.Sprintf("Call in breaker %v successful", breakerName))
		return out, nil

	case err := <-errors:
		cb.Logger.Error(fmt.Sprintf("Got error on channel in breaker %v. Msg: %v", breakerName, err.Error()))
		return nil, err
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
	errors := hystrix.Go(breakerName, func() error {
		// add tracing
		tracing.AddTracingToReqFromContext(ctx, req)
		// perform the request
		err := cb.callRequestWithRetries(req, output)
		return err
	}, fallback)

	select {
	case out := <-output:
		cb.Logger.Debug(fmt.Sprintf("Call in breaker %v successful", breakerName))
		return out, nil

	case err := <-errors:
		cb.Logger.Error(fmt.Sprintf("Got error on channel in breaker %v. Msg: %v", breakerName, err.Error()))
		return nil, err
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

	// perform operation with circuit breaker
	errors := hystrix.Go(breakerName, func() error {
		// generate empty request object for tracing handler to populate
		req := &http.Request{
			Header: http.Header{},
		}
		// add tracing to client request
		tracing.AddTracingToReqFromContext(ctx, req)
		err := cb.callFunctionWithRetries(ctx, fn, output)
		return err
	}, fallback)

	select {
	case out := <-output:
		cb.Logger.Debug(fmt.Sprintf("Call in breaker %v successful", breakerName))
		return out, nil

	case err := <-errors:
		cb.Logger.Error(fmt.Sprintf("Got error on channel in breaker %v. Msg: %v", breakerName, err.Error()))
		return nil, err
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
	errors := hystrix.Go(breakerName, func() error {
		// generate empty request object for tracing handler to populate
		req := &http.Request{
			Header: http.Header{},
		}
		// add tracing to client request
		tracing.AddTracingToReqFromContext(ctx, req)
		err := cb.CallCustomFuncWithRetries(ctx, fn)
		return err
	}, fallback)

	select {
	case err := <-errors:
		cb.Logger.Error(fmt.Sprintf("Got error on channel in breaker %v. Msg: %v", breakerName, err.Error()))
		return err
	default:
		return nil
	}
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

// Deregister publishes a Deregister token to Hystrix/Turbine
func (cb *CircuitBreaker) Deregister() {
	ip, err := util.ResolveIPFromHostsFile()
	if err != nil {
		ip = util.GetIPWithPrefix("10.0.")
	}
	token := DiscoveryToken{
		State:   "DOWN",
		Address: ip,
	}
	bytes, _ := json.Marshal(token)
	_ = cb.AMQPClient.PublishOnQueue(context.TODO(), bytes, "discovery")
	cb.Logger.Info("Sent deregistration token over SpringCloudBus")
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

// DiscoveryToken defines a struct for transmitting the state of a hystrix stream producer.
type DiscoveryToken struct {
	State   string `json:"state"` // UP, RUNNING, DOWN ??
	Address string `json:"address"`
}
