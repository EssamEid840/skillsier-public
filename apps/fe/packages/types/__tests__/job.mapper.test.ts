import { describe, it, expect } from 'vitest';
import { mapJobDTOToDomain } from '../src/mappers/job.mapper';
import type { JobDTO } from '../src/models/job.model';
import { JobStatus, JobBudgetType, JobDuration, ExperienceLevel, JobVisibility } from '../src/enums/job.enum';

describe('Job Mapper', () => {
  it('maps JobDTO to Job domain', () => {
    const dto: JobDTO = {
      job_id: '123',
      title: 'Senior Developer',
      description: 'Build amazing apps',
      status: 'OPEN',
      budget_type: 'HOURLY',
      budget_min: 50,
      budget_max: 100,
      currency: 'USD',
      duration: 'ONE_TO_THREE_MONTHS',
      experience_level: 'EXPERT',
      visibility: 'PUBLIC',
      skills: ['React', 'TypeScript'],
      category: 'Development',
      client_id: 'client-1',
      proposal_count: 5,
      view_count: 100,
      created_at: '2024-01-01T00:00:00Z',
      updated_at: '2024-01-01T00:00:00Z',
    };

    const domain = mapJobDTOToDomain(dto);

    expect(domain.id).toBe('123');
    expect(domain.title).toBe('Senior Developer');
    expect(domain.status).toBe(JobStatus.OPEN);
    expect(domain.budgetType).toBe(JobBudgetType.HOURLY);
    expect(domain.createdAt).toBeInstanceOf(Date);
  });
});