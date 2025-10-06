import React from "react";
import { View, Text, StyleSheet } from "react-native";
import { theme } from "../../theme";

export type DeviceItemProps = {
  name: string;
  lastSeen?: string;
  location?: string;
};

export default function DeviceItem({ name, lastSeen, location }: DeviceItemProps) {
  return (
    <View style={styles.wrap}>
      <Text style={styles.name}>{name}</Text>
      <Text style={styles.meta}>{lastSeen ?? "—"} · {location ?? "—"}</Text>
    </View>
  );
}

const styles = StyleSheet.create({
  wrap: { paddingVertical: 8, borderBottomWidth: 1, borderBottomColor: theme.colors.border },
  name: { color: theme.colors.fg, fontWeight: "600" },
  meta: { color: theme.colors.mutedFg, fontSize: 12 }
});
