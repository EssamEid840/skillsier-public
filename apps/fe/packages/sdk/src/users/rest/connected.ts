import { usersFetch } from "../clients/rest";

export async function getConnectedAccounts() {
  return usersFetch<{ google: boolean }>("rest/connected", { method: "GET" });
}
