// internal/application/user/dto.go
package user

import "time"

// ============================================================================
// USER DTOs
// ============================================================================

// UserDTO represents the complete user data transfer object
type UserDTO struct {
    // Identifiers
    ID           string `json:"id"`
    KeycloakID   string `json:"keycloak_id"`
    Username     string `json:"username"`
    Email        string `json:"email"`
    
    // Personal info
    FirstName    string  `json:"first_name"`
    LastName     string  `json:"last_name"`
    MiddleName   string  `json:"middle_name,omitempty"`
    DisplayName  string  `json:"display_name"`
    
    // Classification
    UserType     string `json:"user_type"`
    Status       string `json:"status"`
    
    // Contact
    PhoneNumber      string `json:"phone_number,omitempty"`
    PhoneCountryCode string `json:"phone_country_code,omitempty"`
    PhoneVerified    bool   `json:"phone_verified"`
    
    // Verification
    EmailVerified       bool       `json:"email_verified"`
    EmailVerifiedAt     *time.Time `json:"email_verified_at,omitempty"`
    IdentityVerified    bool       `json:"identity_verified"`
    IdentityVerifiedAt  *time.Time `json:"identity_verified_at,omitempty"`
    PaymentVerified     bool       `json:"payment_verified"`
    PaymentVerifiedAt   *time.Time `json:"payment_verified_at,omitempty"`
    
    // Security
    TwoFactorEnabled bool       `json:"two_factor_enabled"`
    LastLoginAt      *time.Time `json:"last_login_at,omitempty"`
    
    // Profile
    ProfileCompleteness int  `json:"profile_completeness"`
    OnboardingCompleted bool `json:"onboarding_completed"`
    
    // Location & preferences
    Country  string `json:"country,omitempty"`
    Timezone string `json:"timezone"`
    Language string `json:"language"`
    Currency string `json:"currency"`
    
    // Reputation
    ReputationScore     float64 `json:"reputation_score"`
    TotalRatings        int     `json:"total_ratings"`
    AverageRating       float64 `json:"average_rating"`
    ResponseRatePercent float64 `json:"response_rate_percent"`
    
    // Activity
    IsOnline   bool       `json:"is_online"`
    LastSeenAt *time.Time `json:"last_seen_at,omitempty"`
    
    // Metrics
    TotalEarnings float64 `json:"total_earnings"`
    TotalSpent    float64 `json:"total_spent"`
    TotalJobs     int     `json:"total_jobs"`
    TotalHires    int     `json:"total_hires"`
    CompletedJobs int     `json:"completed_jobs"`
    
    // Badges & features
    IsTopRated  bool       `json:"is_top_rated"`
    IsFeatured  bool       `json:"is_featured"`
    IsPremium   bool       `json:"is_premium"`
    PremiumUntil *time.Time `json:"premium_until,omitempty"`
    
    // Moderation
    HasActiveWarnings bool `json:"has_active_warnings"`
    WarningCount      int  `json:"warning_count"`
    SuspensionCount   int  `json:"suspension_count"`
    
    // Referral
    ReferralCode   string  `json:"referral_code"`
    TotalReferrals int     `json:"total_referrals"`
    
    // Timestamps
    CreatedAt time.Time  `json:"created_at"`
    UpdatedAt time.Time  `json:"updated_at"`
    DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

// PublicUserDTO represents limited public user information
type PublicUserDTO struct {
    ID              string  `json:"id"`
    Username        string  `json:"username"`
    DisplayName     string  `json:"display_name"`
    UserType        string  `json:"user_type"`
    Country         string  `json:"country,omitempty"`
    ReputationScore float64 `json:"reputation_score"`
    TotalRatings    int     `json:"total_ratings"`
    IsTopRated      bool    `json:"is_top_rated"`
    IsFeatured      bool    `json:"is_featured"`
    IsOnline        bool    `json:"is_online"`
    LastSeenAt      *time.Time `json:"last_seen_at,omitempty"`
    MemberSince     time.Time  `json:"member_since"`
}

// UserProfileSummaryDTO represents a summary for listings
type UserProfileSummaryDTO struct {
    ID              string  `json:"id"`
    Username        string  `json:"username"`
    DisplayName     string  `json:"display_name"`
    UserType        string  `json:"user_type"`
    ReputationScore float64 `json:"reputation_score"`
    TotalRatings    int     `json:"total_ratings"`
    IsTopRated      bool    `json:"is_top_rated"`
    IsOnline        bool    `json:"is_online"`
}

// ============================================================================
// CREATE DTOs
// ============================================================================

type CreateUserDTO struct {
    KeycloakID       string `json:"keycloak_id" binding:"required"`
    Username         string `json:"username" binding:"required,min=3,max=50"`
    Email            string `json:"email" binding:"required,email"`
    FirstName        string `json:"first_name" binding:"required,min=1,max=100"`
    LastName         string `json:"last_name" binding:"required,min=1,max=100"`
    UserType         string `json:"user_type" binding:"required,oneof=freelancer client both"`
    Language         string `json:"language" binding:"required"`
    Country          string `json:"country" binding:"required,len=2"`
    Timezone         string `json:"timezone" binding:"required"`
    Currency         string `json:"currency" binding:"required,len=3"`
    ReferredBy       string `json:"referred_by,omitempty"`
}

// ============================================================================
// UPDATE DTOs
// ============================================================================

type UpdateUserDTO struct {
    FirstName        *string `json:"first_name,omitempty" binding:"omitempty,min=1,max=100"`
    LastName         *string `json:"last_name,omitempty" binding:"omitempty,min=1,max=100"`
    MiddleName       *string `json:"middle_name,omitempty" binding:"omitempty,max=100"`
    DisplayName      *string `json:"display_name,omitempty" binding:"omitempty,min=1,max=200"`
    PhoneNumber      *string `json:"phone_number,omitempty"`
    PhoneCountryCode *string `json:"phone_country_code,omitempty"`
    Country          *string `json:"country,omitempty" binding:"omitempty,len=2"`
    Timezone         *string `json:"timezone,omitempty"`
    Language         *string `json:"language,omitempty"`
    Currency         *string `json:"currency,omitempty" binding:"omitempty,len=3"`
}

type UpdateUserStatsDTO struct {
    TotalEarnings *float64 `json:"total_earnings,omitempty"`
    TotalSpent    *float64 `json:"total_spent,omitempty"`
    TotalJobs     *int     `json:"total_jobs,omitempty"`
    TotalHires    *int     `json:"total_hires,omitempty"`
    CompletedJobs *int     `json:"completed_jobs,omitempty"`
    NewRating     *float64 `json:"new_rating,omitempty" binding:"omitempty,min=0,max=5"`
}

// ============================================================================
// LIST RESPONSE DTOs
// ============================================================================

type UserListResponseDTO struct {
    Users      []*UserDTO `json:"users"`
    Total      int64      `json:"total"`
    Page       int        `json:"page"`
    PageSize   int        `json:"page_size"`
    TotalPages int        `json:"total_pages"`
}

type PublicUserListResponseDTO struct {
    Users      []*PublicUserDTO `json:"users"`
    Total      int64            `json:"total"`
    Page       int              `json:"page"`
    PageSize   int              `json:"page_size"`
    TotalPages int              `json:"total_pages"`
}

// ============================================================================
// STATISTICS DTOs
// ============================================================================

type UserStatisticsDTO struct {
    TotalFreelancers        int64   `json:"total_freelancers"`
    TotalClients            int64   `json:"total_clients"`
    TotalActive             int64   `json:"total_active"`
    TotalSuspended          int64   `json:"total_suspended"`
    TotalBanned             int64   `json:"total_banned"`
    TotalVerified           int64   `json:"total_verified"`
    AvgFreelancerReputation float64 `json:"avg_freelancer_reputation"`
    AvgClientReputation     float64 `json:"avg_client_reputation"`
}

// ============================================================================
// ADMIN DTOs
// ============================================================================

type SuspendUserDTO struct {
    Reason      string     `json:"reason" binding:"required"`
    Description string     `json:"description"`
    Duration    *int       `json:"duration"` // Days, nil = indefinite
    EndDate     *time.Time `json:"end_date,omitempty"`
}

type BanUserDTO struct {
    Reason      string `json:"reason" binding:"required"`
    Description string `json:"description"`
    IsPermanent bool   `json:"is_permanent"`
    ExpiresAt   *time.Time `json:"expires_at,omitempty"`
}

type SetPremiumDTO struct {
    ExpiresAt time.Time `json:"expires_at" binding:"required"`
}

// ============================================================================
// VERIFICATION DTOs
// ============================================================================

type VerifyEmailDTO struct {
    UserID string `json:"user_id" binding:"required"`
    Token  string `json:"token" binding:"required"`
}

type VerifyPhoneDTO struct {
    UserID string `json:"user_id" binding:"required"`
    Code   string `json:"code" binding:"required,len=6"`
}

// ============================================================================
// QUERY DTOs
// ============================================================================

type UserSearchQueryDTO struct {
    Query    string  `form:"q"`
    UserType *string `form:"user_type"`
    Country  *string `form:"country"`
    MinRating *float64 `form:"min_rating"`
    IsTopRated *bool  `form:"is_top_rated"`
    IsFeatured *bool  `form:"is_featured"`
    IsOnline   *bool  `form:"is_online"`
    Page       int    `form:"page" binding:"min=1"`
    PageSize   int    `form:"page_size" binding:"min=1,max=100"`
    SortBy     string `form:"sort_by"`
    SortOrder  string `form:"sort_order" binding:"oneof=asc desc"`
}

type UserFilterDTO struct {
    Status         *string    `form:"status"`
    UserType       *string    `form:"user_type"`
    Country        *string    `form:"country"`
    IsVerified     *bool      `form:"is_verified"`
    IsTopRated     *bool      `form:"is_top_rated"`
    IsFeatured     *bool      `form:"is_featured"`
    MinReputation  *float64   `form:"min_reputation"`
    CreatedAfter   *time.Time `form:"created_after"`
    CreatedBefore  *time.Time `form:"created_before"`
    IncludeDeleted bool       `form:"include_deleted"`
    Page           int        `form:"page" binding:"required,min=1"`
    PageSize       int        `form:"page_size" binding:"required,min=1,max=100"`
    SortBy         string     `form:"sort_by"`
    SortOrder      string     `form:"sort_order" binding:"oneof=asc desc"`
}

// ============================================================================
// VALIDATION DTOs
// ============================================================================

type CheckUsernameAvailabilityDTO struct {
    Username string `json:"username" binding:"required,min=3,max=50"`
}

type CheckEmailAvailabilityDTO struct {
    Email string `json:"email" binding:"required,email"`
}

type UsernameAvailabilityResponseDTO struct {
    Available bool   `json:"available"`
    Message   string `json:"message,omitempty"`
}

type EmailAvailabilityResponseDTO struct {
    Available bool   `json:"available"`
    Message   string `json:"message,omitempty"`
}

// ============================================================================
// BULK OPERATION DTOs
// ============================================================================

type BulkUserActionDTO struct {
    UserIDs []string `json:"user_ids" binding:"required,min=1"`
    Action  string   `json:"action" binding:"required,oneof=suspend unsuspend ban unban delete restore feature unfeature"`
    Reason  string   `json:"reason,omitempty"`
}

type BulkUserActionResponseDTO struct {
    SuccessCount int      `json:"success_count"`
    FailedCount  int      `json:"failed_count"`
    FailedUserIDs []string `json:"failed_user_ids,omitempty"`
    Errors       []string `json:"errors,omitempty"`
}

// ============================================================================
// ACTIVITY DTOs
// ============================================================================

type UserActivityDTO struct {
    UserID     string    `json:"user_id"`
    Action     string    `json:"action"`
    IPAddress  string    `json:"ip_address,omitempty"`
    UserAgent  string    `json:"user_agent,omitempty"`
    Timestamp  time.Time `json:"timestamp"`
}

type LoginHistoryDTO struct {
    LastLoginAt       *time.Time `json:"last_login_at"`
    LastLoginIP       string     `json:"last_login_ip,omitempty"`
    LoginCount        int        `json:"login_count"`
    FailedAttempts    int        `json:"failed_login_attempts"`
    LastFailedLoginAt *time.Time `json:"last_failed_login_at,omitempty"`
}

// ============================================================================
// NOTIFICATION PREFERENCE DTOs
// ============================================================================

type UpdateNotificationPreferencesDTO struct {
    MarketingEmails   *bool `json:"marketing_emails,omitempty"`
    ProductUpdates    *bool `json:"product_updates,omitempty"`
    Newsletter        *bool `json:"newsletter,omitempty"`
    SmsNotifications  *bool `json:"sms_notifications,omitempty"`
    PushNotifications *bool `json:"push_notifications,omitempty"`
}

type NotificationPreferencesDTO struct {
    MarketingEmails   bool `json:"marketing_emails"`
    ProductUpdates    bool `json:"product_updates"`
    Newsletter        bool `json:"newsletter"`
    SmsNotifications  bool `json:"sms_notifications"`
    PushNotifications bool `json:"push_notifications"`
}

// ============================================================================
// EXPORT DTOs
// ============================================================================

type UserExportDTO struct {
    Format  string   `form:"format" binding:"required,oneof=csv json xlsx"`
    Fields  []string `form:"fields"`
    UserIDs []string `form:"user_ids"`
}

type UserDataExportResponseDTO struct {
    DownloadURL string    `json:"download_url"`
    ExpiresAt   time.Time `json:"expires_at"`
    FileSize    int64     `json:"file_size"`
    RecordCount int       `json:"record_count"`
}

// ============================================================================
// TRUST & SAFETY DTOs
// ============================================================================

type UserTrustScoreDTO struct {
    UserID            string  `json:"user_id"`
    TrustLevel        string  `json:"trust_level"` // unverified, basic, verified, trusted
    TrustScore        float64 `json:"trust_score"` // 0-100
    EmailVerified     bool    `json:"email_verified"`
    PhoneVerified     bool    `json:"phone_verified"`
    IdentityVerified  bool    `json:"identity_verified"`
    PaymentVerified   bool    `json:"payment_verified"`
    BackgroundChecked bool    `json:"background_checked"`
    ReputationScore   float64 `json:"reputation_score"`
    TotalRatings      int     `json:"total_ratings"`
    HasActiveWarnings bool    `json:"has_active_warnings"`
    AccountAge        int     `json:"account_age_days"`
}

type ReportUserDTO struct {
    ReporterID  string `json:"reporter_id" binding:"required"`
    ReportedID  string `json:"reported_id" binding:"required"`
    Reason      string `json:"reason" binding:"required"`
    Description string `json:"description" binding:"required,min=10"`
    Evidence    []string `json:"evidence,omitempty"` // URLs to screenshots, etc.
}

// ============================================================================
// MAPPER FUNCTIONS
// ============================================================================

// Helper functions would go in mapper.go, but defining interfaces here

type UserMapper interface {
    ToUserDTO(u *user.User) *UserDTO
    ToPublicUserDTO(u *user.User) *PublicUserDTO
    ToUserProfileSummaryDTO(u *user.User) *UserProfileSummaryDTO
    ToUserListResponse(users []*user.User, total int64, page, pageSize int) *UserListResponseDTO
}