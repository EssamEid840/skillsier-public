package review

import (
	"context"
	"encoding/json"
	"fmt"
	"reviews-be/internal/domain/review"
	"reviews-be/internal/domain/outbox"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Service struct {
	reviewRepo review.Repository
	outboxRepo outbox.Repository
	db         *gorm.DB
}

func NewService(reviewRepo review.Repository, outboxRepo outbox.Repository, db *gorm.DB) *Service {
	return &Service{
		reviewRepo: reviewRepo,
		outboxRepo: outboxRepo,
		db:         db,
	}
}

func (s *Service) CreateReview(ctx context.Context, reviewerID uuid.UUID, dto *CreateReviewDTO) (*ReviewResponseDTO, error) {
	// Check if already reviewed
	exists, err := s.reviewRepo.CheckExisting(ctx, dto.ContractID, reviewerID)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, review.ErrAlreadyReviewed
	}

	// TODO: Verify that the contract is completed and reviewer is part of it
	// This requires calling contracts-be or having contract info in cache

	newReview := &review.Review{
		ContractID:        dto.ContractID,
		ReviewerID:        reviewerID,
		RevieweeID:        dto.RevieweeID,
		Rating:            dto.Rating,
		Comment:           dto.Comment,
		QualityOfWork:     dto.QualityOfWork,
		Communication:     dto.Communication,
		Professionalism:   dto.Professionalism,
		DeadlineAdherence: dto.DeadlineAdherence,
	}

	if err := newReview.Validate(); err != nil {
		return nil, err
	}

	// Transaction: Create review + outbox event
	err = s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := s.reviewRepo.Create(ctx, newReview); err != nil {
			return err
		}

		event, err := s.createReviewEvent("review.created", newReview)
		if err != nil {
			return err
		}

		return s.outboxRepo.Create(ctx, event)
	})

	if err != nil {
		return nil, err
	}

	return ToResponseDTO(newReview), nil
}

func (s *Service) GetReview(ctx context.Context, id uuid.UUID) (*ReviewResponseDTO, error) {
	r, err := s.reviewRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return ToResponseDTO(r), nil
}

func (s *Service) GetReceivedReviews(ctx context.Context, userID uuid.UUID, page, pageSize int) (*ListReviewsResponseDTO, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 20
	}

	offset := (page - 1) * pageSize

	reviews, total, err := s.reviewRepo.GetByRevieweeID(ctx, userID, pageSize, offset)
	if err != nil {
		return nil, err
	}

	// Calculate average rating
	avgRating, reviewCount, err := s.reviewRepo.CalculateAverageRating(ctx, userID)
	if err != nil {
		return nil, err
	}

	return ToListResponse(reviews, total, page, pageSize, avgRating, reviewCount), nil
}

func (s *Service) GetGivenReviews(ctx context.Context, userID uuid.UUID, page, pageSize int) (*ListReviewsResponseDTO, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 20
	}

	offset := (page - 1) * pageSize

	reviews, total, err := s.reviewRepo.GetByReviewerID(ctx, userID, pageSize, offset)
	if err != nil {
		return nil, err
	}

	return ToListResponse(reviews, total, page, pageSize, 0, 0), nil
}

func (s *Service) GetUserRating(ctx context.Context, userID uuid.UUID) (*UserRatingDTO, error) {
	avgRating, count, err := s.reviewRepo.CalculateAverageRating(ctx, userID)
	if err != nil {
		return nil, err
	}

	return &UserRatingDTO{
		UserID:        userID,
		AverageRating: avgRating,
		ReviewCount:   count,
	}, nil
}

func (s *Service) createReviewEvent(eventType string, r *review.Review) (*outbox.Event, error) {
	payload := map[string]interface{}{
		"review_id":   r.ID.String(),
		"contract_id": r.ContractID.String(),
		"reviewer_id": r.ReviewerID.String(),
		"reviewee_id": r.RevieweeID.String(),
		"rating":      r.Rating,
	}

	payloadBytes, _ := json.Marshal(payload)
	metadata := map[string]interface{}{"source": "reviews-be"}
	metadataBytes, _ := json.Marshal(metadata)

	return &outbox.Event{
		AggregateID:   r.RevieweeID.String(),
		AggregateType: "review",
		EventType:     eventType,
		Payload:       payloadBytes,
		Metadata:      metadataBytes,
	}, nil
}
