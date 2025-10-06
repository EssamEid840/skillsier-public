export const devtools = (config: any, name?: string) =>
  (set: any, get: any, api: any) =>
    config((...args: any[]) => {
      set(...args);
      // Devtools bridge is handled by Zustand if installed in app
    }, get, api);
