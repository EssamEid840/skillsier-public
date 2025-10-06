import React, { createContext, useMemo, useContext } from "react";
// ✅ import from the package root
import { combine, PosthogClient, AmplitudeClient } from "@skillsier/analytics";

const Ctx = createContext<{ capture: (e: any) => void }>({ capture: () => {} });

export function AnalyticsProvider({ children }: { children: React.ReactNode }) {
  const client = useMemo(() => combine(PosthogClient(), AmplitudeClient()), []);
  return <Ctx.Provider value={{ capture: client.capture }}>{children}</Ctx.Provider>;
}

export const useAnalytics = () => useContext(Ctx);
