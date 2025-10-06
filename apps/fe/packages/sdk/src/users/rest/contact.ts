import { usersFetch } from "../clients/rest";

export type Email = { email: string; primary?: boolean; verified?: boolean };
export type Phone = { phone: string; verified?: boolean };

export async function listEmails() {
  return usersFetch<Email[]>("rest/contact/emails", { method: "GET" });
}

export async function listPhones() {
  return usersFetch<Phone[]>("rest/contact/phones", { method: "GET" });
}
