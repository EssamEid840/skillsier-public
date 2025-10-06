import React from "react";
import { View, Text, StyleSheet, Switch } from "react-native";
import { theme } from "../../theme";

export type AccessibilityControlsProps = {
  reduceMotion: boolean;
  onChangeReduceMotion: (v: boolean) => void;
};

export default function AccessibilityControls({ reduceMotion, onChangeReduceMotion }: AccessibilityControlsProps) {
  return (
    <View style={styles.row}>
      <Text style={styles.label}>Reduce motion</Text>
      <Switch value={reduceMotion} onValueChange={onChangeReduceMotion} />
    </View>
  );
}

const styles = StyleSheet.create({
  row: { flexDirection: "row", alignItems: "center", justifyContent: "space-between" },
  label: { color: theme.colors.fg }
});
