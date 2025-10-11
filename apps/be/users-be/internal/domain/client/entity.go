package client

import (
	"time"
	"github.com/google/uuid"
)

type ClientProfile struct {
	ID               uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	UserID           uuid.UUID `gorm:"type:uuid;not null;unique" json:"user_id"`
	CompanyName      *string   `gorm:"type:varchar(200)" json:"company_name,omitempty"`
	CompanySize      *string   `gorm:"type:varchar(50)" json:"company_size,omitempty"`
	Industry         *string   `gorm:"type:varchar(100)" json:"industry,omitempty"`
	TotalSpent       float64   `gorm:"type:decimal(12,2);default:0" json:"total_spent"`
	TotalJobsPosted  int       `gorm:"type:int;default:0" json:"total_jobs_posted"`
	TotalHired       int       `gorm:"type:int;default:0" json:"total_hired"`
	PaymentVerified  bool      `gorm:"type:boolean;default:false" json:"payment_verified"`
	CreatedAt        time.Time `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt        time.Time `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (ClientProfile) TableName() string {
	return "client_profiles"
}

func (cp *ClientProfile) Validate() error {
	validSizes := map[string]bool{
		"1-10": true, "11-50": true, "51-200": true, "201-500": true, "501+": true,
	}
	if cp.CompanySize != nil && !validSizes[*cp.CompanySize] {
		return ErrInvalidCompanySize
	}
	return nil
}
