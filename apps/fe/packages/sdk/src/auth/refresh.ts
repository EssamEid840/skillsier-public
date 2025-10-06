import { authFetch } from "./clients/rest";

export async function refresh() {
  return authFetch<{ ok: true }>("refresh", { method: "POST" });
}
