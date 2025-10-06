import { publicApiBase } from "../../_env";

type FetchInit = RequestInit & { headers?: Record<string, string> };

function envBase(): string | undefined {
  return publicApiBase();
}

export function usersBaseUrl() {
  const base = envBase();
  return base ? `${base}/api/users` : `/api/users`;
}

export async function usersFetch<T>(path: string, init?: FetchInit): Promise<T> {
  const url = `${usersBaseUrl()}${path.startsWith("/") ? path : `/${path}`}`;
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
    throw new Error(`Users request failed ${res.status}: ${text}`);
  }
  if (res.status === 204) return undefined as T;
  return res.json() as Promise<T>;
}
