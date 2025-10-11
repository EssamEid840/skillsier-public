package postgres

import (
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"users-be/internal/domain/portfolio"
)

type portfolioRepository struct {
	db *gorm.DB
}

func NewPortfolioRepository(db *gorm.DB) portfolio.Repository {
	return &portfolioRepository{db: db}
}

func (r *portfolioRepository) Create(ctx context.Context, p *portfolio.Portfolio) error {
	return r.db.WithContext(ctx).Create(p).Error
}

func (r *portfolioRepository) GetByID(ctx context.Context, id uuid.UUID) (*portfolio.Portfolio, error) {
	var p portfolio.Portfolio
	err := r.db.WithContext(ctx).Preload("Images").Where("id = ?", id).First(&p).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, portfolio.ErrPortfolioNotFound
		}
		return nil, err
	}
	return &p, nil
}

func (r *portfolioRepository) GetByUserID(ctx context.Context, userID uuid.UUID) ([]*portfolio.Portfolio, error) {
	var portfolios []*portfolio.Portfolio
	err := r.db.WithContext(ctx).
		Preload("Images").
		Where("user_id = ?", userID).
		Order("created_at DESC").
		Find(&portfolios).Error
	return portfolios, err
}

func (r *portfolioRepository) Update(ctx context.Context, p *portfolio.Portfolio) error {
	return r.db.WithContext(ctx).Save(p).Error
}

func (r *portfolioRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&portfolio.Portfolio{}, "id = ?", id).Error
}

func (r *portfolioRepository) CreateImage(ctx context.Context, image *portfolio.PortfolioImage) error {
	return r.db.WithContext(ctx).Create(image).Error
}

func (r *portfolioRepository) GetImagesByPortfolioID(ctx context.Context, portfolioID uuid.UUID) ([]*portfolio.PortfolioImage, error) {
	var images []*portfolio.PortfolioImage
	err := r.db.WithContext(ctx).
		Where("portfolio_id = ?", portfolioID).
		Order("display_order ASC").
		Find(&images).Error
	return images, err
}

func (r *portfolioRepository) DeleteImage(ctx context.Context, imageID uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&portfolio.PortfolioImage{}, "id = ?", imageID).Error
}