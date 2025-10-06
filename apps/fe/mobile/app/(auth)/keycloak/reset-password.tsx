import React from "react";
import Screen from "../../../src/navigation/Screen";

export default function ResetPassword() {
  return (
    <Screen title="Reset password" subtitle="If you arrived via an email link, handle the token server-side.">
      {/* For mobile apps, prefer linking to a secure web flow that consumes the token. */}
    </Screen>
  );
}
