import React, { useState, useEffect } from "react";
import { View, Text, TextInput, Pressable, Alert } from "react-native";
import { router, Link } from "expo-router";
import { t, rtl } from "@skillsier/i18n";
import { useAnalytics } from "src/providers/AnalyticsProvider";
import { useAuth } from "src/lib/auth/providers/AuthProvider";

export default function Login() {
  // deps
  const { capture } = useAnalytics();
  const { signInWithGoogle, signInWithPassword } = useAuth();

  // local state
  const [email, setEmail] = useState("");
  const [pass, setPass] = useState("");
  const [busy, setBusy] = useState(false);

  // language (per-screen toggle; you can hoist this to a global app provider if you prefer)
  const [lang, setLang] = useState<"en" | "ar">("en");
  const isRTL = rtl(lang);

  // analytics view
  useEffect(() => {
    capture({ type: "auth_login", method: "google" });
  }, [capture]);

  // actions
  const onGoogle = async () => {
    if (busy) return;
    setBusy(true);
    try {
      await signInWithGoogle();
      router.replace("/features/users");
    } catch (e: any) {
      Alert.alert("Google sign-in failed", e?.message ?? String(e));
    } finally {
      setBusy(false);
    }
  };

  const onPassword = async () => {
    if (busy) return;
    setBusy(true);
    try {
      await signInWithPassword(email.trim(), pass);
      router.replace("/features/users");
    } catch (e: any) {
      Alert.alert("Sign-in failed", e?.message ?? String(e));
    } finally {
      setBusy(false);
    }
  };

  return (
    <View
      style={{
        gap: 12,
        padding: 16,
        // RN-safe layout flip for RTL
        direction: isRTL ? "rtl" : "ltr",
      }}
    >
      {/* Title + Language toggle */}
      <View style={{ flexDirection: "row", justifyContent: "space-between", alignItems: "center" }}>
        <Text style={{ fontSize: 24, fontWeight: "700" }}>
          {t("auth.keycloak.login.title", lang)}
        </Text>

        <View style={{ flexDirection: "row", gap: 8 }}>
          <Pressable
            onPress={() => setLang("en")}
            style={{
              paddingHorizontal: 8,
              paddingVertical: 6,
              borderRadius: 6,
              borderWidth: 1,
              backgroundColor: lang === "en" ? "#111827" : "white",
            }}
          >
            <Text style={{ color: lang === "en" ? "white" : "#111827", fontWeight: "600" }}>EN</Text>
          </Pressable>
          <Pressable
            onPress={() => setLang("ar")}
            style={{
              paddingHorizontal: 8,
              paddingVertical: 6,
              borderRadius: 6,
              borderWidth: 1,
              backgroundColor: lang === "ar" ? "#111827" : "white",
            }}
          >
            <Text style={{ color: lang === "ar" ? "white" : "#111827", fontWeight: "600" }}>AR</Text>
          </Pressable>
        </View>
      </View>

      {/* Email */}
      <TextInput
        value={email}
        onChangeText={setEmail}
        placeholder={t("auth.keycloak.login.email", lang) || "Email"}
        autoCapitalize="none"
        keyboardType="email-address"
        style={{
          borderWidth: 1,
          borderColor: "#e5e7eb",
          padding: 12,
          borderRadius: 8,
          textAlign: isRTL ? "right" : "left",
          backgroundColor: "#e8f0fe",
        }}
      />

      {/* Password */}
      <TextInput
        value={pass}
        onChangeText={setPass}
        placeholder={t("auth.keycloak.login.password", lang) || "Password"}
        secureTextEntry
        style={{
          borderWidth: 1,
          borderColor: "#e5e7eb",
          padding: 12,
          borderRadius: 8,
          textAlign: isRTL ? "right" : "left",
          backgroundColor: "#e8f0fe",
        }}
      />

      {/* Continue (password) */}
      <Pressable
        disabled={busy}
        onPress={onPassword}
        style={{
          backgroundColor: "#111827",
          padding: 12,
          borderRadius: 8,
          alignItems: "center",
          opacity: busy ? 0.7 : 1,
        }}
      >
        <Text style={{ color: "white", fontWeight: "700" }}>
          {t("auth.keycloak.login.submit", lang) || "Continue"}
        </Text>
      </Pressable>

      {/* Continue with Google */}
      <Pressable
        disabled={busy}
        onPress={onGoogle}
        style={{
          backgroundColor: "#2563eb",
          padding: 12,
          borderRadius: 8,
          alignItems: "center",
          opacity: busy ? 0.7 : 1,
        }}
      >
        <Text style={{ color: "white", fontWeight: "700" }}>
          {t("auth.keycloak.login.google", lang) || "Continue with Google"}
        </Text>
      </Pressable>

      {/* Footer */}
      <View style={{ flexDirection: "row", justifyContent: "space-between", marginTop: 8 }}>
        <Link href="/(auth)/keycloak/forgot-password">
          {t("auth.keycloak.forgot", lang)}
        </Link>
        <Link href="/(auth)/keycloak/register">
          {t("auth.keycloak.register.title", lang)}
        </Link>
      </View>
    </View>
  );
}
