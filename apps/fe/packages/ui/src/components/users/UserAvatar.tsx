// fe/packages/ui/src/components/users/UserAvatar.tsx
import React from "react";
import { Image, View, Text, StyleSheet } from "react-native";
import type { StyleProp, ImageStyle, ViewStyle } from "react-native";
import { theme } from "../../theme";

export type UserAvatarProps = {
  uri?: string;
  name?: string;
  size?: number;
  imageStyle?: StyleProp<ImageStyle>;
  containerStyle?: StyleProp<ViewStyle>;
};

function initials(name?: string) {
  if (!name) return "U";
  const parts = name.trim().split(/\s+/);
  const a = parts[0]?.[0] ?? "";
  const b = parts[1]?.[0] ?? "";
  return (a + b).toUpperCase() || a.toUpperCase() || "U";
}

export default function UserAvatar({
  uri,
  name,
  size = 64,
  imageStyle,
  containerStyle
}: UserAvatarProps) {
  const radius = size / 2;

  if (uri) {
    return (
      <Image
        source={{ uri }}
        style={[styles.avatar, { width: size, height: size, borderRadius: radius }, imageStyle]}
        resizeMode="cover"
        accessibilityLabel={`${name ?? "User"} avatar`}
      />
    );
  }

  return (
    <View
      style={[
        styles.fallback,
        { width: size, height: size, borderRadius: radius, backgroundColor: theme.colors.mutedBg },
        containerStyle
      ]}
      accessibilityRole="image"
      aria-label={`${name ?? "User"} avatar`}
    >
      <Text style={[styles.fallbackText, { fontSize: size * 0.38 }]}>{initials(name)}</Text>
    </View>
  );
}

const styles = StyleSheet.create({
  avatar: { backgroundColor: "#ddd" },
  fallback: { alignItems: "center", justifyContent: "center" },
  fallbackText: { color: "#374151", fontWeight: "600" }
});
