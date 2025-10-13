
package postgres

import (
	"context"
	"fmt"
	"time"
	
	"gorm.io/gorm"
	
	"users-be/internal/domain/user"
)

// UserRepository implements user.Repository interface using GORM
type UserRepository struct {
	db *gorm.DB
}

// NewUserRepository creates a new user repository
func NewUserRepository(db *gorm.DB) user.Repository {
	return &UserRepository{db: db}
}

// ============================================================================
// CREATE OPERATIONS
// ============================================================================

func (r *UserRepository) Create(ctx context.Context, u *user.User) error {
	return r.db.WithContext(ctx).Create(u).Error
}

func (r *UserRepository) CreateBatch(ctx context.Context, users []*user.User) error {
	return r.db.WithContext(ctx).CreateInBatches(users, 100).Error
}

// ============================================================================
// READ OPERATIONS - SINGLE ENTITY
// ============================================================================

func (r *UserRepository) FindByID(ctx context.Context, id string) (*user.User, error) {
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

func (r *UserRepository) FindByKeycloakID(ctx context.Context, keycloakID string) (*user.User, error) {
	var u user.User
	err := r.db.WithContext(ctx).Where("keycloak_id = ?", keycloakID).First(&u).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, user.ErrUserNotFoundByKeycloakID
		}
		return nil, err
	}
	return &u, nil
}

func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*user.User, error) {
	var u user.User
	err := r.db.WithContext(ctx).Where("email_value = ?", email).First(&u).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, user.ErrUserNotFoundByEmail
		}
		return nil, err
	}
	return &u, nil
}

func (r *UserRepository) FindByUsername(ctx context.Context, username string) (*user.User, error) {
	var u user.User
	err := r.db.WithContext(ctx).Where("username = ?", username).First(&u).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, user.ErrUserNotFoundByUsername
		}
		return nil, err
	}
	return &u, nil
}

func (r *UserRepository) FindByReferralCode(ctx context.Context, code string) (*user.User, error) {
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

// ============================================================================
// READ OPERATIONS - MULTIPLE ENTITIES
// ============================================================================

func (r *UserRepository) FindByIDs(ctx context.Context, ids []string) ([]*user.User, error) {
	var users []*user.User
	err := r.db.WithContext(ctx).Where("id IN ?", ids).Find(&users).Error
	return users, err
}

func (r *UserRepository) FindByKeycloakIDs(ctx context.Context, keycloakIDs []string) ([]*user.User, error) {
	var users []*user.User
	err := r.db.WithContext(ctx).Where("keycloak_id IN ?", keycloakIDs).Find(&users).Error
	return users, err
}

func (r *UserRepository) FindByEmails(ctx context.Context, emails []string) ([]*user.User, error) {
	var users []*user.User
	err := r.db.WithContext(ctx).Where("email_value IN ?", emails).Find(&users).Error
	return users, err
}

// ============================================================================
// LIST & SEARCH OPERATIONS
// ============================================================================

func (r *UserRepository) List(ctx context.Context, filter user.ListFilter) ([]*user.User, int64, error) {
	var users []*user.User
	var total int64
	
	query := r.db.WithContext(ctx).Model(&user.User{})
	
	// Apply filters
	query = r.applyFilters(query, filter)
	
	// Count total
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	
	// Apply sorting and pagination
	query = r.applySorting(query, filter)
	query = query.Limit(filter.GetLimit()).Offset(filter.GetOffset())
	
	// Execute query
	if err := query.Find(&users).Error; err != nil {
		return nil, 0, err
	}
	
	return users, total, nil
}

func (r *UserRepository) Search(ctx context.Context, query string, filter user.ListFilter) ([]*user.User, int64, error) {
	var users []*user.User
	var total int64
	
	db := r.db.WithContext(ctx).Model(&user.User{})
	
	// Full-text search on multiple fields
	searchQuery := "%" + query + "%"
	db = db.Where(
		"username ILIKE ? OR email_value ILIKE ? OR full_name ILIKE ? OR bio ILIKE ? OR tagline ILIKE ?",
		searchQuery, searchQuery, searchQuery, searchQuery, searchQuery,
	)
	
	// Apply additional filters
	db = r.applyFilters(db, filter)
	
	// Count total
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	
	// Apply sorting and pagination
	db = r.applySorting(db, filter)
	db = db.Limit(filter.GetLimit()).Offset(filter.GetOffset())
	
	// Execute query
	if err := db.Find(&users).Error; err != nil {
		return nil, 0, err
	}
	
	return users, total, nil
}

func (r *UserRepository) FindAll(ctx context.Context) ([]*user.User, error) {
	var users []*user.User
	err := r.db.WithContext(ctx).Find(&users).Error
	return users, err
}

// ============================================================================
// UPDATE OPERATIONS - FULL ENTITY
// ============================================================================

func (r *UserRepository) Update(ctx context.Context, u *user.User) error {
	u.UpdatedAt = time.Now()
	return r.db.WithContext(ctx).Save(u).Error
}

func (r *UserRepository) UpdateBatch(ctx context.Context, users []*user.User) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		for _, u := range users {
			u.UpdatedAt = time.Now()
			if err := tx.Save(u).Error; err != nil {
				return err
			}
		}
		return nil
	})
}

// ============================================================================
// UPDATE OPERATIONS - SPECIFIC FIELDS
// ============================================================================

func (r *UserRepository) UpdateStatus(ctx context.Context, id string, status user.AccountStatus) error {
	return r.db.WithContext(ctx).Model(&user.User{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"status":     status,
			"updated_at": time.Now(),
		}).Error
}

func (r *UserRepository) UpdateVerificationStatus(ctx context.Context, id string, status user.VerificationStatus) error {
	return r.db.WithContext(ctx).Model(&user.User{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"verification_status": status,
			"updated_at":          time.Now(),
		}).Error
}

func (r *UserRepository) UpdateLastSeen(ctx context.Context, id string) error {
	now := time.Now()
	return r.db.WithContext(ctx).Model(&user.User{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"last_seen_at":   &now,
			"last_active_at": &now,
			"is_online":      true,
		}).Error
}

func (r *UserRepository) UpdateOnlineStatus(ctx context.Context, id string, isOnline bool) error {
	return r.db.WithContext(ctx).Model(&user.User{}).
		Where("id = ?", id).
		Update("is_online", isOnline).Error
}

func (r *UserRepository) UpdateProfileCompleteness(ctx context.Context, id string, percentage int) error {
	return r.db.WithContext(ctx).Model(&user.User{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"profile_completeness": percentage,
			"profile_completed":    percentage >= 80,
			"updated_at":           time.Now(),
		}).Error
}

func (r *UserRepository) UpdateRating(ctx context.Context, id string, rating float64, totalReviews int) error {
	return r.db.WithContext(ctx).Model(&user.User{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"rating":        rating,
			"total_reviews": totalReviews,
			"updated_at":    time.Now(),
		}).Error
}

func (r *UserRepository) UpdateStats(ctx context.Context, id string, completedJobs, totalJobs int, successRate float64) error {
	return r.db.WithContext(ctx).Model(&user.User{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"completed_jobs": completedJobs,
			"total_jobs":     totalJobs,
			"success_rate":   successRate,
			"updated_at":     time.Now(),
		}).Error
}

func (r *UserRepository) UpdateEarnings(ctx context.Context, id string, amount float64) error {
	return r.db.WithContext(ctx).Model(&user.User{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"total_earnings": amount,
			"updated_at":     time.Now(),
		}).Error
}

func (r *UserRepository) UpdateSpending(ctx context.Context, id string, amount float64) error {
	return r.db.WithContext(ctx).Model(&user.User{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"total_spent": amount,
			"updated_at":  time.Now(),
		}).Error
}

func (r *UserRepository) UpdateBalance(ctx context.Context, id string, amount float64) error {
	return r.db.WithContext(ctx).Model(&user.User{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"current_balance": amount,
			"updated_at":      time.Now(),
		}).Error
}

// ============================================================================
// UPDATE OPERATIONS - SECURITY & ACTIVITY
// ============================================================================

func (r *UserRepository) RecordLogin(ctx context.Context, id, ipAddress, userAgent string) error {
	now := time.Now()
	return r.db.WithContext(ctx).Model(&user.User{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"last_login_at":   &now,
			"last_login_ip":   ipAddress,
			"last_user_agent": userAgent,
			"login_count":     gorm.Expr("login_count + 1"),
			"login_attempts":  0,
			"is_online":       true,
			"updated_at":      now,
		}).Error
}

func (r *UserRepository) IncrementLoginCount(ctx context.Context, id string) error {
	return r.db.WithContext(ctx).Model(&user.User{}).
		Where("id = ?", id).
		Update("login_count", gorm.Expr("login_count + 1")).Error
}

func (r *UserRepository) IncrementFailedLoginAttempts(ctx context.Context, id string) error {
	return r.db.WithContext(ctx).Model(&user.User{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"login_attempts": gorm.Expr("login_attempts + 1"),
			"updated_at":     time.Now(),
		}).Error
}

func (r *UserRepository) ResetFailedLoginAttempts(ctx context.Context, id string) error {
	return r.db.WithContext(ctx).Model(&user.User{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"login_attempts": 0,
			"locked_until":   nil,
			"updated_at":     time.Now(),
		}).Error
}

func (r *UserRepository) LockAccount(ctx context.Context, id string, until time.Time) error {
	return r.db.WithContext(ctx).Model(&user.User{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"locked_until": &until,
			"updated_at":   time.Now(),
		}).Error
}

func (r *UserRepository) UnlockAccount(ctx context.Context, id string) error {
	return r.db.WithContext(ctx).Model(&user.User{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"locked_until":   nil,
			"login_attempts": 0,
			"updated_at":     time.Now(),
		}).Error
}

// ============================================================================
// UPDATE OPERATIONS - VERIFICATION
// ============================================================================

func (r *UserRepository) VerifyEmail(ctx context.Context, id string) error {
	now := time.Now().Unix()
	return r.db.WithContext(ctx).Model(&user.User{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"email_verified":    true,
			"email_verified_at": &now,
			"status":            user.AccountStatusActive,
			"updated_at":        time.Now(),
		}).Error
}

func (r *UserRepository) VerifyPhone(ctx context.Context, id string) error {
	now := time.Now().Unix()
	return r.db.WithContext(ctx).Model(&user.User{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"phone_verified":    true,
			"phone_verified_at": &now,
			"updated_at":        time.Now(),
		}).Error
}

func (r *UserRepository) VerifyIdentity(ctx context.Context, id string) error {
	return r.db.WithContext(ctx).Model(&user.User{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"identity_verified":    true,
			"verification_status":  user.VerificationStatusVerified,
			"updated_at":           time.Now(),
		}).Error
}

// ============================================================================
// UPDATE OPERATIONS - BADGES & ACHIEVEMENTS
// ============================================================================

func (r *UserRepository) AssignBadge(ctx context.Context, id string, badge user.BadgeType) error {
	var u user.User
	if err := r.db.WithContext(ctx).Where("id = ?", id).First(&u).Error; err != nil {
		return err
	}
	
	// Check if badge already exists
	for _, b := range u.Badges {
		if b == badge {
			return user.ErrBadgeAlreadyAssigned
		}
	}
	
	u.Badges = append(u.Badges, badge)
	u.UpdatedAt = time.Now()
	
	return r.db.WithContext(ctx).Save(&u).Error
}

func (r *UserRepository) RemoveBadge(ctx context.Context, id string, badge user.BadgeType) error {
	var u user.User
	if err := r.db.WithContext(ctx).Where("id = ?", id).First(&u).Error; err != nil {
		return err
	}
	
	newBadges := make([]user.BadgeType, 0)
	for _, b := range u.Badges {
		if b != badge {
			newBadges = append(newBadges, b)
		}
	}
	
	u.Badges = newBadges
	u.UpdatedAt = time.Now()
	
	return r.db.WithContext(ctx).Save(&u).Error
}

func (r *UserRepository) SetFeatured(ctx context.Context, id string, featured bool) error {
	return r.db.WithContext(ctx).Model(&user.User{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"is_featured": featured,
			"updated_at":  time.Now(),
		}).Error
}

func (r *UserRepository) SetTopRated(ctx context.Context, id string, topRated bool) error {
	return r.db.WithContext(ctx).Model(&user.User{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"is_top_rated": topRated,
			"updated_at":   time.Now(),
		}).Error
}

// ============================================================================
// UPDATE OPERATIONS - MODERATION
// ============================================================================

func (r *UserRepository) IncrementWarningCount(ctx context.Context, id string) error {
	return r.db.WithContext(ctx).Model(&user.User{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"warning_count": gorm.Expr("warning_count + 1"),
			"updated_at":    time.Now(),
		}).Error
}

func (r *UserRepository) IncrementSuspensionCount(ctx context.Context, id string) error {
	return r.db.WithContext(ctx).Model(&user.User{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"suspension_count": gorm.Expr("suspension_count + 1"),
			"updated_at":       time.Now(),
		}).Error
}

func (r *UserRepository) IncrementBanCount(ctx context.Context, id string) error {
	return r.db.WithContext(ctx).Model(&user.User{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"ban_count":  gorm.Expr("ban_count + 1"),
			"updated_at": time.Now(),
		}).Error
}

func (r *UserRepository) IncrementFlagCount(ctx context.Context, id string) error {
	return r.db.WithContext(ctx).Model(&user.User{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"flag_count": gorm.Expr("flag_count + 1"),
			"updated_at": time.Now(),
		}).Error
}

func (r *UserRepository) AddNote(ctx context.Context, id, note string) error {
	var u user.User
	if err := r.db.WithContext(ctx).Where("id = ?", id).First(&u).Error; err != nil {
		return err
	}
	
	timestamp := time.Now().Format(time.RFC3339)
	if u.Notes == "" {
		u.Notes = fmt.Sprintf("[%s] %s", timestamp, note)
	} else {
		u.Notes = fmt.Sprintf("%s\n[%s] %s", u.Notes, timestamp, note)
	}
	u.UpdatedAt = time.Now()
	
	return r.db.WithContext(ctx).Save(&u).Error
}

func (r *UserRepository) AddTag(ctx context.Context, id, tag string) error {
	var u user.User
	if err := r.db.WithContext(ctx).Where("id = ?", id).First(&u).Error; err != nil {
		return err
	}
	
	// Check if tag already exists
	for _, t := range u.Tags {
		if t == tag {
			return nil // Already exists
		}
	}
	
	u.Tags = append(u.Tags, tag)
	u.UpdatedAt = time.Now()
	
	return r.db.WithContext(ctx).Save(&u).Error
}

func (r *UserRepository) RemoveTag(ctx context.Context, id, tag string) error {
	var u user.User
	if err := r.db.WithContext(ctx).Where("id = ?", id).First(&u).Error; err != nil {
		return err
	}
	
	newTags := make([]string, 0)
	for _, t := range u.Tags {
		if t != tag {
			newTags = append(newTags, t)
		}
	}
	
	u.Tags = newTags
	u.UpdatedAt = time.Now()
	
	return r.db.WithContext(ctx).Save(&u).Error
}

// ============================================================================
// DELETE OPERATIONS
// ============================================================================

func (r *UserRepository) Delete(ctx context.Context, id string) error {
	now := time.Now()
	return r.db.WithContext(ctx).Model(&user.User{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"deleted_at": &now,
			"status":     user.AccountStatusDeleted,
			"updated_at": now,
		}).Error
}

func (r *UserRepository) SoftDelete(ctx context.Context, id, deletedBy string) error {
	now := time.Now()
	return r.db.WithContext(ctx).Model(&user.User{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"deleted_at": &now,
			"deleted_by": deletedBy,
			"status":     user.AccountStatusDeleted,
			"updated_at": now,
		}).Error
}

func (r *UserRepository) HardDelete(ctx context.Context, id string) error {
	return r.db.WithContext(ctx).Unscoped().Where("id = ?", id).Delete(&user.User{}).Error
}

func (r *UserRepository) RestoreDeleted(ctx context.Context, id string) error {
	return r.db.WithContext(ctx).Model(&user.User{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"deleted_at": nil,
			"deleted_by": nil,
			"status":     user.AccountStatusActive,
			"updated_at": time.Now(),
		}).Error
}

// ============================================================================
// EXISTENCE CHECKS
// ============================================================================

func (r *UserRepository) ExistsByID(ctx context.Context, id string) (bool, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&user.User{}).Where("id = ?", id).Count(&count).Error
	return count > 0, err
}

func (r *UserRepository) ExistsByKeycloakID(ctx context.Context, keycloakID string) (bool, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&user.User{}).Where("keycloak_id = ?", keycloakID).Count(&count).Error
	return count > 0, err
}

func (r *UserRepository) ExistsByEmail(ctx context.Context, email string) (bool, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&user.User{}).Where("email_value = ?", email).Count(&count).Error
	return count > 0, err
}

func (r *UserRepository) ExistsByUsername(ctx context.Context, username string) (bool, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&user.User{}).Where("username = ?", username).Count(&count).Error
	return count > 0, err
}

func (r *UserRepository) FindUnverifiedUsers(ctx context.Context, filter user.ListFilter) ([]*user.User, int64, error) {
	emailVerified := false
	filter.EmailVerified = &emailVerified
	return r.List(ctx, filter)
}

func (r *UserRepository) FindPendingVerification(ctx context.Context, filter user.ListFilter) ([]*user.User, int64, error) {
	status := user.VerificationStatusPending
	filter.VerificationStatus = &status
	return r.List(ctx, filter)
}

func (r *UserRepository) FindVerifiedUsers(ctx context.Context, filter user.ListFilter) ([]*user.User, int64, error) {
	identityVerified := true
	filter.IdentityVerified = &identityVerified
	return r.List(ctx, filter)
}

// ============================================================================
// BUSINESS QUERIES - MODERATION
// ============================================================================

func (r *UserRepository) FindSuspendedUsers(ctx context.Context, filter user.ListFilter) ([]*user.User, int64, error) {
	status := user.AccountStatusSuspended
	filter.Status = &status
	return r.List(ctx, filter)
}

func (r *UserRepository) FindBannedUsers(ctx context.Context, filter user.ListFilter) ([]*user.User, int64, error) {
	status := user.AccountStatusBanned
	filter.Status = &status
	return r.List(ctx, filter)
}

func (r *UserRepository) FindUsersWithWarnings(ctx context.Context) ([]*user.User, error) {
	var users []*user.User
	err := r.db.WithContext(ctx).Where("warning_count > ?", 0).Find(&users).Error
	return users, err
}

func (r *UserRepository) FindFlaggedUsers(ctx context.Context, minFlags int, filter user.ListFilter) ([]*user.User, int64, error) {
	var users []*user.User
	var total int64
	
	query := r.db.WithContext(ctx).Model(&user.User{}).Where("flag_count >= ?", minFlags)
	query = r.applyFilters(query, filter)
	
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	
	query = r.applySorting(query, filter)
	query = query.Limit(filter.GetLimit()).Offset(filter.GetOffset())
	
	if err := query.Find(&users).Error; err != nil {
		return nil, 0, err
	}
	
	return users, total, nil
}

// ============================================================================
// BUSINESS QUERIES - REFERRALS
// ============================================================================

func (r *UserRepository) FindUsersByReferrer(ctx context.Context, referrerID string) ([]*user.User, error) {
	var users []*user.User
	err := r.db.WithContext(ctx).Where("referred_by = ?", referrerID).Find(&users).Error
	return users, err
}

func (r *UserRepository) CountReferrals(ctx context.Context, referrerID string) (int, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&user.User{}).Where("referred_by = ?", referrerID).Count(&count).Error
	return int(count), err
}

// ============================================================================
// ANALYTICS QUERIES - COUNTS
// ============================================================================

func (r *UserRepository) Count(ctx context.Context) (int64, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&user.User{}).Count(&count).Error
	return count, err
}

func (r *UserRepository) CountByUserType(ctx context.Context, userType user.UserType) (int64, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&user.User{}).Where("user_type = ?", userType).Count(&count).Error
	return count, err
}

func (r *UserRepository) CountByStatus(ctx context.Context, status user.AccountStatus) (int64, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&user.User{}).Where("status = ?", status).Count(&count).Error
	return count, err
}

func (r *UserRepository) CountByCountry(ctx context.Context, country string) (int64, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&user.User{}).Where("location_country = ?", country).Count(&count).Error
	return count, err
}

func (r *UserRepository) CountVerified(ctx context.Context) (int64, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&user.User{}).Where("identity_verified = ?", true).Count(&count).Error
	return count, err
}

func (r *UserRepository) CountOnline(ctx context.Context) (int64, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&user.User{}).Where("is_online = ?", true).Count(&count).Error
	return count, err
}

func (r *UserRepository) CountCreatedBetween(ctx context.Context, start, end time.Time) (int64, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&user.User{}).
		Where("created_at BETWEEN ? AND ?", start, end).
		Count(&count).Error
	return count, err
}

// ============================================================================
// ANALYTICS QUERIES - AGGREGATIONS
// ============================================================================

func (r *UserRepository) GetUserStatistics(ctx context.Context) (*user.UserStatistics, error) {
	stats := &user.UserStatistics{
		GeneratedAt: time.Now(),
	}
	
	// Total users
	r.db.WithContext(ctx).Model(&user.User{}).Count(&stats.TotalUsers)
	
	// By status
	r.db.WithContext(ctx).Model(&user.User{}).Where("status = ?", user.AccountStatusActive).Count(&stats.ActiveUsers)
	r.db.WithContext(ctx).Model(&user.User{}).Where("status = ?", user.AccountStatusPending).Count(&stats.PendingUsers)
	r.db.WithContext(ctx).Model(&user.User{}).Where("status = ?", user.AccountStatusInactive).Count(&stats.InactiveUsers)
	r.db.WithContext(ctx).Model(&user.User{}).Where("status = ?", user.AccountStatusSuspended).Count(&stats.SuspendedUsers)
	r.db.WithContext(ctx).Model(&user.User{}).Where("status = ?", user.AccountStatusBanned).Count(&stats.BannedUsers)
	r.db.WithContext(ctx).Model(&user.User{}).Where("status = ?", user.AccountStatusDeleted).Count(&stats.DeletedUsers)
	r.db.WithContext(ctx).Model(&user.User{}).Where("status = ?", user.AccountStatusRestricted).Count(&stats.RestrictedUsers)
	
	// By user type
	r.db.WithContext(ctx).Model(&user.User{}).Where("user_type = ?", user.UserTypeFreelancer).Count(&stats.TotalFreelancers)
	r.db.WithContext(ctx).Model(&user.User{}).Where("user_type = ?", user.UserTypeClient).Count(&stats.TotalClients)
	r.db.WithContext(ctx).Model(&user.User{}).Where("user_type = ?", user.UserTypeBoth).Count(&stats.TotalBoth)
	r.db.WithContext(ctx).Model(&user.User{}).Where("user_type IN ?", []user.UserType{user.UserTypeAdmin, user.UserTypeModerator, user.UserTypeSupport}).Count(&stats.TotalStaff)
	
	// Verification
	r.db.WithContext(ctx).Model(&user.User{}).Where("identity_verified = ?", true).Count(&stats.VerifiedUsers)
	r.db.WithContext(ctx).Model(&user.User{}).Where("identity_verified = ?", false).Count(&stats.UnverifiedUsers)
	r.db.WithContext(ctx).Model(&user.User{}).Where("email_verified = ?", true).Count(&stats.EmailVerifiedUsers)
	r.db.WithContext(ctx).Model(&user.User{}).Where("phone_verified = ?", true).Count(&stats.PhoneVerifiedUsers)
	r.db.WithContext(ctx).Model(&user.User{}).Where("identity_verified = ?", true).Count(&stats.IdentityVerifiedUsers)
	
	if stats.TotalUsers > 0 {
		stats.VerificationRate = float64(stats.VerifiedUsers) / float64(stats.TotalUsers) * 100
	}
	
	// Activity
	r.db.WithContext(ctx).Model(&user.User{}).Where("is_online = ?", true).Count(&stats.OnlineUsers)
	
	today := time.Now().Truncate(24 * time.Hour)
	r.db.WithContext(ctx).Model(&user.User{}).Where("last_login_at >= ?", today).Count(&stats.ActiveToday)
	
	weekAgo := time.Now().AddDate(0, 0, -7)
	r.db.WithContext(ctx).Model(&user.User{}).Where("last_login_at >= ?", weekAgo).Count(&stats.ActiveThisWeek)
	
	monthAgo := time.Now().AddDate(0, -1, 0)
	r.db.WithContext(ctx).Model(&user.User{}).Where("last_login_at >= ?", monthAgo).Count(&stats.ActiveThisMonth)
	
	// Averages
	r.db.WithContext(ctx).Model(&user.User{}).Select("AVG(rating)").Row().Scan(&stats.AverageRating)
	r.db.WithContext(ctx).Model(&user.User{}).Select("AVG(profile_completeness)").Row().Scan(&stats.AverageCompleteness)
	r.db.WithContext(ctx).Model(&user.User{}).Select("AVG(success_rate)").Row().Scan(&stats.AverageSuccessRate)
	r.db.WithContext(ctx).Model(&user.User{}).Select("AVG(completed_jobs)").Row().Scan(&stats.AverageCompletedJobs)
	
	// Badges
	r.db.WithContext(ctx).Model(&user.User{}).Where("is_top_rated = ?", true).Count(&stats.TopRatedCount)
	r.db.WithContext(ctx).Model(&user.User{}).Where("is_rising_talent = ?", true).Count(&stats.RisingTalentCount)
	r.db.WithContext(ctx).Model(&user.User{}).Where("is_expert_vetted = ?", true).Count(&stats.ExpertVettedCount)
	r.db.WithContext(ctx).Model(&user.User{}).Where("is_featured = ?", true).Count(&stats.FeaturedCount)
	
	// Growth
	r.db.WithContext(ctx).Model(&user.User{}).Where("created_at >= ?", today).Count(&stats.UsersCreatedToday)
	yesterday := today.AddDate(0, 0, -1)
	r.db.WithContext(ctx).Model(&user.User{}).Where("created_at >= ? AND created_at < ?", yesterday, today).Count(&stats.UsersCreatedYesterday)
	r.db.WithContext(ctx).Model(&user.User{}).Where("created_at >= ?", weekAgo).Count(&stats.UsersCreatedThisWeek)
	lastWeek := weekAgo.AddDate(0, 0, -7)
	r.db.WithContext(ctx).Model(&user.User{}).Where("created_at >= ? AND created_at < ?", lastWeek, weekAgo).Count(&stats.UsersCreatedLastWeek)
	r.db.WithContext(ctx).Model(&user.User{}).Where("created_at >= ?", monthAgo).Count(&stats.UsersCreatedThisMonth)
	lastMonth := monthAgo.AddDate(0, -1, 0)
	r.db.WithContext(ctx).Model(&user.User{}).Where("created_at >= ? AND created_at < ?", lastMonth, monthAgo).Count(&stats.UsersCreatedLastMonth)
	
	// Calculate growth rates
	if stats.UsersCreatedYesterday > 0 {
		stats.GrowthRateDaily = float64(stats.UsersCreatedToday-stats.UsersCreatedYesterday) / float64(stats.UsersCreatedYesterday) * 100
	}
	if stats.UsersCreatedLastWeek > 0 {
		stats.GrowthRateWeekly = float64(stats.UsersCreatedThisWeek-stats.UsersCreatedLastWeek) / float64(stats.UsersCreatedLastWeek) * 100
	}
	if stats.UsersCreatedLastMonth > 0 {
		stats.GrowthRateMonthly = float64(stats.UsersCreatedThisMonth-stats.UsersCreatedLastMonth) / float64(stats.UsersCreatedLastMonth) * 100
	}
	
	// Moderation
	r.db.WithContext(ctx).Model(&user.User{}).Where("warning_count > ?", 0).Count(&stats.UsersWithWarnings)
	r.db.WithContext(ctx).Model(&user.User{}).Where("flag_count > ?", 0).Count(&stats.FlaggedUsers)
	r.db.WithContext(ctx).Model(&user.User{}).Select("SUM(warning_count)").Row().Scan(&stats.TotalWarnings)
	r.db.WithContext(ctx).Model(&user.User{}).Select("SUM(suspension_count)").Row().Scan(&stats.TotalSuspensions)
	r.db.WithContext(ctx).Model(&user.User{}).Select("SUM(ban_count)").Row().Scan(&stats.TotalBans)
	
	// Referrals
	r.db.WithContext(ctx).Model(&user.User{}).Where("referral_count > ?", 0).Count(&stats.UsersWithReferrals)
	r.db.WithContext(ctx).Model(&user.User{}).Select("SUM(referral_count)").Row().Scan(&stats.TotalReferrals)
	if stats.UsersWithReferrals > 0 {
		stats.AverageReferrals = float64(stats.TotalReferrals) / float64(stats.UsersWithReferrals)
	}
	
	// Financial
	r.db.WithContext(ctx).Model(&user.User{}).Select("SUM(total_earnings)").Row().Scan(&stats.TotalEarningsAll)
	r.db.WithContext(ctx).Model(&user.User{}).Select("SUM(total_spent)").Row().Scan(&stats.TotalSpendingAll)
	if stats.TotalFreelancers > 0 {
		stats.AverageEarnings = stats.TotalEarningsAll / float64(stats.TotalFreelancers)
	}
	if stats.TotalClients > 0 {
		stats.AverageSpending = stats.TotalSpendingAll / float64(stats.TotalClients)
	}
	
	return stats, nil
}

func (r *UserRepository) GetUserGrowthStats(ctx context.Context, days int) (*user.UserGrowthStats, error) {
	endDate := time.Now()
	startDate := endDate.AddDate(0, 0, -days)
	
	stats := &user.UserGrowthStats{
		Period:    fmt.Sprintf("%d days", days),
		StartDate: startDate,
		EndDate:   endDate,
	}
	
	// Total users at end date
	r.db.WithContext(ctx).Model(&user.User{}).Count(&stats.TotalUsers)
	
	// New users in period
	r.db.WithContext(ctx).Model(&user.User{}).
		Where("created_at BETWEEN ? AND ?", startDate, endDate).
		Count(&stats.NewUsers)
	
	// Deleted users in period
	r.db.WithContext(ctx).Model(&user.User{}).
		Where("deleted_at BETWEEN ? AND ?", startDate, endDate).
		Count(&stats.DeletedUsers)
	
	stats.NetGrowth = stats.NewUsers - stats.DeletedUsers
	
	// Calculate rates
	usersAtStart := stats.TotalUsers - stats.NetGrowth
	if usersAtStart > 0 {
		stats.GrowthRate = float64(stats.NetGrowth) / float64(usersAtStart) * 100
		stats.ChurnRate = float64(stats.DeletedUsers) / float64(usersAtStart) * 100
	}
	
	// Active users
	r.db.WithContext(ctx).Model(&user.User{}).
		Where("last_login_at BETWEEN ? AND ?", startDate, endDate).
		Count(&stats.ActiveUsers)
	
	if stats.TotalUsers > 0 {
		stats.RetentionRate = float64(stats.ActiveUsers) / float64(stats.TotalUsers) * 100
	}
	
	// New by type
	r.db.WithContext(ctx).Model(&user.User{}).
		Where("created_at BETWEEN ? AND ? AND user_type = ?", startDate, endDate, user.UserTypeFreelancer).
		Count(&stats.NewFreelancers)
	
	r.db.WithContext(ctx).Model(&user.User{}).
		Where("created_at BETWEEN ? AND ? AND user_type = ?", startDate, endDate, user.UserTypeClient).
		Count(&stats.NewClients)
	
	// Verified users
	r.db.WithContext(ctx).Model(&user.User{}).
		Where("created_at BETWEEN ? AND ? AND identity_verified = ?", startDate, endDate, true).
		Count(&stats.VerifiedUsers)
	
	if stats.NewUsers > 0 {
		stats.VerificationRate = float64(stats.VerifiedUsers) / float64(stats.NewUsers) * 100
	}
	
	return stats, nil
}

func (r *UserRepository) GetTopCountries(ctx context.Context, limit int) ([]user.CountryStats, error) {
	var results []struct {
		Country        string
		CountryCode    string
		UserCount      int64
		FreelancerCount int64
		ClientCount    int64
		AverageRating  float64
	}
	
	err := r.db.WithContext(ctx).Model(&user.User{}).
		Select("location_country as country, location_country_code as country_code, COUNT(*) as user_count, "+
			"COUNT(CASE WHEN user_type = ? THEN 1 END) as freelancer_count, "+
			"COUNT(CASE WHEN user_type = ? THEN 1 END) as client_count, "+
			"AVG(rating) as average_rating", user.UserTypeFreelancer, user.UserTypeClient).
		Group("location_country, location_country_code").
		Order("user_count DESC").
		Limit(limit).
		Scan(&results).Error
	
	if err != nil {
		return nil, err
	}
	
	var total int64
	r.db.WithContext(ctx).Model(&user.User{}).Count(&total)
	
	stats := make([]user.CountryStats, len(results))
	for i, result := range results {
		stats[i] = user.CountryStats{
			Country:         result.Country,
			CountryCode:     result.CountryCode,
			UserCount:       result.UserCount,
			FreelancerCount: result.FreelancerCount,
			ClientCount:     result.ClientCount,
			AverageRating:   result.AverageRating,
		}
		if total > 0 {
			stats[i].Percentage = float64(result.UserCount) / float64(total) * 100
		}
	}
	
	return stats, nil
}

func (r *UserRepository) GetAverageRatingByUserType(ctx context.Context, userType user.UserType) (float64, error) {
	var avgRating float64
	err := r.db.WithContext(ctx).Model(&user.User{}).
		Where("user_type = ?", userType).
		Select("AVG(rating)").
		Row().Scan(&avgRating)
	return avgRating, err
}

// ============================================================================
// BATCH OPERATIONS
// ============================================================================

func (r *UserRepository) UpdateStatusBatch(ctx context.Context, ids []string, status user.AccountStatus) error {
	return r.db.WithContext(ctx).Model(&user.User{}).
		Where("id IN ?", ids).
		Updates(map[string]interface{}{
			"status":     status,
			"updated_at": time.Now(),
		}).Error
}

func (r *UserRepository) VerifyEmailBatch(ctx context.Context, ids []string) error {
	now := time.Now().Unix()
	return r.db.WithContext(ctx).Model(&user.User{}).
		Where("id IN ?", ids).
		Updates(map[string]interface{}{
			"email_verified":    true,
			"email_verified_at": &now,
			"updated_at":        time.Now(),
		}).Error
}

func (r *UserRepository) DeleteBatch(ctx context.Context, ids []string, deletedBy string) error {
	now := time.Now()
	return r.db.WithContext(ctx).Model(&user.User{}).
		Where("id IN ?", ids).
		Updates(map[string]interface{}{
			"deleted_at": &now,
			"deleted_by": deletedBy,
			"status":     user.AccountStatusDeleted,
			"updated_at": now,
		}).Error
}

// ============================================================================
// TRANSACTION SUPPORT
// ============================================================================

func (r *UserRepository) WithTransaction(ctx context.Context, fn func(ctx context.Context) error) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		return fn(ctx)
	})
}

// ============================================================================
// HELPER METHODS
// ============================================================================

func (r *UserRepository) applyFilters(query *gorm.DB, filter user.ListFilter) *gorm.DB {
	// User type filter
	if filter.UserType != nil {
		query = query.Where("user_type = ?", *filter.UserType)
	}
	
	// Status filter
	if filter.Status != nil {
		query = query.Where("status = ?", *filter.Status)
	}
	
	// Verification status filter
	if filter.VerificationStatus != nil {
		query = query.Where("verification_status = ?", *filter.VerificationStatus)
	}
	
	// Location filters
	if filter.Country != "" {
		query = query.Where("location_country = ?", filter.Country)
	}
	if filter.CountryCode != "" {
		query = query.Where("location_country_code = ?", filter.CountryCode)
	}
	if filter.City != "" {
		query = query.Where("location_city = ?", filter.City)
	}
	if filter.Timezone != "" {
		query = query.Where("location_timezone = ?", filter.Timezone)
	}
	
	// Rating filters
	if filter.MinRating > 0 {
		query = query.Where("rating >= ?", filter.MinRating)
	}
	if filter.MaxRating > 0 {
		query = query.Where("rating <= ?", filter.MaxRating)
	}
	
	// Performance filters
	if filter.MinCompletedJobs > 0 {
		query = query.Where("completed_jobs >= ?", filter.MinCompletedJobs)
	}
	if filter.MinSuccessRate > 0 {
		query = query.Where("success_rate >= ?", filter.MinSuccessRate)
	}
	if filter.MinReviews > 0 {
		query = query.Where("total_reviews >= ?", filter.MinReviews)
	}
	
	// Verification filters
	if filter.EmailVerified != nil {
		query = query.Where("email_verified = ?", *filter.EmailVerified)
	}
	if filter.PhoneVerified != nil {
		query = query.Where("phone_verified = ?", *filter.PhoneVerified)
	}
	if filter.IdentityVerified != nil {
		query = query.Where("identity_verified = ?", *filter.IdentityVerified)
	}
	
	// Badge filters
	if filter.IsFeatured != nil {
		query = query.Where("is_featured = ?", *filter.IsFeatured)
	}
	if filter.IsTopRated != nil {
		query = query.Where("is_top_rated = ?", *filter.IsTopRated)
	}
	if filter.IsRisingTalent != nil {
		query = query.Where("is_rising_talent = ?", *filter.IsRisingTalent)
	}
	if filter.IsExpertVetted != nil {
		query = query.Where("is_expert_vetted = ?", *filter.IsExpertVetted)
	}
	
	// Activity filters
	if filter.IsOnline != nil {
		query = query.Where("is_online = ?", *filter.IsOnline)
	}
	if filter.SearchableOnly {
		query = query.Where("searchable_profile = ?", true)
	}
	if filter.AcceptingWork != nil {
		query = query.Where("accepting_work = ?", *filter.AcceptingWork)
	}
	if filter.AvailabilityStatus != nil {
		query = query.Where("availability_status = ?", *filter.AvailabilityStatus)
	}
	
	// Profile completeness
	if filter.MinProfileCompleteness > 0 {
		query = query.Where("profile_completeness >= ?", filter.MinProfileCompleteness)
	}
	if filter.ProfileCompleted != nil {
		query = query.Where("profile_completed = ?", *filter.ProfileCompleted)
	}
	
	// Date filters
	if filter.CreatedAfter != nil {
		query = query.Where("created_at >= ?", *filter.CreatedAfter)
	}
	if filter.CreatedBefore != nil {
		query = query.Where("created_at <= ?", *filter.CreatedBefore)
	}
	if filter.UpdatedAfter != nil {
		query = query.Where("updated_at >= ?", *filter.UpdatedAfter)
	}
	if filter.UpdatedBefore != nil {
		query = query.Where("updated_at <= ?", *filter.UpdatedBefore)
	}
	if filter.LastLoginAfter != nil {
		query = query.Where("last_login_at >= ?", *filter.LastLoginAfter)
	}
	if filter.LastLoginBefore != nil {
		query = query.Where("last_login_at <= ?", *filter.LastLoginBefore)
	}
	
	// Special filters
	if filter.ReferredBy != "" {
		query = query.Where("referred_by = ?", filter.ReferredBy)
	}
	if filter.HasWarnings {
		query = query.Where("warning_count > ?", 0)
	}
	if filter.IsSuspended {
		query = query.Where("status = ?", user.AccountStatusSuspended)
	}
	if filter.IsBanned {
		query = query.Where("status = ?", user.AccountStatusBanned)
	}
	
	// Deleted filter
	if !filter.IncludeDeleted {
		query = query.Where("deleted_at IS NULL")
	}
	
	return query
}

func (r *UserRepository) applySorting(query *gorm.DB, filter user.ListFilter) *gorm.DB {
	orderClause := filter.SortBy + " " + filter.SortOrder
	return query.Order(orderClause)
}

func (r *UserRepository) ExistsByReferralCode(ctx context.Context, code string) (bool, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&user.User{}).Where("referral_code = ?", code).Count(&count).Error
	return count > 0, err
}

// ============================================================================
// BUSINESS QUERIES - FILTERING & SORTING
// ============================================================================

func (r *UserRepository) FindByUserType(ctx context.Context, userType user.UserType, filter user.ListFilter) ([]*user.User, int64, error) {
	filter.UserType = &userType
	return r.List(ctx, filter)
}

func (r *UserRepository) FindByStatus(ctx context.Context, status user.AccountStatus, filter user.ListFilter) ([]*user.User, int64, error) {
	filter.Status = &status
	return r.List(ctx, filter)
}

func (r *UserRepository) FindByCountry(ctx context.Context, country string, filter user.ListFilter) ([]*user.User, int64, error) {
	filter.Country = country
	return r.List(ctx, filter)
}

func (r *UserRepository) FindByCity(ctx context.Context, city string, filter user.ListFilter) ([]*user.User, int64, error) {
	filter.City = city
	return r.List(ctx, filter)
}

func (r *UserRepository) FindByLocation(ctx context.Context, city, country string, filter user.ListFilter) ([]*user.User, int64, error) {
	filter.City = city
	filter.Country = country
	return r.List(ctx, filter)
}

// ============================================================================
// BUSINESS QUERIES - SPECIAL LISTS
// ============================================================================

func (r *UserRepository) FindTopRatedFreelancers(ctx context.Context, limit int) ([]*user.User, error) {
	var users []*user.User
	err := r.db.WithContext(ctx).
		Where("user_type IN ? AND rating >= ?", []user.UserType{user.UserTypeFreelancer, user.UserTypeBoth}, 4.5).
		Order("rating DESC, total_reviews DESC").
		Limit(limit).
		Find(&users).Error
	return users, err
}

func (r *UserRepository) FindTopRatedClients(ctx context.Context, limit int) ([]*user.User, error) {
	var users []*user.User
	err := r.db.WithContext(ctx).
		Where("user_type IN ? AND rating >= ?", []user.UserType{user.UserTypeClient, user.UserTypeBoth}, 4.5).
		Order("rating DESC, total_reviews DESC").
		Limit(limit).
		Find(&users).Error
	return users, err
}

func (r *UserRepository) FindFeaturedUsers(ctx context.Context, userType user.UserType, limit int) ([]*user.User, error) {
	var users []*user.User
	query := r.db.WithContext(ctx).Where("is_featured = ?", true)
	
	if userType != "" {
		query = query.Where("user_type = ?", userType)
	}
	
	err := query.Order("rating DESC").Limit(limit).Find(&users).Error
	return users, err
}

func (r *UserRepository) FindRisingTalent(ctx context.Context, limit int) ([]*user.User, error) {
	var users []*user.User
	err := r.db.WithContext(ctx).
		Where("is_rising_talent = ? AND user_type IN ?", true, []user.UserType{user.UserTypeFreelancer, user.UserTypeBoth}).
		Order("created_at DESC").
		Limit(limit).
		Find(&users).Error
	return users, err
}

func (r *UserRepository) FindExpertVettedFreelancers(ctx context.Context, limit int) ([]*user.User, error) {
	var users []*user.User
	err := r.db.WithContext(ctx).
		Where("is_expert_vetted = ? AND user_type IN ?", true, []user.UserType{user.UserTypeFreelancer, user.UserTypeBoth}).
		Order("rating DESC").
		Limit(limit).
		Find(&users).Error
	return users, err
}

func (r *UserRepository) FindOnlineUsers(ctx context.Context, userType user.UserType) ([]*user.User, error) {
	var users []*user.User
	query := r.db.WithContext(ctx).Where("is_online = ?", true)
	
	if userType != "" {
		query = query.Where("user_type = ?", userType)
	}
	
	err := query.Find(&users).Error
	return users, err
}

func (r *UserRepository) FindRecentlyActive(ctx context.Context, hours int, filter user.ListFilter) ([]*user.User, int64, error) {
	threshold := time.Now().Add(-time.Duration(hours) * time.Hour)
	filter.LastLoginAfter = &threshold
	return r.List(ctx, filter)
}

func (r *UserRepository) FindInactiveUsers(ctx context.Context, days int) ([]*user.User, error) {
	var users []*user.User
	threshold := time.Now().AddDate(0, 0, -days)
	err := r.db.WithContext(ctx).
		Where("last_login_at < ? OR last_login_at IS NULL", threshold).
		Find(&users).Error
	return users, err
}

func (r *UserRepository) FindNewUsers(ctx context.Context, days int, filter user.ListFilter) ([]*user.User, int64, error) {
	threshold := time.Now().AddDate(0, 0, -days)
	filter.CreatedAfter = &threshold
	return r.List(ctx, filter)
}
