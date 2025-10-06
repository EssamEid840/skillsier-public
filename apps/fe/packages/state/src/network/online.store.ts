import { createStore } from "../core/createStore";
export const createOnlineStore = () =>
  createStore<{ online: boolean; setOnline: (v: boolean) => void }>((set) => ({
    online: true,
    setOnline: (v) => set({ online: v })
  }), { name: "network.online", devtools: true });
