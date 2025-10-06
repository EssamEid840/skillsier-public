import { usersFetch } from "../clients/rest";

export async function completeOnboarding() {
  return usersFetch<{ ok: true }>("rest/onboarding/complete", { method: "POST" });
}
