export const pick = <S, K extends keyof S>(keys: K[]) => (s: S) => {
  const out = {} as Pick<S, K>;
  keys.forEach(k => (out[k] = s[k]));
  return out;
};
