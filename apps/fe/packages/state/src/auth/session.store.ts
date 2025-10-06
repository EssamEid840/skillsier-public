import { createStore } from "../core/createStore";
import { webStorage } from "../core/persistors/web"; // Next defaults; mobile can use nativeStorage when creating
export type SessionUser = { id: string; email?: string; displayName?: string; picture?: string };

type State = {
  user: SessionUser | null;
  setUser: (u: SessionUser | null) => void;
  clear: () => void;
};

export const createSessionStore = (persistKey = "session", storage = webStorage) =>
  createStore<State>((set) => ({
    user: null,
    setUser: (u) => set({ user: u }),
    clear: () => set({ user: null })
  }), { name: "session", persist: { key: persistKey, storage }, devtools: true });

export type SessionStore = ReturnType<typeof createSessionStore>;
