package job

import (
	"time"
	"github.com/google/uuid"
)

type SkillDTO struct {
	Name  string `json:"name" binding:"required"`
	Level string `json:"level"`
}

type CreateJobDTO struct {
	Title           string      `json:"title" binding:"required,max=200"`
	Description     string      `json:"description" binding:"required"`
	Category        string      `json:"category" binding:"required"`
	BudgetType      string      `json:"budget_type" binding:"required,oneof=fixed_price hourly"`
	BudgetAmount    *float64    `json:"budget_amount"`
	HourlyRateMin   *float64    `json:"hourly_rate_min"`
	HourlyRateMax   *float64    `json:"hourly_rate_max"`
	Duration        *string     `json:"duration"`
	ExperienceLevel string      `json:"experience_level" binding:"required"`
	RequiredSkills  []SkillDTO  `json:"required_skills"`
}

type UpdateJobDTO struct {
	Title           *string     `json:"title" binding:"omitempty,max=200"`
	Description     *string     `json:"description"`
	Status          *string     `json:"status" binding:"omitempty,oneof=open in_progress completed cancelled"`
	BudgetAmount    *float64    `json:"budget_amount"`
	HourlyRateMin   *float64    `json:"hourly_rate_min"`
	HourlyRateMax   *float64    `json:"hourly_rate_max"`
	Duration        *string     `json:"duration"`
	ExperienceLevel *string     `json:"experience_level"`
}

type JobResponseDTO struct {
	ID              uuid.UUID  `json:"id"`
	ClientID        uuid.UUID  `json:"client_id"`
	Title           string     `json:"title"`
	Description     string     `json:"description"`
	Category        string     `json:"category"`
	BudgetType      string     `json:"budget_type"`
	BudgetAmount    *float64   `json:"budget_amount,omitempty"`
	HourlyRateMin   *float64   `json:"hourly_rate_min,omitempty"`
	HourlyRateMax   *float64   `json:"hourly_rate_max,omitempty"`
	Duration        *string    `json:"duration,omitempty"`
	ExperienceLevel string     `json:"experience_level"`
	Status          string     `json:"status"`
	ProposalCount   int        `json:"proposal_count"`
	RequiredSkills  []SkillDTO `json:"required_skills,omitempty"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`
	ClosedAt        *time.Time `json:"closed_at,omitempty"`
}

type ListJobsResponseDTO struct {
	Jobs       []JobResponseDTO `json:"jobs"`
	Total      int64            `json:"total"`
	Page       int              `json:"page"`
	PageSize   int              `json:"page_size"`
	TotalPages int              `json:"total_pages"`
}
