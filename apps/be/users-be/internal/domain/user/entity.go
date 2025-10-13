// internal/domain/user/entity.go
package user

import (
    "time"
    "gorm.io/gorm"
)

// User represents the core user aggregate root with full enterprise fields
type User struct {
    // Primary identifiers
    ID           string         `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
    KeycloakID   string         `gorm:"uniqueIndex;not null;type:varchar(100)"`
    Username     string         `gorm:"uniqueIndex;not null;type:varchar(100)"`
    Email        string         `gorm:"uniqueIndex;not null;type:varchar(255)"`
    
    // Personal information
    FirstName    string         `gorm:"not null;type:varchar(100)"`
    LastName     string         `gorm:"not null;type:varchar(100)"`
    MiddleName   string         `gorm:"type:varchar(100)"`
    DisplayName  string         `gorm:"type:varchar(200)"`
    
    // User classification
    UserType     UserType       `gorm:"not null;type:varchar(20);index"`
    Status       AccountStatus  `gorm:"not null;type:varchar(20);default:'active';index"`
    
    // Contact verification
    EmailVerified       bool      `gorm:"default:false;index"`
    EmailVerifiedAt     *time.Time
    PhoneNumber         string    `gorm:"type:varchar(20)"`
    PhoneCountryCode    string    `gorm:"type:varchar(5)"`
    PhoneVerified       bool      `gorm:"default:false"`
    PhoneVerifiedAt     *time.Time
    
    // Security & compliance
    TwoFactorEnabled    bool      `gorm:"default:false"`
    TwoFactorMethod     string    `gorm:"type:varchar(20)"` // sms, totp, email
    LastLoginAt         *time.Time
    LastLoginIP         string    `gorm:"type:varchar(45)"`
    LoginCount          int       `gorm:"default:0"`
    FailedLoginAttempts int       `gorm:"default:0"`
    LastFailedLoginAt   *time.Time
    PasswordChangedAt   *time.Time
    
    // Profile completion tracking
    ProfileCompleteness  int      `gorm:"default:0;check:profile_completeness >= 0 AND profile_completeness <= 100"`
    OnboardingCompleted  bool     `gorm:"default:false"`
    OnboardingStep       int      `gorm:"default:0"`
    
    // Geography & localization
    Country              string   `gorm:"type:varchar(2)"` // ISO 3166-1 alpha-2
    Timezone             string   `gorm:"type:varchar(50);default:'UTC'"`
    Language             string   `gorm:"type:varchar(10);default:'en'"`
    Currency             string   `gorm:"type:varchar(3);default:'USD'"` // ISO 4217
    
    // Platform metrics
    ReputationScore      float64  `gorm:"default:0;check:reputation_score >= 0 AND reputation_score <= 5"`
    TotalRatings         int      `gorm:"default:0"`
    AverageRating        float64  `gorm:"default:0"`
    ResponseRatePercent  float64  `gorm:"default:0;check:response_rate_percent >= 0 AND response_rate_percent <= 100"`
    
    // Activity tracking
    IsOnline             bool     `gorm:"default:false;index"`
    LastSeenAt           *time.Time
    TotalTimeOnline      int      `gorm:"default:0"` // in minutes
    
    // Business metrics
    TotalEarnings        float64  `gorm:"default:0;type:decimal(15,2)"`
    TotalSpent           float64  `gorm:"default:0;type:decimal(15,2)"`
    TotalJobs            int      `gorm:"default:0"`
    TotalHires           int      `gorm:"default:0"`
    CompletedJobs        int      `gorm:"default:0"`
    
    // Trust & safety
    IdentityVerified     bool     `gorm:"default:false;index"`
    IdentityVerifiedAt   *time.Time
    PaymentVerified      bool     `gorm:"default:false;index"`
    PaymentVerifiedAt    *time.Time
    BackgroundChecked    bool     `gorm:"default:false"`
    BackgroundCheckedAt  *time.Time
    
    // Flags & moderation
    IsFeatured           bool     `gorm:"default:false;index"`
    IsTopRated           bool     `gorm:"default:false;index"`
    IsPremium            bool     `gorm:"default:false;index"`
    PremiumUntil         *time.Time
    HasActiveWarnings    bool     `gorm:"default:false;index"`
    WarningCount         int      `gorm:"default:0"`
    SuspensionCount      int      `gorm:"default:0"`
    
    // Referral system
    ReferralCode         string   `gorm:"uniqueIndex;type:varchar(20)"`
    ReferredBy           *string  `gorm:"type:uuid"`
    TotalReferrals       int      `gorm:"default:0"`
    ReferralEarnings     float64  `gorm:"default:0;type:decimal(10,2)"`
    
    // Marketing & communications
    MarketingEmailsEnabled    bool `gorm:"default:true"`
    ProductUpdatesEnabled     bool `gorm:"default:true"`
    NewsletterSubscribed      bool `gorm:"default:false"`
    SmsNotificationsEnabled   bool `gorm:"default:false"`
    PushNotificationsEnabled  bool `gorm:"default:true"`
    
    // Admin notes & internal tracking
    InternalNotes        string   `gorm:"type:text"`
    Tags                 string   `gorm:"type:text"` // JSON array of tags
    Metadata             string   `gorm:"type:jsonb"` // Additional flexible metadata
    
    // Soft delete & audit
    CreatedAt    time.Time
    UpdatedAt    time.Time
    DeletedAt    gorm.DeletedAt `gorm:"index"`
    CreatedBy    *string        `gorm:"type:uuid"` // Admin who created (if admin created)
    UpdatedBy    *string        `gorm:"type:uuid"` // Last admin who updated
    DeletedBy    *string        `gorm:"type:uuid"` // Admin who soft-deleted
}

func (User) TableName() string {
    return "users"
}

// UserType represents the type of user account
type UserType string

const (
    UserTypeFreelancer UserType = "freelancer"
    UserTypeClient     UserType = "client"
    UserTypeBoth       UserType = "both" // Can act as both freelancer and client
)

// AccountStatus represents the current status of the user account
type AccountStatus string

const (
    AccountStatusPending   AccountStatus = "pending"   // Email not verified yet
    AccountStatusActive    AccountStatus = "active"    // Normal active account
    AccountStatusSuspended AccountStatus = "suspended" // Temporarily suspended
    AccountStatusBanned    AccountStatus = "banned"    // Permanently banned
    AccountStatusDeleted   AccountStatus = "deleted"   // User-initiated deletion
    AccountStatusInactive  AccountStatus = "inactive"  // Inactive due to inactivity
)

// Business methods for User entity

func (u *User) IsActive() bool {
    return u.Status == AccountStatusActive && u.DeletedAt.Time.IsZero()
}

func (u *User) CanLogin() bool {
    return u.IsActive() && u.EmailVerified
}

func (u *User) IsFreelancer() bool {
    return u.UserType == UserTypeFreelancer || u.UserType == UserTypeBoth
}

func (u *User) IsClient() bool {
    return u.UserType == UserTypeClient || u.UserType == UserTypeBoth
}

func (u *User) IsVerified() bool {
    return u.EmailVerified && u.IdentityVerified && u.PaymentVerified
}

func (u *User) IsTrusted() bool {
    return u.IsVerified() && u.ReputationScore >= 4.0 && !u.HasActiveWarnings
}

func (u *User) CanReceivePayments() bool {
    return u.IsActive() && u.PaymentVerified && u.IsFreelancer()
}

func (u *User) CanHire() bool {
    return u.IsActive() && u.PaymentVerified && u.IsClient()
}

func (u *User) UpdateLoginInfo(ip string) {
    now := time.Now()
    u.LastLoginAt = &now
    u.LastLoginIP = ip
    u.LoginCount++
    u.FailedLoginAttempts = 0
}

func (u *User) RecordFailedLogin() {
    now := time.Now()
    u.FailedLoginAttempts++
    u.LastFailedLoginAt = &now
}

func (u *User) ShouldLockAccount() bool {
    return u.FailedLoginAttempts >= 5
}

func (u *User) CalculateProfileCompleteness(hasProfile, hasSkills, hasExperience, hasEducation, hasPortfolio bool) {
    completeness := 20 // Base for account creation
    
    if u.EmailVerified {
        completeness += 10
    }
    if u.PhoneVerified {
        completeness += 5
    }
    if u.IdentityVerified {
        completeness += 10
    }
    if u.PaymentVerified {
        completeness += 10
    }
    if hasProfile {
        completeness += 15
    }
    if hasSkills {
        completeness += 10
    }
    if hasExperience {
        completeness += 10
    }
    if hasEducation {
        completeness += 5
    }
    if hasPortfolio {
        completeness += 5
    }
    
    u.ProfileCompleteness = completeness
}