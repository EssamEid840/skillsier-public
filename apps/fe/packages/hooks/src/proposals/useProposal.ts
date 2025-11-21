import { useQuery } from '@tanstack/react-query';
import { ProposalsClient } from '@skillsier/api';
import { QUERY_KEYS } from '../config/query-client';

const proposalsClient = new ProposalsClient({
  baseURL: process.env.API_BASE_URL || 'http://localhost:8080',
});

export const useProposal = (proposalId: string, enabled: boolean = true) => {
  return useQuery({
    queryKey: QUERY_KEYS.proposal(proposalId),
    queryFn: async () => {
      return proposalsClient.getProposal(proposalId);
    },
    enabled: enabled && !!proposalId,
  });
};