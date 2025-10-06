import * as AuthSession from "expo-auth-session";
import * as WebBrowser from "expo-web-browser";
import { Platform } from "react-native";

WebBrowser.maybeCompleteAuthSession();

// Good defaults for web + native
export const redirectUri = AuthSession.makeRedirectUri({
  scheme: "skillsier",            // must be allowed in Keycloak client
  preferLocalhost: true,          // web dev
  isTripleSlashed: true,          // android scheme format
  // native: "skillsier://",      // optional explicit native scheme
});

// helpful in logs:
if (__DEV__) console.log("[Auth] redirectUri:", redirectUri, "platform:", Platform.OS);
