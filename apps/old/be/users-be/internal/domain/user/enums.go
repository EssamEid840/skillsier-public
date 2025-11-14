// internal/domain/user/enums.go
package user

import "fmt"

// ============================================================================
// USER TYPE ENUM
// ============================================================================

// UserType defines the type of user (Freelancer, Client, or both)
type UserType string

const (
	UserTypeFreelancer UserType = "FREELANCER"
	UserTypeClient     UserType = "CLIENT"
	UserTypeBoth       UserType = "BOTH"
	UserTypeAdmin      UserType = "ADMIN"      // Platform administrators
	UserTypeModerator  UserType = "MODERATOR"  // Content moderators
	UserTypeSupport    UserType = "SUPPORT"    // Support staff
)

// Valid returns true if the user type is valid
func (ut UserType) Valid() bool {
	switch ut {
	case UserTypeFreelancer, UserTypeClient, UserTypeBoth, UserTypeAdmin, UserTypeModerator, UserTypeSupport:
		return true
	default:
		return false
	}
}

// String returns string representation
func (ut UserType) String() string {
	return string(ut)
}

// IsFreelancer checks if user is a freelancer or both
func (ut UserType) IsFreelancer() bool {
	return ut == UserTypeFreelancer || ut == UserTypeBoth
}

// IsClient checks if user is a client or both
func (ut UserType) IsClient() bool {
	return ut == UserTypeClient || ut == UserTypeBoth
}

// IsStaff checks if user is platform staff
func (ut UserType) IsStaff() bool {
	return ut == UserTypeAdmin || ut == UserTypeModerator || ut == UserTypeSupport
}

// ============================================================================
// ACCOUNT STATUS ENUM
// ============================================================================

// AccountStatus defines the current status of a user account
type AccountStatus string

const (
	AccountStatusPending   AccountStatus = "PENDING"    // Email not verified
	AccountStatusActive    AccountStatus = "ACTIVE"     // Active and verified
	AccountStatusInactive  AccountStatus = "INACTIVE"   // User deactivated account
	AccountStatusSuspended AccountStatus = "SUSPENDED"  // Temporarily suspended
	AccountStatusBanned    AccountStatus = "BANNED"     // Permanently banned
	AccountStatusDeleted   AccountStatus = "DELETED"    // Soft deleted
	AccountStatusRestricted AccountStatus = "RESTRICTED" // Limited functionality
)

// Valid returns true if the account status is valid
func (as AccountStatus) Valid() bool {
	switch as {
	case AccountStatusPending, AccountStatusActive, AccountStatusInactive,
		AccountStatusSuspended, AccountStatusBanned, AccountStatusDeleted, AccountStatusRestricted:
		return true
	default:
		return false
	}
}

// String returns string representation
func (as AccountStatus) String() string {
	return string(as)
}

// CanLogin checks if user can login with this status
func (as AccountStatus) CanLogin() bool {
	return as == AccountStatusActive || as == AccountStatusRestricted
}

// CanPerformActions checks if user can perform platform actions
func (as AccountStatus) CanPerformActions() bool {
	return as == AccountStatusActive
}

// IsBlocked checks if account is blocked (suspended, banned, or deleted)
func (as AccountStatus) IsBlocked() bool {
	return as == AccountStatusSuspended || as == AccountStatusBanned || as == AccountStatusDeleted
}

// ============================================================================
// VERIFICATION STATUS ENUM
// ============================================================================

// VerificationStatus defines the verification status of a user
type VerificationStatus string

const (
	VerificationStatusUnverified      VerificationStatus = "UNVERIFIED"       // No verification submitted
	VerificationStatusPending         VerificationStatus = "PENDING"          // Verification in review
	VerificationStatusVerified        VerificationStatus = "VERIFIED"         // Identity verified
	VerificationStatusRejected        VerificationStatus = "REJECTED"         // Verification rejected
	VerificationStatusExpired         VerificationStatus = "EXPIRED"          // Verification expired
	VerificationStatusRequireResubmit VerificationStatus = "REQUIRE_RESUBMIT" // Need to resubmit docs
)

// Valid returns true if the verification status is valid
func (vs VerificationStatus) Valid() bool {
	switch vs {
	case VerificationStatusUnverified, VerificationStatusPending, VerificationStatusVerified,
		VerificationStatusRejected, VerificationStatusExpired, VerificationStatusRequireResubmit:
		return true
	default:
		return false
	}
}

// String returns string representation
func (vs VerificationStatus) String() string {
	return string(vs)
}

// IsVerified checks if user is verified
func (vs VerificationStatus) IsVerified() bool {
	return vs == VerificationStatusVerified
}

// NeedsVerification checks if user needs to submit verification
func (vs VerificationStatus) NeedsVerification() bool {
	return vs == VerificationStatusUnverified || vs == VerificationStatusRequireResubmit
}

// ============================================================================
// PROFILE VISIBILITY ENUM
// ============================================================================

// ProfileVisibility defines who can see the user's profile
type ProfileVisibility string

const (
	ProfileVisibilityPublic    ProfileVisibility = "PUBLIC"     // Everyone can see
	ProfileVisibilityPrivate   ProfileVisibility = "PRIVATE"    // Only connections
	ProfileVisibilityAnonymous ProfileVisibility = "ANONYMOUS"  // Hidden from search
	ProfileVisibilityRestricted ProfileVisibility = "RESTRICTED" // Limited info visible
)

// Valid returns true if the profile visibility is valid
func (pv ProfileVisibility) Valid() bool {
	switch pv {
	case ProfileVisibilityPublic, ProfileVisibilityPrivate, ProfileVisibilityAnonymous, ProfileVisibilityRestricted:
		return true
	default:
		return false
	}
}

// String returns string representation
func (pv ProfileVisibility) String() string {
	return string(pv)
}

// IsPublic checks if profile is publicly visible
func (pv ProfileVisibility) IsPublic() bool {
	return pv == ProfileVisibilityPublic
}

// ============================================================================
// GENDER ENUM
// ============================================================================

// Gender represents user's gender
type Gender string

const (
	GenderMale              Gender = "MALE"
	GenderFemale            Gender = "FEMALE"
	GenderNonBinary         Gender = "NON_BINARY"
	GenderPreferNotToSay    Gender = "PREFER_NOT_TO_SAY"
	GenderOther             Gender = "OTHER"
)

// Valid returns true if the gender is valid
func (g Gender) Valid() bool {
	switch g {
	case GenderMale, GenderFemale, GenderNonBinary, GenderPreferNotToSay, GenderOther:
		return true
	default:
		return false
	}
}

// String returns string representation
func (g Gender) String() string {
	return string(g)
}

// ============================================================================
// LANGUAGE PROFICIENCY ENUM
// ============================================================================

// LanguageProficiency defines proficiency level for a language
type LanguageProficiency string

const (
	LanguageProficiencyBasic        LanguageProficiency = "BASIC"
	LanguageProficiencyConversational LanguageProficiency = "CONVERSATIONAL"
	LanguageProficiencyFluent       LanguageProficiency = "FLUENT"
	LanguageProficiencyNative       LanguageProficiency = "NATIVE"
)

// Valid returns true if the language proficiency is valid
func (lp LanguageProficiency) Valid() bool {
	switch lp {
	case LanguageProficiencyBasic, LanguageProficiencyConversational, 
		LanguageProficiencyFluent, LanguageProficiencyNative:
		return true
	default:
		return false
	}
}

// String returns string representation
func (lp LanguageProficiency) String() string {
	return string(lp)
}

// Level returns numeric level (1-4) for proficiency
func (lp LanguageProficiency) Level() int {
	switch lp {
	case LanguageProficiencyBasic:
		return 1
	case LanguageProficiencyConversational:
		return 2
	case LanguageProficiencyFluent:
		return 3
	case LanguageProficiencyNative:
		return 4
	default:
		return 0
	}
}

// ============================================================================
// AVAILABILITY STATUS ENUM
// ============================================================================

// AvailabilityStatus defines user's current availability for work
type AvailabilityStatus string

const (
	AvailabilityStatusAvailable      AvailabilityStatus = "AVAILABLE"       // Available for new work
	AvailabilityStatusBusy           AvailabilityStatus = "BUSY"            // Currently working
	AvailabilityStatusUnavailable    AvailabilityStatus = "UNAVAILABLE"     // Not available
	AvailabilityStatusPartTime       AvailabilityStatus = "PART_TIME"       // Available part-time
	AvailabilityStatusFullTime       AvailabilityStatus = "FULL_TIME"       // Available full-time
	AvailabilityStatusOnLeave        AvailabilityStatus = "ON_LEAVE"        // On temporary leave
)

// Valid returns true if the availability status is valid
func (avs AvailabilityStatus) Valid() bool {
	switch avs {
	case AvailabilityStatusAvailable, AvailabilityStatusBusy, AvailabilityStatusUnavailable,
		AvailabilityStatusPartTime, AvailabilityStatusFullTime, AvailabilityStatusOnLeave:
		return true
	default:
		return false
	}
}

// String returns string representation
func (avs AvailabilityStatus) String() string {
	return string(avs)
}

// IsAvailable checks if user is available for work
func (avs AvailabilityStatus) IsAvailable() bool {
	return avs == AvailabilityStatusAvailable || 
		   avs == AvailabilityStatusPartTime || 
		   avs == AvailabilityStatusFullTime
}

// ============================================================================
// BADGE TYPE ENUM (for achievements and recognition)
// ============================================================================

// BadgeType defines types of badges users can earn
type BadgeType string

const (
	BadgeTypeTopRated         BadgeType = "TOP_RATED"
	BadgeTypeRisingTalent     BadgeType = "RISING_TALENT"
	BadgeTypeExpertVetted     BadgeType = "EXPERT_VETTED"
	BadgeTypeTopRatedPlus     BadgeType = "TOP_RATED_PLUS"
	BadgeTypeSpecialized      BadgeType = "SPECIALIZED"
	BadgeTypeClientFavorite   BadgeType = "CLIENT_FAVORITE"
	BadgeTypeReliable         BadgeType = "RELIABLE"
	BadgeTypeCommunicator     BadgeType = "COMMUNICATOR"
	BadgeTypeOnTime           BadgeType = "ON_TIME"
	BadgeTypeHighQuality      BadgeType = "HIGH_QUALITY"
)

// Valid returns true if the badge type is valid
func (bt BadgeType) Valid() bool {
	switch bt {
	case BadgeTypeTopRated, BadgeTypeRisingTalent, BadgeTypeExpertVetted,
		BadgeTypeTopRatedPlus, BadgeTypeSpecialized, BadgeTypeClientFavorite,
		BadgeTypeReliable, BadgeTypeCommunicator, BadgeTypeOnTime, BadgeTypeHighQuality:
		return true
	default:
		return false
	}
}

// String returns string representation
func (bt BadgeType) String() string {
	return string(bt)
}

// ============================================================================
// HELPER FUNCTIONS
// ============================================================================

// ParseUserType parses string to UserType
func ParseUserType(s string) (UserType, error) {
	ut := UserType(s)
	if !ut.Valid() {
		return "", fmt.Errorf("invalid user type: %s", s)
	}
	return ut, nil
}

// ParseAccountStatus parses string to AccountStatus
func ParseAccountStatus(s string) (AccountStatus, error) {
	as := AccountStatus(s)
	if !as.Valid() {
		return "", fmt.Errorf("invalid account status: %s", s)
	}
	return as, nil
}

// ParseVerificationStatus parses string to VerificationStatus
func ParseVerificationStatus(s string) (VerificationStatus, error) {
	vs := VerificationStatus(s)
	if !vs.Valid() {
		return "", fmt.Errorf("invalid verification status: %s", s)
	}
	return vs, nil
}

// ParseProfileVisibility parses string to ProfileVisibility
func ParseProfileVisibility(s string) (ProfileVisibility, error) {
	pv := ProfileVisibility(s)
	if !pv.Valid() {
		return "", fmt.Errorf("invalid profile visibility: %s", s)
	}
	return pv, nil
}