import type {
  JobDTO,
  JobListDTO,
  CreateJobRequest,
  UpdateJobRequest,
  JobListRequest,
} from '@skillsier/types';
import { mockJobListResponse, mockJobs, getMockJob } from './jobs.mock';

export class MockJobsClient {
  async listJobs(params?: JobListRequest): Promise<JobListDTO> {
    await new Promise(resolve => setTimeout(resolve, 500));
    
    let filteredJobs = [...mockJobs];

    if (params?.search) {
      const searchLower = params.search.toLowerCase();
      filteredJobs = filteredJobs.filter(
        job =>
          job.title.toLowerCase().includes(searchLower) ||
          job.description.toLowerCase().includes(searchLower)
      );
    }

    if (params?.status) {
      filteredJobs = filteredJobs.filter(job => job.status === params.status);
    }

    if (params?.budget_type) {
      filteredJobs = filteredJobs.filter(
        job => job.budget_type === params.budget_type
      );
    }

    if (params?.skills && params.skills.length > 0) {
      filteredJobs = filteredJobs.filter(job =>
        params.skills!.some(skill => job.skills.includes(skill))
      );
    }

    const page = params?.page || 1;
    const limit = params?.limit || 10;
    const start = (page - 1) * limit;
    const end = start + limit;
    const paginatedJobs = filteredJobs.slice(start, end);

    return {
      jobs: paginatedJobs,
      pagination: {
        page,
        limit,
        total: filteredJobs.length,
        total_pages: Math.ceil(filteredJobs.length / limit),
      },
    };
  }

  async getJob(jobId: string): Promise<JobDTO> {
    await new Promise(resolve => setTimeout(resolve, 300));
    
    const job = getMockJob(jobId);
    if (!job) {
      throw new Error('Job not found');
    }
    return job;
  }

  async createJob(data: CreateJobRequest): Promise<JobDTO> {
    await new Promise(resolve => setTimeout(resolve, 500));
    
    const newJob: JobDTO = {
      job_id: `job-${Date.now()}`,
      title: data.title,
      description: data.description,
      status: 'DRAFT',
      budget_type: data.budget_type,
      budget_amount: data.budget_amount,
      budget_min: data.budget_min,
      budget_max: data.budget_max,
      currency: data.currency,
      duration: data.duration,
      experience_level: data.experience_level,
      visibility: data.visibility,
      skills: data.skills,
      category: data.category,
      subcategory: data.subcategory,
      client_id: 'current-user',
      proposal_count: 0,
      view_count: 0,
      deadline: data.deadline,
      created_at: new Date().toISOString(),
      updated_at: new Date().toISOString(),
    };
    
    return newJob;
  }

  async updateJob(jobId: string, data: UpdateJobRequest): Promise<JobDTO> {
    await new Promise(resolve => setTimeout(resolve, 500));
    
    const job = getMockJob(jobId);
    if (!job) {
      throw new Error('Job not found');
    }

    return {
      ...job,
      ...data,
      updated_at: new Date().toISOString(),
    };
  }

  async deleteJob(jobId: string): Promise<void> {
    await new Promise(resolve => setTimeout(resolve, 300));
  }

  async publishJob(jobId: string): Promise<JobDTO> {
    await new Promise(resolve => setTimeout(resolve, 500));
    
    const job = getMockJob(jobId);
    if (!job) {
      throw new Error('Job not found');
    }

    return {
      ...job,
      status: 'OPEN',
      published_at: new Date().toISOString(),
      updated_at: new Date().toISOString(),
    };
  }

  async closeJob(jobId: string): Promise<JobDTO> {
    await new Promise(resolve => setTimeout(resolve, 500));
    
    const job = getMockJob(jobId);
    if (!job) {
      throw new Error('Job not found');
    }

    return {
      ...job,
      status: 'CLOSED',
      closed_at: new Date().toISOString(),
      updated_at: new Date().toISOString(),
    };
  }
}