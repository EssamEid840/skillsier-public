export { Card } from './Card';
export type { CardProps } from './Card.types';

// ============================================
// FILE: packages/ui/src/components/Avatar/Avatar.types.ts
// ============================================
export interface AvatarProps {
  src?: string;
  alt?: string;
  size?: 'sm' | 'md' | 'lg' | 'xl';
  fallback?: string;
  className?: string;
}
