package user

import (
	"users-be/internal/domain/user"
)

// ToEntity converts CreateUserDTO to User entity
func (dto *CreateUserDTO) ToEntity() *user.User {
	return &user.User{
		KeycloakID:    dto.KeycloakID,
		Username:      dto.Username,
		Email:         dto.Email,
		FirstName:     dto.FirstName,
		LastName:      dto.LastName,
		EmailVerified: dto.EmailVerified,
		IsActive:      true,
	}
}

// ToResponseDTO converts User entity to UserResponseDTO
func ToResponseDTO(u *user.User) UserResponseDTO {
	return UserResponseDTO{
		ID:              u.ID,
		KeycloakID:      u.KeycloakID,
		Username:        u.Username,
		Email:           u.Email,
		FirstName:       u.FirstName,
		LastName:        u.LastName,
		PhoneNumber:     u.PhoneNumber,
		Bio:             u.Bio,
		ProfileType:     u.ProfileType,
		Profession:      u.Profession,
		HourlyRate:      u.HourlyRate,
		AvailableHours:  u.AvailableHours,
		Country:         u.Country,
		City:            u.City,
		IsActive:        u.IsActive,
		EmailVerified:   u.EmailVerified,
		ProfileComplete: u.ProfileComplete,
		CreatedAt:       u.CreatedAt,
		UpdatedAt:       u.UpdatedAt,
	}
}

// ToResponseDTOList converts a slice of User entities to UserResponseDTOs
func ToResponseDTOList(users []*user.User) []UserResponseDTO {
	dtos := make([]UserResponseDTO, len(users))
	for i, u := range users {
		dtos[i] = ToResponseDTO(u)
	}
	return dtos
}

// ApplyUpdates applies UpdateUserDTO to User entity
func (dto *UpdateUserDTO) ApplyUpdates(u *user.User) {
	if dto.FirstName != nil {
		u.FirstName = *dto.FirstName
	}
	if dto.LastName != nil {
		u.LastName = *dto.LastName
	}
	if dto.PhoneNumber != nil {
		u.PhoneNumber = dto.PhoneNumber
	}
	if dto.Bio != nil {
		u.Bio = dto.Bio
	}
	if dto.ProfileType != nil {
		u.ProfileType = *dto.ProfileType
	}
	if dto.Profession != nil {
		u.Profession = dto.Profession
	}
	if dto.HourlyRate != nil {
		u.HourlyRate = dto.HourlyRate
	}
	if dto.AvailableHours != nil {
		u.AvailableHours = dto.AvailableHours
	}
	if dto.Country != nil {
		u.Country = dto.Country
	}
	if dto.City != nil {
		u.City = dto.City
	}
	
	// Update profile completion status
	u.ProfileComplete = u.IsProfileCompleted()
}