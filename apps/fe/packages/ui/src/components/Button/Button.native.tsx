import * as React from 'react';
import {
  TouchableOpacity,
  Text,
  StyleSheet,
  type TouchableOpacityProps,
} from 'react-native';
import { colors } from '../../tokens';

export interface ButtonProps extends TouchableOpacityProps {
  variant?: 'primary' | 'secondary' | 'outline' | 'ghost' | 'danger';
  size?: 'sm' | 'md' | 'lg';
  children: React.ReactNode;
}

export const Button = React.forwardRef<TouchableOpacity, ButtonProps>(
  (
    { variant = 'primary', size = 'md', children, style, disabled, ...props },
    ref
  ) => {
    const buttonStyles = [
      styles.base,
      styles[variant],
      styles[size],
      disabled && styles.disabled,
      style,
    ];

    const textStyles = [
      styles.text,
      styles[`${variant}Text` as keyof typeof styles],
      styles[`${size}Text` as keyof typeof styles],
    ];

    return (
      <TouchableOpacity
        ref={ref}
        style={buttonStyles}
        disabled={disabled}
        activeOpacity={0.7}
        {...props}
      >
        <Text style={textStyles}>{children}</Text>
      </TouchableOpacity>
    );
  }
);

Button.displayName = 'Button';

const styles = StyleSheet.create({
  base: {
    flexDirection: 'row',
    alignItems: 'center',
    justifyContent: 'center',
    borderRadius: 8,
    gap: 8,
  },
  primary: {
    backgroundColor: colors.primary.DEFAULT,
  },
  secondary: {
    backgroundColor: colors.secondary[100],
  },
  outline: {
    backgroundColor: 'transparent',
    borderWidth: 2,
    borderColor: colors.primary.DEFAULT,
  },
  ghost: {
    backgroundColor: 'transparent',
  },
  danger: {
    backgroundColor: colors.error.DEFAULT,
  },
  disabled: {
    opacity: 0.5,
  },
  sm: {
    height: 36,
    paddingHorizontal: 12,
  },
  md: {
    height: 40,
    paddingHorizontal: 16,
    paddingVertical: 8,
  },
  lg: {
    height: 44,
    paddingHorizontal: 32,
  },
  text: {
    fontWeight: '500',
  },
  primaryText: {
    color: '#FFFFFF',
    fontSize: 14,
  },
  secondaryText: {
    color: colors.secondary[900],
    fontSize: 14,
  },
  outlineText: {
    color: colors.primary.DEFAULT,
    fontSize: 14,
  },
  ghostText: {
    color: colors.secondary[900],
    fontSize: 14,
  },
  dangerText: {
    color: '#FFFFFF',
    fontSize: 14,
  },
  smText: {
    fontSize: 12,
  },
  mdText: {
    fontSize: 14,
  },
  lgText: {
    fontSize: 14,
  },
});