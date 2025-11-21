import { describe, it, expect, beforeEach } from 'vitest';
import { useJobsStore } from '../src/jobs.store';
import { JobStatus, JobBudgetType } from '@skillsier/types';

describe('Jobs Store', () => {
  beforeEach(() => {
    const store = useJobsStore.getState();
    store.resetFilters();
    store.setPage(1);
    store.setViewMode('grid');
  });

  it('initializes with default values', () => {
    const state = useJobsStore.getState();
    
    expect(state.selectedJob).toBeNull();
    expect(state.page).toBe(1);
    expect(state.limit).toBe(10);
    expect(state.sortBy).toBe('recent');
    expect(state.viewMode).toBe('grid');
  });

  it('sets filters correctly', () => {
    const store = useJobsStore.getState();
    
    store.setFilters({
      search: 'developer',
      status: JobStatus.OPEN,
      budgetType: JobBudgetType.HOURLY,
    });

    const state = useJobsStore.getState();
    expect(state.filters.search).toBe('developer');
    expect(state.filters.status).toBe(JobStatus.OPEN);
    expect(state.filters.budgetType).toBe(JobBudgetType.HOURLY);
    expect(state.page).toBe(1); // Should reset to page 1
  });

  it('resets filters', () => {
    const store = useJobsStore.getState();
    
    store.setFilters({ search: 'test', status: JobStatus.OPEN });
    store.resetFilters();

    const state = useJobsStore.getState();
    expect(state.filters.search).toBeUndefined();
    expect(state.filters.status).toBeUndefined();
  });

  it('changes view mode', () => {
    const store = useJobsStore.getState();
    
    store.setViewMode('list');
    expect(useJobsStore.getState().viewMode).toBe('list');
    
    store.setViewMode('grid');
    expect(useJobsStore.getState().viewMode).toBe('grid');
  });

  it('sets pagination', () => {
    const store = useJobsStore.getState();
    
    store.setPage(3);
    expect(useJobsStore.getState().page).toBe(3);
    
    store.setLimit(20);
    const state = useJobsStore.getState();
    expect(state.limit).toBe(20);
    expect(state.page).toBe(1); // Should reset to page 1
  });
});