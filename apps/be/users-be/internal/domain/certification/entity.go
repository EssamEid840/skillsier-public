package certification

import (
	"time"
	"github.com/google/uuid"
)

type Certification struct {
	ID                   uuid.UUID  `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	UserID               uuid.UUID  `gorm:"type:uuid;not null;index" json:"user_id"`
	Name                 string     `gorm:"type:varchar(200);not null" json:"name"`
	IssuingOrganization  string     `gorm:"type:varchar(200);not null" json:"issuing_organization"`
	IssueDate            time.Time  `gorm:"type:date;not null" json:"issue_date"`
	ExpiryDate           *time.Time `gorm:"type:date" json:"expiry_date,omitempty"`
	CredentialID         *string    `gorm:"type:varchar(200)" json:"credential_id,omitempty"`
	CredentialURL        *string    `gorm:"type:text" json:"credential_url,omitempty"`
	CreatedAt            time.Time  `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt            time.Time  `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (Certification) TableName() string {
	return "certifications"
}

func (c *Certification) Validate() error {
	if c.Name == "" {
		return ErrNameRequired
	}
	if c.IssuingOrganization == "" {
		return ErrOrganizationRequired
	}
	if c.ExpiryDate != nil && c.ExpiryDate.Before(c.IssueDate) {
		return ErrInvalidExpiryDate
	}
	return nil
}