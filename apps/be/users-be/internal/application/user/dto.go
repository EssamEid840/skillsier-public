// apps/be/users-be/internal/application/user/dto.go
package user

import (
	"time"
	
	"users-be/internal/domain/user"
)

// ============================================================================
// CREATE USER DTOs
// ============================================================================

type CreateUserDTO struct {
	KeycloakID       string          `json:"keycloak_id"`
	Username         string          `json:"username"`
	Email            string          `json:"email"`
	FirstName        *string         `json:"first_name,omitempty"`
	LastName         *string         `json:"last_name,omitempty"`
	UserType         user.UserType   `json:"user_type"`
	PhoneNumber      *string         `json:"phone_number,omitempty"`
	PhoneCountryCode *string         `json:"phone_country_code,omitempty"`
	City             *string         `json:"city,omitempty"`
	Country          *string         `json:"country,omitempty"`
	CountryCode      *string         `json:"country_code,omitempty"`
	Timezone         *string         `json:"timezone,omitempty"`
	Bio              *string         `json:"bio,omitempty"`
	Tagline          *string         `json:"tagline,omitempty"`
}

// ============================================================================
// UPDATE USER DTOs
// ============================================================================

type UpdateUserDTO struct {
	FirstName         *string `json:"first_name,omitempty"`
	LastName          *string `json:"last_name,omitempty"`
	DisplayName       *string `json:"display_name,omitempty"`
	Bio               *string `json:"bio,omitempty"`
	Tagline           *string `json:"tagline,omitempty"`
	Title             *string `json:"title,omitempty"`
	Website           *string `json:"website,omitempty"`
	ProfilePictureURL *string `json:"profile_picture_url,omitempty"`
	CoverImageURL     *string `json:"cover_image_url,omitempty"`
	PhoneNumber       *string `json:"phone_number,omitempty"`
	PhoneCountryCode  *string `json:"phone_country_code,omitempty"`
	City              *string `json:"city,omitempty"`
	Country           *string `json:"country,omitempty"`
	CountryCode       *string `json:"country_code,omitempty"`
	Timezone          *string `json:"timezone,omitempty"`
}

type UpdateAvailabilityDTO struct {
	Status       *user.AvailabilityStatus `json:"status,omitempty"`
	HoursPerWeek *int                     `json:"hours_per_week,omitempty"`
}

type UpdateSettingsDTO struct {
	ProfileVisibility *user.ProfileVisibility `json:"profile_visibility,omitempty"`
	ShowEmail         *bool                   `json:"show_email,omitempty"`
	ShowPhone         *bool                   `json:"show_phone,omitempty"`
	ShowLocation      *bool                   `json:"show_location,omitempty"`
	SearchableProfile *bool                   `json:"searchable_profile,omitempty"`
	AcceptingWork     *bool                   `json:"accepting_work,omitempty"`
}

// ============================================================================
// RESPONSE DTOs
// ============================================================================

type UserDTO struct {
	ID                   string                   `json:"id"`
	KeycloakID           string                   `json:"keycloak_id"`
	Username             string                   `json:"username"`
	Email                string                   `json:"email"`
	EmailVerified        bool                     `json:"email_verified"`
	FirstName            *string                  `json:"first_name,omitempty"`
	LastName             *string                  `json:"last_name,omitempty"`
	DisplayName          *string                  `json:"display_name,omitempty"`
	FullName             string                   `json:"full_name"`
	UserType             string                   `json:"user_type"`
	Bio                  *string                  `json:"bio,omitempty"`
	Tagline              *string                  `json:"tagline,omitempty"`
	Title                *string                  `json:"title,omitempty"`
	Website              *string                  `json:"website,omitempty"`
	ProfilePictureURL    *string                  `json:"profile_picture_url,omitempty"`
	CoverImageURL        *string                  `json:"cover_image_url,omitempty"`
	ProfileCompleteness  int                      `json:"profile_completeness"`
	PhoneNumber          *string                  `json:"phone_number,omitempty"`
	PhoneCountryCode     *string                  `json:"phone_country_code,omitempty"`
	PhoneVerified        bool                     `json:"phone_verified"`
	City                 *string                  `json:"city,omitempty"`
	Country              *string                  `json:"country,omitempty"`
	CountryCode          *string                  `json:"country_code,omitempty"`
	Timezone             *string                  `json:"timezone,omitempty"`
	Status               string                   `json:"status"`
	VerificationStatus   string                   `json:"verification_status"`
	IsOnline             bool                     `json:"is_online"`
	LastSeenAt           *time.Time               `json:"last_seen_at,omitempty"`
	LastActiveAt         *time.Time               `json:"last_active_at,omitempty"`
	AvailabilityStatus   string                   `json:"availability_status"`
	HoursPerWeek         *int                     `json:"hours_per_week,omitempty"`
	ResponseTime         *int                     `json:"response_time,omitempty"`
	TotalEarned          float64                  `json:"total_earned"`
	TotalSpent           float64                  `json:"total_spent"`
	JobSuccess           *float64                 `json:"job_success,omitempty"`
	TotalJobs            int                      `json:"total_jobs"`
	TotalHires           int                      `json:"total_hires"`
	AvgRating            *float64                 `json:"avg_rating,omitempty"`
	ReviewCount          int                      `json:"review_count"`
	ConnectsBalance      int                      `json:"connects_balance"`
	ProfileVisibility    string                   `json:"profile_visibility"`
	ShowEmail            bool                     `json:"show_email"`
	ShowPhone            bool                     `json:"show_phone"`
	ShowLocation         bool                     `json:"show_location"`
	SearchableProfile    bool                     `json:"searchable_profile"`
	AcceptingWork        bool                     `json:"accepting_work"`
	Badges               []string                 `json:"badges,omitempty"`
	Warnings             int                      `json:"warnings"`
	ReferralCode         *string                  `json:"referral_code,omitempty"`
	ReferredBy           *string                  `json:"referred_by,omitempty"`
	ReferralCount        int                      `json:"referral_count"`
	IsFeatured           bool                     `json:"is_featured"`
	IsTopRated           bool                     `json:"is_top_rated"`
	IsRisingTalent       bool                     `json:"is_rising_talent"`
	IsExpertVetted       bool                     `json:"is_expert_vetted"`
	LastLoginAt          *time.Time               `json:"last_login_at,omitempty"`
	LastLoginIP          *string                  `json:"last_login_ip,omitempty"`
	LoginCount           int                      `json:"login_count"`
	CreatedAt            time.Time                `json:"created_at"`
	UpdatedAt            time.Time                `json:"updated_at"`
}

type UserListDTO struct {
	Users      []UserDTO `json:"users"`
	Total      int64     `json:"total"`
	Page       int       `json:"page"`
	PageSize   int       `json:"page_size"`
	TotalPages int       `json:"total_pages"`
}

type UserSearchResultDTO struct {
	Users       []UserDTO `json:"users"`
	Total       int64     `json:"total"`
	Query       string    `json:"query"`
	SearchTime  int64     `json:"search_time_ms"`
	Aggregations map[string]interface{} `json:"aggregations,omitempty"`
}

type UserStatisticsDTO struct {
	TotalUsers            int64            `json:"total_users"`
	TotalFreelancers      int64            `json:"total_freelancers"`
	TotalClients          int64            `json:"total_clients"`
	ActiveUsers           int64            `json:"active_users"`
	SuspendedUsers        int64            `json:"suspended_users"`
	BannedUsers           int64            `json:"banned_users"`
	VerifiedUsers         int64            `json:"verified_users"`
	UnverifiedUsers       int64            `json:"unverified_users"`
	OnlineUsers           int64            `json:"online_users"`
	NewUsersLast30Days    int64            `json:"new_users_last_30_days"`
	UsersByCountry        map[string]int64 `json:"users_by_country"`
	UsersByUserType       map[string]int64 `json:"users_by_user_type"`
	UsersByStatus         map[string]int64 `json:"users_by_status"`
	AverageCompleteness   float64          `json:"average_completeness"`
	AverageResponseTime   int              `json:"average_response_time"`
}

type UserCountDTO struct {
	Count int64 `json:"count"`
}

type BulkActionResultDTO struct {
	Success      []string `json:"success"`
	Failed       []string `json:"failed"`
	TotalSuccess int      `json:"total_success"`
	TotalFailed  int      `json:"total_failed"`
	Errors       []string `json:"errors,omitempty"`
}

// ============================================================================
// FILTER DTOs
// ============================================================================

type UserFilterDTO struct {
	UserTypes            []string `json:"user_types,omitempty"`
	Statuses             []string `json:"statuses,omitempty"`
	VerificationStatuses []string `json:"verification_statuses,omitempty"`
	Countries            []string `json:"countries,omitempty"`
	Cities               []string `json:"cities,omitempty"`
	MinCompleteness      *int     `json:"min_completeness,omitempty"`
	MaxCompleteness      *int     `json:"max_completeness,omitempty"`
	MinRating            *float64 `json:"min_rating,omitempty"`
	MaxRating            *float64 `json:"max_rating,omitempty"`
	HasBadges            []string `json:"has_badges,omitempty"`
	IsOnline             *bool    `json:"is_online,omitempty"`
	IsFeatured           *bool    `json:"is_featured,omitempty"`
	IsTopRated           *bool    `json:"is_top_rated,omitempty"`
	IsRisingTalent       *bool    `json:"is_rising_talent,omitempty"`
	IsExpertVetted       *bool    `json:"is_expert_vetted,omitempty"`
	CreatedAfter         *time.Time `json:"created_after,omitempty"`
	CreatedBefore        *time.Time `json:"created_before,omitempty"`
	LastSeenAfter        *time.Time `json:"last_seen_after,omitempty"`
	LastSeenBefore       *time.Time `json:"last_seen_before,omitempty"`
	SortBy               string   `json:"sort_by,omitempty"`
	SortOrder            string   `json:"sort_order,omitempty"`
	Page                 int      `json:"page,omitempty"`
	PageSize             int      `json:"page_size,omitempty"`
}


// Add these DTO conversion methods to your service layer

// ToDTO converts user entity to DTO
func ToDTO(u *user.User) *UserDTO {
	if u == nil {
		return nil
	}
	
	return &UserDTO{
		ID:          u.ID,
		KeycloakID:  u.KeycloakID,
		Username:    u.Username,
		Email:       u.Email.Value,
		FirstName:   u.FirstName,
		LastName:    u.LastName,
		FullName:    u.FullName,
		DisplayName: u.DisplayName,
		UserType:    string(u.UserType),
		Status:      string(u.Status),
		// Add other fields as needed
	}
}

// ToDTOList converts user entities to DTO list
func ToDTOList(users []*user.User) []UserDTO {
	result := make([]UserDTO, len(users))
	for i, u := range users {
		result[i] = *ToDTO(u)
	}
	return result
}

// ToListDTO converts user list to list DTO
func ToListDTO(users []*user.User, total int64, page, pageSize int) *UserListDTO {
	return &UserListDTO{
		Users:    ToDTOList(users),
		Total:    total,
		Page:     page,
		PageSize: pageSize,
	}
}

// ToSearchResultDTO converts search results to DTO
func ToSearchResultDTO(users []*user.User, total int64, query string, searchTime int64) *UserSearchResultDTO {
	return &UserSearchResultDTO{
		Users:      ToDTOList(users),
		Total:      total,
		Query:      query,
		SearchTime: searchTime,
	}
}

// ToStatisticsDTO converts statistics to DTO
func ToStatisticsDTO(stats *user.UserStatistics) *UserStatisticsDTO {
	if stats == nil {
		return nil
	}
	
	return &UserStatisticsDTO{
		TotalUsers:           stats.TotalUsers,
		ActiveUsers:          stats.ActiveUsers,
		TotalFreelancers:     stats.TotalFreelancers,
		TotalClients:         stats.TotalClients,
		VerifiedUsers:        stats.VerifiedUsers,
		OnlineUsers:          stats.OnlineUsers,
		// Add other fields as needed
	}
}