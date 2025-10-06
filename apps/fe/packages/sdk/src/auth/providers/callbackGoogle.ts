import { authFetch } from "../clients/rest";

/**
 * Exchange the provider callback params with the BFF to finalize login/link.
 * Pass through the query you received on your callback/deep link.
 */
export async function completeGoogleCallback(query: Record<string, string>) {
  return authFetch<{ ok: true }>("oauth/callback/google", {
    method: "POST",
    body: JSON.stringify(query)
  });
}
