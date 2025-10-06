import React from "react";
import { View, Text, StyleSheet } from "react-native";
import { theme } from "../../theme";

export type SessionItemProps = {
  ip: string;
  userAgent: string;
  time: string;
  current?: boolean;
};

export default function SessionItem({ ip, userAgent, time, current }: SessionItemProps) {
  return (
    <View style={styles.wrap}>
      <Text style={styles.ua}>{userAgent}</Text>
      <Text style={styles.meta}>{ip} · {time} {current ? "· current" : ""}</Text>
    </View>
  );
}

const styles = StyleSheet.create({
  wrap: { paddingVertical: 8, borderBottomWidth: 1, borderBottomColor: theme.colors.border },
  ua: { color: theme.colors.fg },
  meta: { color: theme.colors.mutedFg, fontSize: 12 }
});
