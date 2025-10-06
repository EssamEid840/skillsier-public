// Centralized public API base resolver that works in Next + Expo without @types/node
export function publicApiBase(): string | undefined {
  const g = globalThis as any;

  // Prefer env injected by bundlers (Next/Expo) if present
  const fromProcess =
    g?.process?.env?.NEXT_PUBLIC_API_BASE ??
    g?.process?.env?.EXPO_PUBLIC_API_BASE;

  // Optional: allow setting a global at runtime for tests/dev
  const fromGlobal =
    g?.__NEXT_PUBLIC_API_BASE ??
    g?.__EXPO_PUBLIC_API_BASE;

  return fromProcess ?? fromGlobal ?? undefined;
}
