package certification

import (
	"time"
	"github.com/google/uuid"
)

type CreateCertificationDTO struct {
	Name          string     `json:"name" binding:"required"`
	Issuer        string     `json:"issuer" binding:"required"`
	IssueDate     time.Time  `json:"issue_date" binding:"required"`
	ExpiryDate    *time.Time `json:"expiry_date"`
	CredentialID  string     `json:"credential_id"`
	CredentialURL string     `json:"credential_url"`
}

type UpdateCertificationDTO struct {
	Name          *string    `json:"name,omitempty"`
	Issuer        *string    `json:"issuer,omitempty"`
	IssueDate     *time.Time `json:"issue_date,omitempty"`
	ExpiryDate    *time.Time `json:"expiry_date,omitempty"`
	CredentialID  *string    `json:"credential_id,omitempty"`
	CredentialURL *string    `json:"credential_url,omitempty"`
}

type CertificationResponseDTO struct {
	ID            uuid.UUID  `json:"id"`
	UserID        uuid.UUID  `json:"user_id"`
	Name          string     `json:"name"`
	Issuer        string     `json:"issuer"`
	IssueDate     time.Time  `json:"issue_date"`
	ExpiryDate    *time.Time `json:"expiry_date"`
	CredentialID  string     `json:"credential_id"`
	CredentialURL string     `json:"credential_url"`
	IsExpired     bool       `json:"is_expired"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
}

type ListCertificationsResponseDTO struct {
	Certifications []*CertificationResponseDTO `json:"certifications"`
	Total          int                         `json:"total"`
}