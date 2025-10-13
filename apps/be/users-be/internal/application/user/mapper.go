// internal/application/user/mapper.go
package user

import (
    "time"
    "users-be/internal/domain/user"
)

// ============================================================================
// ENTITY TO DTO MAPPERS
// ============================================================================

func ToUserDTO(u *user.User) *UserDTO {
    if u == nil {
        return nil
    }
    
    return &UserDTO{
        ID:                  u.ID,
        KeycloakID:          u.KeycloakID,
        Username:            u.Username,
        Email:               u.Email,
        FirstName:           u.FirstName,
        LastName:            u.LastName,
        MiddleName:          u.MiddleName,
        DisplayName:         u.DisplayName,
        UserType:            string(u.UserType),
        Status:              string(u.Status),
        PhoneNumber:         u.PhoneNumber,
        PhoneCountryCode:    u.PhoneCountryCode,
        PhoneVerified:       u.PhoneVerified,
        EmailVerified:       u.EmailVerified,
        EmailVerifiedAt:     u.EmailVerifiedAt,
        IdentityVerified:    u.IdentityVerified,
        IdentityVerifiedAt:  u.IdentityVerifiedAt,
        PaymentVerified:     u.PaymentVerified,
        PaymentVerifiedAt:   u.PaymentVerifiedAt,
        TwoFactorEnabled:    u.TwoFactorEnabled,
        LastLoginAt:         u.LastLoginAt,
        ProfileCompleteness: u.ProfileCompleteness,
        OnboardingCompleted: u.OnboardingCompleted,
        Country:             u.Country,
        Timezone:            u.Timezone,
        Language:            u.Language,
        Currency:            u.Currency,
        ReputationScore:     u.ReputationScore,
        TotalRatings:        u.TotalRatings,
        AverageRating:       u.AverageRating,
        ResponseRatePercent: u.ResponseRatePercent,
        IsOnline:            u.IsOnline,
        LastSeenAt:          u.LastSeenAt,
        TotalEarnings:       u.TotalEarnings,
        TotalSpent:          u.TotalSpent,
        TotalJobs:           u.TotalJobs,
        TotalHires:          u.TotalHires,
        CompletedJobs:       u.CompletedJobs,
        IsTopRated:          u.IsTopRated,
        IsFeatured:          u.IsFeatured,
        IsPremium:           u.IsPremium,
        PremiumUntil:        u.PremiumUntil,
        HasActiveWarnings:   u.HasActiveWarnings,
        WarningCount:        u.WarningCount,
        SuspensionCount:     u.SuspensionCount,
        ReferralCode:        u.ReferralCode,
        TotalReferrals:      u.TotalReferrals,
        CreatedAt:           u.CreatedAt,
        UpdatedAt:           u.UpdatedAt,
        DeletedAt:           getDeletedAt(u.DeletedAt),
    }
}

func ToPublicUserDTO(u *user.User) *PublicUserDTO {
    if u == nil {
        return nil
    }
    
    return &PublicUserDTO{
        ID:              u.ID,
        Username:        u.Username,
        DisplayName:     u.DisplayName,
        UserType:        string(u.UserType),
        Country:         u.Country,
        ReputationScore: u.ReputationScore,
        TotalRatings:    u.TotalRatings,
        IsTopRated:      u.IsTopRated,
        IsFeatured:      u.IsFeatured,
        IsOnline:        u.IsOnline,
        LastSeenAt:      u.LastSeenAt,
        MemberSince:     u.CreatedAt,
    }
}

func ToUserProfileSummaryDTO(u *user.User) *UserProfileSummaryDTO {
    if u == nil {
        return nil
    }
    
    return &UserProfileSummaryDTO{
        ID:              u.ID,
        Username:        u.Username,
        DisplayName:     u.DisplayName,
        UserType:        string(u.UserType),
        ReputationScore: u.ReputationScore,
        TotalRatings:    u.TotalRatings,
        IsTopRated:      u.IsTopRated,
        IsOnline:        u.IsOnline,
    }
}

func ToUserTrustScoreDTO(u *user.User) *UserTrustScoreDTO {
    if u == nil {
        return nil
    }
    
    trustScore := calculateTrustScore(u)
    trustLevel := calculateTrustLevel(u)
    accountAgeDays := int(time.Since(u.CreatedAt).Hours() / 24)
    
    return &UserTrustScoreDTO{
        UserID:            u.ID,
        TrustLevel:        trustLevel,
        TrustScore:        trustScore,
        EmailVerified:     u.EmailVerified,
        PhoneVerified:     u.PhoneVerified,
        IdentityVerified:  u.IdentityVerified,
        PaymentVerified:   u.PaymentVerified,
        BackgroundChecked: u.BackgroundChecked,
        ReputationScore:   u.ReputationScore,
        TotalRatings:      u.TotalRatings,
        HasActiveWarnings: u.HasActiveWarnings,
        AccountAge:        accountAgeDays,
    }
}

func ToLoginHistoryDTO(u *user.User) *LoginHistoryDTO {
    if u == nil {
        return nil
    }
    
    return &LoginHistoryDTO{
        LastLoginAt:       u.LastLoginAt,
        LastLoginIP:       u.LastLoginIP,
        LoginCount:        u.LoginCount,
        FailedAttempts:    u.FailedLoginAttempts,
        LastFailedLoginAt: u.LastFailedLoginAt,
    }
}

func ToNotificationPreferencesDTO(u *user.User) *NotificationPreferencesDTO {
    if u == nil {
        return nil
    }
    
    return &NotificationPreferencesDTO{
        MarketingEmails:   u.MarketingEmailsEnabled,
        ProductUpdates:    u.ProductUpdatesEnabled,
        Newsletter:        u.NewsletterSubscribed,
        SmsNotifications:  u.SmsNotificationsEnabled,
        PushNotifications: u.PushNotificationsEnabled,
    }
}

// ============================================================================
// BATCH MAPPERS
// ============================================================================

func ToUserDTOs(users []*user.User) []*UserDTO {
    if users == nil {
        return nil
    }
    
    dtos := make([]*UserDTO, len(users))
    for i, u := range users {
        dtos[i] = ToUserDTO(u)
    }
    return dtos
}

func ToPublicUserDTOs(users []*user.User) []*PublicUserDTO {
    if users == nil {
        return nil
    }
    
    dtos := make([]*PublicUserDTO, len(users))
    for i, u := range users {
        dtos[i] = ToPublicUserDTO(u)
    }
    return dtos
}

func ToUserProfileSummaryDTOs(users []*user.User) []*UserProfileSummaryDTO {
    if users == nil {
        return nil
    }
    
    dtos := make([]*UserProfileSummaryDTO, len(users))
    for i, u := range users {
        dtos[i] = ToUserProfileSummaryDTO(u)
    }
    return dtos
}

// ============================================================================
// DTO TO FILTER MAPPERS
// ============================================================================

func ToListFilter(dto UserFilterDTO) user.ListFilter {
    filter := user.ListFilter{
        Page:           dto.Page,
        PageSize:       dto.PageSize,
        SortBy:         dto.SortBy,
        SortOrder:      dto.SortOrder,
        IncludeDeleted: dto.IncludeDeleted,
    }
    
    if dto.UserType != nil {
        userType := user.UserType(*dto.UserType)
        filter.UserType = &userType
    }
    
    if dto.Status != nil {
        status := user.AccountStatus(*dto.Status)
        filter.Status = &status
    }
    
    if dto.Country != nil {
        filter.Country = dto.Country
    }
    
    if dto.IsVerified != nil {
        filter.IsVerified = dto.IsVerified
    }
    
    if dto.IsTopRated != nil {
        filter.IsTopRated = dto.IsTopRated
    }
    
    if dto.IsFeatured != nil {
        filter.IsFeatured = dto.IsFeatured
    }
    
    if dto.MinReputation != nil {
        filter.MinReputation = dto.MinReputation
    }
    
    if dto.CreatedAfter != nil {
        filter.CreatedAfter = dto.CreatedAfter
    }
    
    if dto.CreatedBefore != nil {
        filter.CreatedBefore = dto.CreatedBefore
    }
    
    return filter
}

func ToSearchFilter(dto UserSearchQueryDTO) user.ListFilter {
    filter := user.ListFilter{
        Page:      dto.Page,
        PageSize:  dto.PageSize,
        SortBy:    dto.SortBy,
        SortOrder: dto.SortOrder,
    }
    
    if dto.UserType != nil {
        userType := user.UserType(*dto.UserType)
        filter.UserType = &userType
    }
    
    if dto.Country != nil {
        filter.Country = dto.Country
    }
    
    if dto.MinRating != nil {
        filter.MinReputation = dto.MinRating
    }
    
    if dto.IsTopRated != nil {
        filter.IsTopRated = dto.IsTopRated
    }
    
    if dto.IsFeatured != nil {
        filter.IsFeatured = dto.IsFeatured
    }
    
    return filter
}

// ============================================================================
// HELPER FUNCTIONS
// ============================================================================

func getDeletedAt(deletedAt gorm.DeletedAt) *time.Time {
    if deletedAt.Valid {
        return &deletedAt.Time
    }
    return nil
}

func calculateTrustScore(u *user.User) float64 {
    score := 0.0
    
    // Email verification (20 points)
    if u.EmailVerified {
        score += 20
    }
    
    // Phone verification (10 points)
    if u.PhoneVerified {
        score += 10
    }
    
    // Identity verification (25 points)
    if u.IdentityVerified {
        score += 25
    }
    
    // Payment verification (15 points)
    if u.PaymentVerified {
        score += 15
    }
    
    // Background check (10 points)
    if u.BackgroundChecked {
        score += 10
    }
    
    // Reputation score (20 points max)
    // Map 0-5 reputation to 0-20 points
    score += (u.ReputationScore / 5.0) * 20
    
    // Penalties
    if u.HasActiveWarnings {
        score -= 10
    }
    
    if u.WarningCount > 0 {
        score -= float64(u.WarningCount) * 2
    }
    
    if u.SuspensionCount > 0 {
        score -= float64(u.SuspensionCount) * 5
    }
    
    // Ensure score is between 0-100
    if score < 0 {
        score = 0
    }
    if score > 100 {
        score = 100
    }
    
    return score
}

func calculateTrustLevel(u *user.User) string {
    score := calculateTrustScore(u)
    
    if score >= 80 && u.EmailVerified && u.IdentityVerified && u.PaymentVerified {
        return "trusted"
    }
    
    if score >= 60 && u.EmailVerified && u.IdentityVerified {
        return "verified"
    }
    
    if score >= 30 && u.EmailVerified {
        return "basic"
    }
    
    return "unverified"
}

// ============================================================================
// VALIDATION HELPERS
// ============================================================================

func ValidateCreateUserDTO(dto CreateUserDTO) error {
    if dto.Username == "" {
        return user.ErrInvalidUsername
    }
    
    if _, err := user.NewEmail(dto.Email); err != nil {
        return err
    }
    
    if dto.UserType != "freelancer" && dto.UserType != "client" && dto.UserType != "both" {
        return user.ErrInvalidUserType
    }
    
    return nil
}

func ValidateUpdateUserDTO(dto UpdateUserDTO) error {
    // Add validation logic here
    return nil
}

// ============================================================================
// COMPARISON HELPERS
// ============================================================================

func HasUserChanged(original *user.User, dto UpdateUserDTO) bool {
    if dto.FirstName != nil && *dto.FirstName != original.FirstName {
        return true
    }
    if dto.LastName != nil && *dto.LastName != original.LastName {
        return true
    }
    if dto.DisplayName != nil && *dto.DisplayName != original.DisplayName {
        return true
    }
    if dto.PhoneNumber != nil && *dto.PhoneNumber != original.PhoneNumber {
        return true
    }
    if dto.Country != nil && *dto.Country != original.Country {
        return true
    }
    if dto.Timezone != nil && *dto.Timezone != original.Timezone {
        return true
    }
    if dto.Language != nil && *dto.Language != original.Language {
        return true
    }
    if dto.Currency != nil && *dto.Currency != original.Currency {
        return true
    }
    return false
}

// ============================================================================
// FORMATTING HELPERS
// ============================================================================

func FormatUserFullName(u *user.User) string {
    if u.MiddleName != "" {
        return u.FirstName + " " + u.MiddleName + " " + u.LastName
    }
    return u.FirstName + " " + u.LastName
}

func FormatUserShortName(u *user.User) string {
    return u.FirstName + " " + string([]rune(u.LastName)[0]) + "."
}

func MaskEmail(email string) string {
    parts := strings.Split(email, "@")
    if len(parts) != 2 {
        return email
    }
    
    username := parts[0]
    domain := parts[1]
    
    if len(username) <= 2 {
        return username[0:1] + "***@" + domain
    }
    
    return username[0:2] + "***@" + domain
}

func MaskPhone(phone string) string {
    if len(phone) <= 4 {
        return "***" + phone
    }
    return "***" + phone[len(phone)-4:]
}

// ============================================================================
// ENRICHMENT HELPERS
// ============================================================================

func EnrichUserDTOWithTrustScore(dto *UserDTO, u *user.User) *UserDTO {
    // Add computed trust score fields
    trustScore := calculateTrustScore(u)
    trustLevel := calculateTrustLevel(u)
    
    // Could extend DTO with these if needed
    // For now, they can be fetched separately via ToUserTrustScoreDTO
    _ = trustScore
    _ = trustLevel
    
    return dto
}

func SanitizeUserDTOForPublic(dto *UserDTO) *PublicUserDTO {
    return &PublicUserDTO{
        ID:              dto.ID,
        Username:        dto.Username,
        DisplayName:     dto.DisplayName,
        UserType:        dto.UserType,
        Country:         dto.Country,
        ReputationScore: dto.ReputationScore,
        TotalRatings:    dto.TotalRatings,
        IsTopRated:      dto.IsTopRated,
        IsFeatured:      dto.IsFeatured,
        IsOnline:        dto.IsOnline,
        LastSeenAt:      dto.LastSeenAt,
        MemberSince:     dto.CreatedAt,
    }
}

// ============================================================================
// IMPORT HELPERS (for gorm)
// ============================================================================

import (
    "strings"
    "gorm.io/gorm"
)