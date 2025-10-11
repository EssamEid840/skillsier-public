package postgres

import (
	"context"
	"users-be/internal/domain/education"
	"github.com/google/uuid"
	"gorm.io/gorm"
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
	if err == gorm.ErrRecordNotFound {
		return nil, education.ErrEducationNotFound
	}
	return &edu, err
}

func (r *educationRepository) GetByUserID(ctx context.Context, userID uuid.UUID) ([]*education.Education, error) {
	var educations []*education.Education
	err := r.db.WithContext(ctx).Where("user_id = ?", userID).
		Order("is_current DESC, start_date DESC").Find(&educations).Error
	return educations, err
}

func (r *educationRepository) Update(ctx context.Context, edu *education.Education) error {
	return r.db.WithContext(ctx).Model(edu).Updates(edu).Error
}

func (r *educationRepository) Delete(ctx context.Context, id uuid.UUID) error {
	result := r.db.WithContext(ctx).Delete(&education.Education{}, "id = ?", id)
	if result.RowsAffected == 0 {
		return education.ErrEducationNotFound
	}
	return result.Error
}

func (r *educationRepository) CountByUserID(ctx context.Context, userID uuid.UUID) (int64, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&education.Education{}).
		Where("user_id = ?", userID).Count(&count).Error
	return count, err
}