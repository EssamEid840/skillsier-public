export interface JobDTO {
  job_id: string;
  title: string;
  description: string;
  status: string;
  budget_type: string;
  budget_amount?: number;
  budget_min?: number;
  budget_max?: number;
  currency: string;
  duration: string;
  experience_level: string;
  visibility: string;
  skills: string[];
  category: string;
  subcategory?: string;
  client_id: string;
  proposal_count: number;
  view_count: number;
  deadline?: string;
  created_at: string;
  updated_at: string;
  published_at?: string;
  closed_at?: string;
}

export interface CreateJobRequest {
  title: string;
  description: string;
  budget_type: string;
  budget_amount?: number;
  budget_min?: number;
  budget_max?: number;
  currency: string;
  duration: string;
  experience_level: string;
  visibility: string;
  skills: string[];
  category: string;
  subcategory?: string;
  deadline?: string;
}

export interface UpdateJobRequest {
  title?: string;
  description?: string;
  budget_amount?: number;
  budget_min?: number;
  budget_max?: number;
  deadline?: string;
  skills?: string[];
}

export interface JobListRequest {
  page?: number;
  limit?: number;
  search?: string;
  status?: string;
  budget_type?: string;
  budget_min?: number;
  budget_max?: number;
  skills?: string[];
  category?: string;
  experience_level?: string;
  duration?: string;
}

export interface JobListDTO {
  jobs: JobDTO[];
  pagination: {
    page: number;
    limit: number;
    total: number;
    total_pages: number;
  };
}