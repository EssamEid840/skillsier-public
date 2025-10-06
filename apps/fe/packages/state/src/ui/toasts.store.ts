import { createStore } from "../core/createStore";

export type Toast = {
  id: string;
  title: string;
  message?: string;
  variant?: "default" | "success" | "error";
};

type State = {
  list: Toast[];
  push: (t: Toast) => void;
  remove: (id: string) => void;
};

export const createToastsStore = () =>
  createStore<State>((set, get) => ({
    list: [],
    push: (t) => set({ list: [t, ...get().list] }),
    remove: (id) =>
      set({
        list: get().list.filter((x: Toast) => x.id !== id),
      }),
  }), { name: "ui.toasts", devtools: true });
