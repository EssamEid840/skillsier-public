import React from "react";
import { View, Text, StyleSheet, Switch } from "react-native";
import { theme } from "../../theme";

export type PrivacyControlsProps = {
  showProfileInSearch: boolean;
  onChangeShowProfileInSearch: (v: boolean) => void;
};

export default function PrivacyControls({ showProfileInSearch, onChangeShowProfileInSearch }: PrivacyControlsProps) {
  return (
    <View style={styles.wrap}>
      <View style={styles.row}>
        <Text style={styles.label}>Appear in search</Text>
        <Switch value={showProfileInSearch} onValueChange={onChangeShowProfileInSearch} />
      </View>
    </View>
  );
}

const styles = StyleSheet.create({
  wrap: { gap: 12 },
  row: { flexDirection: "row", alignItems: "center", justifyContent: "space-between" },
  label: { color: theme.colors.fg }
});
