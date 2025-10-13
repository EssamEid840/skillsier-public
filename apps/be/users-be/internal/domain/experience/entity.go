// internal/domain/experience/entity.go
package experience

import (
    "time"
    "gorm.io/gorm"
)

type WorkExperience struct {
    ID              string         `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
    UserID          string         `gorm:"not null;type:uuid;index"`
    Company         string         `gorm:"not null;type:varchar(200)"`
    CompanyLogo     string         `gorm:"type:varchar(500)"`
    Title           string         `gorm:"not null;type:varchar(200)"`
    EmploymentType  string         `gorm:"type:varchar(50)"` // full-time, part-time, contract, freelance
    Location        string         `gorm:"type:varchar(200)"`
    LocationType    string         `gorm:"type:varchar(30)"` // remote, onsite, hybrid
    Description     string         `gorm:"type:text"`
    StartDate       time.Time      `gorm:"not null"`
    EndDate         *time.Time
    IsCurrent       bool           `gorm:"default:false"`
    Duration        int            `gorm:"default:0"` // months
    Industry        string         `gorm:"type:varchar(100)"`
    Skills          string         `gorm:"type:jsonb"` // JSON array
    Achievements    string         `gorm:"type:jsonb"` // JSON array
    TeamSize        int            `gorm:"default:0"`
    ReportsTo       string         `gorm:"type:varchar(200)"`
    IsVerified      bool           `gorm:"default:false"`
    VerifiedBy      string         `gorm:"type:varchar(100)"`
    VerifiedAt      *time.Time
    DisplayOrder    int            `gorm:"default:0"`
    CreatedAt       time.Time
    UpdatedAt       time.Time
    DeletedAt       gorm.DeletedAt `gorm:"index"`
}

func (w *WorkExperience) CalculateDuration() int {
    endDate := time.Now()
    if w.EndDate != nil {
        endDate = *w.EndDate
    }
    
    years := endDate.Year() - w.StartDate.Year()
    months := int(endDate.Month() - w.StartDate.Month())
    
    totalMonths := years*12 + months
    if totalMonths < 0 {
        totalMonths = 0
    }
    
    return totalMonths
}