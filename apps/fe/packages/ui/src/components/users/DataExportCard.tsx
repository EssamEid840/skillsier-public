import React from "react";
import { View, Text, StyleSheet, Pressable } from "react-native";
import { theme } from "../../theme";

export type DataExportCardProps = {
  onRequestExport?: () => void;
};

export default function DataExportCard({ onRequestExport }: DataExportCardProps) {
  return (
    <View style={styles.card}>
      <Text style={styles.title}>Download your data</Text>
      <Text style={styles.body}>Request an export of your account data as a zip file.</Text>
      {onRequestExport && (
        <Pressable style={styles.btn} onPress={onRequestExport}>
          <Text style={styles.btnText}>Request export</Text>
        </Pressable>
      )}
    </View>
  );
}

const styles = StyleSheet.create({
  card: { borderWidth: 1, borderColor: theme.colors.border, borderRadius: 12, padding: 12, backgroundColor: theme.colors.bg, gap: 8 },
  title: { color: theme.colors.fg, fontWeight: "700" },
  body: { color: theme.colors.mutedFg },
  btn: { backgroundColor: theme.colors.primary, borderRadius: 8, paddingVertical: 8, paddingHorizontal: 12, alignSelf: "flex-start" },
  btnText: { color: theme.colors.primaryText, fontWeight: "600" }
});
