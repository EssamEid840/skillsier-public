import { usersFetch } from "../clients/rest";

export async function changePassword(currentPassword: string, newPassword: string) {
  return usersFetch<{ ok: true }>("rest/security/password", {
    method: "POST",
    body: JSON.stringify({ currentPassword, newPassword })
  });
}
