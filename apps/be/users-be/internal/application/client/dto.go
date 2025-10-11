package client

import (
	"time"
	"github.com/google/uuid"
)

type UpdateClientProfileDTO struct {
	CompanyName *string `json:"company_name"`
	CompanySize *string `json:"company_size"`
	Industry    *string `json:"industry"`
}

type ClientProfileResponseDTO struct {
	ID              uuid.UUID `json:"id"`
	UserID          uuid.UUID `json:"user_id"`
	CompanyName     *string   `json:"company_name,omitempty"`
	CompanySize     *string   `json:"company_size,omitempty"`
	Industry        *string   `json:"industry,omitempty"`
	TotalSpent      float64   `json:"total_spent"`
	TotalJobsPosted int       `json:"total_jobs_posted"`
	TotalHired      int       `json:"total_hired"`
	PaymentVerified bool      `json:"payment_verified"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}