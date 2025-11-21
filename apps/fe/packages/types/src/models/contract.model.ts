export interface ContractDTO {
  contract_id: string;
  job_id: string;
  proposal_id: string;
  client_id: string;
  freelancer_id: string;
  status: string;
  title: string;
  description: string;
  amount: number;
  currency: string;
  start_date: string;
  end_date?: string;
  milestones?: {
    milestone_id: string;
    title: string;
    description?: string;
    amount: number;
    due_date?: string;
    status: string;
    completed_at?: string;
    approved_at?: string;
  }[];
  terms?: string;
  created_at: string;
  updated_at: string;
  signed_at?: string;
  completed_at?: string;
}

export interface CreateContractRequest {
  proposal_id: string;
  title: string;
  description: string;
  amount: number;
  currency: string;
  start_date: string;
  end_date?: string;
  milestones?: {
    title: string;
    description?: string;
    amount: number;
    due_date?: string;
  }[];
  terms?: string;
}