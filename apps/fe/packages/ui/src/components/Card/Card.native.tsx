import React from 'react';
import { View, TouchableOpacity } from 'react-native';
import type { CardProps } from './Card.types';

const paddingClasses = {
  none: '',
  sm: 'p-3',
  md: 'p-4',
  lg: 'p-6',
};

export const Card: React.FC<CardProps> = ({
  children,
  className,
  padding = 'md',
  onPress,
}) => {
  const Component = onPress ? TouchableOpacity : View;

  return (
    <Component
      onPress={onPress}
      activeOpacity={onPress ? 0.7 : 1}
      className={`rounded-lg border border-gray-200 bg-white ${paddingClasses[padding]} ${
        className || ''
      }`}
      style={{
        shadowColor: '#000',
        shadowOffset: { width: 0, height: 2 },
        shadowOpacity: 0.1,
        shadowRadius: 4,
        elevation: 2,
      }}
    >
      {children}
    </Component>
  );
};
