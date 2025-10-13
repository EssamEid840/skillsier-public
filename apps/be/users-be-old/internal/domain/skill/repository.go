package skill

import (
	"context"
	"errors"

	"github.com/google/uuid"
)

var (
	ErrSkillNotFound            = errors.New("skill not found")
	ErrInvalidUserID            = errors.New("invalid user ID")
	ErrSkillNameRequired        = errors.New("skill name is required")
	ErrCategoryRequired         = errors.New("category is required")
	ErrLevelRequired            = errors.New("level is required")
	ErrInvalidLevel             = errors.New("invalid skill level")
	ErrInvalidYearsOfExperience = errors.New("years of experience must be non-negative")
)

type Repository interface {
	Create(ctx context.Context, skill *Skill) error
	GetByID(ctx context.Context, id uuid.UUID) (*Skill, error)
	GetByUserID(ctx context.Context, userID uuid.UUID) ([]*Skill, error)
	Update(ctx context.Context, skill *Skill) error
	Delete(ctx context.Context, id uuid.UUID) error
	GetByUserIDAndName(ctx context.Context, userID uuid.UUID, name string) (*Skill, error)
}