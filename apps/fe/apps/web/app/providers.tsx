'use client';

import * as React from 'react';
import { QueryClientProvider } from '@tanstack/react-query';
import { queryClient } from '@skillsier/hooks';
import { AuthProvider, createAuthAdapter } from '@skillsier/auth';

const authAdapter = createAuthAdapter(
  (process.env.NEXT_PUBLIC_AUTH_PROVIDER as 'dev' | 'keycloak') || 'dev'
);

export function Providers({ children }: { children: React.ReactNode }) {
  return (
    <QueryClientProvider client={queryClient}>
      <AuthProvider adapter={authAdapter}>{children}</AuthProvider>
    </QueryClientProvider>
  );
}