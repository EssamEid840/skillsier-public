import { usersFetch } from "../clients/rest";

export async function startKyc() {
  return usersFetch<{ status: "started" }>("rest/verification/start", { method: "POST" });
}

export async function getKycStatus() {
  return usersFetch<{ status: "todo" | "in_progress" | "done" }>("rest/verification/status", { method: "GET" });
}
