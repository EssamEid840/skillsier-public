import { authFetch } from "./clients/rest";

export async function logout() {
  return authFetch<{ ok: true }>("logout", { method: "POST" });
}
