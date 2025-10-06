import React, { useState } from "react";
import { View, Text, TextInput, Pressable, Alert } from "react-native";
import { backendForgotPassword } from "../../../src/lib/api/auth";

export default function ForgotPassword() {
  const [email, setEmail] = useState("");
  const [busy, setBusy] = useState(false);

  const onSubmit = async () => {
    setBusy(true);
    try {
      await backendForgotPassword(email.trim());
      Alert.alert("Email sent", "Check your inbox for the reset link.");
    } catch (e: any) {
      Alert.alert("Request failed", e?.message ?? String(e));
    } finally {
      setBusy(false);
    }
  };

  return (
    <View style={{ flex: 1, justifyContent: "center", padding: 24, gap: 12 }}>
      <Text style={{ fontWeight: "800", fontSize: 24, textAlign: "center" }}>Forgot password</Text>

      <Text>Email</Text>
      <TextInput autoCapitalize="none" keyboardType="email-address" value={email} onChangeText={setEmail}
        style={{ borderWidth: 1, borderRadius: 8, padding: 12 }} />

      <Pressable disabled={busy} onPress={onSubmit}
        style={{ backgroundColor: "#111827", padding: 12, borderRadius: 8, alignItems: "center" }}>
        <Text style={{ color: "white", fontWeight: "600" }}>{busy ? "Sending…" : "Send reset link"}</Text>
      </Pressable>
    </View>
  );
}
