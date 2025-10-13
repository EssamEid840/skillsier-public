// internal/domain/education/entity.go
package education

import (
    "time"
    "gorm.io/gorm"
)

type Education struct {
    ID                string         `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
    UserID            string         `gorm:"not null;type:uuid;index"`
    School            string         `gorm:"not null;type:varchar(200)"`
    SchoolLogo        string         `gorm:"type:varchar(500)"`
    Degree            string         `gorm:"not null;type:varchar(200)"`
    DegreeType        string         `gorm:"type:varchar(50)"` // associate, bachelor, master, doctorate, certificate
    Field             string         `gorm:"type:varchar(200)"`
    Grade             string         `gorm:"type:varchar(20)"`
    GPA               float64        `gorm:"type:decimal(4,2)"`
    MaxGPA            float64        `gorm:"type:decimal(4,2)"`
    StartDate         time.Time
    EndDate           *time.Time
    GraduationYear    int
    IsCurrent         bool           `gorm:"default:false"`
    Description       string         `gorm:"type:text"`
    Activities        string         `gorm:"type:text"`
    Honors            string         `gorm:"type:jsonb"` // JSON array
    Courses           string         `gorm:"type:jsonb"` // JSON array
    Location          string         `gorm:"type:varchar(200)"`
    Country           string         `gorm:"type:varchar(2)"`
    IsVerified        bool           `gorm:"default:false"`
    VerifiedBy        string         `gorm:"type:varchar(100)"`
    VerifiedAt        *time.Time
    CertificateURL    string         `gorm:"type:varchar(500)"`
    TranscriptURL     string         `gorm:"type:varchar(500)"`
    DisplayOrder      int            `gorm:"default:0"`
    CreatedAt         time.Time
    UpdatedAt         time.Time
    DeletedAt         gorm.DeletedAt `gorm:"index"`
}
