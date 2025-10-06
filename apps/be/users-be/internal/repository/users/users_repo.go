package users

import (
	"context"
	"database/sql"
	"github.com/rs/zerolog/log"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"users.go/m/internal/models"
)

type UsersRepo struct {
	db *sql.DB
}

func NewUsersRepo(db *sql.DB) *UsersRepo {
	return &UsersRepo{db: db}
}

func (repo *UsersRepo) Save(ctx context.Context, exec boil.ContextExecutor, user *models.User) error {
	err := user.Upsert(ctx, exec, true, []string{"id"}, boil.Whitelist("id", "status", "total_amount"), boil.Infer())
	if err != nil {
		log.Err(err).Msg("failed to upsert user")

		return err
	}

	return nil
}

func (repo *UsersRepo) GetUserByID(ctx context.Context, exec boil.ContextExecutor, id string) (*models.User, error) {
	user, err := models.Users(
		models.UserWhere.ID.EQ(id),
	).One(ctx, exec)

	if err != nil {
		log.Err(err).Msg("failed to get user by id")

		return nil, err
	}

	return user, nil
}
