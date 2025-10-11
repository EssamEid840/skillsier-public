package contract

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"contracts-be/internal/domain/contract"
	"contracts-be/internal/domain/outbox"
)

type Service struct {
	contractRepo contract.Repository
	outboxRepo   outbox.Repository
	db           *gorm.DB
}

func NewService(contractRepo contract.Repository, outboxRepo outbox.Repository, db *gorm.DB) *Service {
	return &Service{contractRepo: contractRepo, outboxRepo: outboxRepo, db: db}
}

func (s *Service) GetMyContracts(ctx context.Context, userID uuid.UUID, userType string) ([]*ContractResponseDTO, error) {
	var contracts []*contract.Contract
	var err error

	if userType == "freelancer" {
		contracts, err = s.contractRepo.GetByFreelancerID(ctx, userID)
	} else {
		contracts, err = s.contractRepo.GetByClientID(ctx, userID)
	}

	if err != nil {
		return nil, err
	}

	dtos := make([]*ContractResponseDTO, len(contracts))
	for i, c := range contracts {
		dtos[i] = ToResponseDTO(c)
	}
	return dtos, nil
}

func (s *Service) SubmitMilestone(ctx context.Context, contractID, milestoneID, freelancerID uuid.UUID) error {
	c, err := s.contractRepo.GetByID(ctx, contractID)
	if err != nil {
		return err
	}
	if c.FreelancerID != freelancerID {
		return contract.ErrUnauthorized
	}

	var milestone *contract.ContractMilestone
	for i := range c.Milestones {
		if c.Milestones[i].ID == milestoneID {
			milestone = &c.Milestones[i]
			break
		}
	}
	if milestone == nil {
		return contract.ErrMilestoneNotFound
	}

	now := time.Now()
	milestone.Status = contract.MilestoneStatusSubmitted
	milestone.SubmittedAt = &now

	return s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := s.contractRepo.UpdateMilestone(ctx, milestone); err != nil {
			return err
		}
		event, _ := s.createContractEvent("milestone.submitted", c)
		return s.outboxRepo.Create(ctx, event)
	})
}

func (s *Service) ApproveMilestone(ctx context.Context, contractID, milestoneID, clientID uuid.UUID) error {
	c, err := s.contractRepo.GetByID(ctx, contractID)
	if err != nil {
		return err
	}
	if c.ClientID != clientID {
		return contract.ErrUnauthorized
	}

	var milestone *contract.ContractMilestone
	for i := range c.Milestones {
		if c.Milestones[i].ID == milestoneID {
			milestone = &c.Milestones[i]
			break
		}
	}
	if milestone == nil {
		return contract.ErrMilestoneNotFound
	}

	now := time.Now()
	milestone.Status = contract.MilestoneStatusApproved
	milestone.ApprovedAt = &now

	return s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := s.contractRepo.UpdateMilestone(ctx, milestone); err != nil {
			return err
		}
		event, _ := s.createContractEvent("milestone.approved", c)
		return s.outboxRepo.Create(ctx, event)
	})
}

func (s *Service) CompleteContract(ctx context.Context, contractID, clientID uuid.UUID) error {
	c, err := s.contractRepo.GetByID(ctx, contractID)
	if err != nil {
		return err
	}
	if c.ClientID != clientID {
		return contract.ErrUnauthorized
	}

	now := time.Now()
	c.Status = contract.ContractStatusCompleted
	c.EndDate = &now

	return s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := s.contractRepo.Update(ctx, c); err != nil {
			return err
		}
		event, _ := s.createContractEvent("contract.completed", c)
		return s.outboxRepo.Create(ctx, event)
	})
}

func (s *Service) createContractEvent(eventType string, c *contract.Contract) (*outbox.Event, error) {
	payload := map[string]interface{}{
		"contract_id":   c.ID.String(),
		"job_id":        c.JobID.String(),
		"client_id":     c.ClientID.String(),
		"freelancer_id": c.FreelancerID.String(),
		"status":        string(c.Status),
	}
	payloadBytes, _ := json.Marshal(payload)
	metadata := map[string]interface{}{"source": "contracts-be"}
	metadataBytes, _ := json.Marshal(metadata)
	return &outbox.Event{
		AggregateID:   c.ID.String(),
		AggregateType: "contract",
		EventType:     eventType,
		Payload:       payloadBytes,
		Metadata:      metadataBytes,
	}, nil
}

type ContractResponseDTO struct {
	ID           uuid.UUID               `json:"id"`
	JobID        uuid.UUID               `json:"job_id"`
	ClientID     uuid.UUID               `json:"client_id"`
	FreelancerID uuid.UUID               `json:"freelancer_id"`
	TotalAmount  float64                 `json:"total_amount"`
	Status       contract.ContractStatus `json:"status"`
	StartDate    time.Time               `json:"start_date"`
	EndDate      *time.Time              `json:"end_date"`
	Milestones   []MilestoneDTO          `json:"milestones"`
}

type MilestoneDTO struct {
	ID          uuid.UUID                `json:"id"`
	Description string                   `json:"description"`
	Amount      float64                  `json:"amount"`
	DueDate     time.Time                `json:"due_date"`
	Status      contract.MilestoneStatus `json:"status"`
}

func ToResponseDTO(c *contract.Contract) *ContractResponseDTO {
	milestones := make([]MilestoneDTO, len(c.Milestones))
	for i, m := range c.Milestones {
		milestones[i] = MilestoneDTO{
			ID:          m.ID,
			Description: m.Description,
			Amount:      m.Amount,
			DueDate:     m.DueDate,
			Status:      m.Status,
		}
	}
	return &ContractResponseDTO{
		ID:           c.ID,
		JobID:        c.JobID,
		ClientID:     c.ClientID,
		FreelancerID: c.FreelancerID,
		TotalAmount:  c.TotalAmount,
		Status:       c.Status,
		StartDate:    c.StartDate,
		EndDate:      c.EndDate,
		Milestones:   milestones,
	}
}
