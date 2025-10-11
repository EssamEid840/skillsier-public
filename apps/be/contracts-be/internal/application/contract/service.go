package contract

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
	"contracts-be/internal/domain/contract"
	"contracts-be/internal/domain/outbox"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Service struct {
	contractRepo contract.Repository
	outboxRepo   outbox.Repository
	db           *gorm.DB
}

func NewService(contractRepo contract.Repository, outboxRepo outbox.Repository, db *gorm.DB) *Service {
	return &Service{
		contractRepo: contractRepo,
		outboxRepo:   outboxRepo,
		db:           db,
	}
}

func (s *Service) CreateContract(ctx context.Context, dto *CreateContractDTO) (*ContractResponseDTO, error) {
	now := time.Now()
	
	newContract := &contract.Contract{
		JobID:        dto.JobID,
		ProposalID:   dto.ProposalID,
		ClientID:     dto.ClientID,
		FreelancerID: dto.FreelancerID,
		Title:        dto.Title,
		Description:  dto.Description,
		TotalAmount:  dto.TotalAmount,
		Status:       contract.ContractStatusActive,
		StartDate:    now,
		Terms:        dto.Terms,
	}

	// Create milestones
	for _, m := range dto.Milestones {
		newContract.Milestones = append(newContract.Milestones, contract.Milestone{
			Description: m.Description,
			Amount:      m.Amount,
			DueDate:     m.DueDate,
			Status:      contract.MilestoneStatusPending,
			PaymentStatus: contract.PaymentStatusPending,
		})
	}

	if err := newContract.Validate(); err != nil {
		return nil, err
	}

	// Transaction: Create contract + outbox event
	err := s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := s.contractRepo.Create(ctx, newContract); err != nil {
			return err
		}

		event, err := s.createContractEvent("contract.created", newContract)
		if err != nil {
			return err
		}

		return s.outboxRepo.Create(ctx, event)
	})

	if err != nil {
		return nil, err
	}

	return ToResponseDTO(newContract), nil
}

func (s *Service) GetContract(ctx context.Context, id uuid.UUID) (*ContractResponseDTO, error) {
	c, err := s.contractRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return ToResponseDTO(c), nil
}

func (s *Service) GetMyContracts(ctx context.Context, userID uuid.UUID, role string, page, pageSize int) (*ListContractsResponseDTO, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 20
	}

	offset := (page - 1) * pageSize
	var contracts []*contract.Contract
	var total int64
	var err error

	if role == "freelancer" {
		contracts, total, err = s.contractRepo.GetByFreelancerID(ctx, userID, pageSize, offset)
	} else {
		contracts, total, err = s.contractRepo.GetByClientID(ctx, userID, pageSize, offset)
	}

	if err != nil {
		return nil, err
	}

	return ToListResponse(contracts, total, page, pageSize), nil
}

func (s *Service) SubmitMilestone(ctx context.Context, contractID uuid.UUID, milestoneID uuid.UUID, freelancerID uuid.UUID, dto *SubmitMilestoneDTO) (*MilestoneResponseDTO, error) {
	// Verify contract ownership
	c, err := s.contractRepo.GetByID(ctx, contractID)
	if err != nil {
		return nil, err
	}
	if c.FreelancerID != freelancerID {
		return nil, contract.ErrUnauthorized
	}

	// Get milestone
	milestone, err := s.contractRepo.GetMilestone(ctx, milestoneID)
	if err != nil {
		return nil, err
	}
	if milestone.ContractID != contractID {
		return nil, contract.ErrMilestoneNotFound
	}

	// Update milestone
	now := time.Now()
	milestone.Status = contract.MilestoneStatusSubmitted
	milestone.Deliverables = &dto.Deliverables
	milestone.SubmittedAt = &now

	err = s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := s.contractRepo.UpdateMilestone(ctx, milestone); err != nil {
			return err
		}

		event, err := s.createMilestoneEvent("milestone.submitted", milestone)
		if err != nil {
			return err
		}

		return s.outboxRepo.Create(ctx, event)
	})

	if err != nil {
		return nil, err
	}

	return ToMilestoneResponseDTO(milestone), nil
}

func (s *Service) ApproveMilestone(ctx context.Context, contractID uuid.UUID, milestoneID uuid.UUID, clientID uuid.UUID, dto *ApproveMilestoneDTO) (*MilestoneResponseDTO, error) {
	// Verify contract ownership
	c, err := s.contractRepo.GetByID(ctx, contractID)
	if err != nil {
		return nil, err
	}
	if c.ClientID != clientID {
		return nil, contract.ErrUnauthorized
	}

	// Get milestone
	milestone, err := s.contractRepo.GetMilestone(ctx, milestoneID)
	if err != nil {
		return nil, err
	}
	if milestone.ContractID != contractID {
		return nil, contract.ErrMilestoneNotFound
	}

	// Update milestone
	now := time.Now()
	milestone.Status = contract.MilestoneStatusApproved
	milestone.Feedback = dto.Feedback
	milestone.ApprovedAt = &now

	err = s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := s.contractRepo.UpdateMilestone(ctx, milestone); err != nil {
			return err
		}

		event, err := s.createMilestoneEvent("milestone.approved", milestone)
		if err != nil {
			return err
		}

		return s.outboxRepo.Create(ctx, event)
	})

	if err != nil {
		return nil, err
	}

	return ToMilestoneResponseDTO(milestone), nil
}

func (s *Service) CompleteContract(ctx context.Context, id uuid.UUID, clientID uuid.UUID) (*ContractResponseDTO, error) {
	c, err := s.contractRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if c.ClientID != clientID {
		return nil, contract.ErrUnauthorized
	}

	// Check all milestones are approved
	for _, m := range c.Milestones {
		if m.Status != contract.MilestoneStatusApproved && m.Status != contract.MilestoneStatusPaid {
			return nil, fmt.Errorf("all milestones must be approved before completing contract")
		}
	}

	now := time.Now()
	c.Status = contract.ContractStatusCompleted
	c.EndDate = &now

	err = s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := s.contractRepo.Update(ctx, c); err != nil {
			return err
		}

		event, err := s.createContractEvent("contract.completed", c)
		if err != nil {
			return err
		}

		return s.outboxRepo.Create(ctx, event)
	})

	if err != nil {
		return nil, err
	}

	return ToResponseDTO(c), nil
}

func (s *Service) createContractEvent(eventType string, c *contract.Contract) (*outbox.Event, error) {
	payload := map[string]interface{}{
		"contract_id":   c.ID.String(),
		"job_id":        c.JobID.String(),
		"proposal_id":   c.ProposalID.String(),
		"client_id":     c.ClientID.String(),
		"freelancer_id": c.FreelancerID.String(),
		"total_amount":  c.TotalAmount,
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

func (s *Service) createMilestoneEvent(eventType string, m *contract.Milestone) (*outbox.Event, error) {
	payload := map[string]interface{}{
		"milestone_id": m.ID.String(),
		"contract_id":  m.ContractID.String(),
		"amount":       m.Amount,
		"status":       string(m.Status),
	}

	payloadBytes, _ := json.Marshal(payload)
	metadata := map[string]interface{}{"source": "contracts-be"}
	metadataBytes, _ := json.Marshal(metadata)

	return &outbox.Event{
		AggregateID:   m.ContractID.String(),
		AggregateType: "milestone",
		EventType:     eventType,
		Payload:       payloadBytes,
		Metadata:      metadataBytes,
	}, nil
}
