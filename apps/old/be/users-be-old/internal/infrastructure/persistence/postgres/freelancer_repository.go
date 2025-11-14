package postgres

import (
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"users-be/internal/domain/freelancer"
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
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, freelancer.ErrFreelancerProfileNotFound
		}
		return nil, err
	}
	return &profile, nil
}

func (r *freelancerRepository) Update(ctx context.Context, profile *freelancer.FreelancerProfile) error {
	return r.db.WithContext(ctx).Save(profile).Error
}

func (r *freelancerRepository) Delete(ctx context.Context, userID uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&freelancer.FreelancerProfile{}, "user_id = ?", userID).Error
}