package postgres

import (
	"context"
	"reviews-be/internal/domain/review"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type reviewRepository struct {
	db *gorm.DB
}

func NewReviewRepository(db *gorm.DB) review.Repository {
	return &reviewRepository{db: db}
}

func (r *reviewRepository) Create(ctx context.Context, rev *review.Review) error {
	return r.db.WithContext(ctx).Create(rev).Error
}

func (r *reviewRepository) GetByID(ctx context.Context, id uuid.UUID) (*review.Review, error) {
	var rev review.Review
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&rev).Error
	if err == gorm.ErrRecordNotFound {
		return nil, review.ErrReviewNotFound
	}
	return &rev, err
}

func (r *reviewRepository) GetByContractID(ctx context.Context, contractID uuid.UUID) ([]*review.Review, error) {
	var reviews []*review.Review
	err := r.db.WithContext(ctx).Where("contract_id = ?", contractID).
		Order("created_at DESC").Find(&reviews).Error
	return reviews, err
}

func (r *reviewRepository) GetByRevieweeID(ctx context.Context, revieweeID uuid.UUID, limit, offset int) ([]*review.Review, int64, error) {
	var reviews []*review.Review
	var total int64

	query := r.db.WithContext(ctx).Model(&review.Review{}).Where("reviewee_id = ?", revieweeID)

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := query.Order("created_at DESC").Limit(limit).Offset(offset).Find(&reviews).Error
	return reviews, total, err
}

func (r *reviewRepository) GetByReviewerID(ctx context.Context, reviewerID uuid.UUID, limit, offset int) ([]*review.Review, int64, error) {
	var reviews []*review.Review
	var total int64

	query := r.db.WithContext(ctx).Model(&review.Review{}).Where("reviewer_id = ?", reviewerID)

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := query.Order("created_at DESC").Limit(limit).Offset(offset).Find(&reviews).Error
	return reviews, total, err
}

func (r *reviewRepository) CheckExisting(ctx context.Context, contractID uuid.UUID, reviewerID uuid.UUID) (bool, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&review.Review{}).
		Where("contract_id = ? AND reviewer_id = ?", contractID, reviewerID).
		Count(&count).Error
	return count > 0, err
}

func (r *reviewRepository) CalculateAverageRating(ctx context.Context, userID uuid.UUID) (float64, int64, error) {
	var result struct {
		AvgRating float64
		Count     int64
	}

	err := r.db.WithContext(ctx).Model(&review.Review{}).
		Select("AVG(rating) as avg_rating, COUNT(*) as count").
		Where("reviewee_id = ?", userID).
		Scan(&result).Error

	if err != nil {
		return 0, 0, err
	}

	return result.AvgRating, result.Count, nil
}
