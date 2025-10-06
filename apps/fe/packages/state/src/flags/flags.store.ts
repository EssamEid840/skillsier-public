import { createStore } from "../core/createStore";
export type FlagValue = boolean | number | string;

export const createFlagsStore = (defaults?: Record<string, FlagValue>) =>
  createStore<{ flags: Record<string, FlagValue>; set: (k: string, v: FlagValue) => void }>((set, get) => ({
    flags: { ...(defaults ?? {}) },
    set: (k, v) => set({ flags: { ...get().flags, [k]: v } })
  }), { name: "flags", devtools: true });
