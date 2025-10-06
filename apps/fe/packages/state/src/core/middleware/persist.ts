import type { StorageAdapter } from "../createStore";

export const persist = (config: any, options: { key: string; storage: StorageAdapter }) =>
  (set: any, get: any, api: any) => {
    let initialized = false;
    const load = async () => {
      try {
        const raw = await options.storage.getItem(options.key);
        if (raw) {
          const data = JSON.parse(raw);
          set({ ...get(), ...data }, true);
        }
      } catch {}
      initialized = true;
    };
    load();

    const save = async (state: any) => {
      try {
        await options.storage.setItem(options.key, JSON.stringify(state));
      } catch {}
    };

    return config((partial: any, replace?: boolean) => {
      set(partial, replace);
      if (initialized) save(get());
    }, get, api);
  };
