package initial_entity

// Status represents the lifecycle status of an InitialEntity
type Status string

const (
	// StatusActive indicates the entity is active and operational
	StatusActive Status = "active"
	
	// StatusInactive indicates the entity is temporarily inactive
	StatusInactive Status = "inactive"
	
	// StatusArchived indicates the entity is archived and read-only
	StatusArchived Status = "archived"
)

// AllStatuses returns all valid statuses
func AllStatuses() []Status {
	return []Status{
		StatusActive,
		StatusInactive,
		StatusArchived,
	}
}

// IsValid checks if the status is a valid value
func (s Status) IsValid() bool {
	switch s {
	case StatusActive, StatusInactive, StatusArchived:
		return true
	default:
		return false
	}
}

// String returns the string representation of the status
func (s Status) String() string {
	return string(s)
}

// ParseStatus parses a string into a Status
func ParseStatus(s string) (Status, error) {
	status := Status(s)
	if !status.IsValid() {
		return "", ErrInvalidStatus
	}
	return status, nil
}

// MarshalJSON implements json.Marshaler
func (s Status) MarshalJSON() ([]byte, error) {
	return []byte(`"` + s.String() + `"`), nil
}

// UnmarshalJSON implements json.Unmarshaler
func (s *Status) UnmarshalJSON(data []byte) error {
	// Remove quotes
	str := string(data)
	if len(str) >= 2 && str[0] == '"' && str[str[len(str)-1]] == '"' {
		str = str[1 : len(str)-1]
	}
	
	status, err := ParseStatus(str)
	if err != nil {
		return err
	}
	
	*s = status
	return nil
}