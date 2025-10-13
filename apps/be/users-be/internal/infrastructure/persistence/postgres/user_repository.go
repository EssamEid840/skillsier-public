// internal/infrastructure/persistence/postgres/user_repository.go
package postgres

import (
    "context"
    "fmt"
    "time"
    
    "gorm.io/gorm"
    "users-be/internal/domain/user"
)

type userRepository struct {
    db *gorm.DB
}

func NewUserRepository(db *gorm.DB) user.Repository {
    return &userRepository{db: db}
}

// ============================================================================
// CREATE OPERATIONS
// ============================================================================

func (r *userRepository) Create(ctx context.Context, u *user.User) error {
    return r.db.WithContext(ctx).Create(u).Error
}

func (r *userRepository) CreateBatch(ctx context.Context, users []*user.User) error {
    return r.db.WithContext(ctx).CreateInBatches(users, 100).Error
}

// ============================================================================
// READ OPERATIONS - Single Record
// ============================================================================

func (r *userRepository) FindByID(ctx context.Context, id string) (*user.User, error) {
    var u user.User
    err := r.db.WithContext(ctx).Where("id = ?", id).First(&u).Error
    if err != nil {
        if err == gorm.ErrRecordNotFound {
            return nil, user.ErrUserNotFound
        }
        return nil, err
    }
    return &u, nil
}

func (r *userRepository) FindByKeycloakID(ctx context.Context, keycloakID string) (*user.User, error) {
    var u user.User
    err := r.db.WithContext(ctx).Where("keycloak_id = ?", keycloakID).First(&u).Error
    if err != nil {
        if err == gorm.ErrRecordNotFound {
            return nil, user.ErrUserNotFound
        }
        return nil, err
    }
    return &u, nil
}

func (r *userRepository) FindByEmail(ctx context.Context, email string) (*user.User, error) {
    var u user.User
    err := r.db.WithContext(ctx).Where("email = ?", email).First(&u).Error
    if err != nil {
        if err == gorm.ErrRecordNotFound {
            return nil, user.ErrUserNotFound
        }
        return nil, err
    }
    return &u, nil
}

func (r *userRepository) FindByUsername(ctx context.Context, username string) (*user.User, error) {
    var u user.User
    err := r.db.WithContext(ctx).Where("username = ?", username).First(&u).Error
    if err != nil {
        if err == gorm.ErrRecordNotFound {
            return nil, user.ErrUserNotFound
        }
        return nil, err
    }
    return &u, nil
}

func (r *userRepository) FindByReferralCode(ctx context.Context, code string) (*user.User, error) {
    var u user.User
    err := r.db.WithContext(ctx).Where("referral_code = ?", code).First(&u).Error
    if err != nil {
        if err == gorm.ErrRecordNotFound {
            return nil, user.ErrUserNotFound
        }
        return nil, err
    }
    return &u, nil
}

func (r *userRepository) FindByIDs(ctx context.Context, ids []string) ([]*user.User, error) {
    var users []*user.User
    err := r.db.WithContext(ctx).Where("id IN ?", ids).Find(&users).Error
    return users, err
}

// ============================================================================
// READ OPERATIONS - Lists with Filtering
// ============================================================================

func (r *userRepository) List(ctx context.Context, filter user.ListFilter) ([]*user.User, int64, error) {
    var users []*user.User
    var total int64
    
    query := r.db.WithContext(ctx).Model(&user.User{})
    
    // Apply filters
    query = r.applyFilters(query, filter)
    
    // Count total before pagination
    if err := query.Count(&total).Error; err != nil {
        return nil, 0, err
    }
    
    // Apply sorting
    query = r.applySorting(query, filter)
    
    // Apply pagination
    if filter.PageSize > 0 {
        offset := (filter.Page - 1) * filter.PageSize
        query = query.Offset(offset).Limit(filter.PageSize)
    }
    
    err := query.Find(&users).Error
    return users, total, err
}

func (r *userRepository) Search(ctx context.Context, searchQuery string, filter user.ListFilter) ([]*user.User, int64, error) {
    var users []*user.User
    var total int64
    
    query := r.db.WithContext(ctx).Model(&user.User{})
    
    // Full-text search across multiple fields
    searchPattern := "%" + searchQuery + "%"
    query = query.Where(
        "username ILIKE ? OR email ILIKE ? OR first_name ILIKE ? OR last_name ILIKE ? OR display_name ILIKE ?",
        searchPattern, searchPattern, searchPattern, searchPattern, searchPattern,
    )
    
    // Apply additional filters
    query = r.applyFilters(query, filter)
    
    // Count total
    if err := query.Count(&total).Error; err != nil {
        return nil, 0, err
    }
    
    // Apply sorting
    query = r.applySorting(query, filter)
    
    // Apply pagination
    if filter.PageSize > 0 {
        offset := (filter.Page - 1) * filter.PageSize
        query = query.Offset(offset).Limit(filter.PageSize)
    }
    
    err := query.Find(&users).Error
    return users, total, err
}

// Helper to apply filters
func (r *userRepository) applyFilters(query *gorm.DB, filter user.ListFilter) *gorm.DB {
    if filter.UserType != nil {
        query = query.Where("user_type = ?", *filter.UserType)
    }
    
    if filter.Status != nil {
        query = query.Where("status = ?", *filter.Status)
    }
    
    if filter.Country != nil {
        query = query.Where("country = ?", *filter.Country)
    }
    
    if filter.IsVerified != nil && *filter.IsVerified {
        query = query.Where("email_verified = ? AND identity_verified = ? AND payment_verified = ?", 
            true, true, true)
    }
    
    if filter.IsTopRated != nil {
        query = query.Where("is_top_rated = ?", *filter.IsTopRated)
    }
    
    if filter.IsFeatured != nil {
        query = query.Where("is_featured = ?", *filter.IsFeatured)
    }
    
    if filter.MinReputation != nil {
        query = query.Where("reputation_score >= ?", *filter.MinReputation)
    }
    
    if filter.CreatedAfter != nil {
        query = query.Where("created_at >= ?", *filter.CreatedAfter)
    }
    
    if filter.CreatedBefore != nil {
        query = query.Where("created_at <= ?", *filter.CreatedBefore)
    }
    
    if filter.LastSeenAfter != nil {
        query = query.Where("last_seen_at >= ?", *filter.LastSeenAfter)
    }
    
    if !filter.IncludeDeleted {
        query = query.Where("deleted_at IS NULL")
    }
    
    return query
}

// Helper to apply sorting
func (r *userRepository) applySorting(query *gorm.DB, filter user.ListFilter) *gorm.DB {
    if filter.SortBy == "" {
        return query.Order("created_at DESC")
    }
    
    order := "DESC"
    if filter.SortOrder == "asc" {
        order = "ASC"
    }
    
    return query.Order(fmt.Sprintf("%s %s", filter.SortBy, order))
}

// ============================================================================
// UPDATE OPERATIONS
// ============================================================================

func (r *userRepository) Update(ctx context.Context, u *user.User) error {
    return r.db.WithContext(ctx).Save(u).Error
}

func (r *userRepository) UpdateStatus(ctx context.Context, id string, status user.AccountStatus) error {
    return r.db.WithContext(ctx).
        Model(&user.User{}).
        Where("id = ?", id).
        Update("status", status).Error
}

func (r *userRepository) UpdateLastSeen(ctx context.Context, id string) error {
    now := time.Now()
    return r.db.WithContext(ctx).
        Model(&user.User{}).
        Where("id = ?", id).
        Updates(map[string]interface{}{
            "last_seen_at": now,
            "is_online":    true,
        }).Error
}

func (r *userRepository) UpdateOnlineStatus(ctx context.Context, id string, isOnline bool) error {
    updates := map[string]interface{}{
        "is_online": isOnline,
    }
    if !isOnline {
        updates["last_seen_at"] = time.Now()
    }
    return r.db.WithContext(ctx).
        Model(&user.User{}).
        Where("id = ?", id).
        Updates(updates).Error
}

func (r *userRepository) IncrementLoginCount(ctx context.Context, id string) error {
    return r.db.WithContext(ctx).
        Model(&user.User{}).
        Where("id = ?", id).
        Update("login_count", gorm.Expr("login_count + ?", 1)).Error
}

func (r *userRepository) IncrementFailedLoginAttempts(ctx context.Context, id string) error {
    now := time.Now()
    return r.db.WithContext(ctx).
        Model(&user.User{}).
        Where("id = ?", id).
        Updates(map[string]interface{}{
            "failed_login_attempts": gorm.Expr("failed_login_attempts + ?", 1),
            "last_failed_login_at":  now,
        }).Error
}

func (r *userRepository) ResetFailedLoginAttempts(ctx context.Context, id string) error {
    return r.db.WithContext(ctx).
        Model(&user.User{}).
        Where("id = ?", id).
        Updates(map[string]interface{}{
            "failed_login_attempts": 0,
            "last_failed_login_at":  nil,
        }).Error
}

// ============================================================================
// DELETE OPERATIONS
// ============================================================================

func (r *userRepository) Delete(ctx context.Context, id string) error {
    return r.db.WithContext(ctx).Delete(&user.User{}, "id = ?", id).Error
}

func (r *userRepository) SoftDelete(ctx context.Context, id string, deletedBy string) error {
    return r.db.WithContext(ctx).
        Model(&user.User{}).
        Where("id = ?", id).
        Updates(map[string]interface{}{
            "deleted_at": time.Now(),
            "deleted_by": deletedBy,
            "status":     user.AccountStatusDeleted,
        }).Error
}

func (r *userRepository) HardDelete(ctx context.Context, id string) error {
    return r.db.WithContext(ctx).Unscoped().Delete(&user.User{}, "id = ?", id).Error
}

func (r *userRepository) RestoreDeleted(ctx context.Context, id string) error {
    return r.db.WithContext(ctx).
        Model(&user.User{}).
        Unscoped().
        Where("id = ?", id).
        Updates(map[string]interface{}{
            "deleted_at": nil,
            "deleted_by": nil,
            "status":     user.AccountStatusActive,
        }).Error
}

// ============================================================================
// BUSINESS QUERIES
// ============================================================================

func (r *userRepository) FindTopRatedFreelancers(ctx context.Context, limit int) ([]*user.User, error) {
    var users []*user.User
    err := r.db.WithContext(ctx).
        Where("user_type IN ?", []string{string(user.UserTypeFreelancer), string(user.UserTypeBoth)}).
        Where("status = ?", user.AccountStatusActive).
        Where("is_top_rated = ?", true).
        Order("reputation_score DESC, total_ratings DESC").
        Limit(limit).
        Find(&users).Error
    return users, err
}

func (r *userRepository) FindFeaturedUsers(ctx context.Context, userType user.UserType, limit int) ([]*user.User, error) {
    var users []*user.User
    query := r.db.WithContext(ctx).
        Where("is_featured = ?", true).
        Where("status = ?", user.AccountStatusActive)
    
    if userType != "" {
        query = query.Where("user_type = ?", userType)
    }
    
    err := query.Order("reputation_score DESC").Limit(limit).Find(&users).Error
    return users, err
}

func (r *userRepository) FindOnlineUsers(ctx context.Context, userType user.UserType) ([]*user.User, error) {
    var users []*user.User
    query := r.db.WithContext(ctx).
        Where("is_online = ?", true).
        Where("status = ?", user.AccountStatusActive)
    
    if userType != "" {
        query = query.Where("user_type = ?", userType)
    }
    
    err := query.Order("last_seen_at DESC").Find(&users).Error
    return users, err
}

func (r *userRepository) FindInactiveUsers(ctx context.Context, days int) ([]*user.User, error) {
    var users []*user.User
    cutoffDate := time.Now().AddDate(0, 0, -days)
    err := r.db.WithContext(ctx).
        Where("last_seen_at < ?", cutoffDate).
        Where("status = ?", user.AccountStatusActive).
        Find(&users).Error
    return users, err
}

func (r *userRepository) FindUsersWithWarnings(ctx context.Context) ([]*user.User, error) {
    var users []*user.User
    err := r.db.WithContext(ctx).
        Where("has_active_warnings = ?", true).
        Where("warning_count > ?", 0).
        Order("warning_count DESC").
        Find(&users).Error
    return users, err
}

func (r *userRepository) FindUsersByCountry(ctx context.Context, country string, filter user.ListFilter) ([]*user.User, int64, error) {
    var users []*user.User
    var total int64
    
    query := r.db.WithContext(ctx).Model(&user.User{}).Where("country = ?", country)
    
    query = r.applyFilters(query, filter)
    
    if err := query.Count(&total).Error; err != nil {
        return nil, 0, err
    }
    
    query = r.applySorting(query, filter)
    
    if filter.PageSize > 0 {
        offset := (filter.Page - 1) * filter.PageSize
        query = query.Offset(offset).Limit(filter.PageSize)
    }
    
    err := query.Find(&users).Error
    return users, total, err
}

// ============================================================================
// ANALYTICS QUERIES
// ============================================================================

func (r *userRepository) CountByUserType(ctx context.Context, userType user.UserType) (int64, error) {
    var count int64
    err := r.db.WithContext(ctx).
        Model(&user.User{}).
        Where("user_type = ?", userType).
        Where("deleted_at IS NULL").
        Count(&count).Error
    return count, err
}

func (r *userRepository) CountByStatus(ctx context.Context, status user.AccountStatus) (int64, error) {
    var count int64
    err := r.db.WithContext(ctx).
        Model(&user.User{}).
        Where("status = ?", status).
        Where("deleted_at IS NULL").
        Count(&count).Error
    return count, err
}

func (r *userRepository) CountVerifiedUsers(ctx context.Context) (int64, error) {
    var count int64
    err := r.db.WithContext(ctx).
        Model(&user.User{}).
        Where("email_verified = ? AND identity_verified = ? AND payment_verified = ?", 
            true, true, true).
        Where("deleted_at IS NULL").
        Count(&count).Error
    return count, err
}

func (r *userRepository) GetAverageReputationScore(ctx context.Context, userType user.UserType) (float64, error) {
    var avgScore float64
    query := r.db.WithContext(ctx).
        Model(&user.User{}).
        Where("deleted_at IS NULL").
        Where("total_ratings > ?", 0)
    
    if userType != "" {
        query = query.Where("user_type = ?", userType)
    }
    
    err := query.Select("AVG(reputation_score)").Scan(&avgScore).Error
    return avgScore, err
}

// ============================================================================
// EXISTENCE CHECKS
// ============================================================================

func (r *userRepository) ExistsWithEmail(ctx context.Context, email string) (bool, error) {
    var count int64
    err := r.db.WithContext(ctx).
        Model(&user.User{}).
        Where("email = ?", email).
        Where("deleted_at IS NULL").
        Count(&count).Error
    return count > 0, err
}

func (r *userRepository) ExistsWithUsername(ctx context.Context, username string) (bool, error) {
    var count int64
    err := r.db.WithContext(ctx).
        Model(&user.User{}).
        Where("username = ?", username).
        Where("deleted_at IS NULL").
        Count(&count).Error
    return count > 0, err
}

func (r *userRepository) ExistsWithKeycloakID(ctx context.Context, keycloakID string) (bool, error) {
    var count int64
    err := r.db.WithContext(ctx).
        Model(&user.User{}).
        Where("keycloak_id = ?", keycloakID).
        Where("deleted_at IS NULL").
        Count(&count).Error
    return count > 0, err
}