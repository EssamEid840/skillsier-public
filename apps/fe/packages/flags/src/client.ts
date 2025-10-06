import type { FlagClient, FlagMap, FlagValue } from "./types";

const store: FlagMap = Object.create(null);

export const Flags: FlagClient = {
  get: <T extends FlagValue = FlagValue>(key: string, fallback?: T) => (store[key] as T) ?? (fallback as T),
  set: (key: string, value: FlagValue) => { store[key] = value; },
  all: () => ({ ...store })
};

export function enableDevDefaults() {
  Flags.set("users.profile.passkeys", true);
  Flags.set("users.profile.google_link", true);
}
