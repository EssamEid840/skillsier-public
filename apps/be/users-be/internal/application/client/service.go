package client

import (
	"context"
	"encoding/json"
	"users-be/internal/domain/client"
	"users-be/internal/domain/outbox"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Service struct {
	clientRepo client.Repository
	outboxRepo outbox.Repository
	db         *gorm.DB
}

func NewService(clientRepo client.Repository, outboxRepo outbox.Repository, db *gorm.DB) *Service {
	return &Service{
		clientRepo: clientRepo,
		outboxRepo: outboxRepo,
		db:         db,
	}
}

func (s *Service) GetProfile(ctx context.Context, userID uuid.UUID) (*ClientProfileResponseDTO, error) {
	profile, err := s.clientRepo.GetByUserID(ctx, userID)
	if err != nil {
		if err == client.ErrProfileNotFound {
			// Auto-create profile
			newProfile := &client.ClientProfile{
				UserID: userID,
			}
			if err := s.clientRepo.Create(ctx, newProfile); err != nil {
				return nil, err
			}
			return ToResponseDTO(newProfile), nil
		}
		return nil, err
	}
	return ToResponseDTO(profile), nil
}

func (s *Service) UpdateProfile(ctx context.Context, userID uuid.UUID, dto *UpdateClientProfileDTO) (*ClientProfileResponseDTO, error) {
	profile, err := s.clientRepo.GetByUserID(ctx, userID)
	if err != nil {
		if err == client.ErrProfileNotFound {
			profile = &client.ClientProfile{UserID: userID}
			if err := s.clientRepo.Create(ctx, profile); err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}

	if dto.CompanyName != nil {
		profile.CompanyName = dto.CompanyName
	}
	if dto.CompanySize != nil {
		profile.CompanySize = dto.CompanySize
	}
	if dto.Industry != nil {
		profile.Industry = dto.Industry
	}

	if err := profile.Validate(); err != nil {
		return nil, err
	}

	err = s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := s.clientRepo.Update(ctx, profile); err != nil {
			return err
		}

		event, err := s.createEvent("client.profile.updated", profile)
		if err != nil {
			return err
		}

		return s.outboxRepo.Create(ctx, event)
	})

	if err != nil {
		return nil, err
	}

	return ToResponseDTO(profile), nil
}

func (s *Service) createEvent(eventType string, profile *client.ClientProfile) (*outbox.Event, error) {
	payload := map[string]interface{}{
		"user_id":           profile.UserID.String(),
		"company_name":      profile.CompanyName,
		"total_jobs_posted": profile.TotalJobsPosted,
	}
	payloadBytes, _ := json.Marshal(payload)
	metadata := map[string]interface{}{"source": "users-be"}
	metadataBytes, _ := json.Marshal(metadata)

	return &outbox.Event{
		AggregateID:   profile.UserID.String(),
		AggregateType: "client_profile",
		EventType:     eventType,
		Payload:       payloadBytes,
		Metadata:      metadataBytes,
	}, nil
}
