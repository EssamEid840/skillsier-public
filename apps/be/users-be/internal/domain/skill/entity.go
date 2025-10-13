// internal/domain/skill/entity.go
package skill

import (
    "time"
    "gorm.io/gorm"
)

type Skill struct {
    ID                string         `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
    UserID            string         `gorm:"not null;type:uuid;index:idx_user_skills"`
    SkillCategoryID   string         `gorm:"type:uuid"`
    SkillName         string         `gorm:"not null;type:varchar(100);index"`
    SkillSlug         string         `gorm:"type:varchar(120);index"`
    Proficiency       Proficiency    `gorm:"not null;type:varchar(20)"`
    YearsOfExperience int            `gorm:"default:0"`
    LastUsedYear      int            `gorm:"default:0"`
    IsPrimary         bool           `gorm:"default:false"`
    IsVerified        bool           `gorm:"default:false"`
    VerifiedBy        string         `gorm:"type:varchar(50)"`
    VerifiedAt        *time.Time
    TestScore         float64        `gorm:"type:decimal(5,2);default:0"`
    TestTakenAt       *time.Time
    TestProvider      string         `gorm:"type:varchar(100)"`
    CertificateURL    string         `gorm:"type:varchar(500)"`
    EndorsementCount  int            `gorm:"default:0"`
    ProjectCount      int            `gorm:"default:0"`
    DisplayOrder      int            `gorm:"default:0"`
    Tags              string         `gorm:"type:jsonb"`
    CreatedAt         time.Time
    UpdatedAt         time.Time
    DeletedAt         gorm.DeletedAt `gorm:"index"`
}

type Proficiency string

const (
    ProficiencyBeginner     Proficiency = "beginner"
    ProficiencyIntermediate Proficiency = "intermediate"
    ProficiencyAdvanced     Proficiency = "advanced"
    ProficiencyExpert       Proficiency = "expert"
)

func (s *Skill) CalculateScore() float64 {
    score := 0.0
    
    switch s.Proficiency {
    case ProficiencyBeginner:
        score = 25
    case ProficiencyIntermediate:
        score = 50
    case ProficiencyAdvanced:
        score = 75
    case ProficiencyExpert:
        score = 100
    }
    
    if s.YearsOfExperience > 0 {
        score += float64(s.YearsOfExperience) * 2
    }
    
    if s.IsVerified {
        score += 10
    }
    
    if s.TestScore > 0 {
        score += s.TestScore * 0.2
    }
    
    if s.EndorsementCount > 0 {
        score += float64(s.EndorsementCount) * 0.5
    }
    
    if score > 100 {
        score = 100
    }
    
    return score
}
