import * as React from 'react';
import { View, Text, StyleSheet, type ViewProps } from 'react-native';
import { colors } from '../../tokens';

export interface CardProps extends ViewProps {}

export const Card = React.forwardRef<View, CardProps>(
  ({ style, ...props }, ref) => (
    <View ref={ref} style={[styles.card, style]} {...props} />
  )
);
Card.displayName = 'Card';

export const CardHeader = React.forwardRef<View, ViewProps>(
  ({ style, ...props }, ref) => (
    <View ref={ref} style={[styles.header, style]} {...props} />
  )
);
CardHeader.displayName = 'CardHeader';

export interface CardTitleProps {
  children: React.ReactNode;
  style?: any;
}

export const CardTitle: React.FC<CardTitleProps> = ({ children, style }) => (
  <Text style={[styles.title, style]}>{children}</Text>
);

export interface CardDescriptionProps {
  children: React.ReactNode;
  style?: any;
}

export const CardDescription: React.FC<CardDescriptionProps> = ({
  children,
  style,
}) => <Text style={[styles.description, style]}>{children}</Text>;

export const CardContent = React.forwardRef<View, ViewProps>(
  ({ style, ...props }, ref) => (
    <View ref={ref} style={[styles.content, style]} {...props} />
  )
);
CardContent.displayName = 'CardContent';

export const CardFooter = React.forwardRef<View, ViewProps>(
  ({ style, ...props }, ref) => (
    <View ref={ref} style={[styles.footer, style]} {...props} />
  )
);
CardFooter.displayName = 'CardFooter';

const styles = StyleSheet.create({
  card: {
    borderRadius: 12,
    borderWidth: 1,
    borderColor: colors.secondary[200],
    backgroundColor: '#FFFFFF',
    shadowColor: '#000',
    shadowOffset: { width: 0, height: 1 },
    shadowOpacity: 0.05,
    shadowRadius: 2,
    elevation: 1,
  },
  header: {
    padding: 24,
    gap: 6,
  },
  title: {
    fontSize: 18,
    fontWeight: '600',
    color: colors.secondary[900],
  },
  description: {
    fontSize: 14,
    color: colors.secondary[500],
  },
  content: {
    padding: 24,
    paddingTop: 0,
  },
  footer: {
    flexDirection: 'row',
    alignItems: 'center',
    padding: 24,
    paddingTop: 0,
  },
});