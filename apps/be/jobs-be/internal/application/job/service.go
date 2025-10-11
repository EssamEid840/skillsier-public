package job

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"jobs-be/internal/domain/job"
	"jobs-be/internal/domain/outbox"
)

type Service struct {
	jobRepo    job.Repository
	outboxRepo outbox.Repository
	db         *gorm.DB
}

func NewService(jobRepo job.Repository, outboxRepo outbox.Repository, db *gorm.DB) *Service {
	return &Service{
		jobRepo:    jobRepo,
		outboxRepo: outboxRepo,
		db:         db,
	}
}

func (s *Service) CreateJob(ctx context.Context, clientID uuid.UUID, dto *CreateJobDTO) (*JobResponseDTO, error) {
	j := &job.Job{
		ClientID:        clientID,
		Title:           dto.Title,
		Description:     dto.Description,
		Category:        dto.Category,
		BudgetType:      dto.BudgetType,
		BudgetAmount:    dto.BudgetAmount,
		HourlyRateMin:   dto.HourlyRateMin,
		HourlyRateMax:   dto.HourlyRateMax,
		Duration:        dto.Duration,
		ExperienceLevel: dto.ExperienceLevel,
		Status:          job.JobStatusOpen,
	}

	for _, skillDTO := range dto.Skills {
		j.Skills = append(j.Skills, job.JobSkill{
			Name:     skillDTO.Name,
			Required: skillDTO.Required,
		})
	}

	if err := j.Validate(); err != nil {
		return nil, err
	}

	err := s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := s.jobRepo.Create(ctx, j); err != nil {
			return err
		}

		event, err := s.createJobEvent("job.created", j)
		if err != nil {
			return err
		}

		return s.outboxRepo.Create(ctx, event)
	})

	if err != nil {
		return nil, err
	}

	return ToResponseDTO(j), nil
}

func (s *Service) GetJob(ctx context.Context, id uuid.UUID) (*JobResponseDTO, error) {
	j, err := s.jobRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return ToResponseDTO(j), nil
}

func (s *Service) GetAllJobs(ctx context.Context, page, pageSize int, filters map[string]interface{}) (*ListJobsResponseDTO, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 20
	}

	offset := (page - 1) * pageSize

	jobs, total, err := s.jobRepo.GetAll(ctx, pageSize, offset, filters)
	if err != nil {
		return nil, err
	}

	return ToListResponse(jobs, total, page, pageSize), nil
}

func (s *Service) UpdateJob(ctx context.Context, id, clientID uuid.UUID, dto *UpdateJobDTO) (*JobResponseDTO, error) {
	j, err := s.jobRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if j.ClientID != clientID {
		return nil, job.ErrUnauthorized
	}

	if dto.Title != nil {
		j.Title = *dto.Title
	}
	if dto.Description != nil {
		j.Description = *dto.Description
	}
	if dto.Status != nil {
		j.Status = *dto.Status
	}
	if dto.BudgetAmount != nil {
		j.BudgetAmount = *dto.BudgetAmount
	}
	if dto.HourlyRateMin != nil {
		j.HourlyRateMin = *dto.HourlyRateMin
	}
	if dto.HourlyRateMax != nil {
		j.HourlyRateMax = *dto.HourlyRateMax
	}
	if dto.Duration != nil {
		j.Duration = *dto.Duration
	}
	if dto.ExperienceLevel != nil {
		j.ExperienceLevel = *dto.ExperienceLevel
	}

	err = s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := s.jobRepo.Update(ctx, j); err != nil {
			return err
		}

		event, err := s.createJobEvent("job.updated", j)
		if err != nil {
			return err
		}

		return s.outboxRepo.Create(ctx, event)
	})

	if err != nil {
		return nil, err
	}

	return ToResponseDTO(j), nil
}

func (s *Service) DeleteJob(ctx context.Context, id, clientID uuid.UUID) error {
	j, err := s.jobRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}

	if j.ClientID != clientID {
		return job.ErrUnauthorized
	}

	return s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := s.jobRepo.Delete(ctx, id); err != nil {
			return err
		}

		event, err := s.createJobEvent("job.deleted", j)
		if err != nil {
			return err
		}

		return s.outboxRepo.Create(ctx, event)
	})
}

func (s *Service) GetMyJobs(ctx context.Context, clientID uuid.UUID, page, pageSize int) (*ListJobsResponseDTO, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 20
	}

	offset := (page - 1) * pageSize

	jobs, total, err := s.jobRepo.GetByClientID(ctx, clientID, pageSize, offset)
	if err != nil {
		return nil, err
	}

	return ToListResponse(jobs, total, page, pageSize), nil
}

func (s *Service) createJobEvent(eventType string, j *job.Job) (*outbox.Event, error) {
	payload := map[string]interface{}{
		"job_id":      j.ID.String(),
		"client_id":   j.ClientID.String(),
		"title":       j.Title,
		"status":      string(j.Status),
		"budget_type": string(j.BudgetType),
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal payload: %w", err)
	}

	metadata := map[string]interface{}{"source": "jobs-be"}
	metadataBytes, err := json.Marshal(metadata)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal metadata: %w", err)
	}

	return &outbox.Event{
		AggregateID:   j.ID.String(),
		AggregateType: "job",
		EventType:     eventType,
		Payload:       payloadBytes,
		Metadata:      metadataBytes,
	}, nil
}