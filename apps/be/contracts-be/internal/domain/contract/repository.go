package contract

import (
	"context"
	"errors"

	"github.com/google/uuid"
)

var (
	ErrContractNotFound  = errors.New("contract not found")
	ErrMilestoneNotFound = errors.New("milestone not found")
	ErrUnauthorized      = errors.New("unauthorized")
)

type Repository interface {
	Create(ctx context.Context, contract *Contract) error
	GetByID(ctx context.Context, id uuid.UUID) (*Contract, error)
	GetByFreelancerID(ctx context.Context, freelancerID uuid.UUID) ([]*Contract, error)
	GetByClientID(ctx context.Context, clientID uuid.UUID) ([]*Contract, error)
	Update(ctx context.Context, contract *Contract) error
	UpdateMilestone(ctx context.Context, milestone *ContractMilestone) error
}
