package experience

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"gorm.io/gorm"

	"users-be/internal/domain/experience"
	"users-be/internal/domain/outbox"
)

type Service struct {
	expRepo    experience.Repository
	outboxRepo outbox.Repository
	db         *gorm.DB
}

func NewService(expRepo experience.Repository, outboxRepo outbox.Repository, db *gorm.DB) *Service {
	return &Service{
		expRepo:    expRepo,
		outboxRepo: outboxRepo,
		db:         db,
	}
}

func (s *Service) CreateExperience(ctx context.Context, userID uuid.UUID, dto *CreateWorkExperienceDTO) (*WorkExperienceResponseDTO, error) {
	exp := &experience.WorkExperience{
		UserID:      userID,
		Title:       dto.Title,
		Company:     dto.Company,
		Location:    dto.Location,
		StartDate:   dto.StartDate,
		EndDate:     dto.EndDate,
		IsCurrent:   dto.IsCurrent,
		Description: dto.Description,
		Skills:      pq.StringArray(dto.Skills),
	}

	if err := exp.Validate(); err != nil {
		return nil, err
	}

	err := s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := s.expRepo.Create(ctx, exp); err != nil {
			return err
		}

		event, err := s.createExperienceEvent("experience.added", exp)
		if err != nil {
			return err
		}

		return s.outboxRepo.Create(ctx, event)
	})

	if err != nil {
		return nil, err
	}

	return ToResponseDTO(exp), nil
}

func (s *Service) GetAllExperiences(ctx context.Context, userID uuid.UUID) (*ListWorkExperiencesResponseDTO, error) {
	experiences, err := s.expRepo.GetByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	return ToListResponse(experiences), nil
}

func (s *Service) GetExperience(ctx context.Context, id uuid.UUID, userID uuid.UUID) (*WorkExperienceResponseDTO, error) {
	exp, err := s.expRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if exp.UserID != userID {
		return nil, fmt.Errorf("unauthorized")
	}

	return ToResponseDTO(exp), nil
}

func (s *Service) UpdateExperience(ctx context.Context, id uuid.UUID, userID uuid.UUID, dto *UpdateWorkExperienceDTO) (*WorkExperienceResponseDTO, error) {
	exp, err := s.expRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if exp.UserID != userID {
		return nil, fmt.Errorf("unauthorized")
	}

	if dto.Title != nil {
		exp.Title = *dto.Title
	}
	if dto.Company != nil {
		exp.Company = *dto.Company
	}
	if dto.Location != nil {
		exp.Location = *dto.Location
	}
	if dto.StartDate != nil {
		exp.StartDate = *dto.StartDate
	}
	if dto.EndDate != nil {
		exp.EndDate = dto.EndDate
	}
	if dto.IsCurrent != nil {
		exp.IsCurrent = *dto.IsCurrent
	}
	if dto.Description != nil {
		exp.Description = *dto.Description
	}
	if dto.Skills != nil {
		exp.Skills = pq.StringArray(dto.Skills)
	}

	if err := exp.Validate(); err != nil {
		return nil, err
	}

	err = s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := s.expRepo.Update(ctx, exp); err != nil {
			return err
		}

		event, err := s.createExperienceEvent("experience.updated", exp)
		if err != nil {
			return err
		}

		return s.outboxRepo.Create(ctx, event)
	})

	if err != nil {
		return nil, err
	}

	return ToResponseDTO(exp), nil
}

func (s *Service) DeleteExperience(ctx context.Context, id uuid.UUID, userID uuid.UUID) error {
	exp, err := s.expRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}

	if exp.UserID != userID {
		return fmt.Errorf("unauthorized")
	}

	return s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := s.expRepo.Delete(ctx, id); err != nil {
			return err
		}

		event, err := s.createExperienceEvent("experience.removed", exp)
		if err != nil {
			return err
		}

		return s.outboxRepo.Create(ctx, event)
	})
}

func (s *Service) createExperienceEvent(eventType string, exp *experience.WorkExperience) (*outbox.Event, error) {
	payload := map[string]interface{}{
		"experience_id": exp.ID.String(),
		"user_id":       exp.UserID.String(),
		"title":         exp.Title,
		"company":       exp.Company,
		"is_current":    exp.IsCurrent,
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
		AggregateID:   exp.UserID.String(),
		AggregateType: "user",
		EventType:     eventType,
		Payload:       payloadBytes,
		Metadata:      metadataBytes,
	}, nil
}