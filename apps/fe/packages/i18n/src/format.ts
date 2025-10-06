type Dict = Record<string, any>;
export type Lang = "en" | "ar";

import enUsers from "./locales/en/users.json";
import enAuth from "./locales/en/auth.keycloak.json";
import arUsers from "./locales/ar/users.json";
import arAuth from "./locales/ar/auth.keycloak.json";

const dictionaries: Record<Lang, Dict> = {
  en: { users: enUsers, auth: { keycloak: enAuth } },
  ar: { users: arUsers, auth: { keycloak: arAuth } }
};

export function t(path: string, lang: Lang = "en", fallback?: string): string {
  const parts = path.split(".");
  let cur: any = dictionaries[lang];
  for (const p of parts) {
    if (cur && typeof cur === "object" && p in cur) cur = cur[p];
    else return fallback ?? path;
  }
  return typeof cur === "string" ? cur : fallback ?? path;
}

export function rtl(lang: Lang) { return lang === "ar"; }
