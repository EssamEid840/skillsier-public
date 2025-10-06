import React from "react";
import { FlatList, ViewStyle } from "react-native";
import PortfolioCard, { PortfolioCardProps } from "./PortfolioCard";
import { theme } from "../../theme";

export type PortfolioGridProps = {
  items: (Omit<PortfolioCardProps, "onPress"> & { id: string })[];
  columns?: number;
  gap?: number;
  onItemPress?: (id: string) => void;
  style?: ViewStyle;
};

export default function PortfolioGrid({ items, columns = 2, gap = theme.spacing.md, onItemPress, style }: PortfolioGridProps) {
  return (
    <FlatList
      data={items}
      keyExtractor={(it) => it.id}
      numColumns={columns}
      columnWrapperStyle={{ gap }}
      contentContainerStyle={[{ gap, padding: gap }, style]}
      renderItem={({ item }) => (
        <PortfolioCard
          imageUri={item.imageUri}
          title={item.title}
          onPress={onItemPress ? () => onItemPress(item.id) : undefined}
        />
      )}
    />
  );
}
