import { createStore } from "../core/createStore";
import { webStorage } from "../core/persistors/web";

export type Preferences = { locale: "en" | "ar"; theme: "light" | "dark" };
type State = Preferences & {
  setLocale: (l: "en" | "ar") => void;
  setTheme: (t: "light" | "dark") => void;
};

export const createPreferencesStore = (persistKey = "users.preferences", storage = webStorage) =>
  createStore<State>((set) => ({
    locale: "en",
    theme: "light",
    setLocale: (l) => set({ locale: l }),
    setTheme: (t) => set({ theme: t })
  }), { name: "users.preferences", persist: { key: persistKey, storage }, devtools: true });

export type PreferencesStore = ReturnType<typeof createPreferencesStore>;
