// internal/domain/user/repository.go
package user

import (
	"context"
	"time"
)

// Repository defines the interface for user data persistence
// This follows the repository pattern, keeping domain logic independent of infrastructure
type Repository interface {
	// ========================================================================
	// CREATE OPERATIONS
	// ========================================================================
	
	// Create creates a new user in the database
	Create(ctx context.Context, user *User) error
	
	// CreateBatch creates multiple users in a single transaction
	CreateBatch(ctx context.Context, users []*User) error
	
	// ========================================================================
	// READ OPERATIONS - SINGLE ENTITY
	// ========================================================================
	
	// FindByID retrieves a user by their internal UUID
	FindByID(ctx context.Context, id string) (*User, error)
	
	// FindByKeycloakID retrieves a user by their Keycloak ID (sub claim)
	FindByKeycloakID(ctx context.Context, keycloakID string) (*User, error)
	
	// FindByEmail retrieves a user by their email address
	FindByEmail(ctx context.Context, email string) (*User, error)
	
	// FindByUsername retrieves a user by their username
	FindByUsername(ctx context.Context, username string) (*User, error)
	
	// FindByReferralCode retrieves a user by their referral code
	FindByReferralCode(ctx context.Context, code string) (*User, error)
	
	// ========================================================================
	// READ OPERATIONS - MULTIPLE ENTITIES
	// ========================================================================
	
	// FindByIDs retrieves multiple users by their IDs
	FindByIDs(ctx context.Context, ids []string) ([]*User, error)
	
	// FindByKeycloakIDs retrieves multiple users by their Keycloak IDs
	FindByKeycloakIDs(ctx context.Context, keycloakIDs []string) ([]*User, error)
	
	// FindByEmails retrieves multiple users by their email addresses
	FindByEmails(ctx context.Context, emails []string) ([]*User, error)
	
	// ========================================================================
	// LIST & SEARCH OPERATIONS
	// ========================================================================
	
	// List retrieves a paginated list of users with filters
	List(ctx context.Context, filter ListFilter) ([]*User, int64, error)
	
	// Search performs full-text search on users
	Search(ctx context.Context, query string, filter ListFilter) ([]*User, int64, error)
	
	// FindAll retrieves all users (use with caution, prefer List)
	FindAll(ctx context.Context) ([]*User, error)
	
	// ========================================================================
	// UPDATE OPERATIONS - FULL ENTITY
	// ========================================================================
	
	// Update updates an existing user's information
	Update(ctx context.Context, user *User) error
	
	// UpdateBatch updates multiple users in a single transaction
	UpdateBatch(ctx context.Context, users []*User) error
	
	// ========================================================================
	// UPDATE OPERATIONS - SPECIFIC FIELDS
	// ========================================================================
	
	// UpdateStatus updates only the user's status
	UpdateStatus(ctx context.Context, id string, status AccountStatus) error
	
	// UpdateVerificationStatus updates only the verification status
	UpdateVerificationStatus(ctx context.Context, id string, status VerificationStatus) error
	
	// UpdateLastSeen updates the last seen timestamp
	UpdateLastSeen(ctx context.Context, id string) error
	
	// UpdateOnlineStatus updates the online status
	UpdateOnlineStatus(ctx context.Context, id string, isOnline bool) error
	
	// UpdateProfileCompleteness updates profile completeness percentage
	UpdateProfileCompleteness(ctx context.Context, id string, percentage int) error
	
	// UpdateRating updates user rating and review count
	UpdateRating(ctx context.Context, id string, rating float64, totalReviews int) error
	
	// UpdateStats updates user statistics
	UpdateStats(ctx context.Context, id string, completedJobs, totalJobs int, successRate float64) error
	
	// UpdateEarnings updates total earnings (for freelancers)
	UpdateEarnings(ctx context.Context, id string, amount float64) error
	
	// UpdateSpending updates total spending (for clients)
	UpdateSpending(ctx context.Context, id string, amount float64) error
	
	// UpdateBalance updates current wallet balance
	UpdateBalance(ctx context.Context, id string, amount float64) error
	
	// ========================================================================
	// UPDATE OPERATIONS - SECURITY & ACTIVITY
	// ========================================================================
	
	// RecordLogin records a successful login
	RecordLogin(ctx context.Context, id, ipAddress, userAgent string) error
	
	// IncrementLoginCount increments the login count
	IncrementLoginCount(ctx context.Context, id string) error
	
	// IncrementFailedLoginAttempts increments failed login attempts
	IncrementFailedLoginAttempts(ctx context.Context, id string) error
	
	// ResetFailedLoginAttempts resets failed login attempts to zero
	ResetFailedLoginAttempts(ctx context.Context, id string) error
	
	// LockAccount locks the account until specified time
	LockAccount(ctx context.Context, id string, until time.Time) error
	
	// UnlockAccount unlocks the account
	UnlockAccount(ctx context.Context, id string) error
	
	// ========================================================================
	// UPDATE OPERATIONS - VERIFICATION
	// ========================================================================
	
	// VerifyEmail marks email as verified
	VerifyEmail(ctx context.Context, id string) error
	
	// VerifyPhone marks phone as verified
	VerifyPhone(ctx context.Context, id string) error
	
	// VerifyIdentity marks identity as verified
	VerifyIdentity(ctx context.Context, id string) error
	
	// ========================================================================
	// UPDATE OPERATIONS - BADGES & ACHIEVEMENTS
	// ========================================================================
	
	// AssignBadge assigns a badge to a user
	AssignBadge(ctx context.Context, id string, badge BadgeType) error
	
	// RemoveBadge removes a badge from a user
	RemoveBadge(ctx context.Context, id string, badge BadgeType) error
	
	// SetFeatured sets the featured flag
	SetFeatured(ctx context.Context, id string, featured bool) error
	
	// SetTopRated sets the top rated flag
	SetTopRated(ctx context.Context, id string, topRated bool) error
	
	// ========================================================================
	// UPDATE OPERATIONS - MODERATION
	// ========================================================================
	
	// IncrementWarningCount increments warning count
	IncrementWarningCount(ctx context.Context, id string) error
	
	// IncrementSuspensionCount increments suspension count
	IncrementSuspensionCount(ctx context.Context, id string) error
	
	// IncrementBanCount increments ban count
	IncrementBanCount(ctx context.Context, id string) error
	
	// IncrementFlagCount increments flag count
	IncrementFlagCount(ctx context.Context, id string) error
	
	// AddNote appends a note to user's notes
	AddNote(ctx context.Context, id, note string) error
	
	// AddTag adds a tag to user
	AddTag(ctx context.Context, id, tag string) error
	
	// RemoveTag removes a tag from user
	RemoveTag(ctx context.Context, id, tag string) error
	
	// ========================================================================
	// DELETE OPERATIONS
	// ========================================================================
	
	// Delete soft-deletes a user (sets DeletedAt timestamp)
	Delete(ctx context.Context, id string) error
	
	// SoftDelete soft-deletes a user with deletion metadata
	SoftDelete(ctx context.Context, id, deletedBy string) error
	
	// HardDelete permanently deletes a user from the database
	HardDelete(ctx context.Context, id string) error
	
	// RestoreDeleted restores a soft-deleted user
	RestoreDeleted(ctx context.Context, id string) error
	
	// ========================================================================
	// EXISTENCE CHECKS
	// ========================================================================
	
	// ExistsByID checks if a user exists by ID
	ExistsByID(ctx context.Context, id string) (bool, error)
	
	// ExistsByKeycloakID checks if a user with the given Keycloak ID exists
	ExistsByKeycloakID(ctx context.Context, keycloakID string) (bool, error)
	
	// ExistsByEmail checks if a user with the given email exists
	ExistsByEmail(ctx context.Context, email string) (bool, error)
	
	// ExistsByUsername checks if a user with the given username exists
	ExistsByUsername(ctx context.Context, username string) (bool, error)
	
	// ExistsByReferralCode checks if a referral code is taken
	ExistsByReferralCode(ctx context.Context, code string) (bool, error)
	
	// ========================================================================
	// BUSINESS QUERIES - FILTERING & SORTING
	// ========================================================================
	
	// FindByUserType retrieves users by type (freelancer, client, both)
	FindByUserType(ctx context.Context, userType UserType, filter ListFilter) ([]*User, int64, error)
	
	// FindByStatus retrieves users by account status
	FindByStatus(ctx context.Context, status AccountStatus, filter ListFilter) ([]*User, int64, error)
	
	// FindByCountry retrieves users by country
	FindByCountry(ctx context.Context, country string, filter ListFilter) ([]*User, int64, error)
	
	// FindByCity retrieves users by city
	FindByCity(ctx context.Context, city string, filter ListFilter) ([]*User, int64, error)
	
	// FindByLocation retrieves users by city and country
	FindByLocation(ctx context.Context, city, country string, filter ListFilter) ([]*User, int64, error)
	
	// ========================================================================
	// BUSINESS QUERIES - SPECIAL LISTS
	// ========================================================================
	
	// FindTopRatedFreelancers retrieves top-rated freelancers
	FindTopRatedFreelancers(ctx context.Context, limit int) ([]*User, error)
	
	// FindTopRatedClients retrieves top-rated clients
	FindTopRatedClients(ctx context.Context, limit int) ([]*User, error)
	
	// FindFeaturedUsers retrieves featured users
	FindFeaturedUsers(ctx context.Context, userType UserType, limit int) ([]*User, error)
	
	// FindRisingTalent retrieves rising talent freelancers
	FindRisingTalent(ctx context.Context, limit int) ([]*User, error)
	
	// FindExpertVettedFreelancers retrieves expert-vetted freelancers
	FindExpertVettedFreelancers(ctx context.Context, limit int) ([]*User, error)
	
	// FindOnlineUsers retrieves currently online users
	FindOnlineUsers(ctx context.Context, userType UserType) ([]*User, error)
	
	// FindRecentlyActive retrieves recently active users (within last N hours)
	FindRecentlyActive(ctx context.Context, hours int, filter ListFilter) ([]*User, int64, error)
	
	// FindInactiveUsers retrieves users inactive for N days
	FindInactiveUsers(ctx context.Context, days int) ([]*User, error)
	
	// FindNewUsers retrieves users created within last N days
	FindNewUsers(ctx context.Context, days int, filter ListFilter) ([]*User, int64, error)
	
	// ========================================================================
	// BUSINESS QUERIES - VERIFICATION & COMPLIANCE
	// ========================================================================
	
	// FindUnverifiedUsers retrieves users with unverified email
	FindUnverifiedUsers(ctx context.Context, filter ListFilter) ([]*User, int64, error)
	
	// FindPendingVerification retrieves users with verification pending
	FindPendingVerification(ctx context.Context, filter ListFilter) ([]*User, int64, error)
	
	// FindVerifiedUsers retrieves fully verified users
	FindVerifiedUsers(ctx context.Context, filter ListFilter) ([]*User, int64, error)
	
	// ========================================================================
	// BUSINESS QUERIES - MODERATION
	// ========================================================================
	
	// FindSuspendedUsers retrieves suspended users
	FindSuspendedUsers(ctx context.Context, filter ListFilter) ([]*User, int64, error)
	
	// FindBannedUsers retrieves banned users
	FindBannedUsers(ctx context.Context, filter ListFilter) ([]*User, int64, error)
	
	// FindUsersWithWarnings retrieves users with warnings
	FindUsersWithWarnings(ctx context.Context) ([]*User, error)
	
	// FindFlaggedUsers retrieves users who have been flagged
	FindFlaggedUsers(ctx context.Context, minFlags int, filter ListFilter) ([]*User, int64, error)
	
	// ========================================================================
	// BUSINESS QUERIES - REFERRALS
	// ========================================================================
	
	// FindUsersByReferrer retrieves users referred by a specific user
	FindUsersByReferrer(ctx context.Context, referrerID string) ([]*User, error)
	
	// CountReferrals counts how many users were referred by a user
	CountReferrals(ctx context.Context, referrerID string) (int, error)
	
	// ========================================================================
	// ANALYTICS QUERIES - COUNTS
	// ========================================================================
	
	// Count counts total number of users
	Count(ctx context.Context) (int64, error)
	
	// CountByUserType counts users by type
	CountByUserType(ctx context.Context, userType UserType) (int64, error)
	
	// CountByStatus counts users by status
	CountByStatus(ctx context.Context, status AccountStatus) (int64, error)
	
	// CountByCountry counts users by country
	CountByCountry(ctx context.Context, country string) (int64, error)
	
	// CountVerified counts verified users
	CountVerified(ctx context.Context) (int64, error)
	
	// CountOnline counts currently online users
	CountOnline(ctx context.Context) (int64, error)
	
	// CountCreatedBetween counts users created between two dates
	CountCreatedBetween(ctx context.Context, start, end time.Time) (int64, error)
	
	// ========================================================================
	// ANALYTICS QUERIES - AGGREGATIONS
	// ========================================================================
	
	// GetUserStatistics retrieves comprehensive user statistics
	GetUserStatistics(ctx context.Context) (*UserStatistics, error)
	
	// GetUserGrowthStats retrieves user growth statistics
	GetUserGrowthStats(ctx context.Context, days int) (*UserGrowthStats, error)
	
	// GetTopCountries retrieves top countries by user count
	GetTopCountries(ctx context.Context, limit int) ([]CountryStats, error)
	
	// GetAverageRatingByUserType gets average rating by user type
	GetAverageRatingByUserType(ctx context.Context, userType UserType) (float64, error)
	
	// ========================================================================
	// BATCH OPERATIONS
	// ========================================================================
	
	// UpdateStatusBatch updates status for multiple users
	UpdateStatusBatch(ctx context.Context, ids []string, status AccountStatus) error
	
	// VerifyEmailBatch verifies email for multiple users
	VerifyEmailBatch(ctx context.Context, ids []string) error
	
	// DeleteBatch soft-deletes multiple users
	DeleteBatch(ctx context.Context, ids []string, deletedBy string) error
	
	// ========================================================================
	// TRANSACTION SUPPORT
	// ========================================================================
	
	// WithTransaction executes a function within a transaction
	WithTransaction(ctx context.Context, fn func(ctx context.Context) error) error
}

// ============================================================================
// FILTER TYPES
// ============================================================================

// ListFilter defines filtering options for list queries
type ListFilter struct {
	// Pagination
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
	Page   int `json:"page"`
	
	// Sorting
	SortBy    string `json:"sort_by"`    // Field to sort by
	SortOrder string `json:"sort_order"` // "asc" or "desc"
	
	// Filters
	UserType           *UserType          `json:"user_type,omitempty"`
	Status             *AccountStatus     `json:"status,omitempty"`
	VerificationStatus *VerificationStatus `json:"verification_status,omitempty"`
	Country            string             `json:"country,omitempty"`
	City               string             `json:"city,omitempty"`
	MinRating          float64            `json:"min_rating,omitempty"`
	MaxRating          float64            `json:"max_rating,omitempty"`
	EmailVerified      *bool              `json:"email_verified,omitempty"`
	IdentityVerified   *bool              `json:"identity_verified,omitempty"`
	IsFeatured         *bool              `json:"is_featured,omitempty"`
	IsTopRated         *bool              `json:"is_top_rated,omitempty"`
	IsOnline           *bool              `json:"is_online,omitempty"`
	SearchableOnly     bool               `json:"searchable_only,omitempty"`
	
	// Date filters
	CreatedAfter  *time.Time `json:"created_after,omitempty"`
	CreatedBefore *time.Time `json:"created_before,omitempty"`
	UpdatedAfter  *time.Time `json:"updated_after,omitempty"`
	UpdatedBefore *time.Time `json:"updated_before,omitempty"`
	
	// Include deleted
	IncludeDeleted bool `json:"include_deleted,omitempty"`
}

// NewListFilter creates a new ListFilter with defaults
func NewListFilter() ListFilter {
	return ListFilter{
		Limit:     20,
		Offset:    0,
		Page:      1,
		SortBy:    "created_at",
		SortOrder: "desc",
	}
}

// Validate validates the filter parameters
func (f *ListFilter) Validate() error {
	if f.Limit < 0 {
		return fmt.Errorf("limit cannot be negative")
	}
	if f.Limit > 100 {
		return fmt.Errorf("limit cannot exceed 100")
	}
	if f.Limit == 0 {
		f.Limit = 20
	}
	
	if f.Offset < 0 {
		return fmt.Errorf("offset cannot be negative")
	}
	
	if f.Page < 1 {
		f.Page = 1
	}
	
	// Calculate offset from page if not set
	if f.Offset == 0 && f.Page > 1 {
		f.Offset = (f.Page - 1) * f.Limit
	}
	
	if f.SortOrder != "" && f.SortOrder != "asc" && f.SortOrder != "desc" {
		return fmt.Errorf("sort_order must be 'asc' or 'desc'")
	}
	if f.SortOrder == "" {
		f.SortOrder = "desc"
	}
	
	// Validate sort field
	validSortFields := map[string]bool{
		"id": true, "username": true, "email": true, "created_at": true, 
		"updated_at": true, "rating": true, "total_reviews": true,
		"completed_jobs": true, "success_rate": true, "last_login_at": true,
		"last_seen_at": true, "profile_completeness": true,
	}
	if f.SortBy != "" && !validSortFields[f.SortBy] {
		return ErrInvalidSortField
	}
	if f.SortBy == "" {
		f.SortBy = "created_at"
	}
	
	if f.MinRating < 0 || f.MinRating > 5 {
		return fmt.Errorf("min_rating must be between 0 and 5")
	}
	if f.MaxRating < 0 || f.MaxRating > 5 {
		return fmt.Errorf("max_rating must be between 0 and 5")
	}
	if f.MinRating > f.MaxRating {
		return fmt.Errorf("min_rating cannot be greater than max_rating")
	}
	
	return nil
}

// ============================================================================
// STATISTICS TYPES
// ============================================================================

// UserStatistics contains comprehensive user statistics
type UserStatistics struct {
	TotalUsers          int64   `json:"total_users"`
	ActiveUsers         int64   `json:"active_users"`
	PendingUsers        int64   `json:"pending_users"`
	SuspendedUsers      int64   `json:"suspended_users"`
	BannedUsers         int64   `json:"banned_users"`
	DeletedUsers        int64   `json:"deleted_users"`
	
	TotalFreelancers    int64   `json:"total_freelancers"`
	TotalClients        int64   `json:"total_clients"`
	
	VerifiedUsers       int64   `json:"verified_users"`
	UnverifiedUsers     int64   `json:"unverified_users"`
	
	OnlineUsers         int64   `json:"online_users"`
	
	AverageRating       float64 `json:"average_rating"`
	AverageCompleteness int     `json:"average_completeness"`
	
	TopRatedCount       int64   `json:"top_rated_count"`
	FeaturedCount       int64   `json:"featured_count"`
	
	UsersCreatedToday   int64   `json:"users_created_today"`
	UsersCreatedThisWeek int64  `json:"users_created_this_week"`
	UsersCreatedThisMonth int64 `json:"users_created_this_month"`
}

// UserGrowthStats contains user growth statistics over time
type UserGrowthStats struct {
	Period            string  `json:"period"` // "daily", "weekly", "monthly"
	TotalUsers        int64   `json:"total_users"`
	NewUsers          int64   `json:"new_users"`
	GrowthRate        float64 `json:"growth_rate"` // Percentage
	ActiveUsers       int64   `json:"active_users"`
	RetentionRate     float64 `json:"retention_rate"` // Percentage
	ChurnRate         float64 `json:"churn_rate"` // Percentage
}

// CountryStats contains user statistics by country
type CountryStats struct {
	Country      string `json:"country"`
	CountryCode  string `json:"country_code"`
	UserCount    int64  `json:"user_count"`
	Percentage   float64 `json:"percentage"`
}

// ============================================================================
// QUERY BUILDER HELPERS
// ============================================================================

// SearchQuery defines search parameters
type SearchQuery struct {
	Query      string     `json:"query"`
	Fields     []string   `json:"fields"` // Fields to search in
	UserType   *UserType  `json:"user_type,omitempty"`
	MinRating  float64    `json:"min_rating,omitempty"`
	Country    string     `json:"country,omitempty"`
	Skills     []string   `json:"skills,omitempty"` // For future integration
	Limit      int        `json:"limit"`
	Offset     int        `json:"offset"`
}

// NewSearchQuery creates a new search query with defaults
func NewSearchQuery(query string) SearchQuery {
	return SearchQuery{
		Query: query,
		Fields: []string{
			"username", "full_name", "email", 
			"bio", "tagline", "title", "overview",
		},
		Limit:  20,
		Offset: 0,
	}
}

// Validate validates search query parameters
func (sq *SearchQuery) Validate() error {
	if sq.Query == "" {
		return ErrInvalidSearchQuery
	}
	
	if len(sq.Query) < 2 {
		return fmt.Errorf("search query too short (minimum 2 characters)")
	}
	
	if len(sq.Query) > 200 {
		return fmt.Errorf("search query too long (maximum 200 characters)")
	}
	
	if sq.Limit < 1 {
		sq.Limit = 20
	}
	if sq.Limit > 100 {
		return fmt.Errorf("limit cannot exceed 100")
	}
	
	if sq.Offset < 0 {
		return fmt.Errorf("offset cannot be negative")
	}
	
	if sq.MinRating < 0 || sq.MinRating > 5 {
		return fmt.Errorf("min_rating must be between 0 and 5")
	}
	
	return nil
}

// ============================================================================
// BULK OPERATION TYPES
// ============================================================================

// BulkUpdateRequest represents a bulk update request
type BulkUpdateRequest struct {
	UserIDs   []string               `json:"user_ids"`
	Updates   map[string]interface{} `json:"updates"`
	UpdatedBy string                 `json:"updated_by"`
}

// Validate validates bulk update request
func (bur *BulkUpdateRequest) Validate() error {
	if len(bur.UserIDs) == 0 {
		return ErrEmptyBatchOperation
	}
	
	if len(bur.UserIDs) > 1000 {
		return ErrBatchSizeTooLarge
	}
	
	if len(bur.Updates) == 0 {
		return fmt.Errorf("no updates specified")
	}
	
	return nil
}

// BulkOperationResult represents the result of a bulk operation
type BulkOperationResult struct {
	TotalRequested int      `json:"total_requested"`
	Successful     int      `json:"successful"`
	Failed         int      `json:"failed"`
	Errors         []string `json:"errors,omitempty"`
}

// ============================================================================
// ADVANCED QUERY TYPES
// ============================================================================

// AggregationQuery defines aggregation parameters
type AggregationQuery struct {
	GroupBy   []string               `json:"group_by"`
	Functions map[string]string      `json:"functions"` // e.g., {"rating": "avg", "users": "count"}
	Having    map[string]interface{} `json:"having,omitempty"`
	OrderBy   string                 `json:"order_by,omitempty"`
	Limit     int                    `json:"limit,omitempty"`
}

// TimeSeriesQuery defines time-series aggregation parameters
type TimeSeriesQuery struct {
	Metric      string     `json:"metric"`      // What to measure (count, avg_rating, etc.)
	Interval    string     `json:"interval"`    // "hour", "day", "week", "month"
	StartDate   time.Time  `json:"start_date"`
	EndDate     time.Time  `json:"end_date"`
	GroupBy     []string   `json:"group_by,omitempty"` // Additional grouping
}

// ============================================================================
// HELPER FUNCTIONS
// ============================================================================

// DefaultListFilter returns a filter with sensible defaults
func DefaultListFilter() ListFilter {
	return ListFilter{
		Limit:          20,
		Offset:         0,
		Page:           1,
		SortBy:         "created_at",
		SortOrder:      "desc",
		SearchableOnly: false,
		IncludeDeleted: false,
	}
}

// NewPaginatedFilter creates a filter for paginated queries
func NewPaginatedFilter(page, limit int) ListFilter {
	if limit <= 0 || limit > 100 {
		limit = 20
	}
	if page < 1 {
		page = 1
	}
	
	return ListFilter{
		Limit:     limit,
		Offset:    (page - 1) * limit,
		Page:      page,
		SortBy:    "created_at",
		SortOrder: "desc",
	}
}