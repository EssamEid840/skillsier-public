import React, { forwardRef } from 'react';
import { cn } from '../../lib/utils';
import type { InputProps } from './Input.types';

export const Input = forwardRef<HTMLInputElement, InputProps>(
  (
    {
      value,
      onChange,
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
    },
    ref
  ) => {
    return (
      <div className="w-full">
        {label && (
          <label className="mb-2 block text-sm font-medium text-gray-700">
            {label}
            {required && <span className="ml-1 text-red-500">*</span>}
          </label>
        )}
        <div className="relative">
          {leftIcon && (
            <div className="pointer-events-none absolute inset-y-0 left-0 flex items-center pl-3">
              {leftIcon}
            </div>
          )}
          <input
            ref={ref}
            type={type}
            value={value}
            onChange={onChange}
            placeholder={placeholder}
            disabled={disabled}
            required={required}
            autoComplete={autoComplete}
            autoFocus={autoFocus}
            className={cn(
              'flex h-11 w-full rounded-lg border border-gray-300 bg-white px-4 py-2 text-base transition-colors',
              'placeholder:text-gray-400',
              'focus:border-primary-500 focus:outline-none focus:ring-2 focus:ring-primary-500 focus:ring-offset-0',
              'disabled:cursor-not-allowed disabled:opacity-50',
              leftIcon && 'pl-10',
              rightIcon && 'pr-10',
              error && 'border-red-500 focus:border-red-500 focus:ring-red-500',
              className
            )}
          />
          {rightIcon && (
            <div className="pointer-events-none absolute inset-y-0 right-0 flex items-center pr-3">
              {rightIcon}
            </div>
          )}
        </div>
        {error && <p className="mt-1 text-sm text-red-600">{error}</p>}
      </div>
    );
  }
);

Input.displayName = 'Input';