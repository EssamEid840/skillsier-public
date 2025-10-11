package job

import (
	"context"
	"errors"
	"github.com/google/uuid"
)

var (
	ErrJobNotFound          = errors.New("job not found")
	ErrUnauthorized         = errors.New("unauthorized access")
	ErrTitleRequired        = errors.New("job title is required")
	ErrDescriptionRequired  = errors.New("job description is required")
	ErrBudgetRequired       = errors.New("budget amount is required for fixed price jobs")
	ErrHourlyRateRequired   = errors.New("hourly rate range is required for hourly jobs")
)

type Repository interface {
	Create(ctx context.Context, job *Job) error
	GetByID(ctx context.Context, id uuid.UUID) (*Job, error)
	List(ctx context.Context, filters *ListFilters, limit, offset int) ([]*Job, int64, error)
	Update(ctx context.Context, job *Job) error
	Delete(ctx context.Context, id uuid.UUID) error
	GetByClientID(ctx context.Context, clientID uuid.UUID, limit, offset int) ([]*Job, int64, error)
	IncrementProposalCount(ctx context.Context, jobID uuid.UUID) error
}

type ListFilters struct {
	Category        *string
	BudgetType      *BudgetType
	MinBudget       *float64
	MaxBudget       *float64
	ExperienceLevel *string
	Status          *JobStatus
	Skills          []string
	SearchTerm      *string
}