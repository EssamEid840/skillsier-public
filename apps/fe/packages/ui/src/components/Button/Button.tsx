export { Button } from './Button';
export type { ButtonProps } from './Button.types';

// ============================================
// FILE: packages/ui/src/components/Input/Input.types.ts
// ============================================
import type { ReactNode } from 'react';

export interface InputProps {
  value: string;
  onChangeText?: (text: string) => void;
  onChange?: (e: React.ChangeEvent<HTMLInputElement>) => void;
  placeholder?: string;
  label?: string;
  error?: string;
  disabled?: boolean;
  type?: 'text' | 'email' | 'password' | 'number';
  className?: string;
  leftIcon?: ReactNode;
  rightIcon?: ReactNode;
  required?: boolean;
  autoComplete?: string;
  autoFocus?: boolean;
}