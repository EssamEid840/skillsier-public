import React from 'react';
import { FlashList, FlashListProps } from '@shopify/flash-list';
import { performanceConfig } from '../lib/performance';

export function OptimizedFlashList<T>(props: FlashListProps<T>) {
  return (
    <FlashList
      {...props}
      removeClippedSubviews={performanceConfig.removeClippedSubviews}
      maxToRenderPerBatch={performanceConfig.maxToRenderPerBatch}
      updateCellsBatchingPeriod={performanceConfig.updateCellsBatchingPeriod}
      windowSize={performanceConfig.windowSize}
      // Optimize for 120 FPS
      drawDistance={200}
      estimatedItemSize={props.estimatedItemSize || 100}
      overrideItemLayout={(layout, item, index) => {
        // Pre-calculate layouts for smoother scrolling at high FPS
        if (props.estimatedItemSize) {
          layout.size = props.estimatedItemSize;
        }
      }}
    />
  );
}