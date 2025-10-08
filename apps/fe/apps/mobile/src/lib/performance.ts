import { Platform } from 'react-native';
import { enableFreeze, enableScreens } from 'react-native-screens';
import { configureReanimatedLogger, ReanimatedLogLevel } from 'react-native-reanimated';

// Enable React Native Screens for better performance
enableScreens(true);
enableFreeze(true);

// Configure Reanimated for 120 FPS
configureReanimatedLogger({
  level: ReanimatedLogLevel.warn,
  strict: false,
});

// Frame rate configuration
export const FRAME_RATE = {
  DEFAULT: 60,
  HIGH: 90,
  ULTRA: 120,
};

// Detect device capability
export const getOptimalFrameRate = () => {
  if (Platform.OS === 'ios') {
    // iPhone 13 Pro and later support 120Hz
    return FRAME_RATE.ULTRA;
  } else if (Platform.OS === 'android') {
    // Many Android flagships support 90Hz or 120Hz
    return FRAME_RATE.HIGH;
  }
  return FRAME_RATE.DEFAULT;
};

// Performance utilities
export const performanceConfig = {
  shouldUseNativeDriver: true,
  frameRate: getOptimalFrameRate(),
  removeClippedSubviews: true,
  maxToRenderPerBatch: 10,
  updateCellsBatchingPeriod: 50,
  windowSize: 21,
};
