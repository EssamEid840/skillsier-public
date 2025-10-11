package postgres

import (
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"contracts-be/internal/domain/contract"
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
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, contract.ErrContractNotFound
		}
		return nil, err
	}
	return &c, nil
}

func (r *contractRepository) GetByFreelancerID(ctx context.Context, freelancerID uuid.UUID) ([]*contract.Contract, error) {
	var contracts []*contract.Contract
	err := r.db.WithContext(ctx).Preload("Milestones").Where("freelancer_id = ?", freelancerID).Order("created_at DESC").Find(&contracts).Error
	return contracts, err
}

func (r *contractRepository) GetByClientID(ctx context.Context, clientID uuid.UUID) ([]*contract.Contract, error) {
	var contracts []*contract.Contract
	err := r.db.WithContext(ctx).Preload("Milestones").Where("client_id = ?", clientID).Order("created_at DESC").Find(&contracts).Error
	return contracts, err
}

func (r *contractRepository) Update(ctx context.Context, c *contract.Contract) error {
	return r.db.WithContext(ctx).Save(c).Error
}

func (r *contractRepository) UpdateMilestone(ctx context.Context, m *contract.ContractMilestone) error {
	return r.db.WithContext(ctx).Save(m).Error
}