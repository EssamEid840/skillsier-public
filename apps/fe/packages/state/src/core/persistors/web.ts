import type { StorageAdapter } from "../createStore";

export const webStorage: StorageAdapter = {
  getItem: async (k) => {
    try { return typeof window !== "undefined" ? window.localStorage.getItem(k) : null; } catch { return null; }
  },
  setItem: async (k, v) => { try { if (typeof window !== "undefined") window.localStorage.setItem(k, v); } catch {} },
  removeItem: async (k) => { try { if (typeof window !== "undefined") window.localStorage.removeItem(k); } catch {} }
};
