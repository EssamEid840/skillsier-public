// packages/types/src/entities/contract.ts
// Contract entity types

export type ContractStatus = 'active' | 'completed' | 'cancelled' | 'disputed';
export type PaymentStatus = 'pending' | 'escrowed' | 'released' | 'refunded';

export interface Contract {
  id: string;
  jobId: string;
  proposalId: string;
  clientId: string;
  freelancerId: string;
  title: string;
  description: string;
  amount: number;
  status: ContractStatus;
  startDate: Date;
  endDate?: Date;
  terms: string;
  milestones: ContractMilestone[];
  createdAt: Date;
  updatedAt: Date;
}

export interface ContractMilestone {
  id: string;
  contractId: string;
  description: string;
  amount: number;
  dueDate?: Date;
  status: 'pending' | 'in_progress' | 'submitted' | 'approved' | 'paid';
  paymentStatus: PaymentStatus;
  submittedAt?: Date;
  approvedAt?: Date;
  paidAt?: Date;
}

export interface ContractDetails extends Contract {
  job: {
    id: string;
    title: string;
    category: string;
  };
  client: {
    id: string;
    name: string;
    companyName?: string;
    rating: number;
  };
  freelancer: {
    id: string;
    name: string;
    title: string;
    rating: number;
  };
}

export interface CreateContractRequest {
  jobId: string;
  proposalId: string;
  terms: string;
  milestones: Omit<ContractMilestone, 'id' | 'contractId' | 'status' | 'paymentStatus' | 'submittedAt' | 'approvedAt' | 'paidAt'>[];
}

export interface UpdateContractRequest {
  status?: ContractStatus;
  endDate?: Date;
}

export interface SubmitMilestoneRequest {
  contractId: string;
  milestoneId: string;
  deliverables: string;
  attachments?: string[];
}

export interface ApproveMilestoneRequest {
  contractId: string;
  milestoneId: string;
  feedback?: string;
}