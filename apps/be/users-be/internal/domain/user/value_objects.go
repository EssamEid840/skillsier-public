// internal/domain/user/value_objects.go
package user

import (
    "fmt"
    "regexp"
    "strings"
)

// Email value object with validation
type Email struct {
    Value string
}

func NewEmail(email string) (Email, error) {
    email = strings.TrimSpace(strings.ToLower(email))
    
    // RFC 5322 compliant regex (simplified)
    emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
    if !emailRegex.MatchString(email) {
        return Email{}, ErrInvalidEmail
    }
    
    // Additional validation
    if len(email) > 255 {
        return Email{}, ErrInvalidEmail
    }
    
    // Block disposable email domains
    disposableDomains := []string{"tempmail.com", "throwaway.email", "guerrillamail.com"}
    for _, domain := range disposableDomains {
        if strings.HasSuffix(email, "@"+domain) {
            return Email{}, fmt.Errorf("disposable email addresses not allowed")
        }
    }
    
    return Email{Value: email}, nil
}

// Phone value object with validation
type Phone struct {
    CountryCode string
    Number      string
    FullNumber  string
}

func NewPhone(countryCode, number string) (Phone, error) {
    countryCode = strings.TrimSpace(countryCode)
    number = strings.TrimSpace(strings.ReplaceAll(number, " ", ""))
    
    // Remove common formatting
    number = strings.ReplaceAll(number, "-", "")
    number = strings.ReplaceAll(number, "(", "")
    number = strings.ReplaceAll(number, ")", "")
    
    // Validate country code
    if !strings.HasPrefix(countryCode, "+") {
        countryCode = "+" + countryCode
    }
    
    // E.164 format validation
    phoneRegex := regexp.MustCompile(`^\+?[1-9]\d{1,14}$`)
    fullNumber := countryCode + number
    if !phoneRegex.MatchString(fullNumber) {
        return Phone{}, ErrInvalidPhone
    }
    
    return Phone{
        CountryCode: countryCode,
        Number:      number,
        FullNumber:  fullNumber,
    }, nil
}

// Address value object
type Address struct {
    Street1    string
    Street2    string
    City       string
    State      string
    PostalCode string
    Country    string // ISO 3166-1 alpha-2
}

func (a Address) IsComplete() bool {
    return a.Street1 != "" && a.City != "" && a.Country != ""
}

func (a Address) String() string {
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

// ReferralCode value object
type ReferralCode struct {
    Code string
}

func GenerateReferralCode(username string) ReferralCode {
    // Simple implementation - in production, use more sophisticated algorithm
    code := strings.ToUpper(username[:min(4, len(username))])
    code += fmt.Sprintf("%06d", time.Now().Unix()%1000000)
    return ReferralCode{Code: code}
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}