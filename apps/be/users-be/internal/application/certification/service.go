package certification

import (
	"context"
	"encoding/json"
	"fmt"
	"users-be/internal/domain/certification"
	"users-be/internal/domain/outbox"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

const MaxCertificationsPerUser = 15

type Service struct {
	certRepo   certification.Repository
	outboxRepo outbox.Repository
	db         *gorm.DB
}

func NewService(certRepo certification.Repository, outboxRepo outbox.Repository, db *gorm.DB) *Service {
	return &Service{certRepo: certRepo, outboxRepo: outboxRepo, db: db}
}

func (s *Service) Create(ctx context.Context, userID uuid.UUID, dto *CreateCertificationDTO) (*CertificationResponseDTO, error) {
	count, err := s.certRepo.CountByUserID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to count certifications: %w", err)
	}
	if count >= MaxCertificationsPerUser {
		return nil, certification.ErrMaxCertificationsExceeded
	}

	newCert := &certification.Certification{
		UserID:              userID,
		Name:                dto.Name,
		IssuingOrganization: dto.IssuingOrganization,
		IssueDate:           dto.IssueDate,
		ExpiryDate:          dto.ExpiryDate,
		CredentialID:        dto.CredentialID,
		CredentialURL:       dto.CredentialURL,
	}

	if err := newCert.Validate(); err != nil {
		return nil, err
	}

	err = s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := s.certRepo.Create(ctx, newCert); err != nil {
			return err
		}
		event, err := s.createEvent("certification.added", newCert)
		if err != nil {
			return err
		}
		return s.outboxRepo.Create(ctx, event)
	})
	if err != nil {
		return nil, err
	}
	return ToResponseDTO(newCert), nil
}

func (s *Service) createEvent(eventType string, cert *certification.Certification) (*outbox.Event, error) {
	payload := map[string]interface{}{
		"certification_id":      cert.ID.String(),
		"user_id":               cert.UserID.String(),
		"name":                  cert.Name,
		"issuing_organization":  cert.IssuingOrganization,
	}
	payloadBytes, _ := json.Marshal(payload)
	metadata := map[string]interface{}{"source": "users-be"}
	metadataBytes, _ := json.Marshal(metadata)

	return &outbox.Event{
		AggregateID:   cert.UserID.String(),
		AggregateType: "certification",
		EventType:     eventType,
		Payload:       payloadBytes,
		Metadata:      metadataBytes,
	}, nil
}
