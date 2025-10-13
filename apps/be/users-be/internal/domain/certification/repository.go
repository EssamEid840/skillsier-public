// internal/domain/certification/repository.go
package certification

import "context"

type Repository interface {
    Create(ctx context.Context, cert *Certification) error
    CreateBatch(ctx context.Context, certs []*Certification) error
    Update(ctx context.Context, cert *Certification) error
    FindByID(ctx context.Context, id string) (*Certification, error)
    FindByUserID(ctx context.Context, userID string) ([]*Certification, error)
    FindVerified(ctx context.Context, userID string) ([]*Certification, error)
    FindPending(ctx context.Context, userID string) ([]*Certification, error)
    FindExpired(ctx context.Context, userID string) ([]*Certification, error)
    FindExpiringSoon(ctx context.Context, days int) ([]*Certification, error)
    FindByOrganization(ctx context.Context, org string) ([]*Certification, error)
    FindByCredentialID(ctx context.Context, credentialID string) (*Certification, error)
    Delete(ctx context.Context, id string) error
    DeleteByUserID(ctx context.Context, userID string) error
    UpdateDisplayOrder(ctx context.Context, userID string, certIDs []string) error
    UpdateVerificationStatus(ctx context.Context, id string, status VerificationStatus, verifiedBy string) error
    IncrementViews(ctx context.Context, id string) error
    IncrementEndorsements(ctx context.Context, id string) error
    CountByUser(ctx context.Context, userID string) (int64, error)
    CountVerifiedByUser(ctx context.Context, userID string) (int64, error)
    GetTopCertifications(ctx context.Context, limit int) ([]*Certification, error)
}