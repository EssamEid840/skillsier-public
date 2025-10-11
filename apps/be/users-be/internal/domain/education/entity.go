package education

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Education struct {
	ID           uuid.UUID      `gorm:"type:uuid;primary_key" json:"id"`
	UserID       uuid.UUID      `gorm:"type:uuid;not null;index" json:"user_id"`
	Degree       string         `gorm:"type:varchar(200);not null" json:"degree"`
	Institution  string         `gorm:"type:varchar(200);not null" json:"institution"`
	FieldOfStudy string         `gorm:"type:varchar(200);not null" json:"field_of_study"`
	StartDate    time.Time      `gorm:"not null" json:"start_date"`
	EndDate      *time.Time     `json:"end_date"`
	Description  string         `gorm:"type:text" json:"description"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`
}

func (Education) TableName() string {
	return "educations"
}

func (e *Education) BeforeCreate(tx *gorm.DB) error {
	if e.ID == uuid.Nil {
		e.ID = uuid.New()
	}
	return nil
}

func (e *Education) Validate() error {
	if e.UserID == uuid.Nil {
		return ErrInvalidUserID
	}
	if e.Degree == "" {
		return ErrDegreeRequired
	}
	if e.Institution == "" {
		return ErrInstitutionRequired
	}
	if e.FieldOfStudy == "" {
		return ErrFieldOfStudyRequired
	}
	if e.StartDate.IsZero() {
		return ErrStartDateRequired
	}
	if e.EndDate != nil && e.EndDate.Before(e.StartDate) {
		return ErrInvalidDateRange
	}
	return nil
}