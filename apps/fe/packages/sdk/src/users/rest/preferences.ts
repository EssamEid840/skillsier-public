import { usersFetch } from "../clients/rest";

export type Preferences = { locale: string; theme: "light" | "dark" };

export async function getPreferences() {
  return usersFetch<Preferences>("rest/preferences", { method: "GET" });
}

export async function updatePreferences(patch: Partial<Preferences>) {
  return usersFetch<Preferences>("rest/preferences", {
    method: "PATCH",
    body: JSON.stringify(patch)
  });
}
