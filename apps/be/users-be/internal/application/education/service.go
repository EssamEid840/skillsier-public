package education

import (
	"context"
	"encoding/json"
	"fmt"
	"users-be/internal/domain/education"
	"users-be/internal/domain/outbox"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

const MaxEducationPerUser = 5

type Service struct {
	eduRepo    education.Repository
	outboxRepo outbox.Repository
	db         *gorm.DB
}

func NewService(eduRepo education.Repository, outboxRepo outbox.Repository, db *gorm.DB) *Service {
	return &Service{eduRepo: eduRepo, outboxRepo: outboxRepo, db: db}
}

func (s *Service) Create(ctx context.Context, userID uuid.UUID, dto *CreateEducationDTO) (*EducationResponseDTO, error) {
	count, err := s.eduRepo.CountByUserID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to count education: %w", err)
	}
	if count >= MaxEducationPerUser {
		return nil, education.ErrMaxEducationExceeded
	}

	newEdu := &education.Education{
		UserID:       userID,
		School:       dto.School,
		Degree:       dto.Degree,
		FieldOfStudy: dto.FieldOfStudy,
		StartDate:    dto.StartDate,
		EndDate:      dto.EndDate,
		IsCurrent:    dto.IsCurrent,
		Description:  dto.Description,
	}

	if err := newEdu.Validate(); err != nil {
		return nil, err
	}

	err = s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := s.eduRepo.Create(ctx, newEdu); err != nil {
			return err
		}
		event, err := s.createEvent("education.added", newEdu)
		if err != nil {
			return err
		}
		return s.outboxRepo.Create(ctx, event)
	})
	if err != nil {
		return nil, err
	}
	return ToResponseDTO(newEdu), nil
}

func (s *Service) createEvent(eventType string, edu *education.Education) (*outbox.Event, error) {
	payload := map[string]interface{}{
		"education_id": edu.ID.String(),
		"user_id":      edu.UserID.String(),
		"school":       edu.School,
		"degree":       edu.Degree,
	}
	payloadBytes, _ := json.Marshal(payload)
	metadata := map[string]interface{}{"source": "users-be"}
	metadataBytes, _ := json.Marshal(metadata)

	return &outbox.Event{
		AggregateID:   edu.UserID.String(),
		AggregateType: "education",
		EventType:     eventType,
		Payload:       payloadBytes,
		Metadata:      metadataBytes,
	}, nil
}