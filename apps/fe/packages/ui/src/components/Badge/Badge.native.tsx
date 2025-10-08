import React from 'react';
import { View, Text } from 'react-native';
import type { BadgeProps } from './Badge.types';

const variantClasses = {
  default: 'bg-gray-100',
  success: 'bg-green-100',
  warning: 'bg-yellow-100',
  error: 'bg-red-100',
  info: 'bg-blue-100',
};

const textVariantClasses = {
  default: 'text-gray-800',
  success: 'text-green-800',
  warning: 'text-yellow-800',
  error: 'text-red-800',
  info: 'text-blue-800',
};

export const Badge: React.FC<BadgeProps> = ({
  children,
  variant = 'default',
  size = 'sm',
  className,
}) => {
  return (
    <View
      className={`inline-flex items-center rounded-full ${
        size === 'sm' ? 'px-2 py-0.5' : 'px-2.5 py-1'
      } ${variantClasses[variant]} ${className || ''}`}
    >
      <Text
        className={`${size === 'sm' ? 'text-xs' : 'text-sm'} font-medium ${
          textVariantClasses[variant]
        }`}
      >
        {children}
      </Text>
    </View>
  );
};