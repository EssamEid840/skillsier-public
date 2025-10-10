// packages/shared/src/lib/keycloak/types.ts
// Keycloak type definitions

export interface KeycloakTokens {
  access_token: string;
  refresh_token: string;
  expires_in: number;
  refresh_expires_in: number;
  token_type: string;
  id_token?: string;
  scope: string;
}

export interface KeycloakUserInfo {
  sub: string;
  email_verified: boolean;
  name: string;
  preferred_username: string;
  given_name: string;
  family_name: string;
  email: string;
}

export interface KeycloakError {
  error: string;
  error_description?: string;
}

export interface AuthState {
  isAuthenticated: boolean;
  user: KeycloakUserInfo | null;
  tokens: KeycloakTokens | null;
  isLoading: boolean;
  error: string | null;
}