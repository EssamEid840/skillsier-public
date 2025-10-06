import { usersFetch } from "../clients/rest";

export async function getReferralCode() {
  return usersFetch<{ code: string }>("rest/referrals/code", { method: "GET" });
}
