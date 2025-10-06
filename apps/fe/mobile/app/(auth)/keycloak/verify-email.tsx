import React from "react";
import Screen from "../../../src/navigation/Screen";

export default function VerifyEmail() {
  return (
    <Screen title="Verify your email" subtitle="Check your inbox for a verification link.">
      {/* You can poll your backend for verification status if needed. */}
    </Screen>
  );
}
