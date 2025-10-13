package skill

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SkillLevel string

const (
	SkillLevelBeginner     SkillLevel = "BEGINNER"
	SkillLevelIntermediate SkillLevel = "INTERMEDIATE"
	SkillLevelAdvanced     SkillLevel = "ADVANCED"
	SkillLevelExpert       SkillLevel = "EXPERT"
)

type Skill struct {
	ID                uuid.UUID      `gorm:"type:uuid;primary_key" json:"id"`
	UserID            uuid.UUID      `gorm:"type:uuid;not null;index" json:"user_id"`
	Name              string         `gorm:"type:varchar(100);not null" json:"name"`
	Category          string         `gorm:"type:varchar(50);not null" json:"category"`
	Level             SkillLevel     `gorm:"type:varchar(20);not null" json:"level"`
	YearsOfExperience int            `gorm:"not null" json:"years_of_experience"`
	CreatedAt         time.Time      `json:"created_at"`
	UpdatedAt         time.Time      `json:"updated_at"`
	DeletedAt         gorm.DeletedAt `gorm:"index" json:"-"`
}

func (Skill) TableName() string {
	return "skills"
}

func (s *Skill) BeforeCreate(tx *gorm.DB) error {
	if s.ID == uuid.Nil {
		s.ID = uuid.New()
	}
	return nil
}

func (s *Skill) Validate() error {
	if s.UserID == uuid.Nil {
		return ErrInvalidUserID
	}
	if s.Name == "" {
		return ErrSkillNameRequired
	}
	if s.Category == "" {
		return ErrCategoryRequired
	}
	if s.Level == "" {
		return ErrLevelRequired
	}
	if !s.IsValidLevel() {
		return ErrInvalidLevel
	}
	if s.YearsOfExperience < 0 {
		return ErrInvalidYearsOfExperience
	}
	return nil
}

func (s *Skill) IsValidLevel() bool {
	switch s.Level {
	case SkillLevelBeginner, SkillLevelIntermediate, SkillLevelAdvanced, SkillLevelExpert:
		return true
	}
	return false
}