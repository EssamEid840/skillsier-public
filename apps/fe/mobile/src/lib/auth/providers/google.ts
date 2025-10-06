import { startAuth } from "../keycloak.client";
export const signInWithGoogle = () => startAuth({ useGoogle: true });
