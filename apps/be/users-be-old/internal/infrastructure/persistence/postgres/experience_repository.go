package postgres

import (
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"users-be/internal/domain/experience"
)

type experienceRepository struct {
	db *gorm.DB
}

func NewExperienceRepository(db *gorm.DB) experience.Repository {
	return &experienceRepository{db: db}
}

func (r *experienceRepository) Create(ctx context.Context, exp *experience.WorkExperience) error {
	return r.db.WithContext(ctx).Create(exp).Error
}

func (r *experienceRepository) GetByID(ctx context.Context, id uuid.UUID) (*experience.WorkExperience, error) {
	var exp experience.WorkExperience
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&exp).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, experience.ErrWorkExperienceNotFound
		}
		return nil, err
	}
	return &exp, nil
}

func (r *experienceRepository) GetByUserID(ctx context.Context, userID uuid.UUID) ([]*experience.WorkExperience, error) {
	var experiences []*experience.WorkExperience
	err := r.db.WithContext(ctx).
		Where("user_id = ?", userID).
		Order("is_current DESC, start_date DESC").
		Find(&experiences).Error
	return experiences, err
}

func (r *experienceRepository) Update(ctx context.Context, exp *experience.WorkExperience) error {
	return r.db.WithContext(ctx).Save(exp).Error
}

func (r *experienceRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&experience.WorkExperience{}, "id = ?", id).Error
}