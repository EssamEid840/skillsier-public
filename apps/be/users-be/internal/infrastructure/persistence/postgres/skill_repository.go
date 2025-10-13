// internal/infrastructure/persistence/postgres/skill_repository.go
package postgres

import (
    "context"
    "gorm.io/gorm"
    "users-be/internal/domain/skill"
)

type skillRepository struct {
    db *gorm.DB
}

func NewSkillRepository(db *gorm.DB) skill.Repository {
    return &skillRepository{db: db}
}

func (r *skillRepository) Create(ctx context.Context, s *skill.Skill) error {
    return r.db.WithContext(ctx).Create(s).Error
}

func (r *skillRepository) CreateBatch(ctx context.Context, skills []*skill.Skill) error {
    return r.db.WithContext(ctx).CreateInBatches(skills, 50).Error
}

func (r *skillRepository) Update(ctx context.Context, s *skill.Skill) error {
    return r.db.WithContext(ctx).Save(s).Error
}

func (r *skillRepository) FindByID(ctx context.Context, id string) (*skill.Skill, error) {
    var s skill.Skill
    err := r.db.WithContext(ctx).Where("id = ?", id).First(&s).Error
    if err != nil {
        return nil, err
    }
    return &s, nil
}

func (r *skillRepository) FindByUserID(ctx context.Context, userID string) ([]*skill.Skill, error) {
    var skills []*skill.Skill
    err := r.db.WithContext(ctx).
        Where("user_id = ?", userID).
        Order("is_primary DESC, display_order ASC, proficiency DESC, created_at DESC").
        Find(&skills).Error
    return skills, err
}

func (r *skillRepository) FindByUserIDAndName(ctx context.Context, userID, skillName string) (*skill.Skill, error) {
    var s skill.Skill
    err := r.db.WithContext(ctx).
        Where("user_id = ? AND skill_name = ?", userID, skillName).
        First(&s).Error
    if err != nil {
        return nil, err
    }
    return &s, nil
}

func (r *skillRepository) FindPrimarySkills(ctx context.Context, userID string) ([]*skill.Skill, error) {
    var skills []*skill.Skill
    err := r.db.WithContext(ctx).
        Where("user_id = ? AND is_primary = ?", userID, true).
        Order("display_order ASC").
        Find(&skills).Error
    return skills, err
}

func (r *skillRepository) FindVerifiedSkills(ctx context.Context, userID string) ([]*skill.Skill, error) {
    var skills []*skill.Skill
    err := r.db.WithContext(ctx).
        Where("user_id = ? AND is_verified = ?", userID, true).
        Order("proficiency DESC, years_of_experience DESC").
        Find(&skills).Error
    return skills, err
}

func (r *skillRepository) Search(ctx context.Context, query string) ([]*skill.Skill, error) {
    var skills []*skill.Skill
    pattern := "%" + query + "%"
    err := r.db.WithContext(ctx).
        Where("skill_name ILIKE ?", pattern).
        Limit(100).
        Find(&skills).Error
    return skills, err
}

func (r *skillRepository) FindByCategory(ctx context.Context, categoryID string) ([]*skill.Skill, error) {
    var skills []*skill.Skill
    err := r.db.WithContext(ctx).
        Where("skill_category_id = ?", categoryID).
        Find(&skills).Error
    return skills, err
}

func (r *skillRepository) Delete(ctx context.Context, id string) error {
    return r.db.WithContext(ctx).Delete(&skill.Skill{}, "id = ?", id).Error
}

func (r *skillRepository) DeleteByUserID(ctx context.Context, userID string) error {
    return r.db.WithContext(ctx).Where("user_id = ?", userID).Delete(&skill.Skill{}).Error
}

func (r *skillRepository) UpdateDisplayOrder(ctx context.Context, userID string, skillIDs []string) error {
    return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
        for i, skillID := range skillIDs {
            if err := tx.Model(&skill.Skill{}).
                Where("id = ? AND user_id = ?", skillID, userID).
                Update("display_order", i).Error; err != nil {
                return err
            }
        }
        return nil
    })
}

func (r *skillRepository) IncrementEndorsements(ctx context.Context, id string) error {
    return r.db.WithContext(ctx).
        Model(&skill.Skill{}).
        Where("id = ?", id).
        Update("endorsement_count", gorm.Expr("endorsement_count + ?", 1)).Error
}

func (r *skillRepository) IncrementProjectCount(ctx context.Context, id string) error {
    return r.db.WithContext(ctx).
        Model(&skill.Skill{}).
        Where("id = ?", id).
        Update("project_count", gorm.Expr("project_count + ?", 1)).Error
}

func (r *skillRepository) GetTopSkills(ctx context.Context, limit int) ([]*skill.Skill, error) {
    var skills []*skill.Skill
    err := r.db.WithContext(ctx).
        Select("skill_name, COUNT(*) as count").
        Group("skill_name").
        Order("count DESC").
        Limit(limit).
        Find(&skills).Error
    return skills, err
}

func (r *skillRepository) GetSkillStats(ctx context.Context, skillName string) (map[string]interface{}, error) {
    var stats struct {
        TotalUsers      int64
        AvgExperience   float64
        VerifiedCount   int64
        BeginnerCount   int64
        ExpertCount     int64
    }
    
    err := r.db.WithContext(ctx).Model(&skill.Skill{}).
        Select(`
            COUNT(DISTINCT user_id) as total_users,
            AVG(years_of_experience) as avg_experience,
            COUNT(CASE WHEN is_verified = true THEN 1 END) as verified_count,
            COUNT(CASE WHEN proficiency = 'beginner' THEN 1 END) as beginner_count,
            COUNT(CASE WHEN proficiency = 'expert' THEN 1 END) as expert_count
        `).
        Where("skill_name = ?", skillName).
        Scan(&stats).Error
    
    if err != nil {
        return nil, err
    }
    
    return map[string]interface{}{
        "total_users":      stats.TotalUsers,
        "avg_experience":   stats.AvgExperience,
        "verified_count":   stats.VerifiedCount,
        "beginner_count":   stats.BeginnerCount,
        "expert_count":     stats.ExpertCount,
    }, nil
}

func (r *skillRepository) CountByUser(ctx context.Context, userID string) (int64, error) {
    var count int64
    err := r.db.WithContext(ctx).
        Model(&skill.Skill{}).
        Where("user_id = ?", userID).
        Count(&count).Error
    return count, err
}