package clients

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// KeycloakClient handles interactions with Keycloak
type KeycloakClient struct {
	baseURL      string
	realm        string
	clientID     string
	clientSecret string
	httpClient   *http.Client
	accessToken  string
	tokenExpiry  time.Time
}

// KeycloakConfig holds Keycloak configuration
type KeycloakConfig struct {
	URL          string
	Realm        string
	ClientID     string
	ClientSecret string
	Timeout      time.Duration
}

// TokenResponse represents Keycloak token response
type TokenResponse struct {
	AccessToken      string `json:"access_token"`
	ExpiresIn        int    `json:"expires_in"`
	RefreshExpiresIn int    `json:"refresh_expires_in"`
	RefreshToken     string `json:"refresh_token"`
	TokenType        string `json:"token_type"`
	Scope            string `json:"scope"`
}

// UserInfo represents basic user information from Keycloak
type UserInfo struct {
	Sub               string `json:"sub"`
	EmailVerified     bool   `json:"email_verified"`
	Name              string `json:"name"`
	PreferredUsername string `json:"preferred_username"`
	GivenName         string `json:"given_name"`
	FamilyName        string `json:"family_name"`
	Email             string `json:"email"`
}

// NewKeycloakClient creates a new Keycloak client
func NewKeycloakClient(cfg KeycloakConfig) *KeycloakClient {
	return &KeycloakClient{
		baseURL:      strings.TrimSuffix(cfg.URL, "/"),
		realm:        cfg.Realm,
		clientID:     cfg.ClientID,
		clientSecret: cfg.ClientSecret,
		httpClient: &http.Client{
			Timeout: cfg.Timeout,
		},
	}
}

// GetAccessToken retrieves an access token using client credentials grant
func (k *KeycloakClient) GetAccessToken(ctx context.Context) (string, error) {
	// Return cached token if still valid
	if k.accessToken != "" && time.Now().Before(k.tokenExpiry) {
		return k.accessToken, nil
	}

	tokenURL := fmt.Sprintf("%s/realms/%s/protocol/openid-connect/token", k.baseURL, k.realm)

	data := url.Values{}
	data.Set("grant_type", "client_credentials")
	data.Set("client_id", k.clientID)
	data.Set("client_secret", k.clientSecret)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, tokenURL, strings.NewReader(data.Encode()))
	if err != nil {
		return "", fmt.Errorf("failed to create token request: %w", err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := k.httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to get access token: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("keycloak returned status %d: %s", resp.StatusCode, string(body))
	}

	var tokenResp TokenResponse
	if err := json.NewDecoder(resp.Body).Decode(&tokenResp); err != nil {
		return "", fmt.Errorf("failed to decode token response: %w", err)
	}

	// Cache the token with a 30-second buffer before actual expiry
	k.accessToken = tokenResp.AccessToken
	k.tokenExpiry = time.Now().Add(time.Duration(tokenResp.ExpiresIn-30) * time.Second)

	return k.accessToken, nil
}

// GetUserInfo retrieves user information for a given access token
func (k *KeycloakClient) GetUserInfo(ctx context.Context, userAccessToken string) (*UserInfo, error) {
	userInfoURL := fmt.Sprintf("%s/realms/%s/protocol/openid-connect/userinfo", k.baseURL, k.realm)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, userInfoURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create userinfo request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+userAccessToken)

	resp, err := k.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to get user info: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("keycloak returned status %d: %s", resp.StatusCode, string(body))
	}

	var userInfo UserInfo
	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		return nil, fmt.Errorf("failed to decode userinfo response: %w", err)
	}

	return &userInfo, nil
}

// ValidateToken validates a JWT token against Keycloak
func (k *KeycloakClient) ValidateToken(ctx context.Context, token string) (bool, error) {
	introspectURL := fmt.Sprintf("%s/realms/%s/protocol/openid-connect/token/introspect", k.baseURL, k.realm)

	data := url.Values{}
	data.Set("token", token)
	data.Set("client_id", k.clientID)
	data.Set("client_secret", k.clientSecret)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, introspectURL, strings.NewReader(data.Encode()))
	if err != nil {
		return false, fmt.Errorf("failed to create introspection request: %w", err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := k.httpClient.Do(req)
	if err != nil {
		return false, fmt.Errorf("failed to introspect token: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return false, fmt.Errorf("keycloak returned status %d", resp.StatusCode)
	}

	var result struct {
		Active bool `json:"active"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, fmt.Errorf("failed to decode introspection response: %w", err)
	}

	return result.Active, nil
}

// GetJWKS retrieves the JSON Web Key Set from Keycloak
func (k *KeycloakClient) GetJWKS(ctx context.Context) ([]byte, error) {
	jwksURL := fmt.Sprintf("%s/realms/%s/protocol/openid-connect/certs", k.baseURL, k.realm)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, jwksURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create JWKS request: %w", err)
	}

	resp, err := k.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to get JWKS: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("keycloak returned status %d", resp.StatusCode)
	}

	jwks, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read JWKS response: %w", err)
	}

	return jwks, nil
}

// Logout revokes a refresh token
func (k *KeycloakClient) Logout(ctx context.Context, refreshToken string) error {
	logoutURL := fmt.Sprintf("%s/realms/%s/protocol/openid-connect/logout", k.baseURL, k.realm)

	data := url.Values{}
	data.Set("refresh_token", refreshToken)
	data.Set("client_id", k.clientID)
	data.Set("client_secret", k.clientSecret)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, logoutURL, strings.NewReader(data.Encode()))
	if err != nil {
		return fmt.Errorf("failed to create logout request: %w", err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := k.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to logout: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent && resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("keycloak returned status %d: %s", resp.StatusCode, string(body))
	}

	return nil
}