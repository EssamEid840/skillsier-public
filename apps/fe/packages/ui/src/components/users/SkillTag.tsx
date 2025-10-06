import React from "react";
import { Text, View, StyleSheet, Pressable } from "react-native";
import { theme } from "../../theme";

export type SkillTagProps = {
  label: string;
  onPress?: () => void;
};

export default function SkillTag({ label, onPress }: SkillTagProps) {
  const Cmp = onPress ? Pressable : View;
  return (
    <Cmp style={styles.tag} onPress={onPress} accessibilityRole={onPress ? "button" : undefined}>
      <Text style={styles.text}>{label}</Text>
    </Cmp>
  );
}

const styles = StyleSheet.create({
  tag: {
    paddingVertical: 4,
    paddingHorizontal: 10,
    borderRadius: 999,
    backgroundColor: theme.colors.mutedBg,
    borderWidth: 1,
    borderColor: theme.colors.border
  },
  text: { color: theme.colors.fg, fontSize: theme.typography.size.sm }
});
