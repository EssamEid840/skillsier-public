"use client";

import { useEffect } from "react";

export default function Page() {
  useEffect(() => {
    // Start the Google OAuth flow via Keycloak broker (BFF endpoint)
    window.location.assign("/api/auth/keycloak/oauth/start/google");
  }, []);

  return (
    <main style={{ padding: 24 }}>
      <p>Redirecting to Google…</p>
    </main>
  );
}
