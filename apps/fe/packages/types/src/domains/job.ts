import type {
  JobStatus,
  JobBudgetType,
  JobDuration,
  ExperienceLevel,
  JobVisibility,
} from '../enums/job.enum';

export interface Job {
  id: string;
  title: string;
  description: string;
  status: JobStatus;
  budgetType: JobBudgetType;
  budgetAmount?: number;
  budgetMin?: number;
  budgetMax?: number;
  currency: string;
  duration: JobDuration;
  experienceLevel: ExperienceLevel;
  visibility: JobVisibility;
  skills: string[];
  category: string;
  subcategory?: string;
  clientId: string;
  proposalCount: number;
  viewCount: number;
  deadline?: Date;
  createdAt: Date;
  updatedAt: Date;
  publishedAt?: Date;
  closedAt?: Date;
}

export interface JobFilters {
  search?: string;
  status?: JobStatus;
  budgetType?: JobBudgetType;
  budgetMin?: number;
  budgetMax?: number;
  skills?: string[];
  category?: string;
  experienceLevel?: ExperienceLevel;
  duration?: JobDuration;
}

export interface JobPagination {
  page: number;
  limit: number;
  total: number;
  totalPages: number;
}

export interface JobListResponse {
  jobs: Job[];
  pagination: JobPagination;
}