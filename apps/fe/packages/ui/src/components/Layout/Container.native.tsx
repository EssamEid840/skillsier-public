import * as React from 'react';
import { View, StyleSheet, type ViewProps } from 'react-native';

export interface ContainerProps extends ViewProps {
  maxWidth?: 'sm' | 'md' | 'lg' | 'xl' | '2xl' | 'full';
}

export const Container = React.forwardRef<View, ContainerProps>(
  ({ style, maxWidth = 'xl', ...props }, ref) => {
    return (
      <View
        ref={ref}
        style={[styles.container, maxWidthStyles[maxWidth], style]}
        {...props}
      />
    );
  }
);
Container.displayName = 'Container';

const styles = StyleSheet.create({
  container: {
    width: '100%',
    marginHorizontal: 'auto',
    paddingHorizontal: 16,
  },
});

const maxWidthStyles = StyleSheet.create({
  sm: { maxWidth: 640 },
  md: { maxWidth: 768 },
  lg: { maxWidth: 1024 },
  xl: { maxWidth: 1280 },
  '2xl': { maxWidth: 1536 },
  full: { maxWidth: '100%' },
});