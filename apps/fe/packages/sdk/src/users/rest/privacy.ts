import { usersFetch } from "../clients/rest";

export type Privacy = { showProfileInSearch: boolean };

export async function getPrivacy() {
  return usersFetch<Privacy>("rest/privacy", { method: "GET" });
}

export async function updatePrivacy(patch: Partial<Privacy>) {
  return usersFetch<Privacy>("rest/privacy", {
    method: "PATCH",
    body: JSON.stringify(patch)
  });
}
