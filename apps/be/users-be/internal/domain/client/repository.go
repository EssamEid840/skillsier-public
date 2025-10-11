package client

import (
	"context"
	"errors"
	"github.com/google/uuid"
)

var (
	ErrProfileNotFound      = errors.New("client profile not found")
	ErrInvalidCompanySize   = errors.New("invalid company size")
)

type Repository interface {
	Create(ctx context.Context, profile *ClientProfile) error
	GetByUserID(ctx context.Context, userID uuid.UUID) (*ClientProfile, error)
	Update(ctx context.Context, profile *ClientProfile) error
	UpdateStats(ctx context.Context, userID uuid.UUID, totalSpent float64, totalJobsPosted int, totalHired int) error
}