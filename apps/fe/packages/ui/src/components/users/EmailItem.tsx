import React from "react";
import { View, Text, StyleSheet } from "react-native";
import { theme } from "../../theme";

export type EmailItemProps = {
  email: string;
  primary?: boolean;
  verified?: boolean;
};

export default function EmailItem({ email, primary, verified }: EmailItemProps) {
  return (
    <View style={styles.row}>
      <Text style={styles.email}>{email}</Text>
      <View style={styles.badges}>
        {primary && <Text style={styles.badge}>Primary</Text>}
        {verified ? <Text style={[styles.badge, styles.ok]}>Verified</Text> : <Text style={[styles.badge, styles.warn]}>Unverified</Text>}
      </View>
    </View>
  );
}

const styles = StyleSheet.create({
  row: { flexDirection: "row", alignItems: "center", justifyContent: "space-between" },
  email: { color: theme.colors.fg },
  badges: { flexDirection: "row", gap: 8 },
  badge: { fontSize: 12, color: theme.colors.mutedFg },
  ok: { color: theme.colors.success },
  warn: { color: theme.colors.warning }
});
