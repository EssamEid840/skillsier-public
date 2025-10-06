"use client";
import React, { createContext, useContext, useMemo } from "react";
// ✅ import from the package root, not '/src/...'
import { combine, PosthogClient, AmplitudeClient } from "@skillsier/analytics";

const AnalyticsCtx = createContext<{ capture: (e: any) => void }>({ capture: () => {} });

export function AnalyticsProvider({ children }: { children: React.ReactNode }) {
  const client = useMemo(() => combine(PosthogClient(), AmplitudeClient()), []);
  return <AnalyticsCtx.Provider value={{ capture: client.capture }}>{children}</AnalyticsCtx.Provider>;
}

export function useAnalytics() {
  return useContext(AnalyticsCtx);
}
