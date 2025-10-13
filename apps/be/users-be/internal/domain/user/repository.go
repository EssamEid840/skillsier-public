// internal/domain/user/repository.go
package user

import (
    "context"
)

type Repository interface {
    // Create operations
    Create(ctx context.Context, user *User) error
    CreateBatch(ctx context.Context, users []*User) error
    
    // Read operations
    FindByID(ctx context.Context, id string) (*User, error)
    FindByKeycloakID(ctx context.Context, keycloakID string) (*User, error)
    FindByEmail(ctx context.Context, email string) (*User, error)
    FindByUsername(ctx context.Context, username string) (*User, error)
    FindByReferralCode(ctx context.Context, code string) (*User, error)
    FindByIDs(ctx context.Context, ids []string) ([]*User, error)
    
    // List operations with filters
    List(ctx context.Context, filter ListFilter) ([]*User, int64, error)
    Search(ctx context.Context, query string, filter ListFilter) ([]*User, int64, error)
    
    // Update operations
    Update(ctx context.Context, user *User) error
    UpdateStatus(ctx context.Context, id string, status AccountStatus) error
    UpdateLastSeen(ctx context.Context, id string) error
    UpdateOnlineStatus(ctx context.Context, id string, isOnline bool) error
    IncrementLoginCount(ctx context.Context, id string) error
    IncrementFailedLoginAttempts(ctx context.Context, id string) error
    ResetFailedLoginAttempts(ctx context.Context, id string) error
    
    // Delete operations
    Delete(ctx context.Context, id string) error
    SoftDelete(ctx context.Context, id string, deletedBy string) error
    HardDelete(ctx context.Context, id string) error
    RestoreDeleted(ctx context.Context, id string) error
    
    // Business queries
    FindTopRatedFreelancers(ctx context.Context, limit int) ([]*User, error)
    FindFeaturedUsers(ctx context.Context, userType UserType, limit int) ([]*User, error)
    FindOnlineUsers(ctx context.Context, userType UserType) ([]*User, error)
    FindInactiveUsers(ctx context.Context, days int) ([]*User, error)
    FindUsersWithWarnings(ctx context.Context) ([]*User, error)
    FindUsersByCountry(ctx context.Context, country string, filter ListFilter) ([]*User, int64, error)
    
    // Analytics queries
    CountByUserType(ctx context.Context, userType UserType) (int64, error)
    CountByStatus(ctx context.Context, status AccountStatus) (int64, error)
    CountVerifiedUsers(ctx context.Context) (int64, error)
    GetAverageReputationScore(ctx context.Context, userType UserType) (float64, error)
    
    // Existence checks
    ExistsWithEmail(ctx context.Context, email string) (bool, error)
    ExistsWithUsername(ctx context.Context, username string) (bool, error)
    ExistsWithKeycloakID(ctx context.Context, keycloakID string) (bool, error)
}

// ListFilter provides filtering and pagination for list queries
type ListFilter struct {
    // Pagination
    Page     int
    PageSize int
    
    // Sorting
    SortBy    string // field name
    SortOrder string // asc, desc
    
    // Filters
    UserType       *UserType
    Status         *AccountStatus
    Country        *string
    IsVerified     *bool
    IsTopRated     *bool
    IsFeatured     *bool
    MinReputation  *float64
    
    // Date filters
    CreatedAfter   *time.Time
    CreatedBefore  *time.Time
    LastSeenAfter  *time.Time
    
    // Include deleted
    IncludeDeleted bool
}