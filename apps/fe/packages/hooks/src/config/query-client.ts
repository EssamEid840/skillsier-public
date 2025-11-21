import { QueryClient } from '@tanstack/react-query';

export const queryClient = new QueryClient({
  defaultOptions: {
    queries: {
      staleTime: 1000 * 60 * 5, // 5 minutes
      gcTime: 1000 * 60 * 10, // 10 minutes (formerly cacheTime)
      retry: 1,
      refetchOnWindowFocus: false,
    },
    mutations: {
      retry: 0,
    },
  },
});

export const QUERY_KEYS = {
  // Jobs
  jobs: ['jobs'] as const,
  jobsList: (filters?: Record<string, any>) => ['jobs', 'list', filters] as const,
  job: (id: string) => ['jobs', id] as const,
  
  // Users
  users: ['users'] as const,
  currentUser: ['users', 'me'] as const,
  user: (id: string) => ['users', id] as const,
  userProfile: (id: string) => ['users', id, 'profile'] as const,
  
  // Proposals
  proposals: ['proposals'] as const,
  proposalsList: (filters?: Record<string, any>) => ['proposals', 'list', filters] as const,
  proposal: (id: string) => ['proposals', id] as const,
  
  // Contracts
  contracts: ['contracts'] as const,
  contractsList: (filters?: Record<string, any>) => ['contracts', 'list', filters] as const,
  contract: (id: string) => ['contracts', id] as const,
} as const;