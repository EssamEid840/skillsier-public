// ==========================================
// FILE: shared/eventbus/consumer_factory.go
// Factory for creating event consumers with handlers
// ==========================================
package eventbus

import (
	"context"
	"log"
	"users-be/internal/infrastructure/messaging/kafka"
	"users-be/internal/config"
)

type MessageHandler func(ctx context.Context, message []byte) error

type ConsumerFactory struct {
	kafkaConfig *config.KafkaConfig
}

func NewConsumerFactory(kafkaConfig *config.KafkaConfig) *ConsumerFactory {
	return &ConsumerFactory{
		kafkaConfig: kafkaConfig,
	}
}

func (f *ConsumerFactory) CreateConsumer(topics []string, handler MessageHandler) (*kafka.Consumer, error) {
	consumer, err := kafka.NewConsumer(
		f.kafkaConfig,
		topics,
		handler,
	)
	if err != nil {
		return nil, err
	}

	return consumer, nil
}

func (f *ConsumerFactory) StartConsumer(ctx context.Context, consumer *kafka.Consumer) {
	go func() {
		if err := consumer.Start(ctx); err != nil && err != context.Canceled {
			log.Printf("Consumer error: %v", err)
		}
	}()
}