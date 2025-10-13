// internal/domain/certification/entity.go
package certification

import (
    "time"
    "gorm.io/gorm"
)

type Certification struct {
    ID                      string             `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
    UserID                  string             `gorm:"not null;type:uuid;index"`
    Name                    string             `gorm:"not null;type:varchar(300)"`
    IssuingOrganization     string             `gorm:"not null;type:varchar(200)"`
    OrganizationLogo        string             `gorm:"type:varchar(500)"`
    IssueDate               time.Time          `gorm:"not null"`
    ExpiryDate              *time.Time
    DoesNotExpire           bool               `gorm:"default:false"`
    CredentialID            string             `gorm:"type:varchar(200)"`
    CredentialURL           string             `gorm:"type:varchar(500)"`
    CertificateURL          string             `gorm:"type:varchar(500)"`
    BadgeURL                string             `gorm:"type:varchar(500)"`
    VerificationStatus      VerificationStatus `gorm:"type:varchar(20);default:'pending'"`
    VerifiedBy              string             `gorm:"type:varchar(100)"`
    VerifiedAt              *time.Time
    RejectionReason         string             `gorm:"type:text"`
    Category                string             `gorm:"type:varchar(100)"`
    SkillsCovered           string             `gorm:"type:jsonb"` // JSON array
    Description             string             `gorm:"type:text"`
    IssuedBy                string             `gorm:"type:varchar(100)"` // manual, automatic, imported
    ImportSource            string             `gorm:"type:varchar(50)"` // linkedin, coursera, udemy, etc
    Score                   float64            `gorm:"type:decimal(5,2)"`
    MaxScore                float64            `gorm:"type:decimal(5,2)"`
    PassingScore            float64            `gorm:"type:decimal(5,2)"`
    IsPrimary               bool               `gorm:"default:false"`
    DisplayOrder            int                `gorm:"default:0"`
    ViewCount               int                `gorm:"default:0"`
    EndorsementCount        int                `gorm:"default:0"`
    CreatedAt               time.Time
    UpdatedAt               time.Time
    DeletedAt               gorm.DeletedAt     `gorm:"index"`
}

type VerificationStatus string

const (
    VerificationPending  VerificationStatus = "pending"
    VerificationVerified VerificationStatus = "verified"
    VerificationRejected VerificationStatus = "rejected"
    VerificationExpired  VerificationStatus = "expired"
)

func (c *Certification) IsExpired() bool {
    if c.DoesNotExpire || c.ExpiryDate == nil {
        return false
    }
    return c.ExpiryDate.Before(time.Now())
}

func (c *Certification) IsValid() bool {
    return c.VerificationStatus == VerificationVerified && !c.IsExpired()
}
