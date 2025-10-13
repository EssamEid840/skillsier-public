package client

import (
	"context"
	"errors"

	"github.com/google/uuid"
)

var (
	ErrClientProfileNotFound = errors.New("client profile not found")
	ErrInvalidUserID         = errors.New("invalid user ID")
)

type Repository interface {
	Create(ctx context.Context, profile *ClientProfile) error
	GetByUserID(ctx context.Context, userID uuid.UUID) (*ClientProfile, error)
	Update(ctx context.Context, profile *ClientProfile) error
	Delete(ctx context.Context, userID uuid.UUID) error
}