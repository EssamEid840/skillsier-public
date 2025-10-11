package proposal

import (
	"context"
	"encoding/json"
	"fmt"
	"gorm.io/gorm"
	"proposals-be/internal/domain/outbox"
)

type Service struct {
	proposalRepo proposal.Repository
	outboxRepo   outbox.Repository
	db           *gorm.DB
}

func NewService(proposalRepo proposal.Repository, outboxRepo outbox.Repository, db *gorm.DB) *Service {
	return &Service{proposalRepo: proposalRepo, outboxRepo: outboxRepo, db: db}
}

func (s *Service) CreateProposal(ctx context.Context, freelancerID uuid.UUID, dto *CreateProposalDTO) (*ProposalResponseDTO, error) {
	duplicate, err := s.proposalRepo.CheckDuplicate(ctx, dto.JobID, freelancerID)
	if err != nil {
		return nil, err
	}
	if duplicate {
		return nil, proposal.ErrDuplicateProposal
	}

	p := &proposal.Proposal{
		JobID:        dto.JobID,
		FreelancerID: freelancerID,
		CoverLetter:  dto.CoverLetter,
		ProposedRate: dto.ProposedRate,
		DeliveryTime: dto.DeliveryTime,
		Status:       proposal.ProposalStatusPending,
	}

	for _, m := range dto.Milestones {
		p.Milestones = append(p.Milestones, proposal.ProposalMilestone{
			Description: m.Description,
			Amount:      m.Amount,
			DueDate:     m.DueDate,
		})
	}

	err = s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := s.proposalRepo.Create(ctx, p); err != nil {
			return err
		}
		event, _ := s.createProposalEvent("proposal.created", p)
		return s.outboxRepo.Create(ctx, event)
	})

	if err != nil {
		return nil, err
	}
	return ToResponseDTO(p), nil
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

func (s *Service) WithdrawProposal(ctx context.Context, id, freelancerID uuid.UUID) error {
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
		event, _ := s.createProposalEvent("proposal.withdrawn", p)
		return s.outboxRepo.Create(ctx, event)
	})
}

func (s *Service) createProposalEvent(eventType string, p *proposal.Proposal) (*outbox.Event, error) {
	payload := map[string]interface{}{
		"proposal_id":   p.ID.String(),
		"job_id":        p.JobID.String(),
		"freelancer_id": p.FreelancerID.String(),
		"status":        string(p.Status),
	}
	payloadBytes, _ := json.Marshal(payload)
	metadata := map[string]interface{}{"source": "proposals-be"}
	metadataBytes, _ := json.Marshal(metadata)
	return &outbox.Event{
		AggregateID:   p.ID.String(),
		AggregateType: "proposal",
		EventType:     eventType,
		Payload:       payloadBytes,
		Metadata:      metadataBytes,
	}, nil
}

func ToResponseDTO(p *proposal.Proposal) *ProposalResponseDTO {
	milestones := make([]ProposalMilestoneDTO, len(p.Milestones))
	for i, m := range p.Milestones {
		milestones[i] = ProposalMilestoneDTO{
			Description: m.Description,
			Amount:      m.Amount,
			DueDate:     m.DueDate,
		}
	}
	return &ProposalResponseDTO{
		ID:           p.ID,
		JobID:        p.JobID,
		FreelancerID: p.FreelancerID,
		CoverLetter:  p.CoverLetter,
		ProposedRate: p.ProposedRate,
		DeliveryTime: p.DeliveryTime,
		Status:       p.Status,
		Milestones:   milestones,
		CreatedAt:    p.CreatedAt,
		UpdatedAt:    p.UpdatedAt,
	}
}

func ToListResponse(proposals []*proposal.Proposal, total int64, page, pageSize int) *ListProposalsResponseDTO {
	dtos := make([]*ProposalResponseDTO, len(proposals))
	for i, p := range proposals {
		dtos[i] = ToResponseDTO(p)
	}
	totalPages := int(total) / pageSize
	if int(total)%pageSize > 0 {
		totalPages++
	}
	return &ListProposalsResponseDTO{
		Proposals:  dtos,
		Total:      total,
		Page:       page,
		PageSize:   pageSize,
		TotalPages: totalPages,
	}
}