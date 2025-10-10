// apps/mobile/app/auth/callback.tsx
// OAuth callback handler for mobile

import { useEffect } from 'react';
import { View, Text, ActivityIndicator } from 'react-native';
import { useRouter } from 'expo-router';

export default function AuthCallbackScreen() {
  const router = useRouter();

  useEffect(() => {
    // The actual OAuth handling is done by expo-web-browser
    // This screen just shows a loading state briefly before redirecting
    const timer = setTimeout(() => {
      router.replace('/(tabs)/dashboard');
    }, 1000);

    return () => clearTimeout(timer);
  }, [router]);

  return (
    <View className="flex-1 items-center justify-center bg-white">
      <ActivityIndicator size="large" color="#3B82F6" />
      <Text className="mt-4 text-gray-600">Completing sign in...</Text>
    </View>
  );
}