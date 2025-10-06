'use client';

type Props = {
  /** Where to land after a successful login (default: /users) */
  returnTo?: string;
  className?: string;
};

export default function OAuthGoogleButton({ returnTo = '/users', className }: Props) {
  const href = `/api/auth/keycloak/oauth/start/google?ru=${encodeURIComponent(returnTo)}`;

  return (
    <a
      href={href}
      className={
        className ??
        'inline-flex items-center justify-center rounded-md border px-4 py-2 text-sm font-medium hover:bg-muted'
      }
      aria-label="Continue with Google"
    >
      <svg aria-hidden="true" width="18" height="18" viewBox="0 0 48 48" className="mr-2">
        <path fill="#FFC107" d="M43.6 20.5H42V20H24v8h11.3A12.9 12.9 0 0 1 24 36.9a13 13 0 1 1 9.2-22.2l5.7-5.7A21 21 0 1 0 24 45c10.5 0 21-7.6 21-21 0-1.6-.2-2.8-.4-3.5z"/>
        <path fill="#FF3D00" d="M6.3 14.7l6.6 4.8A12.9 12.9 0 0 1 24 11.1c3.2 0 6 .9 8.4 2.6l6.3-6.3A21 21 0 0 0 3 14.7z"/>
        <path fill="#4CAF50" d="M24 45c5.4 0 10.4-1.8 14.3-4.9l-6.6-5a13 13 0 0 1-19.1-6.8l-6.6 5.1A21 21 0 0 0 24 45z"/>
        <path fill="#1976D2" d="M43.6 20.5H42V20H24v8h11.3c-1.3 3.9-4.9 6.9-9.3 6.9v0c2.3 0 4.5-.7 6.3-1.9l6.6 5C36.7 41.2 31.7 43 26.3 43c11.7 0 21.3-9.3 21.3-21 0-1.6-.2-2.8-.4-3.5z"/>
      </svg>
      Continue with Google
    </a>
  );
}
