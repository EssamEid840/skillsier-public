/** Compose the Google IdP authorize URL via Keycloak */
export function composeGoogleStartUrl(authorizeUrl: string, opts: {
  client_id: string;
  redirect_uri: string;
  scope?: string;
  state?: string;
  code_challenge?: string;
  code_challenge_method?: "S256";
  kc_idp_hint?: "google";
  prompt?: string;
}) {
  const url = new URL(authorizeUrl);
  url.searchParams.set("response_type", "code");
  url.searchParams.set("client_id", opts.client_id);
  url.searchParams.set("redirect_uri", opts.redirect_uri);
  url.searchParams.set("scope", opts.scope ?? "openid profile email");
  if (opts.state) url.searchParams.set("state", opts.state);
  if (opts.code_challenge) url.searchParams.set("code_challenge", opts.code_challenge);
  if (opts.code_challenge_method) url.searchParams.set("code_challenge_method", opts.code_challenge_method);
  // Keycloak-specific: route to Google IdP
  url.searchParams.set("kc_idp_hint", opts.kc_idp_hint ?? "google");
  if (opts.prompt) url.searchParams.set("prompt", opts.prompt);
  return url.toString();
}
