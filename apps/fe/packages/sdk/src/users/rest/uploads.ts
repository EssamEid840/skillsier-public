import { usersFetch } from "../clients/rest";

export async function getUploadUrl(kind: "avatar" | "portfolio") {
  return usersFetch<{ url: string }>("rest/uploads/url", {
    method: "POST",
    body: JSON.stringify({ kind })
  });
}
