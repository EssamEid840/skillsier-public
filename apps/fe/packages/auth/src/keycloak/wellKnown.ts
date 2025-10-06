import type { WellKnown } from "./types";

/** Fetch OpenID well-known from Keycloak realm */
export async function fetchWellKnown(url: string): Promise<WellKnown> {
  const res = await fetch(url);
  if (!res.ok) throw new Error(`WellKnown fetch failed ${res.status}`);
  return res.json() as Promise<WellKnown>;
}
