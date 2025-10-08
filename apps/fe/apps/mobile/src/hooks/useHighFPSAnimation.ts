import { useEffect } from 'react';
import { useSharedValue, withTiming, Easing } from 'react-native-reanimated';
import { getOptimalFrameRate } from '../lib/performance';

export const useHighFPSAnimation = (initialValue: number = 0) => {
  const animatedValue = useSharedValue(initialValue);
  const frameRate = getOptimalFrameRate();

  const animate = (toValue: number, duration: number = 300) => {
    animatedValue.value = withTiming(toValue, {
      duration,
      easing: Easing.bezier(0.25, 0.1, 0.25, 1),
    });
  };

  return { animatedValue, animate, frameRate };
};
