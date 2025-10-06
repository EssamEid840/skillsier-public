import { authFetch } from "./clients/rest";

export type RegisterBody = {
  email: string;
  password: string;
  displayName?: string;
};

export async function register(body: RegisterBody) {
  return authFetch<{ ok: true }>("register", {
    method: "POST",
    body: JSON.stringify(body)
  });
}
