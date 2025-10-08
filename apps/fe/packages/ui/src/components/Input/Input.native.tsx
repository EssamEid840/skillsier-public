import React from 'react';
import { TextInput, Text, View } from 'react-native';
import type { InputProps } from './Input.types';

export const Input: React.FC<InputProps> = ({
  value,
  onChangeText,
  placeholder,
  label,
  error,
  disabled,
  type = 'text',
  className,
  leftIcon,
  rightIcon,
  required,
  autoComplete,
  autoFocus,
}) => {
  const getSecureTextEntry = () => type === 'password';
  const getKeyboardType = () => {
    switch (type) {
      case 'email':
        return 'email-address';
      case 'number':
        return 'numeric';
      default:
        return 'default';
    }
  };

  return (
    <View className="w-full">
      {label && (
        <Text className="mb-2 text-sm font-medium text-gray-700">
          {label}
          {required && <Text className="ml-1 text-red-500">*</Text>}
        </Text>
      )}
      <View className="relative">
        {leftIcon && (
          <View className="absolute inset-y-0 left-0 flex items-center pl-3 z-10">
            {leftIcon}
          </View>
        )}
        <TextInput
          value={value}
          onChangeText={onChangeText}
          placeholder={placeholder}
          editable={!disabled}
          secureTextEntry={getSecureTextEntry()}
          keyboardType={getKeyboardType()}
          autoComplete={autoComplete}
          autoFocus={autoFocus}
          placeholderTextColor="#9ca3af"
          className={`h-11 w-full rounded-lg border ${
            error ? 'border-red-500' : 'border-gray-300'
          } bg-white px-4 py-2 text-base ${leftIcon ? 'pl-10' : ''} ${
            rightIcon ? 'pr-10' : ''
          } ${disabled ? 'opacity-50' : ''} ${className || ''}`}
        />
        {rightIcon && (
          <View className="absolute inset-y-0 right-0 flex items-center pr-3">
            {rightIcon}
          </View>
        )}
      </View>
      {error && <Text className="mt-1 text-sm text-red-600">{error}</Text>}
    </View>
  );
};
