package postgres

import (
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"users-be/internal/domain/education"
)

type educationRepository struct {
	db *gorm.DB
}

func NewEducationRepository(db *gorm.DB) education.Repository {
	return &educationRepository{db: db}
}

func (r *educationRepository) Create(ctx context.Context, edu *education.Education) error {
	return r.db.WithContext(ctx).Create(edu).Error
}

func (r *educationRepository) GetByID(ctx context.Context, id uuid.UUID) (*education.Education, error) {
	var edu education.Education
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&edu).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, education.ErrEducationNotFound
		}
		return nil, err
	}
	return &edu, nil
}

func (r *educationRepository) GetByUserID(ctx context.Context, userID uuid.UUID) ([]*education.Education, error) {
	var educations []*education.Education
	err := r.db.WithContext(ctx).
		Where("user_id = ?", userID).
		Order("start_date DESC").
		Find(&educations).Error
	return educations, err
}

func (r *educationRepository) Update(ctx context.Context, edu *education.Education) error {
	return r.db.WithContext(ctx).Save(edu).Error
}

func (r *educationRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&education.Education{}, "id = ?", id).Error
}