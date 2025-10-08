export const queryKeys = {
  auth: {
    me: ['auth', 'me'] as const,
  },
  users: {
    profile: ['users', 'profile'] as const,
    freelancerProfile: ['users', 'freelancer-profile'] as const,
    clientProfile: ['users', 'client-profile'] as const,
    skills: ['users', 'skills'] as const,
    workExperience: ['users', 'work-experience'] as const,
    education: ['users', 'education'] as const,
    certifications: ['users', 'certifications'] as const,
    portfolio: ['users', 'portfolio'] as const,
    stats: ['users', 'stats'] as const,
    earnings: ['users', 'earnings'] as const,
    reviews: ['users', 'reviews'] as const,
    byId: (id: string) => ['users', id] as const,
  },
  jobs: {
    all: ['jobs'] as const,
    list: (filters?: Record<string, unknown>) => ['jobs', 'list', filters] as const,
    detail: (id: string) => ['jobs', 'detail', id] as const,
    myJobs: ['jobs', 'my-jobs'] as const,
  },
  proposals: {
    all: ['proposals'] as const,
    myProposals: ['proposals', 'my-proposals'] as const,
    detail: (id: string) => ['proposals', 'detail', id] as const,
  },
  contracts: {
    all: ['contracts'] as const,
    myContracts: ['contracts', 'my-contracts'] as const,
    detail: (id: string) => ['contracts', 'detail', id] as const,
  },
} as const;
