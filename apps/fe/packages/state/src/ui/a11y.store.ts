import { createStore } from "../core/createStore";
export const createA11yStore = () =>
  createStore<{ reduceMotion: boolean; setReduceMotion: (v: boolean) => void }>((set) => ({
    reduceMotion: false,
    setReduceMotion: (v) => set({ reduceMotion: v })
  }), { name: "ui.a11y", devtools: true });
