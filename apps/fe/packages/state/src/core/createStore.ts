import { create } from "zustand";
import { devtools as _devtools } from "./middleware/devtools";
import { persist as _persist } from "./middleware/persist";
import { immer as _immer } from "./middleware/immer";
import { logger as _logger } from "./middleware/logger";

type MW<S> = (config: any, name?: string) => any;

export function createStore<S>(
  init: (set: any, get: any) => S,
  opts?: { name?: string; persist?: { key: string; storage: StorageAdapter }; devtools?: boolean; logger?: boolean; immer?: boolean }
) {
  let config: any = (set: any, get: any) => init(set, get);
  if (opts?.immer) config = _immer(config);
  if (opts?.logger) config = _logger(config);
  if (opts?.persist) config = _persist(config, opts.persist);
  const withDevtools = opts?.devtools ? _devtools(config, opts.name) : config;
  return create<S>()(withDevtools as any);
}

export type StorageAdapter = {
  getItem(key: string): Promise<string | null>;
  setItem(key: string, value: string): Promise<void>;
  removeItem(key: string): Promise<void>;
};
