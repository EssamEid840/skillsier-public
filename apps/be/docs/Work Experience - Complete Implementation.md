// ==========================================
// FILE: internal/domain/experience/entity.go
// ==========================================
package experience

import (
	"time"
	"github.com/google/uuid"
)

type WorkExperience struct {
	ID          uuid.UUID  `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	UserID      uuid.UUID  `gorm:"type:uuid;not null;index" json:"user_id"`
	Title       string     `gorm:"type:varchar(200);not null" json:"title"`
	Company     string     `gorm:"type:varchar(200);not null" json:"company"`
	Location    *string    `gorm:"type:varchar(200)" json:"location,omitempty"`
	StartDate   time.Time  `gorm:"type:date;not null" json:"start_date"`
	EndDate     *time.Time `gorm:"type:date" json:"end_date,omitempty"`
	IsCurrent   bool       `gorm:"type:boolean;default:false" json:"is_current"`
	Description *string    `gorm:"type:text" json:"description,omitempty"`
	CreatedAt   time.Time  `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time  `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (WorkExperience) TableName() string {
	return "work_experiences"
}

func (we *WorkExperience) Validate() error {
	if we.Title == "" {
		return ErrTitleRequired
	}
	if we.Company == "" {
		return ErrCompanyRequired
	}
	if we.EndDate != nil && !we.IsCurrent && we.EndDate.Before(we.StartDate) {
		return ErrInvalidDateRange
	}
	if we.IsCurrent && we.EndDate != nil {
		return ErrCurrentWithEndDate
	}
	return nil
}

// ==========================================
// FILE: internal/domain/experience/repository.go
// ==========================================
package experience

import (
	"context"
	"errors"
	"github.com/google/uuid"
)

var (
	ErrExperienceNotFound  = errors.New("work experience not found")
	ErrTitleRequired       = errors.New("title is required")
	ErrCompanyRequired     = errors.New("company name is required")
	ErrInvalidDateRange    = errors.New("end date must be after start date")
	ErrCurrentWithEndDate  = errors.New("current position cannot have end date")
	ErrMaxExperienceExceeded = errors.New("maximum number of experiences exceeded")
)

type Repository interface {
	Create(ctx context.Context, exp *WorkExperience) error
	GetByID(ctx context.Context, id uuid.UUID) (*WorkExperience, error)
	GetByUserID(ctx context.Context, userID uuid.UUID) ([]*WorkExperience, error)
	Update(ctx context.Context, exp *WorkExperience) error
	Delete(ctx context.Context, id uuid.UUID) error
	CountByUserID(ctx context.Context, userID uuid.UUID) (int64, error)
}

// ==========================================
// FILE: internal/infrastructure/persistence/postgres/experience_repository.go
// ==========================================
package postgres

import (
	"context"
	"users-be/internal/domain/experience"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type experienceRepository struct {
	db *gorm.DB
}

func NewExperienceRepository(db *gorm.DB) experience.Repository {
	return &experienceRepository{db: db}
}

func (r *experienceRepository) Create(ctx context.Context, exp *experience.WorkExperience) error {
	return r.db.WithContext(ctx).Create(exp).Error
}

func (r *experienceRepository) GetByID(ctx context.Context, id uuid.UUID) (*experience.WorkExperience, error) {
	var exp experience.WorkExperience
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&exp).Error
	if err == gorm.ErrRecordNotFound {
		return nil, experience.ErrExperienceNotFound
	}
	return &exp, err
}

func (r *experienceRepository) GetByUserID(ctx context.Context, userID uuid.UUID) ([]*experience.WorkExperience, error) {
	var experiences []*experience.WorkExperience
	err := r.db.WithContext(ctx).Where("user_id = ?", userID).
		Order("is_current DESC, start_date DESC").Find(&experiences).Error
	return experiences, err
}

func (r *experienceRepository) Update(ctx context.Context, exp *experience.WorkExperience) error {
	return r.db.WithContext(ctx).Model(exp).Updates(exp).Error
}

func (r *experienceRepository) Delete(ctx context.Context, id uuid.UUID) error {
	result := r.db.WithContext(ctx).Delete(&experience.WorkExperience{}, "id = ?", id)
	if result.RowsAffected == 0 {
		return experience.ErrExperienceNotFound
	}
	return result.Error
}

func (r *experienceRepository) CountByUserID(ctx context.Context, userID uuid.UUID) (int64, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&experience.WorkExperience{}).
		Where("user_id = ?", userID).Count(&count).Error
	return count, err
}

// ==========================================
// FILE: internal/application/experience/dto.go
// ==========================================
package experience

import (
	"time"
	"github.com/google/uuid"
)

type CreateWorkExperienceDTO struct {
	Title       string     `json:"title" binding:"required,max=200"`
	Company     string     `json:"company" binding:"required,max=200"`
	Location    *string    `json:"location" binding:"omitempty,max=200"`
	StartDate   time.Time  `json:"start_date" binding:"required"`
	EndDate     *time.Time `json:"end_date"`
	IsCurrent   bool       `json:"is_current"`
	Description *string    `json:"description" binding:"omitempty,max=5000"`
}

type UpdateWorkExperienceDTO struct {
	Title       *string    `json:"title" binding:"omitempty,max=200"`
	Company     *string    `json:"company" binding:"omitempty,max=200"`
	Location    *string    `json:"location" binding:"omitempty,max=200"`
	StartDate   *time.Time `json:"start_date"`
	EndDate     *time.Time `json:"end_date"`
	IsCurrent   *bool      `json:"is_current"`
	Description *string    `json:"description" binding:"omitempty,max=5000"`
}

type WorkExperienceResponseDTO struct {
	ID          uuid.UUID  `json:"id"`
	UserID      uuid.UUID  `json:"user_id"`
	Title       string     `json:"title"`
	Company     string     `json:"company"`
	Location    *string    `json:"location,omitempty"`
	StartDate   time.Time  `json:"start_date"`
	EndDate     *time.Time `json:"end_date,omitempty"`
	IsCurrent   bool       `json:"is_current"`
	Description *string    `json:"description,omitempty"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

// ==========================================
// FILE: internal/application/experience/mapper.go
// ==========================================
package experience

import "users-be/internal/domain/experience"

func ToResponseDTO(exp *experience.WorkExperience) *WorkExperienceResponseDTO {
	return &WorkExperienceResponseDTO{
		ID:          exp.ID,
		UserID:      exp.UserID,
		Title:       exp.Title,
		Company:     exp.Company,
		Location:    exp.Location,
		StartDate:   exp.StartDate,
		EndDate:     exp.EndDate,
		IsCurrent:   exp.IsCurrent,
		Description: exp.Description,
		CreatedAt:   exp.CreatedAt,
		UpdatedAt:   exp.UpdatedAt,
	}
}

func ToResponseDTOList(experiences []*experience.WorkExperience) []*WorkExperienceResponseDTO {
	result := make([]*WorkExperienceResponseDTO, len(experiences))
	for i, exp := range experiences {
		result[i] = ToResponseDTO(exp)
	}
	return result
}

// ==========================================
// FILE: internal/application/experience/service.go
// ==========================================
package experience

import (
	"context"
	"encoding/json"
	"fmt"
	"users-be/internal/domain/experience"
	"users-be/internal/domain/outbox"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

const MaxExperiencePerUser = 10

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
	count, err := s.expRepo.CountByUserID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to count experiences: %w", err)
	}
	if count >= MaxExperiencePerUser {
		return nil, experience.ErrMaxExperienceExceeded
	}

	newExp := &experience.WorkExperience{
		UserID:      userID,
		Title:       dto.Title,
		Company:     dto.Company,
		Location:    dto.Location,
		StartDate:   dto.StartDate,
		EndDate:     dto.EndDate,
		IsCurrent:   dto.IsCurrent,
		Description: dto.Description,
	}

	if err := newExp.Validate(); err != nil {
		return nil, err
	}

	return s.createWithEvent(ctx, newExp)
}

func (s *Service) GetExperiences(ctx context.Context, userID uuid.UUID) ([]*WorkExperienceResponseDTO, error) {
	experiences, err := s.expRepo.GetByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	return ToResponseDTOList(experiences), nil
}

func (s *Service) GetExperienceByID(ctx context.Context, id uuid.UUID, userID uuid.UUID) (*WorkExperienceResponseDTO, error) {
	exp, err := s.expRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if exp.UserID != userID {
		return nil, experience.ErrExperienceNotFound
	}
	return ToResponseDTO(exp), nil
}

func (s *Service) UpdateExperience(ctx context.Context, id uuid.UUID, userID uuid.UUID, dto *UpdateWorkExperienceDTO) (*WorkExperienceResponseDTO, error) {
	exp, err := s.expRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if exp.UserID != userID {
		return nil, experience.ErrExperienceNotFound
	}

	if dto.Title != nil {
		exp.Title = *dto.Title
	}
	if dto.Company != nil {
		exp.Company = *dto.Company
	}
	if dto.Location != nil {
		exp.Location = dto.Location
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
		exp.Description = dto.Description
	}

	if err := exp.Validate(); err != nil {
		return nil, err
	}

	return s.updateWithEvent(ctx, exp)
}

func (s *Service) DeleteExperience(ctx context.Context, id uuid.UUID, userID uuid.UUID) error {
	exp, err := s.expRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if exp.UserID != userID {
		return experience.ErrExperienceNotFound
	}

	return s.deleteWithEvent(ctx, exp)
}

func (s *Service) createWithEvent(ctx context.Context, exp *experience.WorkExperience) (*WorkExperienceResponseDTO, error) {
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

func (s *Service) updateWithEvent(ctx context.Context, exp *experience.WorkExperience) (*WorkExperienceResponseDTO, error) {
	err := s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
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

func (s *Service) deleteWithEvent(ctx context.Context, exp *experience.WorkExperience) error {
	return s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := s.expRepo.Delete(ctx, exp.ID); err != nil {
			return err
		}
		event, err := s.createExperienceEvent("experience.deleted", exp)
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
		AggregateType: "work_experience",
		EventType:     eventType,
		Payload:       payloadBytes,
		Metadata:      metadataBytes,
	}, nil
}

// ==========================================
// FILE: internal/interfaces/http/handlers/experience_handler.go
// ==========================================
package handlers

import (
	"net/http"
	"users-be/internal/application/experience"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ExperienceHandler struct {
	service *experience.Service
}

func NewExperienceHandler(service *experience.Service) *ExperienceHandler {
	return &ExperienceHandler{service: service}
}

func (h *ExperienceHandler) Create(c *gin.Context) {
	userID, err := uuid.Parse(c.GetString("user_id"))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid user"})
		return
	}

	var dto experience.CreateWorkExperienceDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := h.service.CreateExperience(c.Request.Context(), userID, &dto)
	if err != nil {
		statusCode := http.StatusInternalServerError
		if err == experience.ErrMaxExperienceExceeded {
			statusCode = http.StatusBadRequest
		}
		c.JSON(statusCode, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, result)
}

func (h *ExperienceHandler) GetAll(c *gin.Context) {
	userID, err := uuid.Parse(c.GetString("user_id"))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid user"})
		return
	}

	experiences, err := h.service.GetExperiences(c.Request.Context(), userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"experiences": experiences})
}

func (h *ExperienceHandler) Update(c *gin.Context) {
	userID, err := uuid.Parse(c.GetString("user_id"))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid user"})
		return
	}

	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid experience ID"})
		return
	}

	var dto experience.UpdateWorkExperienceDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := h.service.UpdateExperience(c.Request.Context(), id, userID, &dto)
	if err != nil {
		statusCode := http.StatusInternalServerError
		if err == experience.ErrExperienceNotFound {
			statusCode = http.StatusNotFound
		}
		c.JSON(statusCode, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

func (h *ExperienceHandler) Delete(c *gin.Context) {
	userID, err := uuid.Parse(c.GetString("user_id"))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid user"})
		return
	}

	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid experience ID"})
		return
	}

	if err := h.service.DeleteExperience(c.Request.Context(), id, userID); err != nil {
		statusCode := http.StatusInternalServerError
		if err == experience.ErrExperienceNotFound {
			statusCode = http.StatusNotFound
		}
		c.JSON(statusCode, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "work experience deleted successfully"})
}