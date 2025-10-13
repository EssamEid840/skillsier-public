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

type Service struct {
	certRepo   certification.Repository
	outboxRepo outbox.Repository
	db         *gorm.DB
}

func NewService(certRepo certification.Repository, outboxRepo outbox.Repository, db *gorm.DB) *Service {
	return &Service{certRepo: certRepo, outboxRepo: outboxRepo, db: db}
}

func (s *Service) CreateCertification(ctx context.Context, userID uuid.UUID, dto *CreateCertificationDTO) (*CertificationResponseDTO, error) {
	cert := &certification.Certification{
		UserID:        userID,
		Name:          dto.Name,
		Issuer:        dto.Issuer,
		IssueDate:     dto.IssueDate,
		ExpiryDate:    dto.ExpiryDate,
		CredentialID:  dto.CredentialID,
		CredentialURL: dto.CredentialURL,
	}

	if err := cert.Validate(); err != nil {
		return nil, err
	}

	err := s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := s.certRepo.Create(ctx, cert); err != nil {
			return err
		}
		event, _ := s.createCertEvent("certification.added", cert)
		return s.outboxRepo.Create(ctx, event)
	})

	if err != nil {
		return nil, err
	}
	return ToResponseDTO(cert), nil
}

func (s *Service) GetAllCertifications(ctx context.Context, userID uuid.UUID) (*ListCertificationsResponseDTO, error) {
	certs, err := s.certRepo.GetByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	return ToListResponse(certs), nil
}

func (s *Service) UpdateCertification(ctx context.Context, id, userID uuid.UUID, dto *UpdateCertificationDTO) (*CertificationResponseDTO, error) {
	cert, err := s.certRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if cert.UserID != userID {
		return nil, fmt.Errorf("unauthorized")
	}

	if dto.Name != nil {
		cert.Name = *dto.Name
	}
	if dto.Issuer != nil {
		cert.Issuer = *dto.Issuer
	}
	if dto.IssueDate != nil {
		cert.IssueDate = *dto.IssueDate
	}
	if dto.ExpiryDate != nil {
		cert.ExpiryDate = dto.ExpiryDate
	}
	if dto.CredentialID != nil {
		cert.CredentialID = *dto.CredentialID
	}
	if dto.CredentialURL != nil {
		cert.CredentialURL = *dto.CredentialURL
	}

	err = s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := s.certRepo.Update(ctx, cert); err != nil {
			return err
		}
		event, _ := s.createCertEvent("certification.updated", cert)
		return s.outboxRepo.Create(ctx, event)
	})

	if err != nil {
		return nil, err
	}
	return ToResponseDTO(cert), nil
}

func (s *Service) DeleteCertification(ctx context.Context, id, userID uuid.UUID) error {
	cert, err := s.certRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if cert.UserID != userID {
		return fmt.Errorf("unauthorized")
	}

	return s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := s.certRepo.Delete(ctx, id); err != nil {
			return err
		}
		event, _ := s.createCertEvent("certification.removed", cert)
		return s.outboxRepo.Create(ctx, event)
	})
}

func (s *Service) createCertEvent(eventType string, cert *certification.Certification) (*outbox.Event, error) {
	payload := map[string]interface{}{
		"certification_id": cert.ID.String(),
		"user_id":          cert.UserID.String(),
		"name":             cert.Name,
		"issuer":           cert.Issuer,
	}
	payloadBytes, _ := json.Marshal(payload)
	metadata := map[string]interface{}{"source": "users-be"}
	metadataBytes, _ := json.Marshal(metadata)

	return &outbox.Event{
		AggregateID:   cert.UserID.String(),
		AggregateType: "user",
		EventType:     eventType,
		Payload:       payloadBytes,
		Metadata:      metadataBytes,
	}, nil
}