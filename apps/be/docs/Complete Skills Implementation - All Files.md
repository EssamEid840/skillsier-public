// ==========================================
// FILE: internal/domain/skill/entity.go
// ==========================================
package skill

import (
	"time"
	"github.com/google/uuid"
)

type Skill struct {
	ID           uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	UserID       uuid.UUID `gorm:"type:uuid;not null;index" json:"user_id"`
	Name         string    `gorm:"type:varchar(100);not null" json:"name"`
	Level        string    `gorm:"type:varchar(50);not null" json:"level"`
	YearsOfExp   *int      `gorm:"type:int" json:"years_of_experience,omitempty"`
	Endorsements int       `gorm:"type:int;default:0" json:"endorsements"`
	IsPrimary    bool      `gorm:"type:boolean;default:false" json:"is_primary"`
	CreatedAt    time.Time `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt    time.Time `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (Skill) TableName() string {
	return "skills"
}

func (s *Skill) Validate() error {
	validLevels := map[string]bool{
		"beginner": true, "intermediate": true, "advanced": true, "expert": true,
	}
	if !validLevels[s.Level] {
		return ErrInvalidSkillLevel
	}
	if s.Name == "" {
		return ErrSkillNameRequired
	}
	if s.YearsOfExp != nil && *s.YearsOfExp < 0 {
		return ErrInvalidYearsOfExp
	}
	return nil
}

// ==========================================
// FILE: internal/domain/skill/repository.go
// ==========================================
package skill

import (
	"context"
	"errors"
	"github.com/google/uuid"
)

var (
	ErrSkillNotFound       = errors.New("skill not found")
	ErrInvalidSkillLevel   = errors.New("invalid skill level")
	ErrSkillNameRequired   = errors.New("skill name is required")
	ErrInvalidYearsOfExp   = errors.New("years of experience cannot be negative")
	ErrMaxSkillsExceeded   = errors.New("maximum number of skills exceeded")
)

type Repository interface {
	Create(ctx context.Context, skill *Skill) error
	GetByID(ctx context.Context, id uuid.UUID) (*Skill, error)
	GetByUserID(ctx context.Context, userID uuid.UUID) ([]*Skill, error)
	Update(ctx context.Context, skill *Skill) error
	Delete(ctx context.Context, id uuid.UUID) error
	CountByUserID(ctx context.Context, userID uuid.UUID) (int64, error)
}

// ==========================================
// FILE: internal/infrastructure/persistence/postgres/skill_repository.go
// ==========================================
package postgres

import (
	"context"
	"users-be/internal/domain/skill"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type skillRepository struct {
	db *gorm.DB
}

func NewSkillRepository(db *gorm.DB) skill.Repository {
	return &skillRepository{db: db}
}

func (r *skillRepository) Create(ctx context.Context, s *skill.Skill) error {
	return r.db.WithContext(ctx).Create(s).Error
}

func (r *skillRepository) GetByID(ctx context.Context, id uuid.UUID) (*skill.Skill, error) {
	var s skill.Skill
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&s).Error
	if err == gorm.ErrRecordNotFound {
		return nil, skill.ErrSkillNotFound
	}
	return &s, err
}

func (r *skillRepository) GetByUserID(ctx context.Context, userID uuid.UUID) ([]*skill.Skill, error) {
	var skills []*skill.Skill
	err := r.db.WithContext(ctx).Where("user_id = ?", userID).
		Order("is_primary DESC, created_at DESC").Find(&skills).Error
	return skills, err
}

func (r *skillRepository) Update(ctx context.Context, s *skill.Skill) error {
	return r.db.WithContext(ctx).Model(s).Updates(s).Error
}

func (r *skillRepository) Delete(ctx context.Context, id uuid.UUID) error {
	result := r.db.WithContext(ctx).Delete(&skill.Skill{}, "id = ?", id)
	if result.RowsAffected == 0 {
		return skill.ErrSkillNotFound
	}
	return result.Error
}

func (r *skillRepository) CountByUserID(ctx context.Context, userID uuid.UUID) (int64, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&skill.Skill{}).
		Where("user_id = ?", userID).Count(&count).Error
	return count, err
}

// ==========================================
// FILE: internal/application/skill/dto.go
// ==========================================
package skill

import (
	"time"
	"github.com/google/uuid"
)

type CreateSkillDTO struct {
	Name       string `json:"name" binding:"required,max=100"`
	Level      string `json:"level" binding:"required,oneof=beginner intermediate advanced expert"`
	YearsOfExp *int   `json:"years_of_experience" binding:"omitempty,min=0"`
	IsPrimary  bool   `json:"is_primary"`
}

type UpdateSkillDTO struct {
	Name       *string `json:"name" binding:"omitempty,max=100"`
	Level      *string `json:"level" binding:"omitempty,oneof=beginner intermediate advanced expert"`
	YearsOfExp *int    `json:"years_of_experience" binding:"omitempty,min=0"`
	IsPrimary  *bool   `json:"is_primary"`
}

type SkillResponseDTO struct {
	ID           uuid.UUID `json:"id"`
	UserID       uuid.UUID `json:"user_id"`
	Name         string    `json:"name"`
	Level        string    `json:"level"`
	YearsOfExp   *int      `json:"years_of_experience,omitempty"`
	Endorsements int       `json:"endorsements"`
	IsPrimary    bool      `json:"is_primary"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// ==========================================
// FILE: internal/application/skill/mapper.go
// ==========================================
package skill

import "users-be/internal/domain/skill"

func ToResponseDTO(s *skill.Skill) *SkillResponseDTO {
	return &SkillResponseDTO{
		ID:           s.ID,
		UserID:       s.UserID,
		Name:         s.Name,
		Level:        s.Level,
		YearsOfExp:   s.YearsOfExp,
		Endorsements: s.Endorsements,
		IsPrimary:    s.IsPrimary,
		CreatedAt:    s.CreatedAt,
		UpdatedAt:    s.UpdatedAt,
	}
}

func ToResponseDTOList(skills []*skill.Skill) []*SkillResponseDTO {
	result := make([]*SkillResponseDTO, len(skills))
	for i, s := range skills {
		result[i] = ToResponseDTO(s)
	}
	return result
}

// ==========================================
// FILE: internal/application/skill/service.go
// ==========================================
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

// ==========================================
// FILE: internal/interfaces/http/handlers/skill_handler.go
// ==========================================
package handlers

import (
	"net/http"
	"users-be/internal/application/skill"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type SkillHandler struct {
	skillService *skill.Service
}

func NewSkillHandler(skillService *skill.Service) *SkillHandler {
	return &SkillHandler{skillService: skillService}
}

// CreateSkill handles POST /users/profile/skills
func (h *SkillHandler) CreateSkill(c *gin.Context) {
	userID, err := uuid.Parse(c.GetString("user_id"))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid user"})
		return
	}

	var dto skill.CreateSkillDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := h.skillService.CreateSkill(c.Request.Context(), userID, &dto)
	if err != nil {
		statusCode := http.StatusInternalServerError
		if err == skill.ErrMaxSkillsExceeded {
			statusCode = http.StatusBadRequest
		}
		c.JSON(statusCode, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, result)
}

// GetSkills handles GET /users/profile/skills
func (h *SkillHandler) GetSkills(c *gin.Context) {
	userID, err := uuid.Parse(c.GetString("user_id"))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid user"})
		return
	}

	skills, err := h.skillService.GetSkills(c.Request.Context(), userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"skills": skills})
}

// UpdateSkill handles PATCH /users/profile/skills/:id
func (h *SkillHandler) UpdateSkill(c *gin.Context) {
	userID, err := uuid.Parse(c.GetString("user_id"))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid user"})
		return
	}

	skillID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid skill ID"})
		return
	}

	var dto skill.UpdateSkillDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := h.skillService.UpdateSkill(c.Request.Context(), skillID, userID, &dto)
	if err != nil {
		statusCode := http.StatusInternalServerError
		if err == skill.ErrSkillNotFound {
			statusCode = http.StatusNotFound
		}
		c.JSON(statusCode, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

// DeleteSkill handles DELETE /users/profile/skills/:id
func (h *SkillHandler) DeleteSkill(c *gin.Context) {
	userID, err := uuid.Parse(c.GetString("user_id"))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid user"})
		return
	}

	skillID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid skill ID"})
		return
	}

	if err := h.skillService.DeleteSkill(c.Request.Context(), skillID, userID); err != nil {
		statusCode := http.StatusInternalServerError
		if err == skill.ErrSkillNotFound {
			statusCode = http.StatusNotFound
		}
		c.JSON(statusCode, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "skill deleted successfully"})
}