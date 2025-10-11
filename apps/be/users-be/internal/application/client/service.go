package client

import (
	"context"
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
	return &Service{clientRepo: clientRepo, outboxRepo: outboxRepo, db: db}
}

func (s *Service) GetProfile(ctx context.Context, userID uuid.UUID) (*ClientProfileResponseDTO, error) {
	profile, err := s.clientRepo.GetByUserID(ctx, userID)
	if err != nil {
		if err == client.ErrClientProfileNotFound {
			profile = &client.ClientProfile{UserID: userID}
			if err := s.clientRepo.Create(ctx, profile); err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}
	return ToResponseDTO(profile), nil
}

func (s *Service) UpdateProfile(ctx context.Context, userID uuid.UUID, dto *UpdateClientProfileDTO) (*ClientProfileResponseDTO, error) {
	profile, err := s.clientRepo.GetByUserID(ctx, userID)
	if err != nil {
		if err == client.ErrClientProfileNotFound {
			profile = &client.ClientProfile{UserID: userID}
			s.clientRepo.Create(ctx, profile)
		} else {
			return nil, err
		}
	}
	if dto.CompanyName != nil {
		profile.CompanyName = *dto.CompanyName
	}
	if dto.CompanySize != nil {
		profile.CompanySize = *dto.CompanySize
	}
	if dto.Industry != nil {
		profile.Industry = *dto.Industry
	}
	if dto.Website != nil {
		profile.Website = *dto.Website
	}
	if err := profile.Validate(); err != nil {
		return nil, err
	}
	if err := s.clientRepo.Update(ctx, profile); err != nil {
		return nil, err
	}
	return ToResponseDTO(profile), nil
}