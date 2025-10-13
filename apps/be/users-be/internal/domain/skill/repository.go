// internal/domain/skill/repository.go
package skill

import "context"

type Repository interface {
    Create(ctx context.Context, skill *Skill) error
    CreateBatch(ctx context.Context, skills []*Skill) error
    Update(ctx context.Context, skill *Skill) error
    FindByID(ctx context.Context, id string) (*Skill, error)
    FindByUserID(ctx context.Context, userID string) ([]*Skill, error)
    FindByUserIDAndName(ctx context.Context, userID, skillName string) (*Skill, error)
    FindPrimarySkills(ctx context.Context, userID string) ([]*Skill, error)
    FindVerifiedSkills(ctx context.Context, userID string) ([]*Skill, error)
    Search(ctx context.Context, query string) ([]*Skill, error)
    FindByCategory(ctx context.Context, categoryID string) ([]*Skill, error)
    Delete(ctx context.Context, id string) error
    DeleteByUserID(ctx context.Context, userID string) error
    UpdateDisplayOrder(ctx context.Context, userID string, skillIDs []string) error
    IncrementEndorsements(ctx context.Context, id string) error
    IncrementProjectCount(ctx context.Context, id string) error
    GetTopSkills(ctx context.Context, limit int) ([]*Skill, error)
    GetSkillStats(ctx context.Context, skillName string) (map[string]interface{}, error)
    CountByUser(ctx context.Context, userID string) (int64, error)
}