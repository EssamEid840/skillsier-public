// internal/application/user/validators.go
package user

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"
	
	userDomain "users-be/internal/domain/user"
)

// ============================================================================
// EMAIL VALIDATION
// ============================================================================

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

// ValidateEmail validates email format
func ValidateEmail(email string) error {
	email = strings.TrimSpace(strings.ToLower(email))
	
	if email == "" {
		return userDomain.ErrEmailRequired
	}
	
	if len(email) > 255 {
		return userDomain.ErrEmailTooLong
	}
	
	if !emailRegex.MatchString(email) {
		return userDomain.ErrInvalidEmailFormat
	}
	
	// Check for disposable email domains
	if isDisposableEmail(email) {
		return fmt.Errorf("disposable email addresses are not allowed")
	}
	
	return nil
}

// isDisposableEmail checks if email is from disposable provider
func isDisposableEmail(email string) bool {
	disposableDomains := map[string]bool{
		"tempmail.com": true, "guerrillamail.com": true, "10minutemail.com": true,
		"mailinator.com": true, "throwaway.email": true, "temp-mail.org": true,
		"maildrop.cc": true, "getnada.com": true, "trashmail.com": true,
	}
	
	parts := strings.Split(email, "@")
	if len(parts) != 2 {
		return false
	}
	
	domain := strings.ToLower(parts[1])
	return disposableDomains[domain]
}

// ============================================================================
// USERNAME VALIDATION
// ============================================================================

var usernameRegex = regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9_-]*$`)

// ValidateUsername validates username format
func ValidateUsername(username string) error {
	username = strings.TrimSpace(username)
	
	if username == "" {
		return userDomain.ErrUsernameRequired
	}
	
	if len(username) < 3 {
		return userDomain.ErrUsernameTooShort
	}
	
	if len(username) > 50 {
		return userDomain.ErrUsernameTooLong
	}
	
	// Must start with letter
	if !unicode.IsLetter(rune(username[0])) {
		return userDomain.ErrInvalidUsernameFormat
	}
	
	// Can only contain alphanumeric, underscore, hyphen
	if !usernameRegex.MatchString(username) {
		return userDomain.ErrInvalidUsernameFormat
	}
	
	// Cannot start or end with special characters
	if strings.HasPrefix(username, "_") || strings.HasPrefix(username, "-") ||
		strings.HasSuffix(username, "_") || strings.HasSuffix(username, "-") {
		return userDomain.ErrInvalidUsernameFormat
	}
	
	// Cannot have consecutive special characters
	if strings.Contains(username, "__") || strings.Contains(username, "--") ||
		strings.Contains(username, "_-") || strings.Contains(username, "-_") {
		return userDomain.ErrInvalidUsernameFormat
	}
	
	// Check for reserved usernames
	if isReservedUsername(username) {
		return fmt.Errorf("username is reserved")
	}
	
	return nil
}

// isReservedUsername checks if username is reserved
func isReservedUsername(username string) bool {
	reserved := map[string]bool{
		"admin": true, "root": true, "system": true, "support": true,
		"help": true, "api": true, "www": true, "ftp": true, "mail": true,
		"webmaster": true, "hostmaster": true, "postmaster": true,
		"info": true, "contact": true, "abuse": true, "security": true,
		"billing": true, "sales": true, "marketing": true,
		"noreply": true, "no-reply": true, "donotreply": true,
	}
	
	return reserved[strings.ToLower(username)]
}

// ============================================================================
// PASSWORD VALIDATION
// ============================================================================

// ValidatePassword validates password strength
func ValidatePassword(password string) error {
	if password == "" {
		return fmt.Errorf("password is required")
	}
	
	if len(password) < 8 {
		return fmt.Errorf("password must be at least 8 characters")
	}
	
	if len(password) > 128 {
		return fmt.Errorf("password must not exceed 128 characters")
	}
	
	var (
		hasUpper   bool
		hasLower   bool
		hasNumber  bool
		hasSpecial bool
	)
	
	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}
	
	if !hasUpper {
		return fmt.Errorf("password must contain at least one uppercase letter")
	}
	if !hasLower {
		return fmt.Errorf("password must contain at least one lowercase letter")
	}
	if !hasNumber {
		return fmt.Errorf("password must contain at least one number")
	}
	if !hasSpecial {
		return fmt.Errorf("password must contain at least one special character")
	}
	
	// Check for common passwords
	if isCommonPassword(password) {
		return fmt.Errorf("password is too common")
	}
	
	return nil
}

// isCommonPassword checks if password is commonly used
func isCommonPassword(password string) bool {
	commonPasswords := map[string]bool{
		"password": true, "Password1": true, "12345678": true,
		"password123": true, "Passw0rd": true, "qwerty123": true,
	}
	
	return commonPasswords[password]
}

// ============================================================================
// PHONE VALIDATION
// ============================================================================

var phoneRegex = regexp.MustCompile(`^\+?[1-9]\d{1,14}$`)

// ValidatePhone validates phone number
func ValidatePhone(countryCode, number string) error {
	if countryCode == "" {
		return fmt.Errorf("country code is required")
	}
	
	if number == "" {
		return fmt.Errorf("phone number is required")
	}
	
	// Remove all non-digits except leading +
	cleanNumber := regexp.MustCompile(`[^\d+]`).ReplaceAllString(number, "")
	
	if len(cleanNumber) < 7 || len(cleanNumber) > 15 {
		return fmt.Errorf("invalid phone number length")
	}
	
	// Validate E.164 format
	fullNumber := countryCode + cleanNumber
	if !phoneRegex.MatchString(fullNumber) {
		return fmt.Errorf("invalid phone number format")
	}
	
	return nil
}

// ============================================================================
// NAME VALIDATION
// ============================================================================

// ValidateName validates first/last name
func ValidateName(name string, fieldName string) error {
	name = strings.TrimSpace(name)
	
	if name == "" {
		return fmt.Errorf("%s is required", fieldName)
	}
	
	if len(name) < 2 {
		return fmt.Errorf("%s must be at least 2 characters", fieldName)
	}
	
	if len(name) > 100 {
		return fmt.Errorf("%s must not exceed 100 characters", fieldName)
	}
	
	// Should only contain letters, spaces, hyphens, apostrophes
	nameRegex := regexp.MustCompile(`^[a-zA-Z\s'-]+$`)
	if !nameRegex.MatchString(name) {
		return fmt.Errorf("%s contains invalid characters", fieldName)
	}
	
	// Should not start or end with special characters
	if strings.HasPrefix(name, "-") || strings.HasPrefix(name, "'") ||
		strings.HasSuffix(name, "-") || strings.HasSuffix(name, "'") {
		return fmt.Errorf("%s format is invalid", fieldName)
	}
	
	return nil
}

// ============================================================================
// BIO/TAGLINE VALIDATION
// ============================================================================

// ValidateBio validates bio content
func ValidateBio(bio string) error {
	bio = strings.TrimSpace(bio)
	
	if len(bio) > 5000 {
		return userDomain.ErrInvalidBioLength
	}
	
	// Check for suspicious patterns
	if containsSuspiciousContent(bio) {
		return fmt.Errorf("bio contains prohibited content")
	}
	
	return nil
}

// ValidateTagline validates tagline
func ValidateTagline(tagline string) error {
	tagline = strings.TrimSpace(tagline)
	
	if len(tagline) > 200 {
		return userDomain.ErrInvalidTaglineLength
	}
	
	if containsSuspiciousContent(tagline) {
		return fmt.Errorf("tagline contains prohibited content")
	}
	
	return nil
}

// containsSuspiciousContent checks for prohibited content
func containsSuspiciousContent(text string) bool {
	text = strings.ToLower(text)
	
	// Check for spam keywords
	spamKeywords := []string{
		"bitcoin", "crypto", "forex", "lottery", "prize",
		"click here", "buy now", "limited offer", "guarantee",
	}
	
	for _, keyword := range spamKeywords {
		if strings.Contains(text, keyword) {
			return true
		}
	}
	
	// Check for excessive URLs
	urlCount := strings.Count(text, "http://") + strings.Count(text, "https://")
	if urlCount > 3 {
		return true
	}
	
	// Check for excessive special characters
	specialCharCount := 0
	for _, char := range text {
		if unicode.IsPunct(char) || unicode.IsSymbol(char) {
			specialCharCount++
		}
	}
	
	if specialCharCount > len(text)/3 {
		return true
	}
	
	return false
}

// ============================================================================
// URL VALIDATION
// ============================================================================

var urlRegex = regexp.MustCompile(`^https?://[^\s/$.?#].[^\s]*$`)

// ValidateURL validates URL format
func ValidateURL(url string) error {
	url = strings.TrimSpace(url)
	
	if url == "" {
		return nil // URL is optional
	}
	
	if len(url) > 500 {
		return fmt.Errorf("URL too long (max 500 characters)")
	}
	
	if !urlRegex.MatchString(url) {
		return fmt.Errorf("invalid URL format")
	}
	
	// Must use HTTPS for security
	if !strings.HasPrefix(url, "https://") {
		return fmt.Errorf("URL must use HTTPS")
	}
	
	return nil
}

// ============================================================================
// COUNTRY CODE VALIDATION
// ============================================================================

// ValidateCountryCode validates ISO 3166-1 alpha-2 country code
func ValidateCountryCode(code string) error {
	code = strings.ToUpper(strings.TrimSpace(code))
	
	if code == "" {
		return fmt.Errorf("country code is required")
	}
	
	if len(code) != 2 {
		return fmt.Errorf("country code must be 2 characters (ISO 3166-1 alpha-2)")
	}
	
	// Validate it's only letters
	for _, char := range code {
		if !unicode.IsLetter(char) {
			return fmt.Errorf("country code must contain only letters")
		}
	}
	
	return nil
}

// ============================================================================
// USER TYPE VALIDATION
// ============================================================================

// ValidateUserType validates user type
func ValidateUserType(userType string) error {
	ut := userDomain.UserType(userType)
	if !ut.Valid() {
		return userDomain.ErrInvalidUserType
	}
	return nil
}

// ============================================================================
// RATING VALIDATION
// ============================================================================

// ValidateRating validates rating value
func ValidateRating(rating float64) error {
	if rating < 0 || rating > 5 {
		return userDomain.ErrInvalidRating
	}
	return nil
}

// ============================================================================
// DTO VALIDATORS
// ============================================================================

// ValidateCreateUserDTO validates CreateUserDTO
func ValidateCreateUserDTO(dto *CreateUserDTO) error {
	errors := userDomain.NewValidationErrors()
	
	if err := ValidateEmail(dto.Email); err != nil {
		errors.Add("email", err.Error(), dto.Email)
	}
	
	if err := ValidateUsername(dto.Username); err != nil {
		errors.Add("username", err.Error(), dto.Username)
	}
	
	if err := ValidateName(dto.FirstName, "first_name"); err != nil {
		errors.Add("first_name", err.Error(), dto.FirstName)
	}
	
	if err := ValidateName(dto.LastName, "last_name"); err != nil {
		errors.Add("last_name", err.Error(), dto.LastName)
	}
	
	if err := ValidateUserType(dto.UserType); err != nil {
		errors.Add("user_type", err.Error(), dto.UserType)
	}
	
	if dto.PhoneNumber != "" && dto.PhoneCountryCode != "" {
		if err := ValidatePhone(dto.PhoneCountryCode, dto.PhoneNumber); err != nil {
			errors.Add("phone", err.Error(), dto.PhoneNumber)
		}
	}
	
	if dto.Bio != "" {
		if err := ValidateBio(dto.Bio); err != nil {
			errors.Add("bio", err.Error(), dto.Bio)
		}
	}
	
	if dto.Tagline != "" {
		if err := ValidateTagline(dto.Tagline); err != nil {
			errors.Add("tagline", err.Error(), dto.Tagline)
		}
	}
	
	if dto.CountryCode != "" {
		if err := ValidateCountryCode(dto.CountryCode); err != nil {
			errors.Add("country_code", err.Error(), dto.CountryCode)
		}
	}
	
	if errors.HasErrors() {
		return errors
	}
	
	return nil
}

// ValidateUpdateUserDTO validates UpdateUserDTO
func ValidateUpdateUserDTO(dto *UpdateUserDTO) error {
	errors := userDomain.NewValidationErrors()
	
	if dto.FirstName != nil {
		if err := ValidateName(*dto.FirstName, "first_name"); err != nil {
			errors.Add("first_name", err.Error(), *dto.FirstName)
		}
	}
	
	if dto.LastName != nil {
		if err := ValidateName(*dto.LastName, "last_name"); err != nil {
			errors.Add("last_name", err.Error(), *dto.LastName)
		}
	}
	
	if dto.Bio != nil {
		if err := ValidateBio(*dto.Bio); err != nil {
			errors.Add("bio", err.Error(), *dto.Bio)
		}
	}
	
	if dto.Tagline != nil {
		if err := ValidateTagline(*dto.Tagline); err != nil {
			errors.Add("tagline", err.Error(), *dto.Tagline)
		}
	}
	
	if dto.Website != nil && *dto.Website != "" {
		if err := ValidateURL(*dto.Website); err != nil {
			errors.Add("website", err.Error(), *dto.Website)
		}
	}
	
	if dto.PhoneNumber != nil && dto.PhoneCountryCode != nil {
		if err := ValidatePhone(*dto.PhoneCountryCode, *dto.PhoneNumber); err != nil {
			errors.Add("phone", err.Error(), *dto.PhoneNumber)
		}
	}
	
	if dto.CountryCode != nil && *dto.CountryCode != "" {
		if err := ValidateCountryCode(*dto.CountryCode); err != nil {
			errors.Add("country_code", err.Error(), *dto.CountryCode)
		}
	}
	
	if errors.HasErrors() {
		return errors
	}
	
	return nil
}

// ValidateUpdateProfileDTO validates UpdateProfileDTO
func ValidateUpdateProfileDTO(dto *UpdateProfileDTO) error {
	errors := userDomain.NewValidationErrors()
	
	if dto.Bio != "" {
		if err := ValidateBio(dto.Bio); err != nil {
			errors.Add("bio", err.Error(), dto.Bio)
		}
	}
	
	if dto.Tagline != "" {
		if err := ValidateTagline(dto.Tagline); err != nil {
			errors.Add("tagline", err.Error(), dto.Tagline)
		}
	}
	
	if dto.Title != "" {
		if len(dto.Title) > 200 {
			errors.Add("title", "Title too long (max 200 characters)", dto.Title)
		}
	}
	
	if errors.HasErrors() {
		return errors
	}
	
	return nil
}

// ValidateUpdateAvailabilityDTO validates UpdateAvailabilityDTO
func ValidateUpdateAvailabilityDTO(dto *UpdateAvailabilityDTO) error {
	errors := userDomain.NewValidationErrors()
	
	status := userDomain.AvailabilityStatus(dto.Status)
	if !status.Valid() {
		errors.Add("status", "Invalid availability status", dto.Status)
	}
	
	if dto.HoursPerWeek < 0 || dto.HoursPerWeek > 168 {
		errors.Add("hours_per_week", "Hours per week must be between 0 and 168", dto.HoursPerWeek)
	}
	
	if errors.HasErrors() {
		return errors
	}
	
	return nil
}

// ValidateUpdateSettingsDTO validates UpdateSettingsDTO
func ValidateUpdateSettingsDTO(dto *UpdateSettingsDTO) error {
	errors := userDomain.NewValidationErrors()
	
	if dto.ProfileVisibility != nil {
		visibility := userDomain.ProfileVisibility(*dto.ProfileVisibility)
		if !visibility.Valid() {
			errors.Add("profile_visibility", "Invalid profile visibility", *dto.ProfileVisibility)
		}
	}
	
	if errors.HasErrors() {
		return errors
	}
	
	return nil
}

// ============================================================================
// BATCH VALIDATION
// ============================================================================

// ValidateBatchSize validates batch operation size
func ValidateBatchSize(size int) error {
	if size <= 0 {
		return userDomain.ErrEmptyBatchOperation
	}
	
	if size > 1000 {
		return userDomain.ErrBatchSizeTooLarge
	}
	
	return nil
}

// ValidateUserIDs validates array of user IDs
func ValidateUserIDs(ids []string) error {
	if len(ids) == 0 {
		return userDomain.ErrEmptyBatchOperation
	}
	
	for i, id := range ids {
		if id == "" {
			return fmt.Errorf("user ID at index %d is empty", i)
		}
	}
	
	return nil
}

// ============================================================================
// SANITIZATION HELPERS
// ============================================================================

// SanitizeUsername sanitizes username
func SanitizeUsername(username string) string {
	return strings.TrimSpace(strings.ToLower(username))
}

// SanitizeEmail sanitizes email
func SanitizeEmail(email string) string {
	return strings.TrimSpace(strings.ToLower(email))
}

// SanitizeName sanitizes name (first/last name)
func SanitizeName(name string) string {
	name = strings.TrimSpace(name)
	// Capitalize first letter of each word
	words := strings.Fields(name)
	for i, word := range words {
		if len(word) > 0 {
			words[i] = strings.ToUpper(string(word[0])) + strings.ToLower(word[1:])
		}
	}
	return strings.Join(words, " ")
}

// SanitizeText sanitizes general text input
func SanitizeText(text string) string {
	text = strings.TrimSpace(text)
	// Remove excessive whitespace
	text = regexp.MustCompile(`\s+`).ReplaceAllString(text, " ")
	return text
}

// StripHTML strips HTML tags from text
func StripHTML(text string) string {
	htmlTagRegex := regexp.MustCompile(`<[^>]*>`)
	return htmlTagRegex.ReplaceAllString(text, "")
}

// ============================================================================
// VALIDATION HELPER FUNCTIONS
// ============================================================================

// IsValidUUID checks if string is valid UUID
func IsValidUUID(id string) bool {
	uuidRegex := regexp.MustCompile(`^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}// internal/application/user/validators.go
package user

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"
	
	userDomain "users-be/internal/domain/user"
)

// ============================================================================
// EMAIL VALIDATION
// ============================================================================

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

// ValidateEmail validates email format
func ValidateEmail(email string) error {
	email = strings.TrimSpace(strings.ToLower(email))
	
	if email == "" {
		return userDomain.ErrEmailRequired
	}
	
	if len(email) > 255 {
		return userDomain.ErrEmailTooLong
	}
	
	if !emailRegex.MatchString(email) {
		return userDomain.ErrInvalidEmailFormat
	}
	
	// Check for disposable email domains
	if isDisposableEmail(email) {
		return fmt.Errorf("disposable email addresses are not allowed")
	}
	
	return nil
}

// isDisposableEmail checks if email is from disposable provider
func isDisposableEmail(email string) bool {
	disposableDomains := map[string]bool{
		"tempmail.com": true, "guerrillamail.com": true, "10minutemail.com": true,
		"mailinator.com": true, "throwaway.email": true, "temp-mail.org": true,
		"maildrop.cc": true, "getnada.com": true, "trashmail.com": true,
	}
	
	parts := strings.Split(email, "@")
	if len(parts) != 2 {
		return false
	}
	
	domain := strings.ToLower(parts[1])
	return disposableDomains[domain]
}

// ============================================================================
// USERNAME VALIDATION
// ============================================================================

var usernameRegex = regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9_-]*$`)

// ValidateUsername validates username format
func ValidateUsername(username string) error {
	username = strings.TrimSpace(username)
	
	if username == "" {
		return userDomain.ErrUsernameRequired
	}
	
	if len(username) < 3 {
		return userDomain.ErrUsernameTooShort
	}
	
	if len(username) > 50 {
		return userDomain.ErrUsernameTooLong
	}
	
	// Must start with letter
	if !unicode.IsLetter(rune(username[0])) {
		return userDomain.ErrInvalidUsernameFormat
	}
	
	// Can only contain alphanumeric, underscore, hyphen
	if !usernameRegex.MatchString(username) {
		return userDomain.ErrInvalidUsernameFormat
	}
	
	// Cannot start or end with special characters
	if strings.HasPrefix(username, "_") || strings.HasPrefix(username, "-") ||
		strings.HasSuffix(username, "_") || strings.HasSuffix(username, "-") {
		return userDomain.ErrInvalidUsernameFormat
	}
	
	// Cannot have consecutive special characters
	if strings.Contains(username, "__") || strings.Contains(username, "--") ||
		strings.Contains(username, "_-") || strings.Contains(username, "-_") {
		return userDomain.ErrInvalidUsernameFormat
	}
	
	// Check for reserved usernames
	if isReservedUsername(username) {
		return fmt.Errorf("username is reserved")
	}
	
	return nil
}

// isReservedUsername checks if username is reserved
func isReservedUsername(username string) bool {
	reserved := map[string]bool{
		"admin": true, "root": true, "system": true, "support": true,
		"help": true, "api": true, "www": true, "ftp": true, "mail": true,
		"webmaster": true, "hostmaster": true, "postmaster": true,
		"info": true, "contact": true, "abuse": true, "security": true,
		"billing": true, "sales": true, "marketing": true,
		"noreply": true, "no-reply": true, "donotreply": true,
	}
	
	return reserved[strings.ToLower(username)]
}

// ============================================================================
// PASSWORD VALIDATION
// ============================================================================

// ValidatePassword validates password strength
func ValidatePassword(password string) error {
	if password == "" {
		return fmt.Errorf("password is required")
	}
	
	if len(password) < 8 {
		return fmt.Errorf("password must be at least 8 characters")
	}
	
	if len(password) > 128 {
		return fmt.Errorf("password must not exceed 128 characters")
	}
	
	var (
		hasUpper   bool
		hasLower   bool
		hasNumber  bool
		hasSpecial bool
	)
	
	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}
	
	if !hasUpper {
		return fmt.Errorf("password must contain at least one uppercase letter")
	}
	if !hasLower {
		return fmt.Errorf("password must contain at least one lowercase letter")
	}
	if !hasNumber {
		return fmt.Errorf("password must contain at least one number")
	}
	if !hasSpecial {
		return fmt.Errorf("password must contain at least one special character")
	}
	
	// Check for common passwords
	if isCommonPassword(password) {
		return fmt.Errorf("password is too common")
	}
	
	return nil
}

// isCommonPassword checks if password is commonly used
func isCommonPassword(password string) bool {
	commonPasswords := map[string]bool{
		"password": true, "Password1": true, "12345678": true,
		"password123": true, "Passw0rd": true, "qwerty123": true,
	}
	
	return commonPasswords[password]
}

// ============================================================================
// PHONE VALIDATION
// ============================================================================

var phoneRegex = regexp.MustCompile(`^\+?[1-9]\d{1,14}$`)

// ValidatePhone validates phone number
func ValidatePhone(countryCode, number string) error {
	if countryCode == "" {
		return fmt.Errorf("country code is required")
	}
	
	if number == "" {
		return fmt.Errorf("phone number is required")
	}
	
	// Remove all non-digits except leading +
	cleanNumber := regexp.MustCompile(`[^\d+]`).ReplaceAllString(number, "")
	
	if len(cleanNumber) < 7 || len(cleanNumber) > 15 {
		return fmt.Errorf("invalid phone number length")
	}
	
	// Validate E.164 format
	fullNumber := countryCode + cleanNumber
	if !phoneRegex.MatchString(fullNumber) {
		return fmt.Errorf("invalid phone number format")
	}
	
	return nil
}

// ============================================================================
// NAME VALIDATION
// ============================================================================

// ValidateName validates first/last name
func ValidateName(name string, fieldName string) error {
	name = strings.TrimSpace(name)
	
	if name == "" {
		return fmt.Errorf("%s is required", fieldName)
	}
	
	if len(name) < 2 {
		return fmt.Errorf("%s must be at least 2 characters", fieldName)
	}
	
	if len(name) > 100 {
		return fmt.Errorf("%s must not exceed 100 characters", fieldName)
	}
	
	// Should only contain letters, spaces, hyphens, apostrophes
	nameRegex := regexp.MustCompile(`^[a-zA-Z\s'-]+$`)
	if !nameRegex.MatchString(name) {
		return fmt.Errorf("%s contains invalid characters", fieldName)
	}
	
	// Should not start or end with special characters
	if strings.HasPrefix(name, "-") || strings.HasPrefix(name, "'") ||
		strings.HasSuffix(name, "-") || strings.HasSuffix(name, "'") {
		return fmt.Errorf("%s format is invalid", fieldName)
	}
	
	return nil
}

// ============================================================================
// BIO/TAGLINE VALIDATION
// ============================================================================

// ValidateBio validates bio content
func ValidateBio(bio string) error {
	bio = strings.TrimSpace(bio)
	
	if len(bio) > 5000 {
		return userDomain.ErrInvalidBioLength
	}
	
	// Check for suspicious patterns
	if containsSuspiciousContent(bio) {
		return fmt.Errorf("bio contains prohibited content")
	}
	
	return nil
}

// ValidateTagline validates tagline
func ValidateTagline(tagline string) error {
	tagline = strings.TrimSpace(tagline)
	
	if len(tagline) > 200 {
		return userDomain.ErrInvalidTaglineLength
	}
	
	if containsSuspiciousContent(tagline) {
		return fmt.Errorf("tagline contains prohibited content")
	}
	
	return nil
}

// containsSuspiciousContent checks for prohibited content
func containsSuspiciousContent(text string) bool {
	text = strings.ToLower(text)
	
	// Check for spam keywords
	spamKeywords := []string{
		"bitcoin", "crypto", "forex", "lottery", "prize",
		"click here", "buy now", "limited offer", "guarantee",
	}
	
	for _, keyword := range spamKeywords {
		if strings.Contains(text, keyword) {
			return true
		}
	}
	
	// Check for excessive URLs
	urlCount := strings.Count(text, "http://") + strings.Count(text, "https://")
	if urlCount > 3 {
		return true
	}
	
	// Check for excessive special characters
	specialCharCount := 0
	for _, char := range text {
		if unicode.IsPunct(char) || unicode.IsSymbol(char) {
			specialCharCount++
		}
	}
	
	if specialCharCount > len(text)/3 {
		return true
	}
	
	return false
}

// ============================================================================
// URL VALIDATION
// ============================================================================

var urlRegex = regexp.MustCompile(`^https?://[^\s/$.?#].[^\s]*$`)

// ValidateURL validates URL format
func ValidateURL(url string) error {
	url = strings.TrimSpace(url)
	
	if url == "" {
		return nil // URL is optional
	}
	
	if len(url) > 500 {
		return fmt.Errorf("URL too long (max 500 characters)")
	}
	
	if !urlRegex.MatchString(url) {
		return fmt.Errorf("invalid URL format")
	}
	
	// Must use HTTPS for security
	if !strings.HasPrefix(url, "https://") {
		return fmt.Errorf("URL must use HTTPS")
	}
	
	return nil
}

// ============================================================================
// COUNTRY CODE VALIDATION
// ============================================================================

// ValidateCountryCode validates ISO 3166-1 alpha-2 country code
func ValidateCountryCode(code string) error {
	code = strings.ToUpper(strings.TrimSpace(code))
	
	if code == "" {
		return fmt.Errorf("country code is required")
	}
	
	if len(code) != 2 {
		return fmt.Errorf("country code must be 2 characters (ISO 3166-1 alpha-2)")
	}
	
	// Validate it's only letters
	for _, char := range code {
		if !unicode.IsLetter(char) {
			return fmt.Errorf("country code must contain only letters")
		}
	}
	
	return nil
}

// ============================================================================
// USER TYPE VALIDATION
// ============================================================================

// ValidateUserType validates user type
func ValidateUserType(userType string) error {
	ut := userDomain.UserType(userType)
	if !ut.Valid() {
		return userDomain.ErrInvalidUserType
	}
	return nil
}

// ============================================================================
// RATING VALIDATION
// ============================================================================

// ValidateRating validates rating value
func ValidateRating(rating float64) error {
	if rating < 0 || rating > 5 {
		return userDomain.ErrInvalidRating
	}
	return nil
}

// ============================================================================
// DTO VALIDATORS
// ============================================================================

)
	return uuidRegex.MatchString(strings.ToLower(id))
}

// IsValidKeycloakID checks if string is valid Keycloak ID format
func IsValidKeycloakID(id string) bool {
	// Keycloak IDs are UUIDs
	return IsValidUUID(id)
}

// IsValidTimezone checks if timezone is valid
func IsValidTimezone(timezone string) bool {
	// Basic validation - just check if it's not empty and has reasonable length
	timezone = strings.TrimSpace(timezone)
	return timezone != "" && len(timezone) <= 100
}

// IsValidLanguageCode checks if language code is valid (ISO 639-1)
func IsValidLanguageCode(code string) bool {
	code = strings.ToLower(strings.TrimSpace(code))
	return len(code) == 2 && regexp.MustCompile(`^[a-z]{2}// internal/application/user/validators.go
package user

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"
	
	userDomain "users-be/internal/domain/user"
)

// ============================================================================
// EMAIL VALIDATION
// ============================================================================

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

// ValidateEmail validates email format
func ValidateEmail(email string) error {
	email = strings.TrimSpace(strings.ToLower(email))
	
	if email == "" {
		return userDomain.ErrEmailRequired
	}
	
	if len(email) > 255 {
		return userDomain.ErrEmailTooLong
	}
	
	if !emailRegex.MatchString(email) {
		return userDomain.ErrInvalidEmailFormat
	}
	
	// Check for disposable email domains
	if isDisposableEmail(email) {
		return fmt.Errorf("disposable email addresses are not allowed")
	}
	
	return nil
}

// isDisposableEmail checks if email is from disposable provider
func isDisposableEmail(email string) bool {
	disposableDomains := map[string]bool{
		"tempmail.com": true, "guerrillamail.com": true, "10minutemail.com": true,
		"mailinator.com": true, "throwaway.email": true, "temp-mail.org": true,
		"maildrop.cc": true, "getnada.com": true, "trashmail.com": true,
	}
	
	parts := strings.Split(email, "@")
	if len(parts) != 2 {
		return false
	}
	
	domain := strings.ToLower(parts[1])
	return disposableDomains[domain]
}

// ============================================================================
// USERNAME VALIDATION
// ============================================================================

var usernameRegex = regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9_-]*$`)

// ValidateUsername validates username format
func ValidateUsername(username string) error {
	username = strings.TrimSpace(username)
	
	if username == "" {
		return userDomain.ErrUsernameRequired
	}
	
	if len(username) < 3 {
		return userDomain.ErrUsernameTooShort
	}
	
	if len(username) > 50 {
		return userDomain.ErrUsernameTooLong
	}
	
	// Must start with letter
	if !unicode.IsLetter(rune(username[0])) {
		return userDomain.ErrInvalidUsernameFormat
	}
	
	// Can only contain alphanumeric, underscore, hyphen
	if !usernameRegex.MatchString(username) {
		return userDomain.ErrInvalidUsernameFormat
	}
	
	// Cannot start or end with special characters
	if strings.HasPrefix(username, "_") || strings.HasPrefix(username, "-") ||
		strings.HasSuffix(username, "_") || strings.HasSuffix(username, "-") {
		return userDomain.ErrInvalidUsernameFormat
	}
	
	// Cannot have consecutive special characters
	if strings.Contains(username, "__") || strings.Contains(username, "--") ||
		strings.Contains(username, "_-") || strings.Contains(username, "-_") {
		return userDomain.ErrInvalidUsernameFormat
	}
	
	// Check for reserved usernames
	if isReservedUsername(username) {
		return fmt.Errorf("username is reserved")
	}
	
	return nil
}

// isReservedUsername checks if username is reserved
func isReservedUsername(username string) bool {
	reserved := map[string]bool{
		"admin": true, "root": true, "system": true, "support": true,
		"help": true, "api": true, "www": true, "ftp": true, "mail": true,
		"webmaster": true, "hostmaster": true, "postmaster": true,
		"info": true, "contact": true, "abuse": true, "security": true,
		"billing": true, "sales": true, "marketing": true,
		"noreply": true, "no-reply": true, "donotreply": true,
	}
	
	return reserved[strings.ToLower(username)]
}

// ============================================================================
// PASSWORD VALIDATION
// ============================================================================

// ValidatePassword validates password strength
func ValidatePassword(password string) error {
	if password == "" {
		return fmt.Errorf("password is required")
	}
	
	if len(password) < 8 {
		return fmt.Errorf("password must be at least 8 characters")
	}
	
	if len(password) > 128 {
		return fmt.Errorf("password must not exceed 128 characters")
	}
	
	var (
		hasUpper   bool
		hasLower   bool
		hasNumber  bool
		hasSpecial bool
	)
	
	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}
	
	if !hasUpper {
		return fmt.Errorf("password must contain at least one uppercase letter")
	}
	if !hasLower {
		return fmt.Errorf("password must contain at least one lowercase letter")
	}
	if !hasNumber {
		return fmt.Errorf("password must contain at least one number")
	}
	if !hasSpecial {
		return fmt.Errorf("password must contain at least one special character")
	}
	
	// Check for common passwords
	if isCommonPassword(password) {
		return fmt.Errorf("password is too common")
	}
	
	return nil
}

// isCommonPassword checks if password is commonly used
func isCommonPassword(password string) bool {
	commonPasswords := map[string]bool{
		"password": true, "Password1": true, "12345678": true,
		"password123": true, "Passw0rd": true, "qwerty123": true,
	}
	
	return commonPasswords[password]
}

// ============================================================================
// PHONE VALIDATION
// ============================================================================

var phoneRegex = regexp.MustCompile(`^\+?[1-9]\d{1,14}$`)

// ValidatePhone validates phone number
func ValidatePhone(countryCode, number string) error {
	if countryCode == "" {
		return fmt.Errorf("country code is required")
	}
	
	if number == "" {
		return fmt.Errorf("phone number is required")
	}
	
	// Remove all non-digits except leading +
	cleanNumber := regexp.MustCompile(`[^\d+]`).ReplaceAllString(number, "")
	
	if len(cleanNumber) < 7 || len(cleanNumber) > 15 {
		return fmt.Errorf("invalid phone number length")
	}
	
	// Validate E.164 format
	fullNumber := countryCode + cleanNumber
	if !phoneRegex.MatchString(fullNumber) {
		return fmt.Errorf("invalid phone number format")
	}
	
	return nil
}

// ============================================================================
// NAME VALIDATION
// ============================================================================

// ValidateName validates first/last name
func ValidateName(name string, fieldName string) error {
	name = strings.TrimSpace(name)
	
	if name == "" {
		return fmt.Errorf("%s is required", fieldName)
	}
	
	if len(name) < 2 {
		return fmt.Errorf("%s must be at least 2 characters", fieldName)
	}
	
	if len(name) > 100 {
		return fmt.Errorf("%s must not exceed 100 characters", fieldName)
	}
	
	// Should only contain letters, spaces, hyphens, apostrophes
	nameRegex := regexp.MustCompile(`^[a-zA-Z\s'-]+$`)
	if !nameRegex.MatchString(name) {
		return fmt.Errorf("%s contains invalid characters", fieldName)
	}
	
	// Should not start or end with special characters
	if strings.HasPrefix(name, "-") || strings.HasPrefix(name, "'") ||
		strings.HasSuffix(name, "-") || strings.HasSuffix(name, "'") {
		return fmt.Errorf("%s format is invalid", fieldName)
	}
	
	return nil
}

// ============================================================================
// BIO/TAGLINE VALIDATION
// ============================================================================

// ValidateBio validates bio content
func ValidateBio(bio string) error {
	bio = strings.TrimSpace(bio)
	
	if len(bio) > 5000 {
		return userDomain.ErrInvalidBioLength
	}
	
	// Check for suspicious patterns
	if containsSuspiciousContent(bio) {
		return fmt.Errorf("bio contains prohibited content")
	}
	
	return nil
}

// ValidateTagline validates tagline
func ValidateTagline(tagline string) error {
	tagline = strings.TrimSpace(tagline)
	
	if len(tagline) > 200 {
		return userDomain.ErrInvalidTaglineLength
	}
	
	if containsSuspiciousContent(tagline) {
		return fmt.Errorf("tagline contains prohibited content")
	}
	
	return nil
}

// containsSuspiciousContent checks for prohibited content
func containsSuspiciousContent(text string) bool {
	text = strings.ToLower(text)
	
	// Check for spam keywords
	spamKeywords := []string{
		"bitcoin", "crypto", "forex", "lottery", "prize",
		"click here", "buy now", "limited offer", "guarantee",
	}
	
	for _, keyword := range spamKeywords {
		if strings.Contains(text, keyword) {
			return true
		}
	}
	
	// Check for excessive URLs
	urlCount := strings.Count(text, "http://") + strings.Count(text, "https://")
	if urlCount > 3 {
		return true
	}
	
	// Check for excessive special characters
	specialCharCount := 0
	for _, char := range text {
		if unicode.IsPunct(char) || unicode.IsSymbol(char) {
			specialCharCount++
		}
	}
	
	if specialCharCount > len(text)/3 {
		return true
	}
	
	return false
}

// ============================================================================
// URL VALIDATION
// ============================================================================

var urlRegex = regexp.MustCompile(`^https?://[^\s/$.?#].[^\s]*$`)

// ValidateURL validates URL format
func ValidateURL(url string) error {
	url = strings.TrimSpace(url)
	
	if url == "" {
		return nil // URL is optional
	}
	
	if len(url) > 500 {
		return fmt.Errorf("URL too long (max 500 characters)")
	}
	
	if !urlRegex.MatchString(url) {
		return fmt.Errorf("invalid URL format")
	}
	
	// Must use HTTPS for security
	if !strings.HasPrefix(url, "https://") {
		return fmt.Errorf("URL must use HTTPS")
	}
	
	return nil
}

// ============================================================================
// COUNTRY CODE VALIDATION
// ============================================================================

// ValidateCountryCode validates ISO 3166-1 alpha-2 country code
func ValidateCountryCode(code string) error {
	code = strings.ToUpper(strings.TrimSpace(code))
	
	if code == "" {
		return fmt.Errorf("country code is required")
	}
	
	if len(code) != 2 {
		return fmt.Errorf("country code must be 2 characters (ISO 3166-1 alpha-2)")
	}
	
	// Validate it's only letters
	for _, char := range code {
		if !unicode.IsLetter(char) {
			return fmt.Errorf("country code must contain only letters")
		}
	}
	
	return nil
}

// ============================================================================
// USER TYPE VALIDATION
// ============================================================================

// ValidateUserType validates user type
func ValidateUserType(userType string) error {
	ut := userDomain.UserType(userType)
	if !ut.Valid() {
		return userDomain.ErrInvalidUserType
	}
	return nil
}

// ============================================================================
// RATING VALIDATION
// ============================================================================

// ValidateRating validates rating value
func ValidateRating(rating float64) error {
	if rating < 0 || rating > 5 {
		return userDomain.ErrInvalidRating
	}
	return nil
}

// ============================================================================
// DTO VALIDATORS
// ============================================================================

).MatchString(code)
}

// IsValidCurrencyCode checks if currency code is valid (ISO 4217)
func IsValidCurrencyCode(code string) bool {
	code = strings.ToUpper(strings.TrimSpace(code))
	return len(code) == 3 && regexp.MustCompile(`^[A-Z]{3}// internal/application/user/validators.go
package user

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"
	
	userDomain "users-be/internal/domain/user"
)

// ============================================================================
// EMAIL VALIDATION
// ============================================================================

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

// ValidateEmail validates email format
func ValidateEmail(email string) error {
	email = strings.TrimSpace(strings.ToLower(email))
	
	if email == "" {
		return userDomain.ErrEmailRequired
	}
	
	if len(email) > 255 {
		return userDomain.ErrEmailTooLong
	}
	
	if !emailRegex.MatchString(email) {
		return userDomain.ErrInvalidEmailFormat
	}
	
	// Check for disposable email domains
	if isDisposableEmail(email) {
		return fmt.Errorf("disposable email addresses are not allowed")
	}
	
	return nil
}

// isDisposableEmail checks if email is from disposable provider
func isDisposableEmail(email string) bool {
	disposableDomains := map[string]bool{
		"tempmail.com": true, "guerrillamail.com": true, "10minutemail.com": true,
		"mailinator.com": true, "throwaway.email": true, "temp-mail.org": true,
		"maildrop.cc": true, "getnada.com": true, "trashmail.com": true,
	}
	
	parts := strings.Split(email, "@")
	if len(parts) != 2 {
		return false
	}
	
	domain := strings.ToLower(parts[1])
	return disposableDomains[domain]
}

// ============================================================================
// USERNAME VALIDATION
// ============================================================================

var usernameRegex = regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9_-]*$`)

// ValidateUsername validates username format
func ValidateUsername(username string) error {
	username = strings.TrimSpace(username)
	
	if username == "" {
		return userDomain.ErrUsernameRequired
	}
	
	if len(username) < 3 {
		return userDomain.ErrUsernameTooShort
	}
	
	if len(username) > 50 {
		return userDomain.ErrUsernameTooLong
	}
	
	// Must start with letter
	if !unicode.IsLetter(rune(username[0])) {
		return userDomain.ErrInvalidUsernameFormat
	}
	
	// Can only contain alphanumeric, underscore, hyphen
	if !usernameRegex.MatchString(username) {
		return userDomain.ErrInvalidUsernameFormat
	}
	
	// Cannot start or end with special characters
	if strings.HasPrefix(username, "_") || strings.HasPrefix(username, "-") ||
		strings.HasSuffix(username, "_") || strings.HasSuffix(username, "-") {
		return userDomain.ErrInvalidUsernameFormat
	}
	
	// Cannot have consecutive special characters
	if strings.Contains(username, "__") || strings.Contains(username, "--") ||
		strings.Contains(username, "_-") || strings.Contains(username, "-_") {
		return userDomain.ErrInvalidUsernameFormat
	}
	
	// Check for reserved usernames
	if isReservedUsername(username) {
		return fmt.Errorf("username is reserved")
	}
	
	return nil
}

// isReservedUsername checks if username is reserved
func isReservedUsername(username string) bool {
	reserved := map[string]bool{
		"admin": true, "root": true, "system": true, "support": true,
		"help": true, "api": true, "www": true, "ftp": true, "mail": true,
		"webmaster": true, "hostmaster": true, "postmaster": true,
		"info": true, "contact": true, "abuse": true, "security": true,
		"billing": true, "sales": true, "marketing": true,
		"noreply": true, "no-reply": true, "donotreply": true,
	}
	
	return reserved[strings.ToLower(username)]
}

// ============================================================================
// PASSWORD VALIDATION
// ============================================================================

// ValidatePassword validates password strength
func ValidatePassword(password string) error {
	if password == "" {
		return fmt.Errorf("password is required")
	}
	
	if len(password) < 8 {
		return fmt.Errorf("password must be at least 8 characters")
	}
	
	if len(password) > 128 {
		return fmt.Errorf("password must not exceed 128 characters")
	}
	
	var (
		hasUpper   bool
		hasLower   bool
		hasNumber  bool
		hasSpecial bool
	)
	
	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}
	
	if !hasUpper {
		return fmt.Errorf("password must contain at least one uppercase letter")
	}
	if !hasLower {
		return fmt.Errorf("password must contain at least one lowercase letter")
	}
	if !hasNumber {
		return fmt.Errorf("password must contain at least one number")
	}
	if !hasSpecial {
		return fmt.Errorf("password must contain at least one special character")
	}
	
	// Check for common passwords
	if isCommonPassword(password) {
		return fmt.Errorf("password is too common")
	}
	
	return nil
}

// isCommonPassword checks if password is commonly used
func isCommonPassword(password string) bool {
	commonPasswords := map[string]bool{
		"password": true, "Password1": true, "12345678": true,
		"password123": true, "Passw0rd": true, "qwerty123": true,
	}
	
	return commonPasswords[password]
}

// ============================================================================
// PHONE VALIDATION
// ============================================================================

var phoneRegex = regexp.MustCompile(`^\+?[1-9]\d{1,14}$`)

// ValidatePhone validates phone number
func ValidatePhone(countryCode, number string) error {
	if countryCode == "" {
		return fmt.Errorf("country code is required")
	}
	
	if number == "" {
		return fmt.Errorf("phone number is required")
	}
	
	// Remove all non-digits except leading +
	cleanNumber := regexp.MustCompile(`[^\d+]`).ReplaceAllString(number, "")
	
	if len(cleanNumber) < 7 || len(cleanNumber) > 15 {
		return fmt.Errorf("invalid phone number length")
	}
	
	// Validate E.164 format
	fullNumber := countryCode + cleanNumber
	if !phoneRegex.MatchString(fullNumber) {
		return fmt.Errorf("invalid phone number format")
	}
	
	return nil
}

// ============================================================================
// NAME VALIDATION
// ============================================================================

// ValidateName validates first/last name
func ValidateName(name string, fieldName string) error {
	name = strings.TrimSpace(name)
	
	if name == "" {
		return fmt.Errorf("%s is required", fieldName)
	}
	
	if len(name) < 2 {
		return fmt.Errorf("%s must be at least 2 characters", fieldName)
	}
	
	if len(name) > 100 {
		return fmt.Errorf("%s must not exceed 100 characters", fieldName)
	}
	
	// Should only contain letters, spaces, hyphens, apostrophes
	nameRegex := regexp.MustCompile(`^[a-zA-Z\s'-]+$`)
	if !nameRegex.MatchString(name) {
		return fmt.Errorf("%s contains invalid characters", fieldName)
	}
	
	// Should not start or end with special characters
	if strings.HasPrefix(name, "-") || strings.HasPrefix(name, "'") ||
		strings.HasSuffix(name, "-") || strings.HasSuffix(name, "'") {
		return fmt.Errorf("%s format is invalid", fieldName)
	}
	
	return nil
}

// ============================================================================
// BIO/TAGLINE VALIDATION
// ============================================================================

// ValidateBio validates bio content
func ValidateBio(bio string) error {
	bio = strings.TrimSpace(bio)
	
	if len(bio) > 5000 {
		return userDomain.ErrInvalidBioLength
	}
	
	// Check for suspicious patterns
	if containsSuspiciousContent(bio) {
		return fmt.Errorf("bio contains prohibited content")
	}
	
	return nil
}

// ValidateTagline validates tagline
func ValidateTagline(tagline string) error {
	tagline = strings.TrimSpace(tagline)
	
	if len(tagline) > 200 {
		return userDomain.ErrInvalidTaglineLength
	}
	
	if containsSuspiciousContent(tagline) {
		return fmt.Errorf("tagline contains prohibited content")
	}
	
	return nil
}

// containsSuspiciousContent checks for prohibited content
func containsSuspiciousContent(text string) bool {
	text = strings.ToLower(text)
	
	// Check for spam keywords
	spamKeywords := []string{
		"bitcoin", "crypto", "forex", "lottery", "prize",
		"click here", "buy now", "limited offer", "guarantee",
	}
	
	for _, keyword := range spamKeywords {
		if strings.Contains(text, keyword) {
			return true
		}
	}
	
	// Check for excessive URLs
	urlCount := strings.Count(text, "http://") + strings.Count(text, "https://")
	if urlCount > 3 {
		return true
	}
	
	// Check for excessive special characters
	specialCharCount := 0
	for _, char := range text {
		if unicode.IsPunct(char) || unicode.IsSymbol(char) {
			specialCharCount++
		}
	}
	
	if specialCharCount > len(text)/3 {
		return true
	}
	
	return false
}

// ============================================================================
// URL VALIDATION
// ============================================================================

var urlRegex = regexp.MustCompile(`^https?://[^\s/$.?#].[^\s]*$`)

// ValidateURL validates URL format
func ValidateURL(url string) error {
	url = strings.TrimSpace(url)
	
	if url == "" {
		return nil // URL is optional
	}
	
	if len(url) > 500 {
		return fmt.Errorf("URL too long (max 500 characters)")
	}
	
	if !urlRegex.MatchString(url) {
		return fmt.Errorf("invalid URL format")
	}
	
	// Must use HTTPS for security
	if !strings.HasPrefix(url, "https://") {
		return fmt.Errorf("URL must use HTTPS")
	}
	
	return nil
}

// ============================================================================
// COUNTRY CODE VALIDATION
// ============================================================================

// ValidateCountryCode validates ISO 3166-1 alpha-2 country code
func ValidateCountryCode(code string) error {
	code = strings.ToUpper(strings.TrimSpace(code))
	
	if code == "" {
		return fmt.Errorf("country code is required")
	}
	
	if len(code) != 2 {
		return fmt.Errorf("country code must be 2 characters (ISO 3166-1 alpha-2)")
	}
	
	// Validate it's only letters
	for _, char := range code {
		if !unicode.IsLetter(char) {
			return fmt.Errorf("country code must contain only letters")
		}
	}
	
	return nil
}

// ============================================================================
// USER TYPE VALIDATION
// ============================================================================

// ValidateUserType validates user type
func ValidateUserType(userType string) error {
	ut := userDomain.UserType(userType)
	if !ut.Valid() {
		return userDomain.ErrInvalidUserType
	}
	return nil
}

// ============================================================================
// RATING VALIDATION
// ============================================================================

// ValidateRating validates rating value
func ValidateRating(rating float64) error {
	if rating < 0 || rating > 5 {
		return userDomain.ErrInvalidRating
	}
	return nil
}

// ============================================================================
// DTO VALIDATORS
// ============================================================================

).MatchString(code)
}

// ============================================================================
// CONTENT MODERATION
// ============================================================================

// ContainsProfanity checks if text contains profanity
func ContainsProfanity(text string) bool {
	text = strings.ToLower(text)
	
	// Basic profanity list (should be expanded or use external service)
	profanityList := []string{
		// Add profanity words here
	}
	
	for _, word := range profanityList {
		if strings.Contains(text, word) {
			return true
		}
	}
	
	return false
}

// ContainsContactInfo checks if text contains contact information
func ContainsContactInfo(text string) bool {
	text = strings.ToLower(text)
	
	// Check for email pattern
	if emailRegex.MatchString(text) {
		return true
	}
	
	// Check for phone pattern
	phonePattern := regexp.MustCompile(`\d{3}[-.\s]?\d{3}[-.\s]?\d{4}`)
	if phonePattern.MatchString(text) {
		return true
	}
	
	// Check for skype, telegram, whatsapp mentions
	contactKeywords := []string{"skype:", "telegram:", "whatsapp:", "call me", "text me"}
	for _, keyword := range contactKeywords {
		if strings.Contains(text, keyword) {
			return true
		}
	}
	
	return false
}

// ============================================================================
// BUSINESS RULE VALIDATORS
// ============================================================================

// ValidateProfileCompleteness validates profile meets minimum requirements
func ValidateProfileCompleteness(completeness int) error {
	if completeness < 0 || completeness > 100 {
		return fmt.Errorf("profile completeness must be between 0 and 100")
	}
	return nil
}

// ValidateMinimumAge validates user meets minimum age requirement
func ValidateMinimumAge(age int) error {
	const minimumAge = 18
	
	if age < minimumAge {
		return fmt.Errorf("user must be at least %d years old", minimumAge)
	}
	
	return nil
}

// ValidateMaximumAge validates user age is reasonable
func ValidateMaximumAge(age int) error {
	const maximumAge = 120
	
	if age > maximumAge {
		return fmt.Errorf("invalid age")
	}
	
	return nil
}