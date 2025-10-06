import React from "react";
import { View, Text, StyleSheet } from "react-native";
import { theme } from "../../theme";

export type CountryPickerProps = {
  value: string; // ISO or display label
};

export default function CountryPicker({ value }: CountryPickerProps) {
  return (
    <View style={styles.wrap}>
      <Text style={styles.label}>Country</Text>
      <Text style={styles.value}>{value}</Text>
      {/* Plug a real picker/list later */}
    </View>
  );
}

const styles = StyleSheet.create({
  wrap: { gap: 4 },
  label: { color: theme.colors.mutedFg },
  value: { color: theme.colors.fg, fontWeight: "600" }
});
