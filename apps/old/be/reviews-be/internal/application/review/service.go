package review

import (
	"context"
	"encoding/json"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"reviews-be/internal/domain/outbox"
	"reviews-be/internal/domain/review"
)

type Service struct {
	reviewRepo review.Repository
	outboxRepo outbox.Repository
	db         *gorm.DB
}

func NewService(reviewRepo review.Repository, outboxRepo outbox.Repository, db *gorm.DB) *Service {
	return &Service{reviewRepo: reviewRepo, outboxRepo: outboxRepo, db: db}
}

type CreateReviewDTO struct {
	ContractID uuid.UUID `json:"contract_id" binding:"required"`
	RevieweeID uuid.UUID `json:"reviewee_id" binding:"required"`
	Rating     int       `json:"rating" binding:"required,min=1,max=5"`
	Comment    string    `json:"comment"`
}

type ReviewResponseDTO struct {
	ID         uuid.UUID `json:"id"`
	ContractID uuid.UUID `json:"contract_id"`
	ReviewerID uuid.UUID `json:"reviewer_id"`
	RevieweeID uuid.UUID `json:"reviewee_id"`
	Rating     int       `json:"rating"`
	Comment    string    `json:"comment"`
}

func (s *Service) CreateReview(ctx context.Context, reviewerID uuid.UUID, dto *CreateReviewDTO) (*ReviewResponseDTO, error) {
	exists, err := s.reviewRepo.CheckExisting(ctx, dto.ContractID, reviewerID)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, review.ErrAlreadyReviewed
	}

	rev := &review.Review{
		ContractID: dto.ContractID,
		ReviewerID: reviewerID,
		RevieweeID: dto.RevieweeID,
		Rating:     dto.Rating,
		Comment:    dto.Comment,
	}

	if err := rev.Validate(); err != nil {
		return nil, err
	}

	err = s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := s.reviewRepo.Create(ctx, rev); err != nil {
			return err
		}
		event, _ := s.createReviewEvent("review.created", rev)
		return s.outboxRepo.Create(ctx, event)
	})

	if err != nil {
		return nil, err
	}

	return &ReviewResponseDTO{
		ID:         rev.ID,
		ContractID: rev.ContractID,
		ReviewerID: rev.ReviewerID,
		RevieweeID: rev.RevieweeID,
		Rating:     rev.Rating,
		Comment:    rev.Comment,
	}, nil
}

func (s *Service) GetReceivedReviews(ctx context.Context, userID uuid.UUID) ([]*ReviewResponseDTO, error) {
	reviews, err := s.reviewRepo.GetByRevieweeID(ctx, userID)
	if err != nil {
		return nil, err
	}
	dtos := make([]*ReviewResponseDTO, len(reviews))
	for i, r := range reviews {
		dtos[i] = &ReviewResponseDTO{
			ID:         r.ID,
			ContractID: r.ContractID,
			ReviewerID: r.ReviewerID,
			RevieweeID: r.RevieweeID,
			Rating:     r.Rating,
			Comment:    r.Comment,
		}
	}
	return dtos, nil
}

func (s *Service) createReviewEvent(eventType string, r *review.Review) (*outbox.Event, error) {
	payload := map[string]interface{}{
		"review_id":   r.ID.String(),
		"reviewee_id": r.RevieweeID.String(),
		"rating":      r.Rating,
	}
	payloadBytes, _ := json.Marshal(payload)
	metadata := map[string]interface{}{"source": "reviews-be"}
	metadataBytes, _ := json.Marshal(metadata)
	return &outbox.Event{
		AggregateID:   r.ID.String(),
		AggregateType: "review",
		EventType:     eventType,
		Payload:       payloadBytes,
		Metadata:      metadataBytes,
	}, nil
}