/** Exchange auth code for tokens (BFF or native flow). */
export async function exchangeCodeForTokens(tokenUrl: string, params: {
  client_id: string;
  code: string;
  redirect_uri: string;
  code_verifier: string;
  client_secret?: string;
}) {
  const body = new URLSearchParams({
    grant_type: "authorization_code",
    client_id: params.client_id,
    code: params.code,
    redirect_uri: params.redirect_uri,
    code_verifier: params.code_verifier
  });
  if (params.client_secret) body.set("client_secret", params.client_secret);

  const res = await fetch(tokenUrl, { method: "POST", body });
  if (!res.ok) throw new Error(`exchangeCodeForTokens failed ${res.status}`);
  return res.json() as Promise<{ access_token: string; refresh_token?: string; id_token?: string; expires_in?: number }>;
}

/** Refresh tokens */
export async function refreshTokens(tokenUrl: string, params: {
  client_id: string;
  refresh_token: string;
  client_secret?: string;
}) {
  const body = new URLSearchParams({
    grant_type: "refresh_token",
    client_id: params.client_id,
    refresh_token: params.refresh_token
  });
  if (params.client_secret) body.set("client_secret", params.client_secret);

  const res = await fetch(tokenUrl, { method: "POST", body });
  if (!res.ok) throw new Error(`refreshTokens failed ${res.status}`);
  return res.json() as Promise<{ access_token: string; refresh_token?: string; id_token?: string; expires_in?: number }>;
}
