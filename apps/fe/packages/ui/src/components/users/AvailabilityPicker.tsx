import React from "react";
import { View, Text, Pressable, StyleSheet } from "react-native";
import { theme } from "../../theme";

export type AvailabilityPickerProps = {
  available: boolean;
  onChange: (next: boolean) => void;
};

export default function AvailabilityPicker({ available, onChange }: AvailabilityPickerProps) {
  return (
    <View style={styles.wrap}>
      <Text style={styles.label}>Availability</Text>
      <Pressable style={[styles.btn, available && styles.btnOn]} onPress={() => onChange(!available)} accessibilityRole="switch" aria-checked={available}>
        <Text style={[styles.btnText, available && styles.btnTextOn]}>{available ? "Available" : "Away"}</Text>
      </Pressable>
    </View>
  );
}

const styles = StyleSheet.create({
  wrap: { gap: 8 },
  label: { color: theme.colors.mutedFg },
  btn: { paddingVertical: 8, paddingHorizontal: 12, borderRadius: 8, backgroundColor: theme.colors.mutedBg },
  btnOn: { backgroundColor: theme.colors.success },
  btnText: { color: theme.colors.fg },
  btnTextOn: { color: theme.colors.primaryText, fontWeight: "700" }
});
