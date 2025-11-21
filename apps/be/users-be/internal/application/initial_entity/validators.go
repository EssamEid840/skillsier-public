package initial_entity

import (
	"regexp"
	"strings"

	"skillsier.dev/apps/be/users-be/internal/domain/initial_entity"
)

var (
	// nameRegex validates entity names (alphanumeric, spaces, hyphens, underscores)
	nameRegex = regexp.MustCompile(`^[a-zA-Z0-9\s\-_]+$`)
	
	// reservedNames are names that cannot be used
	reservedNames = map[string]bool{
		"admin":     true,
		"system":    true,
		"root":      true,
		"null":      true,
		"undefined": true,
	}
)

// ValidateName validates an entity name
func ValidateName(name string) error {
	// Check length
	if len(name) < 3 {
		return initial_entity.ErrNameTooShort
	}
	if len(name) > 255 {
		return initial_entity.ErrNameTooLong
	}

	// Check format
	if !nameRegex.MatchString(name) {
		return initial_entity.NewValidationError("name", "must contain only alphanumeric characters, spaces, hyphens, or underscores")
	}

	// Check for reserved names
	normalized := strings.ToLower(strings.TrimSpace(name))
	if reservedNames[normalized] {
		return initial_entity.NewValidationError("name", "is a reserved name and cannot be used")
	}

	// Check for profanity or inappropriate content (simplified check)
	if containsProfanity(normalized) {
		return initial_entity.NewValidationError("name", "contains inappropriate content")
	}

	return nil
}

// ValidateDescription validates an entity description
func ValidateDescription(description string) error {
	if len(description) > 1000 {
		return initial_entity.NewValidationError("description", "must not exceed 1000 characters")
	}

	// Check for inappropriate content
	if containsProfanity(strings.ToLower(description)) {
		return initial_entity.NewValidationError("description", "contains inappropriate content")
	}

	return nil
}

// ValidateTags validates entity tags
func ValidateTags(tags []string) error {
	if len(tags) > 20 {
		return initial_entity.NewValidationError("tags", "cannot exceed 20 tags")
	}

	seen := make(map[string]bool)
	for _, tag := range tags {
		// Check length
		if len(tag) < 1 || len(tag) > 50 {
			return initial_entity.NewValidationError("tags", "each tag must be between 1 and 50 characters")
		}

		// Check for duplicates
		normalized := strings.ToLower(strings.TrimSpace(tag))
		if seen[normalized] {
			return initial_entity.NewValidationError("tags", "contains duplicate tags")
		}
		seen[normalized] = true

		// Check format (alphanumeric and hyphens only)
		if !regexp.MustCompile(`^[a-zA-Z0-9\-]+$`).MatchString(tag) {
			return initial_entity.NewValidationError("tags", "tags must contain only alphanumeric characters and hyphens")
		}
	}

	return nil
}

// ValidateProperties validates entity properties
func ValidateProperties(properties map[string]string) error {
	if len(properties) > 50 {
		return initial_entity.NewValidationError("properties", "cannot exceed 50 properties")
	}

	for key, value := range properties {
		// Check key length
		if len(key) < 1 || len(key) > 100 {
			return initial_entity.NewValidationError("properties", "property keys must be between 1 and 100 characters")
		}

		// Check value length
		if len(value) > 500 {
			return initial_entity.NewValidationError("properties", "property values must not exceed 500 characters")
		}

		// Check key format (alphanumeric, underscore, hyphen, dot)
		if !regexp.MustCompile(`^[a-zA-Z0-9_\-\.]+$`).MatchString(key) {
			return initial_entity.NewValidationError("properties", "property keys must contain only alphanumeric characters, underscores, hyphens, or dots")
		}
	}

	return nil
}

// containsProfanity checks if text contains profanity (simplified implementation)
// In production, use a proper profanity filter library
func containsProfanity(text string) bool {
	// Simplified profanity list (expand in production)
	profanityList := []string{
		"badword1",
		"badword2",
		// Add more as needed
	}

	for _, word := range profanityList {
		if strings.Contains(text, word) {
			return true
		}
	}

	return false
}

// ValidateStatusTransition validates if a status transition is allowed
func ValidateStatusTransition(currentStatus, newStatus initial_entity.Status) error {
	// Define allowed transitions
	allowedTransitions := map[initial_entity.Status][]initial_entity.Status{
		initial_entity.StatusActive:   {initial_entity.StatusInactive, initial_entity.StatusArchived},
		initial_entity.StatusInactive: {initial_entity.StatusActive, initial_entity.StatusArchived},
		initial_entity.StatusArchived: {}, // Cannot transition from archived
	}

	allowed, exists := allowedTransitions[currentStatus]
	if !exists {
		return initial_entity.ErrInvalidStatus
	}

	for _, s := range allowed {
		if s == newStatus {
			return nil
		}
	}

	return initial_entity.ErrInvalidStatusTransition
}