package education

import (
	"context"
	"errors"
	"github.com/google/uuid"
)

var (
	ErrEducationNotFound = errors.New("education not found")
	ErrSchoolRequired    = errors.New("school name is required")
	ErrDegreeRequired    = errors.New("degree is required")
	ErrInvalidDateRange  = errors.New("end date must be after start date")
	ErrMaxEducationExceeded = errors.New("maximum number of education entries exceeded")
)

type Repository interface {
	Create(ctx context.Context, edu *Education) error
	GetByID(ctx context.Context, id uuid.UUID) (*Education, error)
	GetByUserID(ctx context.Context, userID uuid.UUID) ([]*Education, error)
	Update(ctx context.Context, edu *Education) error
	Delete(ctx context.Context, id uuid.UUID) error
	CountByUserID(ctx context.Context, userID uuid.UUID) (int64, error)
}