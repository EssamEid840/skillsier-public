package client

import "github.com/google/uuid"

type UpdateClientProfileDTO struct {
	CompanyName *string `json:"company_name,omitempty"`
	CompanySize *string `json:"company_size,omitempty"`
	Industry    *string `json:"industry,omitempty"`
	Website     *string `json:"website,omitempty"`
}

type ClientProfileResponseDTO struct {
	ID          uuid.UUID `json:"id"`
	UserID      uuid.UUID `json:"user_id"`
	CompanyName string    `json:"company_name"`
	CompanySize string    `json:"company_size"`
	Industry    string    `json:"industry"`
	Website     string    `json:"website"`
	TotalSpent  float64   `json:"total_spent"`
	JobsPosted  int       `json:"jobs_posted"`
	HireRate    float64   `json:"hire_rate"`
	Rating      float64   `json:"rating"`
	ReviewCount int       `json:"review_count"`
	IsVerified  bool      `json:"is_verified"`
}