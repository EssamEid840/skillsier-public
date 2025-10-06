import React from "react";
import { View, Text } from "react-native";
import { Flags, enableDevDefaults } from "@skillsier/flags";

export default function PasskeysScreen() {
  enableDevDefaults();
  const show = Flags.get<boolean>("users.profile.passkeys", false);

  return (
    <View style={{ padding: 16 }}>
      <Text style={{ fontSize: 18, fontWeight: "600" }}>Passkeys</Text>
      {show ? <Text>Passkey Ready ✅</Text> : <Text>Disabled</Text>}
    </View>
  );
}
