package skill

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"users-be/internal/domain/outbox"
	"users-be/internal/domain/skill"
)

type Service struct {
	skillRepo  skill.Repository
	outboxRepo outbox.Repository
	db         *gorm.DB
}

func NewService(skillRepo skill.Repository, outboxRepo outbox.Repository, db *gorm.DB) *Service {
	return &Service{
		skillRepo:  skillRepo,
		outboxRepo: outboxRepo,
		db:         db,
	}
}

func (s *Service) CreateSkill(ctx context.Context, userID uuid.UUID, dto *CreateSkillDTO) (*SkillResponseDTO, error) {
	existing, _ := s.skillRepo.GetByUserIDAndName(ctx, userID, dto.Name)
	if existing != nil {
		return nil, fmt.Errorf("skill already exists")
	}

	sk := &skill.Skill{
		UserID:            userID,
		Name:              dto.Name,
		Category:          dto.Category,
		Level:             dto.Level,
		YearsOfExperience: dto.YearsOfExperience,
	}

	if err := sk.Validate(); err != nil {
		return nil, err
	}

	err := s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := s.skillRepo.Create(ctx, sk); err != nil {
			return err
		}

		event, err := s.createSkillEvent("skill.added", sk)
		if err != nil {
			return err
		}

		return s.outboxRepo.Create(ctx, event)
	})

	if err != nil {
		return nil, err
	}

	return ToResponseDTO(sk), nil
}

func (s *Service) GetAllSkills(ctx context.Context, userID uuid.UUID) (*ListSkillsResponseDTO, error) {
	skills, err := s.skillRepo.GetByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	return ToListResponse(skills), nil
}

func (s *Service) GetSkill(ctx context.Context, id uuid.UUID, userID uuid.UUID) (*SkillResponseDTO, error) {
	sk, err := s.skillRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if sk.UserID != userID {
		return nil, fmt.Errorf("unauthorized")
	}

	return ToResponseDTO(sk), nil
}

func (s *Service) UpdateSkill(ctx context.Context, id uuid.UUID, userID uuid.UUID, dto *UpdateSkillDTO) (*SkillResponseDTO, error) {
	sk, err := s.skillRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if sk.UserID != userID {
		return nil, fmt.Errorf("unauthorized")
	}

	if dto.Level != nil {
		sk.Level = *dto.Level
	}
	if dto.YearsOfExperience != nil {
		sk.YearsOfExperience = *dto.YearsOfExperience
	}

	if err := sk.Validate(); err != nil {
		return nil, err
	}

	err = s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := s.skillRepo.Update(ctx, sk); err != nil {
			return err
		}

		event, err := s.createSkillEvent("skill.updated", sk)
		if err != nil {
			return err
		}

		return s.outboxRepo.Create(ctx, event)
	})

	if err != nil {
		return nil, err
	}

	return ToResponseDTO(sk), nil
}

func (s *Service) DeleteSkill(ctx context.Context, id uuid.UUID, userID uuid.UUID) error {
	sk, err := s.skillRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}

	if sk.UserID != userID {
		return fmt.Errorf("unauthorized")
	}

	return s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := s.skillRepo.Delete(ctx, id); err != nil {
			return err
		}

		event, err := s.createSkillEvent("skill.removed", sk)
		if err != nil {
			return err
		}

		return s.outboxRepo.Create(ctx, event)
	})
}

func (s *Service) createSkillEvent(eventType string, sk *skill.Skill) (*outbox.Event, error) {
	payload := map[string]interface{}{
		"skill_id":            sk.ID.String(),
		"user_id":             sk.UserID.String(),
		"name":                sk.Name,
		"category":            sk.Category,
		"level":               string(sk.Level),
		"years_of_experience": sk.YearsOfExperience,
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
		AggregateID:   sk.UserID.String(),
		AggregateType: "user",
		EventType:     eventType,
		Payload:       payloadBytes,
		Metadata:      metadataBytes,
	}, nil
}