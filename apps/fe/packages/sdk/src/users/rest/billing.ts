import { usersFetch } from "../clients/rest";

export type BillingProfile = { country: string; currency: string };

export async function getBillingProfile() {
  return usersFetch<BillingProfile>("rest/billing/profile", { method: "GET" });
}
