package certification

import (
	"context"
	"errors"

	"github.com/google/uuid"
)

var (
	ErrCertificationNotFound = errors.New("certification not found")
	ErrInvalidUserID         = errors.New("invalid user ID")
	ErrNameRequired          = errors.New("name is required")
	ErrIssuerRequired        = errors.New("issuer is required")
	ErrIssueDateRequired     = errors.New("issue date is required")
	ErrInvalidDateRange      = errors.New("expiry date must be after issue date")
)

type Repository interface {
	Create(ctx context.Context, certification *Certification) error
	GetByID(ctx context.Context, id uuid.UUID) (*Certification, error)
	GetByUserID(ctx context.Context, userID uuid.UUID) ([]*Certification, error)
	Update(ctx context.Context, certification *Certification) error
	Delete(ctx context.Context, id uuid.UUID) error
}