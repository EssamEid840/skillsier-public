import { authFetch } from "./clients/rest";

export type LoginBody = {
  username: string;
  password: string;
};

export type LoginResponse = {
  ok: true;
};

export async function login(body: LoginBody) {
  return authFetch<LoginResponse>("login", {
    method: "POST",
    body: JSON.stringify(body)
  });
}
