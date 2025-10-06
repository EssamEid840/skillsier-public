import { publicApiBase } from "../../_env";

type FetchInit = RequestInit & { headers?: Record<string, string> };

function envBase(): string | undefined {
  return publicApiBase();
}

export function authBaseUrl() {
  const base = envBase();
  return base ? `${base}/api/auth/keycloak` : `/api/auth/keycloak`;
}

export async function authFetch<T>(path: string, init?: FetchInit): Promise<T> {
  const url = `${authBaseUrl()}${path.startsWith("/") ? path : `/${path}`}`;
  const res = await fetch(url, {
    ...init,
    headers: {
      "Content-Type": "application/json",
      ...(init?.headers || {})
    },
    credentials: "include"
  });
  if (!res.ok) {
    const text = await res.text().catch(() => "");
    throw new Error(`Auth request failed ${res.status}: ${text}`);
  }
  if (res.status === 204) return undefined as T;
  return res.json() as Promise<T>;
}
