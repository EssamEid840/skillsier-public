package education

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"users-be/internal/domain/education"
	"users-be/internal/domain/outbox"
)

type Service struct {
	eduRepo    education.Repository
	outboxRepo outbox.Repository
	db         *gorm.DB
}

func NewService(eduRepo education.Repository, outboxRepo outbox.Repository, db *gorm.DB) *Service {
	return &Service{
		eduRepo:    eduRepo,
		outboxRepo: outboxRepo,
		db:         db,
	}
}

func (s *Service) CreateEducation(ctx context.Context, userID uuid.UUID, dto *CreateEducationDTO) (*EducationResponseDTO, error) {
	edu := &education.Education{
		UserID:       userID,
		Degree:       dto.Degree,
		Institution:  dto.Institution,
		FieldOfStudy: dto.FieldOfStudy,
		StartDate:    dto.StartDate,
		EndDate:      dto.EndDate,
		Description:  dto.Description,
	}

	if err := edu.Validate(); err != nil {
		return nil, err
	}

	err := s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := s.eduRepo.Create(ctx, edu); err != nil {
			return err
		}

		event, err := s.createEducationEvent("education.added", edu)
		if err != nil {
			return err
		}

		return s.outboxRepo.Create(ctx, event)
	})

	if err != nil {
		return nil, err
	}

	return ToResponseDTO(edu), nil
}

func (s *Service) GetAllEducations(ctx context.Context, userID uuid.UUID) (*ListEducationsResponseDTO, error) {
	educations, err := s.eduRepo.GetByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	return ToListResponse(educations), nil
}

func (s *Service) GetEducation(ctx context.Context, id uuid.UUID, userID uuid.UUID) (*EducationResponseDTO, error) {
	edu, err := s.eduRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if edu.UserID != userID {
		return nil, fmt.Errorf("unauthorized")
	}

	return ToResponseDTO(edu), nil
}

func (s *Service) UpdateEducation(ctx context.Context, id uuid.UUID, userID uuid.UUID, dto *UpdateEducationDTO) (*EducationResponseDTO, error) {
	edu, err := s.eduRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if edu.UserID != userID {
		return nil, fmt.Errorf("unauthorized")
	}

	if dto.Degree != nil {
		edu.Degree = *dto.Degree
	}
	if dto.Institution != nil {
		edu.Institution = *dto.Institution
	}
	if dto.FieldOfStudy != nil {
		edu.FieldOfStudy = *dto.FieldOfStudy
	}
	if dto.StartDate != nil {
		edu.StartDate = *dto.StartDate
	}
	if dto.EndDate != nil {
		edu.EndDate = dto.EndDate
	}
	if dto.Description != nil {
		edu.Description = *dto.Description
	}

	if err := edu.Validate(); err != nil {
		return nil, err
	}

	err = s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := s.eduRepo.Update(ctx, edu); err != nil {
			return err
		}

		event, err := s.createEducationEvent("education.updated", edu)
		if err != nil {
			return err
		}

		return s.outboxRepo.Create(ctx, event)
	})

	if err != nil {
		return nil, err
	}

	return ToResponseDTO(edu), nil
}

func (s *Service) DeleteEducation(ctx context.Context, id uuid.UUID, userID uuid.UUID) error {
	edu, err := s.eduRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}

	if edu.UserID != userID {
		return fmt.Errorf("unauthorized")
	}

	return s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := s.eduRepo.Delete(ctx, id); err != nil {
			return err
		}

		event, err := s.createEducationEvent("education.removed", edu)
		if err != nil {
			return err
		}

		return s.outboxRepo.Create(ctx, event)
	})
}

func (s *Service) createEducationEvent(eventType string, edu *education.Education) (*outbox.Event, error) {
	payload := map[string]interface{}{
		"education_id":   edu.ID.String(),
		"user_id":        edu.UserID.String(),
		"degree":         edu.Degree,
		"institution":    edu.Institution,
		"field_of_study": edu.FieldOfStudy,
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal payload: %w", err)
	}

	metadata := map[string]interface{}{"source": "users-be"}
	metadataBytes, err := json.Marshal(metadata)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal metadata: %w", err)
	}

	return &outbox.Event{
		AggregateID:   edu.UserID.String(),
		AggregateType: "user",
		EventType:     eventType,
		Payload:       payloadBytes,
		Metadata:      metadataBytes,
	}, nil
}