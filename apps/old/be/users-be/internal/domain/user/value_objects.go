// internal/domain/user/value_objects.go
package user

import (
	"fmt"
	"regexp"
	"strings"
)

// ============================================================================
// EMAIL VALUE OBJECT
// ============================================================================

// Email represents a validated email address
type Email struct {
	Value       string `gorm:"type:varchar(255);not null;uniqueIndex" json:"value"`
	Verified    bool   `gorm:"default:false" json:"verified"`
	VerifiedAt  *int64 `gorm:"type:bigint" json:"verified_at,omitempty"`
	Primary     bool   `gorm:"default:true" json:"primary"`
}

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

// NewEmail creates a new Email value object with validation
func NewEmail(email string) (Email, error) {
	email = strings.TrimSpace(strings.ToLower(email))
	
	if email == "" {
		return Email{}, ErrEmailRequired
	}
	
	if len(email) > 255 {
		return Email{}, ErrEmailTooLong
	}
	
	if !emailRegex.MatchString(email) {
		return Email{}, ErrInvalidEmailFormat
	}
	
	return Email{
		Value:    email,
		Verified: false,
		Primary:  true,
	}, nil
}

// Validate checks if the email is valid
func (e Email) Validate() error {
	if e.Value == "" {
		return ErrEmailRequired
	}
	if !emailRegex.MatchString(e.Value) {
		return ErrInvalidEmailFormat
	}
	return nil
}

// Normalize returns normalized email (lowercase, trimmed)
func (e Email) Normalize() string {
	return strings.TrimSpace(strings.ToLower(e.Value))
}

// Domain returns the domain part of the email
func (e Email) Domain() string {
	parts := strings.Split(e.Value, "@")
	if len(parts) == 2 {
		return parts[1]
	}
	return ""
}

// IsDisposable checks if email is from a disposable email provider
func (e Email) IsDisposable() bool {
	disposableDomains := map[string]bool{
		"tempmail.com": true, "guerrillamail.com": true, "10minutemail.com": true,
		"mailinator.com": true, "throwaway.email": true, "temp-mail.org": true,
	}
	return disposableDomains[e.Domain()]
}

// ============================================================================
// PHONE VALUE OBJECT
// ============================================================================

// Phone represents a validated phone number
type Phone struct {
	CountryCode    string `gorm:"type:varchar(5)" json:"country_code"`
	Number         string `gorm:"type:varchar(20)" json:"number"`
	FullNumber     string `gorm:"type:varchar(25)" json:"full_number"`
	Verified       bool   `gorm:"default:false" json:"verified"`
	VerifiedAt     *int64 `gorm:"type:bigint" json:"verified_at,omitempty"`
	Primary        bool   `gorm:"default:true" json:"primary"`
	Type           string `gorm:"type:varchar(20)" json:"type"` // mobile, landline, voip
}

// NewPhone creates a new Phone value object with validation
func NewPhone(countryCode, number string) (Phone, error) {
	countryCode = strings.TrimSpace(countryCode)
	number = strings.TrimSpace(number)
	
	// Remove non-numeric characters except +
	number = regexp.MustCompile(`[^\d+]`).ReplaceAllString(number, "")
	
	if countryCode == "" {
		return Phone{}, fmt.Errorf("country code is required")
	}
	
	if number == "" {
		return Phone{}, fmt.Errorf("phone number is required")
	}
	
	// Basic validation
	if len(number) < 7 || len(number) > 15 {
		return Phone{}, fmt.Errorf("invalid phone number length")
	}
	
	fullNumber := countryCode + number
	
	return Phone{
		CountryCode: countryCode,
		Number:      number,
		FullNumber:  fullNumber,
		Verified:    false,
		Primary:     true,
		Type:        "mobile",
	}, nil
}

// Validate checks if the phone is valid
func (p Phone) Validate() error {
	if p.CountryCode == "" {
		return fmt.Errorf("country code is required")
	}
	if p.Number == "" {
		return fmt.Errorf("phone number is required")
	}
	if len(p.Number) < 7 || len(p.Number) > 15 {
		return fmt.Errorf("invalid phone number length")
	}
	return nil
}

// Format returns formatted phone number (e.g., +1 (555) 123-4567)
func (p Phone) Format() string {
	if p.CountryCode == "+1" && len(p.Number) == 10 {
		return fmt.Sprintf("%s (%s) %s-%s", 
			p.CountryCode, 
			p.Number[0:3], 
			p.Number[3:6], 
			p.Number[6:10])
	}
	return p.FullNumber
}

// ============================================================================
// ADDRESS VALUE OBJECT
// ============================================================================

// Address represents a physical address
type Address struct {
	Street1    string  `gorm:"type:varchar(200)" json:"street1"`
	Street2    string  `gorm:"type:varchar(200)" json:"street2,omitempty"`
	City       string  `gorm:"type:varchar(100);index" json:"city"`
	State      string  `gorm:"type:varchar(100)" json:"state,omitempty"`
	PostalCode string  `gorm:"type:varchar(20)" json:"postal_code"`
	Country    string  `gorm:"type:varchar(100);index" json:"country"`
	CountryCode string `gorm:"type:varchar(3)" json:"country_code"` // ISO 3166-1 alpha-2
	Latitude   float64 `gorm:"type:decimal(10,8)" json:"latitude,omitempty"`
	Longitude  float64 `gorm:"type:decimal(11,8)" json:"longitude,omitempty"`
	Timezone   string  `gorm:"type:varchar(50)" json:"timezone,omitempty"`
	Verified   bool    `gorm:"default:false" json:"verified"`
	Primary    bool    `gorm:"default:true" json:"primary"`
	Type       string  `gorm:"type:varchar(20)" json:"type"` // home, office, billing, shipping
}

// NewAddress creates a new Address value object
func NewAddress(street1, city, country, countryCode string) (Address, error) {
	street1 = strings.TrimSpace(street1)
	city = strings.TrimSpace(city)
	country = strings.TrimSpace(country)
	countryCode = strings.ToUpper(strings.TrimSpace(countryCode))
	
	if city == "" {
		return Address{}, fmt.Errorf("city is required")
	}
	
	if country == "" {
		return Address{}, fmt.Errorf("country is required")
	}
	
	if countryCode != "" && len(countryCode) != 2 {
		return Address{}, fmt.Errorf("country code must be 2 characters (ISO 3166-1 alpha-2)")
	}
	
	return Address{
		Street1:     street1,
		City:        city,
		Country:     country,
		CountryCode: countryCode,
		Verified:    false,
		Primary:     true,
		Type:        "home",
	}, nil
}

// Validate checks if the address is valid
func (a Address) Validate() error {
	if a.City == "" {
		return fmt.Errorf("city is required")
	}
	if a.Country == "" {
		return fmt.Errorf("country is required")
	}
	if a.CountryCode != "" && len(a.CountryCode) != 2 {
		return fmt.Errorf("invalid country code format")
	}
	return nil
}

// FullAddress returns formatted full address string
func (a Address) FullAddress() string {
	parts := []string{}
	
	if a.Street1 != "" {
		parts = append(parts, a.Street1)
	}
	if a.Street2 != "" {
		parts = append(parts, a.Street2)
	}
	if a.City != "" {
		parts = append(parts, a.City)
	}
	if a.State != "" {
		parts = append(parts, a.State)
	}
	if a.PostalCode != "" {
		parts = append(parts, a.PostalCode)
	}
	if a.Country != "" {
		parts = append(parts, a.Country)
	}
	
	return strings.Join(parts, ", ")
}

// ============================================================================
// LOCATION VALUE OBJECT (For searchability and timezone)
// ============================================================================

// Location represents a geographic location with timezone
type Location struct {
	City        string  `gorm:"type:varchar(100);index" json:"city"`
	Country     string  `gorm:"type:varchar(100);index" json:"country"`
	CountryCode string  `gorm:"type:varchar(3);index" json:"country_code"` // ISO 3166-1 alpha-2
	Timezone    string  `gorm:"type:varchar(50)" json:"timezone"`
	Latitude    float64 `gorm:"type:decimal(10,8)" json:"latitude,omitempty"`
	Longitude   float64 `gorm:"type:decimal(11,8)" json:"longitude,omitempty"`
}

// NewLocation creates a new Location value object
func NewLocation(city, country, countryCode, timezone string) (Location, error) {
	city = strings.TrimSpace(city)
	country = strings.TrimSpace(country)
	countryCode = strings.ToUpper(strings.TrimSpace(countryCode))
	timezone = strings.TrimSpace(timezone)
	
	if city == "" {
		return Location{}, fmt.Errorf("city is required")
	}
	
	if country == "" {
		return Location{}, fmt.Errorf("country is required")
	}
	
	if countryCode != "" && len(countryCode) != 2 {
		return Location{}, fmt.Errorf("country code must be 2 characters")
	}
	
	return Location{
		City:        city,
		Country:     country,
		CountryCode: countryCode,
		Timezone:    timezone,
	}, nil
}

// Validate checks if the location is valid
func (l Location) Validate() error {
	if l.City == "" {
		return fmt.Errorf("city is required")
	}
	if l.Country == "" {
		return fmt.Errorf("country is required")
	}
	return nil
}

// FullLocation returns formatted location string
func (l Location) FullLocation() string {
	if l.City != "" && l.Country != "" {
		return fmt.Sprintf("%s, %s", l.City, l.Country)
	}
	if l.City != "" {
		return l.City
	}
	return l.Country
}