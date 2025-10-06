import { Redirect } from "expo-router";

export default function Index() {
  // No imperative router calls before the root layout mounts
  return <Redirect href="/(auth)/keycloak/login" />;
}
