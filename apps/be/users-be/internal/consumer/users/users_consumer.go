package users

import (
	"context"
	"encoding/json"
	"github.com/rs/zerolog/log"
	"users.go/m/internal/infrastructure/event"
	"users.go/m/internal/usecases/users"
	"time"
)

type UserCreatedEvent struct {
	ID          string     `json:"id"`
	Status      string     `json:"status"`
	UserID      string     `json:"user_id"`
	CreatedAt   *time.Time `json:"created_at"` // или `string`, если null приходит без ISO-формата
	UpdatedAt   *time.Time `json:"updated_at"`
	TotalAmount float64    `json:"total_amount,string"` // строка → float64
}

type UserConsumer struct {
	confirmUserUseCase *users.ConfirmUserUseCase
	consumer           event.EventConsumer
	workerCount        int
}

func NewUserConsumer(
	consumer event.EventConsumer,
	confirmUserUseCase *users.ConfirmUserUseCase,
) *UserConsumer {
	return &UserConsumer{
		consumer:           consumer,
		workerCount:        5,
		confirmUserUseCase: confirmUserUseCase,
	}
}

func (c *UserConsumer) Start(ctx context.Context) {
	workerChan := make(chan event.ConsumerMessage, c.workerCount)

	go func() {
		consumerMessageChan, err := c.consumer.Start(ctx)
		if err != nil {
			panic(err)
		}

		for {
			select {
			case <-ctx.Done():
				return
			case consumerMessage := <-consumerMessageChan:
				workerChan <- consumerMessage
			}
		}
	}()

	for range c.workerCount {
		go c.worker(ctx, workerChan)
	}
}

func (c *UserConsumer) worker(ctx context.Context, channel <-chan event.ConsumerMessage) {
	for {
		select {
		case consumerMessage := <-channel:
			c.handleRecord(ctx, consumerMessage)
		case <-ctx.Done():
			return
		}
	}
}

func (c *UserConsumer) handleRecord(ctx context.Context, msg event.ConsumerMessage) {
	var event UserCreatedEvent
	err := json.Unmarshal(msg.Value, &event)
	if err != nil {
		log.Err(err).Msg("failed to unmarshal event as user_created_event")
		return
	}

	err = c.confirmUserUseCase.Execute(ctx, users.ConfirmUserInput{UserID: event.ID})
	if err != nil {
		log.Err(err).Msg("failed to confirm user")
		return
	}
}
