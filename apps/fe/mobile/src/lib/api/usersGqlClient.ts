type GraphQLResponse<T> = { data?: T; errors?: { message: string }[] };

const BASE = process.env.EXPO_PUBLIC_USERS_BE_URL ?? "http://localhost:4000";

export async function gql<T>(query: string, variables?: Record<string, any>, token?: string) {
  const r = await fetch(`${BASE}/graphql`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
      ...(token ? { Authorization: `Bearer ${token}` } : {}),
    },
    body: JSON.stringify({ query, variables }),
  });
  if (!r.ok) throw new Error(`GraphQL HTTP ${r.status}`);
  const json = (await r.json()) as GraphQLResponse<T>;
  if (json.errors?.length) throw new Error(json.errors[0].message);
  return json.data as T;
}
