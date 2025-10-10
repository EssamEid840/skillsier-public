// packages/types/src/entities/job.ts
// Job/Gig entity types

export type JobStatus = 'draft' | 'open' | 'in_progress' | 'completed' | 'cancelled';
export type JobType = 'fixed' | 'hourly';
export type ExperienceLevel = 'entry' | 'intermediate' | 'expert';

export interface Job {
  id: string;
  clientId: string;
  title: string;
  description: string;
  category: string;
  subcategory?: string;
  type: JobType;
  budget?: number;
  hourlyRateMin?: number;
  hourlyRateMax?: number;
  estimatedDuration?: string;
  experienceLevel: ExperienceLevel;
  skills: string[];
  status: JobStatus;
  attachments: string[];
  proposals: number;
  deadline?: Date;
  createdAt: Date;
  updatedAt: Date;
}

export interface JobDetails extends Job {
  client: {
    id: string;
    name: string;
    companyName?: string;
    rating: number;
    reviewCount: number;
    totalSpent: number;
    location?: string;
  };
}

export interface JobFilters {
  category?: string;
  subcategory?: string;
  type?: JobType;
  budgetMin?: number;
  budgetMax?: number;
  experienceLevel?: ExperienceLevel;
  skills?: string[];
  search?: string;
}

export interface CreateJobRequest {
  title: string;
  description: string;
  category: string;
  subcategory?: string;
  type: JobType;
  budget?: number;
  hourlyRateMin?: number;
  hourlyRateMax?: number;
  estimatedDuration?: string;
  experienceLevel: ExperienceLevel;
  skills: string[];
  deadline?: Date;
}

export interface UpdateJobRequest extends Partial<CreateJobRequest> {
  status?: JobStatus;
}