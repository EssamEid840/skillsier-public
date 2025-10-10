// packages/types/src/entities/proposal.ts
// Proposal entity types

export type ProposalStatus = 'pending' | 'accepted' | 'rejected' | 'withdrawn';

export interface Proposal {
  id: string;
  jobId: string;
  freelancerId: string;
  coverLetter: string;
  bidAmount: number;
  estimatedDuration: string;
  milestones?: Milestone[];
  status: ProposalStatus;
  attachments: string[];
  createdAt: Date;
  updatedAt: Date;
}

export interface Milestone {
  id: string;
  description: string;
  amount: number;
  dueDate?: Date;
  status: 'pending' | 'in_progress' | 'completed';
}

export interface ProposalDetails extends Proposal {
  job: {
    id: string;
    title: string;
    budget?: number;
    type: string;
  };
  freelancer: {
    id: string;
    name: string;
    title: string;
    rating: number;
    reviewCount: number;
    hourlyRate: number;
    completionRate: number;
    totalJobs: number;
  };
}

export interface CreateProposalRequest {
  jobId: string;
  coverLetter: string;
  bidAmount: number;
  estimatedDuration: string;
  milestones?: Omit<Milestone, 'id' | 'status'>[];
}

export interface UpdateProposalRequest extends Partial<CreateProposalRequest> {
  status?: ProposalStatus;
}