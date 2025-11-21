export interface ProposalDTO {
  proposal_id: string;
  job_id: string;
  freelancer_id: string;
  status: string;
  cover_letter: string;
  proposed_rate?: number;
  proposed_amount?: number;
  currency: string;
  estimated_duration?: string;
  deliverables?: string;
  milestones?: {
    title: string;
    description?: string;
    amount: number;
    due_date?: string;
  }[];
  attachments?: string[];
  created_at: string;
  updated_at: string;
  submitted_at?: string;
}

export interface CreateProposalRequest {
  job_id: string;
  cover_letter: string;
  proposed_rate?: number;
  proposed_amount?: number;
  currency: string;
  estimated_duration?: string;
  deliverables?: string;
  milestones?: {
    title: string;
    description?: string;
    amount: number;
    due_date?: string;
  }[];
  attachments?: string[];
}

export interface UpdateProposalRequest {
  cover_letter?: string;
  proposed_rate?: number;
  proposed_amount?: number;
  estimated_duration?: string;
  deliverables?: string;
}