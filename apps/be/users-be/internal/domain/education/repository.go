// internal/domain/education/repository.go
package education

import "context"

type Repository interface {
    Create(ctx context.Context, edu *Education) error
    CreateBatch(ctx context.Context, educations []*Education) error
    Update(ctx context.Context, edu *Education) error
    FindByID(ctx context.Context, id string) (*Education, error)
    FindByUserID(ctx context.Context, userID string) ([]*Education, error)
    FindVerified(ctx context.Context, userID string) ([]*Education, error)
    FindCurrent(ctx context.Context, userID string) ([]*Education, error)
    FindBySchool(ctx context.Context, school string) ([]*Education, error)
    Delete(ctx context.Context, id string) error
    DeleteByUserID(ctx context.Context, userID string) error
    UpdateDisplayOrder(ctx context.Context, userID string, eduIDs []string) error
    CountByUser(ctx context.Context, userID string) (int64, error)
    GetHighestDegree(ctx context.Context, userID string) (*Education, error)
}