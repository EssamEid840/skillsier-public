// apps/fe/mobile/app/(auth)/keycloak/social/google.tsx
import React, { useEffect } from "react";
import { Text, View, ActivityIndicator } from "react-native";
import { router } from "expo-router";
import { useAuth } from "src/lib/auth/providers/AuthProvider";
import { useAnalytics } from "src/providers/AnalyticsProvider";

export default function GoogleSocial() {
  const { signInWithGoogle } = useAuth();
  const { capture } = useAnalytics();

  useEffect(() => {
    (async () => {
      try {
        await signInWithGoogle();
        capture({ type: "auth_login", method: "google" });
        router.replace("/features/users");
      } catch (e) {
        console.warn(e);
      }
    })();
  }, [capture, signInWithGoogle]);

  return (
    <View style={{ flex: 1, justifyContent: "center", alignItems: "center", gap: 12 }}>
      <ActivityIndicator />
      <Text>Signing in with Google…</Text>
    </View>
  );
}
