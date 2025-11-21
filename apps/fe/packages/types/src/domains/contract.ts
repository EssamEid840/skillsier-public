import type { ContractStatus } from '../enums/status.enum';

export interface Contract {
  id: string;
  jobId: string;
  proposalId: string;
  clientId: string;
  freelancerId: string;
  status: ContractStatus;
  title: string;
  description: string;
  amount: number;
  currency: string;
  startDate: Date;
  endDate?: Date;
  milestones?: ContractMilestone[];
  terms?: string;
  createdAt: Date;
  updatedAt: Date;
  signedAt?: Date;
  completedAt?: Date;
}

export interface ContractMilestone {
  id: string;
  title: string;
  description?: string;
  amount: number;
  dueDate?: Date;
  status: 'PENDING' | 'IN_PROGRESS' | 'COMPLETED' | 'APPROVED' | 'DISPUTED';
  completedAt?: Date;
  approvedAt?: Date;
}