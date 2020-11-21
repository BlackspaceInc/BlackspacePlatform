package rabbitmq

import (
	"strings"

	"go.uber.org/zap"
)

func ParseAndCreateQueueReference(consumingQueues, producingQueues string, numProducerQueues, numConsumerQueues int, logger *zap.Logger) Queues {
	// parse both string based in comma seperator
	consumerQueueSet := strings.SplitN(consumingQueues, ",", numProducerQueues)
	producerQueueSet := strings.SplitN(producingQueues, ",", numConsumerQueues)

	logger.Info("Producer Queues", zap.Any("Queues", producerQueueSet))
	logger.Info("Consumer Queues", zap.Any("Queues", consumerQueueSet))

	queueToExchangeMapping := make(map[string]Exchange)

	// populate the queue to exchange mapping
	populateQueueToExchangeMapping(consumerQueueSet, queueToExchangeMapping, logger)
	populateQueueToExchangeMapping(producerQueueSet, queueToExchangeMapping, logger)

	// initialize a queue from the queue to exchange mapping
	return InitiateQueues(queueToExchangeMapping)
}

func ParseAndCreateQueueReferences(consumingQueues, producingQueues []string, numProducerQueues, numConsumerQueues int, logger *zap.Logger) Queues {
	// parse both string based in comma seperator
	logger.Info("Producer Queues", zap.Any("Queues", producingQueues))
	logger.Info("Consumer Queues", zap.Any("Queues", consumingQueues))

	queueToExchangeMapping := make(map[string]Exchange)

	// populate the queue to exchange mapping
	populateQueueToExchangeMapping(producingQueues, queueToExchangeMapping, logger)
	populateQueueToExchangeMapping(consumingQueues, queueToExchangeMapping, logger)

	// initialize a queue from the queue to exchange mapping
	return InitiateQueues(queueToExchangeMapping)
}

func populateQueueToExchangeMapping(queueSet []string, queueToExchangeMapping map[string]Exchange, logger *zap.Logger) {
	// from parsed set of queues extract the queue name and exchange type and create a queue instance
	for _, queueName := range queueSet {
		if queueName != "" {
			queueDetails := strings.SplitN(queueName, ":", 2)
			logger.Info("details", zap.Any("queue details", queueDetails))

			name := queueDetails[0]
			exchangeType := queueDetails[1]

			// TODO: Look into the impact of each respective params and enable where fitting
			exchange := Exchange{
				ExchangeName: name,
				ExchangeType: exchangeType,
				Durable:      true,
				AutoDelete:   false,
				Internal:     false,
				NoWait:       false,
				Args:         nil,
			}

			if _, ok := queueToExchangeMapping[queueName]; !ok {
				queueToExchangeMapping[name] = exchange
			}
		}
	}
}
