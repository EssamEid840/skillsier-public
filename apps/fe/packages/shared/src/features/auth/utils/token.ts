// packages/shared/src/features/auth/utils/token.ts
// Token management utilities

export interface DecodedToken {
  exp: number;
  iat: number;
  sub: string;
  email?: string;
  name?: string;
  preferred_username?: string;
  [key: string]: any;
}

/**
 * Decode JWT token (without verification)
 */
export function decodeToken(token: string): DecodedToken | null {
  try {
    const parts = token.split('.');
    if (parts.length !== 3) {
      return null;
    }

    const payload = parts[1];
    const decoded = JSON.parse(atob(payload.replace(/-/g, '+').replace(/_/g, '/')));
    return decoded;
  } catch (error) {
    console.error('Failed to decode token:', error);
    return null;
  }
}

/**
 * Check if token is expired
 */
export function isTokenExpired(token: string): boolean {
  const decoded = decodeToken(token);
  if (!decoded || !decoded.exp) {
    return true;
  }

  // Check with 5 minute buffer
  const expirationTime = decoded.exp * 1000;
  const currentTime = Date.now();
  const bufferTime = 5 * 60 * 1000; // 5 minutes

  return currentTime >= expirationTime - bufferTime;
}

/**
 * Get token expiration time
 */
export function getTokenExpiration(token: string): Date | null {
  const decoded = decodeToken(token);
  if (!decoded || !decoded.exp) {
    return null;
  }

  return new Date(decoded.exp * 1000);
}

/**
 * Get time until token expires (in seconds)
 */
export function getTimeUntilExpiration(token: string): number {
  const decoded = decodeToken(token);
  if (!decoded || !decoded.exp) {
    return 0;
  }

  const expirationTime = decoded.exp * 1000;
  const currentTime = Date.now();
  const timeRemaining = Math.max(0, expirationTime - currentTime);

  return Math.floor(timeRemaining / 1000);
}

/**
 * Extract user info from token
 */
export function getUserInfoFromToken(token: string): {
  id: string;
  email?: string;
  name?: string;
  username?: string;
} | null {
  const decoded = decodeToken(token);
  if (!decoded) {
    return null;
  }

  return {
    id: decoded.sub,
    email: decoded.email,
    name: decoded.name,
    username: decoded.preferred_username,
  };
}

/**
 * Check if token has required scopes
 */
export function hasRequiredScopes(token: string, requiredScopes: string[]): boolean {
  const decoded = decodeToken(token);
  if (!decoded || !decoded.scope) {
    return false;
  }

  const tokenScopes = decoded.scope.split(' ');
  return requiredScopes.every(scope => tokenScopes.includes(scope));
}