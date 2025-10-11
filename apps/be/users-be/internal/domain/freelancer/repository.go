package freelancer

import (
	"context"
	"errors"
	"github.com/google/uuid"
)

var (
	ErrProfileNotFound        = errors.New("freelancer profile not found")
	ErrInvalidHourlyRate      = errors.New("hourly rate cannot be negative")
	ErrInvalidAvailableHours  = errors.New("available hours cannot be negative")
)

type Repository interface {
	Create(ctx context.Context, profile *FreelancerProfile) error
	GetByUserID(ctx context.Context, userID uuid.UUID) (*FreelancerProfile, error)
	Update(ctx context.Context, profile *FreelancerProfile) error
	UpdateStats(ctx context.Context, userID uuid.UUID, totalJobs int, totalEarnings float64, successRate float64) error
}