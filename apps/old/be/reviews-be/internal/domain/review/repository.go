package review

import (
	"context"
	"errors"

	"github.com/google/uuid"
)

var (
	ErrReviewNotFound    = errors.New("review not found")
	ErrInvalidRating     = errors.New("rating must be between 1 and 5")
	ErrInvalidUserID     = errors.New("invalid user ID")
	ErrAlreadyReviewed   = errors.New("already reviewed this contract")
)

type Repository interface {
	Create(ctx context.Context, review *Review) error
	GetByID(ctx context.Context, id uuid.UUID) (*Review, error)
	GetByRevieweeID(ctx context.Context, revieweeID uuid.UUID) ([]*Review, error)
	GetByReviewerID(ctx context.Context, reviewerID uuid.UUID) ([]*Review, error)
	GetAverageRating(ctx context.Context, userID uuid.UUID) (float64, int, error)
	CheckExisting(ctx context.Context, contractID, reviewerID uuid.UUID) (bool, error)
}
