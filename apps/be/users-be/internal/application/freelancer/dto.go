package freelancer

import (
	"time"
	"github.com/google/uuid"
)

type UpdateFreelancerProfileDTO struct {
	Title          *string  `json:"title"`
	Overview       *string  `json:"overview"`
	HourlyRate     *float64 `json:"hourly_rate"`
	AvailableHours *int     `json:"available_hours"`
	ResponseTime   *int     `json:"response_time"`
}

type FreelancerProfileResponseDTO struct {
	ID             uuid.UUID `json:"id"`
	UserID         uuid.UUID `json:"user_id"`
	Title          *string   `json:"title,omitempty"`
	Overview       *string   `json:"overview,omitempty"`
	HourlyRate     *float64  `json:"hourly_rate,omitempty"`
	AvailableHours *int      `json:"available_hours,omitempty"`
	ResponseTime   *int      `json:"response_time,omitempty"`
	TotalJobs      int       `json:"total_jobs"`
	TotalEarnings  float64   `json:"total_earnings"`
	SuccessRate    float64   `json:"success_rate"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
