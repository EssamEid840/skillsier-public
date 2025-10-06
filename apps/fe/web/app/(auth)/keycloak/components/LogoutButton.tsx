'use client';

import { useState } from 'react';

export default function LogoutButton({ redirectTo = '/' }: { redirectTo?: string }) {
  const [busy, setBusy] = useState(false);
  const onClick = async () => {
    if (busy) return;
    setBusy(true);
    try {
      await fetch('/api/auth/keycloak/logout', { method: 'POST' });
    } finally {
      window.location.href = redirectTo;
    }
  };
  return (
    <button
      onClick={onClick}
      disabled={busy}
      className="px-3 py-2 rounded-xl border"
      aria-busy={busy}
    >
      {busy ? 'Signing out…' : 'Logout'}
    </button>
  );
}
