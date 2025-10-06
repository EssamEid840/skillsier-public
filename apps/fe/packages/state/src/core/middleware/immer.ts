import { produce } from "immer";
export const immer = (config: any) =>
  (set: any, get: any, api: any) =>
    config((fn: any, replace?: boolean) => {
      const nextState = typeof fn === "function" ? produce(get(), fn) : fn;
      set(nextState, replace);
    }, get, api);
