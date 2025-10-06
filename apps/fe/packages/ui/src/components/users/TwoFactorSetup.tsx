import React from "react";
import { View, Text, Pressable, StyleSheet } from "react-native";
import { theme } from "../../theme";

export type TwoFactorSetupProps = {
  enabled: boolean;
  onSetup?: () => void;
  onDisable?: () => void;
};

export default function TwoFactorSetup({ enabled, onSetup, onDisable }: TwoFactorSetupProps) {
  return (
    <View style={styles.row}>
      <Text style={styles.label}>Two-factor authentication</Text>
      {enabled ? (
        <Pressable style={[styles.btn, styles.danger]} onPress={onDisable}><Text style={styles.btnText}>Disable</Text></Pressable>
      ) : (
        <Pressable style={styles.btn} onPress={onSetup}><Text style={styles.btnText}>Enable</Text></Pressable>
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
