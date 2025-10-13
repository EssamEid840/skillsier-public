package kafka

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"

	"users-be/internal/config"

	"github.com/IBM/sarama"
)

// Consumer wraps a Kafka consumer
type Consumer struct {
	consumer sarama.ConsumerGroup
	topics   []string
	handler  sarama.ConsumerGroupHandler
}

// MessageHandler is a function that processes Kafka messages
type MessageHandler func(ctx context.Context, message *sarama.ConsumerMessage) error

// NewConsumer creates a new Kafka consumer
func NewConsumer(cfg *config.KafkaConfig, topics []string, handler MessageHandler) (*Consumer, error) {
	// Configure Sarama
	config := sarama.NewConfig()
	config.Version = sarama.V3_0_0_0
	config.Consumer.Return.Errors = true
	config.Consumer.Offsets.Initial = sarama.OffsetOldest
	config.Consumer.Group.Rebalance.GroupStrategies = []sarama.BalanceStrategy{
		sarama.NewBalanceStrategyRoundRobin(),
	}

	// Configure SASL authentication
	config.Net.SASL.Enable = true
	config.Net.SASL.Mechanism = sarama.SASLTypeSCRAMSHA512
	config.Net.SASL.User = cfg.SASLUsername
	config.Net.SASL.Password = cfg.SASLPassword
	config.Net.SASL.SCRAMClientGeneratorFunc = func() sarama.SCRAMClient {
		return &XDGSCRAMClient{HashGeneratorFcn: SHA512}
	}

	// Configure TLS
	config.Net.TLS.Enable = true
	if cfg.SkipVerify {
		config.Net.TLS.Config = &tls.Config{
			InsecureSkipVerify: true,
		}
	}

	// Create consumer group
	consumerGroup, err := sarama.NewConsumerGroup(cfg.Brokers, cfg.ConsumerGroup, config)
	if err != nil {
		return nil, fmt.Errorf("failed to create consumer group: %w", err)
	}

	return &Consumer{
		consumer: consumerGroup,
		topics:   topics,
		handler:  &consumerGroupHandler{messageHandler: handler},
	}, nil
}

// Start begins consuming messages from Kafka
func (c *Consumer) Start(ctx context.Context) error {
	for {
		// Check if context is cancelled
		if err := ctx.Err(); err != nil {
			return err
		}

		// Consume messages
		if err := c.consumer.Consume(ctx, c.topics, c.handler); err != nil {
			log.Printf("Error consuming from Kafka: %v", err)
			// Continue consuming even on error
		}
	}
}

// Close closes the Kafka consumer
func (c *Consumer) Close() error {
	return c.consumer.Close()
}

// consumerGroupHandler implements sarama.ConsumerGroupHandler
type consumerGroupHandler struct {
	messageHandler MessageHandler
}

// Setup is called when a new session is started
func (h *consumerGroupHandler) Setup(sarama.ConsumerGroupSession) error {
	return nil
}

// Cleanup is called when a session is ended
func (h *consumerGroupHandler) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

// ConsumeClaim processes messages from a partition
func (h *consumerGroupHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		// Process the message
		if err := h.messageHandler(session.Context(), message); err != nil {
			log.Printf("Error processing message: %v", err)
			// Continue processing even on error
			// In production, you might want to implement dead letter queue
		}
		
		// Mark message as processed
		session.MarkMessage(message, "")
	}
	return nil
}