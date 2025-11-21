import { create } from 'zustand';
import type { Proposal, ProposalFilters } from '@skillsier/types';

interface ProposalsState {
  // State
  selectedProposal: Proposal | null;
  filters: ProposalFilters;
  draftProposal: Partial<Proposal> | null;

  // Actions
  setSelectedProposal: (proposal: Proposal | null) => void;
  setFilters: (filters: Partial<ProposalFilters>) => void;
  resetFilters: () => void;
  setDraftProposal: (draft: Partial<Proposal> | null) => void;
  updateDraftProposal: (updates: Partial<Proposal>) => void;
  clearDraftProposal: () => void;
}

const defaultFilters: ProposalFilters = {
  status: undefined,
  jobId: undefined,
  freelancerId: undefined,
};

export const useProposalsStore = create<ProposalsState>((set) => ({
  // Initial state
  selectedProposal: null,
  filters: defaultFilters,
  draftProposal: null,

  // Actions
  setSelectedProposal: (proposal) => set({ selectedProposal: proposal }),

  setFilters: (newFilters) =>
    set((state) => ({
      filters: { ...state.filters, ...newFilters },
    })),

  resetFilters: () => set({ filters: defaultFilters }),

  setDraftProposal: (draft) => set({ draftProposal: draft }),

  updateDraftProposal: (updates) =>
    set((state) => ({
      draftProposal: state.draftProposal
        ? { ...state.draftProposal, ...updates }
        : updates,
    })),

  clearDraftProposal: () => set({ draftProposal: null }),
}));