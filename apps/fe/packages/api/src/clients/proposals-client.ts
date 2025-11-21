import { BaseApiClient, type ApiClientConfig } from '../lib/base-client';
import type {
  ProposalDTO,
  CreateProposalRequest,
  UpdateProposalRequest,
} from '@skillsier/types';

export class ProposalsClient extends BaseApiClient {
  constructor(config: ApiClientConfig) {
    super(config);
  }

  async listProposals(params?: {
    job_id?: string;
    status?: string;
  }): Promise<ProposalDTO[]> {
    const queryParams = new URLSearchParams();
    if (params?.job_id) queryParams.append('job_id', params.job_id);
    if (params?.status) queryParams.append('status', params.status);

    const query = queryParams.toString();
    return this.get<ProposalDTO[]>(`/v1/proposals${query ? `?${query}` : ''}`);
  }

  async getProposal(proposalId: string): Promise<ProposalDTO> {
    return this.get<ProposalDTO>(`/v1/proposals/${proposalId}`);
  }

  async createProposal(data: CreateProposalRequest): Promise<ProposalDTO> {
    return this.post<ProposalDTO>('/v1/proposals', data);
  }

  async updateProposal(
    proposalId: string,
    data: UpdateProposalRequest
  ): Promise<ProposalDTO> {
    return this.patch<ProposalDTO>(`/v1/proposals/${proposalId}`, data);
  }

  async submitProposal(proposalId: string): Promise<ProposalDTO> {
    return this.post<ProposalDTO>(`/v1/proposals/${proposalId}/submit`);
  }

  async withdrawProposal(proposalId: string): Promise<ProposalDTO> {
    return this.post<ProposalDTO>(`/v1/proposals/${proposalId}/withdraw`);
  }
}