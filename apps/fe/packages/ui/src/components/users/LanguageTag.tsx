import React from "react";
import { View, Text, StyleSheet } from "react-native";
import { theme } from "../../theme";

export type LanguageTagProps = {
  language: string;
  level?: "basic" | "conversational" | "fluent" | "native";
};

export default function LanguageTag({ language, level }: LanguageTagProps) {
  return (
    <View style={styles.wrap}>
      <Text style={styles.lang}>{language}</Text>
      {level && <Text style={styles.level}>{level}</Text>}
    </View>
  );
}

const styles = StyleSheet.create({
  wrap: {
    paddingVertical: 4,
    paddingHorizontal: 10,
    borderRadius: 8,
    backgroundColor: theme.colors.mutedBg,
    borderWidth: 1,
    borderColor: theme.colors.border,
    flexDirection: "row",
    gap: 6
  },
  lang: { color: theme.colors.fg },
  level: { color: theme.colors.mutedFg, fontSize: theme.typography.size.sm }
});
