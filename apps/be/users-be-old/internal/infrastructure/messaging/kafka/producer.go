package kafka

import (
	"crypto/tls"
	"fmt"

	"users-be/internal/config"

	"github.com/IBM/sarama"
)

// Producer wraps a Kafka producer
type Producer struct {
	producer sarama.SyncProducer
}

// NewProducer creates a new Kafka producer
func NewProducer(cfg *config.KafkaConfig) (*Producer, error) {
	// Configure Sarama
	config := sarama.NewConfig()
	config.Version = sarama.V3_0_0_0
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true
	config.Producer.RequiredAcks = sarama.WaitForAll // Wait for all replicas
	config.Producer.Retry.Max = 5
	config.Producer.Compression = sarama.CompressionSnappy

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

	// Create producer
	producer, err := sarama.NewSyncProducer(cfg.Brokers, config)
	if err != nil {
		return nil, fmt.Errorf("failed to create producer: %w", err)
	}

	return &Producer{producer: producer}, nil
}

// Publish sends a message to a Kafka topic
func (p *Producer) Publish(topic string, key, value []byte) error {
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Key:   sarama.ByteEncoder(key),
		Value: sarama.ByteEncoder(value),
	}

	partition, offset, err := p.producer.SendMessage(msg)
	if err != nil {
		return fmt.Errorf("failed to publish message: %w", err)
	}

	// Log successful publish (you might want to use structured logging)
	fmt.Printf("Message published to topic=%s partition=%d offset=%d\n", topic, partition, offset)
	return nil
}

// Close closes the Kafka producer
func (p *Producer) Close() error {
	return p.producer.Close()
}