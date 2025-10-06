import type { Jwks } from "./types";

/** Fetch JWKS (public keys) for ID/Access token verification */
export async function fetchJwks(jwksUri: string): Promise<Jwks> {
  const res = await fetch(jwksUri);
  if (!res.ok) throw new Error(`JWKS fetch failed ${res.status}`);
  return res.json() as Promise<Jwks>;
}
