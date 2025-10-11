package contract

import (
	"time"
	"github.com/google/uuid"
)

type MilestoneDTO struct {
	Description string     `json:"description" binding:"required"`
	Amount      float64    `json:"amount" binding:"required,gt=0"`
	DueDate     *time.Time `json:"due_date"`
}

type CreateContractDTO struct {
	JobID        uuid.UUID      `json:"job_id" binding:"required"`
	ProposalID   uuid.UUID      `json:"proposal_id" binding:"required"`
	ClientID     uuid.UUID      `json:"client_id" binding:"required"`
	FreelancerID uuid.UUID      `json:"freelancer_id" binding:"required"`
	Title        string         `json:"title" binding:"required"`
	Description  string         `json:"description" binding:"required"`
	TotalAmount  float64        `json:"total_amount" binding:"required,gt=0"`
	Terms        string         `json:"terms" binding:"required"`
	Milestones   []MilestoneDTO `json:"milestones" binding:"required,min=1"`
}

type SubmitMilestoneDTO struct {
	Deliverables string   `json:"deliverables" binding:"required"`
	Attachments  []string `json:"attachments"`
}

type ApproveMilestoneDTO struct {
	Feedback *string `json:"feedback"`
}

type MilestoneResponseDTO struct {
	ID            uuid.UUID  `json:"id"`
	ContractID    uuid.UUID  `json:"contract_id"`
	Description   string     `json:"description"`
	Amount        float64    `json:"amount"`
	DueDate       *time.Time `json:"due_date,omitempty"`
	Status        string     `json:"status"`
	PaymentStatus string     `json:"payment_status"`
	SubmittedAt   *time.Time `json:"submitted_at,omitempty"`
	ApprovedAt    *time.Time `json:"approved_at,omitempty"`
	PaidAt        *time.Time `json:"paid_at,omitempty"`
	Deliverables  *string    `json:"deliverables,omitempty"`
	Feedback      *string    `json:"feedback,omitempty"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
}

type ContractResponseDTO struct {
	ID           uuid.UUID              `json:"id"`
	JobID        uuid.UUID              `json:"job_id"`
	ProposalID   uuid.UUID              `json:"proposal_id"`
	ClientID     uuid.UUID              `json:"client_id"`
	FreelancerID uuid.UUID              `json:"freelancer_id"`
	Title        string                 `json:"title"`
	Description  string                 `json:"description"`
	TotalAmount  float64                `json:"total_amount"`
	Status       string                 `json:"status"`
	StartDate    time.Time              `json:"start_date"`
	EndDate      *time.Time             `json:"end_date,omitempty"`
	Terms        string                 `json:"terms"`
	Milestones   []MilestoneResponseDTO `json:"milestones"`
	CreatedAt    time.Time              `json:"created_at"`
	UpdatedAt    time.Time              `json:"updated_at"`
}

type ListContractsResponseDTO struct {
	Contracts  []ContractResponseDTO `json:"contracts"`
	Total      int64                 `json:"total"`
	Page       int                   `json:"page"`
	PageSize   int                   `json:"page_size"`
	TotalPages int                   `json:"total_pages"`
}