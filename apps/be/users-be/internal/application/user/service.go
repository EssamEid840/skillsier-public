// internal/application/user/service.go
package user

import (
    "context"
    "fmt"
    "time"
    
    "gorm.io/gorm"
    "users-be/internal/domain/outbox"
    "users-be/internal/domain/user"
    "users-be/internal/infrastructure/cache/redis"
)

type Service struct {
    repo       user.Repository
    outboxRepo outbox.Repository
    db         *gorm.DB
    cache      *redis.UserCache
}

func NewService(
    repo user.Repository,
    outboxRepo outbox.Repository,
    db *gorm.DB,
    cache *redis.UserCache,
) *Service {
    return &Service{
        repo:       repo,
        outboxRepo: outboxRepo,
        db:         db,
        cache:      cache,
    }
}

// ============================================================================
// CREATE OPERATIONS
// ============================================================================

func (s *Service) CreateUser(ctx context.Context, dto CreateUserDTO) (*UserDTO, error) {
    // Validate email uniqueness
    exists, err := s.repo.ExistsWithEmail(ctx, dto.Email)
    if err != nil {
        return nil, fmt.Errorf("failed to check email existence: %w", err)
    }
    if exists {
        return nil, user.ErrEmailTaken
    }
    
    // Validate username uniqueness
    exists, err = s.repo.ExistsWithUsername(ctx, dto.Username)
    if err != nil {
        return nil, fmt.Errorf("failed to check username existence: %w", err)
    }
    if exists {
        return nil, user.ErrUsernameTaken
    }
    
    // Validate email format
    email, err := user.NewEmail(dto.Email)
    if err != nil {
        return nil, err
    }
    
    // Generate referral code
    referralCode := user.GenerateReferralCode(dto.Username)
    
    // Create user entity
    u := &user.User{
        KeycloakID:          dto.KeycloakID,
        Username:            dto.Username,
        Email:               email.Value,
        FirstName:           dto.FirstName,
        LastName:            dto.LastName,
        DisplayName:         fmt.Sprintf("%s %s", dto.FirstName, dto.LastName),
        UserType:            user.UserType(dto.UserType),
        Status:              user.AccountStatusPending,
        EmailVerified:       false,
        ProfileCompleteness: 20, // Base score for account creation
        Language:            dto.Language,
        Country:             dto.Country,
        Timezone:            dto.Timezone,
        Currency:            dto.Currency,
        ReferralCode:        referralCode.Code,
    }
    
    // Handle referral if provided
    if dto.ReferredBy != "" {
        referrer, err := s.repo.FindByReferralCode(ctx, dto.ReferredBy)
        if err == nil {
            u.ReferredBy = &referrer.ID
        }
    }
    
    // Transaction: Create user + publish event
    err = s.db.Transaction(func(tx *gorm.DB) error {
        if err := s.repo.Create(ctx, u); err != nil {
            return fmt.Errorf("failed to create user: %w", err)
        }
        
        // Publish user.created event
        event := &outbox.Event{
            AggregateID:   u.ID,
            AggregateType: "user",
            EventType:     "user.created",
            Payload:       s.buildUserCreatedEventPayload(u),
        }
        
        if err := s.outboxRepo.Create(ctx, event); err != nil {
            return fmt.Errorf("failed to create outbox event: %w", err)
        }
        
        // Update referrer stats if applicable
        if u.ReferredBy != nil {
            if err := s.incrementReferralCount(ctx, *u.ReferredBy); err != nil {
                // Log but don't fail transaction
                fmt.Printf("Warning: failed to increment referral count: %v\n", err)
            }
        }
        
        return nil
    })
    
    if err != nil {
        return nil, err
    }
    
    // Cache the user
    if err := s.cache.Set(ctx, u); err != nil {
        // Log but don't fail
        fmt.Printf("Warning: failed to cache user: %v\n", err)
    }
    
    return ToUserDTO(u), nil
}

func (s *Service) CreateBulkUsers(ctx context.Context, dtos []CreateUserDTO) ([]*UserDTO, error) {
    users := make([]*user.User, 0, len(dtos))
    
    for _, dto := range dtos {
        email, err := user.NewEmail(dto.Email)
        if err != nil {
            return nil, fmt.Errorf("invalid email for user %s: %w", dto.Username, err)
        }
        
        referralCode := user.GenerateReferralCode(dto.Username)
        
        u := &user.User{
            KeycloakID:          dto.KeycloakID,
            Username:            dto.Username,
            Email:               email.Value,
            FirstName:           dto.FirstName,
            LastName:            dto.LastName,
            DisplayName:         fmt.Sprintf("%s %s", dto.FirstName, dto.LastName),
            UserType:            user.UserType(dto.UserType),
            Status:              user.AccountStatusPending,
            ProfileCompleteness: 20,
            Language:            dto.Language,
            Country:             dto.Country,
            Timezone:            dto.Timezone,
            Currency:            dto.Currency,
            ReferralCode:        referralCode.Code,
        }
        
        users = append(users, u)
    }
    
    if err := s.repo.CreateBatch(ctx, users); err != nil {
        return nil, fmt.Errorf("failed to create users in batch: %w", err)
    }
    
    // Convert to DTOs
    dtoList := make([]*UserDTO, len(users))
    for i, u := range users {
        dtoList[i] = ToUserDTO(u)
    }
    
    return dtoList, nil
}

// ============================================================================
// READ OPERATIONS
// ============================================================================

func (s *Service) GetUserByID(ctx context.Context, id string) (*UserDTO, error) {
    // Try cache first
    u, err := s.cache.Get(ctx, id)
    if err == nil && u != nil {
        return ToUserDTO(u), nil
    }
    
    // Fetch from database
    u, err = s.repo.FindByID(ctx, id)
    if err != nil {
        return nil, err
    }
    
    // Cache for next time
    _ = s.cache.Set(ctx, u)
    
    return ToUserDTO(u), nil
}

func (s *Service) GetUserByEmail(ctx context.Context, email string) (*UserDTO, error) {
    u, err := s.repo.FindByEmail(ctx, email)
    if err != nil {
        return nil, err
    }
    return ToUserDTO(u), nil
}

func (s *Service) GetUserByUsername(ctx context.Context, username string) (*UserDTO, error) {
    u, err := s.repo.FindByUsername(ctx, username)
    if err != nil {
        return nil, err
    }
    return ToUserDTO(u), nil
}

func (s *Service) GetUserByKeycloakID(ctx context.Context, keycloakID string) (*UserDTO, error) {
    u, err := s.repo.FindByKeycloakID(ctx, keycloakID)
    if err != nil {
        return nil, err
    }
    return ToUserDTO(u), nil
}

func (s *Service) GetUsersByIDs(ctx context.Context, ids []string) ([]*UserDTO, error) {
    users, err := s.repo.FindByIDs(ctx, ids)
    if err != nil {
        return nil, err
    }
    
    dtos := make([]*UserDTO, len(users))
    for i, u := range users {
        dtos[i] = ToUserDTO(u)
    }
    
    return dtos, nil
}

func (s *Service) ListUsers(ctx context.Context, filter user.ListFilter) (*UserListResponseDTO, error) {
    users, total, err := s.repo.List(ctx, filter)
    if err != nil {
        return nil, err
    }
    
    dtos := make([]*UserDTO, len(users))
    for i, u := range users {
        dtos[i] = ToUserDTO(u)
    }
    
    return &UserListResponseDTO{
        Users:      dtos,
        Total:      total,
        Page:       filter.Page,
        PageSize:   filter.PageSize,
        TotalPages: calculateTotalPages(total, filter.PageSize),
    }, nil
}

func (s *Service) SearchUsers(ctx context.Context, query string, filter user.ListFilter) (*UserListResponseDTO, error) {
    users, total, err := s.repo.Search(ctx, query, filter)
    if err != nil {
        return nil, err
    }
    
    dtos := make([]*UserDTO, len(users))
    for i, u := range users {
        dtos[i] = ToUserDTO(u)
    }
    
    return &UserListResponseDTO{
        Users:      dtos,
        Total:      total,
        Page:       filter.Page,
        PageSize:   filter.PageSize,
        TotalPages: calculateTotalPages(total, filter.PageSize),
    }, nil
}

// ============================================================================
// UPDATE OPERATIONS
// ============================================================================

func (s *Service) UpdateUser(ctx context.Context, id string, dto UpdateUserDTO) (*UserDTO, error) {
    u, err := s.repo.FindByID(ctx, id)
    if err != nil {
        return nil, err
    }
    
    // Check if user can be updated
    if u.Status == user.AccountStatusBanned {
        return nil, user.ErrUserBanned
    }
    
    // Apply updates
    updated := false
    
    if dto.FirstName != nil && *dto.FirstName != u.FirstName {
        u.FirstName = *dto.FirstName
        updated = true
    }
    
    if dto.LastName != nil && *dto.LastName != u.LastName {
        u.LastName = *dto.LastName
        updated = true
    }
    
    if dto.DisplayName != nil && *dto.DisplayName != u.DisplayName {
        u.DisplayName = *dto.DisplayName
        updated = true
    }
    
    if dto.PhoneNumber != nil && *dto.PhoneNumber != u.PhoneNumber {
        // Validate phone
        if dto.PhoneCountryCode != nil {
            _, err := user.NewPhone(*dto.PhoneCountryCode, *dto.PhoneNumber)
            if err != nil {
                return nil, err
            }
            u.PhoneCountryCode = *dto.PhoneCountryCode
        }
        u.PhoneNumber = *dto.PhoneNumber
        u.PhoneVerified = false
        u.PhoneVerifiedAt = nil
        updated = true
    }
    
    if dto.Country != nil && *dto.Country != u.Country {
        u.Country = *dto.Country
        updated = true
    }
    
    if dto.Timezone != nil && *dto.Timezone != u.Timezone {
        u.Timezone = *dto.Timezone
        updated = true
    }
    
    if dto.Language != nil && *dto.Language != u.Language {
        u.Language = *dto.Language
        updated = true
    }
    
    if dto.Currency != nil && *dto.Currency != u.Currency {
        u.Currency = *dto.Currency
        updated = true
    }
    
    if !updated {
        return ToUserDTO(u), nil
    }
    
    // Transaction: Update user + publish event
    err = s.db.Transaction(func(tx *gorm.DB) error {
        if err := s.repo.Update(ctx, u); err != nil {
            return fmt.Errorf("failed to update user: %w", err)
        }
        
        // Publish user.updated event
        event := &outbox.Event{
            AggregateID:   u.ID,
            AggregateType: "user",
            EventType:     "user.updated",
            Payload:       s.buildUserUpdatedEventPayload(u),
        }
        
        if err := s.outboxRepo.Create(ctx, event); err != nil {
            return fmt.Errorf("failed to create outbox event: %w", err)
        }
        
        return nil
    })
    
    if err != nil {
        return nil, err
    }
    
    // Invalidate cache
    _ = s.cache.Invalidate(ctx, id)
    
    return ToUserDTO(u), nil
}

func (s *Service) VerifyEmail(ctx context.Context, userID string) (*UserDTO, error) {
    u, err := s.repo.FindByID(ctx, userID)
    if err != nil {
        return nil, err
    }
    
    if u.EmailVerified {
        return ToUserDTO(u), nil
    }
    
    now := time.Now()
    u.EmailVerified = true
    u.EmailVerifiedAt = &now
    
    // Update status to active if pending
    if u.Status == user.AccountStatusPending {
        u.Status = user.AccountStatusActive
    }
    
    // Recalculate profile completeness
    u.ProfileCompleteness += 10
    
    err = s.db.Transaction(func(tx *gorm.DB) error {
        if err := s.repo.Update(ctx, u); err != nil {
            return err
        }
        
        // Publish user.verified event
        event := &outbox.Event{
            AggregateID:   u.ID,
            AggregateType: "user",
            EventType:     "user.verified",
            Payload:       s.buildUserVerifiedEventPayload(u),
        }
        
        return s.outboxRepo.Create(ctx, event)
    })
    
    if err != nil {
        return nil, err
    }
    
    _ = s.cache.Invalidate(ctx, userID)
    
    return ToUserDTO(u), nil
}

func (s *Service) VerifyPhone(ctx context.Context, userID string) (*UserDTO, error) {
    u, err := s.repo.FindByID(ctx, userID)
    if err != nil {
        return nil, err
    }
    
    if u.PhoneVerified {
        return ToUserDTO(u), nil
    }
    
    now := time.Now()
    u.PhoneVerified = true
    u.PhoneVerifiedAt = &now
    u.ProfileCompleteness += 5
    
    if err := s.repo.Update(ctx, u); err != nil {
        return nil, err
    }
    
    _ = s.cache.Invalidate(ctx, userID)
    
    return ToUserDTO(u), nil
}

func (s *Service) VerifyIdentity(ctx context.Context, userID string) (*UserDTO, error) {
    u, err := s.repo.FindByID(ctx, userID)
    if err != nil {
        return nil, err
    }
    
    now := time.Now()
    u.IdentityVerified = true
    u.IdentityVerifiedAt = &now
    u.ProfileCompleteness += 10
    
    if err := s.repo.Update(ctx, u); err != nil {
        return nil, err
    }
    
    _ = s.cache.Invalidate(ctx, userID)
    
    return ToUserDTO(u), nil
}

func (s *Service) VerifyPayment(ctx context.Context, userID string) (*UserDTO, error) {
    u, err := s.repo.FindByID(ctx, userID)
    if err != nil {
        return nil, err
    }
    
    now := time.Now()
    u.PaymentVerified = true
    u.PaymentVerifiedAt = &now
    u.ProfileCompleteness += 10
    
    if err := s.repo.Update(ctx, u); err != nil {
        return nil, err
    }
    
    _ = s.cache.Invalidate(ctx, userID)
    
    return ToUserDTO(u), nil
}

func (s *Service) UpdateOnlineStatus(ctx context.Context, userID string, isOnline bool) error {
    if err := s.repo.UpdateOnlineStatus(ctx, userID, isOnline); err != nil {
        return err
    }
    
    _ = s.cache.Invalidate(ctx, userID)
    return nil
}

func (s *Service) RecordLogin(ctx context.Context, userID, ip string) error {
    u, err := s.repo.FindByID(ctx, userID)
    if err != nil {
        return err
    }
    
    u.UpdateLoginInfo(ip)
    
    if err := s.repo.Update(ctx, u); err != nil {
        return err
    }
    
    _ = s.cache.Invalidate(ctx, userID)
    return nil
}

func (s *Service) RecordFailedLogin(ctx context.Context, userID string) error {
    u, err := s.repo.FindByID(ctx, userID)
    if err != nil {
        return err
    }
    
    u.RecordFailedLogin()
    
    if err := s.repo.Update(ctx, u); err != nil {
        return err
    }
    
    if u.ShouldLockAccount() {
        if err := s.SuspendUser(ctx, userID, "Too many failed login attempts", "system"); err != nil {
            return err
        }
    }
    
    _ = s.cache.Invalidate(ctx, userID)
    return nil
}

// ============================================================================
// DELETE OPERATIONS
// ============================================================================

func (s *Service) DeleteUser(ctx context.Context, userID string) error {
    u, err := s.repo.FindByID(ctx, userID)
    if err != nil {
        return err
    }
    
    // Soft delete
    if err := s.repo.SoftDelete(ctx, userID, userID); err != nil {
        return err
    }
    
    // Publish event
    event := &outbox.Event{
        AggregateID:   u.ID,
        AggregateType: "user",
        EventType:     "user.deleted",
        Payload:       fmt.Sprintf(`{"user_id":"%s"}`, u.ID),
    }
    
    _ = s.outboxRepo.Create(ctx, event)
    
    _ = s.cache.Invalidate(ctx, userID)
    
    return nil
}

func (s *Service) HardDeleteUser(ctx context.Context, userID string) error {
    if err := s.repo.HardDelete(ctx, userID); err != nil {
        return err
    }
    
    _ = s.cache.Invalidate(ctx, userID)
    return nil
}

func (s *Service) RestoreUser(ctx context.Context, userID string) (*UserDTO, error) {
    if err := s.repo.RestoreDeleted(ctx, userID); err != nil {
        return nil, err
    }
    
    u, err := s.repo.FindByID(ctx, userID)
    if err != nil {
        return nil, err
    }
    
    _ = s.cache.Invalidate(ctx, userID)
    
    return ToUserDTO(u), nil
}

// ============================================================================
// ADMIN OPERATIONS
// ============================================================================

func (s *Service) SuspendUser(ctx context.Context, userID, reason, suspendedBy string) error {
    u, err := s.repo.FindByID(ctx, userID)
    if err != nil {
        return err
    }
    
    u.Status = user.AccountStatusSuspended
    u.SuspensionCount++
    
    err = s.db.Transaction(func(tx *gorm.DB) error {
        if err := s.repo.Update(ctx, u); err != nil {
            return err
        }
        
        event := &outbox.Event{
            AggregateID:   u.ID,
            AggregateType: "user",
            EventType:     "user.suspended",
            Payload:       fmt.Sprintf(`{"user_id":"%s","reason":"%s","suspended_by":"%s"}`, u.ID, reason, suspendedBy),
        }
        
        return s.outboxRepo.Create(ctx, event)
    })
    
    if err != nil {
        return err
    }
    
    _ = s.cache.Invalidate(ctx, userID)
    return nil
}

func (s *Service) UnsuspendUser(ctx context.Context, userID string) error {
    if err := s.repo.UpdateStatus(ctx, userID, user.AccountStatusActive); err != nil {
        return err
    }
    
    _ = s.cache.Invalidate(ctx, userID)
    return nil
}

func (s *Service) BanUser(ctx context.Context, userID, reason, bannedBy string) error {
    u, err := s.repo.FindByID(ctx, userID)
    if err != nil {
        return err
    }
    
    u.Status = user.AccountStatusBanned
    
    err = s.db.Transaction(func(tx *gorm.DB) error {
        if err := s.repo.Update(ctx, u); err != nil {
            return err
        }
        
        event := &outbox.Event{
            AggregateID:   u.ID,
            AggregateType: "user",
            EventType:     "user.banned",
            Payload:       fmt.Sprintf(`{"user_id":"%s","reason":"%s","banned_by":"%s"}`, u.ID, reason, bannedBy),
        }
        
        return s.outboxRepo.Create(ctx, event)
    })
    
    if err != nil {
        return err
    }
    
    _ = s.cache.Invalidate(ctx, userID)
    return