package user

import (
	"time"

	"github.com/google/uuid"
)

// CreateUserDTO is the data transfer object for creating a user
type CreateUserDTO struct {
	KeycloakID    string  `json:"keycloak_id" binding:"required"`
	Username      string  `json:"username" binding:"required"`
	Email         string  `json:"email" binding:"required,email"`
	FirstName     string  `json:"first_name"`
	LastName      string  `json:"last_name"`
	EmailVerified bool    `json:"email_verified"`
}

// UpdateUserDTO is the data transfer object for updating a user
type UpdateUserDTO struct {
	FirstName      *string  `json:"first_name"`
	LastName       *string  `json:"last_name"`
	PhoneNumber    *string  `json:"phone_number"`
	Bio            *string  `json:"bio"`
	ProfileType    *string  `json:"profile_type"` // freelancer, client, both
	Profession     *string  `json:"profession"`
	HourlyRate     *float64 `json:"hourly_rate"`
	AvailableHours *int     `json:"available_hours"`
	Country        *string  `json:"country"`
	City           *string  `json:"city"`
}

// UserResponseDTO is the data transfer object for user responses
type UserResponseDTO struct {
	ID              uuid.UUID  `json:"id"`
	KeycloakID      string     `json:"keycloak_id"`
	Username        string     `json:"username"`
	Email           string     `json:"email"`
	FirstName       string     `json:"first_name"`
	LastName        string     `json:"last_name"`
	PhoneNumber     *string    `json:"phone_number,omitempty"`
	Bio             *string    `json:"bio,omitempty"`
	ProfileType     string     `json:"profile_type"`
	Profession      *string    `json:"profession,omitempty"`
	HourlyRate      *float64   `json:"hourly_rate,omitempty"`
	AvailableHours  *int       `json:"available_hours,omitempty"`
	Country         *string    `json:"country,omitempty"`
	City            *string    `json:"city,omitempty"`
	IsActive        bool       `json:"is_active"`
	EmailVerified   bool       `json:"email_verified"`
	ProfileComplete bool       `json:"profile_complete"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`
}

// ListUsersResponseDTO is the response for listing users
type ListUsersResponseDTO struct {
	Users      []UserResponseDTO `json:"users"`
	Total      int64             `json:"total"`
	Page       int               `json:"page"`
	PageSize   int               `json:"page_size"`
	TotalPages int               `json:"total_pages"`
}