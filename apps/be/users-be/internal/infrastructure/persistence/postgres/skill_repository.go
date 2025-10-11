package postgres

import (
	"context"
	"users-be/internal/domain/skill"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type skillRepository struct {
	db *gorm.DB
}

func NewSkillRepository(db *gorm.DB) skill.Repository {
	return &skillRepository{db: db}
}

func (r *skillRepository) Create(ctx context.Context, s *skill.Skill) error {
	return r.db.WithContext(ctx).Create(s).Error
}

func (r *skillRepository) GetByID(ctx context.Context, id uuid.UUID) (*skill.Skill, error) {
	var s skill.Skill
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&s).Error
	if err == gorm.ErrRecordNotFound {
		return nil, skill.ErrSkillNotFound
	}
	return &s, err
}

func (r *skillRepository) GetByUserID(ctx context.Context, userID uuid.UUID) ([]*skill.Skill, error) {
	var skills []*skill.Skill
	err := r.db.WithContext(ctx).Where("user_id = ?", userID).
		Order("is_primary DESC, created_at DESC").Find(&skills).Error
	return skills, err
}

func (r *skillRepository) Update(ctx context.Context, s *skill.Skill) error {
	return r.db.WithContext(ctx).Model(s).Updates(s).Error
}

func (r *skillRepository) Delete(ctx context.Context, id uuid.UUID) error {
	result := r.db.WithContext(ctx).Delete(&skill.Skill{}, "id = ?", id)
	if result.RowsAffected == 0 {
		return skill.ErrSkillNotFound
	}
	return result.Error
}

func (r *skillRepository) CountByUserID(ctx context.Context, userID uuid.UUID) (int64, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&skill.Skill{}).
		Where("user_id = ?", userID).Count(&count).Error
	return count, err
}