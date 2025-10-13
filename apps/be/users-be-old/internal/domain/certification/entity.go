package certification

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Certification struct {
	ID            uuid.UUID      `gorm:"type:uuid;primary_key" json:"id"`
	UserID        uuid.UUID      `gorm:"type:uuid;not null;index" json:"user_id"`
	Name          string         `gorm:"type:varchar(200);not null" json:"name"`
	Issuer        string         `gorm:"type:varchar(200);not null" json:"issuer"`
	IssueDate     time.Time      `gorm:"not null" json:"issue_date"`
	ExpiryDate    *time.Time     `json:"expiry_date"`
	CredentialID  string         `gorm:"type:varchar(200)" json:"credential_id"`
	CredentialURL string         `gorm:"type:varchar(500)" json:"credential_url"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"`
}

func (Certification) TableName() string {
	return "certifications"
}

func (c *Certification) BeforeCreate(tx *gorm.DB) error {
	if c.ID == uuid.Nil {
		c.ID = uuid.New()
	}
	return nil
}

func (c *Certification) Validate() error {
	if c.UserID == uuid.Nil {
		return ErrInvalidUserID
	}
	if c.Name == "" {
		return ErrNameRequired
	}
	if c.Issuer == "" {
		return ErrIssuerRequired
	}
	if c.IssueDate.IsZero() {
		return ErrIssueDateRequired
	}
	if c.ExpiryDate != nil && c.ExpiryDate.Before(c.IssueDate) {
		return ErrInvalidDateRange
	}
	return nil
}

func (c *Certification) IsExpired() bool {
	if c.ExpiryDate == nil {
		return false
	}
	return time.Now().After(*c.ExpiryDate)
}