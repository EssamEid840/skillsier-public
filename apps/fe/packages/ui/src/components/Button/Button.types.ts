import type { ReactNode } from 'react';

export interface ButtonProps {
  children: ReactNode;
  variant?: 'primary' | 'secondary' | 'outline' | 'ghost' | 'destructive';
  size?: 'sm' | 'md' | 'lg';
  disabled?: boolean;
  loading?: boolean;
  onPress?: () => void;
  onClick?: () => void;
  className?: string;
  type?: 'button' | 'submit' | 'reset';
  fullWidth?: boolean;
  icon?: ReactNode;
  iconPosition?: 'left' | 'right';
}