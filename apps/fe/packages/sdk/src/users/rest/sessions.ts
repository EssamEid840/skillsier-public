import { usersFetch } from "../clients/rest";

export type Session = { id: string; ip: string; userAgent: string; time: string; current?: boolean };

export async function listSessions() {
  return usersFetch<Session[]>("rest/sessions", { method: "GET" });
}
