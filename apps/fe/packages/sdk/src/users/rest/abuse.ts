import { usersFetch } from "../clients/rest";

export async function reportAbuse(targetUserId: string, reason: string) {
  return usersFetch<{ ok: true }>("rest/abuse/report", {
    method: "POST",
    body: JSON.stringify({ targetUserId, reason })
  });
}
