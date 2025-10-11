package proposal

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProposalStatus string

const (
	ProposalStatusPending   ProposalStatus = "PENDING"
	ProposalStatusAccepted  ProposalStatus = "ACCEPTED"
	ProposalStatusRejected  ProposalStatus = "REJECTED"
	ProposalStatusWithdrawn ProposalStatus = "WITHDRAWN"
)

type Proposal struct {
	ID             uuid.UUID        `gorm:"type:uuid;primary_key" json:"id"`
	JobID          uuid.UUID        `gorm:"type:uuid;not null;index" json:"job_id"`
	FreelancerID   uuid.UUID        `gorm:"type:uuid;not null;index" json:"freelancer_id"`
	CoverLetter    string           `gorm:"type:text;not null" json:"cover_letter"`
	ProposedRate   float64          `gorm:"type:decimal(10,2);not null" json:"proposed_rate"`
	DeliveryTime   int              `gorm:"not null" json:"delivery_time"`
	Status         ProposalStatus   `gorm:"type:varchar(20);not null;default:'PENDING';index" json:"status"`
	Milestones     []ProposalMilestone `gorm:"foreignKey:ProposalID" json:"milestones"`
	CreatedAt      time.Time        `json:"created_at"`
	UpdatedAt      time.Time        `json:"updated_at"`
	DeletedAt      gorm.DeletedAt   `gorm:"index" json:"-"`
}

func (Proposal) TableName() string {
	return "proposals"
}

func (p *Proposal) BeforeCreate(tx *gorm.DB) error {
	if p.ID == uuid.Nil {
		p.ID = uuid.New()
	}
	if p.Status == "" {
		p.Status = ProposalStatusPending
	}
	return nil
}

type ProposalMilestone struct {
	ID          uuid.UUID      `gorm:"type:uuid;primary_key" json:"id"`
	ProposalID  uuid.UUID      `gorm:"type:uuid;not null;index" json:"proposal_id"`
	Description string         `gorm:"type:varchar(500);not null" json:"description"`
	Amount      float64        `gorm:"type:decimal(10,2);not null" json:"amount"`
	DueDate     time.Time      `gorm:"not null" json:"due_date"`
	CreatedAt   time.Time      `json:"created_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}

func (ProposalMilestone) TableName() string {
	return "proposal_milestones"
}

func (pm *ProposalMilestone) BeforeCreate(tx *gorm.DB) error {
	if pm.ID == uuid.Nil {
		pm.ID = uuid.New()
	}
	return nil
}