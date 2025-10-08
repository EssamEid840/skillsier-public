// packages/ui/src/components/Badge/Badge.tsx
import { Platform } from 'react-native';

export { Badge } from Platform.select({
  web: () => require('./Badge.web'),
  default: () => require('./Badge.native'),
})()!;

export type { BadgeProps } from './Badge.types';