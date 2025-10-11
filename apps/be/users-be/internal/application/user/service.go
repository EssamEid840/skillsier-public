package user

import (
	"context"
	"encoding/json"
	"fmt"
	"math"

	"users-be/internal/domain/outbox"
	"users-be/internal/domain/user"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Service handles user business logic
type Service struct {
	userRepo   user.Repository
	outboxRepo outbox.Repository
	db         *gorm.DB
}

// NewService creates a new user service
func NewService(userRepo user.Repository, outboxRepo outbox.Repository, db *gorm.DB) *Service {
	return &Service{
		userRepo:   userRepo,
		outboxRepo: outboxRepo,
		db:         db,
	}
}

// CreateUser creates a new user and publishes a user.created event
func (s *Service) CreateUser(ctx context.Context, dto *CreateUserDTO) (*UserResponseDTO, error) {
	// Check if user already exists
	exists, err := s.userRepo.ExistsByKeycloakID(ctx, dto.KeycloakID)
	if err != nil {
		return nil, fmt.Errorf("failed to check user existence: %w", err)
	}
	if exists {
		return nil, fmt.Errorf("user with Keycloak ID %s already exists", dto.KeycloakID)
	}

	// Convert DTO to entity
	newUser := dto.ToEntity()

	// Use transaction to ensure atomicity (user creation + outbox event)
	err = s.db.Transaction(func(tx *gorm.DB) error {
		// Create user
		if err := s.userRepo.Create(ctx, newUser); err != nil {
			return fmt.Errorf("failed to create user: %w", err)
		}

		// Create outbox event for user.created
		event, err := s.createUserEvent("user.created", newUser)
		if err != nil {
			return fmt.Errorf("failed to create user event: %w", err)
		}

		if err := s.outboxRepo.Create(ctx, event); err != nil {
			return fmt.Errorf("failed to create outbox event: %w", err)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	// Convert to response DTO
	response := ToResponseDTO(newUser)
	return &response, nil
}

// GetUser retrieves a user by ID
func (s *Service) GetUser(ctx context.Context, id uuid.UUID) (*UserResponseDTO, error) {
	u, err := s.userRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	response := ToResponseDTO(u)
	return &response, nil
}

// GetUserByKeycloakID retrieves a user by Keycloak ID
func (s *Service) GetUserByKeycloakID(ctx context.Context, keycloakID string) (*UserResponseDTO, error) {
	u, err := s.userRepo.FindByKeycloakID(ctx, keycloakID)
	if err != nil {
		return nil, err
	}

	response := ToResponseDTO(u)
	return &response, nil
}

// UpdateUser updates user information
func (s *Service) UpdateUser(ctx context.Context, id uuid.UUID, dto *UpdateUserDTO) (*UserResponseDTO, error) {
	// Find existing user
	u, err := s.userRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	// Store old values for event
	oldUser := *u

	// Apply updates
	dto.ApplyUpdates(u)

	// Use transaction for update + outbox event
	err = s.db.Transaction(func(tx *gorm.DB) error {
		// Update user
		if err := s.userRepo.Update(ctx, u); err != nil {
			return fmt.Errorf("failed to update user: %w", err)
		}

		// Create outbox event for user.updated
		event, err := s.createUserUpdatedEvent(u, &oldUser)
		if err != nil {
			return fmt.Errorf("failed to create update event: %w", err)
		}

		if err := s.outboxRepo.Create(ctx, event); err != nil {
			return fmt.Errorf("failed to create outbox event: %w", err)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	response := ToResponseDTO(u)
	return &response, nil
}

// DeleteUser soft-deletes a user
func (s *Service) DeleteUser(ctx context.Context, id uuid.UUID) error {
	// Verify user exists
	u, err := s.userRepo.FindByID(ctx, id)
	if err != nil {
		return err
	}

	// Use transaction for delete + outbox event
	return s.db.Transaction(func(tx *gorm.DB) error {
		// Delete user
		if err := s.userRepo.Delete(ctx, id); err != nil {
			return fmt.Errorf("failed to delete user: %w", err)
		}

		// Create outbox event for user.deleted
		event, err := s.createUserEvent("user.deleted", u)
		if err != nil {
			return fmt.Errorf("failed to create delete event: %w", err)
		}

		if err := s.outboxRepo.Create(ctx, event); err != nil {
			return fmt.Errorf("failed to create outbox event: %w", err)
		}

		return nil
	})
}

// ListUsers retrieves a paginated list of users
func (s *Service) ListUsers(ctx context.Context, page, pageSize int) (*ListUsersResponseDTO, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 20
	}

	offset := (page - 1) * pageSize

	users, total, err := s.userRepo.List(ctx, pageSize, offset)
	if err != nil {
		return nil, err
	}

	totalPages := int(math.Ceil(float64(total) / float64(pageSize)))

	return &ListUsersResponseDTO{
		Users:      ToResponseDTOList(users),
		Total:      total,
		Page:       page,
		PageSize:   pageSize,
		TotalPages: totalPages,
	}, nil
}

// createUserEvent creates an outbox event for user operations
func (s *Service) createUserEvent(eventType string, u *user.User) (*outbox.Event, error) {
	// Create event payload
	payload := map[string]interface{}{
		"user_id":       u.ID.String(),
		"keycloak_id":   u.KeycloakID,
		"username":      u.Username,
		"email":         u.Email,
		"first_name":    u.FirstName,
		"last_name":     u.LastName,
		"profile_type":  u.ProfileType,
		"is_active":     u.IsActive,
		"email_verified": u.EmailVerified,
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal payload: %w", err)
	}

	// Create metadata
	metadata := map[string]interface{}{
		"source": "users-be",
	}
	metadataBytes, err := json.Marshal(metadata)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal metadata: %w", err)
	}

	return &outbox.Event{
		AggregateID:   u.ID.String(),
		AggregateType: "user",
		EventType:     eventType,
		Version:       1,
		Payload:       payloadBytes,
		Metadata:      metadataBytes,
		Topic:         "user-events",
		Status:        outbox.EventStatusPending,
	}, nil
}

// createUserUpdatedEvent creates an event with old and new values
func (s *Service) createUserUpdatedEvent(newUser, oldUser *user.User) (*outbox.Event, error) {
	payload := map[string]interface{}{
		"user_id":     newUser.ID.String(),
		"keycloak_id": newUser.KeycloakID,
		"changes": map[string]interface{}{
			"old": map[string]interface{}{
				"first_name":    oldUser.FirstName,
				"last_name":     oldUser.LastName,
				"profile_type":  oldUser.ProfileType,
				"phone_number":  oldUser.PhoneNumber,
				"bio":           oldUser.Bio,
				"profession":    oldUser.Profession,
				"hourly_rate":   oldUser.HourlyRate,
				"country":       oldUser.Country,
				"city":          oldUser.City,
			},
			"new": map[string]interface{}{
				"first_name":    newUser.FirstName,
				"last_name":     newUser.LastName,
				"profile_type":  newUser.ProfileType,
				"phone_number":  newUser.PhoneNumber,
				"bio":           newUser.Bio,
				"profession":    newUser.Profession,
				"hourly_rate":   newUser.HourlyRate,
				"country":       newUser.Country,
				"city":          newUser.City,
			},
		},
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal payload: %w", err)
	}

	metadata := map[string]interface{}{
		"source": "users-be",
	}
	metadataBytes, err := json.Marshal(metadata)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal metadata: %w", err)
	}

	return &outbox.Event{
		AggregateID:   newUser.ID.String(),
		AggregateType: "user",
		EventType:     "user.updated",
		Version:       1,
		Payload:       payloadBytes,
		Metadata:      metadataBytes,
		Topic:         "user-events",
		Status:        outbox.EventStatusPending,
	}, nil
}