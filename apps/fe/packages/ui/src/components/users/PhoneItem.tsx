import React from "react";
import { View, Text, StyleSheet } from "react-native";
import { theme } from "../../theme";

export type PhoneItemProps = {
  phone: string;
  verified?: boolean;
};

export default function PhoneItem({ phone, verified }: PhoneItemProps) {
  return (
    <View style={styles.row}>
      <Text style={styles.phone}>{phone}</Text>
      <Text style={[styles.badge, verified ? styles.ok : styles.warn]}>{verified ? "Verified" : "Unverified"}</Text>
    </View>
  );
}

const styles = StyleSheet.create({
  row: { flexDirection: "row", alignItems: "center", justifyContent: "space-between" },
  phone: { color: theme.colors.fg },
  badge: { fontSize: 12 },
  ok: { color: theme.colors.success },
  warn: { color: theme.colors.warning }
});
