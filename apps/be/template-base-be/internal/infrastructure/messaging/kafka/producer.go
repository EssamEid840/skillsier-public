package kafka

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/IBM/sarama"

	"<module>/internal/config"
)

// Producer wraps Kafka producer functionality
type Producer struct {
	producer sarama.SyncProducer
	config   config.KafkaConfig
}

// NewProducer creates a new Kafka producer with SASL_SSL configuration
func NewProducer(cfg config.KafkaConfig) (*Producer, error) {
	saramaConfig := sarama.NewConfig()

	// Producer settings
	saramaConfig.Producer.RequiredAcks = sarama.RequiredAcks(cfg.ProducerConfig.RequiredAcks)
	saramaConfig.Producer.Return.Successes = true
	saramaConfig.Producer.Return.Errors = true
	saramaConfig.Producer.Retry.Max = cfg.ProducerConfig.MaxRetries
	saramaConfig.Producer.Retry.Backoff = cfg.ProducerConfig.RetryBackoff
	saramaConfig.Producer.Idempotent = cfg.ProducerConfig.IdempotentWrites
	saramaConfig.Producer.MaxMessageBytes = cfg.ProducerConfig.MaxMessageBytes

	// Compression
	switch cfg.ProducerConfig.Compression {
	case "gzip":
		saramaConfig.Producer.Compression = sarama.CompressionGZIP
	case "snappy":
		saramaConfig.Producer.Compression = sarama.CompressionSnappy
	case "lz4":
		saramaConfig.Producer.Compression = sarama.CompressionLZ4
	case "zstd":
		saramaConfig.Producer.Compression = sarama.CompressionZSTD
	default:
		saramaConfig.Producer.Compression = sarama.CompressionNone
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
				return &SCRAMClient{HashGeneratorFcn: SHA256}
			}
		case "SCRAM-SHA-512":
			saramaConfig.Net.SASL.Mechanism = sarama.SASLTypeSCRAMSHA512
			saramaConfig.Net.SASL.SCRAMClientGeneratorFunc = func() sarama.SCRAMClient {
				return &SCRAMClient{HashGeneratorFcn: SHA512}
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
			// Note: In production, you should configure proper TLS
			// saramaConfig.Net.TLS.Config = &tls.Config{InsecureSkipVerify: true}
		}
	}

	// Metadata settings
	saramaConfig.Metadata.Retry.Max = 3
	saramaConfig.Metadata.Retry.Backoff = 250 * time.Millisecond

	// Version
	saramaConfig.Version = sarama.V3_0_0_0

	// Create producer
	producer, err := sarama.NewSyncProducer(cfg.Brokers, saramaConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to create Kafka producer: %w", err)
	}

	log.Printf("✓ Kafka producer connected: %v (mechanism: %s)", cfg.Brokers, cfg.SASLMechanism)

	return &Producer{
		producer: producer,
		config:   cfg,
	}, nil
}

// Publish publishes a message to a Kafka topic
func (p *Producer) Publish(ctx context.Context, topic string, payload []byte) error {
	message := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.ByteEncoder(payload),
		Timestamp: time.Now(),
	}

	// Add headers if needed
	message.Headers = []sarama.RecordHeader{
		{
			Key:   []byte("producer"),
			Value: []byte(p.config.TopicPrefix + "-be"),
		},
		{
			Key:   []byte("timestamp"),
			Value: []byte(time.Now().Format(time.RFC3339)),
		},
	}

	partition, offset, err := p.producer.SendMessage(message)
	if err != nil {
		return fmt.Errorf("failed to publish message to topic %s: %w", topic, err)
	}

	log.Printf("→ Published message to topic %s (partition: %d, offset: %d)", topic, partition, offset)

	return nil
}

// PublishWithKey publishes a message to a Kafka topic with a partition key
func (p *Producer) PublishWithKey(ctx context.Context, topic string, key []byte, payload []byte) error {
	message := &sarama.ProducerMessage{
		Topic:     topic,
		Key:       sarama.ByteEncoder(key),
		Value:     sarama.ByteEncoder(payload),
		Timestamp: time.Now(),
	}

	partition, offset, err := p.producer.SendMessage(message)
	if err != nil {
		return fmt.Errorf("failed to publish message to topic %s: %w", topic, err)
	}

	log.Printf("→ Published keyed message to topic %s (partition: %d, offset: %d)", topic, partition, offset)

	return nil
}

// PublishBatch publishes multiple messages in a batch
func (p *Producer) PublishBatch(ctx context.Context, topic string, payloads [][]byte) error {
	messages := make([]*sarama.ProducerMessage, len(payloads))
	
	for i, payload := range payloads {
		messages[i] = &sarama.ProducerMessage{
			Topic:     topic,
			Value:     sarama.ByteEncoder(payload),
			Timestamp: time.Now(),
		}
	}

	err := p.producer.SendMessages(messages)
	if err != nil {
		return fmt.Errorf("failed to publish batch to topic %s: %w", topic, err)
	}

	log.Printf("→ Published batch of %d messages to topic %s", len(payloads), topic)

	return nil
}

// Close closes the producer
func (p *Producer) Close() error {
	if err := p.producer.Close(); err != nil {
		return fmt.Errorf("failed to close Kafka producer: %w", err)
	}

	log.Println("✓ Kafka producer closed")
	return nil
}