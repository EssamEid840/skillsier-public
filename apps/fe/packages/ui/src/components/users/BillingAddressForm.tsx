import React from "react";
import { View, Text, TextInput, StyleSheet } from "react-native";
import { theme } from "../../theme";

export type BillingAddress = {
  line1: string;
  line2?: string;
  city: string;
  state?: string;
  postalCode?: string;
  country: string;
};

export type BillingAddressFormProps = {
  value: BillingAddress;
  onChange: (next: BillingAddress) => void;
};

export default function BillingAddressForm({ value, onChange }: BillingAddressFormProps) {
  const set = (k: keyof BillingAddress, v: string) => onChange({ ...value, [k]: v });
  return (
    <View style={styles.wrap}>
      <TextInput placeholder="Address line 1" style={styles.in} value={value.line1} onChangeText={(t) => set("line1", t)} />
      <TextInput placeholder="Address line 2" style={styles.in} value={value.line2} onChangeText={(t) => set("line2", t)} />
      <TextInput placeholder="City" style={styles.in} value={value.city} onChangeText={(t) => set("city", t)} />
      <TextInput placeholder="State" style={styles.in} value={value.state} onChangeText={(t) => set("state", t)} />
      <TextInput placeholder="Postal code" style={styles.in} value={value.postalCode} onChangeText={(t) => set("postalCode", t)} />
      <TextInput placeholder="Country" style={styles.in} value={value.country} onChangeText={(t) => set("country", t)} />
    </View>
  );
}

const styles = StyleSheet.create({
  wrap: { gap: 10 },
  in: { borderWidth: 1, borderColor: theme.colors.border, borderRadius: 8, padding: 10, color: theme.colors.fg }
});
