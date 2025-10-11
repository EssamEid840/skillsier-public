package postgres

import (
	"context"
	"contracts-be/internal/domain/contract"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type contractRepository struct {
	db *gorm.DB
}

func NewContractRepository(db *gorm.DB) contract.Repository {
	return &contractRepository{db: db}
}

func (r *contractRepository) Create(ctx context.Context, c *contract.Contract) error {
	return r.db.WithContext(ctx).Create(c).Error
}

func (r *contractRepository) GetByID(ctx context.Context, id uuid.UUID) (*contract.Contract, error) {
	var c contract.Contract
	err := r.db.WithContext(ctx).Preload("Milestones").Where("id = ?", id).First(&c).Error
	if err == gorm.ErrRecordNotFound {
		return nil, contract.ErrContractNotFound
	}
	return &c, err
}

func (r *contractRepository) List(ctx context.Context, filters *contract.ListFilters, limit, offset int) ([]*contract.Contract, int64, error) {
	var contracts []*contract.Contract
	var total int64

	query := r.db.WithContext(ctx).Model(&contract.Contract{}).Preload("Milestones")

	if filters != nil {
		if filters.ClientID != nil {
			query = query.Where("client_id = ?", *filters.ClientID)
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

	err := query.Order("created_at DESC").Limit(limit).Offset(offset).Find(&contracts).Error
	return contracts, total, err
}

func (r *contractRepository) Update(ctx context.Context, c *contract.Contract) error {
	return r.db.WithContext(ctx).Model(c).Updates(c).Error
}

func (r *contractRepository) GetByFreelancerID(ctx context.Context, freelancerID uuid.UUID, limit, offset int) ([]*contract.Contract, int64, error) {
	var contracts []*contract.Contract
	var total int64

	query := r.db.WithContext(ctx).Model(&contract.Contract{}).
		Where("freelancer_id = ?", freelancerID).Preload("Milestones")

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := query.Order("created_at DESC").Limit(limit).Offset(offset).Find(&contracts).Error
	return contracts, total, err
}

func (r *contractRepository) GetByClientID(ctx context.Context, clientID uuid.UUID, limit, offset int) ([]*contract.Contract, int64, error) {
	var contracts []*contract.Contract
	var total int64

	query := r.db.WithContext(ctx).Model(&contract.Contract{}).
		Where("client_id = ?", clientID).Preload("Milestones")

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := query.Order("created_at DESC").Limit(limit).Offset(offset).Find(&contracts).Error
	return contracts, total, err
}

func (r *contractRepository) GetMilestone(ctx context.Context, milestoneID uuid.UUID) (*contract.Milestone, error) {
	var milestone contract.Milestone
	err := r.db.WithContext(ctx).Where("id = ?", milestoneID).First(&milestone).Error
	if err == gorm.ErrRecordNotFound {
		return nil, contract.ErrMilestoneNotFound
	}
	return &milestone, err
}

func (r *contractRepository) UpdateMilestone(ctx context.Context, milestone *contract.Milestone) error {
	return r.db.WithContext(ctx).Model(milestone).Updates(milestone).Error
}
