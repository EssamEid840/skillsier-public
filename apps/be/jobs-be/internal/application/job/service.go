package job

import (
	"context"
	"encoding/json"
	"fmt"
	"jobs-be/internal/domain/job"
	"jobs-be/internal/domain/outbox"
	"github.com/google/uuid"
	"gorm.io/gorm"
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
	newJob := &job.Job{
		ClientID:        clientID,
		Title:           dto.Title,
		Description:     dto.Description,
		Category:        dto.Category,
		BudgetType:      job.BudgetType(dto.BudgetType),
		BudgetAmount:    dto.BudgetAmount,
		HourlyRateMin:   dto.HourlyRateMin,
		HourlyRateMax:   dto.HourlyRateMax,
		Duration:        dto.Duration,
		ExperienceLevel: dto.ExperienceLevel,
		Status:          job.JobStatusOpen,
	}

	// Add required skills
	for _, skillDTO := range dto.RequiredSkills {
		newJob.RequiredSkills = append(newJob.RequiredSkills, job.JobSkill{
			Name:  skillDTO.Name,
			Level: skillDTO.Level,
		})
	}

	if err := newJob.Validate(); err != nil {
		return nil, err
	}

	// Transaction: Create job + outbox event
	err := s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := s.jobRepo.Create(ctx, newJob); err != nil {
			return err
		}

		event, err := s.createJobEvent("job.created", newJob)
		if err != nil {
			return err
		}

		return s.outboxRepo.Create(ctx, event)
	})

	if err != nil {
		return nil, err
	}

	return ToResponseDTO(newJob), nil
}

func (s *Service) GetJob(ctx context.Context, id uuid.UUID) (*JobResponseDTO, error) {
	j, err := s.jobRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return ToResponseDTO(j), nil
}

func (s *Service) ListJobs(ctx context.Context, filters *job.ListFilters, page, pageSize int) (*ListJobsResponseDTO, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 20
	}

	offset := (page - 1) * pageSize

	jobs, total, err := s.jobRepo.List(ctx, filters, pageSize, offset)
	if err != nil {
		return nil, err
	}

	return ToListResponse(jobs, total, page, pageSize), nil
}

func (s *Service) UpdateJob(ctx context.Context, id uuid.UUID, clientID uuid.UUID, dto *UpdateJobDTO) (*JobResponseDTO, error) {
	j, err := s.jobRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	// Verify ownership
	if j.ClientID != clientID {
		return nil, job.ErrUnauthorized
	}

	// Update fields
	if dto.Title != nil {
		j.Title = *dto.Title
	}
	if dto.Description != nil {
		j.Description = *dto.Description
	}
	if dto.Status != nil {
		j.Status = job.JobStatus(*dto.Status)
	}
	// ... update other fields ...

	// Transaction: Update + event
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

func (s *Service) DeleteJob(ctx context.Context, id uuid.UUID, clientID uuid.UUID) error {
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

// Continue with HTTP handlers, router, main.go, config, etc.
// Similar to users-be structure