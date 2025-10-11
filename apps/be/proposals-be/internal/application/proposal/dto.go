package proposal

import (
	"time"
	"github.com/google/uuid"
)

type MilestoneDTO struct {
	Description string     `json:"description" binding:"required"`
	Amount      float64    `json:"amount" binding:"required,gt=0"`
	DueDate     *time.Time `json:"due_date"`
}

type CreateProposalDTO struct {
	JobID             uuid.UUID      `json:"job_id" binding:"required"`
	CoverLetter       string         `json:"cover_letter" binding:"required"`
	BidAmount         float64        `json:"bid_amount" binding:"required,gt=0"`
	EstimatedDuration string         `json:"estimated_duration" binding:"required"`
	Milestones        []MilestoneDTO `json:"milestones"`
}

type UpdateProposalDTO struct {
	CoverLetter       *string  `json:"cover_letter"`
	BidAmount         *float64 `json:"bid_amount" binding:"omitempty,gt=0"`
	EstimatedDuration *string  `json:"estimated_duration"`
}

type ProposalResponseDTO struct {
	ID                uuid.UUID      `json:"id"`
	JobID             uuid.UUID      `json:"job_id"`
	FreelancerID      uuid.UUID      `json:"freelancer_id"`
	CoverLetter       string         `json:"cover_letter"`
	BidAmount         float64        `json:"bid_amount"`
	EstimatedDuration string         `json:"estimated_duration"`
	Status            string         `json:"status"`
	Milestones        []MilestoneDTO `json:"milestones,omitempty"`
	CreatedAt         time.Time      `json:"created_at"`
	UpdatedAt         time.Time      `json:"updated_at"`
}

type ListProposalsResponseDTO struct {
	Proposals  []ProposalResponseDTO `json:"proposals"`
	Total      int64                 `json:"total"`
	Page       int                   `json:"page"`
	PageSize   int                   `json:"page_size"`
	TotalPages int                   `json:"total_pages"`
}