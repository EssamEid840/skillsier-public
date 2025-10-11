package postgres

import (
	"context"
	"users-be/internal/domain/freelancer"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type freelancerRepository struct {
	db *gorm.DB
}

func NewFreelancerRepository(db *gorm.DB) freelancer.Repository {
	return &freelancerRepository{db: db}
}

func (r *freelancerRepository) Create(ctx context.Context, profile *freelancer.FreelancerProfile) error {
	return r.db.WithContext(ctx).Create(profile).Error
}

func (r *freelancerRepository) GetByUserID(ctx context.Context, userID uuid.UUID) (*freelancer.FreelancerProfile, error) {
	var profile freelancer.FreelancerProfile
	err := r.db.WithContext(ctx).Where("user_id = ?", userID).First(&profile).Error
	if err == gorm.ErrRecordNotFound {
		return nil, freelancer.ErrProfileNotFound
	}
	return &profile, err
}

func (r *freelancerRepository) Update(ctx context.Context, profile *freelancer.FreelancerProfile) error {
	return r.db.WithContext(ctx).Model(profile).Updates(profile).Error
}

func (r *freelancerRepository) UpdateStats(ctx context.Context, userID uuid.UUID, totalJobs int, totalEarnings float64, successRate float64) error {
	return r.db.WithContext(ctx).Model(&freelancer.FreelancerProfile{}).
		Where("user_id = ?", userID).
		Updates(map[string]interface{}{
			"total_jobs":     totalJobs,
			"total_earnings": totalEarnings,
			"success_rate":   successRate,
		}).Error
}