import React from "react";
import { useLocalSearchParams } from "expo-router";
import Screen from "../../../../src/navigation/Screen";
export default function PortfolioItem(){
  const { itemId } = useLocalSearchParams<{ itemId: string }>();
  return <Screen title={`Portfolio Item: ${itemId ?? ""}`} />;
}
