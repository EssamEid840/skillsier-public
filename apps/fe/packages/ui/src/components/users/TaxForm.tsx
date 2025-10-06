import React from "react";
import { View, TextInput, StyleSheet } from "react-native";
import { theme } from "../../theme";

export type TaxFormValue = {
  tin?: string; // tax identification number
  vat?: string;
};

export type TaxFormProps = {
  value: TaxFormValue;
  onChange: (next: TaxFormValue) => void;
};

export default function TaxForm({ value, onChange }: TaxFormProps) {
  const set = (k: keyof TaxFormValue, v: string) => onChange({ ...value, [k]: v });
  return (
    <View style={styles.wrap}>
      <TextInput placeholder="TIN" style={styles.in} value={value.tin} onChangeText={(t) => set("tin", t)} />
      <TextInput placeholder="VAT" style={styles.in} value={value.vat} onChangeText={(t) => set("vat", t)} />
    </View>
  );
}

const styles = StyleSheet.create({
  wrap: { gap: 10 },
  in: { borderWidth: 1, borderColor: theme.colors.border, borderRadius: 8, padding: 10, color: theme.colors.fg }
});
