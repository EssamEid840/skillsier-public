import React from "react";
import { View, Text, StyleSheet } from "react-native";
import { theme } from "../../theme";

export type PayoutMethod = {
  type: "bank" | "paypal";
  masked?: string; // e.g., ****1234
};

export type PayoutMethodCardProps = {
  method: PayoutMethod;
};

export default function PayoutMethodCard({ method }: PayoutMethodCardProps) {
  return (
    <View style={styles.card}>
      <Text style={styles.title}>{method.type.toUpperCase()}</Text>
      {!!method.masked && <Text style={styles.muted}>{method.masked}</Text>}
    </View>
  );
}

const styles = StyleSheet.create({
  card: { borderWidth: 1, borderColor: theme.colors.border, borderRadius: 12, padding: 12, backgroundColor: theme.colors.bg },
  title: { color: theme.colors.fg, fontWeight: "600" },
  muted: { color: theme.colors.mutedFg, marginTop: 4 }
});
