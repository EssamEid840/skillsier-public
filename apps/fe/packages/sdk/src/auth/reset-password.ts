import { authFetch } from "./clients/rest";

export async function resetPassword(token: string, newPassword: string) {
  return authFetch<{ ok: true }>("reset-password", {
    method: "POST",
    body: JSON.stringify({ token, newPassword })
  });
}
