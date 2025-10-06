type Req = { query: string; variables?: Record<string, unknown>; operationName?: string };

export async function usersGql<T = any>(req: Req): Promise<T> {
  const r = await fetch("/api/users-graphql", {
    method: "POST",
    headers: { "content-type": "application/json" },
    body: JSON.stringify(req),
    cache: "no-store",
  });
  const j = await r.json();
  if (j.errors) throw j.errors;
  return j.data;
}
