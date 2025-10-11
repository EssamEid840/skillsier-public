package proposal

import (
	"context"
	"encoding/json"
	"fmt"
	"proposals-be/internal/domain/proposal"
	"proposals-be/internal/domain/outbox"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Service struct {
	proposalRepo proposal.Repository
	outboxRepo   outbox.Repository
	db           *gorm.DB
}

func NewService(proposalRepo proposal.Repository, outboxRepo outbox.Repository, db *gorm.DB) *Service {
	return &Service{
		proposalRepo: proposalRepo,
		outboxRepo:   outboxRepo,
		db:           db,
	}
}

func (s *Service) CreateProposal(ctx context.Context, freelancerID uuid.UUID, dto *CreateProposalDTO) (*ProposalResponseDTO, error) {
	// Check if already submitted for this job
	exists, err := s.proposalRepo.CheckExisting(ctx, dto.JobID, freelancerID)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, proposal.ErrJobAlreadySubmitted
	}

	newProposal := &proposal.Proposal{
		JobID:             dto.JobID,
		FreelancerID:      freelancerID,
		CoverLetter:       dto.CoverLetter,
		BidAmount:         dto.BidAmount,
		EstimatedDuration: dto.EstimatedDuration,
		Status:            proposal.ProposalStatusPending,
	}

	// Add milestones if provided
	for _, m := range dto.Milestones {
		newProposal.Milestones = append(newProposal.Milestones, proposal.ProposalMilestone{
			Description: m.Description,
			Amount:      m.Amount,
			DueDate:     m.DueDate,
		})
	}

	if err := newProposal.Validate(); err != nil {
		return nil, err
	}

	// Transaction: Create proposal + outbox event
	err = s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := s.proposalRepo.Create(ctx, newProposal); err != nil {
			return err
		}

		event, err := s.createProposalEvent("proposal.created", newProposal)
		if err != nil {
			return err
		}

		return s.outboxRepo.Create(ctx, event)
	})

	if err != nil {
		return nil, err
	}

	return ToResponseDTO(newProposal), nil
}

func (s *Service) GetProposal(ctx context.Context, id uuid.UUID) (*ProposalResponseDTO, error) {
	p, err := s.proposalRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return ToResponseDTO(p), nil
}

func (s *Service) ListProposals(ctx context.Context, filters *proposal.ListFilters, page, pageSize int) (*ListProposalsResponseDTO, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 20
	}

	offset := (page - 1) * pageSize

	proposals, total, err := s.proposalRepo.List(ctx, filters, pageSize, offset)
	if err != nil {
		return nil, err
	}

	return ToListResponse(proposals, total, page, pageSize), nil
}

func (s *Service) UpdateProposal(ctx context.Context, id uuid.UUID, freelancerID uuid.UUID, dto *UpdateProposalDTO) (*ProposalResponseDTO, error) {
	p, err := s.proposalRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	// Verify ownership
	if p.FreelancerID != freelancerID {
		return nil, proposal.ErrUnauthorized
	}

	// Can only update pending proposals
	if p.Status != proposal.ProposalStatusPending {
		return nil, fmt.Errorf("can only update pending proposals")
	}

	// Update fields
	if dto.CoverLetter != nil {
		p.CoverLetter = *dto.CoverLetter
	}
	if dto.BidAmount != nil {
		p.BidAmount = *dto.BidAmount
	}
	if dto.EstimatedDuration != nil {
		p.EstimatedDuration = *dto.EstimatedDuration
	}

	// Transaction: Update + event
	err = s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := s.proposalRepo.Update(ctx, p); err != nil {
			return err
		}

		event, err := s.createProposalEvent("proposal.updated", p)
		if err != nil {
			return err
		}

		return s.outboxRepo.Create(ctx, event)
	})

	if err != nil {
		return nil, err
	}

	return ToResponseDTO(p), nil
}

func (s *Service) WithdrawProposal(ctx context.Context, id uuid.UUID, freelancerID uuid.UUID) error {
	p, err := s.proposalRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}

	if p.FreelancerID != freelancerID {
		return proposal.ErrUnauthorized
	}

	if p.Status != proposal.ProposalStatusPending {
		return fmt.Errorf("can only withdraw pending proposals")
	}

	p.Status = proposal.ProposalStatusWithdrawn

	return s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := s.proposalRepo.Update(ctx, p); err != nil {
			return err
		}

		event, err := s.createProposalEvent("proposal.withdrawn", p)
		if err != nil {
			return err
		}

		return s.outboxRepo.Create(ctx, event)
	})
}

func (s *Service) GetMyProposals(ctx context.Context, freelancerID uuid.UUID, page, pageSize int) (*ListProposalsResponseDTO, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 20
	}

	offset := (page - 1) * pageSize

	proposals, total, err := s.proposalRepo.GetByFreelancerID(ctx, freelancerID, pageSize, offset)
	if err != nil {
		return nil, err
	}

	return ToListResponse(proposals, total, page, pageSize), nil
}

func (s *Service) createProposalEvent(eventType string, p *proposal.Proposal) (*outbox.Event, error) {
	payload := map[string]interface{}{
		"proposal_id":    p.ID.String(),
		"job_id":         p.JobID.String(),
		"freelancer_id":  p.FreelancerID.String(),
		"bid_amount":     p.BidAmount,
		"status":         string(p.Status),
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal payload: %w", err)
	}

	metadata := map[string]interface{}{"source": "proposals-be"}
	metadataBytes, err := json.Marshal(metadata)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal metadata: %w", err)
	}

	return &outbox.Event{
		AggregateID:   p.ID.String(),
		AggregateType: "proposal",
		EventType:     eventType,
		Payload:       payloadBytes,
		Metadata:      metadataBytes,
	}, nil
}
