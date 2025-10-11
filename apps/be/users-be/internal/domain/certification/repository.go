package certification

import (
	"context"
	"errors"
	"github.com/google/uuid"
)

var (
	ErrCertificationNotFound = errors.New("certification not found")
	ErrNameRequired          = errors.New("certification name is required")
	ErrOrganizationRequired  = errors.New("issuing organization is required")
	ErrInvalidExpiryDate     = errors.New("expiry date must be after issue date")
	ErrMaxCertificationsExceeded = errors.New("maximum number of certifications exceeded")
)

type Repository interface {
	Create(ctx context.Context, cert *Certification) error
	GetByID(ctx context.Context, id uuid.UUID) (*Certification, error)
	GetByUserID(ctx context.Context, userID uuid.UUID) ([]*Certification, error)
	Update(ctx context.Context, cert *Certification) error
	Delete(ctx context.Context, id uuid.UUID) error
	CountByUserID(ctx context.Context, userID uuid.UUID) (int64, error)
}