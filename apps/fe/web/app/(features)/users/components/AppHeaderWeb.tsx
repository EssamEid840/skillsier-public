import Link from 'next/link';
import { headers } from 'next/headers';

async function getMe() {
  const h = await headers();
  const cookie = h.get('cookie') ?? '';
  const proto = h.get('x-forwarded-proto') ?? 'http';
  const host = h.get('host') ?? 'localhost:3000';
  const base = process.env.NEXT_PUBLIC_BASE_URL ?? `${proto}://${host}`;

  const res = await fetch(`${base}/api/auth/keycloak/me`, {
    headers: { cookie },
    cache: 'no-store',
  }).catch(() => null);

  if (!res || !res.ok) return null;
  return res.json();
}

export default async function AppHeaderWeb() {
  const me = await getMe();

  return (
    <header className="flex items-center justify-between p-4 border-b">
      <Link href="/" className="font-semibold">Skillsier</Link>
      <nav className="flex items-center gap-3">
        {me ? (
          <>
            <span className="text-sm opacity-75">
              {me.email ?? me.preferred_username ?? 'Signed in'}
            </span>
            <a href="/api/auth/keycloak/logout" className="px-3 py-2 rounded-xl border">
              Logout
            </a>
          </>
        ) : (
          <Link href="/keycloak/login" className="px-3 py-2 rounded-xl border">Login</Link>
        )}
      </nav>
    </header>
  );
}
