import React from "react";
import Screen from "../../../src/navigation/Screen";

export default function MfaChallenge() {
  return (
    <Screen title="Two-factor challenge" subtitle="Enter the code from your authenticator app or SMS.">
      {/* Implement TOTP/SMS challenge UI once your Keycloak policy demands it. */}
    </Screen>
  );
}
