package review

import (
	"context"
	"errors"
	"github.com/google/uuid"
)

var (
	ErrReviewNotFound         = errors.New("review not found")
	ErrInvalidRating          = errors.New("rating must be between 1 and 5")
	ErrInvalidDetailedRating  = errors.New("detailed ratings must be between 1 and 5")
	ErrAlreadyReviewed        = errors.New("you have already reviewed this contract")
	ErrUnauthorized           = errors.New("unauthorized to create review")
)

type Repository interface {
	Create(ctx context.Context, review *Review) error
	GetByID(ctx context.Context, id uuid.UUID) (*Review, error)
	GetByContractID(ctx context.Context, contractID uuid.UUID) ([]*Review, error)
	GetByRevieweeID(ctx context.Context, revieweeID uuid.UUID, limit, offset int) ([]*Review, int64, error)
	GetByReviewerID(ctx context.Context, reviewerID uuid.UUID, limit, offset int) ([]*Review, int64, error)
	CheckExisting(ctx context.Context, contractID uuid.UUID, reviewerID uuid.UUID) (bool, error)
	CalculateAverageRating(ctx context.Context, userID uuid.UUID) (float64, int64, error)
}