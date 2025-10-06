import React from "react";
import Screen from "../../../src/navigation/Screen";

export default function Register() {
  return (
    <Screen title="Create your account" subtitle="Use Google sign-in from the login screen for now.">
      {/* TODO: If you want a custom form, proxy to your backend to call Keycloak Admin endpoints safely. */}
    </Screen>
  );
}
