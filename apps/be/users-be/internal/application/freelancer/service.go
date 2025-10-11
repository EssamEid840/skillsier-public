package freelancer

import (
	"context"
	"encoding/json"
	"fmt"
	"users-be/internal/domain/freelancer"
	"users-be/internal/domain/outbox"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Service struct {
	freelancerRepo freelancer.Repository
	outboxRepo     outbox.Repository
	db             *gorm.DB
}

func NewService(freelancerRepo freelancer.Repository, outboxRepo outbox.Repository, db *gorm.DB) *Service {
	return &Service{
		freelancerRepo: freelancerRepo,
		outboxRepo:     outboxRepo,
		db:             db,
	}
}

func (s *Service) GetProfile(ctx context.Context, userID uuid.UUID) (*FreelancerProfileResponseDTO, error) {
	profile, err := s.freelancerRepo.GetByUserID(ctx, userID)
	if err != nil {
		if err == freelancer.ErrProfileNotFound {
			// Auto-create profile if it doesn't exist
			newProfile := &freelancer.FreelancerProfile{
				UserID: userID,
			}
			if err := s.freelancerRepo.Create(ctx, newProfile); err != nil {
				return nil, err
			}
			return ToResponseDTO(newProfile), nil
		}
		return nil, err
	}
	return ToResponseDTO(profile), nil
}

func (s *Service) UpdateProfile(ctx context.Context, userID uuid.UUID, dto *UpdateFreelancerProfileDTO) (*FreelancerProfileResponseDTO, error) {
	profile, err := s.freelancerRepo.GetByUserID(ctx, userID)
	if err != nil {
		if err == freelancer.ErrProfileNotFound {
			// Auto-create if doesn't exist
			profile = &freelancer.FreelancerProfile{UserID: userID}
			if err := s.freelancerRepo.Create(ctx, profile); err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}

	// Update fields
	if dto.Title != nil {
		profile.Title = dto.Title
	}
	if dto.Overview != nil {
		profile.Overview = dto.Overview
	}
	if dto.HourlyRate != nil {
		profile.HourlyRate = dto.HourlyRate
	}
	if dto.AvailableHours != nil {
		profile.AvailableHours = dto.AvailableHours
	}
	if dto.ResponseTime != nil {
		profile.ResponseTime = dto.ResponseTime
	}

	if err := profile.Validate(); err != nil {
		return nil, err
	}

	err = s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := s.freelancerRepo.Update(ctx, profile); err != nil {
			return err
		}

		event, err := s.createEvent("freelancer.profile.updated", profile)
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

func (s *Service) createEvent(eventType string, profile *freelancer.FreelancerProfile) (*outbox.Event, error) {
	payload := map[string]interface{}{
		"user_id":       profile.UserID.String(),
		"hourly_rate":   profile.HourlyRate,
		"total_jobs":    profile.TotalJobs,
		"total_earnings": profile.TotalEarnings,
	}
	payloadBytes, _ := json.Marshal(payload)
	metadata := map[string]interface{}{"source": "users-be"}
	metadataBytes, _ := json.Marshal(metadata)

	return &outbox.Event{
		AggregateID:   profile.UserID.String(),
		AggregateType: "freelancer_profile",
		EventType:     eventType,
		Payload:       payloadBytes,
		Metadata:      metadataBytes,
	}, nil
}