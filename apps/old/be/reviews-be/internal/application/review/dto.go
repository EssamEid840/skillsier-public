package review

import (
	"time"
	"github.com/google/uuid"
	"math"
)

type CreateReviewDTO struct {
	ContractID        uuid.UUID `json:"contract_id" binding:"required"`
	RevieweeID        uuid.UUID `json:"reviewee_id" binding:"required"`
	Rating            int       `json:"rating" binding:"required,min=1,max=5"`
	Comment           *string   `json:"comment"`
	QualityOfWork     *int      `json:"quality_of_work" binding:"omitempty,min=1,max=5"`
	Communication     *int      `json:"communication" binding:"omitempty,min=1,max=5"`
	Professionalism   *int      `json:"professionalism" binding:"omitempty,min=1,max=5"`
	DeadlineAdherence *int      `json:"deadline_adherence" binding:"omitempty,min=1,max=5"`
}

type ReviewResponseDTO struct {
	ID                uuid.UUID  `json:"id"`
	ContractID        uuid.UUID  `json:"contract_id"`
	ReviewerID        uuid.UUID  `json:"reviewer_id"`
	RevieweeID        uuid.UUID  `json:"reviewee_id"`
	Rating            int        `json:"rating"`
	Comment           *string    `json:"comment,omitempty"`
	QualityOfWork     *int       `json:"quality_of_work,omitempty"`
	Communication     *int       `json:"communication,omitempty"`
	Professionalism   *int       `json:"professionalism,omitempty"`
	DeadlineAdherence *int       `json:"deadline_adherence,omitempty"`
	CreatedAt         time.Time  `json:"created_at"`
	UpdatedAt         time.Time  `json:"updated_at"`
}

type ListReviewsResponseDTO struct {
	Reviews       []ReviewResponseDTO `json:"reviews"`
	Total         int64               `json:"total"`
	Page          int                 `json:"page"`
	PageSize      int                 `json:"page_size"`
	TotalPages    int                 `json:"total_pages"`
	AverageRating float64             `json:"average_rating,omitempty"`
	ReviewCount   int64               `json:"review_count,omitempty"`
}

type UserRatingDTO struct {
	UserID        uuid.UUID `json:"user_id"`
	AverageRating float64   `json:"average_rating"`
	ReviewCount   int64     `json:"review_count"`
}