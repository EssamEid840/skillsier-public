import { createStore } from "../core/createStore";
export type Session = { id: string; ip: string; ua: string; time: string; current?: boolean };
export const createSessionsStore = () =>
  createStore<{ list: Session[]; set: (s: Session[]) => void }>((set) => ({
    list: [],
    set: (s) => set({ list: s })
  }), { name: "users.sessions", devtools: true });
