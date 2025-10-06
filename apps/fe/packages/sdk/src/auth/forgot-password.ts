import { authFetch } from "./clients/rest";

export async function forgotPassword(email: string) {
  return authFetch<{ ok: true }>("forgot-password", {
    method: "POST",
    body: JSON.stringify({ email })
  });
}
