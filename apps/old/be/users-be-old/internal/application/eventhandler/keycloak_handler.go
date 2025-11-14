package eventhandler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"users-be/internal/application/user"

	"github.com/IBM/sarama"
)

// KeycloakEvent represents the actual structure of events from Keycloak
type KeycloakEvent struct {
	Type           string                 `json:"type"`           // e.g., "LOGIN", "USER-CREATE"
	RealmID        string                 `json:"realmId"`
	ID             string                 `json:"id"`
	Time           int64                  `json:"time"`
	ClientID       string                 `json:"clientId"`
	UserID         string                 `json:"userId"`         // Admin/actor who triggered the event
	IPAddress      string                 `json:"ipAddress"`
	ResourcePath   string                 `json:"resourcePath"`   // e.g., "users/b263791a-9545-412e-865c-9700a368d257"
	Representation string                 `json:"representation"` // JSON string of user data
	Details        map[string]interface{} `json:"details"`
	BridgeMetadata map[string]interface{} `json:"bridgeMetadata"`
}

// UserRepresentation represents the user data in the representation field
type UserRepresentation struct {
	Username      string                 `json:"username"`
	FirstName     string                 `json:"firstName"`
	LastName      string                 `json:"lastName"`
	Email         string                 `json:"email"`
	EmailVerified bool                   `json:"emailVerified"`
	Enabled       bool                   `json:"enabled"`
	Attributes    map[string]interface{} `json:"attributes"`
}

// KeycloakEventHandler handles events from Keycloak
type KeycloakEventHandler struct {
	userService *user.Service
}

// NewKeycloakEventHandler creates a new Keycloak event handler
func NewKeycloakEventHandler(userService *user.Service) *KeycloakEventHandler {
	return &KeycloakEventHandler{
		userService: userService,
	}
}

// HandleMessage processes a Kafka message containing a Keycloak event
func (h *KeycloakEventHandler) HandleMessage(ctx context.Context, message *sarama.ConsumerMessage) error {
	log.Printf("Received Keycloak event from topic=%s partition=%d offset=%d",
		message.Topic, message.Partition, message.Offset)

	// Parse the message
	var event KeycloakEvent
	if err := json.Unmarshal(message.Value, &event); err != nil {
		log.Printf("Failed to unmarshal Keycloak event: %v", err)
		return fmt.Errorf("failed to unmarshal Keycloak event: %w", err)
	}

	log.Printf("Processing Keycloak event: type=%s resourcePath=%s",
		event.Type, event.ResourcePath)

	// Route to appropriate handler based on event type
	switch event.Type {
	case "LOGIN":
		return h.handleUserLogin(ctx, &event)
	case "USER-CREATE", "REGISTER":
		return h.handleUserCreate(ctx, &event)
	case "USER-UPDATE", "UPDATE_PROFILE":
		return h.handleUserUpdate(ctx, &event)
	case "USER-DELETE":
		return h.handleUserDelete(ctx, &event)
	default:
		// Log unknown events but don't fail
		log.Printf("Ignoring Keycloak event type: %s", event.Type)
		return nil
	}
}

// handleUserCreate creates a new user when they are created in Keycloak
func (h *KeycloakEventHandler) handleUserCreate(ctx context.Context, event *KeycloakEvent) error {
	log.Printf("Handling user creation event...")

	// Extract user ID from resourcePath (e.g., "users/b263791a-9545-412e-865c-9700a368d257")
	var keycloakUserID string
	if event.ResourcePath != "" {
		parts := strings.Split(event.ResourcePath, "/")
		if len(parts) == 2 && parts[0] == "users" {
			keycloakUserID = parts[1]
		}
	}

	if keycloakUserID == "" {
		log.Printf("ERROR: Could not extract user ID from resourcePath: %s", event.ResourcePath)
		return fmt.Errorf("missing user ID in event")
	}

	log.Printf("Extracted Keycloak user ID: %s", keycloakUserID)

	// Parse the representation field (it's a JSON string)
	var userRep UserRepresentation
	if event.Representation != "" {
		if err := json.Unmarshal([]byte(event.Representation), &userRep); err != nil {
			log.Printf("ERROR: Failed to parse user representation: %v", err)
			return fmt.Errorf("failed to parse user representation: %w", err)
		}
	} else {
		log.Printf("ERROR: Representation field is empty")
		return fmt.Errorf("empty representation field")
	}

	log.Printf("Parsed user data: username=%s email=%s firstName=%s lastName=%s",
		userRep.Username, userRep.Email, userRep.FirstName, userRep.LastName)

	// Check if user already exists
	existingUser, err := h.userService.GetUserByKeycloakID(ctx, keycloakUserID)
	if err == nil && existingUser != nil {
		log.Printf("User %s already exists in database, skipping creation", keycloakUserID)
		return nil
	}

	// Create user DTO
	dto := &user.CreateUserDTO{
		KeycloakID:    keycloakUserID,
		Username:      userRep.Username,
		Email:         userRep.Email,
		FirstName:     userRep.FirstName,
		LastName:      userRep.LastName,
		EmailVerified: userRep.EmailVerified,
	}

	// Create user in database
	createdUser, err := h.userService.CreateUser(ctx, dto)
	if err != nil {
		log.Printf("ERROR: Failed to create user: %v", err)
		return fmt.Errorf("failed to create user from Keycloak event: %w", err)
	}

	log.Printf("✓ Successfully created user: id=%s username=%s email=%s keycloak_id=%s",
		createdUser.ID, createdUser.Username, createdUser.Email, createdUser.KeycloakID)

	return nil
}

// handleUserDelete handles user deletion from Keycloak
func (h *KeycloakEventHandler) handleUserDelete(ctx context.Context, event *KeycloakEvent) error {
	log.Printf("Handling user deletion event...")

	// Extract user ID from resourcePath
	var keycloakUserID string
	if event.ResourcePath != "" {
		parts := strings.Split(event.ResourcePath, "/")
		if len(parts) == 2 && parts[0] == "users" {
			keycloakUserID = parts[1]
		}
	}

	if keycloakUserID == "" {
		log.Printf("ERROR: Could not extract user ID from resourcePath: %s", event.ResourcePath)
		return fmt.Errorf("missing user ID in event")
	}

	log.Printf("Extracted Keycloak user ID for deletion: %s", keycloakUserID)

	// Get existing user
	existingUser, err := h.userService.GetUserByKeycloakID(ctx, keycloakUserID)
	if err != nil {
		log.Printf("User %s not found in database, nothing to delete", keycloakUserID)
		return nil // Not an error if user doesn't exist
	}

	// Delete user (soft delete)
	if err := h.userService.DeleteUser(ctx, existingUser.ID); err != nil {
		log.Printf("ERROR: Failed to delete user: %v", err)
		return fmt.Errorf("failed to delete user from Keycloak event: %w", err)
	}

	log.Printf("✓ Successfully deleted user: id=%s username=%s keycloak_id=%s",
		existingUser.ID, existingUser.Username, keycloakUserID)

	return nil
}

// handleUserLogin handles user login events
func (h *KeycloakEventHandler) handleUserLogin(ctx context.Context, event *KeycloakEvent) error {
	log.Printf("User login event: userId=%s clientId=%s", event.UserID, event.ClientID)

	// Extract username from details if available
	username := ""
	if event.Details != nil {
		if un, ok := event.Details["username"].(string); ok {
			username = un
		}
	}

	// Try to find user by Keycloak ID
	existingUser, err := h.userService.GetUserByKeycloakID(ctx, event.UserID)
	if err != nil || existingUser == nil {
		log.Printf("User %s not found in database (username: %s). This might be an admin or the user wasn't synced yet.",
			event.UserID, username)
		return nil
	}

	log.Printf("User %s (username: %s) logged in", existingUser.ID, existingUser.Username)

	// You could update last_login timestamp here if needed
	return nil
}

// handleUserUpdate handles user profile updates from Keycloak
func (h *KeycloakEventHandler) handleUserUpdate(ctx context.Context, event *KeycloakEvent) error {
	log.Printf("Handling user update event...")

	// Extract user ID from resourcePath
	var keycloakUserID string
	if event.ResourcePath != "" {
		parts := strings.Split(event.ResourcePath, "/")
		if len(parts) == 2 && parts[0] == "users" {
			keycloakUserID = parts[1]
		}
	}

	if keycloakUserID == "" {
		log.Printf("ERROR: Could not extract user ID from resourcePath: %s", event.ResourcePath)
		return fmt.Errorf("missing user ID in event")
	}

	// Parse the representation field
	var userRep UserRepresentation
	if event.Representation != "" {
		if err := json.Unmarshal([]byte(event.Representation), &userRep); err != nil {
			log.Printf("ERROR: Failed to parse user representation: %v", err)
			return fmt.Errorf("failed to parse user representation: %w", err)
		}
	} else {
		log.Printf("WARNING: Representation field is empty for update event")
		return nil
	}

	// Get existing user
	existingUser, err := h.userService.GetUserByKeycloakID(ctx, keycloakUserID)
	if err != nil {
		log.Printf("User %s not found, creating from update event", keycloakUserID)
		// User doesn't exist, create them
		dto := &user.CreateUserDTO{
			KeycloakID:    keycloakUserID,
			Username:      userRep.Username,
			Email:         userRep.Email,
			FirstName:     userRep.FirstName,
			LastName:      userRep.LastName,
			EmailVerified: userRep.EmailVerified,
		}
		_, err := h.userService.CreateUser(ctx, dto)
		return err
	}

	// Create update DTO
	updateDTO := &user.UpdateUserDTO{}
	
	if userRep.FirstName != "" && userRep.FirstName != existingUser.FirstName {
		updateDTO.FirstName = &userRep.FirstName
	}
	if userRep.LastName != "" && userRep.LastName != existingUser.LastName {
		updateDTO.LastName = &userRep.LastName
	}

	// Only update if there are actual changes
	if updateDTO.FirstName != nil || updateDTO.LastName != nil {
		_, err := h.userService.UpdateUser(ctx, existingUser.ID, updateDTO)
		if err != nil {
			return fmt.Errorf("failed to update user from Keycloak event: %w", err)
		}
		log.Printf("✓ Successfully updated user: id=%s", existingUser.ID)
	} else {
		log.Printf("No changes detected for user: id=%s", existingUser.ID)
	}

	return nil
}

