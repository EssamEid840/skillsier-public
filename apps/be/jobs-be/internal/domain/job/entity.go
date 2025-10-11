
import (
	"time"
	"github.com/google/uuid"
)

type JobStatus string
type BudgetType string

const (
	JobStatusOpen       JobStatus = "open"
	JobStatusInProgress JobStatus = "in_progress"
	JobStatusCompleted  JobStatus = "completed"
	JobStatusCancelled  JobStatus = "cancelled"
	
	BudgetTypeFixed  BudgetType = "fixed_price"
	BudgetTypeHourly BudgetType = "hourly"
)

type Job struct {
	ID              uuid.UUID  `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	ClientID        uuid.UUID  `gorm:"type:uuid;not null;index" json:"client_id"`
	Title           string     `gorm:"type:varchar(200);not null" json:"title"`
	Description     string     `gorm:"type:text;not null" json:"description"`
	Category        string     `gorm:"type:varchar(100)" json:"category"`
	BudgetType      BudgetType `gorm:"type:varchar(50);not null" json:"budget_type"`
	BudgetAmount    *float64   `gorm:"type:decimal(12,2)" json:"budget_amount,omitempty"`
	HourlyRateMin   *float64   `gorm:"type:decimal(10,2)" json:"hourly_rate_min,omitempty"`
	HourlyRateMax   *float64   `gorm:"type:decimal(10,2)" json:"hourly_rate_max,omitempty"`
	Duration        *string    `gorm:"type:varchar(100)" json:"duration,omitempty"`
	ExperienceLevel string     `gorm:"type:varchar(50)" json:"experience_level"`
	Status          JobStatus  `gorm:"type:varchar(50);not null;default:'open'" json:"status"`
	ProposalCount   int        `gorm:"type:int;default:0" json:"proposal_count"`
	RequiredSkills  []JobSkill `gorm:"foreignKey:JobID" json:"required_skills,omitempty"`
	CreatedAt       time.Time  `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt       time.Time  `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
	ClosedAt        *time.Time `gorm:"type:timestamp" json:"closed_at,omitempty"`
}

type JobSkill struct {
	ID    uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	JobID uuid.UUID `gorm:"type:uuid;not null;index" json:"job_id"`
	Name  string    `gorm:"type:varchar(100);not null" json:"name"`
	Level string    `gorm:"type:varchar(50)" json:"level"`
}

func (Job) TableName() string {
	return "jobs"
}

func (JobSkill) TableName() string {
	return "job_skills"
}

func (j *Job) Validate() error {
	if j.Title == "" {
		return ErrTitleRequired
	}
	if j.Description == "" {
		return ErrDescriptionRequired
	}
	if j.BudgetType == BudgetTypeFixed && j.BudgetAmount == nil {
		return ErrBudgetRequired
	}
	if j.BudgetType == BudgetTypeHourly && (j.HourlyRateMin == nil || j.HourlyRateMax == nil) {
		return ErrHourlyRateRequired
	}
	return nil
}