// apps/fe/mobile/src/lib/auth/providers/AuthProvider.tsx
import React, {
  createContext,
  useCallback,
  useContext,
  useEffect,
  useMemo,
  useRef,
  useState,
  PropsWithChildren,
} from "react";
import { AppState } from "react-native";
import { jwtDecode } from "jwt-decode";

import {
  saveTokens,
  getTokens as loadTokens,
  clearTokens,
  Tokens,
} from "../secure-store";

import {
  startAuth,
  passwordLogin,
  refreshToken as kcRefresh,
  logoutRequest,
} from "../keycloak.client";

type IdPayload = {
  sub: string;
  email?: string;
  email_verified?: boolean;
  name?: string;
  given_name?: string;
  family_name?: string;
  preferred_username?: string;
  // add more if you map extra claims
};

type User = {
  sub?: string;
  email?: string;
  emailVerified?: boolean;
  name?: string;
  givenName?: string;
  familyName?: string;
  username?: string;
} | null;

type AuthContextType = {
  user: User;
  tokens: Tokens | null;
  isAuthenticated: boolean;

  // Sign-in methods
  signInWithGoogle: () => Promise<void>;
  signInWithPassword: (email: string, password: string) => Promise<void>;

  // Session
  signOut: () => Promise<void>;
  getAccessToken: () => Promise<string | null>; // auto-refreshes if needed
};

const AuthContext = createContext<AuthContextType | undefined>(undefined);

function mapUserFromIdToken(idToken?: string): User {
  if (!idToken) return {};
  try {
    const p = jwtDecode<IdPayload>(idToken);
    return {
      sub: p.sub,
      email: p.email,
      emailVerified: p.email_verified,
      name: p.name,
      givenName: p.given_name,
      familyName: p.family_name,
      username: p.preferred_username,
    };
  } catch {
    return {};
  }
}

function now() {
  return Date.now();
}

function msUntilExpiry(expiresAt?: number | null) {
  if (!expiresAt) return 0;
  return Math.max(0, expiresAt - now());
}

export function AuthProvider({ children }: PropsWithChildren) {
  const [tokens, setTokens] = useState<Tokens | null>(null);
  const [user, setUser] = useState<User>(null);
  const timerRef = useRef<ReturnType<typeof setTimeout> | null>(null);

  const scheduleRefresh = useCallback(
    (t: Tokens | null) => {
      if (timerRef.current) {
        clearTimeout(timerRef.current);
        timerRef.current = null;
      }
      if (!t?.expiresAt) return;

      // refresh 45s before expiry, add small jitter (0–10s)
      const drift = Math.floor(Math.random() * 10000);
      const refreshIn = Math.max(0, t.expiresAt - now() - 45000 - drift);

      timerRef.current = setTimeout(async () => {
        try {
          if (t?.refreshToken) {
            const r = await kcRefresh(t.refreshToken);
            const expiresAt = now() + (r.expires_in - 30) * 1000;
            const updated: Tokens = {
              accessToken: r.access_token,
              refreshToken: r.refresh_token ?? t.refreshToken,
              idToken: r.id_token ?? t.idToken,
              expiresAt,
            };
            await saveTokens(updated);
            setTokens(updated);
            setUser(mapUserFromIdToken(updated.idToken));
            scheduleRefresh(updated);
          }
        } catch (e) {
          // refresh failed — clear session
          await clearTokens();
          setTokens(null);
          setUser(null);
        }
      }, refreshIn);
    },
    []
  );

  // Load session on mount
  useEffect(() => {
    let mounted = true;
    (async () => {
      const stored = await loadTokens();
      if (!mounted) return;

      if (stored) {
        // If nearly expired (<= 30s), try refresh
        const remaining = msUntilExpiry(stored.expiresAt);
        if (remaining <= 30000 && stored.refreshToken) {
          try {
            const r = await kcRefresh(stored.refreshToken);
            const expiresAt = now() + (r.expires_in - 30) * 1000;
            const updated: Tokens = {
              accessToken: r.access_token,
              refreshToken: r.refresh_token ?? stored.refreshToken,
              idToken: r.id_token ?? stored.idToken,
              expiresAt,
            };
            await saveTokens(updated);
            setTokens(updated);
            setUser(mapUserFromIdToken(updated.idToken));
            scheduleRefresh(updated);
            return;
          } catch {
            await clearTokens();
          }
        }
        setTokens(stored);
        setUser(mapUserFromIdToken(stored.idToken));
        scheduleRefresh(stored);
      }
    })();
    return () => {
      mounted = false;
      if (timerRef.current) clearTimeout(timerRef.current);
    };
  }, [scheduleRefresh]);

  // Refresh when app returns to foreground if close to expiry
  useEffect(() => {
    const sub = AppState.addEventListener("change", async (s) => {
      if (s !== "active") return;
      if (!tokens?.expiresAt || !tokens.refreshToken) return;
      if (msUntilExpiry(tokens.expiresAt) > 60000) return; // > 60s left, skip
      try {
        const r = await kcRefresh(tokens.refreshToken);
        const expiresAt = now() + (r.expires_in - 30) * 1000;
        const updated: Tokens = {
          accessToken: r.access_token,
          refreshToken: r.refresh_token ?? tokens.refreshToken,
          idToken: r.id_token ?? tokens.idToken,
          expiresAt,
        };
        await saveTokens(updated);
        setTokens(updated);
        setUser(mapUserFromIdToken(updated.idToken));
        scheduleRefresh(updated);
      } catch {
        await clearTokens();
        setTokens(null);
        setUser(null);
      }
    });
    return () => sub.remove();
  }, [tokens, scheduleRefresh]);

  const signInWithGoogle = useCallback(async () => {
    const res = await startAuth({ useGoogle: true });
    const expiresAt = now() + (res.expires_in - 30) * 1000;
    const saved: Tokens = {
      accessToken: res.access_token,
      refreshToken: res.refresh_token,
      idToken: res.id_token,
      expiresAt,
    };
    await saveTokens(saved);
    setTokens(saved);
    setUser(mapUserFromIdToken(saved.idToken));
    scheduleRefresh(saved);
  }, [scheduleRefresh]);

  const signInWithPassword = useCallback(
    async (email: string, password: string) => {
      const res = await passwordLogin(email, password);
      const expiresAt = now() + (res.expires_in - 30) * 1000;
      const saved: Tokens = {
        accessToken: res.access_token,
        refreshToken: res.refresh_token,
        idToken: res.id_token,
        expiresAt,
      };
      await saveTokens(saved);
      setTokens(saved);
      setUser(mapUserFromIdToken(saved.idToken));
      scheduleRefresh(saved);
    },
    [scheduleRefresh]
  );

  const getAccessToken = useCallback(async (): Promise<string | null> => {
    if (!tokens) return null;

    // Still fresh enough?
    if (msUntilExpiry(tokens.expiresAt) > 30000) {
      return tokens.accessToken;
    }

    // Try a refresh if we can
    if (tokens.refreshToken) {
      try {
        const r = await kcRefresh(tokens.refreshToken);
        const expiresAt = now() + (r.expires_in - 30) * 1000;
        const updated: Tokens = {
          accessToken: r.access_token,
          refreshToken: r.refresh_token ?? tokens.refreshToken,
          idToken: r.id_token ?? tokens.idToken,
          expiresAt,
        };
        await saveTokens(updated);
        setTokens(updated);
        setUser(mapUserFromIdToken(updated.idToken));
        scheduleRefresh(updated);
        return updated.accessToken;
      } catch {
        await clearTokens();
        setTokens(null);
        setUser(null);
        return null;
      }
    }

    // No refresh token available
    await clearTokens();
    setTokens(null);
    setUser(null);
    return null;
  }, [tokens, scheduleRefresh]);

  const signOut = useCallback(async () => {
    try {
      if (tokens?.refreshToken) {
        await logoutRequest(tokens.refreshToken);
      }
    } catch {
      // ignore network/logout failures
    } finally {
      await clearTokens();
      setTokens(null);
      setUser(null);
      if (timerRef.current) {
        clearTimeout(timerRef.current);
        timerRef.current = null;
      }
    }
  }, [tokens]);

  const value = useMemo<AuthContextType>(
    () => ({
      user,
      tokens,
      isAuthenticated: !!tokens?.accessToken,
      signInWithGoogle,
      signInWithPassword,
      signOut,
      getAccessToken,
    }),
    [user, tokens, signInWithGoogle, signInWithPassword, signOut, getAccessToken]
  );

  return <AuthContext.Provider value={value}>{children}</AuthContext.Provider>;
}

export function useAuth(): AuthContextType {
  const ctx = useContext(AuthContext);
  if (!ctx) throw new Error("useAuth must be used within <AuthProvider>");
  return ctx;
}
