package freelancer

import (
	"context"
	"errors"

	"github.com/google/uuid"
)

var (
	ErrFreelancerProfileNotFound = errors.New("freelancer profile not found")
	ErrInvalidUserID             = errors.New("invalid user ID")
	ErrInvalidHourlyRate         = errors.New("hourly rate must be non-negative")
	ErrInvalidAvailability       = errors.New("invalid availability status")
)

type Repository interface {
	Create(ctx context.Context, profile *FreelancerProfile) error
	GetByUserID(ctx context.Context, userID uuid.UUID) (*FreelancerProfile, error)
	Update(ctx context.Context, profile *FreelancerProfile) error
	Delete(ctx context.Context, userID uuid.UUID) error
}