import { createStore } from "../core/createStore";
export const createSecurityStore = () =>
  createStore<{ passkeysEnabled: boolean; setEnabled: (v: boolean) => void }>((set) => ({
    passkeysEnabled: false,
    setEnabled: (v) => set({ passkeysEnabled: v })
  }), { name: "users.security", devtools: true });
