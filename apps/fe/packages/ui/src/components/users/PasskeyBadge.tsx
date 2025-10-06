import React from "react";
import { Text, StyleSheet } from "react-native";
import { theme } from "../../theme";

export default function PasskeyBadge() {
  return <Text style={styles.badge}>Passkey Ready</Text>;
}

const styles = StyleSheet.create({
  badge: { color: theme.colors.success, fontWeight: "600" }
});
