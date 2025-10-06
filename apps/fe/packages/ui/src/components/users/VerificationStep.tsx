import React from "react";
import { View, Text, StyleSheet } from "react-native";
import { theme } from "../../theme";

export type VerificationStepProps = {
  index: number;
  label: string;
  status: "todo" | "in_progress" | "done";
};

export default function VerificationStep({ index, label, status }: VerificationStepProps) {
  const color = status === "done" ? theme.colors.success : status === "in_progress" ? theme.colors.warning : theme.colors.mutedFg;
  return (
    <View style={styles.row}>
      <View style={[styles.dot, { backgroundColor: color }]} />
      <Text style={styles.text}>{index}. {label}</Text>
    </View>
  );
}

const styles = StyleSheet.create({
  row: { flexDirection: "row", alignItems: "center", gap: 10 },
  dot: { width: 10, height: 10, borderRadius: 5 },
  text: { color: theme.colors.fg }
});
