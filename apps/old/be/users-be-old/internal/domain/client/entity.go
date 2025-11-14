package client

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ClientProfile struct {
	ID            uuid.UUID      `gorm:"type:uuid;primary_key" json:"id"`
	UserID        uuid.UUID      `gorm:"type:uuid;not null;unique;index" json:"user_id"`
	CompanyName   string         `gorm:"type:varchar(200)" json:"company_name"`
	CompanySize   string         `gorm:"type:varchar(50)" json:"company_size"`
	Industry      string         `gorm:"type:varchar(100)" json:"industry"`
	Website       string         `gorm:"type:varchar(500)" json:"website"`
	TotalSpent    float64        `gorm:"type:decimal(15,2);default:0" json:"total_spent"`
	JobsPosted    int            `gorm:"default:0" json:"jobs_posted"`
	HireRate      float64        `gorm:"type:decimal(5,2);default:0" json:"hire_rate"`
	Rating        float64        `gorm:"type:decimal(3,2);default:0" json:"rating"`
	ReviewCount   int            `gorm:"default:0" json:"review_count"`
	IsVerified    bool           `gorm:"default:false" json:"is_verified"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"`
}

func (ClientProfile) TableName() string {
	return "client_profiles"
}

func (c *ClientProfile) BeforeCreate(tx *gorm.DB) error {
	if c.ID == uuid.Nil {
		c.ID = uuid.New()
	}
	return nil
}

func (c *ClientProfile) Validate() error {
	if c.UserID == uuid.Nil {
		return ErrInvalidUserID
	}
	return nil
}