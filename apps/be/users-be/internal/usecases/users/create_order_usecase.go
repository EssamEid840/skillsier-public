package users

import (
	"context"
	"database/sql"
	"encoding/json"
	"github.com/ericlagergren/decimal"
	"github.com/friendsofgo/errors"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/types"
	"users.go/m/internal/entities/outbox"
	"users.go/m/internal/infrastructure/event"
	"users.go/m/internal/models"
	"users.go/m/internal/repository/users"
	outbox2 "users.go/m/internal/repository/outbox"
	"users.go/m/internal/uow"
	"time"
)

type CreateUserUseCase struct {
	userRepo      *users.UsersRepo
	outboxRepo     *outbox2.OutboxRepo
	uow            *uow.UnitOfWork
	eventPublisher event.EventPublisher
}

func NewCreateUserUseCase(
	userRepo *users.UsersRepo,
	uow *uow.UnitOfWork,
	outboxRepo *outbox2.OutboxRepo,
	eventPublisher event.EventPublisher,
) *CreateUserUseCase {
	return &CreateUserUseCase{
		userRepo:      userRepo,
		uow:            uow,
		outboxRepo:     outboxRepo,
		eventPublisher: eventPublisher,
	}
}

type CreateUserInput struct {
	UserID      string  `json:"user_id" binding:"required"`
	TotalAmount float64 `json:"total_amount" binding:"required"`
}

func (uc *CreateUserUseCase) Execute(ctx context.Context, input CreateUserInput) error {
	totalAmountDecimal := new(decimal.Big).SetFloat64(input.TotalAmount)

	user := &models.User{
		ID:          uuid.New().String(),
		UserID:      input.UserID,
		TotalAmount: types.NewDecimal(totalAmountDecimal),
		Status:      models.UserStatusPending,
	}

	userPayload, err := json.Marshal(user)
	if err != nil {
		log.Err(err).Msg("failed to marshal user")

		return errors.New("failed to marshal user")
	}

	err = uc.uow.Do(ctx, func(tx *sql.Tx) (err error) {
		if err := uc.userRepo.Save(ctx, tx, user); err != nil {
			return err
		}

		event := &models.Outbox{
			ID:            uuid.New().String(),
			AggregateType: string(outbox.UserOutboxAggregateType),
			AggregateID:   user.ID,
			Type:          string(outbox.OutboxEventTypeUserCreated),
			Payload:       userPayload,
			SentAt:        null.TimeFromPtr(nil),
			OccurredAt:    null.TimeFrom(time.Now()),
		}

		return uc.outboxRepo.Save(ctx, tx, event)
	})

	if err != nil {
		log.Err(err).Msg("failed to save user")

		return errors.New("failed to save user")
	}

	return nil
}
