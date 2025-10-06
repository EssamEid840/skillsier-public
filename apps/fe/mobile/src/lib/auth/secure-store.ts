// apps/fe/mobile/src/lib/auth/secure-store.ts
import { Platform } from "react-native";
import * as SecureStore from "expo-secure-store";

const KEY = "kc_tokens_v1";
const isWeb = Platform.OS === "web";

// Safe handle to localStorage without DOM typings
function getLocalStorage(): { getItem?: (k: string) => string | null; setItem?: (k: string, v: string) => void; removeItem?: (k: string) => void } | undefined {
  try {
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    const ls = (globalThis as any)?.localStorage;
    if (!ls) return undefined;
    // simple feature test
    if (typeof ls.getItem === "function" && typeof ls.setItem === "function") return ls;
  } catch {
    // ignore SSR or restricted environments
  }
  return undefined;
}

const ls = getLocalStorage();

export type Tokens = {
  accessToken: string;
  refreshToken?: string;
  idToken?: string;
  expiresAt?: number;
};

async function setItem(key: string, value: string) {
  if (isWeb && ls?.setItem) {
    try { ls.setItem(key, value); } catch { /* ignore */ }
  } else {
    await SecureStore.setItemAsync(key, value);
  }
}

async function getItem(key: string) {
  if (isWeb && ls?.getItem) {
    try { return ls.getItem(key); } catch { return null; }
  }
  return await SecureStore.getItemAsync(key);
}

async function deleteItem(key: string) {
  if (isWeb && ls?.removeItem) {
    try { ls.removeItem(key); } catch { /* ignore */ }
  } else {
    await SecureStore.deleteItemAsync(key);
  }
}

export async function saveTokens(t: Tokens) {
  await setItem(KEY, JSON.stringify(t));
}

export async function getTokens(): Promise<Tokens | null> {
  const v = await getItem(KEY);
  return v ? (JSON.parse(v) as Tokens) : null;
}

export async function clearTokens() {
  await deleteItem(KEY);
}
