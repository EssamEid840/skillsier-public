// packages/shared/src/lib/keycloak/config.ts
// Shared Keycloak configuration

export interface KeycloakConfig {
  url: string;
  realm: string;
  clientId: string;
}

export function getKeycloakConfig(platform: 'web' | 'mobile'): KeycloakConfig {
  if (platform === 'web') {
    return {
      url: process.env.NEXT_PUBLIC_KEYCLOAK_URL || '',
      realm: process.env.NEXT_PUBLIC_KEYCLOAK_REALM || '',
      clientId: process.env.NEXT_PUBLIC_KEYCLOAK_CLIENT_ID || '',
    };
  } else {
    return {
      url: process.env.EXPO_PUBLIC_KEYCLOAK_URL || '',
      realm: process.env.EXPO_PUBLIC_KEYCLOAK_REALM || '',
      clientId: process.env.EXPO_PUBLIC_KEYCLOAK_CLIENT_ID || '',
    };
  }
}

export const KEYCLOAK_ENDPOINTS = {
  auth: (url: string, realm: string) =>
    `${url}/realms/${realm}/protocol/openid-connect/auth`,
  token: (url: string, realm: string) =>
    `${url}/realms/${realm}/protocol/openid-connect/token`,
  logout: (url: string, realm: string) =>
    `${url}/realms/${realm}/protocol/openid-connect/logout`,
  userInfo: (url: string, realm: string) =>
    `${url}/realms/${realm}/protocol/openid-connect/userinfo`,
  jwks: (url: string, realm: string) =>
    `${url}/realms/${realm}/protocol/openid-connect/certs`,
};