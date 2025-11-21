import { create } from 'zustand';
import type { Contract } from '@skillsier/types';

interface ContractsState {
  // State
  selectedContract: Contract | null;
  activeContractId: string | null;

  // Actions
  setSelectedContract: (contract: Contract | null) => void;
  setActiveContractId: (contractId: string | null) => void;
}

export const useContractsStore = create<ContractsState>((set) => ({
  // Initial state
  selectedContract: null,
  activeContractId: null,

  // Actions
  setSelectedContract: (contract) => set({ selectedContract: contract }),
  setActiveContractId: (contractId) => set({ activeContractId: contractId }),
}));