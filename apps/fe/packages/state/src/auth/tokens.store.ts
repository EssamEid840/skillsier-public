import { createStore } from "../core/createStore";
import { webStorage } from "../core/persistors/web";

type Tokens = { accessToken?: string; refreshToken?: string };
type State = Tokens & { setTokens: (t: Tokens) => void; clear: () => void; };

export const createTokensStore = (persistKey = "tokens", storage = webStorage) =>
  createStore<State>((set) => ({
    accessToken: undefined,
    refreshToken: undefined,
    setTokens: (t) => set(t),
    clear: () => set({ accessToken: undefined, refreshToken: undefined })
  }), { name: "tokens", persist: { key: persistKey, storage }, devtools: true });

export type TokensStore = ReturnType<typeof createTokensStore>;
