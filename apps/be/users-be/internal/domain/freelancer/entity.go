package freelancer

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Availability string

const (
	AvailabilityAvailable    Availability = "AVAILABLE"
	AvailabilityBusy         Availability = "BUSY"
	AvailabilityNotAvailable Availability = "NOT_AVAILABLE"
)

type FreelancerProfile struct {
	ID                uuid.UUID      `gorm:"type:uuid;primary_key" json:"id"`
	UserID            uuid.UUID      `gorm:"type:uuid;not null;unique;index" json:"user_id"`
	ProfessionalTitle string         `gorm:"type:varchar(200)" json:"professional_title"`
	Overview          string         `gorm:"type:text" json:"overview"`
	HourlyRate        float64        `gorm:"type:decimal(10,2)" json:"hourly_rate"`
	Availability      Availability   `gorm:"type:varchar(20);default:'AVAILABLE'" json:"availability"`
	TotalJobs         int            `gorm:"default:0" json:"total_jobs"`
	TotalEarnings     float64        `gorm:"type:decimal(15,2);default:0" json:"total_earnings"`
	SuccessRate       float64        `gorm:"type:decimal(5,2);default:0" json:"success_rate"`
	Rating            float64        `gorm:"type:decimal(3,2);default:0" json:"rating"`
	ReviewCount       int            `gorm:"default:0" json:"review_count"`
	CreatedAt         time.Time      `json:"created_at"`
	UpdatedAt         time.Time      `json:"updated_at"`
	DeletedAt         gorm.DeletedAt `gorm:"index" json:"-"`
}

func (FreelancerProfile) TableName() string {
	return "freelancer_profiles"
}

func (f *FreelancerProfile) BeforeCreate(tx *gorm.DB) error {
	if f.ID == uuid.Nil {
		f.ID = uuid.New()
	}
	if f.Availability == "" {
		f.Availability = AvailabilityAvailable
	}
	return nil
}

func (f *FreelancerProfile) Validate() error {
	if f.UserID == uuid.Nil {
		return ErrInvalidUserID
	}
	if f.HourlyRate < 0 {
		return ErrInvalidHourlyRate
	}
	if !f.IsValidAvailability() {
		return ErrInvalidAvailability
	}
	return nil
}

func (f *FreelancerProfile) IsValidAvailability() bool {
	switch f.Availability {
	case AvailabilityAvailable, AvailabilityBusy, AvailabilityNotAvailable:
		return true
	}
	return false
}