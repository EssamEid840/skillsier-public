import { authFetch, authBaseUrl } from "../clients/rest";

/**
 * Web: returns redirect URL (server may redirect already).
 * Mobile: you likely open this URL in a browser (deep links handle callback).
 */
export async function startGoogleOAuth(): Promise<{ url: string }> {
  // If server redirects, you might not even reach this JSON.
  // Keep BFF returning JSON with a 'url' for mobile predictability.
  return authFetch<{ url: string }>("oauth/start/google", { method: "GET" });
}

/** Utility to build the start URL directly (handy for <a href> on web). */
export function startGoogleUrl() {
  return `${authBaseUrl()}/oauth/start/google`;
}
