// internal/infrastructure/persistence/postgres/education_repository.go
package postgres

import (
    "context"
    "gorm.io/gorm"
    "users-be/internal/domain/education"
)

type educationRepository struct {
    db *gorm.DB
}

func NewEducationRepository(db *gorm.DB) education.Repository {
    return &educationRepository{db: db}
}

func (r *educationRepository) Create(ctx context.Context, edu *education.Education) error {
    return r.db.WithContext(ctx).Create(edu).Error
}

func (r *educationRepository) CreateBatch(ctx context.Context, educations []*education.Education) error {
    return r.db.WithContext(ctx).CreateInBatches(educations, 50).Error
}

func (r *educationRepository) Update(ctx context.Context, edu *education.Education) error {
    return r.db.WithContext(ctx).Save(edu).Error
}

func (r *educationRepository) FindByID(ctx context.Context, id string) (*education.Education, error) {
    var edu education.Education
    err := r.db.WithContext(ctx).Where("id = ?", id).First(&edu).Error
    if err != nil {
        if err == gorm.ErrRecordNotFound {
            return nil, education.ErrEducationNotFound
        }
        return nil, err
    }
    return &edu, nil
}

func (r *educationRepository) FindByUserID(ctx context.Context, userID string) ([]*education.Education, error) {
    var educations []*education.Education
    err := r.db.WithContext(ctx).
        Where("user_id = ?", userID).
        Order("is_current DESC, graduation_year DESC, display_order ASC").
        Find(&educations).Error
    return educations, err
}

func (r *educationRepository) FindVerified(ctx context.Context, userID string) ([]*education.Education, error) {
    var educations []*education.Education
    err := r.db.WithContext(ctx).
        Where("user_id = ? AND is_verified = ?", userID, true).
        Order("graduation_year DESC").
        Find(&educations).Error
    return educations, err
}

func (r *educationRepository) FindCurrent(ctx context.Context, userID string) ([]*education.Education, error) {
    var educations []*education.Education
    err := r.db.WithContext(ctx).
        Where("user_id = ? AND is_current = ?", userID, true).
        Find(&educations).Error
    return educations, err
}

func (r *educationRepository) FindBySchool(ctx context.Context, school string) ([]*education.Education, error) {
    var educations []*education.Education
    err := r.db.WithContext(ctx).
        Where("school ILIKE ?", "%"+school+"%").
        Order("graduation_year DESC").
        Find(&educations).Error
    return educations, err
}

func (r *educationRepository) Delete(ctx context.Context, id string) error {
    return r.db.WithContext(ctx).Delete(&education.Education{}, "id = ?", id).Error
}

func (r *educationRepository) DeleteByUserID(ctx context.Context, userID string) error {
    return r.db.WithContext(ctx).Where("user_id = ?", userID).Delete(&education.Education{}).Error
}

func (r *educationRepository) UpdateDisplayOrder(ctx context.Context, userID string, eduIDs []string) error {
    return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
        for i, eduID := range eduIDs {
            if err := tx.Model(&education.Education{}).
                Where("id = ? AND user_id = ?", eduID, userID).
                Update("display_order", i).Error; err != nil {
                return err
            }
        }
        return nil
    })
}

func (r *educationRepository) CountByUser(ctx context.Context, userID string) (int64, error) {
    var count int64
    err := r.db.WithContext(ctx).
        Model(&education.Education{}).
        Where("user_id = ?", userID).
        Count(&count).Error
    return count, err
}

func (r *educationRepository) GetHighestDegree(ctx context.Context, userID string) (*education.Education, error) {
    var edu education.Education
    degreeOrder := map[string]int{
        "doctorate":  5,
        "master":     4,
        "bachelor":   3,
        "associate":  2,
        "certificate": 1,
    }
    
    err := r.db.WithContext(ctx).
        Where("user_id = ?", userID).
        Order("CASE degree_type " +
            "WHEN 'doctorate' THEN 5 " +
            "WHEN 'master' THEN 4 " +
            "WHEN 'bachelor' THEN 3 " +
            "WHEN 'associate' THEN 2 " +
            "WHEN 'certificate' THEN 1 " +
            "ELSE 0 END DESC").
        First(&edu).Error
    
    if err != nil {
        if err == gorm.ErrRecordNotFound {
            return nil, education.ErrEducationNotFound
        }
        return nil, err
    }
    return &edu, nil
}