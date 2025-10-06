import type { WellKnown } from "./types";

export function getAuthEndpoints(wk: WellKnown) {
  return {
    authorize: wk.authorization_endpoint,
    token: wk.token_endpoint,
    userinfo: wk.userinfo_endpoint,
    jwks: wk.jwks_uri,
    logout: wk.end_session_endpoint
  };
}
