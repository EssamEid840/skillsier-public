package outbox

import (
	"context"

	"skillsier.dev/apps/be/users-be/internal/infrastructure/messaging/kafka"
)

// KafkaPublisher implements platform-shared outbox.MessagePublisher using Kafka
type KafkaPublisher struct {
	producer *kafka.Producer
}

// NewKafkaPublisher creates a new Kafka message publisher
func NewKafkaPublisher(producer *kafka.Producer) *KafkaPublisher {
	return &KafkaPublisher{
		producer: producer,
	}
}

// Publish publishes a message to Kafka
func (p *KafkaPublisher) Publish(ctx context.Context, topic string, key string, payload []byte) error {
	// Use PublishWithKey if key is provided, otherwise use regular Publish
	if key != "" {
		return p.producer.PublishWithKey(ctx, topic, []byte(key), payload)
	}
	return p.producer.Publish(ctx, topic, payload)
}