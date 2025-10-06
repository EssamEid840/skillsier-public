// Tries Expo SecureStore if available; falls back to in-memory
let mem: Record<string, string> = {};
type MaybeSecureStore = { getItemAsync(k: string): Promise<string|null>; setItemAsync(k: string,v: string): Promise<void>; deleteItemAsync(k: string): Promise<void>; };
const g = globalThis as any;
const SS: MaybeSecureStore | undefined = g?.ExpoSecureStore ?? g?.SecureStore;

export const Secure = {
  getItem: async (k: string) => (SS ? SS.getItemAsync(k) : mem[k] ?? null),
  setItem: async (k: string, v: string) => (SS ? SS.setItemAsync(k, v) : (mem[k] = v, undefined)),
  removeItem: async (k: string) => (SS ? SS.deleteItemAsync(k) : delete mem[k])
};
export default Secure;
