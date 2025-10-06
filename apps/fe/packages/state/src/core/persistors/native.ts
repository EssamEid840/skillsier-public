import type { StorageAdapter } from "../createStore";
let mem: Record<string, string> = {};
export const nativeStorage: StorageAdapter = {
  getItem: async (k) => mem[k] ?? null,
  setItem: async (k, v) => { mem[k] = v; },
  removeItem: async (k) => { delete mem[k]; }
};
