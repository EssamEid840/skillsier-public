import { useMutation, useQueryClient } from '@tanstack/react-query';
import { ProposalsClient } from '@skillsier/api';
import type { CreateProposalRequest } from '@skillsier/types';
import { QUERY_KEYS } from '../config/query-client';

const proposalsClient = new ProposalsClient({
  baseURL: process.env.API_BASE_URL || 'http://localhost:8080',
});

export const useCreateProposal = () => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: async (data: CreateProposalRequest) => {
      return proposalsClient.createProposal(data);
    },
    onSuccess: (data) => {
      queryClient.invalidateQueries({ queryKey: QUERY_KEYS.proposals });
      queryClient.invalidateQueries({ queryKey: QUERY_KEYS.job(data.job_id) });
    },
  });
};