export async function usersRest(path: string, init?: RequestInit) {
  const url = `/api/users/rest/${path.replace(/^\/+/, "")}`;
  const r = await fetch(url, init);
  if (!r.ok) throw new Error(`usersRest ${r.status}`);
  const ct = r.headers.get("content-type") || "";
  return ct.includes("application/json") ? r.json() : r.text();
}
