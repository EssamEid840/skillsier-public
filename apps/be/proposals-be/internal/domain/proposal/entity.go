package proposal

import (
	"time"
	"github.com/google/uuid"
)

type ProposalStatus string

const (
	ProposalStatusPending   ProposalStatus = "pending"
	ProposalStatusAccepted  ProposalStatus = "accepted"
	ProposalStatusRejected  ProposalStatus = "rejected"
	ProposalStatusWithdrawn ProposalStatus = "withdrawn"
)

type Proposal struct {
	ID                uuid.UUID          `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	JobID             uuid.UUID          `gorm:"type:uuid;not null;index" json:"job_id"`
	FreelancerID      uuid.UUID          `gorm:"type:uuid;not null;index" json:"freelancer_id"`
	CoverLetter       string             `gorm:"type:text;not null" json:"cover_letter"`
	BidAmount         float64            `gorm:"type:decimal(12,2);not null" json:"bid_amount"`
	EstimatedDuration string             `gorm:"type:varchar(100);not null" json:"estimated_duration"`
	Status            ProposalStatus     `gorm:"type:varchar(50);not null;default:'pending'" json:"status"`
	Milestones        []ProposalMilestone `gorm:"foreignKey:ProposalID" json:"milestones,omitempty"`
	CreatedAt         time.Time          `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt         time.Time          `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
}

type ProposalMilestone struct {
	ID          uuid.UUID  `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	ProposalID  uuid.UUID  `gorm:"type:uuid;not null;index" json:"proposal_id"`
	Description string     `gorm:"type:text;not null" json:"description"`
	Amount      float64    `gorm:"type:decimal(12,2);not null" json:"amount"`
	DueDate     *time.Time `gorm:"type:date" json:"due_date,omitempty"`
}

func (Proposal) TableName() string {
	return "proposals"
}

func (ProposalMilestone) TableName() string {
	return "proposal_milestones"
}

func (p *Proposal) Validate() error {
	if p.CoverLetter == "" {
		return ErrCoverLetterRequired
	}
	if p.BidAmount <= 0 {
		return ErrInvalidBidAmount
	}
	if p.EstimatedDuration == "" {
		return ErrDurationRequired
	}
	return nil
}