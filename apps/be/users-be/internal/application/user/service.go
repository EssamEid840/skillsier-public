// apps/be/users-be/internal/application/user/service.go
package user

import (
	"context"
	"fmt"
	"time"
	
	"users-be/internal/domain/user"
	
	"gorm.io/gorm"
)

// Service handles user business logic
type Service struct {
	repo user.Repository
	db   *gorm.DB
}

// NewService creates a new user service
func NewService(repo user.Repository, db *gorm.DB) *Service {
	return &Service{
		repo: repo,
		db:   db,
	}
}

// ============================================================================
// CREATE OPERATIONS
// ============================================================================

// CreateUser creates a new user
func (s *Service) CreateUser(ctx context.Context, dto *CreateUserDTO) (*UserDTO, error) {
	// Check if user already exists
	exists, err := s.repo.ExistsByKeycloakID(ctx, dto.KeycloakID)
	if err != nil {
		return nil, fmt.Errorf("failed to check user existence: %w", err)
	}
	if exists {
		return nil, user.ErrUserAlreadyExists
	}
	
	// Check if username is taken
	exists, err = s.repo.ExistsByUsername(ctx, dto.Username)
	if err != nil {
		return nil, fmt.Errorf("failed to check username: %w", err)
	}
	if exists {
		return nil, user.ErrUsernameTaken
	}
	
	// Check if email is taken
	exists, err = s.repo.ExistsByEmail(ctx, dto.Email)
	if err != nil {
		return nil, fmt.Errorf("failed to check email: %w", err)
	}
	if exists {
		return nil, user.ErrEmailTaken
	}
	
	// Convert DTO to entity
	newUser, err := dto.ToEntity()
	if err != nil {
		return nil, fmt.Errorf("failed to convert DTO to entity: %w", err)
	}
	
	// Create user in transaction
	err = s.db.Transaction(func(tx *gorm.DB) error {
		if err := s.repo.Create(ctx, newUser); err != nil {
			return fmt.Errorf("failed to create user: %w", err)
		}
		
		// TODO: Publish user.created event to outbox
		
		return nil
	})
	
	if err != nil {
		return nil, err
	}
	
	return ToDTO(newUser), nil
}

// ============================================================================
// READ OPERATIONS
// ============================================================================

// GetUser retrieves a user by ID
func (s *Service) GetUser(ctx context.Context, id string) (*UserDTO, error) {
	u, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return ToDTO(u), nil
}

// GetUserByKeycloakID retrieves a user by Keycloak ID
func (s *Service) GetUserByKeycloakID(ctx context.Context, keycloakID string) (*UserDTO, error) {
	u, err := s.repo.FindByKeycloakID(ctx, keycloakID)
	if err != nil {
		return nil, err
	}
	return ToDTO(u), nil
}

// GetUserByEmail retrieves a user by email
func (s *Service) GetUserByEmail(ctx context.Context, email string) (*UserDTO, error) {
	u, err := s.repo.FindByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	return ToDTO(u), nil
}

// GetUserByUsername retrieves a user by username
func (s *Service) GetUserByUsername(ctx context.Context, username string) (*UserDTO, error) {
	u, err := s.repo.FindByUsername(ctx, username)
	if err != nil {
		return nil, err
	}
	return ToDTO(u), nil
}

// GetUsersByIDs retrieves multiple users by their IDs
func (s *Service) GetUsersByIDs(ctx context.Context, ids []string) ([]UserDTO, error) {
	users, err := s.repo.FindByIDs(ctx, ids)
	if err != nil {
		return nil, err
	}
	return ToDTOList(users), nil
}

// ListUsers retrieves a paginated list of users
func (s *Service) ListUsers(ctx context.Context, filterDTO UserFilterDTO) (*UserListDTO, error) {
	filter := filterDTO.ToListFilter()
	users, total, err := s.repo.List(ctx, filter)
	if err != nil {
		return nil, err
	}
	
	return ToListDTO(users, total, filter.Page, filter.PageSize), nil
}

// SearchUsers performs full-text search on users
func (s *Service) SearchUsers(ctx context.Context, query string, filterDTO UserFilterDTO) (*UserSearchResultDTO, error) {
	startTime := time.Now()
	filter := filterDTO.ToListFilter()
	
	users, total, err := s.repo.Search(ctx, query, filter)
	if err != nil {
		return nil, err
	}
	
	searchTime := time.Since(startTime).Milliseconds()
	return ToSearchResultDTO(users, total, query, searchTime), nil
}

// GetUserStatistics retrieves comprehensive user statistics
func (s *Service) GetUserStatistics(ctx context.Context) (*UserStatisticsDTO, error) {
	stats, err := s.repo.GetStatistics(ctx)
	if err != nil {
		return nil, err
	}
	return ToStatisticsDTO(stats), nil
}

// ============================================================================
// UPDATE OPERATIONS
// ============================================================================

// UpdateUser updates user information
func (s *Service) UpdateUser(ctx context.Context, id string, dto *UpdateUserDTO) (*UserDTO, error) {
	u, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	
	// Apply updates
	if err := dto.ApplyUpdates(u); err != nil {
		return nil, fmt.Errorf("failed to apply updates: %w", err)
	}
	
	// Save in transaction
	err = s.db.Transaction(func(tx *gorm.DB) error {
		if err := s.repo.Update(ctx, u); err != nil {
			return fmt.Errorf("failed to update user: %w", err)
		}
		
		// TODO: Publish user.updated event to outbox
		
		return nil
	})
	
	if err != nil {
		return nil, err
	}
	
	return ToDTO(u), nil
}

// UpdateAvailability updates user availability status
func (s *Service) UpdateAvailability(ctx context.Context, id string, dto *UpdateAvailabilityDTO) (*UserDTO, error) {
	u, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	
	dto.ApplyUpdates(u)
	
	if err := s.repo.Update(ctx, u); err != nil {
		return nil, err
	}
	
	return ToDTO(u), nil
}

// UpdateSettings updates user privacy and settings
func (s *Service) UpdateSettings(ctx context.Context, id string, dto *UpdateSettingsDTO) (*UserDTO, error) {
	u, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	
	dto.ApplyUpdates(u)
	
	if err := s.repo.Update(ctx, u); err != nil {
		return nil, err
	}
	
	return ToDTO(u), nil
}

// ============================================================================
// VERIFICATION OPERATIONS
// ============================================================================

// VerifyEmail marks user's email as verified
func (s *Service) VerifyEmail(ctx context.Context, id string) (*UserDTO, error) {
	if err := s.repo.VerifyEmail(ctx, id); err != nil {
		return nil, err
	}
	
	// Update verification status if all verifications complete
	u, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	
	if u.Email.Verified && (u.Phone == nil || u.Phone.Verified) {
		if err := s.repo.UpdateVerificationStatus(ctx, id, user.VerificationStatusVerified); err != nil {
			return nil, err
		}
		u.VerificationStatus = user.VerificationStatusVerified
	}
	
	// TODO: Publish user.verified event
	
	return ToDTO(u), nil
}

// VerifyPhone marks user's phone as verified
func (s *Service) VerifyPhone(ctx context.Context, id string) (*UserDTO, error) {
	if err := s.repo.VerifyPhone(ctx, id); err != nil {
		return nil, err
	}
	
	u, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	
	// TODO: Publish user.verified event
	
	return ToDTO(u), nil
}

// VerifyIdentity marks user's identity as verified
func (s *Service) VerifyIdentity(ctx context.Context, id string) (*UserDTO, error) {
	if err := s.repo.VerifyIdentity(ctx, id); err != nil {
		return nil, err
	}
	
	u, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	
	// TODO: Publish user.verified event
	
	return ToDTO(u), nil
}

// ============================================================================
// MODERATION OPERATIONS
// ============================================================================

// SuspendUser suspends a user account
func (s *Service) SuspendUser(ctx context.Context, id, reason, suspendedBy string) (*UserDTO, error) {
	u, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	
	if u.Status == user.StatusSuspended {
		return nil, user.ErrUserAlreadySuspended
	}
	
	err = s.db.Transaction(func(tx *gorm.DB) error {
		if err := s.repo.UpdateStatus(ctx, id, user.StatusSuspended); err != nil {
			return err
		}
		
		// TODO: Create suspension record
		// TODO: Publish user.suspended event
		
		return nil
	})
	
	if err != nil {
		return nil, err
	}
	
	u.Status = user.StatusSuspended
	return ToDTO(u), nil
}

// UnsuspendUser reactivates a suspended user
func (s *Service) UnsuspendUser(ctx context.Context, id, unsuspendedBy string) (*UserDTO, error) {
	u, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	
	if u.Status != user.StatusSuspended {
		return nil, user.ErrUserNotSuspended
	}
	
	err = s.db.Transaction(func(tx *gorm.DB) error {
		if err := s.repo.UpdateStatus(ctx, id, user.StatusActive); err != nil {
			return err
		}
		
		// TODO: Update suspension record
		// TODO: Publish user.unsuspended event
		
		return nil
	})
	
	if err != nil {
		return nil, err
	}
	
	u.Status = user.StatusActive
	return ToDTO(u), nil
}

// BanUser permanently bans a user
func (s *Service) BanUser(ctx context.Context, id, reason, bannedBy string) (*UserDTO, error) {
	u, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	
	if u.Status == user.StatusBanned {
		return nil, user.ErrUserAlreadyBanned
	}
	
	err = s.db.Transaction(func(tx *gorm.DB) error {
		if err := s.repo.UpdateStatus(ctx, id, user.StatusBanned); err != nil {
			return err
		}
		
		// TODO: Create ban record
		// TODO: Publish user.banned event
		
		return nil
	})
	
	if err != nil {
		return nil, err
	}
	
	u.Status = user.StatusBanned
	return ToDTO(u), nil
}

// RestoreUser restores a banned user
func (s *Service) RestoreUser(ctx context.Context, id, restoredBy string) (*UserDTO, error) {
	u, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	
	err = s.db.Transaction(func(tx *gorm.DB) error {
		if err := s.repo.UpdateStatus(ctx, id, user.StatusActive); err != nil {
			return err
		}
		
		// Clear warnings and flags
		if err := s.repo.ClearWarnings(ctx, id); err != nil {
			return err
		}
		if err := s.repo.ClearFlags(ctx, id); err != nil {
			return err
		}
		
		// TODO: Publish user.restored event
		
		return nil
	})
	
	if err != nil {
		return nil, err
	}
	
	u.Status = user.StatusActive
	u.Warnings = 0
	return ToDTO(u), nil
}

// AddWarning adds a warning to user account
func (s *Service) AddWarning(ctx context.Context, id, reason, issuedBy string) (*UserDTO, error) {
	err := s.db.Transaction(func(tx *gorm.DB) error {
		if err := s.repo.IncrementWarnings(ctx, id); err != nil {
			return err
		}
		
		// TODO: Create warning record
		// TODO: Check if warnings exceed threshold for auto-suspension
		
		return nil
	})
	
	if err != nil {
		return nil, err
	}
	
	u, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	
	return ToDTO(u), nil
}

// ============================================================================
// BADGE OPERATIONS
// ============================================================================

// AssignBadge assigns a badge to a user
func (s *Service) AssignBadge(ctx context.Context, id string, badgeType user.BadgeType) (*UserDTO, error) {
	if err := s.repo.AddBadge(ctx, id, string(badgeType)); err != nil {
		return nil, err
	}
	
	// Update related flags based on badge type
	switch badgeType {
	case user.BadgeTypeTopRated:
		// Update is_top_rated flag
		if err := s.repo.Update(ctx, &user.User{ID: id, IsTopRated: true}); err != nil {
			return nil, err
		}
	case user.BadgeTypeRisingTalent:
		if err := s.repo.Update(ctx, &user.User{ID: id, IsRisingTalent: true}); err != nil {
			return nil, err
		}
	case user.BadgeTypeExpertVetted:
		if err := s.repo.Update(ctx, &user.User{ID: id, IsExpertVetted: true}); err != nil {
			return nil, err
		}
	}
	
	// TODO: Publish badge.awarded event
	
	u, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	
	return ToDTO(u), nil
}

// RemoveBadge removes a badge from a user
func (s *Service) RemoveBadge(ctx context.Context, id string, badgeType user.BadgeType) (*UserDTO, error) {
	if err := s.repo.RemoveBadge(ctx, id, string(badgeType)); err != nil {
		return nil, err
	}
	
	// Update related flags
	switch badgeType {
	case user.BadgeTypeTopRated:
		if err := s.repo.Update(ctx, &user.User{ID: id, IsTopRated: false}); err != nil {
			return nil, err
		}
	case user.BadgeTypeRisingTalent:
		if err := s.repo.Update(ctx, &user.User{ID: id, IsRisingTalent: false}); err != nil {
			return nil, err
		}
	case user.BadgeTypeExpertVetted:
		if err := s.repo.Update(ctx, &user.User{ID: id, IsExpertVetted: false}); err != nil {
			return nil, err
		}
	}
	
	u, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	
	return ToDTO(u), nil
}

// ============================================================================
// ACTIVITY TRACKING
// ============================================================================

// RecordLogin records a user login
func (s *Service) RecordLogin(ctx context.Context, id, ipAddress, userAgent string) error {
	return s.repo.RecordLogin(ctx, id, ipAddress, userAgent)
}

// UpdateLastSeen updates user's last seen timestamp
func (s *Service) UpdateLastSeen(ctx context.Context, id string) error {
	return s.repo.UpdateLastSeen(ctx, id)
}

// SetOnlineStatus sets user's online status
func (s *Service) SetOnlineStatus(ctx context.Context, id string, isOnline bool) error {
	return s.repo.UpdateOnlineStatus(ctx, id, isOnline)
}

// ============================================================================
// FEATURED & SPECIAL OPERATIONS
// ============================================================================

// SetFeatured marks a user as featured
func (s *Service) SetFeatured(ctx context.Context, id string, featured bool) (*UserDTO, error) {
	u, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	
	u.IsFeatured = featured
	if err := s.repo.Update(ctx, u); err != nil {
		return nil, err
	}
	
	return ToDTO(u), nil
}

// ============================================================================
// REFERRAL OPERATIONS
// ============================================================================

// GetReferrals gets all users referred by this user
func (s *Service) GetReferrals(ctx context.Context, referrerID string) ([]UserDTO, error) {
	users, err := s.repo.FindByReferrer(ctx, referrerID)
	if err != nil {
		return nil, err
	}
	return ToDTOList(users), nil
}

// GetReferralCount gets count of users referred
func (s *Service) GetReferralCount(ctx context.Context, referrerID string) (int64, error) {
	return s.repo.CountReferrals(ctx, referrerID)
}

// ============================================================================
// CONNECTS OPERATIONS
// ============================================================================

// AddConnects adds connects to user balance
func (s *Service) AddConnects(ctx context.Context, id string, amount int) error {
	if amount <= 0 {
		return fmt.Errorf("amount must be positive")
	}
	return s.repo.AddConnects(ctx, id, amount)
}

// DeductConnects deducts connects from user balance
func (s *Service) DeductConnects(ctx context.Context, id string, amount int) error {
	if amount <= 0 {
		return fmt.Errorf("amount must be positive")
	}
	
	// Check balance
	balance, err := s.repo.GetConnectsBalance(ctx, id)
	if err != nil {
		return err
	}
	
	if balance < amount {
		return user.ErrInsufficientConnects
	}
	
	return s.repo.DeductConnects(ctx, id, amount)
}

// GetConnectsBalance gets user's connects balance
func (s *Service) GetConnectsBalance(ctx context.Context, id string) (int, error) {
	return s.repo.GetConnectsBalance(ctx, id)
}

// ============================================================================
// BULK OPERATIONS
// ============================================================================

// BulkUpdateStatus updates status for multiple users
func (s *Service) BulkUpdateStatus(ctx context.Context, ids []string, status user.AccountStatus) (*BulkActionResultDTO, error) {
	result := &BulkActionResultDTO{
		Success: []string{},
		Failed:  []string{},
		Errors:  []string{},
	}
	
	for _, id := range ids {
		if err := s.repo.UpdateStatus(ctx, id, status); err != nil {
			result.Failed = append(result.Failed, id)
			result.Errors = append(result.Errors, fmt.Sprintf("%s: %v", id, err))
		} else {
			result.Success = append(result.Success, id)
		}
	}
	
	result.TotalSuccess = len(result.Success)
	result.TotalFailed = len(result.Failed)
	
	return result, nil
}

// ============================================================================
// DELETE OPERATIONS
// ============================================================================

// DeleteUser soft-deletes a user
func (s *Service) DeleteUser(ctx context.Context, id, deletedBy string) error {
	return s.db.Transaction(func(tx *gorm.DB) error {
		if err := s.repo.SoftDelete(ctx, id, deletedBy); err != nil {
			return err
		}
		
		// TODO: Publish user.deleted event
		
		return nil
	})
}

// RestoreDeletedUser restores a soft-deleted user
func (s *Service) RestoreDeletedUser(ctx context.Context, id string) (*UserDTO, error) {
	if err := s.repo.RestoreDeleted(ctx, id); err != nil {
		return nil, err
	}
	
	u, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	
	return ToDTO(u), nil
}

// ============================================================================
// BUSINESS QUERIES
// ============================================================================

// GetTopRatedFreelancers retrieves top-rated freelancers
func (s *Service) GetTopRatedFreelancers(ctx context.Context, limit int) ([]UserDTO, error) {
	users, err := s.repo.FindTopRatedFreelancers(ctx, limit)
	if err != nil {
		return nil, err
	}
	return ToDTOList(users), nil
}

// GetTopRatedClients retrieves top-rated clients
func (s *Service) GetTopRatedClients(ctx context.Context, limit int) ([]UserDTO, error) {
	users, err := s.repo.FindTopRatedClients(ctx, limit)
	if err != nil {
		return nil, err
	}
	return ToDTOList(users), nil
}

// GetFeaturedUsers retrieves featured users
func (s *Service) GetFeaturedUsers(ctx context.Context, userType user.UserType, limit int) ([]UserDTO, error) {
	users, err := s.repo.FindFeaturedUsers(ctx, userType, limit)
	if err != nil {
		return nil, err
	}
	return ToDTOList(users), nil
}

// GetRisingTalent retrieves rising talent freelancers
func (s *Service) GetRisingTalent(ctx context.Context, limit int) ([]UserDTO, error) {
	users, err := s.repo.FindRisingTalent(ctx, limit)
	if err != nil {
		return nil, err
	}
	return ToDTOList(users), nil
}

// GetOnlineUsers retrieves currently online users
func (s *Service) GetOnlineUsers(ctx context.Context, userType user.UserType) ([]UserDTO, error) {
	users, err := s.repo.FindOnlineUsers(ctx, userType)
	if err != nil {
		return nil, err
	}
	return ToDTOList(users), nil
}