"use client";

import React, { createContext, useContext, useMemo } from "react";
import { usePathname } from "next/navigation";
import { getMessages } from "../i18n";
type Dict = Record<string, string | Record<string, any>>;
type Ctx = { locale: "en" | "ar"; t: (key: string, vars?: Record<string, string | number>) => string };

const I18nCtx = createContext<Ctx | null>(null);

export function I18nProvider({ children }: { children: React.ReactNode }) {
  // naive locale inference from path: /ar/...
  const pathname = usePathname();
  const locale: "en" | "ar" = pathname?.startsWith("/ar") ? "ar" : "en";
  const dict: Dict = useMemo(() => getMessages(locale), [locale]);

  const t = (key: string, vars: Record<string, string | number> = {}) => {
    const val = key.split(".").reduce<any>((acc, k) => (acc ? acc[k] : undefined), dict);
    if (typeof val !== "string") return key;
    return val.replace(/\{(\w+)\}/g, (_, k) => String(vars[k] ?? `{${k}}`));
  };

  const value = useMemo(() => ({ locale, t }), [locale, dict]);

  return <I18nCtx.Provider value={value}>{children}</I18nCtx.Provider>;
}

export const useI18n = () => {
  const ctx = useContext(I18nCtx);
  if (!ctx) throw new Error("useI18n must be used within I18nProvider");
  return ctx;
};
