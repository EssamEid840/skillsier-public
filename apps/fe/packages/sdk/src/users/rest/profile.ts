import { usersFetch } from "../clients/rest";

export type Profile = {
  id: string;
  handle: string;
  displayName?: string;
  summary?: string;
  avatarUrl?: string;
};

export async function getMyProfile() {
  return usersFetch<Profile>("rest/profile/me", { method: "GET" });
}

export async function updateMyProfile(patch: Partial<Profile>) {
  return usersFetch<Profile>("rest/profile/me", {
    method: "PATCH",
    body: JSON.stringify(patch)
  });
}
