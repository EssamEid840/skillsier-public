package freelancer

import (
	"context"
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
	return &Service{freelancerRepo: freelancerRepo, outboxRepo: outboxRepo, db: db}
}

func (s *Service) GetProfile(ctx context.Context, userID uuid.UUID) (*FreelancerProfileResponseDTO, error) {
	profile, err := s.freelancerRepo.GetByUserID(ctx, userID)
	if err != nil {
		if err == freelancer.ErrFreelancerProfileNotFound {
			profile = &freelancer.FreelancerProfile{UserID: userID}
			if err := s.freelancerRepo.Create(ctx, profile); err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}
	return ToResponseDTO(profile), nil
}

func (s *Service) UpdateProfile(ctx context.Context, userID uuid.UUID, dto *UpdateFreelancerProfileDTO) (*FreelancerProfileResponseDTO, error) {
	profile, err := s.freelancerRepo.GetByUserID(ctx, userID)
	if err != nil {
		if err == freelancer.ErrFreelancerProfileNotFound {
			profile = &freelancer.FreelancerProfile{UserID: userID}
			s.freelancerRepo.Create(ctx, profile)
		} else {
			return nil, err
		}
	}
	if dto.ProfessionalTitle != nil {
		profile.ProfessionalTitle = *dto.ProfessionalTitle
	}
	if dto.Overview != nil {
		profile.Overview = *dto.Overview
	}
	if dto.HourlyRate != nil {
		profile.HourlyRate = *dto.HourlyRate
	}
	if dto.Availability != nil {
		profile.Availability = *dto.Availability
	}
	if err := profile.Validate(); err != nil {
		return nil, err
	}
	if err := s.freelancerRepo.Update(ctx, profile); err != nil {
		return nil, err
	}
	return ToResponseDTO(profile), nil
}