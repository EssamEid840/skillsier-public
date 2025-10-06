import React from "react";
import { View, Text, Pressable, StyleSheet } from "react-native";
import { theme } from "../../theme";

export type ConnectedAccountItemProps = {
  provider: "google";
  linked: boolean;
  onLink?: () => void;
  onUnlink?: () => void;
};

export default function ConnectedAccountItem({ provider, linked, onLink, onUnlink }: ConnectedAccountItemProps) {
  return (
    <View style={styles.row}>
      <Text style={styles.label}>{provider === "google" ? "Google" : provider}</Text>
      {linked ? (
        <Pressable style={[styles.btn, styles.danger]} onPress={onUnlink}><Text style={styles.btnText}>Unlink</Text></Pressable>
      ) : (
        <Pressable style={styles.btn} onPress={onLink}><Text style={styles.btnText}>Link</Text></Pressable>
      )}
    </View>
  );
}

const styles = StyleSheet.create({
  row: { flexDirection: "row", alignItems: "center", justifyContent: "space-between" },
  label: { color: theme.colors.fg },
  btn: { backgroundColor: theme.colors.primary, paddingVertical: 8, paddingHorizontal: 12, borderRadius: 8 },
  danger: { backgroundColor: theme.colors.danger },
  btnText: { color: theme.colors.primaryText, fontWeight: "600" }
});
