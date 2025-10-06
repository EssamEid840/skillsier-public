import React from "react";
import { View, Text, Switch, StyleSheet } from "react-native";
import { theme } from "../../theme";

export type VisibilityToggleProps = {
  visible: boolean;
  onChange: (v: boolean) => void;
};

export default function VisibilityToggle({ visible, onChange }: VisibilityToggleProps) {
  return (
    <View style={styles.row}>
      <Text style={styles.label}>Public profile</Text>
      <Switch value={visible} onValueChange={onChange} />
    </View>
  );
}

const styles = StyleSheet.create({
  row: { flexDirection: "row", alignItems: "center", justifyContent: "space-between" },
  label: { color: theme.colors.fg }
});
