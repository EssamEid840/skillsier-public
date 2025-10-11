package postgres

import (
	"context"
	"users-be/internal/domain/client"
	"github.com/google/uuid"
	"gorm.io/gorm"
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
	if err == gorm.ErrRecordNotFound {
		return nil, client.ErrProfileNotFound
	}
	return &profile, err
}

func (r *clientRepository) Update(ctx context.Context, profile *client.ClientProfile) error {
	return r.db.WithContext(ctx).Model(profile).Updates(profile).Error
}

func (r *clientRepository) UpdateStats(ctx context.Context, userID uuid.UUID, totalSpent float64, totalJobsPosted int, totalHired int) error {
	return r.db.WithContext(ctx).Model(&client.ClientProfile{}).
		Where("user_id = ?", userID).
		Updates(map[string]interface{}{
			"total_spent":       totalSpent,
			"total_jobs_posted": totalJobsPosted,
			"total_hired":       totalHired,
		}).Error
}