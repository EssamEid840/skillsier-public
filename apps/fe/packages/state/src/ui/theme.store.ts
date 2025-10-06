import { createStore } from "../core/createStore";
import { webStorage } from "../core/persistors/web";

export const createThemeStore = (persistKey = "ui.theme", storage = webStorage) =>
  createStore<{ theme: "light" | "dark"; setTheme: (t: "light" | "dark") => void }>((set) => ({
    theme: "light",
    setTheme: (t) => set({ theme: t })
  }), { name: "ui.theme", persist: { key: persistKey, storage }, devtools: true });
