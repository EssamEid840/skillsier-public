import React from "react";
import { View, Text, StyleSheet } from "react-native";
import UserAvatar from "./UserAvatar";
import { theme } from "../../theme";

export type ProfileHeaderProps = {
  name: string;
  handle?: string;
  avatarUrl?: string;
};

export default function ProfileHeader({ name, handle, avatarUrl }: ProfileHeaderProps) {
  return (
    <View style={styles.wrap}>
      <UserAvatar uri={avatarUrl} name={name} size={72} />
      <View style={styles.texts}>
        <Text style={styles.name}>{name}</Text>
        {!!handle && <Text style={styles.handle}>@{handle}</Text>}
      </View>
    </View>
  );
}

const styles = StyleSheet.create({
  wrap: { flexDirection: "row", alignItems: "center", gap: theme.spacing.md },
  texts: { gap: 2 },
  name: { fontSize: theme.typography.size.xl, fontWeight: "700", color: theme.colors.fg },
  handle: { fontSize: theme.typography.size.sm, color: theme.colors.mutedFg }
});
