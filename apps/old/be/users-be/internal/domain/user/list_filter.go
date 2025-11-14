// internal/domain/user/list_filter.go
package user

import (
	"fmt"
	"time"
)

// ListFilter defines filtering options for list queries
type ListFilter struct {
	// Pagination
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
	Page   int `json:"page"`
	
	// Sorting
	SortBy    string `json:"sort_by"`
	SortOrder string `json:"sort_order"`
	
	// Filters - User Type & Status
	UserType           *UserType          `json:"user_type,omitempty"`
	Status             *AccountStatus     `json:"status,omitempty"`
	VerificationStatus *VerificationStatus `json:"verification_status,omitempty"`
	
	// Filters - Location
	Country     string `json:"country,omitempty"`
	CountryCode string `json:"country_code,omitempty"`
	City        string `json:"city,omitempty"`
	Timezone    string `json:"timezone,omitempty"`
	
	// Filters - Ratings & Performance
	MinRating          float64 `json:"min_rating,omitempty"`
	MaxRating          float64 `json:"max_rating,omitempty"`
	MinCompletedJobs   int     `json:"min_completed_jobs,omitempty"`
	MinSuccessRate     float64 `json:"min_success_rate,omitempty"`
	MinReviews         int     `json:"min_reviews,omitempty"`
	
	// Filters - Verification
	EmailVerified    *bool `json:"email_verified,omitempty"`
	PhoneVerified    *bool `json:"phone_verified,omitempty"`
	IdentityVerified *bool `json:"identity_verified,omitempty"`
	
	// Filters - Badges & Features
	IsFeatured     *bool      `json:"is_featured,omitempty"`
	IsTopRated     *bool      `json:"is_top_rated,omitempty"`
	IsRisingTalent *bool      `json:"is_rising_talent,omitempty"`
	IsExpertVetted *bool      `json:"is_expert_vetted,omitempty"`
	HasBadge       *BadgeType `json:"has_badge,omitempty"`
	
	// Filters - Activity
	IsOnline       *bool  `json:"is_online,omitempty"`
	SearchableOnly bool   `json:"searchable_only,omitempty"`
	AcceptingWork  *bool  `json:"accepting_work,omitempty"`
	AvailabilityStatus *AvailabilityStatus `json:"availability_status,omitempty"`
	
	// Filters - Profile Completeness
	MinProfileCompleteness int  `json:"min_profile_completeness,omitempty"`
	ProfileCompleted       *bool `json:"profile_completed,omitempty"`
	
	// Date Filters
	CreatedAfter  *time.Time `json:"created_after,omitempty"`
	CreatedBefore *time.Time `json:"created_before,omitempty"`
	UpdatedAfter  *time.Time `json:"updated_after,omitempty"`
	UpdatedBefore *time.Time `json:"updated_before,omitempty"`
	LastLoginAfter  *time.Time `json:"last_login_after,omitempty"`
	LastLoginBefore *time.Time `json:"last_login_before,omitempty"`
	
	// Special Filters
	ReferredBy      string `json:"referred_by,omitempty"`
	HasWarnings     bool   `json:"has_warnings,omitempty"`
	IsSuspended     bool   `json:"is_suspended,omitempty"`
	IsBanned        bool   `json:"is_banned,omitempty"`
	IncludeDeleted  bool   `json:"include_deleted,omitempty"`
	
	// Search
	SearchQuery string   `json:"search_query,omitempty"`
	SearchFields []string `json:"search_fields,omitempty"`
}

// NewListFilter creates a new ListFilter with defaults
func NewListFilter() ListFilter {
	return ListFilter{
		Limit:          20,
		Offset:         0,
		Page:           1,
		SortBy:         "created_at",
		SortOrder:      "desc",
		SearchableOnly: false,
		IncludeDeleted: false,
		SearchFields:   []string{"username", "full_name", "email", "bio", "tagline"},
	}
}

// Validate validates the filter parameters
func (f *ListFilter) Validate() error {
	// Validate limit
	if f.Limit < 0 {
		return fmt.Errorf("limit cannot be negative")
	}
	if f.Limit > 100 {
		return fmt.Errorf("limit cannot exceed 100")
	}
	if f.Limit == 0 {
		f.Limit = 20
	}
	
	// Validate offset
	if f.Offset < 0 {
		return fmt.Errorf("offset cannot be negative")
	}
	
	// Validate page
	if f.Page < 1 {
		f.Page = 1
	}
	
	// Calculate offset from page if needed
	if f.Offset == 0 && f.Page > 1 {
		f.Offset = (f.Page - 1) * f.Limit
	}
	
	// Validate sort order
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
		"last_seen_at": true, "profile_completeness": true, "full_name": true,
		"total_earnings": true, "total_spent": true,
	}
	if f.SortBy != "" && !validSortFields[f.SortBy] {
		return ErrInvalidSortField
	}
	if f.SortBy == "" {
		f.SortBy = "created_at"
	}
	
	// Validate ratings
	if f.MinRating < 0 || f.MinRating > 5 {
		return fmt.Errorf("min_rating must be between 0 and 5")
	}
	if f.MaxRating < 0 || f.MaxRating > 5 {
		return fmt.Errorf("max_rating must be between 0 and 5")
	}
	if f.MinRating > 0 && f.MaxRating > 0 && f.MinRating > f.MaxRating {
		return fmt.Errorf("min_rating cannot be greater than max_rating")
	}
	
	// Validate success rate
	if f.MinSuccessRate < 0 || f.MinSuccessRate > 100 {
		return fmt.Errorf("min_success_rate must be between 0 and 100")
	}
	
	// Validate profile completeness
	if f.MinProfileCompleteness < 0 || f.MinProfileCompleteness > 100 {
		return fmt.Errorf("min_profile_completeness must be between 0 and 100")
	}
	
	// Validate date ranges
	if f.CreatedAfter != nil && f.CreatedBefore != nil && f.CreatedAfter.After(*f.CreatedBefore) {
		return fmt.Errorf("created_after cannot be after created_before")
	}
	if f.UpdatedAfter != nil && f.UpdatedBefore != nil && f.UpdatedAfter.After(*f.UpdatedBefore) {
		return fmt.Errorf("updated_after cannot be after updated_before")
	}
	if f.LastLoginAfter != nil && f.LastLoginBefore != nil && f.LastLoginAfter.After(*f.LastLoginBefore) {
		return fmt.Errorf("last_login_after cannot be after last_login_before")
	}
	
	return nil
}

// ApplyDefaults applies default values to empty fields
func (f *ListFilter) ApplyDefaults() {
	if f.Limit == 0 {
		f.Limit = 20
	}
	if f.Page == 0 {
		f.Page = 1
	}
	if f.SortBy == "" {
		f.SortBy = "created_at"
	}
	if f.SortOrder == "" {
		f.SortOrder = "desc"
	}
	if f.SearchFields == nil || len(f.SearchFields) == 0 {
		f.SearchFields = []string{"username", "full_name", "email", "bio", "tagline"}
	}
}

// HasFilters checks if any filters are applied
func (f *ListFilter) HasFilters() bool {
	return f.UserType != nil ||
		f.Status != nil ||
		f.VerificationStatus != nil ||
		f.Country != "" ||
		f.City != "" ||
		f.MinRating > 0 ||
		f.MaxRating > 0 ||
		f.MinCompletedJobs > 0 ||
		f.MinSuccessRate > 0 ||
		f.EmailVerified != nil ||
		f.IdentityVerified != nil ||
		f.IsFeatured != nil ||
		f.IsTopRated != nil ||
		f.IsOnline != nil ||
		f.SearchQuery != "" ||
		f.CreatedAfter != nil ||
		f.CreatedBefore != nil
}

// GetOffset returns the calculated offset
func (f *ListFilter) GetOffset() int {
	if f.Offset > 0 {
		return f.Offset
	}
	if f.Page > 1 {
		return (f.Page - 1) * f.Limit
	}
	return 0
}

// GetLimit returns the limit
func (f *ListFilter) GetLimit() int {
	if f.Limit <= 0 {
		return 20
	}
	if f.Limit > 100 {
		return 100
	}
	return f.Limit
}

// ToMap converts filter to map for logging/debugging
func (f *ListFilter) ToMap() map[string]interface{} {
	m := make(map[string]interface{})
	
	m["limit"] = f.Limit
	m["offset"] = f.Offset
	m["page"] = f.Page
	m["sort_by"] = f.SortBy
	m["sort_order"] = f.SortOrder
	
	if f.UserType != nil {
		m["user_type"] = *f.UserType
	}
	if f.Status != nil {
		m["status"] = *f.Status
	}
	if f.Country != "" {
		m["country"] = f.Country
	}
	if f.City != "" {
		m["city"] = f.City
	}
	if f.MinRating > 0 {
		m["min_rating"] = f.MinRating
	}
	if f.SearchQuery != "" {
		m["search_query"] = f.SearchQuery
	}
	
	return m
}