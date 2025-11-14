package job

import (
	"time"

	"github.com/google/uuid"

	"jobs-be/internal/domain/job"
)

type JobSkillDTO struct {
	Name     string `json:"name"`
	Required bool   `json:"required"`
}

type CreateJobDTO struct {
	Title           string           `json:"title" binding:"required"`
	Description     string           `json:"description" binding:"required"`
	Category        string           `json:"category" binding:"required"`
	BudgetType      job.BudgetType   `json:"budget_type" binding:"required"`
	BudgetAmount    float64          `json:"budget_amount"`
	HourlyRateMin   float64          `json:"hourly_rate_min"`
	HourlyRateMax   float64          `json:"hourly_rate_max"`
	Duration        string           `json:"duration"`
	ExperienceLevel string           `json:"experience_level"`
	Skills          []JobSkillDTO    `json:"skills"`
}

type UpdateJobDTO struct {
	Title           *string        `json:"title,omitempty"`
	Description     *string        `json:"description,omitempty"`
	Status          *job.JobStatus `json:"status,omitempty"`
	BudgetAmount    *float64       `json:"budget_amount,omitempty"`
	HourlyRateMin   *float64       `json:"hourly_rate_min,omitempty"`
	HourlyRateMax   *float64       `json:"hourly_rate_max,omitempty"`
	Duration        *string        `json:"duration,omitempty"`
	ExperienceLevel *string        `json:"experience_level,omitempty"`
}

type JobResponseDTO struct {
	ID              uuid.UUID       `json:"id"`
	ClientID        uuid.UUID       `json:"client_id"`
	Title           string          `json:"title"`
	Description     string          `json:"description"`
	Category        string          `json:"category"`
	Status          job.JobStatus   `json:"status"`
	BudgetType      job.BudgetType  `json:"budget_type"`
	BudgetAmount    float64         `json:"budget_amount"`
	HourlyRateMin   float64         `json:"hourly_rate_min"`
	HourlyRateMax   float64         `json:"hourly_rate_max"`
	Duration        string          `json:"duration"`
	ExperienceLevel string          `json:"experience_level"`
	ProposalCount   int             `json:"proposal_count"`
	Skills          []JobSkillDTO   `json:"skills"`
	CreatedAt       time.Time       `json:"created_at"`
	UpdatedAt       time.Time       `json:"updated_at"`
}

type ListJobsResponseDTO struct {
	Jobs       []*JobResponseDTO `json:"jobs"`
	Total      int64             `json:"total"`
	Page       int               `json:"page"`
	PageSize   int               `json:"page_size"`
	TotalPages int               `json:"total_pages"`
}