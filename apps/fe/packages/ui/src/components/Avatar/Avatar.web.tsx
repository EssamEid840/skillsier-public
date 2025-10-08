import React from 'react';
import { cn } from '../../lib/utils';
import type { AvatarProps } from './Avatar.types';

const sizeVariants = {
  sm: 'h-8 w-8 text-xs',
  md: 'h-10 w-10 text-sm',
  lg: 'h-12 w-12 text-base',
  xl: 'h-16 w-16 text-lg',
};

export const Avatar: React.FC<AvatarProps> = ({
  src,
  alt = 'Avatar',
  size = 'md',
  fallback,
  className,
}) => {
  const [imgError, setImgError] = React.useState(false);

  const displayFallback = fallback || alt.charAt(0).toUpperCase();

  return (
    <div
      className={cn(
        'relative inline-flex items-center justify-center overflow-hidden rounded-full bg-gray-200',
        sizeVariants[size],
        className
      )}
    >
      {src && !imgError ? (
        <img
          src={src}
          alt={alt}
          className="h-full w-full object-cover"
          onError={() => setImgError(true)}
        />
      ) : (
        <span className="font-medium text-gray-600">{displayFallback}</span>
      )}
    </div>
  );
};