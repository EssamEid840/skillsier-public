import React from 'react';
import { cn } from '../../lib/utils';
import type { CardProps } from './Card.types';

const paddingVariants = {
  none: '',
  sm: 'p-3',
  md: 'p-4',
  lg: 'p-6',
};

const shadowVariants = {
  none: '',
  sm: 'shadow-sm',
  md: 'shadow-md',
  lg: 'shadow-lg',
};

export const Card: React.FC<CardProps> = ({
  children,
  className,
  padding = 'md',
  shadow = 'sm',
  onPress,
  hoverable = false,
}) => {
  const Component = onPress ? 'button' : 'div';

  return (
    <Component
      onClick={onPress}
      className={cn(
        'rounded-lg border border-gray-200 bg-white',
        paddingVariants[padding],
        shadowVariants[shadow],
        hoverable && 'transition-shadow hover:shadow-md',
        onPress && 'cursor-pointer transition-transform hover:scale-[1.02]',
        className
      )}
    >
      {children}
    </Component>
  );
};
