package postgres

import (
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"proposals-be/internal/domain/proposal"
)

type proposalRepository struct {
	db *gorm.DB
}

func NewProposalRepository(db *gorm.DB) proposal.Repository {
	return &proposalRepository{db: db}
}

func (r *proposalRepository) Create(ctx context.Context, p *proposal.Proposal) error {
	return r.db.WithContext(ctx).Create(p).Error
}

func (r *proposalRepository) GetByID(ctx context.Context, id uuid.UUID) (*proposal.Proposal, error) {
	var p proposal.Proposal
	err := r.db.WithContext(ctx).Preload("Milestones").Where("id = ?", id).First(&p).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, proposal.ErrProposalNotFound
		}
		return nil, err
	}
	return &p, nil
}

func (r *proposalRepository) GetByJobID(ctx context.Context, jobID uuid.UUID) ([]*proposal.Proposal, error) {
	var proposals []*proposal.Proposal
	err := r.db.WithContext(ctx).
		Preload("Milestones").
		Where("job_id = ?", jobID).
		Order("created_at DESC").
		Find(&proposals).Error
	return proposals, err
}

func (r *proposalRepository) GetByFreelancerID(ctx context.Context, freelancerID uuid.UUID, limit, offset int) ([]*proposal.Proposal, int64, error) {
	var proposals []*proposal.Proposal
	var total int64

	query := r.db.WithContext(ctx).
		Model(&proposal.Proposal{}).
		Preload("Milestones").
		Where("freelancer_id = ?", freelancerID)

	query.Count(&total)

	err := query.Order("created_at DESC").
		Limit(limit).
		Offset(offset).
		Find(&proposals).Error

	return proposals, total, err
}

func (r *proposalRepository) Update(ctx context.Context, p *proposal.Proposal) error {
	return r.db.WithContext(ctx).Save(p).Error
}

func (r *proposalRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&proposal.Proposal{}, "id = ?", id).Error
}

func (r *proposalRepository) CheckDuplicate(ctx context.Context, jobID, freelancerID uuid.UUID) (bool, error) {
	var count int64
	err := r.db.WithContext(ctx).
		Model(&proposal.Proposal{}).
		Where("job_id = ? AND freelancer_id = ? AND status != ?", jobID, freelancerID, proposal.ProposalStatusWithdrawn).
		Count(&count).Error
	return count > 0, err
}