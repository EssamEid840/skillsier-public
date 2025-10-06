import React from "react";
import { View, Image, Text, StyleSheet, Pressable } from "react-native";
import { theme } from "../../theme";

export type PortfolioCardProps = {
  imageUri?: string;
  title: string;
  onPress?: () => void;
};

export default function PortfolioCard({ imageUri, title, onPress }: PortfolioCardProps) {
  return (
    <Pressable onPress={onPress} style={styles.card} accessibilityRole={onPress ? "button" : undefined}>
      <View style={styles.media}>
        {imageUri ? (
          <Image source={{ uri: imageUri }} style={styles.img} resizeMode="cover" />
        ) : (
          <View style={styles.ph} />
        )}
      </View>
      <Text numberOfLines={1} style={styles.title}>{title}</Text>
    </Pressable>
  );
}

const styles = StyleSheet.create({
  card: { width: "100%", borderRadius: 12, overflow: "hidden", backgroundColor: theme.colors.bg, borderWidth: 1, borderColor: theme.colors.border },
  media: { width: "100%", aspectRatio: 16/9, backgroundColor: theme.colors.mutedBg },
  img: { width: "100%", height: "100%" },
  ph: { flex: 1, backgroundColor: "#e5e7eb" },
  title: { padding: 10, fontWeight: "600", color: theme.colors.fg }
});
