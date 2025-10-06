"use client";

import { useEffect } from "react";

export default function Page() {
  useEffect(() => {
    (async () => {
      try {
        await fetch("/api/auth/keycloak/logout", { method: "POST" });
      } finally {
        window.location.replace("/");
      }
    })();
  }, []);

  return (
    <main style={{ padding: 24 }}>
      <p>Signing you out…</p>
    </main>
  );
}
