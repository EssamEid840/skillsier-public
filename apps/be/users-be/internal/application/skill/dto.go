// internal/application/skill/dto.go
package skill

import "time"

type SkillDTO struct {
    ID                string    `json:"id"`
    UserID            string    `json:"user_id"`
    SkillName         string    `json:"skill_name"`
    Proficiency       string    `json:"proficiency"`
    YearsOfExperience int       `json:"years_of_experience"`
    LastUsedYear      int       `json:"last_used_year"`
    IsPrimary         bool      `json:"is_primary"`
    IsVerified        bool      `json:"is_verified"`
    TestScore         float64   `json:"test_score"`
    EndorsementCount  int       `json:"endorsement_count"`
    ProjectCount      int       `json:"project_count"`
    DisplayOrder      int       `json:"display_order"`
    SkillScore        float64   `json:"skill_score"`
    CreatedAt         time.Time `json:"created_at"`
    UpdatedAt         time.Time `json:"updated_at"`
}

type AddSkillDTO struct {
    UserID            string  `json:"user_id" binding:"required"`
    SkillName         string  `json:"skill_name" binding:"required,min=2,max=100"`
    Proficiency       string  `json:"proficiency" binding:"required,oneof=beginner intermediate advanced expert"`
    YearsOfExperience int     `json:"years_of_experience" binding:"min=0,max=50"`
    LastUsedYear      int     `json:"last_used_year" binding:"min=1980,max=2100"`
    IsPrimary         bool    `json:"is_primary"`
}

type UpdateSkillDTO struct {
    Proficiency       *string `json:"proficiency" binding:"omitempty,oneof=beginner intermediate advanced expert"`
    YearsOfExperience *int    `json:"years_of_experience" binding:"omitempty,min=0,max=50"`
    LastUsedYear      *int    `json:"last_used_year" binding:"omitempty,min=1980,max=2100"`
    IsPrimary         *bool   `json:"is_primary"`
    TestScore         *float64 `json:"test_score" binding:"omitempty,min=0,max=100"`
}

type ReorderSkillsDTO struct {
    SkillIDs []string `json:"skill_ids" binding:"required,min=1"`
}