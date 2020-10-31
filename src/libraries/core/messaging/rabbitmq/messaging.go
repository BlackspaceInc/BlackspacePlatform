package rabbitmq

import (
	"context"
	"errors"

	"github.com/opentracing/opentracing-go"
	"github.com/streadway/amqp"
	"go.uber.org/zap"

	"github.com/BlackspaceInc/common/tracing"
)

// IMessagingClient defines an interface for connection and consuming messages
type IRabbitMQMessagingClient interface {
	PublishOnQueue(ctx context.Context, body []byte, exchange string) error
	SubscribeToQueue(consumerExchangeName string, handlerFunc func(amqp.Delivery) error) error
	Close()
}

var DefaultConnectionString = "amqp://guest:guest@localhost:5672"

type RabbitMQClient struct {
	Conn          *amqp.Connection
	Chann         *amqp.Channel
	Logger        *zap.Logger
	QueueBindings map[string]*amqp.Queue
}

type Queues struct {
	QueueDetails []QueueDetails
	Exchanges    []Exchange
}

type Exchange struct {
	ExchangeName string
	ExchangeType string
	Durable      bool
	AutoDelete   bool
	Internal     bool
	NoWait       bool
	Args         *amqp.Table
}

type QueueDetails struct {
	QueueName    string
	ExchangeName string
	Durable      bool
	AutoDelete   bool
	Exclusive    bool
	NoWait       bool
	Args         *amqp.Table
}

// InitiateQueues takes as input a queue to exchange mapping an returns reference to a queue
func InitiateQueues(queueToExchangeMap map[string]Exchange) Queues {
	var queue Queues
	for key, value := range queueToExchangeMap {
		details := QueueDetails{
			QueueName:    key,
			ExchangeName: value.ExchangeName,
			Durable:      value.Durable,
			AutoDelete:   value.AutoDelete,
			Exclusive:    false,
			NoWait:       value.NoWait,
			Args:         value.Args,
		}

		queue.QueueDetails = append(queue.QueueDetails, details)
		queue.Exchanges = append(queue.Exchanges, value)
	}

	return queue
}

// New creates an amqp broker connection and binds a set of input queues to each specific exchange
// Note all queues that the services writes to and consumes must be present and passed as input parameters
func New(url string, queues Queues, logger *zap.Logger) (*RabbitMQClient, error) {
	queueBindingsMap := make(map[string]*amqp.Queue)

	if url == " " {
		logger.Info("empty amqp connection string, reverting to default one")
		// fallback to default connection string if one isn't specified
		url = DefaultConnectionString
	}

	// Connect to the rabbitMQ instance
	connection, err := amqp.Dial(url)
	if err != nil {
		return nil, err
	}

	channel, err := connection.Channel()
	if err != nil {
		return nil, err
	}

	// create the exchanges and tie such exchanges to a each created queue
	for _, exchange := range queues.Exchanges {
		var currArgs amqp.Table
		if exchange.Args == nil {
			currArgs = nil
		} else {
			currArgs = *exchange.Args
		}
		// declare an exchange type names exchange name which we'll use to publish messages
		err := channel.ExchangeDeclare(
			exchange.ExchangeName,
			exchange.ExchangeType,
			exchange.Durable,    // durable : true
			exchange.AutoDelete, // autodelete : false
			exchange.Internal,   // internal : false
			exchange.NoWait,     // nowait : false
			currArgs)            // nil
		if err != nil {
			return nil, err
		}
	}

	// declare queues and bind queues to such exchanges
	for _, details := range queues.QueueDetails {
		var currArgs amqp.Table
		if details.Args == nil {
			currArgs = nil
		} else {
			currArgs = *details.Args
		}

		q, err := channel.QueueDeclare(
			details.QueueName,
			details.Durable,    // durable : true
			details.AutoDelete, // autodelete : false
			details.Exclusive,  // exclusive : false
			details.NoWait,     // nowait : false
			currArgs)           // nil
		if err != nil {
			return nil, err
		}
		err = channel.QueueBind(
			details.QueueName,
			details.ExchangeName,
			details.ExchangeName,
			details.NoWait, // nowait : false
			currArgs)

		if err != nil {
			return nil, err
		}
		// add each queue to a hashmap comprised of the exchange to queue mapping
		queueBindingsMap[details.ExchangeName] = &q
	}

	return &RabbitMQClient{Conn: connection, Chann: channel, Logger: logger, QueueBindings: queueBindingsMap}, nil
}

// PublishOnQueue publishes a message on a queue based on exchange name
func (m *RabbitMQClient) PublishOnQueue(ctx context.Context, body []byte, exchange string) error {
	if ctx == nil {
		return m.publish(body, exchange)
	}

	return m.publishWithContext(ctx, body, exchange)
}

// publish publishes a message byte string on a queue tied to an exchange
func (m *RabbitMQClient) publish(body []byte, exchangeName string) error {
	return m.publishOnQueue(amqp.Publishing{Body: body}, exchangeName)
}

// publishWithContext published a message byte string on a queue with an associated context useful for tracing
func (m *RabbitMQClient) publishWithContext(ctx context.Context, body []byte, exchangeName string) error {
	message := m.buildMessage(ctx, body)
	return m.publishOnQueue(message, exchangeName)
}

// publishOnQueue publishes a message on a queue
func (m *RabbitMQClient) publishOnQueue(msg amqp.Publishing, exchangeName string) error {
	if m.Conn == nil {
		m.Logger.Error("attempted to send message before a connection was initialized.")
		return errors.New("attempted to send message before connection was initialized")
	}

	// attempt to extract a reference to the queue from the bindings map present in the rabbitMQ client object
	// Note. if the queue does infact exist, then it must have been boound to the input exchange
	if m.queueExists(exchangeName) {
		// send the message to the channel
		if err := m.Chann.Publish(exchangeName, exchangeName, false, false, msg); err != nil {
			m.Logger.Error(err.Error())
			return err
		}
		return nil
	}

	// return an error as the exchange name must have been tied to a queue upon instantiation of the
	// amqp connection object
	return errors.New("invalid exchange name. no queue bound to the exchange. please reinitialize the RabbitMqClient object with the proper set of queues and exchange names")
}

// SubscribeToQueue subscribes a callback to a queue tied to an exchange
func (m *RabbitMQClient) SubscribeToQueue(consumerExchangeName string, callback func(amqp.Delivery) error) error {
	if m.Conn == nil {
		m.Logger.Error("attempted to send message before a connection was initialized.")
		return errors.New("attempted to send message before connection was initialized")
	}

	defer m.Chann.Close()
	defer m.Conn.Close()

	// check that a queue tied to the consumer exchange name exists in our exchange to queue mapping
	if m.queueExists(consumerExchangeName) {
		queue := m.QueueBindings[consumerExchangeName]
		// get queued messages
		// TODO: consider implementing workers goroutines that routinely poll for rabbitmq and using channels distribute operation
		msgs, err := m.Chann.Consume(queue.Name, consumerExchangeName, false, false, false, false, nil)
		if err != nil {
			m.Logger.Error(err.Error())
			return err
		}

		// TODO: make a background go routine handle this
		m.Logger.Info("commencing processing of messages from queue", zap.String("QueueName", queue.Name))
		for d := range msgs {
			// Invoke the handlerFunc func we passed as parameter.
			if err := callback(d); err != nil {
				m.Logger.Error("Handlerfunc returned an error. Requeuing message", zap.Error(err))
				_ = d.Nack(false, true)
			}
			// acknowledge the message
			_ = d.Ack(false)
		}
		m.Logger.Info("successfully processed set of messages from queue", zap.String("QueueName", queue.Name))
	}

	// return an error as the exchange name must have been tied to a queue upon instantiation of the
	// amqp connection object
	return errors.New("invalid exchange name. no queue bound to the exchange. please reinitialize the RabbitMqClient object with the proper set of queues and exchange names")
}

// Close closes the connection to the AMQP-broker, if available.
func (m *RabbitMQClient) Close() {
	if m.Conn != nil {
		m.Logger.Info("Closing connection to AMQP broker")
		_ = m.Conn.Close()
	}
}

// queueExists asserts a queue pointer reference exists in the queuebindings object
func (m *RabbitMQClient) queueExists(exchange string) bool {
	if _, ok := m.QueueBindings[exchange]; ok {
		return ok
	}
	return false
}

//  buildMessage traces the interaction between a client and a queue
func (m *RabbitMQClient) buildMessage(ctx context.Context, body []byte) amqp.Publishing {
	publishing := amqp.Publishing{
		ContentType: "application/json",
		Body:        body, // Our JSON body as []byte
	}
	if ctx != nil {
		child := tracing.StartChildSpanFromContext(ctx, "messaging")
		defer child.Finish()
		var val = make(opentracing.TextMapCarrier)
		err := tracing.AddTracingToTextMapCarrier(child, val)
		if err != nil {
			m.Logger.Error("Error injecting span context.", zap.Error(err))
		} else {
			publishing.Headers = tracing.CarrierToMap(val)
		}
	}
	return publishing
}
