import React from "react";
import { View, Text, StyleSheet, Pressable } from "react-native";
import { theme } from "../../theme";

export type ReferralCardProps = {
  code: string;
  onShare?: () => void;
};

export default function ReferralCard({ code, onShare }: ReferralCardProps) {
  return (
    <View style={styles.card}>
      <Text style={styles.label}>Your referral code</Text>
      <Text style={styles.code}>{code}</Text>
      {onShare && (
        <Pressable style={styles.btn} onPress={onShare}>
          <Text style={styles.btnText}>Share</Text>
        </Pressable>
      )}
    </View>
  );
}

const styles = StyleSheet.create({
  card: { borderWidth: 1, borderColor: theme.colors.border, borderRadius: 12, padding: 12, backgroundColor: theme.colors.bg, alignItems: "flex-start", gap: 8 },
  label: { color: theme.colors.mutedFg },
  code: { color: theme.colors.fg, fontWeight: "700", fontSize: 18 },
  btn: { backgroundColor: theme.colors.primary, borderRadius: 8, paddingVertical: 8, paddingHorizontal: 12 },
  btnText: { color: theme.colors.primaryText, fontWeight: "600" }
});
