import { usersFetch } from "../clients/rest";

export type Token = { id: string; label: string; createdAt: string };

export async function listTokens() {
  return usersFetch<Token[]>("rest/developer/tokens", { method: "GET" });
}
