/**
 * Tiny native storage shim. In Expo you can replace this file by aliasing
 * to SecureStore, MMKV, or AsyncStorage if desired.
 */
let memory: Record<string, string> = {};

export async function getItem(key: string) { return memory[key] ?? null; }
export async function setItem(key: string, val: string) { memory[key] = val; }
export async function removeItem(key: string) { delete memory[key]; }
