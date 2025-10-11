package contract

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ContractStatus string
type MilestoneStatus string

const (
	ContractStatusActive    ContractStatus = "ACTIVE"
	ContractStatusCompleted ContractStatus = "COMPLETED"
	ContractStatusDisputed  ContractStatus = "DISPUTED"

	MilestoneStatusPending   MilestoneStatus = "PENDING"
	MilestoneStatusSubmitted MilestoneStatus = "SUBMITTED"
	MilestoneStatusApproved  MilestoneStatus = "APPROVED"
	MilestoneStatusRejected  MilestoneStatus = "REJECTED"
)

type Contract struct {
	ID           uuid.UUID      `gorm:"type:uuid;primary_key" json:"id"`
	JobID        uuid.UUID      `gorm:"type:uuid;not null;index" json:"job_id"`
	ClientID     uuid.UUID      `gorm:"type:uuid;not null;index" json:"client_id"`
	FreelancerID uuid.UUID      `gorm:"type:uuid;not null;index" json:"freelancer_id"`
	ProposalID   uuid.UUID      `gorm:"type:uuid;not null" json:"proposal_id"`
	TotalAmount  float64        `gorm:"type:decimal(15,2);not null" json:"total_amount"`
	Status       ContractStatus `gorm:"type:varchar(20);not null;default:'ACTIVE';index" json:"status"`
	StartDate    time.Time      `gorm:"not null" json:"start_date"`
	EndDate      *time.Time     `json:"end_date"`
	Milestones   []ContractMilestone `gorm:"foreignKey:ContractID" json:"milestones"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`
}

func (Contract) TableName() string {
	return "contracts"
}

func (c *Contract) BeforeCreate(tx *gorm.DB) error {
	if c.ID == uuid.Nil {
		c.ID = uuid.New()
	}
	if c.Status == "" {
		c.Status = ContractStatusActive
	}
	return nil
}

type ContractMilestone struct {
	ID          uuid.UUID       `gorm:"type:uuid;primary_key" json:"id"`
	ContractID  uuid.UUID       `gorm:"type:uuid;not null;index" json:"contract_id"`
	Description string          `gorm:"type:varchar(500);not null" json:"description"`
	Amount      float64         `gorm:"type:decimal(10,2);not null" json:"amount"`
	DueDate     time.Time       `gorm:"not null" json:"due_date"`
	Status      MilestoneStatus `gorm:"type:varchar(20);not null;default:'PENDING'" json:"status"`
	SubmittedAt *time.Time      `json:"submitted_at"`
	ApprovedAt  *time.Time      `json:"approved_at"`
	CreatedAt   time.Time       `json:"created_at"`
	UpdatedAt   time.Time       `json:"updated_at"`
	DeletedAt   gorm.DeletedAt  `gorm:"index" json:"-"`
}

func (ContractMilestone) TableName() string {
	return "contract_milestones"
}

func (cm *ContractMilestone) BeforeCreate(tx *gorm.DB) error {
	if cm.ID == uuid.Nil {
		cm.ID = uuid.New()
	}
	if cm.Status == "" {
		cm.Status = MilestoneStatusPending
	}
	return nil
}
