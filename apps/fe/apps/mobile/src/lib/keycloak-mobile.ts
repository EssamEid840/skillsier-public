// apps/mobile/src/lib/keycloak-mobile.ts
// Mobile Keycloak integration with WebBrowser for OAuth

import * as WebBrowser from 'expo-web-browser';
import * as Linking from 'expo-linking';
import { Platform } from 'react-native';
import { MMKV } from 'react-native-mmkv';

// Initialize MMKV storage
const storage = new MMKV();

// Keycloak Configuration
const KEYCLOAK_URL = process.env.EXPO_PUBLIC_KEYCLOAK_URL!;
const KEYCLOAK_REALM = process.env.EXPO_PUBLIC_KEYCLOAK_REALM!;
const CLIENT_ID = process.env.EXPO_PUBLIC_KEYCLOAK_CLIENT_ID!;

// For OAuth redirects
WebBrowser.maybeCompleteAuthSession();

// Helper to generate random string
function generateRandomString(length: number): string {
  const charset = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-._~';
  let result = '';
  const randomValues = new Uint8Array(length);
  crypto.getRandomValues(randomValues);
  
  for (let i = 0; i < length; i++) {
    result += charset[randomValues[i] % charset.length];
  }
  return result;
}

// Generate code verifier for PKCE
function generateCodeVerifier(): string {
  return generateRandomString(128);
}

// Generate code challenge from verifier
async function generateCodeChallenge(verifier: string): Promise<string> {
  const encoder = new TextEncoder();
  const data = encoder.encode(verifier);
  const hash = await crypto.subtle.digest('SHA-256', data);
  
  // Convert to base64url
  const base64 = btoa(String.fromCharCode(...new Uint8Array(hash)));
  return base64
    .replace(/\+/g, '-')
    .replace(/\//g, '_')
    .replace(/=/g, '');
}

// Build authorization URL
async function buildAuthUrl(params: {
  redirectUri: string;
  state: string;
  codeVerifier: string;
  idpHint?: string;
}): Promise<string> {
  const { redirectUri, state, codeVerifier, idpHint } = params;
  const codeChallenge = await generateCodeChallenge(codeVerifier);

  const authUrl = new URL(
    `/realms/${KEYCLOAK_REALM}/protocol/openid-connect/auth`,
    KEYCLOAK_URL
  );

  authUrl.searchParams.set('client_id', CLIENT_ID);
  authUrl.searchParams.set('redirect_uri', redirectUri);
  authUrl.searchParams.set('response_type', 'code');
  authUrl.searchParams.set('scope', 'openid email profile');
  authUrl.searchParams.set('state', state);
  authUrl.searchParams.set('code_challenge', codeChallenge);
  authUrl.searchParams.set('code_challenge_method', 'S256');

  if (idpHint) {
    authUrl.searchParams.set('kc_idp_hint', idpHint);
  }

  return authUrl.toString();
}

// Exchange code for tokens
async function exchangeCodeForTokens(params: {
  code: string;
  redirectUri: string;
  codeVerifier: string;
}): Promise<TokenResponse> {
  const { code, redirectUri, codeVerifier } = params;

  const tokenUrl = `${KEYCLOAK_URL}/realms/${KEYCLOAK_REALM}/protocol/openid-connect/token`;

  const body = new URLSearchParams({
    grant_type: 'authorization_code',
    code,
    redirect_uri: redirectUri,
    client_id: CLIENT_ID,
    code_verifier: codeVerifier,
  });

  const response = await fetch(tokenUrl, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/x-www-form-urlencoded',
    },
    body: body.toString(),
  });

  if (!response.ok) {
    const error = await response.text();
    throw new Error(`Token exchange failed: ${error}`);
  }

  return response.json();
}

// Main login function
export async function loginWithProvider(provider: 'google' | 'local' = 'google') {
  try {
    // Generate PKCE parameters
    const state = generateRandomString(32);
    const codeVerifier = generateCodeVerifier();

    // Store code verifier for later use
    storage.set('code_verifier', codeVerifier);
    storage.set('oauth_state', state);

    // Get redirect URI
    const redirectUri = Linking.createURL('/auth/callback');

    // Build auth URL
    const authUrl = await buildAuthUrl({
      redirectUri,
      state,
      codeVerifier,
      idpHint: provider !== 'local' ? provider : undefined,
    });

    // Open browser for authentication
    const result = await WebBrowser.openAuthSessionAsync(
      authUrl,
      redirectUri
    );

    if (result.type === 'success') {
      const { url } = result;
      const params = new URL(url).searchParams;
      const code = params.get('code');
      const returnedState = params.get('state');

      // Verify state
      const savedState = storage.getString('oauth_state');
      if (returnedState !== savedState) {
        throw new Error('State mismatch - possible CSRF attack');
      }

      if (!code) {
        throw new Error('No authorization code received');
      }

      // Exchange code for tokens
      const savedCodeVerifier = storage.getString('code_verifier');
      if (!savedCodeVerifier) {
        throw new Error('Code verifier not found');
      }

      const tokens = await exchangeCodeForTokens({
        code,
        redirectUri,
        codeVerifier: savedCodeVerifier,
      });

      // Store tokens
      storage.set('access_token', tokens.access_token);
      storage.set('refresh_token', tokens.refresh_token);
      storage.set('token_expires_at', Date.now() + tokens.expires_in * 1000);

      // Clean up temporary storage
      storage.delete('code_verifier');
      storage.delete('oauth_state');

      return tokens;
    }

    throw new Error('Authentication cancelled or failed');
  } catch (error) {
    console.error('Login error:', error);
    throw error;
  }
}

// Refresh access token
export async function refreshAccessToken(): Promise<TokenResponse> {
  const refreshToken = storage.getString('refresh_token');
  if (!refreshToken) {
    throw new Error('No refresh token available');
  }

  const tokenUrl = `${KEYCLOAK_URL}/realms/${KEYCLOAK_REALM}/protocol/openid-connect/token`;

  const body = new URLSearchParams({
    grant_type: 'refresh_token',
    refresh_token: refreshToken,
    client_id: CLIENT_ID,
  });

  const response = await fetch(tokenUrl, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/x-www-form-urlencoded',
    },
    body: body.toString(),
  });

  if (!response.ok) {
    // Refresh token expired or invalid
    storage.delete('access_token');
    storage.delete('refresh_token');
    storage.delete('token_expires_at');
    throw new Error('Token refresh failed');
  }

  const tokens = await response.json();

  // Update stored tokens
  storage.set('access_token', tokens.access_token);
  storage.set('refresh_token', tokens.refresh_token);
  storage.set('token_expires_at', Date.now() + tokens.expires_in * 1000);

  return tokens;
}

// Get valid access token (refresh if needed)
export async function getAccessToken(): Promise<string | null> {
  const accessToken = storage.getString('access_token');
  const expiresAt = storage.getNumber('token_expires_at');

  if (!accessToken || !expiresAt) {
    return null;
  }

  // Check if token is expired (with 5 minute buffer)
  if (Date.now() >= expiresAt - 300000) {
    try {
      const tokens = await refreshAccessToken();
      return tokens.access_token;
    } catch {
      return null;
    }
  }

  return accessToken;
}

// Logout
export async function logout(): Promise<void> {
  const refreshToken = storage.getString('refresh_token');

  if (refreshToken) {
    const logoutUrl = `${KEYCLOAK_URL}/realms/${KEYCLOAK_REALM}/protocol/openid-connect/logout`;

    try {
      await fetch(logoutUrl, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/x-www-form-urlencoded',
        },
        body: new URLSearchParams({
          client_id: CLIENT_ID,
          refresh_token: refreshToken,
        }).toString(),
      });
    } catch (error) {
      console.error('Logout error:', error);
    }
  }

  // Clear all stored tokens
  storage.delete('access_token');
  storage.delete('refresh_token');
  storage.delete('token_expires_at');
  storage.delete('user_info');
}

// Check if user is authenticated
export function isAuthenticated(): boolean {
  const token = storage.getString('access_token');
  const expiresAt = storage.getNumber('token_expires_at');
  
  if (!token || !expiresAt) {
    return false;
  }

  return Date.now() < expiresAt;
}

// Types
export interface TokenResponse {
  access_token: string;
  refresh_token: string;
  expires_in: number;
  refresh_expires_in: number;
  token_type: string;
  id_token?: string;
  scope: string;
}