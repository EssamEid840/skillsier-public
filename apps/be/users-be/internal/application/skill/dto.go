package skill

import (
	"time"
	"github.com/google/uuid"
)

type CreateSkillDTO struct {
	Name       string `json:"name" binding:"required,max=100"`
	Level      string `json:"level" binding:"required,oneof=beginner intermediate advanced expert"`
	YearsOfExp *int   `json:"years_of_experience" binding:"omitempty,min=0"`
	IsPrimary  bool   `json:"is_primary"`
}

type UpdateSkillDTO struct {
	Name       *string `json:"name" binding:"omitempty,max=100"`
	Level      *string `json:"level" binding:"omitempty,oneof=beginner intermediate advanced expert"`
	YearsOfExp *int    `json:"years_of_experience" binding:"omitempty,min=0"`
	IsPrimary  *bool   `json:"is_primary"`
}

type SkillResponseDTO struct {
	ID           uuid.UUID `json:"id"`
	UserID       uuid.UUID `json:"user_id"`
	Name         string    `json:"name"`
	Level        string    `json:"level"`
	YearsOfExp   *int      `json:"years_of_experience,omitempty"`
	Endorsements int       `json:"endorsements"`
	IsPrimary    bool      `json:"is_primary"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}