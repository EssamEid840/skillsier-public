export interface PressableProps {
  children: React.ReactNode;
  onPress?: () => void;
  className?: string;
  disabled?: boolean;
  style?: Record<string, unknown>;
}
