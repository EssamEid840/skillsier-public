import * as AuthSession from "expo-auth-session";
import Constants from "expo-constants";
import { Platform } from "react-native";
import { redirectUri } from "./deep-links";

type TokenResponse = {
  access_token: string;
  refresh_token?: string;
  id_token?: string;
  expires_in: number;
  token_type: string;
};

const KC = Constants.expoConfig?.extra?.keycloak as {
  issuer: string;          // https://<host>/realms/<realm>
  clientId: string;        // public client id for mobile
  scopes: string;          // "openid profile email"
  kcIdpHintGoogle?: string;// "google"
  directAccess?: boolean;  // enable username/password (Direct Access Grants)
};

const discovery = {
  authorizationEndpoint: `${KC.issuer}/protocol/openid-connect/auth`,
  tokenEndpoint:         `${KC.issuer}/protocol/openid-connect/token`,
  endSessionEndpoint:    `${KC.issuer}/protocol/openid-connect/logout`,
};

export async function startAuth({ useGoogle = false }: { useGoogle?: boolean }) {
  const request = new AuthSession.AuthRequest({
    clientId: KC.clientId,
    redirectUri,
    responseType: AuthSession.ResponseType.Code,
    usePKCE: true,
    scopes: KC.scopes.split(/\s+/),
    extraParams: useGoogle && KC.kcIdpHintGoogle ? { kc_idp_hint: KC.kcIdpHintGoogle } : undefined,
  });

  await request.makeAuthUrlAsync(discovery);

  // On web, a popup may be blocked by the browser. Users must allow popups.
  const result = await request.promptAsync(discovery);
  if (result.type !== "success" || !result.params.code) throw new Error("Login cancelled/failed.");

  const exchanged = await AuthSession.exchangeCodeAsync(
    {
      clientId: KC.clientId,
      code: result.params.code,
      redirectUri, // required here
      extraParams: { code_verifier: request.codeVerifier! },
    },
    discovery
  );

  return {
    access_token: exchanged.accessToken!,
    refresh_token: exchanged.refreshToken,
    id_token: exchanged.idToken,
    expires_in: exchanged.expiresIn ?? 3600,
    token_type: exchanged.tokenType ?? "Bearer",
  } satisfies TokenResponse;
}

// Username/password (Direct Access Grants) — requires a PUBLIC client with “Direct Access Grants” enabled in Keycloak
export async function passwordLogin(username: string, password: string) {
  if (!KC.directAccess) throw new Error("Direct Access Grants not enabled for this client.");
  const body = new URLSearchParams({
    grant_type: "password",
    client_id: KC.clientId,
    username,
    password,
    scope: KC.scopes,
  });
  const res = await fetch(discovery.tokenEndpoint, {
    method: "POST",
    headers: { "Content-Type": "application/x-www-form-urlencoded" },
    body: body.toString(), // RN fetch needs a string
  });
  if (!res.ok) throw new Error(`Password login failed: ${res.status}`);
  return (await res.json()) as TokenResponse;
}

export async function refreshToken(refresh_token: string) {
  const body = new URLSearchParams({
    grant_type: "refresh_token",
    client_id: KC.clientId,
    refresh_token,
  });
  const res = await fetch(discovery.tokenEndpoint, {
    method: "POST",
    headers: { "Content-Type": "application/x-www-form-urlencoded" },
    body: body.toString(),
  });
  if (!res.ok) throw new Error(`Refresh failed: ${res.status}`);
  return (await res.json()) as TokenResponse;
}

export async function logoutRequest(refresh_token?: string) {
  const body = new URLSearchParams({
    client_id: KC.clientId,
    ...(refresh_token ? { refresh_token } : {}),
  });
  await fetch(discovery.endSessionEndpoint, {
    method: "POST",
    headers: { "Content-Type": "application/x-www-form-urlencoded" },
    body: body.toString(),
  });
}
