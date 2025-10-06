import React from "react";
import { View, Text, StyleSheet } from "react-native";
import { theme } from "../../theme";

export type TimezonePickerProps = {
  value: string; // e.g. "Africa/Cairo"
};

export default function TimezonePicker({ value }: TimezonePickerProps) {
  return (
    <View style={styles.wrap}>
      <Text style={styles.label}>Timezone</Text>
      <Text style={styles.value}>{value}</Text>
      {/* Integrate full picker later; keeping this component cross-platform */}
    </View>
  );
}

const styles = StyleSheet.create({
  wrap: { gap: 4 },
  label: { color: theme.colors.mutedFg },
  value: { color: theme.colors.fg, fontWeight: "600" }
});
