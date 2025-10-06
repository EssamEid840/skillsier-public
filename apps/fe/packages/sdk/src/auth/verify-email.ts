import { authFetch } from "./clients/rest";

export async function sendVerifyEmail() {
  return authFetch<{ ok: true }>("send-verify-email", { method: "POST" });
}
