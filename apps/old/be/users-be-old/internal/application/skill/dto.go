package skill

import (
	"time"

	"github.com/google/uuid"

	"users-be/internal/domain/skill"
)

type CreateSkillDTO struct {
	Name              string           `json:"name" binding:"required"`
	Category          string           `json:"category" binding:"required"`
	Level             skill.SkillLevel `json:"level" binding:"required"`
	YearsOfExperience int              `json:"years_of_experience" binding:"required,min=0"`
}

type UpdateSkillDTO struct {
	Level             *skill.SkillLevel `json:"level,omitempty"`
	YearsOfExperience *int              `json:"years_of_experience,omitempty"`
}

type SkillResponseDTO struct {
	ID                uuid.UUID        `json:"id"`
	UserID            uuid.UUID        `json:"user_id"`
	Name              string           `json:"name"`
	Category          string           `json:"category"`
	Level             skill.SkillLevel `json:"level"`
	YearsOfExperience int              `json:"years_of_experience"`
	CreatedAt         time.Time        `json:"created_at"`
	UpdatedAt         time.Time        `json:"updated_at"`
}

type ListSkillsResponseDTO struct {
	Skills []*SkillResponseDTO `json:"skills"`
	Total  int                 `json:"total"`
}