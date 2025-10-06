import { createStore } from "../core/createStore";

export type Notice = { id: string; title: string; body?: string; read?: boolean };

type State = {
  items: Notice[];
  push: (n: Notice) => void;
  markRead: (id: string) => void;
};

export const createNotificationsStore = () =>
  createStore<State>((set, get) => ({
    items: [],
    push: (n) => set({ items: [n, ...get().items] }),
    markRead: (id) =>
      set({
        items: get().items.map((i: Notice) =>
          i.id === id ? { ...i, read: true } : i
        ),
      }),
  }), { name: "notifications", devtools: true });
