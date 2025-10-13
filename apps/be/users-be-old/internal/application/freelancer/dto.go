package freelancer

import (
	"github.com/google/uuid"
	"users-be/internal/domain/freelancer"
)

type UpdateFreelancerProfileDTO struct {
	ProfessionalTitle *string                      `json:"professional_title,omitempty"`
	Overview          *string                      `json:"overview,omitempty"`
	HourlyRate        *float64                     `json:"hourly_rate,omitempty"`
	Availability      *freelancer.Availability     `json:"availability,omitempty"`
}

type FreelancerProfileResponseDTO struct {
	ID                uuid.UUID                `json:"id"`
	UserID            uuid.UUID                `json:"user_id"`
	ProfessionalTitle string                   `json:"professional_title"`
	Overview          string                   `json:"overview"`
	HourlyRate        float64                  `json:"hourly_rate"`
	Availability      freelancer.Availability  `json:"availability"`
	TotalJobs         int                      `json:"total_jobs"`
	TotalEarnings     float64                  `json:"total_earnings"`
	SuccessRate       float64                  `json:"success_rate"`
	Rating            float64                  `json:"rating"`
	ReviewCount       int                      `json:"review_count"`
}