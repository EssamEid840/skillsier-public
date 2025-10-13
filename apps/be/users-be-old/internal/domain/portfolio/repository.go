package portfolio

import (
	"context"
	"errors"

	"github.com/google/uuid"
)

var (
	ErrPortfolioNotFound = errors.New("portfolio not found")
	ErrInvalidUserID     = errors.New("invalid user ID")
	ErrTitleRequired     = errors.New("title is required")
	ErrImageNotFound     = errors.New("portfolio image not found")
)

type Repository interface {
	Create(ctx context.Context, portfolio *Portfolio) error
	GetByID(ctx context.Context, id uuid.UUID) (*Portfolio, error)
	GetByUserID(ctx context.Context, userID uuid.UUID) ([]*Portfolio, error)
	Update(ctx context.Context, portfolio *Portfolio) error
	Delete(ctx context.Context, id uuid.UUID) error
	
	CreateImage(ctx context.Context, image *PortfolioImage) error
	GetImagesByPortfolioID(ctx context.Context, portfolioID uuid.UUID) ([]*PortfolioImage, error)
	DeleteImage(ctx context.Context, imageID uuid.UUID) error
}