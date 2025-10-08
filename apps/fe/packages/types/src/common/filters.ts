// ============================================
// FILE: packages/types/src/common/filters.ts
// ============================================
export interface CourseFilters {
  category?: string;
  level?: CourseLevel;
  minRating?: number;
  maxPrice?: number;
  tags?: string[];
  search?: string;
}

export interface UserFilters {
  role?: UserRole;
  search?: string;
  emailVerified?: boolean;
}