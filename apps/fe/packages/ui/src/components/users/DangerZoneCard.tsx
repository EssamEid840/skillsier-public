import React from "react";
import { View, Text, StyleSheet, Pressable } from "react-native";
import { theme } from "../../theme";

export type DangerZoneCardProps = {
  onDelete?: () => void;
};

export default function DangerZoneCard({ onDelete }: DangerZoneCardProps) {
  return (
    <View style={styles.card}>
      <Text style={styles.title}>Danger zone</Text>
      <Text style={styles.body}>Delete your account and all associated data. This action is irreversible.</Text>
      {onDelete && (
        <Pressable style={styles.btn} onPress={onDelete}>
          <Text style={styles.btnText}>Delete account</Text>
        </Pressable>
      )}
    </View>
  );
}

const styles = StyleSheet.create({
  card: { borderWidth: 1, borderColor: theme.colors.danger, borderRadius: 12, padding: 12, backgroundColor: "#fff5f5", gap: 8 },
  title: { color: theme.colors.danger, fontWeight: "700" },
  body: { color: "#7f1d1d" },
  btn: { backgroundColor: theme.colors.danger, borderRadius: 8, paddingVertical: 8, paddingHorizontal: 12, alignSelf: "flex-start" },
  btnText: { color: theme.colors.primaryText, fontWeight: "700" }
});
