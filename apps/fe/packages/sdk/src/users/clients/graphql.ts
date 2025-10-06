import { publicApiBase } from "../../_env";

function envBase(): string | undefined {
  return publicApiBase();
}

export function usersGraphqlUrl() {
  const base = envBase();
  return base ? `${base}/api/users-graphql` : `/api/users-graphql`;
}

export async function gqlRequest<TData = any, TVariables = Record<string, any>>(
  query: string,
  variables?: TVariables
): Promise<TData> {
  const res = await fetch(usersGraphqlUrl(), {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    credentials: "include",
    body: JSON.stringify({ query, variables })
  });
  if (!res.ok) {
    const text = await res.text().catch(() => "");
    throw new Error(`GraphQL request failed ${res.status}: ${text}`);
  }
  const payload = await res.json();
  if (payload.errors?.length) {
    throw new Error(`GraphQL errors: ${JSON.stringify(payload.errors)}`);
  }
  return payload.data as TData;
}
