package portfolio

import (
	"context"
	"errors"
	"github.com/google/uuid"
)

var (
	ErrPortfolioNotFound      = errors.New("portfolio not found")
	ErrTitleRequired          = errors.New("title is required")
	ErrMaxPortfolioExceeded   = errors.New("maximum number of portfolio items exceeded")
	ErrMaxImagesExceeded      = errors.New("maximum number of images per portfolio item exceeded")
)

type Repository interface {
	Create(ctx context.Context, portfolio *Portfolio) error
	GetByID(ctx context.Context, id uuid.UUID) (*Portfolio, error)
	GetByUserID(ctx context.Context, userID uuid.UUID) ([]*Portfolio, error)
	Update(ctx context.Context, portfolio *Portfolio) error
	Delete(ctx context.Context, id uuid.UUID) error
	CountByUserID(ctx context.Context, userID uuid.UUID) (int64, error)
	
	// Image operations
	AddImage(ctx context.Context, image *PortfolioImage) error
	GetImages(ctx context.Context, portfolioID uuid.UUID) ([]*PortfolioImage, error)
	CountImages(ctx context.Context, portfolioID uuid.UUID) (int64, error)
	DeleteImage(ctx context.Context, imageID uuid.UUID) error
}