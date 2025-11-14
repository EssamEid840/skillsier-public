// apps/be/users-be/internal/application/user/mapper.go
package user

import (
	"users-be/internal/domain/user"
)

// ============================================================================
// DOMAIN TO DTO MAPPERS
// ============================================================================

// ToDTO converts a User entity to UserDTO
func ToDTO(u *user.User) *UserDTO {
	if u == nil {
		return nil
	}
	
	dto := &UserDTO{
		ID:                   u.ID,
		KeycloakID:           u.KeycloakID,
		Username:             u.Username,
		Email:                u.Email.Value,
		EmailVerified:        u.Email.Verified,
		FirstName:            u.FirstName,
		LastName:             u.LastName,
		DisplayName:          u.DisplayName,
		FullName:             u.FullName,
		UserType:             string(u.UserType),
		Bio:                  u.Bio,
		Tagline:              u.Tagline,
		Title:                u.Title,
		Website:              u.Website,
		ProfilePictureURL:    u.ProfilePictureURL,
		CoverImageURL:        u.CoverImageURL,
		ProfileCompleteness:  u.ProfileCompleteness,
		PhoneVerified:        u.Phone != nil && u.Phone.Verified,
		City:                 u.City,
		Country:              u.Country,
		CountryCode:          u.CountryCode,
		Timezone:             u.Timezone,
		Status:               string(u.Status),
		VerificationStatus:   string(u.VerificationStatus),
		IsOnline:             u.IsOnline,
		LastSeenAt:           u.LastSeenAt,
		LastActiveAt:         u.LastActiveAt,
		AvailabilityStatus:   string(u.AvailabilityStatus),
		HoursPerWeek:         u.HoursPerWeek,
		ResponseTime:         u.ResponseTime,
		TotalEarned:          u.TotalEarned,
		TotalSpent:           u.TotalSpent,
		JobSuccess:           u.JobSuccess,
		TotalJobs:            u.TotalJobs,
		TotalHires:           u.TotalHires,
		AvgRating:            u.AvgRating,
		ReviewCount:          u.ReviewCount,
		ConnectsBalance:      u.ConnectsBalance,
		ProfileVisibility:    string(u.ProfileVisibility),
		ShowEmail:            u.ShowEmail,
		ShowPhone:            u.ShowPhone,
		ShowLocation:         u.ShowLocation,
		SearchableProfile:    u.SearchableProfile,
		AcceptingWork:        u.AcceptingWork,
		Badges:               u.Badges,
		Warnings:             u.Warnings,
		ReferralCode:         u.ReferralCode,
		ReferredBy:           u.ReferredBy,
		ReferralCount:        u.ReferralCount,
		IsFeatured:           u.IsFeatured,
		IsTopRated:           u.IsTopRated,
		IsRisingTalent:       u.IsRisingTalent,
		IsExpertVetted:       u.IsExpertVetted,
		LastLoginAt:          u.LastLoginAt,
		LastLoginIP:          u.LastLoginIP,
		LoginCount:           u.LoginCount,
		CreatedAt:            u.CreatedAt,
		UpdatedAt:            u.UpdatedAt,
	}
	
	// Add phone info if exists
	if u.Phone != nil {
		dto.PhoneNumber = &u.Phone.Number
		dto.PhoneCountryCode = &u.Phone.CountryCode
	}
	
	return dto
}

// ToDTOList converts a slice of User entities to UserDTOs
func ToDTOList(users []*user.User) []UserDTO {
	if users == nil {
		return []UserDTO{}
	}
	
	dtos := make([]UserDTO, 0, len(users))
	for _, u := range users {
		if dto := ToDTO(u); dto != nil {
			dtos = append(dtos, *dto)
		}
	}
	return dtos
}

// ToListDTO converts users slice with pagination info to UserListDTO
func ToListDTO(users []*user.User, total int64, page, pageSize int) *UserListDTO {
	totalPages := int(total) / pageSize
	if int(total)%pageSize > 0 {
		totalPages++
	}
	
	return &UserListDTO{
		Users:      ToDTOList(users),
		Total:      total,
		Page:       page,
		PageSize:   pageSize,
		TotalPages: totalPages,
	}
}

// ToSearchResultDTO converts users slice to search result DTO
func ToSearchResultDTO(users []*user.User, total int64, query string, searchTime int64) *UserSearchResultDTO {
	return &UserSearchResultDTO{
		Users:      ToDTOList(users),
		Total:      total,
		Query:      query,
		SearchTime: searchTime,
	}
}

// ToStatisticsDTO converts UserStatistics entity to DTO
func ToStatisticsDTO(stats *user.UserStatistics) *UserStatisticsDTO {
	if stats == nil {
		return nil
	}
	
	return &UserStatisticsDTO{
		TotalUsers:           stats.TotalUsers,
		TotalFreelancers:     stats.TotalFreelancers,
		TotalClients:         stats.TotalClients,
		ActiveUsers:          stats.ActiveUsers,
		SuspendedUsers:       stats.SuspendedUsers,
		BannedUsers:          stats.BannedUsers,
		VerifiedUsers:        stats.VerifiedUsers,
		UnverifiedUsers:      stats.UnverifiedUsers,
		OnlineUsers:          stats.OnlineUsers,
		NewUsersLast30Days:   stats.NewUsersLast30Days,
		UsersByCountry:       stats.UsersByCountry,
		UsersByUserType:      stats.UsersByUserType,
		UsersByStatus:        stats.UsersByStatus,
		AverageCompleteness:  stats.AverageCompleteness,
		AverageResponseTime:  stats.AverageResponseTime,
	}
}

// ============================================================================
// DTO TO DOMAIN MAPPERS
// ============================================================================

// ToEntity converts CreateUserDTO to User entity
func (dto *CreateUserDTO) ToEntity() (*user.User, error) {
	u := &user.User{
		KeycloakID: dto.KeycloakID,
		Username:   dto.Username,
		FirstName:  dto.FirstName,
		LastName:   dto.LastName,
		UserType:   dto.UserType,
		City:       dto.City,
		Country:    dto.Country,
		CountryCode: dto.CountryCode,
		Timezone:   dto.Timezone,
		Bio:        dto.Bio,
		Tagline:    dto.Tagline,
	}
	
	// Set email
	email, err := user.NewEmail(dto.Email)
	if err != nil {
		return nil, err
	}
	u.Email = email
	
	// Set phone if provided
	if dto.PhoneNumber != nil && *dto.PhoneNumber != "" {
		countryCode := "+1" // default
		if dto.PhoneCountryCode != nil {
			countryCode = *dto.PhoneCountryCode
		}
		phone, err := user.NewPhone(*dto.PhoneNumber, countryCode)
		if err != nil {
			return nil, err
		}
		u.Phone = phone
	}
	
	// Generate full name
	u.FullName = u.GenerateFullName()
	
	// Set default values
	u.Status = user.StatusActive
	u.VerificationStatus = user.VerificationUnverified
	u.AvailabilityStatus = user.AvailabilityAvailable
	u.ProfileVisibility = user.VisibilityPublic
	u.ShowEmail = false
	u.ShowPhone = false
	u.ShowLocation = true
	u.SearchableProfile = true
	u.AcceptingWork = true
	u.ConnectsBalance = 0
	u.ProfileCompleteness = u.CalculateCompleteness()
	
	return u, nil
}

// ApplyUpdates applies UpdateUserDTO to User entity
func (dto *UpdateUserDTO) ApplyUpdates(u *user.User) error {
	if dto.FirstName != nil {
		u.FirstName = dto.FirstName
	}
	if dto.LastName != nil {
		u.LastName = dto.LastName
	}
	if dto.DisplayName != nil {
		u.DisplayName = dto.DisplayName
	}
	if dto.Bio != nil {
		u.Bio = dto.Bio
	}
	if dto.Tagline != nil {
		u.Tagline = dto.Tagline
	}
	if dto.Title != nil {
		u.Title = dto.Title
	}
	if dto.Website != nil {
		u.Website = dto.Website
	}
	if dto.ProfilePictureURL != nil {
		u.ProfilePictureURL = dto.ProfilePictureURL
	}
	if dto.CoverImageURL != nil {
		u.CoverImageURL = dto.CoverImageURL
	}
	if dto.City != nil {
		u.City = dto.City
	}
	if dto.Country != nil {
		u.Country = dto.Country
	}
	if dto.CountryCode != nil {
		u.CountryCode = dto.CountryCode
	}
	if dto.Timezone != nil {
		u.Timezone = dto.Timezone
	}
	
	// Update phone if provided
	if dto.PhoneNumber != nil && *dto.PhoneNumber != "" {
		countryCode := "+1" // default
		if dto.PhoneCountryCode != nil {
			countryCode = *dto.PhoneCountryCode
		}
		phone, err := user.NewPhone(*dto.PhoneNumber, countryCode)
		if err != nil {
			return err
		}
		u.Phone = phone
	}
	
	// Regenerate full name if first or last name changed
	if dto.FirstName != nil || dto.LastName != nil {
		u.FullName = u.GenerateFullName()
	}
	
	// Recalculate profile completeness
	u.ProfileCompleteness = u.CalculateCompleteness()
	
	return nil
}

// ApplyAvailabilityUpdates applies UpdateAvailabilityDTO to User entity
func (dto *UpdateAvailabilityDTO) ApplyUpdates(u *user.User) {
	if dto.Status != nil {
		u.AvailabilityStatus = *dto.Status
	}
	if dto.HoursPerWeek != nil {
		u.HoursPerWeek = dto.HoursPerWeek
	}
}

// ApplySettingsUpdates applies UpdateSettingsDTO to User entity
func (dto *UpdateSettingsDTO) ApplyUpdates(u *user.User) {
	if dto.ProfileVisibility != nil {
		u.ProfileVisibility = *dto.ProfileVisibility
	}
	if dto.ShowEmail != nil {
		u.ShowEmail = *dto.ShowEmail
	}
	if dto.ShowPhone != nil {
		u.ShowPhone = *dto.ShowPhone
	}
	if dto.ShowLocation != nil {
		u.ShowLocation = *dto.ShowLocation
	}
	if dto.SearchableProfile != nil {
		u.SearchableProfile = *dto.SearchableProfile
	}
	if dto.AcceptingWork != nil {
		u.AcceptingWork = *dto.AcceptingWork
	}
}

// ============================================================================
// FILTER MAPPERS
// ============================================================================

// ToListFilter converts UserFilterDTO to domain ListFilter
func (dto *UserFilterDTO) ToListFilter() user.ListFilter {
	filter := user.ListFilter{}
	
	// Convert string slices to domain types
	if len(dto.UserTypes) > 0 {
		filter.UserTypes = make([]user.UserType, len(dto.UserTypes))
		for i, t := range dto.UserTypes {
			filter.UserTypes[i] = user.UserType(t)
		}
	}
	
	if len(dto.Statuses) > 0 {
		filter.Statuses = make([]user.AccountStatus, len(dto.Statuses))
		for i, s := range dto.Statuses {
			filter.Statuses[i] = user.AccountStatus(s)
		}
	}
	
	if len(dto.VerificationStatuses) > 0 {
		filter.VerificationStatuses = make([]user.VerificationStatus, len(dto.VerificationStatuses))
		for i, v := range dto.VerificationStatuses {
			filter.VerificationStatuses[i] = user.VerificationStatus(v)
		}
	}
	
	filter.Countries = dto.Countries
	filter.Cities = dto.Cities
	filter.MinCompleteness = dto.MinCompleteness
	filter.MaxCompleteness = dto.MaxCompleteness
	filter.MinRating = dto.MinRating
	filter.MaxRating = dto.MaxRating
	filter.HasBadges = dto.HasBadges
	filter.IsOnline = dto.IsOnline
	filter.IsFeatured = dto.IsFeatured
	filter.IsTopRated = dto.IsTopRated
	filter.IsRisingTalent = dto.IsRisingTalent
	filter.IsExpertVetted = dto.IsExpertVetted
	filter.CreatedAfter = dto.CreatedAfter
	filter.CreatedBefore = dto.CreatedBefore
	filter.LastSeenAfter = dto.LastSeenAfter
	filter.LastSeenBefore = dto.LastSeenBefore
	filter.SortBy = dto.SortBy
	filter.SortOrder = dto.SortOrder
	filter.Page = dto.Page
	filter.PageSize = dto.PageSize
	
	return filter
}