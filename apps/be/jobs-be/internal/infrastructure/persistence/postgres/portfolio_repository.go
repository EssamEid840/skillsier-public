package postgres

import (
	"context"
	"users-be/internal/domain/portfolio"
	"github.com/google/uuid"
	"gorm.io/gorm"
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
	if err == gorm.ErrRecordNotFound {
		return nil, portfolio.ErrPortfolioNotFound
	}
	return &p, err
}

func (r *portfolioRepository) GetByUserID(ctx context.Context, userID uuid.UUID) ([]*portfolio.Portfolio, error) {
	var portfolios []*portfolio.Portfolio
	err := r.db.WithContext(ctx).Preload("Images", func(db *gorm.DB) *gorm.DB {
		return db.Order("display_order ASC")
	}).Where("user_id = ?", userID).
		Order("created_at DESC").Find(&portfolios).Error
	return portfolios, err
}

func (r *portfolioRepository) Update(ctx context.Context, p *portfolio.Portfolio) error {
	return r.db.WithContext(ctx).Model(p).Updates(p).Error
}

func (r *portfolioRepository) Delete(ctx context.Context, id uuid.UUID) error {
	// This will cascade delete images due to foreign key constraint
	result := r.db.WithContext(ctx).Delete(&portfolio.Portfolio{}, "id = ?", id)
	if result.RowsAffected == 0 {
		return portfolio.ErrPortfolioNotFound
	}
	return result.Error
}

func (r *portfolioRepository) CountByUserID(ctx context.Context, userID uuid.UUID) (int64, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&portfolio.Portfolio{}).
		Where("user_id = ?", userID).Count(&count).Error
	return count, err
}

func (r *portfolioRepository) AddImage(ctx context.Context, image *portfolio.PortfolioImage) error {
	return r.db.WithContext(ctx).Create(image).Error
}

func (r *portfolioRepository) GetImages(ctx context.Context, portfolioID uuid.UUID) ([]*portfolio.PortfolioImage, error) {
	var images []*portfolio.PortfolioImage
	err := r.db.WithContext(ctx).Where("portfolio_id = ?", portfolioID).
		Order("display_order ASC").Find(&images).Error
	return images, err
}

func (r *portfolioRepository) CountImages(ctx context.Context, portfolioID uuid.UUID) (int64, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&portfolio.PortfolioImage{}).
		Where("portfolio_id = ?", portfolioID).Count(&count).Error
	return count, err
}

func (r *portfolioRepository) DeleteImage(ctx context.Context, imageID uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&portfolio.PortfolioImage{}, "id = ?", imageID).Error
}
