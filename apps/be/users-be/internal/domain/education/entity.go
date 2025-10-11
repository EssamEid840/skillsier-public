package education

import (
	"time"
	"github.com/google/uuid"
)

type Education struct {
	ID            uuid.UUID  `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	UserID        uuid.UUID  `gorm:"type:uuid;not null;index" json:"user_id"`
	School        string     `gorm:"type:varchar(200);not null" json:"school"`
	Degree        string     `gorm:"type:varchar(200);not null" json:"degree"`
	FieldOfStudy  *string    `gorm:"type:varchar(200)" json:"field_of_study,omitempty"`
	StartDate     time.Time  `gorm:"type:date;not null" json:"start_date"`
	EndDate       *time.Time `gorm:"type:date" json:"end_date,omitempty"`
	IsCurrent     bool       `gorm:"type:boolean;default:false" json:"is_current"`
	Description   *string    `gorm:"type:text" json:"description,omitempty"`
	CreatedAt     time.Time  `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt     time.Time  `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (Education) TableName() string {
	return "educations"
}

func (e *Education) Validate() error {
	if e.School == "" {
		return ErrSchoolRequired
	}
	if e.Degree == "" {
		return ErrDegreeRequired
	}
	if e.EndDate != nil && !e.IsCurrent && e.EndDate.Before(e.StartDate) {
		return ErrInvalidDateRange
	}
	return nil
}