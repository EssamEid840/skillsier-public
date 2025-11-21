import { describe, it, expect } from 'vitest';
import { MockJobsClient } from '../src/mocks/jobs-client.mock';

describe('MockJobsClient', () => {
  const client = new MockJobsClient();

  it('lists jobs successfully', async () => {
    const result = await client.listJobs({ page: 1, limit: 10 });
    
    expect(result.jobs).toBeDefined();
    expect(result.jobs.length).toBeGreaterThan(0);
    expect(result.pagination.page).toBe(1);
    expect(result.pagination.limit).toBe(10);
  });

  it('gets a single job by id', async () => {
    const job = await client.getJob('1');
    
    expect(job).toBeDefined();
    expect(job.job_id).toBe('1');
    expect(job.title).toBeTruthy();
  });

  it('filters jobs by search term', async () => {
    const result = await client.listJobs({ search: 'Full-Stack' });
    
    expect(result.jobs.length).toBeGreaterThan(0);
    expect(result.jobs[0]?.title.includes('Full-Stack')).toBe(true);
  });

  it('creates a new job', async () => {
    const newJob = await client.createJob({
      title: 'Test Job',
      description: 'Test Description',
      budget_type: 'HOURLY',
      budget_min: 30,
      budget_max: 60,
      currency: 'USD',
      duration: 'ONE_TO_THREE_MONTHS',
      experience_level: 'INTERMEDIATE',
      visibility: 'PUBLIC',
      skills: ['JavaScript'],
      category: 'Development',
    });

    expect(newJob).toBeDefined();
    expect(newJob.title).toBe('Test Job');
    expect(newJob.status).toBe('DRAFT');
  });
});