// internal/infrastructure/persistence/postgres/profile_repository.go
package postgres

import (
    "context"
    "fmt"
    "time"
    "gorm.io/gorm"
    "users-be/internal/domain/profile"
)

type profileRepository struct {
    db *gorm.DB
}

func NewProfileRepository(db *gorm.DB) profile.Repository {
    return &profileRepository{db: db}
}

// CREATE
func (r *profileRepository) Create(ctx context.Context, p *profile.Profile) error {
    return r.db.WithContext(ctx).Create(p).Error
}

func (r *profileRepository) CreateBatch(ctx context.Context, profiles []*profile.Profile) error {
    return r.db.WithContext(ctx).CreateInBatches(profiles, 100).Error
}

// READ - Single
func (r *profileRepository) FindByID(ctx context.Context, id string) (*profile.Profile, error) {
    var p profile.Profile
    err := r.db.WithContext(ctx).Where("id = ?", id).First(&p).Error
    if err != nil {
        if err == gorm.ErrRecordNotFound {
            return nil, profile.ErrProfileNotFound
        }
        return nil, err
    }
    return &p, nil
}

func (r *profileRepository) FindByUserID(ctx context.Context, userID string) (*profile.Profile, error) {
    var p profile.Profile
    err := r.db.WithContext(ctx).Where("user_id = ?", userID).First(&p).Error
    if err != nil {
        if err == gorm.ErrRecordNotFound {
            return nil, profile.ErrProfileNotFound
        }
        return nil, err
    }
    return &p, nil
}

func (r *profileRepository) FindByUserIDs(ctx context.Context, userIDs []string) ([]*profile.Profile, error) {
    var profiles []*profile.Profile
    err := r.db.WithContext(ctx).Where("user_id IN ?", userIDs).Find(&profiles).Error
    return profiles, err
}

// READ - Lists
func (r *profileRepository) List(ctx context.Context, filter profile.ListFilter) ([]*profile.Profile, int64, error) {
    var profiles []*profile.Profile
    var total int64
    
    query := r.db.WithContext(ctx).Model(&profile.Profile{})
    query = r.applyFilters(query, filter)
    
    if err := query.Count(&total).Error; err != nil {
        return nil, 0, err
    }
    
    query = r.applySorting(query, filter)
    
    if filter.PageSize > 0 {
        offset := (filter.Page - 1) * filter.PageSize
        query = query.Offset(offset).Limit(filter.PageSize)
    }
    
    err := query.Find(&profiles).Error
    return profiles, total, err
}

func (r *profileRepository) Search(ctx context.Context, searchQuery string, filter profile.ListFilter) ([]*profile.Profile, int64, error) {
    var profiles []*profile.Profile
    var total int64
    
    query := r.db.WithContext(ctx).Model(&profile.Profile{})
    
    searchPattern := "%" + searchQuery + "%"
    query = query.Where(
        "title ILIKE ? OR bio ILIKE ? OR tagline ILIKE ? OR overview ILIKE ? OR industry ILIKE ?",
        searchPattern, searchPattern, searchPattern, searchPattern, searchPattern,
    )
    
    query = r.applyFilters(query, filter)
    
    if err := query.Count(&total).Error; err != nil {
        return nil, 0, err
    }
    
    query = r.applySorting(query, filter)
    
    if filter.PageSize > 0 {
        offset := (filter.Page - 1) * filter.PageSize
        query = query.Offset(offset).Limit(filter.PageSize)
    }
    
    err := query.Find(&profiles).Error
    return profiles, total, err
}

// READ - Business Queries
func (r *profileRepository) FindAvailableForWork(ctx context.Context, limit int) ([]*profile.Profile, error) {
    var profiles []*profile.Profile
    err := r.db.WithContext(ctx).
        Where("availability_status = ?", "available").
        Where("is_public = ?", true).
        Where("searchable_profile = ?", true).
        Order("quality_score DESC, completion_percentage DESC").
        Limit(limit).
        Find(&profiles).Error
    return profiles, err
}

func (r *profileRepository) FindByLocation(ctx context.Context, country, city string) ([]*profile.Profile, error) {
    var profiles []*profile.Profile
    query := r.db.WithContext(ctx).Where("country = ?", country)
    if city != "" {
        query = query.Where("city = ?", city)
    }
    err := query.Find(&profiles).Error
    return profiles, err
}

func (r *profileRepository) FindByRateRange(ctx context.Context, minRate, maxRate float64) ([]*profile.Profile, error) {
    var profiles []*profile.Profile
    err := r.db.WithContext(ctx).
        Where("hourly_rate >= ? AND hourly_rate <= ?", minRate, maxRate).
        Order("hourly_rate ASC").
        Find(&profiles).Error
    return profiles, err
}

func (r *profileRepository) FindFeatured(ctx context.Context, limit int) ([]*profile.Profile, error) {
    var profiles []*profile.Profile
    // Join with users table to check is_featured flag
    err := r.db.WithContext(ctx).
        Joins("JOIN users ON users.id = profiles.user_id").
        Where("users.is_featured = ?", true).
        Where("profiles.is_public = ?", true).
        Limit(limit).
        Find(&profiles).Error
    return profiles, err
}

func (r *profileRepository) FindRecentlyUpdated(ctx context.Context, limit int) ([]*profile.Profile, error) {
    var profiles []*profile.Profile
    err := r.db.WithContext(ctx).
        Order("updated_at DESC").
        Limit(limit).
        Find(&profiles).Error
    return profiles, err
}

func (r *profileRepository) FindIncomplete(ctx context.Context) ([]*profile.Profile, error) {
    var profiles []*profile.Profile
    err := r.db.WithContext(ctx).
        Where("completion_percentage < ?", 80).
        Order("completion_percentage ASC").
        Find(&profiles).Error
    return profiles, err
}

func (r *profileRepository) FindWithLowQuality(ctx context.Context, threshold float64) ([]*profile.Profile, error) {
    var profiles []*profile.Profile
    err := r.db.WithContext(ctx).
        Where("quality_score < ?", threshold).
        Order("quality_score ASC").
        Find(&profiles).Error
    return profiles, err
}

// UPDATE
func (r *profileRepository) Update(ctx context.Context, p *profile.Profile) error {
    return r.db.WithContext(ctx).Save(p).Error
}

func (r *profileRepository) UpdateCompletionPercentage(ctx context.Context, userID string, percentage int) error {
    return r.db.WithContext(ctx).
        Model(&profile.Profile{}).
        Where("user_id = ?", userID).
        Update("completion_percentage", percentage).Error
}

func (r *profileRepository) UpdateQualityScore(ctx context.Context, userID string, score float64) error {
    now := time.Now()
    return r.db.WithContext(ctx).
        Model(&profile.Profile{}).
        Where("user_id = ?", userID).
        Updates(map[string]interface{}{
            "quality_score":      score,
            "last_quality_check": now,
        }).Error
}

func (r *profileRepository) IncrementViews(ctx context.Context, userID string) error {
    return r.db.WithContext(ctx).
        Model(&profile.Profile{}).
        Where("user_id = ?", userID).
        Updates(map[string]interface{}{
            "profile_views":            gorm.Expr("profile_views + ?", 1),
            "profile_views_this_week":  gorm.Expr("profile_views_this_week + ?", 1),
            "profile_views_this_month": gorm.Expr("profile_views_this_month + ?", 1),
            "last_viewed_at":           time.Now(),
        }).Error
}

func (r *profileRepository) UpdateAvailabilityStatus(ctx context.Context, userID, status string) error {
    return r.db.WithContext(ctx).
        Model(&profile.Profile{}).
        Where("user_id = ?", userID).
        Update("availability_status", status).Error
}

// DELETE
func (r *profileRepository) Delete(ctx context.Context, id string) error {
    return r.db.WithContext(ctx).Delete(&profile.Profile{}, "id = ?", id).Error
}

func (r *profileRepository) SoftDelete(ctx context.Context, id string) error {
    return r.db.WithContext(ctx).
        Model(&profile.Profile{}).
        Where("id = ?", id).
        Update("deleted_at", time.Now()).Error
}

// ANALYTICS
func (r *profileRepository) CountByCountry(ctx context.Context, country string) (int64, error) {
    var count int64
    err := r.db.WithContext(ctx).
        Model(&profile.Profile{}).
        Where("country = ?", country).
        Count(&count).Error
    return count, err
}

func (r *profileRepository) GetAverageQualityScore(ctx context.Context) (float64, error) {
    var avgScore float64
    err := r.db.WithContext(ctx).
        Model(&profile.Profile{}).
        Select("AVG(quality_score)").
        Scan(&avgScore).Error
    return avgScore, err
}

func (r *profileRepository) GetAverageCompletionRate(ctx context.Context) (float64, error) {
    var avgCompletion float64
    err := r.db.WithContext(ctx).
        Model(&profile.Profile{}).
        Select("AVG(completion_percentage)").
        Scan(&avgCompletion).Error
    return avgCompletion, err
}

// HELPERS
func (r *profileRepository) applyFilters(query *gorm.DB, filter profile.ListFilter) *gorm.DB {
    if filter.Country != nil {
        query = query.Where("country = ?", *filter.Country)
    }
    if filter.City != nil {
        query = query.Where("city = ?", *filter.City)
    }
    if filter.MinRate != nil {
        query = query.Where("hourly_rate >= ?", *filter.MinRate)
    }
    if filter.MaxRate != nil {
        query = query.Where("hourly_rate <= ?", *filter.MaxRate)
    }
    if filter.AvailabilityStatus != nil {
        query = query.Where("availability_status = ?", *filter.AvailabilityStatus)
    }
    if filter.MinYearsExperience != nil {
        query = query.Where("years_of_experience >= ?", *filter.MinYearsExperience)
    }
    if filter.Industry != nil {
        query = query.Where("industry = ?", *filter.Industry)
    }
    if filter.IsPublic != nil {
        query = query.Where("is_public = ?", *filter.IsPublic)
    }
    if filter.MinQualityScore != nil {
        query = query.Where("quality_score >= ?", *filter.MinQualityScore)
    }
    if filter.MinCompletion != nil {
        query = query.Where("completion_percentage >= ?", *filter.MinCompletion)
    }
    return query
}

func (r *profileRepository) applySorting(query *gorm.DB, filter profile.ListFilter) *gorm.DB {
    if filter.SortBy == "" {
        return query.Order("updated_at DESC")
    }
    
    order := "DESC"
    if filter.SortOrder == "asc" {
        order = "ASC"
    }
    
    return query.Order(fmt.Sprintf("%s %s", filter.SortBy, order))
}