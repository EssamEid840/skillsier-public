package keycloak

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// Client provides access to Keycloak Admin API
type Client struct {
	config      *Config
	httpClient  *http.Client
	accessToken string
	tokenExpiry time.Time
}

// NewClient creates a new Keycloak admin API client
func NewClient(config *Config) *Client {
	return &Client{
		config: config,
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

// getAccessToken obtains an access token for the admin API
func (c *Client) getAccessToken(ctx context.Context) (string, error) {
	// Return cached token if still valid
	if c.accessToken != "" && time.Now().Before(c.tokenExpiry) {
		return c.accessToken, nil
	}

	// Request new token
	tokenURL := fmt.Sprintf("%s/realms/%s/protocol/openid-connect/token",
		c.config.BaseURL, c.config.Realm)

	data := url.Values{}
	data.Set("grant_type", "client_credentials")
	data.Set("client_id", c.config.ClientID)
	data.Set("client_secret", c.config.ClientSecret)

	req, err := http.NewRequestWithContext(ctx, "POST", tokenURL, strings.NewReader(data.Encode()))
	if err != nil {
		return "", fmt.Errorf("failed to create token request: %w", err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to request token: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("token request failed with status %d: %s", resp.StatusCode, string(body))
	}

	var tokenResp struct {
		AccessToken string `json:"access_token"`
		ExpiresIn   int    `json:"expires_in"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&tokenResp); err != nil {
		return "", fmt.Errorf("failed to decode token response: %w", err)
	}

	// Cache the token (with 30 second buffer before expiry)
	c.accessToken = tokenResp.AccessToken
	c.tokenExpiry = time.Now().Add(time.Duration(tokenResp.ExpiresIn-30) * time.Second)

	return c.accessToken, nil
}

// GetUser retrieves a user by ID from Keycloak
func (c *Client) GetUser(ctx context.Context, userID string) (map[string]interface{}, error) {
	token, err := c.getAccessToken(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get access token: %w", err)
	}

	userURL := fmt.Sprintf("%s/admin/realms/%s/users/%s",
		c.config.BaseURL, c.config.Realm, userID)

	req, err := http.NewRequestWithContext(ctx, "GET", userURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("get user failed with status %d: %s", resp.StatusCode, string(body))
	}

	var user map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return nil, fmt.Errorf("failed to decode user: %w", err)
	}

	return user, nil
}

// UpdateUser updates a user in Keycloak
func (c *Client) UpdateUser(ctx context.Context, userID string, updates map[string]interface{}) error {
	token, err := c.getAccessToken(ctx)
	if err != nil {
		return fmt.Errorf("failed to get access token: %w", err)
	}

	userURL := fmt.Sprintf("%s/admin/realms/%s/users/%s",
		c.config.BaseURL, c.config.Realm, userID)

	body, err := json.Marshal(updates)
	if err != nil {
		return fmt.Errorf("failed to marshal updates: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "PUT", userURL, bytes.NewReader(body))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent && resp.StatusCode != http.StatusOK {
		respBody, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("update user failed with status %d: %s", resp.StatusCode, string(respBody))
	}

	return nil
}

// AssignRole assigns a role to a user
func (c *Client) AssignRole(ctx context.Context, userID, roleName string) error {
	token, err := c.getAccessToken(ctx)
	if err != nil {
		return fmt.Errorf("failed to get access token: %w", err)
	}

	// First, get the role ID
	roleURL := fmt.Sprintf("%s/admin/realms/%s/roles/%s",
		c.config.BaseURL, c.config.Realm, roleName)

	req, err := http.NewRequestWithContext(ctx, "GET", roleURL, nil)
	if err != nil {
		return fmt.Errorf("failed to create role request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to get role: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("get role failed with status %d: %s", resp.StatusCode, string(body))
	}

	var role map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&role); err != nil {
		return fmt.Errorf("failed to decode role: %w", err)
	}

	// Assign the role to the user
	assignURL := fmt.Sprintf("%s/admin/realms/%s/users/%s/role-mappings/realm",
		c.config.BaseURL, c.config.Realm, userID)

	roleData := []map[string]interface{}{role}
	body, err := json.Marshal(roleData)
	if err != nil {
		return fmt.Errorf("failed to marshal role: %w", err)
	}

	req, err = http.NewRequestWithContext(ctx, "POST", assignURL, bytes.NewReader(body))
	if err != nil {
		return fmt.Errorf("failed to create assign request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	resp, err = c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to assign role: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent && resp.StatusCode != http.StatusOK {
		respBody, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("assign role failed with status %d: %s", resp.StatusCode, string(respBody))
	}

	return nil
}

// RemoveRole removes a role from a user
func (c *Client) RemoveRole(ctx context.Context, userID, roleName string) error {
	token, err := c.getAccessToken(ctx)
	if err != nil {
		return fmt.Errorf("failed to get access token: %w", err)
	}

	// Get the role
	roleURL := fmt.Sprintf("%s/admin/realms/%s/roles/%s",
		c.config.BaseURL, c.config.Realm, roleName)

	req, err := http.NewRequestWithContext(ctx, "GET", roleURL, nil)
	if err != nil {
		return fmt.Errorf("failed to create role request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to get role: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("get role failed with status %d: %s", resp.StatusCode, string(body))
	}

	var role map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&role); err != nil {
		return fmt.Errorf("failed to decode role: %w", err)
	}

	// Remove the role
	removeURL := fmt.Sprintf("%s/admin/realms/%s/users/%s/role-mappings/realm",
		c.config.BaseURL, c.config.Realm, userID)

	roleData := []map[string]interface{}{role}
	body, err := json.Marshal(roleData)
	if err != nil {
		return fmt.Errorf("failed to marshal role: %w", err)
	}

	req, err = http.NewRequestWithContext(ctx, "DELETE", removeURL, bytes.NewReader(body))
	if err != nil {
		return fmt.Errorf("failed to create remove request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	resp, err = c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to remove role: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent && resp.StatusCode != http.StatusOK {
		respBody, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("remove role failed with status %d: %s", resp.StatusCode, string(respBody))
	}

	return nil
}