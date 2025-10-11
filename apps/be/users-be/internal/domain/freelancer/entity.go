package freelancer

import (
	"time"
	"github.com/google/uuid"
)

type FreelancerProfile struct {
	ID            uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	UserID        uuid.UUID `gorm:"type:uuid;not null;unique" json:"user_id"`
	Title         *string   `gorm:"type:varchar(200)" json:"title,omitempty"`
	Overview      *string   `gorm:"type:text" json:"overview,omitempty"`
	HourlyRate    *float64  `gorm:"type:decimal(10,2)" json:"hourly_rate,omitempty"`
	AvailableHours *int     `gorm:"type:int" json:"available_hours,omitempty"`
	ResponseTime  *int      `gorm:"type:int" json:"response_time,omitempty"` // in hours
	TotalJobs     int       `gorm:"type:int;default:0" json:"total_jobs"`
	TotalEarnings float64   `gorm:"type:decimal(12,2);default:0" json:"total_earnings"`
	SuccessRate   float64   `gorm:"type:decimal(5,2);default:0" json:"success_rate"`
	CreatedAt     time.Time `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt     time.Time `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (FreelancerProfile) TableName() string {
	return "freelancer_profiles"
}

func (fp *FreelancerProfile) Validate() error {
	if fp.HourlyRate != nil && *fp.HourlyRate < 0 {
		return ErrInvalidHourlyRate
	}
	if fp.AvailableHours != nil && *fp.AvailableHours < 0 {
		return ErrInvalidAvailableHours
	}
	return nil
}