export * from './colors';
export * from './spacing';
export * from './typography';

export const theme = {
  borderRadius: {
    sm: 4,
    md: 8,
    lg: 12,
    full: 9999,
  },
  shadows: {
    sm: '0 1px 2px 0 rgba(0, 0, 0, 0.05)',
    md: '0 4px 6px -1px rgba(0, 0, 0, 0.1)',
    lg: '0 10px 15px -3px rgba(0, 0, 0, 0.1)',
    xl: '0 20px 25px -5px rgba(0, 0, 0, 0.1)',
  },
} as const;