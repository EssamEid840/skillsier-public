"use client";

import React from "react";
import Link from "next/link";
import { useAnalytics } from "@/src/providers/AnalyticsProvider";
import { t, rtl } from "@skillsier/i18n";
import OAuthGoogleButton from "../components/OAuthGoogleButton";
import LoginForm from "../components/LoginForm";

// Small helper to read/write the preferred language in localStorage
function useLang() {
  const [lang, setLang] = React.useState<"en" | "ar">(() => {
    if (typeof window === "undefined") return "en";
    const saved = window.localStorage.getItem("lang");
    return (saved === "ar" ? "ar" : "en") as "en" | "ar";
  });

  React.useEffect(() => {
    if (typeof window !== "undefined") window.localStorage.setItem("lang", lang);
  }, [lang]);

  return { lang, setLang };
}

export default function LoginPage() {
  const { capture } = useAnalytics();
  const { lang, setLang } = useLang();
  const isRTL = rtl(lang);

  React.useEffect(() => {
    // simple view event; change to your preferred naming
    capture({ type: "auth_login", method: "google" });
  }, [capture]);

  return (
    <main className="min-h-screen flex items-start justify-center p-6">
      <div className="w-full max-w-md space-y-4" dir={isRTL ? "rtl" : "ltr"}>
        {/* Title + Language toggle */}
        <div className="flex items-center justify-between">
          <h1 className="text-2xl font-semibold">
            {t("auth.keycloak.login.title", lang)}
          </h1>

          <div className="flex gap-2">
            <button
              type="button"
              onClick={() => setLang("en")}
              className={`px-2 py-1 rounded border ${lang === "en" ? "bg-black text-white" : "bg-white"}`}
              aria-pressed={lang === "en"}
            >
              EN
            </button>
            <button
              type="button"
              onClick={() => setLang("ar")}
              className={`px-2 py-1 rounded border ${lang === "ar" ? "bg-black text-white" : "bg-white"}`}
              aria-pressed={lang === "ar"}
            >
              AR
            </button>
          </div>
        </div>

        {/* Your existing password form */}
        <LoginForm />

        {/* Divider */}
        <div className="grid place-items-center text-sm text-muted-foreground">or</div>

        {/* Google button via SDK */}
        <OAuthGoogleButton />

        {/* Footer links */}
        <div className="flex justify-between text-sm mt-2">
          <Link href="/keycloak/forgot-password" className="underline">
            {t("auth.keycloak.forgot", lang)}
          </Link>
          <Link href="/keycloak/register" className="underline">
            {t("auth.keycloak.register.title", lang)}
          </Link>
        </div>
      </div>
    </main>
  );
}
