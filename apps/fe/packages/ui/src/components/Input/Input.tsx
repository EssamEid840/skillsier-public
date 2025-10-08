export { Input } from './Input';
export type { InputProps } from './Input.types';

// ============================================
// FILE: packages/ui/src/components/Card/Card.types.ts
// ============================================
import type { ReactNode } from 'react';

export interface CardProps {
  children: ReactNode;
  className?: string;
  padding?: 'none' | 'sm' | 'md' | 'lg';
  shadow?: 'none' | 'sm' | 'md' | 'lg';
  onPress?: () => void;
  hoverable?: boolean;
}