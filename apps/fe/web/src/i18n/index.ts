import enUsers from "./messages/en/users.json";
import enAuth from "./messages/en/auth.keycloak.json";
import arUsers from "./messages/ar/users.json";
import arAuth from "./messages/ar/auth.keycloak.json";

export function getMessages(locale: "en" | "ar") {
  if (locale === "ar") {
    return { users: arUsers, auth: arAuth } as const;
  }
  return { users: enUsers, auth: enAuth } as const;
}
