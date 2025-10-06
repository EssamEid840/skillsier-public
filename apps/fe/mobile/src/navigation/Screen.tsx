import React from "react";
import { ScrollView, Text, View } from "react-native";

export default function Screen({
  title,
  children,
  subtitle,
}: {
  title: string;
  subtitle?: string;
  children?: React.ReactNode;
}) {
  return (
    <ScrollView contentContainerStyle={{ padding: 16 }}>
      <View style={{ gap: 8 }}>
        <Text style={{ fontSize: 22, fontWeight: "700" }}>{title}</Text>
        {subtitle ? (
          <Text style={{ color: "#6b7280" }}>{subtitle}</Text>
        ) : null}
        <View style={{ height: 1, backgroundColor: "#e5e7eb", marginVertical: 8 }} />
        <View style={{ gap: 12 }}>{children}</View>
      </View>
    </ScrollView>
  );
}
