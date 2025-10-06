package users

import (
	"context"
	"database/sql"
	"github.com/friendsofgo/errors"
	"github.com/rs/zerolog/log"
	"users.go/m/internal/models"
	"users.go/m/internal/repository/users"
	"users.go/m/internal/uow"
)

type ConfirmUserUseCase struct {
	userRepo *users.UsersRepo
	uow       *uow.UnitOfWork
}

func NewConfirmUserUseCase(
	userRepo *users.UsersRepo,
	uow *uow.UnitOfWork,
) *ConfirmUserUseCase {
	return &ConfirmUserUseCase{
		userRepo: userRepo,
		uow:       uow,
	}
}

type ConfirmUserInput struct {
	UserID string `json:"user_id" binding:"required"`
}

func (uc *ConfirmUserUseCase) Execute(ctx context.Context, input ConfirmUserInput) error {
	err := uc.uow.Do(ctx, func(tx *sql.Tx) (err error) {
		user, err := uc.userRepo.GetUserByID(ctx, tx, input.UserID)

		if err != nil {
			return err
		}

		if user.Status != models.UserStatusPending {
			return errors.New("user status is not pending")
		}

		user.Status = models.UserStatusConfirmed
		if err := uc.userRepo.Save(ctx, tx, user); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		log.Err(err).Msg("failed to save user")

		return errors.New("failed to save user")
	}

	return nil
}
