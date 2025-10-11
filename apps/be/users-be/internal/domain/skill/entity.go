package skill

import (
	"time"
	"github.com/google/uuid"
)

type Skill struct {
	ID           uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	UserID       uuid.UUID `gorm:"type:uuid;not null;index" json:"user_id"`
	Name         string    `gorm:"type:varchar(100);not null" json:"name"`
	Level        string    `gorm:"type:varchar(50);not null" json:"level"`
	YearsOfExp   *int      `gorm:"type:int" json:"years_of_experience,omitempty"`
	Endorsements int       `gorm:"type:int;default:0" json:"endorsements"`
	IsPrimary    bool      `gorm:"type:boolean;default:false" json:"is_primary"`
	CreatedAt    time.Time `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt    time.Time `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (Skill) TableName() string {
	return "skills"
}

func (s *Skill) Validate() error {
	validLevels := map[string]bool{
		"beginner": true, "intermediate": true, "advanced": true, "expert": true,
	}
	if !validLevels[s.Level] {
		return ErrInvalidSkillLevel
	}
	if s.Name == "" {
		return ErrSkillNameRequired
	}
	if s.YearsOfExp != nil && *s.YearsOfExp < 0 {
		return ErrInvalidYearsOfExp
	}
	return nil
}
