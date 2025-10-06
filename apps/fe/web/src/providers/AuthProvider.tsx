"use client";

import React, { createContext, useContext, useEffect, useState } from "react";

type User = { sub: string; email?: string; name?: string } | null;
type AuthCtx = { user: User; loading: boolean; refresh: () => Promise<void> };

const Ctx = createContext<AuthCtx | null>(null);

export function AuthProvider({ children }: { children: React.ReactNode }) {
  const [user, setUser] = useState<User>(null);
  const [loading, setLoading] = useState(true);

  const refresh = async () => {
    try {
      const r = await fetch("/api/auth/keycloak/me", { cache: "no-store" });
      const j = await r.json();
      setUser(j.authenticated ? (j.user || null) : null);
    } catch {
      setUser(null);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    refresh();
  }, []);

  return <Ctx.Provider value={{ user, loading, refresh }}>{children}</Ctx.Provider>;
}

export function useAuth() {
  const ctx = useContext(Ctx);
  if (!ctx) throw new Error("useAuth must be used within AuthProvider");
  return ctx;
}
