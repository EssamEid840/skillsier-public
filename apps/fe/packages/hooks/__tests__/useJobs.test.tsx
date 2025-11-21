import { describe, it, expect } from 'vitest';
import { renderHook, waitFor } from '@testing-library/react';
import { QueryClient, QueryClientProvider } from '@tanstack/react-query';
import { useJobs } from '../src/jobs/useJobs';
import * as React from 'react';

const createWrapper = () => {
  const queryClient = new QueryClient({
    defaultOptions: {
      queries: {
        retry: false,
      },
    },
  });

  return ({ children }: { children: React.ReactNode }) => (
    <QueryClientProvider client={queryClient}>{children}</QueryClientProvider>
  );
};

describe('useJobs', () => {
  it('fetches jobs successfully', async () => {
    const { result } = renderHook(() => useJobs(), {
      wrapper: createWrapper(),
    });

    expect(result.current.isLoading).toBe(true);

    await waitFor(() => {
      expect(result.current.isSuccess).toBe(true);
    });

    expect(result.current.data?.jobs).toBeDefined();
    expect(result.current.data?.jobs.length).toBeGreaterThan(0);
    expect(result.current.data?.pagination).toBeDefined();
  });

  it('applies filters correctly', async () => {
    const { result } = renderHook(
      () =>
        useJobs({
          search: 'Full-Stack',
        }),
      {
        wrapper: createWrapper(),
      }
    );

    await waitFor(() => {
      expect(result.current.isSuccess).toBe(true);
    });

    expect(result.current.data?.jobs).toBeDefined();
  });
});