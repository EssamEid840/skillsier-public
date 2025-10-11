package job

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type JobStatus string
type BudgetType string

const (
	JobStatusOpen       JobStatus = "OPEN"
	JobStatusInProgress JobStatus = "IN_PROGRESS"
	JobStatusCompleted  JobStatus = "COMPLETED"
	JobStatusClosed     JobStatus = "CLOSED"

	BudgetTypeFixed  BudgetType = "FIXED"
	BudgetTypeHourly BudgetType = "HOURLY"
)

type Job struct {
	ID              uuid.UUID      `gorm:"type:uuid;primary_key" json:"id"`
	ClientID        uuid.UUID      `gorm:"type:uuid;not null;index" json:"client_id"`
	Title           string         `gorm:"type:varchar(200);not null" json:"title"`
	Description     string         `gorm:"type:text;not null" json:"description"`
	Category        string         `gorm:"type:varchar(100);not null;index" json:"category"`
	Status          JobStatus      `gorm:"type:varchar(20);not null;default:'OPEN';index" json:"status"`
	BudgetType      BudgetType     `gorm:"type:varchar(20);not null" json:"budget_type"`
	BudgetAmount    float64        `gorm:"type:decimal(15,2)" json:"budget_amount"`
	HourlyRateMin   float64        `gorm:"type:decimal(10,2)" json:"hourly_rate_min"`
	HourlyRateMax   float64        `gorm:"type:decimal(10,2)" json:"hourly_rate_max"`
	Duration        string         `gorm:"type:varchar(50)" json:"duration"`
	ExperienceLevel string         `gorm:"type:varchar(50)" json:"experience_level"`
	ProposalCount   int            `gorm:"default:0" json:"proposal_count"`
	Skills          []JobSkill     `gorm:"foreignKey:JobID" json:"skills"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"-"`
}

func (Job) TableName() string {
	return "jobs"
}

func (j *Job) BeforeCreate(tx *gorm.DB) error {
	if j.ID == uuid.Nil {
		j.ID = uuid.New()
	}
	if j.Status == "" {
		j.Status = JobStatusOpen
	}
	return nil
}

func (j *Job) Validate() error {
	if j.ClientID == uuid.Nil {
		return ErrInvalidClientID
	}
	if j.Title == "" {
		return ErrTitleRequired
	}
	if j.Description == "" {
		return ErrDescriptionRequired
	}
	if j.Category == "" {
		return ErrCategoryRequired
	}
	if j.BudgetType == "" {
		return ErrBudgetTypeRequired
	}
	return nil
}