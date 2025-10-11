package postgres

import (
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"users-be/internal/domain/skill"
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
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, skill.ErrSkillNotFound
		}
		return nil, err
	}
	return &s, nil
}

func (r *skillRepository) GetByUserID(ctx context.Context, userID uuid.UUID) ([]*skill.Skill, error) {
	var skills []*skill.Skill
	err := r.db.WithContext(ctx).Where("user_id = ?", userID).Order("created_at DESC").Find(&skills).Error
	return skills, err
}

func (r *skillRepository) Update(ctx context.Context, s *skill.Skill) error {
	return r.db.WithContext(ctx).Save(s).Error
}

func (r *skillRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&skill.Skill{}, "id = ?", id).Error
}

func (r *skillRepository) GetByUserIDAndName(ctx context.Context, userID uuid.UUID, name string) (*skill.Skill, error) {
	var s skill.Skill
	err := r.db.WithContext(ctx).Where("user_id = ? AND name = ?", userID, name).First(&s).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, skill.ErrSkillNotFound
		}
		return nil, err
	}
	return &s, nil
}