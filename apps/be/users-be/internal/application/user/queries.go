// internal/application/user/queries.go
package user

import (
	"context"
	userDomain "users-be/internal/domain/user"
)

// ============================================================================
// QUERY DEFINITIONS (CQRS Pattern)
// ============================================================================

// GetUserByIDQuery represents a query to get user by ID
type GetUserByIDQuery struct {
	UserID string
}

// GetUserByKeycloakIDQuery represents a query to get user by Keycloak ID
type GetUserByKeycloakIDQuery struct {
	KeycloakID string
}

// GetUserByEmailQuery represents a query to get user by email
type GetUserByEmailQuery struct {
	Email string
}

// GetUserByUsernameQuery represents a query to get user by username
type GetUserByUsernameQuery struct {
	Username string
}

// GetPublicProfileQuery represents a query to get public profile
type GetPublicProfileQuery struct {
	UserID string
}

// ListUsersQuery represents a query to list users
type ListUsersQuery struct {
	Filter userDomain.ListFilter
}

// SearchUsersQuery represents a query to search users
type SearchUsersQuery struct {
	SearchQuery string
	Filter      userDomain.ListFilter
}

// GetTopRatedFreelancersQuery represents a query to get top rated freelancers
type GetTopRatedFreelancersQuery struct {
	Limit int
}

// GetFeaturedUsersQuery represents a query to get featured users
type GetFeaturedUsersQuery struct {
	UserType userDomain.UserType
	Limit    int
}

// GetOnlineUsersQuery represents a query to get online users
type GetOnlineUsersQuery struct {
	UserType userDomain.UserType
}

// GetInactiveUsersQuery represents a query to get inactive users
type GetInactiveUsersQuery struct {
	Days int
}

// GetUsersWithWarningsQuery represents a query to get users with warnings
type GetUsersWithWarningsQuery struct{}

// GetUsersByCountryQuery represents a query to get users by country
type GetUsersByCountryQuery struct {
	Country string
	Filter  userDomain.ListFilter
}

// GetUserStatisticsQuery represents a query to get user statistics
type GetUserStatisticsQuery struct{}

// GetUserGrowthStatsQuery represents a query to get growth statistics
type GetUserGrowthStatsQuery struct {
	Days int
}

// GetTopCountriesQuery represents a query to get top countries
type GetTopCountriesQuery struct {
	Limit int
}

// GetUsersByStatusQuery represents a query to get users by status
type GetUsersByStatusQuery struct {
	Status userDomain.AccountStatus
	Filter userDomain.ListFilter
}

// GetVerifiedUsersQuery represents a query to get verified users
type GetVerifiedUsersQuery struct {
	Filter userDomain.ListFilter
}

// GetUnverifiedUsersQuery represents a query to get unverified users
type GetUnverifiedUsersQuery struct {
	Filter userDomain.ListFilter
}

// GetRecentlyActiveUsersQuery represents a query to get recently active users
type GetRecentlyActiveUsersQuery struct {
	Hours  int
	Filter userDomain.ListFilter
}

// GetNewUsersQuery represents a query to get new users
type GetNewUsersQuery struct {
	Days   int
	Filter userDomain.ListFilter
}

// CountUsersByTypeQuery represents a query to count users by type
type CountUsersByTypeQuery struct {
	UserType userDomain.UserType
}

// CountUsersByStatusQuery represents a query to count users by status
type CountUsersByStatusQuery struct {
	Status userDomain.AccountStatus
}

// CountOnlineUsersQuery represents a query to count online users
type CountOnlineUsersQuery struct{}

// GetUsersByReferrerQuery represents a query to get users by referrer
type GetUsersByReferrerQuery struct {
	ReferrerID string
}

// ============================================================================
// QUERY HANDLER
// ============================================================================

// QueryHandler handles user queries
type QueryHandler struct {
	service *Service
}

// NewQueryHandler creates a new query handler
func NewQueryHandler(service *Service) *QueryHandler {
	return &QueryHandler{service: service}
}

// HandleGetUserByID handles get user by ID query
func (h *QueryHandler) HandleGetUserByID(ctx context.Context, query GetUserByIDQuery) (*UserDTO, error) {
	return h.service.GetUserByID(ctx, query.UserID)
}

// HandleGetUserByKeycloakID handles get user by Keycloak ID query
func (h *QueryHandler) HandleGetUserByKeycloakID(ctx context.Context, query GetUserByKeycloakIDQuery) (*UserDTO, error) {
	return h.service.GetUserByKeycloakID(ctx, query.KeycloakID)
}

// HandleGetUserByEmail handles get user by email query
func (h *QueryHandler) HandleGetUserByEmail(ctx context.Context, query GetUserByEmailQuery) (*UserDTO, error) {
	return h.service.GetUserByEmail(ctx, query.Email)
}

// HandleGetUserByUsername handles get user by username query
func (h *QueryHandler) HandleGetUserByUsername(ctx context.Context, query GetUserByUsernameQuery) (*UserDTO, error) {
	return h.service.GetUserByUsername(ctx, query.Username)
}

// HandleGetPublicProfile handles get public profile query
func (h *QueryHandler) HandleGetPublicProfile(ctx context.Context, query GetPublicProfileQuery) (*PublicUserDTO, error) {
	return h.service.GetPublicProfile(ctx, query.UserID)
}

// HandleListUsers handles list users query
func (h *QueryHandler) HandleListUsers(ctx context.Context, query ListUsersQuery) (*UserListResponseDTO, error) {
	return h.service.ListUsers(ctx, query.Filter)
}

// HandleSearchUsers handles search users query
func (h *QueryHandler) HandleSearchUsers(ctx context.Context, query SearchUsersQuery) (*UserListResponseDTO, error) {
	return h.service.SearchUsers(ctx, query.SearchQuery, query.Filter)
}

// HandleGetUserStatistics handles get user statistics query
func (h *QueryHandler) HandleGetUserStatistics(ctx context.Context, query GetUserStatisticsQuery) (*userDomain.UserStatistics, error) {
	return h.service.GetUserStatistics(ctx)
}

// HandleGetUserGrowthStats handles get growth statistics query
func (h *QueryHandler) HandleGetUserGrowthStats(ctx context.Context, query GetUserGrowthStatsQuery) (*userDomain.UserGrowthStats, error) {
	return h.service.GetUserGrowthStats(ctx, query.Days)
}

// ============================================================================
// QUERY VALIDATION
// ============================================================================

// Validate validates GetUserByIDQuery
func (q GetUserByIDQuery) Validate() error {
	if q.UserID == "" {
		return userDomain.ErrInvalidUserID
	}
	return nil
}

// Validate validates GetUserByKeycloakIDQuery
func (q GetUserByKeycloakIDQuery) Validate() error {
	if q.KeycloakID == "" {
		return userDomain.ErrInvalidKeycloakID
	}
	return nil
}

// Validate validates GetUserByEmailQuery
func (q GetUserByEmailQuery) Validate() error {
	if q.Email == "" {
		return userDomain.ErrEmailRequired
	}
	return nil
}

// Validate validates GetUserByUsernameQuery
func (q GetUserByUsernameQuery) Validate() error {
	if q.Username == "" {
		return userDomain.ErrUsernameRequired
	}
	return nil
}

// Validate validates ListUsersQuery
func (q ListUsersQuery) Validate() error {
	return q.Filter.Validate()
}

// Validate validates SearchUsersQuery
func (q SearchUsersQuery) Validate() error {
	if q.SearchQuery == "" {
		return userDomain.ErrInvalidSearchQuery
	}
	if len(q.SearchQuery) < 2 {
		return userDomain.ErrInvalidSearchQuery
	}
	return q.Filter.Validate()
}

// Validate validates GetTopRatedFreelancersQuery
func (q GetTopRatedFreelancersQuery) Validate() error {
	if q.Limit <= 0 {
		return userDomain.ErrInvalidInput
	}
	if q.Limit > 100 {
		return userDomain.ErrInvalidInput
	}
	return nil
}

// Validate validates GetUsersByCountryQuery
func (q GetUsersByCountryQuery) Validate() error {
	if q.Country == "" {
		return userDomain.ErrInvalidInput
	}
	return q.Filter.Validate()
}

// Validate validates GetUserGrowthStatsQuery
func (q GetUserGrowthStatsQuery) Validate() error {
	if q.Days <= 0 {
		return userDomain.ErrInvalidInput
	}
	if q.Days > 365 {
		return userDomain.ErrInvalidInput
	}
	return nil
}

// ============================================================================
// QUERY RESULT TYPES
// ============================================================================

// UserQueryResult represents a single user query result
type UserQueryResult struct {
	User  *UserDTO
	Error error
}

// UsersQueryResult represents multiple users query result
type UsersQueryResult struct {
	Users []*UserDTO
	Total int64
	Error error
}

// StatisticsQueryResult represents statistics query result
type StatisticsQueryResult struct {
	Statistics *userDomain.UserStatistics
	Error      error
}