import type { ProposalStatus } from '../enums/status.enum';

export interface Proposal {
  id: string;
  jobId: string;
  freelancerId: string;
  status: ProposalStatus;
  coverLetter: string;
  proposedRate?: number;
  proposedAmount?: number;
  currency: string;
  estimatedDuration?: string;
  deliverables?: string;
  milestones?: ProposalMilestone[];
  attachments?: string[];
  createdAt: Date;
  updatedAt: Date;
  submittedAt?: Date;
}

export interface ProposalMilestone {
  title: string;
  description?: string;
  amount: number;
  dueDate?: Date;
}

export interface ProposalFilters {
  status?: ProposalStatus;
  jobId?: string;
  freelancerId?: string;
}