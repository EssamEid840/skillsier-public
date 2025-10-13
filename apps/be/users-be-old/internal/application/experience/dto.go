package experience

import (
	"time"

	"github.com/google/uuid"
)

type CreateWorkExperienceDTO struct {
	Title       string    `json:"title" binding:"required"`
	Company     string    `json:"company" binding:"required"`
	Location    string    `json:"location"`
	StartDate   time.Time `json:"start_date" binding:"required"`
	EndDate     *time.Time `json:"end_date"`
	IsCurrent   bool      `json:"is_current"`
	Description string    `json:"description"`
	Skills      []string  `json:"skills"`
}

type UpdateWorkExperienceDTO struct {
	Title       *string    `json:"title,omitempty"`
	Company     *string    `json:"company,omitempty"`
	Location    *string    `json:"location,omitempty"`
	StartDate   *time.Time `json:"start_date,omitempty"`
	EndDate     *time.Time `json:"end_date,omitempty"`
	IsCurrent   *bool      `json:"is_current,omitempty"`
	Description *string    `json:"description,omitempty"`
	Skills      []string   `json:"skills,omitempty"`
}

type WorkExperienceResponseDTO struct {
	ID          uuid.UUID  `json:"id"`
	UserID      uuid.UUID  `json:"user_id"`
	Title       string     `json:"title"`
	Company     string     `json:"company"`
	Location    string     `json:"location"`
	StartDate   time.Time  `json:"start_date"`
	EndDate     *time.Time `json:"end_date"`
	IsCurrent   bool       `json:"is_current"`
	Description string     `json:"description"`
	Skills      []string   `json:"skills"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

type ListWorkExperiencesResponseDTO struct {
	Experiences []*WorkExperienceResponseDTO `json:"experiences"`
	Total       int                          `json:"total"`
}