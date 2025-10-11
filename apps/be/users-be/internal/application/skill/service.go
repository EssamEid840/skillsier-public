package skill

import (
	"context"
	"encoding/json"
	"fmt"
	"users-be/internal/domain/skill"
	"users-be/internal/domain/outbox"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

const MaxSkillsPerUser = 20

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
	// Check max skills limit
	count, err := s.skillRepo.CountByUserID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to count skills: %w", err)
	}
	if count >= MaxSkillsPerUser {
		return nil, skill.ErrMaxSkillsExceeded
	}

	newSkill := &skill.Skill{
		UserID:     userID,
		Name:       dto.Name,
		Level:      dto.Level,
		YearsOfExp: dto.YearsOfExp,
		IsPrimary:  dto.IsPrimary,
	}

	if err := newSkill.Validate(); err != nil {
		return nil, err
	}

	// Use transaction for atomicity
	return s.createWithEvent(ctx, newSkill)
}

func (s *Service) GetSkills(ctx context.Context, userID uuid.UUID) ([]*SkillResponseDTO, error) {
	skills, err := s.skillRepo.GetByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	return ToResponseDTOList(skills), nil
}

func (s *Service) GetSkillByID(ctx context.Context, id uuid.UUID, userID uuid.UUID) (*SkillResponseDTO, error) {
	s, err := s.skillRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if s.UserID != userID {
		return nil, skill.ErrSkillNotFound
	}
	return ToResponseDTO(s), nil
}

func (s *Service) UpdateSkill(ctx context.Context, id uuid.UUID, userID uuid.UUID, dto *UpdateSkillDTO) (*SkillResponseDTO, error) {
	s, err := s.skillRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if s.UserID != userID {
		return nil, skill.ErrSkillNotFound
	}

	// Update fields
	if dto.Name != nil {
		s.Name = *dto.Name
	}
	if dto.Level != nil {
		s.Level = *dto.Level
	}
	if dto.YearsOfExp != nil {
		s.YearsOfExp = dto.YearsOfExp
	}
	if dto.IsPrimary != nil {
		s.IsPrimary = *dto.IsPrimary
	}

	if err := s.Validate(); err != nil {
		return nil, err
	}

	return s.updateWithEvent(ctx, s)
}

func (s *Service) DeleteSkill(ctx context.Context, id uuid.UUID, userID uuid.UUID) error {
	sk, err := s.skillRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if sk.UserID != userID {
		return skill.ErrSkillNotFound
	}

	return s.deleteWithEvent(ctx, sk)
}

// Helper methods for event creation
func (s *Service) createWithEvent(ctx context.Context, sk *skill.Skill) (*SkillResponseDTO, error) {
	err := s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := s.skillRepo.Create(ctx, sk); err != nil {
			return err
		}
		event, err := s.createSkillEvent("skill.created", sk)
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

func (s *Service) updateWithEvent(ctx context.Context, sk *skill.Skill) (*SkillResponseDTO, error) {
	err := s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
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

func (s *Service) deleteWithEvent(ctx context.Context, sk *skill.Skill) error {
	return s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := s.skillRepo.Delete(ctx, sk.ID); err != nil {
			return err
		}
		event, err := s.createSkillEvent("skill.deleted", sk)
		if err != nil {
			return err
		}
		return s.outboxRepo.Create(ctx, event)
	})
}

func (s *Service) createSkillEvent(eventType string, sk *skill.Skill) (*outbox.Event, error) {
	payload := map[string]interface{}{
		"skill_id":             sk.ID.String(),
		"user_id":              sk.UserID.String(),
		"name":                 sk.Name,
		"level":                sk.Level,
		"years_of_experience":  sk.YearsOfExp,
		"is_primary":           sk.IsPrimary,
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
		AggregateType: "skill",
		EventType:     eventType,
		Payload:       payloadBytes,
		Metadata:      metadataBytes,
	}, nil
}
