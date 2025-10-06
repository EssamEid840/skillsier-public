import React, { useMemo } from "react";
import { View, Text, Pressable, StyleSheet } from "react-native";
import { theme } from "../../theme";

export type ThemeToggleProps = {
  mode: "light" | "dark";
  onToggle: () => void;
};

export default function ThemeToggle({ mode, onToggle }: ThemeToggleProps) {
  const label = useMemo(() => (mode === "light" ? "Switch to dark" : "Switch to light"), [mode]);
  return (
    <View style={styles.row}>
      <Text style={styles.label}>Theme</Text>
      <Pressable style={styles.btn} onPress={onToggle}>
        <Text style={styles.btnText}>{label}</Text>
      </Pressable>
    </View>
  );
}

const styles = StyleSheet.create({
  row: { flexDirection: "row", alignItems: "center", justifyContent: "space-between" },
  label: { color: theme.colors.fg },
  btn: { backgroundColor: theme.colors.primary, paddingVertical: 8, paddingHorizontal: 12, borderRadius: 8 },
  btnText: { color: theme.colors.primaryText, fontWeight: "600" }
});
