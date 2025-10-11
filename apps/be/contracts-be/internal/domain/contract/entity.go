package contract

import (
	"time"
	"github.com/google/uuid"
)

type ContractStatus string
type MilestoneStatus string
type PaymentStatus string

const (
	ContractStatusActive    ContractStatus = "active"
	ContractStatusCompleted ContractStatus = "completed"
	ContractStatusCancelled ContractStatus = "cancelled"
	ContractStatusDisputed  ContractStatus = "disputed"
	
	MilestoneStatusPending   MilestoneStatus = "pending"
	MilestoneStatusInProgress MilestoneStatus = "in_progress"
	MilestoneStatusSubmitted MilestoneStatus = "submitted"
	MilestoneStatusApproved  MilestoneStatus = "approved"
	MilestoneStatusPaid      MilestoneStatus = "paid"
	
	PaymentStatusPending  PaymentStatus = "pending"
	PaymentStatusEscrowed PaymentStatus = "escrowed"
	PaymentStatusReleased PaymentStatus = "released"
	PaymentStatusRefunded PaymentStatus = "refunded"
)

type Contract struct {
	ID           uuid.UUID  `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	JobID        uuid.UUID  `gorm:"type:uuid;not null;index" json:"job_id"`
	ProposalID   uuid.UUID  `gorm:"type:uuid;not null;unique" json:"proposal_id"`
	ClientID     uuid.UUID  `gorm:"type:uuid;not null;index" json:"client_id"`
	FreelancerID uuid.UUID  `gorm:"type:uuid;not null;index" json:"freelancer_id"`
	Title        string     `gorm:"type:varchar(200);not null" json:"title"`
	Description  string     `gorm:"type:text;not null" json:"description"`
	TotalAmount  float64    `gorm:"type:decimal(12,2);not null" json:"total_amount"`
	Status       ContractStatus `gorm:"type:varchar(50);not null;default:'active'" json:"status"`
	StartDate    time.Time  `gorm:"type:date;not null" json:"start_date"`
	EndDate      *time.Time `gorm:"type:date" json:"end_date,omitempty"`
	Terms        string     `gorm:"type:text" json:"terms"`
	Milestones   []Milestone `gorm:"foreignKey:ContractID" json:"milestones,omitempty"`
	CreatedAt    time.Time  `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt    time.Time  `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
}

type Milestone struct {
	ID            uuid.UUID       `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	ContractID    uuid.UUID       `gorm:"type:uuid;not null;index" json:"contract_id"`
	Description   string          `gorm:"type:text;not null" json:"description"`
	Amount        float64         `gorm:"type:decimal(12,2);not null" json:"amount"`
	DueDate       *time.Time      `gorm:"type:date" json:"due_date,omitempty"`
	Status        MilestoneStatus `gorm:"type:varchar(50);not null;default:'pending'" json:"status"`
	PaymentStatus PaymentStatus   `gorm:"type:varchar(50);not null;default:'pending'" json:"payment_status"`
	SubmittedAt   *time.Time      `gorm:"type:timestamp" json:"submitted_at,omitempty"`
	ApprovedAt    *time.Time      `gorm:"type:timestamp" json:"approved_at,omitempty"`
	PaidAt        *time.Time      `gorm:"type:timestamp" json:"paid_at,omitempty"`
	Deliverables  *string         `gorm:"type:text" json:"deliverables,omitempty"`
	Feedback      *string         `gorm:"type:text" json:"feedback,omitempty"`
	CreatedAt     time.Time       `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt     time.Time       `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (Contract) TableName() string {
	return "contracts"
}

func (Milestone) TableName() string {
	return "contract_milestones"
}

func (c *Contract) Validate() error {
	if c.Title == "" {
		return ErrTitleRequired
	}
	if c.TotalAmount <= 0 {
		return ErrInvalidAmount
	}
	return nil
}
