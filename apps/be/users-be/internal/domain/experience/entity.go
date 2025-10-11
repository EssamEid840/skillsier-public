package experience

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type WorkExperience struct {
	ID          uuid.UUID      `gorm:"type:uuid;primary_key" json:"id"`
	UserID      uuid.UUID      `gorm:"type:uuid;not null;index" json:"user_id"`
	Title       string         `gorm:"type:varchar(200);not null" json:"title"`
	Company     string         `gorm:"type:varchar(200);not null" json:"company"`
	Location    string         `gorm:"type:varchar(200)" json:"location"`
	StartDate   time.Time      `gorm:"not null" json:"start_date"`
	EndDate     *time.Time     `json:"end_date"`
	IsCurrent   bool           `gorm:"default:false" json:"is_current"`
	Description string         `gorm:"type:text" json:"description"`
	Skills      pq.StringArray `gorm:"type:text[]" json:"skills"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}

func (WorkExperience) TableName() string {
	return "work_experiences"
}

func (w *WorkExperience) BeforeCreate(tx *gorm.DB) error {
	if w.ID == uuid.Nil {
		w.ID = uuid.New()
	}
	return nil
}

func (w *WorkExperience) Validate() error {
	if w.UserID == uuid.Nil {
		return ErrInvalidUserID
	}
	if w.Title == "" {
		return ErrTitleRequired
	}
	if w.Company == "" {
		return ErrCompanyRequired
	}
	if w.StartDate.IsZero() {
		return ErrStartDateRequired
	}
	if !w.IsCurrent && w.EndDate == nil {
		return ErrEndDateRequired
	}
	if w.EndDate != nil && w.EndDate.Before(w.StartDate) {
		return ErrInvalidDateRange
	}
	return nil
}