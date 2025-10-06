export type FlagValue = boolean | number | string;
export type FlagMap = Record<string, FlagValue>;

export type FlagClient = {
  get: <T extends FlagValue = FlagValue>(key: string, fallback?: T) => T;
  set: (key: string, value: FlagValue) => void;
  all: () => FlagMap;
};
