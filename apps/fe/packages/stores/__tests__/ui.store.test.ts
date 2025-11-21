import { describe, it, expect, beforeEach } from 'vitest';
import { useUIStore } from '../src/ui.store';

describe('UI Store', () => {
  beforeEach(() => {
    const store = useUIStore.getState();
    store.setTheme('system');
    store.setSidebarOpen(true);
    store.setSidebarCollapsed(false);
    store.setMobileMenuOpen(false);
    store.closeModal();
  });

  it('initializes with default values', () => {
    const state = useUIStore.getState();
    
    expect(state.theme).toBe('system');
    expect(state.sidebarOpen).toBe(true);
    expect(state.sidebarCollapsed).toBe(false);
    expect(state.mobileMenuOpen).toBe(false);
    expect(state.activeModal).toBeNull();
  });

  it('toggles sidebar', () => {
    const store = useUIStore.getState();
    
    store.toggleSidebar();
    expect(useUIStore.getState().sidebarOpen).toBe(false);
    
    store.toggleSidebar();
    expect(useUIStore.getState().sidebarOpen).toBe(true);
  });

  it('manages modal state', () => {
    const store = useUIStore.getState();
    const modalData = { jobId: '123' };
    
    store.openModal('createJob', modalData);
    let state = useUIStore.getState();
    expect(state.activeModal).toBe('createJob');
    expect(state.modalData).toEqual(modalData);
    
    store.closeModal();
    state = useUIStore.getState();
    expect(state.activeModal).toBeNull();
    expect(state.modalData).toBeNull();
  });

  it('manages toasts', () => {
    const store = useUIStore.getState();
    
    store.addToast({
      type: 'success',
      message: 'Job created successfully',
    });
    
    let state = useUIStore.getState();
    expect(state.toasts.length).toBe(1);
    expect(state.toasts[0]?.message).toBe('Job created successfully');
    
    const toastId = state.toasts[0]?.id || '';
    store.removeToast(toastId);
    
    state = useUIStore.getState();
    expect(state.toasts.length).toBe(0);
  });

  it('changes theme', () => {
    const store = useUIStore.getState();
    
    store.setTheme('dark');
    expect(useUIStore.getState().theme).toBe('dark');
    
    store.setTheme('light');
    expect(useUIStore.getState().theme).toBe('light');
  });
});