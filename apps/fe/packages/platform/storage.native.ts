// Minimal async storage shim for native; swap to AsyncStorage/MMKV later if you like
let mem: Record<string, string> = {};
export async function getItem(key: string) { return mem[key] ?? null; }
export async function setItem(key: string, value: string) { mem[key] = value; }
export async function removeItem(key: string) { delete mem[key]; }
