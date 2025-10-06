import { createStore } from "../core/createStore";
export type Email = { email: string; primary?: boolean; verified?: boolean };
export type Phone = { phone: string; verified?: boolean };

export const createContactStore = () =>
  createStore<{ emails: Email[]; phones: Phone[]; setEmails: (e: Email[]) => void; setPhones: (p: Phone[]) => void }>((set) => ({
    emails: [],
    phones: [],
    setEmails: (e) => set({ emails: e }),
    setPhones: (p) => set({ phones: p })
  }), { name: "users.contact", devtools: true });
