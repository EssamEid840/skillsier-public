package postgres

import (
	"context"
	"proposals-be/internal/domain/proposal"
	"github.com/google/uuid"
	"gorm.io/gorm"
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
	if err == gorm.ErrRecordNotFound {
		return nil, proposal.ErrProposalNotFound
	}
	return &p, err
}

func (r *proposalRepository) List(ctx context.Context, filters *proposal.ListFilters, limit, offset int) ([]*proposal.Proposal, int64, error) {
	var proposals []*proposal.Proposal
	var total int64

	query := r.db.WithContext(ctx).Model(&proposal.Proposal{}).Preload("Milestones")

	if filters != nil {
		if filters.JobID != nil {
			query = query.Where("job_id = ?", *filters.JobID)
		}
		if filters.FreelancerID != nil {
			query = query.Where("freelancer_id = ?", *filters.FreelancerID)
		}
		if filters.Status != nil {
			query = query.Where("status = ?", *filters.Status)
		}
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := query.Order("created_at DESC").Limit(limit).Offset(offset).Find(&proposals).Error
	return proposals, total, err
}

func (r *proposalRepository) Update(ctx context.Context, p *proposal.Proposal) error {
	return r.db.WithContext(ctx).Model(p).Updates(p).Error
}

func (r *proposalRepository) Delete(ctx context.Context, id uuid.UUID) error {
	result := r.db.WithContext(ctx).Delete(&proposal.Proposal{}, "id = ?", id)
	if result.RowsAffected == 0 {
		return proposal.ErrProposalNotFound
	}
	return result.Error
}

func (r *proposalRepository) GetByJobID(ctx context.Context, jobID uuid.UUID, limit, offset int) ([]*proposal.Proposal, int64, error) {
	var proposals []*proposal.Proposal
	var total int64

	query := r.db.WithContext(ctx).Model(&proposal.Proposal{}).Where("job_id = ?", jobID).Preload("Milestones")

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := query.Order("created_at DESC").Limit(limit).Offset(offset).Find(&proposals).Error
	return proposals, total, err
}

func (r *proposalRepository) GetByFreelancerID(ctx context.Context, freelancerID uuid.UUID, limit, offset int) ([]*proposal.Proposal, int64, error) {
	var proposals []*proposal.Proposal
	var total int64

	query := r.db.WithContext(ctx).Model(&proposal.Proposal{}).Where("freelancer_id = ?", freelancerID).Preload("Milestones")

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := query.Order("created_at DESC").Limit(limit).Offset(offset).Find(&proposals).Error
	return proposals, total, err
}

func (r *proposalRepository) CheckExisting(ctx context.Context, jobID uuid.UUID, freelancerID uuid.UUID) (bool, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&proposal.Proposal{}).
		Where("job_id = ? AND freelancer_id = ?", jobID, freelancerID).
		Count(&count).Error
	return count > 0, err
}