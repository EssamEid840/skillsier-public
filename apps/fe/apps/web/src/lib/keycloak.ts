// apps/web/src/lib/keycloak.ts
// Production-ready Keycloak configuration with Google SSO support

import crypto from 'crypto';

// Keycloak Configuration
export const keycloakConfig = {
  url: process.env.NEXT_PUBLIC_KEYCLOAK_URL!,
  realm: process.env.NEXT_PUBLIC_KEYCLOAK_REALM!,
  clientId: process.env.NEXT_PUBLIC_KEYCLOAK_CLIENT_ID!,
  clientSecret: process.env.KEYCLOAK_CLIENT_SECRET!,
  issuerUrl: process.env.KEYCLOAK_ISSUER_URL!,
  mgmtClientId: process.env.KEYCLOAK_MGMT_CLIENT_ID!,
  mgmtClientSecret: process.env.KEYCLOAK_MGMT_CLIENT_SECRET!,
};

// Validate configuration
if (typeof window === 'undefined') {
  const required = [
    'NEXT_PUBLIC_KEYCLOAK_URL',
    'NEXT_PUBLIC_KEYCLOAK_REALM',
    'NEXT_PUBLIC_KEYCLOAK_CLIENT_ID',
    'KEYCLOAK_CLIENT_SECRET',
    'KEYCLOAK_ISSUER_URL',
    'KEYCLOAK_MGMT_CLIENT_ID',
    'KEYCLOAK_MGMT_CLIENT_SECRET',
  ];

  const missing = required.filter((key) => !process.env[key]);
  if (missing.length > 0) {
    throw new Error(
      `Missing required Keycloak environment variables: ${missing.join(', ')}`
    );
  }
}

// PKCE Helper Functions for Enhanced Security
export function generateCodeVerifier(): string {
  return crypto.randomBytes(32).toString('base64url');
}

export function generateCodeChallenge(verifier: string): string {
  return crypto.createHash('sha256').update(verifier).digest('base64url');
}

// Build Authorization URL with PKCE
export function buildAuthorizationUrl(params: {
  redirectUri: string;
  state: string;
  codeVerifier: string;
  idpHint?: string; // For direct Google SSO
}): string {
  const { redirectUri, state, codeVerifier, idpHint } = params;
  const codeChallenge = generateCodeChallenge(codeVerifier);

  const authUrl = new URL(
    `/realms/${keycloakConfig.realm}/protocol/openid-connect/auth`,
    keycloakConfig.url
  );

  authUrl.searchParams.set('client_id', keycloakConfig.clientId);
  authUrl.searchParams.set('redirect_uri', redirectUri);
  authUrl.searchParams.set('response_type', 'code');
  authUrl.searchParams.set('scope', 'openid email profile');
  authUrl.searchParams.set('state', state);
  authUrl.searchParams.set('code_challenge', codeChallenge);
  authUrl.searchParams.set('code_challenge_method', 'S256');

  // Direct Google SSO if specified
  if (idpHint) {
    authUrl.searchParams.set('kc_idp_hint', idpHint);
  }

  return authUrl.toString();
}

// Exchange Authorization Code for Tokens
export async function exchangeCodeForTokens(params: {
  code: string;
  redirectUri: string;
  codeVerifier: string;
}): Promise<TokenResponse> {
  const { code, redirectUri, codeVerifier } = params;

  const tokenUrl = `${keycloakConfig.url}/realms/${keycloakConfig.realm}/protocol/openid-connect/token`;

  const body = new URLSearchParams({
    grant_type: 'authorization_code',
    code,
    redirect_uri: redirectUri,
    client_id: keycloakConfig.clientId,
    client_secret: keycloakConfig.clientSecret,
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

// Get Management Access Token (for user registration)
export async function getManagementToken(): Promise<string> {
  const tokenUrl = `${keycloakConfig.url}/realms/${keycloakConfig.realm}/protocol/openid-connect/token`;

  const body = new URLSearchParams({
    grant_type: 'client_credentials',
    client_id: keycloakConfig.mgmtClientId,
    client_secret: keycloakConfig.mgmtClientSecret,
  });

  const response = await fetch(tokenUrl, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/x-www-form-urlencoded',
    },
    body: body.toString(),
  });

  if (!response.ok) {
    throw new Error('Failed to get management token');
  }

  const data = await response.json();
  return data.access_token;
}

// Create User in Keycloak
export async function createKeycloakUser(params: {
  email: string;
  firstName: string;
  lastName: string;
  password: string;
}): Promise<string> {
  const { email, firstName, lastName, password } = params;
  const token = await getManagementToken();

  const userUrl = `${keycloakConfig.url}/admin/realms/${keycloakConfig.realm}/users`;

  const userData = {
    email,
    emailVerified: true,
    enabled: true,
    firstName,
    lastName,
    username: email,
    credentials: [
      {
        type: 'password',
        value: password,
        temporary: false,
      },
    ],
  };

  const response = await fetch(userUrl, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
      Authorization: `Bearer ${token}`,
    },
    body: JSON.stringify(userData),
  });

  if (!response.ok) {
    const error = await response.text();
    throw new Error(`Failed to create user: ${error}`);
  }

  // Extract user ID from Location header
  const location = response.headers.get('location');
  if (!location) {
    throw new Error('User created but ID not returned');
  }

  const userId = location.split('/').pop()!;
  return userId;
}

// Refresh Access Token
export async function refreshAccessToken(refreshToken: string): Promise<TokenResponse> {
  const tokenUrl = `${keycloakConfig.url}/realms/${keycloakConfig.realm}/protocol/openid-connect/token`;

  const body = new URLSearchParams({
    grant_type: 'refresh_token',
    refresh_token: refreshToken,
    client_id: keycloakConfig.clientId,
    client_secret: keycloakConfig.clientSecret,
  });

  const response = await fetch(tokenUrl, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/x-www-form-urlencoded',
    },
    body: body.toString(),
  });

  if (!response.ok) {
    throw new Error('Token refresh failed');
  }

  return response.json();
}

// Logout from Keycloak
export async function logoutFromKeycloak(refreshToken: string): Promise<void> {
  const logoutUrl = `${keycloakConfig.url}/realms/${keycloakConfig.realm}/protocol/openid-connect/logout`;

  const body = new URLSearchParams({
    client_id: keycloakConfig.clientId,
    client_secret: keycloakConfig.clientSecret,
    refresh_token: refreshToken,
  });

  await fetch(logoutUrl, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/x-www-form-urlencoded',
    },
    body: body.toString(),
  });
}

// Verify JWT Token
export async function verifyToken(token: string): Promise<UserInfo> {
  const userInfoUrl = `${keycloakConfig.url}/realms/${keycloakConfig.realm}/protocol/openid-connect/userinfo`;

  const response = await fetch(userInfoUrl, {
    headers: {
      Authorization: `Bearer ${token}`,
    },
  });

  if (!response.ok) {
    throw new Error('Token verification failed');
  }

  return response.json();
}

// Types
export interface TokenResponse {
  access_token: string;
  refresh_token: string;
  expires_in: number;
  refresh_expires_in: number;
  token_type: string;
  id_token?: string;
  'not-before-policy'?: number;
  session_state?: string;
  scope: string;
}

export interface UserInfo {
  sub: string;
  email_verified: boolean;
  name: string;
  preferred_username: string;
  given_name: string;
  family_name: string;
  email: string;
}