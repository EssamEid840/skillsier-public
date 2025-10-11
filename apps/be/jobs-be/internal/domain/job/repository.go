package job

import (
	"context"
	"errors"

	"github.com/google/uuid"
)

var (
	ErrJobNotFound          = errors.New("job not found")
	ErrInvalidClientID      = errors.New("invalid client ID")
	ErrTitleRequired        = errors.New("title is required")
	ErrDescriptionRequired  = errors.New("description is required")
	ErrCategoryRequired     = errors.New("category is required")
	ErrBudgetTypeRequired   = errors.New("budget type is required")
	ErrUnauthorized         = errors.New("unauthorized")
)

type Repository interface {
	Create(ctx context.Context, job *Job) error
	GetByID(ctx context.Context, id uuid.UUID) (*Job, error)
	GetAll(ctx context.Context, limit, offset int, filters map[string]interface{}) ([]*Job, int64, error)
	GetByClientID(ctx context.Context, clientID uuid.UUID, limit, offset int) ([]*Job, int64, error)
	Update(ctx context.Context, job *Job) error
	Delete(ctx context.Context, id uuid.UUID) error
	IncrementProposalCount(ctx context.Context, jobID uuid.UUID) error
}