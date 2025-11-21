import { create } from 'zustand';
import { persist } from 'zustand/middleware';
import type { Job, JobFilters } from '@skillsier/types';

interface JobsState {
  // State
  selectedJob: Job | null;
  filters: JobFilters;
  page: number;
  limit: number;
  sortBy: 'recent' | 'budget' | 'proposals';
  viewMode: 'grid' | 'list';

  // Actions
  setSelectedJob: (job: Job | null) => void;
  setFilters: (filters: Partial<JobFilters>) => void;
  resetFilters: () => void;
  setPage: (page: number) => void;
  setLimit: (limit: number) => void;
  setSortBy: (sortBy: 'recent' | 'budget' | 'proposals') => void;
  setViewMode: (mode: 'grid' | 'list') => void;
}

const defaultFilters: JobFilters = {
  search: undefined,
  status: undefined,
  budgetType: undefined,
  budgetMin: undefined,
  budgetMax: undefined,
  skills: undefined,
  category: undefined,
  experienceLevel: undefined,
  duration: undefined,
};

export const useJobsStore = create<JobsState>()(
  persist(
    (set) => ({
      // Initial state
      selectedJob: null,
      filters: defaultFilters,
      page: 1,
      limit: 10,
      sortBy: 'recent',
      viewMode: 'grid',

      // Actions
      setSelectedJob: (job) => set({ selectedJob: job }),

      setFilters: (newFilters) =>
        set((state) => ({
          filters: { ...state.filters, ...newFilters },
          page: 1, // Reset to page 1 when filters change
        })),

      resetFilters: () =>
        set({
          filters: defaultFilters,
          page: 1,
        }),

      setPage: (page) => set({ page }),

      setLimit: (limit) => set({ limit, page: 1 }),

      setSortBy: (sortBy) => set({ sortBy }),

      setViewMode: (mode) => set({ viewMode: mode }),
    }),
    {
      name: 'skillsier-jobs-storage',
      partialize: (state) => ({
        filters: state.filters,
        limit: state.limit,
        sortBy: state.sortBy,
        viewMode: state.viewMode,
      }),
    }
  )
);