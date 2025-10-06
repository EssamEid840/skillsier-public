// Web "secure" storage just uses localStorage; for secrets, keep them in httpOnly cookies via BFF
export const Secure = {
  getItem: async (k: string) => {
    try { return typeof window !== "undefined" ? window.localStorage.getItem(k) : null; } catch { return null; }
  },
  setItem: async (k: string, v: string) => {
    try { if (typeof window !== "undefined") window.localStorage.setItem(k, v); } catch {}
  },
  removeItem: async (k: string) => {
    try { if (typeof window !== "undefined") window.localStorage.removeItem(k); } catch {}
  }
};
export default Secure;
