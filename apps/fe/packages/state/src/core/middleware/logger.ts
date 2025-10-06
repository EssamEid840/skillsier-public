// No Node types needed — read env via globalThis
export const logger = (config: any) =>
  (set: any, get: any, api: any) =>
    config((partial: any, replace?: boolean) => {
      const g = globalThis as any;
      const nodeEnv: string | undefined = g?.process?.env?.NODE_ENV;
      const isProd = nodeEnv === "production";

      if (!isProd) {
        // eslint-disable-next-line no-console
        console.log("[store]", typeof partial === "function" ? "fn(update)" : partial);
      }
      set(partial, replace);
    }, get, api);
