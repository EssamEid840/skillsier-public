import { BaseApiClient, type ApiClientConfig } from '../lib/base-client';
import type {
  JobDTO,
  JobListDTO,
  CreateJobRequest,
  UpdateJobRequest,
  JobListRequest,
} from '@skillsier/types';

export class JobsClient extends BaseApiClient {
  constructor(config: ApiClientConfig) {
    super(config);
  }

  async listJobs(params?: JobListRequest): Promise<JobListDTO> {
    const queryParams = new URLSearchParams();
    
    if (params?.page) queryParams.append('page', params.page.toString());
    if (params?.limit) queryParams.append('limit', params.limit.toString());
    if (params?.search) queryParams.append('search', params.search);
    if (params?.status) queryParams.append('status', params.status);
    if (params?.budget_type) queryParams.append('budget_type', params.budget_type);
    if (params?.budget_min) queryParams.append('budget_min', params.budget_min.toString());
    if (params?.budget_max) queryParams.append('budget_max', params.budget_max.toString());
    if (params?.category) queryParams.append('category', params.category);
    if (params?.experience_level) queryParams.append('experience_level', params.experience_level);
    if (params?.duration) queryParams.append('duration', params.duration);
    if (params?.skills) {
      params.skills.forEach(skill => queryParams.append('skills[]', skill));
    }

    const query = queryParams.toString();
    return this.get<JobListDTO>(`/v1/jobs${query ? `?${query}` : ''}`);
  }

  async getJob(jobId: string): Promise<JobDTO> {
    return this.get<JobDTO>(`/v1/jobs/${jobId}`);
  }

  async createJob(data: CreateJobRequest): Promise<JobDTO> {
    return this.post<JobDTO>('/v1/jobs', data);
  }

  async updateJob(jobId: string, data: UpdateJobRequest): Promise<JobDTO> {
    return this.patch<JobDTO>(`/v1/jobs/${jobId}`, data);
  }

  async deleteJob(jobId: string): Promise<void> {
    return this.delete<void>(`/v1/jobs/${jobId}`);
  }

  async publishJob(jobId: string): Promise<JobDTO> {
    return this.post<JobDTO>(`/v1/jobs/${jobId}/publish`);
  }

  async closeJob(jobId: string): Promise<JobDTO> {
    return this.post<JobDTO>(`/v1/jobs/${jobId}/close`);
  }
}