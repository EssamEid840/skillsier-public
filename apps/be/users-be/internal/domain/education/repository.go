package education

import (
	"context"
	"errors"

	"github.com/google/uuid"
)

var (
	ErrEducationNotFound     = errors.New("education not found")
	ErrInvalidUserID         = errors.New("invalid user ID")
	ErrDegreeRequired        = errors.New("degree is required")
	ErrInstitutionRequired   = errors.New("institution is required")
	ErrFieldOfStudyRequired  = errors.New("field of study is required")
	ErrStartDateRequired     = errors.New("start date is required")
	ErrInvalidDateRange      = errors.New("end date must be after start date")
)

type Repository interface {
	Create(ctx context.Context, education *Education) error
	GetByID(ctx context.Context, id uuid.UUID) (*Education, error)
	GetByUserID(ctx context.Context, userID uuid.UUID) ([]*Education, error)
	Update(ctx context.Context, education *Education) error
	Delete(ctx context.Context, id uuid.UUID) error
}