export const pick =
  <T, K extends keyof T>(keys: readonly K[]) =>
  (s: T) =>
    keys.reduce((acc, k) => ((acc[k] = s[k]), acc), {} as Pick<T, K>);
