package proposal

import (
	"time"

	"github.com/google/uuid"

	"proposals-be/internal/domain/proposal"
)

type ProposalMilestoneDTO struct {
	Description string    `json:"description" binding:"required"`
	Amount      float64   `json:"amount" binding:"required"`
	DueDate     time.Time `json:"due_date" binding:"required"`
}

type CreateProposalDTO struct {
	JobID        uuid.UUID              `json:"job_id" binding:"required"`
	CoverLetter  string                 `json:"cover_letter" binding:"required"`
	ProposedRate float64                `json:"proposed_rate" binding:"required"`
	DeliveryTime int                    `json:"delivery_time" binding:"required"`
	Milestones   []ProposalMilestoneDTO `json:"milestones"`
}

type UpdateProposalDTO struct {
	CoverLetter  *string  `json:"cover_letter,omitempty"`
	ProposedRate *float64 `json:"proposed_rate,omitempty"`
	DeliveryTime *int     `json:"delivery_time,omitempty"`
}

type ProposalResponseDTO struct {
	ID           uuid.UUID               `json:"id"`
	JobID        uuid.UUID               `json:"job_id"`
	FreelancerID uuid.UUID               `json:"freelancer_id"`
	CoverLetter  string                  `json:"cover_letter"`
	ProposedRate float64                 `json:"proposed_rate"`
	DeliveryTime int                     `json:"delivery_time"`
	Status       proposal.ProposalStatus `json:"status"`
	Milestones   []ProposalMilestoneDTO  `json:"milestones"`
	CreatedAt    time.Time               `json:"created_at"`
	UpdatedAt    time.Time               `json:"updated_at"`
}

type ListProposalsResponseDTO struct {
	Proposals  []*ProposalResponseDTO `json:"proposals"`
	Total      int64                  `json:"total"`
	Page       int                    `json:"page"`
	PageSize   int                    `json:"page_size"`
	TotalPages int                    `json:"total_pages"`
}
