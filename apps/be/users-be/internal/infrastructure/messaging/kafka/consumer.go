package kafka

import (
	"context"
	"fmt"
	"log"

	"github.com/IBM/sarama"

	"skillsier.dev/apps/be/users-be/internal/config"
)

// Consumer wraps Kafka consumer functionality
type Consumer struct {
	consumer sarama.ConsumerGroup
	config   config.KafkaConfig
}

// MessageHandler is a function that processes a Kafka message
type MessageHandler func(ctx context.Context, message *sarama.ConsumerMessage) error

// NewConsumer creates a new Kafka consumer with SASL_SSL configuration
func NewConsumer(cfg config.KafkaConfig) (*Consumer, error) {
	saramaConfig := sarama.NewConfig()

	// Consumer settings
	saramaConfig.Consumer.Return.Errors = true
	saramaConfig.Consumer.Offsets.Initial = sarama.OffsetOldest
	saramaConfig.Consumer.Group.Rebalance.Strategy = sarama.NewBalanceStrategyRoundRobin()
	
	// Session settings
	saramaConfig.Consumer.Group.Session.Timeout = cfg.ConsumerConfig.SessionTimeout
	
	// Fetch settings
	saramaConfig.Consumer.Fetch.Min = int32(cfg.ConsumerConfig.FetchMinBytes)
	saramaConfig.Consumer.MaxProcessingTime = cfg.ConsumerConfig.MaxProcessingTime

	// Auto-commit settings
	if cfg.ConsumerConfig.EnableAutoCommit {
		saramaConfig.Consumer.Offsets.AutoCommit.Enable = true
		saramaConfig.Consumer.Offsets.AutoCommit.Interval = cfg.ConsumerConfig.AutoCommitInterval
	} else {
		saramaConfig.Consumer.Offsets.AutoCommit.Enable = false
	}

	// SASL configuration
	if cfg.SecurityProtocol == "SASL_SSL" {
		saramaConfig.Net.SASL.Enable = true
		saramaConfig.Net.SASL.User = cfg.SASLUsername
		saramaConfig.Net.SASL.Password = cfg.SASLPassword

		// SCRAM mechanism
		switch cfg.SASLMechanism {
		case "SCRAM-SHA-256":
			saramaConfig.Net.SASL.Mechanism = sarama.SASLTypeSCRAMSHA256
			saramaConfig.Net.SASL.SCRAMClientGeneratorFunc = func() sarama.SCRAMClient {
				return &XDGSCRAMClient{HashGeneratorFcn: SHA256}
			}
		case "SCRAM-SHA-512":
			saramaConfig.Net.SASL.Mechanism = sarama.SASLTypeSCRAMSHA512
			saramaConfig.Net.SASL.SCRAMClientGeneratorFunc = func() sarama.SCRAMClient {
				return &XDGSCRAMClient{HashGeneratorFcn: SHA512}
			}
		case "PLAIN":
			saramaConfig.Net.SASL.Mechanism = sarama.SASLTypePlaintext
		default:
			return nil, fmt.Errorf("unsupported SASL mechanism: %s", cfg.SASLMechanism)
		}

		// TLS configuration
		saramaConfig.Net.TLS.Enable = true
		if cfg.TLSSkipVerify {
			log.Println("⚠ WARNING: TLS certificate verification is disabled")
		}
	}

	// Version
	saramaConfig.Version = sarama.V3_0_0_0

	// Create consumer group
	consumer, err := sarama.NewConsumerGroup(cfg.Brokers, cfg.ConsumerGroup, saramaConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to create Kafka consumer: %w", err)
	}

	log.Printf("✓ Kafka consumer created: group=%s, brokers=%v", cfg.ConsumerGroup, cfg.Brokers)

	return &Consumer{
		consumer: consumer,
		config:   cfg,
	}, nil
}

// Subscribe subscribes to topics and processes messages with the given handler
func (c *Consumer) Subscribe(ctx context.Context, topics []string, handler MessageHandler) error {
	log.Printf("→ Subscribing to topics: %v", topics)

	consumerHandler := &consumerGroupHandler{
		handler: handler,
	}

	for {
		select {
		case <-ctx.Done():
			log.Println("✓ Consumer stopped (context cancelled)")
			return nil
		default:
			// Consume messages
			if err := c.consumer.Consume(ctx, topics, consumerHandler); err != nil {
				log.Printf("⚠ Consumer error: %v", err)
				return fmt.Errorf("consumer error: %w", err)
			}
		}
	}
}

// Close closes the consumer
func (c *Consumer) Close() error {
	if err := c.consumer.Close(); err != nil {
		return fmt.Errorf("failed to close Kafka consumer: %w", err)
	}

	log.Println("✓ Kafka consumer closed")
	return nil
}

// consumerGroupHandler implements sarama.ConsumerGroupHandler
type consumerGroupHandler struct {
	handler MessageHandler
}

// Setup is called when the consumer group is being set up
func (h *consumerGroupHandler) Setup(sarama.ConsumerGroupSession) error {
	log.Println("→ Consumer group setup")
	return nil
}

// Cleanup is called when the consumer group is being torn down
func (h *consumerGroupHandler) Cleanup(sarama.ConsumerGroupSession) error {
	log.Println("→ Consumer group cleanup")
	return nil
}

// ConsumeClaim processes messages from a partition
func (h *consumerGroupHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		// Process message
		if err := h.handler(session.Context(), message); err != nil {
			log.Printf("⚠ Failed to process message (topic=%s, partition=%d, offset=%d): %v",
				message.Topic, message.Partition, message.Offset, err)
			// Continue processing even on error
		} else {
			log.Printf("✓ Processed message (topic=%s, partition=%d, offset=%d)",
				message.Topic, message.Partition, message.Offset)
		}

		// Mark message as consumed
		session.MarkMessage(message, "")
	}

	return nil
}