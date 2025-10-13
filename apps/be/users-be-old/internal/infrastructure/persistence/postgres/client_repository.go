package postgres

import (
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"users-be/internal/domain/client"
)

type clientRepository struct {
	db *gorm.DB
}

func NewClientRepository(db *gorm.DB) client.Repository {
	return &clientRepository{db: db}
}

func (r *clientRepository) Create(ctx context.Context, profile *client.ClientProfile) error {
	return r.db.WithContext(ctx).Create(profile).Error
}

func (r *clientRepository) GetByUserID(ctx context.Context, userID uuid.UUID) (*client.ClientProfile, error) {
	var profile client.ClientProfile
	err := r.db.WithContext(ctx).Where("user_id = ?", userID).First(&profile).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, client.ErrClientProfileNotFound
		}
		return nil, err
	}
	return &profile, nil
}

func (r *clientRepository) Update(ctx context.Context, profile *client.ClientProfile) error {
	return r.db.WithContext(ctx).Save(profile).Error
}

func (r *clientRepository) Delete(ctx context.Context, userID uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&client.ClientProfile{}, "user_id = ?", userID).Error
}