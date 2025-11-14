package education

import (
	"time"

	"github.com/google/uuid"
)

type CreateEducationDTO struct {
	Degree       string     `json:"degree" binding:"required"`
	Institution  string     `json:"institution" binding:"required"`
	FieldOfStudy string     `json:"field_of_study" binding:"required"`
	StartDate    time.Time  `json:"start_date" binding:"required"`
	EndDate      *time.Time `json:"end_date"`
	Description  string     `json:"description"`
}

type UpdateEducationDTO struct {
	Degree       *string    `json:"degree,omitempty"`
	Institution  *string    `json:"institution,omitempty"`
	FieldOfStudy *string    `json:"field_of_study,omitempty"`
	StartDate    *time.Time `json:"start_date,omitempty"`
	EndDate      *time.Time `json:"end_date,omitempty"`
	Description  *string    `json:"description,omitempty"`
}

type EducationResponseDTO struct {
	ID           uuid.UUID  `json:"id"`
	UserID       uuid.UUID  `json:"user_id"`
	Degree       string     `json:"degree"`
	Institution  string     `json:"institution"`
	FieldOfStudy string     `json:"field_of_study"`
	StartDate    time.Time  `json:"start_date"`
	EndDate      *time.Time `json:"end_date"`
	Description  string     `json:"description"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
}

type ListEducationsResponseDTO struct {
	Educations []*EducationResponseDTO `json:"educations"`
	Total      int                     `json:"total"`
}