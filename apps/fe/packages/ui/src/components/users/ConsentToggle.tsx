import React from "react";
import { View, Text, Switch, StyleSheet } from "react-native";
import { theme } from "../../theme";

export type ConsentToggleProps = {
  label: string;
  value: boolean;
  onChange: (v: boolean) => void;
};

export default function ConsentToggle({ label, value, onChange }: ConsentToggleProps) {
  return (
    <View style={styles.row}>
      <Text style={styles.label}>{label}</Text>
      <Switch value={value} onValueChange={onChange} />
    </View>
  );
}

const styles = StyleSheet.create({
  row: { flexDirection: "row", alignItems: "center", justifyContent: "space-between" },
  label: { color: theme.colors.fg }
});
