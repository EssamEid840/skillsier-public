import React from 'react';
import { View, Image, Text } from 'react-native';
import type { AvatarProps } from './Avatar.types';

const sizeValues = {
  sm: 32,
  md: 40,
  lg: 48,
  xl: 64,
};

export const Avatar: React.FC<AvatarProps> = ({
  src,
  alt = 'Avatar',
  size = 'md',
  fallback,
  className,
}) => {
  const [imgError, setImgError] = React.useState(false);
  const displayFallback = fallback || alt.charAt(0).toUpperCase();
  const sizeValue = sizeValues[size];

  return (
    <View
      className={`items-center justify-center overflow-hidden rounded-full bg-gray-200 ${
        className || ''
      }`}
      style={{ width: sizeValue, height: sizeValue }}
    >
      {src && !imgError ? (
        <Image
          source={{ uri: src }}
          style={{ width: sizeValue, height: sizeValue }}
          onError={() => setImgError(true)}
        />
      ) : (
        <Text className="font-medium text-gray-600">{displayFallback}</Text>
      )}
    </View>
  );
};
