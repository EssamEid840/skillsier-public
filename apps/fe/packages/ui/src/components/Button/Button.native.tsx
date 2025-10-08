import React from 'react';
import { TouchableOpacity, Text, ActivityIndicator, View } from 'react-native';
import type { ButtonProps } from './Button.types';

export const Button: React.FC<ButtonProps> = ({
  children,
  variant = 'primary',
  size = 'md',
  disabled,
  loading,
  onPress,
  className,
  fullWidth,
  icon,
  iconPosition = 'left',
}) => {
  const getVariantStyles = () => {
    const base = 'rounded-lg flex-row items-center justify-center';
    const variants = {
      primary: 'bg-primary',
      secondary: 'bg-gray-100',
      outline: 'border-2 border-gray-300 bg-transparent',
      ghost: 'bg-transparent',
      destructive: 'bg-red-600',
    };
    return `${base} ${variants[variant]}`;
  };

  const getSizeStyles = () => {
    const sizes = {
      sm: 'h-9 px-3',
      md: 'h-11 px-5',
      lg: 'h-12 px-6',
    };
    return sizes[size];
  };

  const getTextStyles = () => {
    const textVariants = {
      primary: 'text-white font-medium',
      secondary: 'text-gray-900 font-medium',
      outline: 'text-gray-900 font-medium',
      ghost: 'text-gray-900 font-medium',
      destructive: 'text-white font-medium',
    };
    return textVariants[variant];
  };

  return (
    <TouchableOpacity
      onPress={onPress}
      disabled={disabled || loading}
      className={`${getVariantStyles()} ${getSizeStyles()} ${fullWidth ? 'w-full' : ''} ${
        disabled || loading ? 'opacity-50' : ''
      } ${className || ''}`}
      activeOpacity={0.7}
    >
      {loading ? (
        <View className="flex-row items-center">
          <ActivityIndicator color={variant === 'primary' ? 'white' : 'black'} className="mr-2" />
          <Text className={getTextStyles()}>Loading...</Text>
        </View>
      ) : (
        <View className="flex-row items-center">
          {icon && iconPosition === 'left' && <View className="mr-2">{icon}</View>}
          <Text className={getTextStyles()}>{children}</Text>
          {icon && iconPosition === 'right' && <View className="ml-2">{icon}</View>}
        </View>
      )}
    </TouchableOpacity>
  );
};