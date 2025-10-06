import React from "react";
import { View, Text, Pressable, StyleSheet } from "react-native";
import { theme } from "../../theme";

export type CurrencyPickerProps = {
  value: string;
  onPick: (v: string) => void;
  options?: string[];
};

const DEFAULTS = ["USD", "EUR", "EGP"];

export default function CurrencyPicker({ value, onPick, options = DEFAULTS }: CurrencyPickerProps) {
  return (
    <View style={styles.wrap}>
      <Text style={styles.value}>Currency: {value}</Text>
      <View style={styles.row}>
        {options.map(opt => (
          <Pressable key={opt} style={[styles.btn, value === opt && styles.btnOn]} onPress={() => onPick(opt)}>
            <Text style={[styles.btnText, value === opt && styles.btnTextOn]}>{opt}</Text>
          </Pressable>
        ))}
      </View>
    </View>
  );
}

const styles = StyleSheet.create({
  wrap: { gap: 8 },
  value: { color: theme.colors.fg },
  row: { flexDirection: "row", flexWrap: "wrap", gap: 8 },
  btn: { paddingVertical: 6, paddingHorizontal: 10, borderRadius: 8, backgroundColor: theme.colors.mutedBg },
  btnOn: { backgroundColor: theme.colors.primary },
  btnText: { color: theme.colors.fg },
  btnTextOn: { color: theme.colors.primaryText, fontWeight: "700" }
});
