import { createStore } from "../core/createStore";
export const createConnectedStore = () =>
  createStore<{ google: boolean; setGoogle: (v: boolean) => void }>((set) => ({
    google: false,
    setGoogle: (v) => set({ google: v })
  }), { name: "users.connected", devtools: true });
