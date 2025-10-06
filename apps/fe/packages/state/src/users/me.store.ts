import { createStore } from "../core/createStore";
import { webStorage } from "../core/persistors/web";
import type { SessionUser } from "../auth/session.store";

type State = {
  me: SessionUser | null;
  setMe: (v: SessionUser | null) => void;
};

export const createMeStore = (persistKey = "users.me", storage = webStorage) =>
  createStore<State>((set) => ({
    me: null,
    setMe: (v) => set({ me: v })
  }), { name: "users.me", persist: { key: persistKey, storage }, devtools: true });

export type MeStore = ReturnType<typeof createMeStore>;
