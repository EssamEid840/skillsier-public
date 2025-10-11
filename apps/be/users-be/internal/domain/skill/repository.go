package skill

import (
	"context"
	"errors"
	"github.com/google/uuid"
)

var (
	ErrSkillNotFound       = errors.New("skill not found")
	ErrInvalidSkillLevel   = errors.New("invalid skill level")
	ErrSkillNameRequired   = errors.New("skill name is required")
	ErrInvalidYearsOfExp   = errors.New("years of experience cannot be negative")
	ErrMaxSkillsExceeded   = errors.New("maximum number of skills exceeded")
)

type Repository interface {
	Create(ctx context.Context, skill *Skill) error
	GetByID(ctx context.Context, id uuid.UUID) (*Skill, error)
	GetByUserID(ctx context.Context, userID uuid.UUID) ([]*Skill, error)
	Update(ctx context.Context, skill *Skill) error
	Delete(ctx context.Context, id uuid.UUID) error
	CountByUserID(ctx context.Context, userID uuid.UUID) (int64, error)
}