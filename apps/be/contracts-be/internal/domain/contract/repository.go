package contract

import (
	"context"
	"errors"
	"github.com/google/uuid"
)

var (
	ErrContractNotFound  = errors.New("contract not found")
	ErrMilestoneNotFound = errors.New("milestone not found")
	ErrUnauthorized      = errors.New("unauthorized access")
	ErrTitleRequired     = errors.New("title is required")
	ErrInvalidAmount     = errors.New("amount must be greater than zero")
	ErrInvalidStatus     = errors.New("invalid status transition")
)

type Repository interface {
	Create(ctx context.Context, contract *Contract) error
	GetByID(ctx context.Context, id uuid.UUID) (*Contract, error)
	List(ctx context.Context, filters *ListFilters, limit, offset int) ([]*Contract, int64, error)
	Update(ctx context.Context, contract *Contract) error
	GetByFreelancerID(ctx context.Context, freelancerID uuid.UUID, limit, offset int) ([]*Contract, int64, error)
	GetByClientID(ctx context.Context, clientID uuid.UUID, limit, offset int) ([]*Contract, int64, error)
	
	// Milestone operations
	GetMilestone(ctx context.Context, milestoneID uuid.UUID) (*Milestone, error)
	UpdateMilestone(ctx context.Context, milestone *Milestone) error
}

type ListFilters struct {
	ClientID     *uuid.UUID
	FreelancerID *uuid.UUID
	Status       *ContractStatus
}