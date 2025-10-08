// ============================================
// FILE: packages/types/src/common/pagination.ts
// ============================================
export interface PaginationParams {
  page?: number;
  pageSize?: number;
  sortBy?: string;
  sortOrder?: 'asc' | 'desc';
}

export interface CursorPaginationParams {
  cursor?: string;
  limit?: number;
}