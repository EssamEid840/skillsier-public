// packages/ui/src/components/Avatar/Avatar.tsx
import { Platform } from 'react-native';

export { Avatar } from Platform.select({
  web: () => require('./Avatar.web'),
  default: () => require('./Avatar.native'),
})()!;

export type { AvatarProps } from './Avatar.types';