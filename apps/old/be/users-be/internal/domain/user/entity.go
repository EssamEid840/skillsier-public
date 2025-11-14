// internal/domain/user/entity.go
package user

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

// User represents the core user aggregate root
// This is the foundation entity for all users (Freelancers, Clients, Staff)
type User struct {
	// ========================================================================
	// IDENTITY & CORE INFO
	// ========================================================================
	ID          string `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	KeycloakID  string `gorm:"type:varchar(255);uniqueIndex:idx_keycloak_id;not null" json:"keycloak_id"`
	Username    string `gorm:"type:varchar(50);uniqueIndex:idx_username;not null" json:"username"`
	Email       Email  `gorm:"embedded;embeddedPrefix:email_" json:"email"`
	
	// ========================================================================
	// PERSONAL INFORMATION
	// ========================================================================
	FirstName   string  `gorm:"type:varchar(100);not null" json:"first_name"`
	LastName    string  `gorm:"type:varchar(100);not null" json:"last_name"`
	FullName    string  `gorm:"type:varchar(200);index:idx_full_name" json:"full_name"` // Computed: FirstName + LastName
	DisplayName string  `gorm:"type:varchar(200)" json:"display_name"` // User's preferred display name
	MiddleName  string  `gorm:"type:varchar(100)" json:"middle_name,omitempty"`
	Nickname    string  `gorm:"type:varchar(100)" json:"nickname,omitempty"`
	
	// ========================================================================
	// CONTACT INFORMATION
	// ========================================================================
	Phone       *Phone    `gorm:"embedded;embeddedPrefix:phone_" json:"phone,omitempty"`
	Address     *Address  `gorm:"embedded;embeddedPrefix:address_" json:"address,omitempty"`
	Location    *Location `gorm:"embedded;embeddedPrefix:location_" json:"location,omitempty"`
	
	// ========================================================================
	// DEMOGRAPHICS
	// ========================================================================
	DateOfBirth  *time.Time `gorm:"type:date" json:"date_of_birth,omitempty"`
	Gender       Gender     `gorm:"type:varchar(30)" json:"gender,omitempty"`
	Nationality  string     `gorm:"type:varchar(100)" json:"nationality,omitempty"`
	Languages    []string   `gorm:"type:jsonb" json:"languages,omitempty"` // Array of language codes
	
	// ========================================================================
	// PROFILE MEDIA
	// ========================================================================
	ProfilePictureURL string `gorm:"type:text" json:"profile_picture_url,omitempty"`
	CoverImageURL     string `gorm:"type:text" json:"cover_image_url,omitempty"`
	VideoIntroURL     string `gorm:"type:text" json:"video_intro_url,omitempty"`
	ThumbnailURL      string `gorm:"type:text" json:"thumbnail_url,omitempty"` // Small profile pic
	
	// ========================================================================
	// USER TYPE & ROLES
	// ========================================================================
	UserType         UserType   `gorm:"type:varchar(20);not null;index:idx_user_type" json:"user_type"`
	AdditionalTypes  []UserType `gorm:"type:jsonb" json:"additional_types,omitempty"` // For users who are both freelancer & client
	
	// ========================================================================
	// ACCOUNT STATUS & VERIFICATION
	// ========================================================================
	Status             AccountStatus      `gorm:"type:varchar(20);not null;index:idx_status;default:'PENDING'" json:"status"`
	VerificationStatus VerificationStatus `gorm:"type:varchar(30);default:'UNVERIFIED'" json:"verification_status"`
	EmailVerified      bool               `gorm:"default:false" json:"email_verified"`
	PhoneVerified      bool               `gorm:"default:false" json:"phone_verified"`
	IdentityVerified   bool               `gorm:"default:false" json:"identity_verified"`
	
	// ========================================================================
	// PROFILE CONTENT
	// ========================================================================
	Bio             string `gorm:"type:text" json:"bio,omitempty"`
	Tagline         string `gorm:"type:varchar(200)" json:"tagline,omitempty"` // Professional headline
	Title           string `gorm:"type:varchar(200)" json:"title,omitempty"`   // Professional title
	Overview        string `gorm:"type:text" json:"overview,omitempty"`        // Extended bio
	Website         string `gorm:"type:varchar(500)" json:"website,omitempty"`
	
	// ========================================================================
	// SOCIAL LINKS
	// ========================================================================
	SocialLinks map[string]string `gorm:"type:jsonb" json:"social_links,omitempty"` // linkedin, github, twitter, etc.
	
	// ========================================================================
	// PROFILE SETTINGS
	// ========================================================================
	ProfileVisibility  ProfileVisibility `gorm:"type:varchar(20);default:'PUBLIC'" json:"profile_visibility"`
	ShowEmail          bool              `gorm:"default:false" json:"show_email"`
	ShowPhone          bool              `gorm:"default:false" json:"show_phone"`
	ShowLocation       bool              `gorm:"default:true" json:"show_location"`
	SearchableProfile  bool              `gorm:"default:true;index:idx_searchable" json:"searchable_profile"`
	AcceptingWork      bool              `gorm:"default:true" json:"accepting_work"`
	
	// ========================================================================
	// AVAILABILITY
	// ========================================================================
	AvailabilityStatus AvailabilityStatus `gorm:"type:varchar(20);default:'AVAILABLE'" json:"availability_status"`
	HoursPerWeek       int                `gorm:"default:0" json:"hours_per_week,omitempty"` // Available hours per week
	PreferredWorkType  string             `gorm:"type:varchar(50)" json:"preferred_work_type,omitempty"` // remote, onsite, hybrid
	
	// ========================================================================
	// RATINGS & REPUTATION
	// ========================================================================
	Rating               float64 `gorm:"type:decimal(3,2);default:0.00;index:idx_rating" json:"rating"` // 0.00 to 5.00
	TotalReviews         int     `gorm:"default:0" json:"total_reviews"`
	TotalJobs            int     `gorm:"default:0" json:"total_jobs"`
	CompletedJobs        int     `gorm:"default:0;index:idx_completed_jobs" json:"completed_jobs"`
	SuccessRate          float64 `gorm:"type:decimal(5,2);default:0.00" json:"success_rate"` // 0-100%
	ResponseTime         int     `gorm:"default:0" json:"response_time"` // Average response time in minutes
	RecommendationScore  float64 `gorm:"type:decimal(5,2);default:0.00" json:"recommendation_score"` // 0-100
	
	// ========================================================================
	// EARNINGS & SPENDING (Cached from financial-be)
	// ========================================================================
	TotalEarnings  float64 `gorm:"type:decimal(15,2);default:0.00" json:"total_earnings,omitempty"`  // For freelancers
	TotalSpent     float64 `gorm:"type:decimal(15,2);default:0.00" json:"total_spent,omitempty"`     // For clients
	CurrentBalance float64 `gorm:"type:decimal(15,2);default:0.00" json:"current_balance,omitempty"` // Wallet balance
	
	// ========================================================================
	// BADGES & ACHIEVEMENTS
	// ========================================================================
	Badges          []BadgeType `gorm:"type:jsonb" json:"badges,omitempty"`
	IsFeatured      bool        `gorm:"default:false;index:idx_featured" json:"is_featured"`
	IsTopRated      bool        `gorm:"default:false;index:idx_top_rated" json:"is_top_rated"`
	IsRisingTalent  bool        `gorm:"default:false" json:"is_rising_talent"`
	IsExpertVetted  bool        `gorm:"default:false" json:"is_expert_vetted"`
	
	// ========================================================================
	// PROFILE COMPLETENESS
	// ========================================================================
	ProfileCompleteness int  `gorm:"default:0" json:"profile_completeness"` // 0-100%
	ProfileCompleted    bool `gorm:"default:false" json:"profile_completed"`
	
	// ========================================================================
	// SECURITY & COMPLIANCE
	// ========================================================================
	TwoFactorEnabled       bool   `gorm:"default:false" json:"two_factor_enabled"`
	TwoFactorSecret        string `gorm:"type:varchar(255)" json:"-"` // Never expose in JSON
	BackupCodes            []string `gorm:"type:jsonb" json:"-"` // Never expose
	SecurityQuestion       string `gorm:"type:text" json:"-"`
	SecurityAnswer         string `gorm:"type:text" json:"-"`
	LoginAttempts          int    `gorm:"default:0" json:"login_attempts"`
	LockedUntil            *time.Time `gorm:"type:timestamp" json:"locked_until,omitempty"`
	LastPasswordChange     *time.Time `gorm:"type:timestamp" json:"last_password_change,omitempty"`
	PasswordResetRequired  bool   `gorm:"default:false" json:"password_reset_required"`
	
	// ========================================================================
	// ACTIVITY TRACKING
	// ========================================================================
	LastLoginAt     *time.Time `gorm:"type:timestamp;index:idx_last_login" json:"last_login_at,omitempty"`
	LastSeenAt      *time.Time `gorm:"type:timestamp" json:"last_seen_at,omitempty"`
	LastActiveAt    *time.Time `gorm:"type:timestamp" json:"last_active_at,omitempty"`
	IsOnline        bool       `gorm:"default:false;index:idx_online" json:"is_online"`
	LoginCount      int        `gorm:"default:0" json:"login_count"`
	LastLoginIP     string     `gorm:"type:varchar(45)" json:"last_login_ip,omitempty"` // IPv4 or IPv6
	LastUserAgent   string     `gorm:"type:text" json:"last_user_agent,omitempty"`
	
	// ========================================================================
	// REFERRALS & MARKETING
	// ========================================================================
	ReferralCode    string `gorm:"type:varchar(50);uniqueIndex:idx_referral_code" json:"referral_code,omitempty"`
	ReferredBy      string `gorm:"type:uuid;index:idx_referred_by" json:"referred_by,omitempty"` // User ID who referred
	ReferralCount   int    `gorm:"default:0" json:"referral_count"` // Number of successful referrals
	MarketingOptIn  bool   `gorm:"default:false" json:"marketing_opt_in"`
	NewsletterOptIn bool   `gorm:"default:false" json:"newsletter_opt_in"`
	
	// ========================================================================
	// ADMIN & MODERATION
	// ========================================================================
	WarningCount    int        `gorm:"default:0" json:"warning_count"`
	SuspensionCount int        `gorm:"default:0" json:"suspension_count"`
	BanCount        int        `gorm:"default:0" json:"ban_count"`
	FlagCount       int        `gorm:"default:0" json:"flag_count"` // Times user was flagged
	Notes           string     `gorm:"type:text" json:"notes,omitempty"` // Admin notes
	Tags            []string   `gorm:"type:jsonb" json:"tags,omitempty"` // Admin tags
	
	// ========================================================================
	// PREFERENCES (stored as JSON for flexibility)
	// ========================================================================
	Preferences map[string]interface{} `gorm:"type:jsonb" json:"preferences,omitempty"`
	Settings    map[string]interface{} `gorm:"type:jsonb" json:"settings,omitempty"`
	
	// ========================================================================
	// METADATA
	// ========================================================================
	Metadata map[string]interface{} `gorm:"type:jsonb" json:"metadata,omitempty"` // Flexible key-value storage
	
	// ========================================================================
	// TIMESTAMPS
	// ========================================================================
	CreatedAt time.Time  `gorm:"autoCreateTime;index:idx_created_at" json:"created_at"`
	UpdatedAt time.Time  `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt *time.Time `gorm:"index:idx_deleted_at" json:"deleted_at,omitempty"`
	DeletedBy string     `gorm:"type:uuid" json:"deleted_by,omitempty"` // Admin who deleted
	
	// ========================================================================
	// AUDIT TRAIL
	// ========================================================================
	CreatedBy       string `gorm:"type:uuid" json:"created_by,omitempty"`
	UpdatedBy       string `gorm:"type:uuid" json:"updated_by,omitempty"`
	LastModifiedBy  string `gorm:"type:uuid" json:"last_modified_by,omitempty"`
	Version         int    `gorm:"default:1" json:"version"` // Optimistic locking
}

// ============================================================================
// CONSTRUCTOR
// ============================================================================

// NewUser creates a new User entity with required fields
func NewUser(keycloakID, username, email, firstName, lastName string, userType UserType) (*User, error) {
	// Validate required fields
	if keycloakID == "" {
		return nil, ErrInvalidKeycloakID
	}
	if username == "" {
		return nil, ErrUsernameRequired
	}
	if email == "" {
		return nil, ErrEmailRequired
	}
	if firstName == "" {
		return nil, ErrFirstNameRequired
	}
	if lastName == "" {
		return nil, ErrLastNameRequired
	}
	if !userType.Valid() {
		return nil, ErrInvalidUserType
	}
	
	// Create and validate email
	emailVO, err := NewEmail(email)
	if err != nil {
		return nil, err
	}
	
	// Validate username
	if err := validateUsername(username); err != nil {
		return nil, err
	}
	
	// Create user
	user := &User{
		KeycloakID:          keycloakID,
		Username:            sanitizeUsername(username),
		Email:               emailVO,
		FirstName:           strings.TrimSpace(firstName),
		LastName:            strings.TrimSpace(lastName),
		FullName:            fmt.Sprintf("%s %s", strings.TrimSpace(firstName), strings.TrimSpace(lastName)),
		DisplayName:         strings.TrimSpace(firstName), // Default to first name
		UserType:            userType,
		Status:              AccountStatusPending, // Starts as pending until email verified
		VerificationStatus:  VerificationStatusUnverified,
		ProfileVisibility:   ProfileVisibilityPublic,
		SearchableProfile:   true,
		AcceptingWork:       userType.IsFreelancer(), // Freelancers default to accepting work
		AvailabilityStatus:  AvailabilityStatusAvailable,
		Rating:              0.00,
		ProfileCompleteness: 20, // Basic info provided = 20%
		AdditionalTypes:     []UserType{},
		Badges:              []BadgeType{},
		Preferences:         make(map[string]interface{}),
		Settings:            make(map[string]interface{}),
		Metadata:            make(map[string]interface{}),
		SocialLinks:         make(map[string]string),
		Version:             1,
	}
	
	// Generate referral code
	user.ReferralCode = generateReferralCode(username)
	
	return user, nil
}

// ============================================================================
// BUSINESS METHODS
// ============================================================================

// Activate activates the user account
func (u *User) Activate() error {
	if u.Status == AccountStatusBanned {
		return ErrCannotActivateBannedUser
	}
	if u.Status == AccountStatusDeleted {
		return ErrCannotReactivateDeletedUser
	}
	if u.Status == AccountStatusActive {
		return ErrAlreadyActive
	}
	
	u.Status = AccountStatusActive
	u.UpdatedAt = time.Now()
	u.Version++
	return nil
}

// Suspend suspends the user account
func (u *User) Suspend(reason string, suspendedBy string) error {
	if u.Status == AccountStatusSuspended {
		return ErrAlreadySuspended
	}
	if u.Status == AccountStatusBanned {
		return fmt.Errorf("cannot suspend banned user")
	}
	if u.Status == AccountStatusDeleted {
		return fmt.Errorf("cannot suspend deleted user")
	}
	
	u.Status = AccountStatusSuspended
	u.SuspensionCount++
	u.UpdatedAt = time.Now()
	u.LastModifiedBy = suspendedBy
	u.Version++
	
	// Add note
	if u.Notes == "" {
		u.Notes = fmt.Sprintf("Suspended: %s", reason)
	} else {
		u.Notes += fmt.Sprintf("\n[%s] Suspended: %s", time.Now().Format(time.RFC3339), reason)
	}
	
	return nil
}

// Ban permanently bans the user account
func (u *User) Ban(reason string, bannedBy string) error {
	if u.Status == AccountStatusBanned {
		return ErrAlreadyBanned
	}
	
	u.Status = AccountStatusBanned
	u.BanCount++
	u.UpdatedAt = time.Now()
	u.LastModifiedBy = bannedBy
	u.Version++
	
	// Add note
	if u.Notes == "" {
		u.Notes = fmt.Sprintf("Banned: %s", reason)
	} else {
		u.Notes += fmt.Sprintf("\n[%s] Banned: %s", time.Now().Format(time.RFC3339), reason)
	}
	
	return nil
}

// Restore restores a suspended/banned account
func (u *User) Restore(restoredBy string) error {
	if u.Status != AccountStatusSuspended && u.Status != AccountStatusBanned {
		return fmt.Errorf("can only restore suspended or banned accounts")
	}
	
	u.Status = AccountStatusActive
	u.UpdatedAt = time.Now()
	u.LastModifiedBy = restoredBy
	u.Version++
	
	// Add note
	if u.Notes == "" {
		u.Notes = "Account restored"
	} else {
		u.Notes += fmt.Sprintf("\n[%s] Account restored", time.Now().Format(time.RFC3339))
	}
	
	return nil
}

// SoftDelete soft deletes the user
func (u *User) SoftDelete(deletedBy string) error {
	if u.Status == AccountStatusDeleted {
		return ErrAlreadyDeleted
	}
	
	now := time.Now()
	u.Status = AccountStatusDeleted
	u.DeletedAt = &now
	u.DeletedBy = deletedBy
	u.UpdatedAt = now
	u.Version++
	
	return nil
}

// VerifyEmail marks email as verified
func (u *User) VerifyEmail() {
	now := time.Now().Unix()
	u.Email.Verified = true
	u.Email.VerifiedAt = &now
	u.EmailVerified = true
	u.UpdatedAt = time.Now()
	
	// Auto-activate if pending
	if u.Status == AccountStatusPending {
		u.Status = AccountStatusActive
	}
	
	u.Version++
}

// VerifyPhone marks phone as verified
func (u *User) VerifyPhone() {
	if u.Phone != nil {
		now := time.Now().Unix()
		u.Phone.Verified = true
		u.Phone.VerifiedAt = &now
		u.PhoneVerified = true
		u.UpdatedAt = time.Now()
		u.Version++
	}
}

// VerifyIdentity marks identity as verified
func (u *User) VerifyIdentity() {
	u.IdentityVerified = true
	u.VerificationStatus = VerificationStatusVerified
	u.UpdatedAt = time.Now()
	u.Version++
}

// UpdateProfile updates profile information
func (u *User) UpdateProfile(bio, tagline, title string) error {
	if len(bio) > 5000 {
		return ErrInvalidBioLength
	}
	if len(tagline) > 200 {
		return ErrInvalidTaglineLength
	}
	
	u.Bio = strings.TrimSpace(bio)
	u.Tagline = strings.TrimSpace(tagline)
	u.Title = strings.TrimSpace(title)
	u.UpdatedAt = time.Now()
	u.Version++
	
	// Recalculate profile completeness
	u.CalculateProfileCompleteness()
	
	return nil
}

// UpdateLastSeen updates last seen timestamp
func (u *User) UpdateLastSeen() {
	now := time.Now()
	u.LastSeenAt = &now
	u.LastActiveAt = &now
	u.IsOnline = true
}

// SetOffline marks user as offline
func (u *User) SetOffline() {
	u.IsOnline = false
}

// RecordLogin records a successful login
func (u *User) RecordLogin(ipAddress, userAgent string) {
	now := time.Now()
	u.LastLoginAt = &now
	u.LastLoginIP = ipAddress
	u.LastUserAgent = userAgent
	u.LoginCount++
	u.LoginAttempts = 0 // Reset failed attempts on successful login
	u.IsOnline = true
	u.UpdatedAt = now
}

// IncrementFailedLogin increments failed login attempts
func (u *User) IncrementFailedLogin() {
	u.LoginAttempts++
	u.UpdatedAt = time.Now()
	
	// Lock account after 5 failed attempts for 15 minutes
	if u.LoginAttempts >= 5 {
		lockUntil := time.Now().Add(15 * time.Minute)
		u.LockedUntil = &lockUntil
	}
}

// ResetFailedLogins resets failed login attempts
func (u *User) ResetFailedLogins() {
	u.LoginAttempts = 0
	u.LockedUntil = nil
	u.UpdatedAt = time.Now()
}

// IsLocked checks if account is currently locked
func (u *User) IsLocked() bool {
	if u.LockedUntil == nil {
		return false
	}
	return time.Now().Before(*u.LockedUntil)
}

// AddWarning adds a warning to the user account
func (u *User) AddWarning(reason string, issuedBy string) {
	u.WarningCount++
	u.UpdatedAt = time.Now()
	u.LastModifiedBy = issuedBy
	
	if u.Notes == "" {
		u.Notes = fmt.Sprintf("Warning: %s", reason)
	} else {
		u.Notes += fmt.Sprintf("\n[%s] Warning: %s", time.Now().Format(time.RFC3339), reason)
	}
	
	u.Version++
}

// AssignBadge assigns a badge to the user
func (u *User) AssignBadge(badge BadgeType) error {
	if !badge.Valid() {
		return ErrInvalidBadgeType
	}
	
	// Check if badge already assigned
	for _, b := range u.Badges {
		if b == badge {
			return ErrBadgeAlreadyAssigned
		}
	}
	
	u.Badges = append(u.Badges, badge)
	u.UpdatedAt = time.Now()
	u.Version++
	
	// Update badge flags
	switch badge {
	case BadgeTypeTopRated, BadgeTypeTopRatedPlus:
		u.IsTopRated = true
	case BadgeTypeRisingTalent:
		u.IsRisingTalent = true
	case BadgeTypeExpertVetted:
		u.IsExpertVetted = true
	}
	
	return nil
}

// RemoveBadge removes a badge from the user
func (u *User) RemoveBadge(badge BadgeType) {
	newBadges := make([]BadgeType, 0)
	for _, b := range u.Badges {
		if b != badge {
			newBadges = append(newBadges, b)
		}
	}
	u.Badges = newBadges
	u.UpdatedAt = time.Now()
	u.Version++
	
	// Update badge flags
	switch badge {
	case BadgeTypeTopRated, BadgeTypeTopRatedPlus:
		u.IsTopRated = false
	case BadgeTypeRisingTalent:
		u.IsRisingTalent = false
	case BadgeTypeExpertVetted:
		u.IsExpertVetted = false
	}
}

// UpdateRating updates user's rating
func (u *User) UpdateRating(rating float64, totalReviews int) error {
	if rating < 0 || rating > 5 {
		return ErrInvalidRating
	}
	
	u.Rating = rating
	u.TotalReviews = totalReviews
	u.UpdatedAt = time.Now()
	u.Version++
	
	return nil
}

// UpdateStats updates user statistics
func (u *User) UpdateStats(completedJobs, totalJobs int, successRate float64) error {
	if successRate < 0 || successRate > 100 {
		return ErrInvalidCompletionRate
	}
	
	u.CompletedJobs = completedJobs
	u.TotalJobs = totalJobs
	u.SuccessRate = successRate
	u.UpdatedAt = time.Now()
	u.Version++
	
	return nil
}

// UpdateEarnings updates total earnings (for freelancers)
func (u *User) UpdateEarnings(amount float64) {
	u.TotalEarnings = amount
	u.UpdatedAt = time.Now()
	u.Version++
}

// UpdateSpending updates total spending (for clients)
func (u *User) UpdateSpending(amount float64) {
	u.TotalSpent = amount
	u.UpdatedAt = time.Now()
	u.Version++
}

// CalculateProfileCompleteness calculates profile completeness percentage
func (u *User) CalculateProfileCompleteness() int {
	completeness := 0
	
	// Basic info (20%)
	if u.FirstName != "" && u.LastName != "" && u.Email.Value != "" {
		completeness += 20
	}
	
	// Profile picture (10%)
	if u.ProfilePictureURL != "" {
		completeness += 10
	}
	
	// Bio/Overview (15%)
	if u.Bio != "" || u.Overview != "" {
		completeness += 15
	}
	
	// Tagline/Title (10%)
	if u.Tagline != "" || u.Title != "" {
		completeness += 10
	}
	
	// Location (10%)
	if u.Location != nil && u.Location.City != "" {
		completeness += 10
	}
	
	// Phone (5%)
	if u.Phone != nil && u.Phone.Number != "" {
		completeness += 5
	}
	
	// Email verified (10%)
	if u.EmailVerified {
		completeness += 10
	}
	
	// Identity verified (10%)
	if u.IdentityVerified {
		completeness += 10
	}
	
	// Social links (5%)
	if len(u.SocialLinks) > 0 {
		completeness += 5
	}
	
	// Video intro (5%)
	if u.VideoIntroURL != "" {
		completeness += 5
	}
	
	u.ProfileCompleteness = completeness
	u.ProfileCompleted = completeness >= 80
	
	return completeness
}

// SetFeatured marks user as featured
func (u *User) SetFeatured(featured bool) {
	u.IsFeatured = featured
	u.UpdatedAt = time.Now()
	u.Version++
}

// UpdateAvailability updates availability status
func (u *User) UpdateAvailability(status AvailabilityStatus, hoursPerWeek int) error {
	if !status.Valid() {
		return ErrInvalidAvailability
	}
	if hoursPerWeek < 0 || hoursPerWeek > 168 {
		return ErrHoursPerWeekInvalid
	}
	
	u.AvailabilityStatus = status
	u.HoursPerWeek = hoursPerWeek
	u.AcceptingWork = status.IsAvailable()
	u.UpdatedAt = time.Now()
	u.Version++
	
	return nil
}

// ============================================================================
// VALIDATION METHODS
// ============================================================================

// Validate validates the user entity
func (u *User) Validate() error {
	errors := NewValidationErrors()
	
	// Required fields
	if u.KeycloakID == "" {
		errors.Add("keycloak_id", "Keycloak ID is required", u.KeycloakID)
	}
	if u.Username == "" {
		errors.Add("username", "Username is required", u.Username)
	}
	if u.Email.Value == "" {
		errors.Add("email", "Email is required", u.Email.Value)
	}
	if u.FirstName == "" {
		errors.Add("first_name", "First name is required", u.FirstName)
	}
	if u.LastName == "" {
		errors.Add("last_name", "Last name is required", u.LastName)
	}
	
	// Validate enums
	if !u.UserType.Valid() {
		errors.Add("user_type", "Invalid user type", u.UserType)
	}
	if !u.Status.Valid() {
		errors.Add("status", "Invalid account status", u.Status)
	}
	if !u.VerificationStatus.Valid() {
		errors.Add("verification_status", "Invalid verification status", u.VerificationStatus)
	}
	
	// Validate email
	if err := u.Email.Validate(); err != nil {
		errors.Add("email", err.Error(), u.Email.Value)
	}
	
	// Validate username
	if err := validateUsername(u.Username); err != nil {
		errors.Add("username", err.Error(), u.Username)
	}
	
	// Validate phone if present
	if u.Phone != nil {
		if err := u.Phone.Validate(); err != nil {
			errors.Add("phone", err.Error(), u.Phone.Number)
		}
	}
	
	// Validate address if present
	if u.Address != nil {
		if err := u.Address.Validate(); err != nil {
			errors.Add("address", err.Error(), u.Address)
		}
	}
	
	// Validate location if present
	if u.Location != nil {
		if err := u.Location.Validate(); err != nil {
			errors.Add("location", err.Error(), u.Location)
		}
	}
	
	// Validate lengths
	if len(u.Bio) > 5000 {
		errors.Add("bio", "Bio exceeds maximum length of 5000 characters", len(u.Bio))
	}
	if len(u.Tagline) > 200 {
		errors.Add("tagline", "Tagline exceeds maximum length of 200 characters", len(u.Tagline))
	}
	
	// Validate rating
	if u.Rating < 0 || u.Rating > 5 {
		errors.Add("rating", "Rating must be between 0 and 5", u.Rating)
	}
	
	// Validate success rate
	if u.SuccessRate < 0 || u.SuccessRate > 100 {
		errors.Add("success_rate", "Success rate must be between 0 and 100", u.SuccessRate)
	}
	
	// Validate hours per week
	if u.HoursPerWeek < 0 || u.HoursPerWeek > 168 {
		errors.Add("hours_per_week", "Hours per week must be between 0 and 168", u.HoursPerWeek)
	}
	
	if errors.HasErrors() {
		return errors
	}
	
	return nil
}

// ============================================================================
// QUERY HELPERS
// ============================================================================

// CanPerformActions checks if user can perform platform actions
func (u *User) CanPerformActions() bool {
	return u.Status.CanPerformActions() && !u.IsLocked()
}

// CanLogin checks if user can login
func (u *User) CanLogin() bool {
	return u.Status.CanLogin() && !u.IsLocked()
}

// IsActive checks if account is active
func (u *User) IsActive() bool {
	return u.Status == AccountStatusActive
}

// IsBlocked checks if account is blocked
func (u *User) IsBlocked() bool {
	return u.Status.IsBlocked()
}

// IsVerified checks if user is fully verified
func (u *User) IsVerified() bool {
	return u.EmailVerified && u.IdentityVerified
}

// HasCompletedProfile checks if profile is sufficiently complete
func (u *User) HasCompletedProfile() bool {
	return u.ProfileCompleteness >= 80
}

// IsFreelancer checks if user is a freelancer
func (u *User) IsFreelancer() bool {
	return u.UserType.IsFreelancer()
}

// IsClient checks if user is a client
func (u *User) IsClient() bool {
	return u.UserType.IsClient()
}

// IsStaff checks if user is platform staff
func (u *User) IsStaff() bool {
	return u.UserType.IsStaff()
}

// ============================================================================
// UTILITY METHODS
// ============================================================================

// GetFullName returns the full name
func (u *User) GetFullName() string {
	if u.FullName != "" {
		return u.FullName
	}
	return fmt.Sprintf("%s %s", u.FirstName, u.LastName)
}

// GetDisplayName returns the preferred display name
func (u *User) GetDisplayName() string {
	if u.DisplayName != "" {
		return u.DisplayName
	}
	return u.FirstName
}

// GetProfileURL returns the profile URL
func (u *User) GetProfileURL() string {
	return fmt.Sprintf("/users/%s", u.Username)
}

// GetAge calculates age from date of birth
func (u *User) GetAge() int {
	if u.DateOfBirth == nil {
		return 0
	}
	now := time.Now()
	age := now.Year() - u.DateOfBirth.Year()
	if now.YearDay() < u.DateOfBirth.YearDay() {
		age--
	}
	return age
}

// ============================================================================
// PRIVATE HELPER FUNCTIONS
// ============================================================================

// validateUsername validates username format
func validateUsername(username string) error {
	if username == "" {
		return ErrUsernameRequired
	}
	
	username = strings.TrimSpace(username)
	
	if len(username) < 3 {
		return ErrUsernameTooShort
	}
	if len(username) > 50 {
		return ErrUsernameTooLong
	}
	
	// Username must start with letter, contain only alphanumeric, underscore, hyphen
	if !usernameRegex.MatchString(username) {
		return ErrInvalidUsernameFormat
	}
	
	// Cannot start or end with special characters
	if strings.HasPrefix(username, "_") || strings.HasPrefix(username, "-") ||
		strings.HasSuffix(username, "_") || strings.HasSuffix(username, "-") {
		return ErrInvalidUsernameFormat
	}
	
	return nil
}

// sanitizeUsername sanitizes username
func sanitizeUsername(username string) string {
	return strings.TrimSpace(strings.ToLower(username))
}

// generateReferralCode generates a unique referral code
func generateReferralCode(username string) string {
	// Use first 6 chars of username + 4 random chars
	code := strings.ToUpper(username)
	if len(code) > 6 {
		code = code[:6]
	}
	// In production, add random characters
	return code + fmt.Sprintf("%04d", time.Now().Unix()%10000)
}

// ============================================================================
// REGEX PATTERNS
// ============================================================================

// usernameRegex validates username format (starts with letter, contains alphanumeric, underscore, hyphen)
var usernameRegex = regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9_-]*// internal/domain/user/entity.go
package user

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

// User represents the core user aggregate root
// This is the foundation entity for all users (Freelancers, Clients, Staff)
type User struct {
	// ========================================================================
	// IDENTITY & CORE INFO
	// ========================================================================
	ID          string `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	KeycloakID  string `gorm:"type:varchar(255);uniqueIndex:idx_keycloak_id;not null" json:"keycloak_id"`
	Username    string `gorm:"type:varchar(50);uniqueIndex:idx_username;not null" json:"username"`
	Email       Email  `gorm:"embedded;embeddedPrefix:email_" json:"email"`
	
	// ========================================================================
	// PERSONAL INFORMATION
	// ========================================================================
	FirstName   string  `gorm:"type:varchar(100);not null" json:"first_name"`
	LastName    string  `gorm:"type:varchar(100);not null" json:"last_name"`
	FullName    string  `gorm:"type:varchar(200);index:idx_full_name" json:"full_name"` // Computed: FirstName + LastName
	DisplayName string  `gorm:"type:varchar(200)" json:"display_name"` // User's preferred display name
	MiddleName  string  `gorm:"type:varchar(100)" json:"middle_name,omitempty"`
	Nickname    string  `gorm:"type:varchar(100)" json:"nickname,omitempty"`
	
	// ========================================================================
	// CONTACT INFORMATION
	// ========================================================================
	Phone       *Phone    `gorm:"embedded;embeddedPrefix:phone_" json:"phone,omitempty"`
	Address     *Address  `gorm:"embedded;embeddedPrefix:address_" json:"address,omitempty"`
	Location    *Location `gorm:"embedded;embeddedPrefix:location_" json:"location,omitempty"`
	
	// ========================================================================
	// DEMOGRAPHICS
	// ========================================================================
	DateOfBirth  *time.Time `gorm:"type:date" json:"date_of_birth,omitempty"`
	Gender       Gender     `gorm:"type:varchar(30)" json:"gender,omitempty"`
	Nationality  string     `gorm:"type:varchar(100)" json:"nationality,omitempty"`
	Languages    []string   `gorm:"type:jsonb" json:"languages,omitempty"` // Array of language codes
	
	// ========================================================================
	// PROFILE MEDIA
	// ========================================================================
	ProfilePictureURL string `gorm:"type:text" json:"profile_picture_url,omitempty"`
	CoverImageURL     string `gorm:"type:text" json:"cover_image_url,omitempty"`
	VideoIntroURL     string `gorm:"type:text" json:"video_intro_url,omitempty"`
	ThumbnailURL      string `gorm:"type:text" json:"thumbnail_url,omitempty"` // Small profile pic
	
	// ========================================================================
	// USER TYPE & ROLES
	// ========================================================================
	UserType         UserType   `gorm:"type:varchar(20);not null;index:idx_user_type" json:"user_type"`
	AdditionalTypes  []UserType `gorm:"type:jsonb" json:"additional_types,omitempty"` // For users who are both freelancer & client
	
	// ========================================================================
	// ACCOUNT STATUS & VERIFICATION
	// ========================================================================
	Status             AccountStatus      `gorm:"type:varchar(20);not null;index:idx_status;default:'PENDING'" json:"status"`
	VerificationStatus VerificationStatus `gorm:"type:varchar(30);default:'UNVERIFIED'" json:"verification_status"`
	EmailVerified      bool               `gorm:"default:false" json:"email_verified"`
	PhoneVerified      bool               `gorm:"default:false" json:"phone_verified"`
	IdentityVerified   bool               `gorm:"default:false" json:"identity_verified"`
	
	// ========================================================================
	// PROFILE CONTENT
	// ========================================================================
	Bio             string `gorm:"type:text" json:"bio,omitempty"`
	Tagline         string `gorm:"type:varchar(200)" json:"tagline,omitempty"` // Professional headline
	Title           string `gorm:"type:varchar(200)" json:"title,omitempty"`   // Professional title
	Overview        string `gorm:"type:text" json:"overview,omitempty"`        // Extended bio
	Website         string `gorm:"type:varchar(500)" json:"website,omitempty"`
	
	// ========================================================================
	// SOCIAL LINKS
	// ========================================================================
	SocialLinks map[string]string `gorm:"type:jsonb" json:"social_links,omitempty"` // linkedin, github, twitter, etc.
	
	// ========================================================================
	// PROFILE SETTINGS
	// ========================================================================
	ProfileVisibility  ProfileVisibility `gorm:"type:varchar(20);default:'PUBLIC'" json:"profile_visibility"`
	ShowEmail          bool              `gorm:"default:false" json:"show_email"`
	ShowPhone          bool              `gorm:"default:false" json:"show_phone"`
	ShowLocation       bool              `gorm:"default:true" json:"show_location"`
	SearchableProfile  bool              `gorm:"default:true;index:idx_searchable" json:"searchable_profile"`
	AcceptingWork      bool              `gorm:"default:true" json:"accepting_work"`
	
	// ========================================================================
	// AVAILABILITY
	// ========================================================================
	AvailabilityStatus AvailabilityStatus `gorm:"type:varchar(20);default:'AVAILABLE'" json:"availability_status"`
	HoursPerWeek       int                `gorm:"default:0" json:"hours_per_week,omitempty"` // Available hours per week
	PreferredWorkType  string             `gorm:"type:varchar(50)" json:"preferred_work_type,omitempty"` // remote, onsite, hybrid
	
	// ========================================================================
	// RATINGS & REPUTATION
	// ========================================================================
	Rating               float64 `gorm:"type:decimal(3,2);default:0.00;index:idx_rating" json:"rating"` // 0.00 to 5.00
	TotalReviews         int     `gorm:"default:0" json:"total_reviews"`
	TotalJobs            int     `gorm:"default:0" json:"total_jobs"`
	CompletedJobs        int     `gorm:"default:0;index:idx_completed_jobs" json:"completed_jobs"`
	SuccessRate          float64 `gorm:"type:decimal(5,2);default:0.00" json:"success_rate"` // 0-100%
	ResponseTime         int     `gorm:"default:0" json:"response_time"` // Average response time in minutes
	RecommendationScore  float64 `gorm:"type:decimal(5,2);default:0.00" json:"recommendation_score"` // 0-100
	
	// ========================================================================
	// EARNINGS & SPENDING (Cached from financial-be)
	// ========================================================================
	TotalEarnings  float64 `gorm:"type:decimal(15,2);default:0.00" json:"total_earnings,omitempty"`  // For freelancers
	TotalSpent     float64 `gorm:"type:decimal(15,2);default:0.00" json:"total_spent,omitempty"`     // For clients
	CurrentBalance float64 `gorm:"type:decimal(15,2);default:0.00" json:"current_balance,omitempty"` // Wallet balance
	
	// ========================================================================
	// BADGES & ACHIEVEMENTS
	// ========================================================================
	Badges          []BadgeType `gorm:"type:jsonb" json:"badges,omitempty"`
	IsFeatured      bool        `gorm:"default:false;index:idx_featured" json:"is_featured"`
	IsTopRated      bool        `gorm:"default:false;index:idx_top_rated" json:"is_top_rated"`
	IsRisingTalent  bool        `gorm:"default:false" json:"is_rising_talent"`
	IsExpertVetted  bool        `gorm:"default:false" json:"is_expert_vetted"`
	
	// ========================================================================
	// PROFILE COMPLETENESS
	// ========================================================================
	ProfileCompleteness int  `gorm:"default:0" json:"profile_completeness"` // 0-100%
	ProfileCompleted    bool `gorm:"default:false" json:"profile_completed"`
	
	// ========================================================================
	// SECURITY & COMPLIANCE
	// ========================================================================
	TwoFactorEnabled       bool   `gorm:"default:false" json:"two_factor_enabled"`
	TwoFactorSecret        string `gorm:"type:varchar(255)" json:"-"` // Never expose in JSON
	BackupCodes            []string `gorm:"type:jsonb" json:"-"` // Never expose
	SecurityQuestion       string `gorm:"type:text" json:"-"`
	SecurityAnswer         string `gorm:"type:text" json:"-"`
	LoginAttempts          int    `gorm:"default:0" json:"login_attempts"`
	LockedUntil            *time.Time `gorm:"type:timestamp" json:"locked_until,omitempty"`
	LastPasswordChange     *time.Time `gorm:"type:timestamp" json:"last_password_change,omitempty"`
	PasswordResetRequired  bool   `gorm:"default:false" json:"password_reset_required"`
	
	// ========================================================================
	// ACTIVITY TRACKING
	// ========================================================================
	LastLoginAt     *time.Time `gorm:"type:timestamp;index:idx_last_login" json:"last_login_at,omitempty"`
	LastSeenAt      *time.Time `gorm:"type:timestamp" json:"last_seen_at,omitempty"`
	LastActiveAt    *time.Time `gorm:"type:timestamp" json:"last_active_at,omitempty"`
	IsOnline        bool       `gorm:"default:false;index:idx_online" json:"is_online"`
	LoginCount      int        `gorm:"default:0" json:"login_count"`
	LastLoginIP     string     `gorm:"type:varchar(45)" json:"last_login_ip,omitempty"` // IPv4 or IPv6
	LastUserAgent   string     `gorm:"type:text" json:"last_user_agent,omitempty"`
	
	// ========================================================================
	// REFERRALS & MARKETING
	// ========================================================================
	ReferralCode    string `gorm:"type:varchar(50);uniqueIndex:idx_referral_code" json:"referral_code,omitempty"`
	ReferredBy      string `gorm:"type:uuid;index:idx_referred_by" json:"referred_by,omitempty"` // User ID who referred
	ReferralCount   int    `gorm:"default:0" json:"referral_count"` // Number of successful referrals
	MarketingOptIn  bool   `gorm:"default:false" json:"marketing_opt_in"`
	NewsletterOptIn bool   `gorm:"default:false" json:"newsletter_opt_in"`
	
	// ========================================================================
	// ADMIN & MODERATION
	// ========================================================================
	WarningCount    int        `gorm:"default:0" json:"warning_count"`
	SuspensionCount int        `gorm:"default:0" json:"suspension_count"`
	BanCount        int        `gorm:"default:0" json:"ban_count"`
	FlagCount       int        `gorm:"default:0" json:"flag_count"` // Times user was flagged
	Notes           string     `gorm:"type:text" json:"notes,omitempty"` // Admin notes
	Tags            []string   `gorm:"type:jsonb" json:"tags,omitempty"` // Admin tags
	
	// ========================================================================
	// PREFERENCES (stored as JSON for flexibility)
	// ========================================================================
	Preferences map[string]interface{} `gorm:"type:jsonb" json:"preferences,omitempty"`
	Settings    map[string]interface{} `gorm:"type:jsonb" json:"settings,omitempty"`
	
	// ========================================================================
	// METADATA
	// ========================================================================
	Metadata map[string]interface{} `gorm:"type:jsonb" json:"metadata,omitempty"` // Flexible key-value storage
	
	// ========================================================================
	// TIMESTAMPS
	// ========================================================================
	CreatedAt time.Time  `gorm:"autoCreateTime;index:idx_created_at" json:"created_at"`
	UpdatedAt time.Time  `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt *time.Time `gorm:"index:idx_deleted_at" json:"deleted_at,omitempty"`
	DeletedBy string     `gorm:"type:uuid" json:"deleted_by,omitempty"` // Admin who deleted
	
	// ========================================================================
	// AUDIT TRAIL
	// ========================================================================
	CreatedBy       string `gorm:"type:uuid" json:"created_by,omitempty"`
	UpdatedBy       string `gorm:"type:uuid" json:"updated_by,omitempty"`
	LastModifiedBy  string `gorm:"type:uuid" json:"last_modified_by,omitempty"`
	Version         int    `gorm:"default:1" json:"version"` // Optimistic locking
}

// ============================================================================
// CONSTRUCTOR
// ============================================================================

// NewUser creates a new User entity with required fields
func NewUser(keycloakID, username, email, firstName, lastName string, userType UserType) (*User, error) {
	// Validate required fields
	if keycloakID == "" {
		return nil, ErrInvalidKeycloakID
	}
	if username == "" {
		return nil, ErrUsernameRequired
	}
	if email == "" {
		return nil, ErrEmailRequired
	}
	if firstName == "" {
		return nil, ErrFirstNameRequired
	}
	if lastName == "" {
		return nil, ErrLastNameRequired
	}
	if !userType.Valid() {
		return nil, ErrInvalidUserType
	}
	
	// Create and validate email
	emailVO, err := NewEmail(email)
	if err != nil {
		return nil, err
	}
	
	// Validate username
	if err := validateUsername(username); err != nil {
		return nil, err
	}
	
	// Create user
	user := &User{
		KeycloakID:          keycloakID,
		Username:            sanitizeUsername(username),
		Email:               emailVO,
		FirstName:           strings.TrimSpace(firstName),
		LastName:            strings.TrimSpace(lastName),
		FullName:            fmt.Sprintf("%s %s", strings.TrimSpace(firstName), strings.TrimSpace(lastName)),
		DisplayName:         strings.TrimSpace(firstName), // Default to first name
		UserType:            userType,
		Status:              AccountStatusPending, // Starts as pending until email verified
		VerificationStatus:  VerificationStatusUnverified,
		ProfileVisibility:   ProfileVisibilityPublic,
		SearchableProfile:   true,
		AcceptingWork:       userType.IsFreelancer(), // Freelancers default to accepting work
		AvailabilityStatus:  AvailabilityStatusAvailable,
		Rating:              0.00,
		ProfileCompleteness: 20, // Basic info provided = 20%
		AdditionalTypes:     []UserType{},
		Badges:              []BadgeType{},
		Preferences:         make(map[string]interface{}),
		Settings:            make(map[string]interface{}),
		Metadata:            make(map[string]interface{}),
		SocialLinks:         make(map[string]string),
		Version:             1,
	}
	
	// Generate referral code
	user.ReferralCode = generateReferralCode(username)
	
	return user, nil
}

// ============================================================================
// BUSINESS METHODS
// ============================================================================

// Activate activates the user account
func (u *User) Activate() error {
	if u.Status == AccountStatusBanned {
		return ErrCannotActivateBannedUser
	}
	if u.Status == AccountStatusDeleted {
		return ErrCannotReactivateDeletedUser
	}
	if u.Status == AccountStatusActive {
		return ErrAlreadyActive
	}
	
	u.Status = AccountStatusActive
	u.UpdatedAt = time.Now()
	u.Version++
	return nil
}

// Suspend suspends the user account
func (u *User) Suspend(reason string, suspendedBy string) error {
	if u.Status == AccountStatusSuspended {
		return ErrAlreadySuspended
	}
	if u.Status == AccountStatusBanned {
		return fmt.Errorf("cannot suspend banned user")
	}
	if u.Status == AccountStatusDeleted {
		return fmt.Errorf("cannot suspend deleted user")
	}
	
	u.Status = AccountStatusSuspended
	u.SuspensionCount++
	u.UpdatedAt = time.Now()
	u.LastModifiedBy = suspendedBy
	u.Version++
	
	// Add note
	if u.Notes == "" {
		u.Notes = fmt.Sprintf("Suspended: %s", reason)
	} else {
		u.Notes += fmt.Sprintf("\n[%s] Suspended: %s", time.Now().Format(time.RFC3339), reason)
	}
	
	return nil
}

// Ban permanently bans the user account
func (u *User) Ban(reason string, bannedBy string) error {
	if u.Status == AccountStatusBanned {
		return ErrAlreadyBanned
	}
	
	u.Status = AccountStatusBanned
	u.BanCount++
	u.UpdatedAt = time.Now()
	u.LastModifiedBy = bannedBy
	u.Version++
	
	// Add note
	if u.Notes == "" {
		u.Notes = fmt.Sprintf("Banned: %s", reason)
	} else {
		u.Notes += fmt.Sprintf("\n[%s] Banned: %s", time.Now().Format(time.RFC3339), reason)
	}
	
	return nil
}

// Restore restores a suspended/banned account
func (u *User) Restore(restoredBy string) error {
	if u.Status != AccountStatusSuspended && u.Status != AccountStatusBanned {
		return fmt.Errorf("can only restore suspended or banned accounts")
	}
	
	u.Status = AccountStatusActive
	u.UpdatedAt = time.Now()
	u.LastModifiedBy = restoredBy
	u.Version++
	
	// Add note
	if u.Notes == "" {
		u.Notes = "Account restored"
	} else {
		u.Notes += fmt.Sprintf("\n[%s] Account restored", time.Now().Format(time.RFC3339))
	}
	
	return nil
}

// SoftDelete soft deletes the user
func (u *User) SoftDelete(deletedBy string) error {
	if u.Status == AccountStatusDeleted {
		return ErrAlreadyDeleted
	}
	
	now := time.Now()
	u.Status = AccountStatusDeleted
	u.DeletedAt = &now
	u.DeletedBy = deletedBy
	u.UpdatedAt = now
	u.Version++
	
	return nil
}

// VerifyEmail marks email as verified
func (u *User) VerifyEmail() {
	now := time.Now().Unix()
	u.Email.Verified = true
	u.Email.VerifiedAt = &now
	u.EmailVerified = true
	u.UpdatedAt = time.Now()
	
	// Auto-activate if pending
	if u.Status == AccountStatusPending {
		u.Status = AccountStatusActive
	}
	
	u.Version++
}

// VerifyPhone marks phone as verified
func (u *User) VerifyPhone() {
	if u.Phone != nil {
		now := time.Now().Unix()
		u.Phone.Verified = true
		u.Phone.VerifiedAt = &now
		u.PhoneVerified = true
		u.UpdatedAt = time.Now()
		u.Version++
	}
}

// VerifyIdentity marks identity as verified
func (u *User) VerifyIdentity() {
	u.IdentityVerified = true
	u.VerificationStatus = VerificationStatusVerified
	u.UpdatedAt = time.Now()
	u.Version++
}

// UpdateProfile updates profile information
func (u *User) UpdateProfile(bio, tagline, title string) error {
	if len(bio) > 5000 {
		return ErrInvalidBioLength
	}
	if len(tagline) > 200 {
		return ErrInvalidTaglineLength
	}
	
	u.Bio = strings.TrimSpace(bio)
	u.Tagline = strings.TrimSpace(tagline)
	u.Title = strings.TrimSpace(title)
	u.UpdatedAt = time.Now()
	u.Version++
	
	// Recalculate profile completeness
	u.CalculateProfileCompleteness()
	
	return nil
}

// UpdateLastSeen updates last seen timestamp
func (u *User) UpdateLastSeen() {
	now := time.Now()
	u.LastSeenAt = &now
	u.LastActiveAt = &now
	u.IsOnline = true
}

// SetOffline marks user as offline
func (u *User) SetOffline() {
	u.IsOnline = false
}

// RecordLogin records a successful login
func (u *User) RecordLogin(ipAddress, userAgent string) {
	now := time.Now()
	u.LastLoginAt = &now
	u.LastLoginIP = ipAddress
	u.LastUserAgent = userAgent
	u.LoginCount++
	u.LoginAttempts = 0 // Reset failed attempts on successful login
	u.IsOnline = true
	u.UpdatedAt = now
}

// IncrementFailedLogin increments failed login attempts
func (u *User) IncrementFailedLogin() {
	u.LoginAttempts++
	u.UpdatedAt = time.Now()
	
	// Lock account after 5 failed attempts for 15 minutes
	if u.LoginAttempts >= 5 {
		lockUntil := time.Now().Add(15 * time.Minute)
		u.LockedUntil = &lockUntil
	}
}

// ResetFailedLogins resets failed login attempts
func (u *User) ResetFailedLogins() {
	u.LoginAttempts = 0
	u.LockedUntil = nil
	u.UpdatedAt = time.Now()
}

// IsLocked checks if account is currently locked
func (u *User) IsLocked() bool {
	if u.LockedUntil == nil {
		return false
	}
	return time.Now().Before(*u.LockedUntil)
}

// AddWarning adds a warning to the user account
func (u *User) AddWarning(reason string, issuedBy string) {
	u.WarningCount++
	u.UpdatedAt = time.Now()
	u.LastModifiedBy = issuedBy
	
	if u.Notes == "" {
		u.Notes = fmt.Sprintf("Warning: %s", reason)
	} else {
		u.Notes += fmt.Sprintf("\n[%s] Warning: %s", time.Now().Format(time.RFC3339), reason)
	}
	
	u.Version++
}

// AssignBadge assigns a badge to the user
func (u *User) AssignBadge(badge BadgeType) error {
	if !badge.Valid() {
		return ErrInvalidBadgeType
	}
	
	// Check if badge already assigned
	for _, b := range u.Badges {
		if b == badge {
			return ErrBadgeAlreadyAssigned
		}
	}
	
	u.Badges = append(u.Badges, badge)
	u.UpdatedAt = time.Now()
	u.Version++
	
	// Update badge flags
	switch badge {
	case BadgeTypeTopRated, BadgeTypeTopRatedPlus:
		u.IsTopRated = true
	case BadgeTypeRisingTalent:
		u.IsRisingTalent = true
	case BadgeTypeExpertVetted:
		u.IsExpertVetted = true
	}
	
	return nil
}

// RemoveBadge removes a badge from the user
func (u *User) RemoveBadge(badge BadgeType) {
	newBadges := make([]BadgeType, 0)
	for _, b := range u.Badges {
		if b != badge {
			newBadges = append(newBadges, b)
		}
	}
	u.Badges = newBadges
	u.UpdatedAt = time.Now()
	u.Version++
	
	// Update badge flags
	switch badge {
	case BadgeTypeTopRated, BadgeTypeTopRatedPlus:
		u.IsTopRated = false
	case BadgeTypeRisingTalent:
		u.IsRisingTalent = false
	case BadgeTypeExpertVetted:
		u.IsExpertVetted = false
	}
}

// UpdateRating updates user's rating
func (u *User) UpdateRating(rating float64, totalReviews int) error {
	if rating < 0 || rating > 5 {
		return ErrInvalidRating
	}
	
	u.Rating = rating
	u.TotalReviews = totalReviews
	u.UpdatedAt = time.Now()
	u.Version++
	
	return nil
}

// UpdateStats updates user statistics
func (u *User) UpdateStats(completedJobs, totalJobs int, successRate float64) error {
	if successRate < 0 || successRate > 100 {
		return ErrInvalidCompletionRate
	}
	
	u.CompletedJobs = completedJobs
	u.TotalJobs = totalJobs
	u.SuccessRate = successRate
	u.UpdatedAt = time.Now()
	u.Version++
	
	return nil
}

// UpdateEarnings updates total earnings (for freelancers)
func (u *User) UpdateEarnings(amount float64) {
	u.TotalEarnings = amount
	u.UpdatedAt = time.Now()
	u.Version++
}

// UpdateSpending updates total spending (for clients)
func (u *User) UpdateSpending(amount float64) {
	u.TotalSpent = amount
	u.UpdatedAt = time.Now()
	u.Version++
}

// CalculateProfileCompleteness calculates profile completeness percentage
func (u *User) CalculateProfileCompleteness() int {
	completeness := 0
	
	// Basic info (20%)
	if u.FirstName != "" && u.LastName != "" && u.Email.Value != "" {
		completeness += 20
	}
	
	// Profile picture (10%)
	if u.ProfilePictureURL != "" {
		completeness += 10
	}
	
	// Bio/Overview (15%)
	if u.Bio != "" || u.Overview != "" {
		completeness += 15
	}
	
	// Tagline/Title (10%)
	if u.Tagline != "" || u.Title != "" {
		completeness += 10
	}
	
	// Location (10%)
	if u.Location != nil && u.Location.City != "" {
		completeness += 10
	}
	
	// Phone (5%)
	if u.Phone != nil && u.Phone.Number != "" {
		completeness += 5
	}
	
	// Email verified (10%)
	if u.EmailVerified {
		completeness += 10
	}
	
	// Identity verified (10%)
	if u.IdentityVerified {
		completeness += 10
	}
	
	// Social links (5%)
	if len(u.SocialLinks) > 0 {
		completeness += 5
	}
	
	// Video intro (5%)
	if u.VideoIntroURL != "" {
		completeness += 5
	}
	
	u.ProfileCompleteness = completeness
	u.ProfileCompleted = completeness >= 80
	
	return completeness
}

// SetFeatured marks user as featured
func (u *User) SetFeatured(featured bool) {
	u.IsFeatured = featured
	u.UpdatedAt = time.Now()
	u.Version++
}

// UpdateAvailability updates availability status
func (u *User) UpdateAvailability(status AvailabilityStatus, hoursPerWeek int) error {
	if !status.Valid() {
		return ErrInvalidAvailability
	}
	if hoursPerWeek < 0 || hoursPerWeek > 168 {
		return ErrHoursPerWeekInvalid
	}
	
	u.AvailabilityStatus = status
	u.HoursPerWeek = hoursPerWeek
	u.AcceptingWork = status.IsAvailable()
	u.UpdatedAt = time.Now()
	u.Version++
	
	return nil
}

// ============================================================================
// VALIDATION METHODS
// ============================================================================

// Validate validates the user entity
func (u *User) Validate() error {
	errors := NewValidationErrors()
	
	// Required fields
	if u.KeycloakID == "" {
		errors.Add("keycloak_id", "Keycloak ID is required", u.KeycloakID)
	}
	if u.Username == "" {
		errors.Add("username", "Username is required", u.Username)
	}
	if u.Email.Value == "" {
		errors.Add("email", "Email is required", u.Email.Value)
	}
	if u.FirstName == "" {
		errors.Add("first_name", "First name is required", u.FirstName)
	}
	if u.LastName == "" {
		errors.Add("last_name", "Last name is required", u.LastName)
	}
	
	// Validate enums
	if !u.UserType.Valid() {
		errors.Add("user_type", "Invalid user type", u.UserType)
	}
	if !u.Status.Valid() {
		errors.Add("status", "Invalid account status", u.Status)
	}
	if !u.VerificationStatus.Valid() {
		errors.Add("verification_status", "Invalid verification status", u.VerificationStatus)
	}
	
	// Validate email
	if err := u.Email.Validate(); err != nil {
		errors.Add("email", err.Error(), u.Email.Value)
	}
	
	// Validate username
	if err := validateUsername(u.Username); err != nil {
		errors.Add("username", err.Error(), u.Username)
	}
	
)

// ============================================================================
// TABLE NAME (GORM)
// ============================================================================

// TableName specifies the table name for GORM
func (User) TableName() string {
	return "users"
}