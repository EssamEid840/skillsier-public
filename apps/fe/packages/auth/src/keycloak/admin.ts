/**
 * Minimal admin token exchange helpers (server-to-server or BFF usage).
 * Keep light; full admin client is handled server-side.
 */
export async function clientCredentialsToken(tokenUrl: string, clientId: string, clientSecret: string) {
  const body = new URLSearchParams({
    grant_type: "client_credentials",
    client_id: clientId,
    client_secret: clientSecret
  });
  const res = await fetch(tokenUrl, { method: "POST", body });
  if (!res.ok) throw new Error(`clientCredentialsToken failed ${res.status}`);
  return res.json() as Promise<{ access_token: string; expires_in: number; token_type: string }>;
}
