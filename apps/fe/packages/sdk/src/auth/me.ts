import { authFetch } from "./clients/rest";

export type Me = {
  id: string;
  email: string;
  displayName?: string;
  picture?: string;
};

export async function me() {
  return authFetch<Me>("me", { method: "GET" });
}
