package experience

import (
	"time"

	"github.com/google/uuid"
)

// WorkExperience represents a freelancer's work experience
type WorkExperience struct {
	ID          uuid.UUID  `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	UserID      uuid.UUID  `gorm:"type:uuid;not null;index" json:"user_id"`
	Title       string     `gorm:"type:varchar(200);not null" json:"title"`
	Company     string     `gorm:"type:varchar(200);not null" json:"company"`
	Location    *string    `gorm:"type:varchar(200)" json:"location,omitempty"`
	StartDate   time.Time  `gorm:"type:date;not null" json:"start_date"`
	EndDate     *time.Time `gorm:"type:date" json:"end_date,omitempty"`
	IsCurrent   bool       `gorm:"type:boolean;default:false" json:"is_current"`
	Description *string    `gorm:"type:text" json:"description,omitempty"`
	CreatedAt   time.Time  `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time  `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
}

// TableName specifies the table name for GORM
func (WorkExperience) TableName() string {
	return "work_experiences"
}

// Validate checks if the work experience data is valid
func (we *WorkExperience) Validate() error {
	if we.Title == "" {
		return ErrTitleRequired
	}

	if we.Company == "" {
		return ErrCompanyRequired
	}

	if we.EndDate != nil && !we.IsCurrent && we.EndDate.Before(we.StartDate) {
		return ErrInvalidDateRange
	}

	if we.IsCurrent && we.EndDate != nil {
		return ErrCurrentWithEndDate
	}

	return nil
}