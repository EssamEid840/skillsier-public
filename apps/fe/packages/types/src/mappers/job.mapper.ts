import type { Job, JobListResponse } from '../domains/job';
import type { JobDTO, JobListDTO } from '../models/job.model';
import {
  JobStatus,
  JobBudgetType,
  JobDuration,
  ExperienceLevel,
  JobVisibility,
} from '../enums/job.enum';

export const mapJobDTOToDomain = (dto: JobDTO): Job => ({
  id: dto.job_id,
  title: dto.title,
  description: dto.description,
  status: dto.status as JobStatus,
  budgetType: dto.budget_type as JobBudgetType,
  budgetAmount: dto.budget_amount,
  budgetMin: dto.budget_min,
  budgetMax: dto.budget_max,
  currency: dto.currency,
  duration: dto.duration as JobDuration,
  experienceLevel: dto.experience_level as ExperienceLevel,
  visibility: dto.visibility as JobVisibility,
  skills: dto.skills,
  category: dto.category,
  subcategory: dto.subcategory,
  clientId: dto.client_id,
  proposalCount: dto.proposal_count,
  viewCount: dto.view_count,
  deadline: dto.deadline ? new Date(dto.deadline) : undefined,
  createdAt: new Date(dto.created_at),
  updatedAt: new Date(dto.updated_at),
  publishedAt: dto.published_at ? new Date(dto.published_at) : undefined,
  closedAt: dto.closed_at ? new Date(dto.closed_at) : undefined,
});

export const mapJobListDTOToDomain = (dto: JobListDTO): JobListResponse => ({
  jobs: dto.jobs.map(mapJobDTOToDomain),
  pagination: {
    page: dto.pagination.page,
    limit: dto.pagination.limit,
    total: dto.pagination.total,
    totalPages: dto.pagination.total_pages,
  },
});