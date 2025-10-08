import { useEffect } from 'react';
import { Stack } from 'expo-router';
import { QueryClient, QueryClientProvider } from '@tanstack/react-query';
import { setTokenStorage } from '@skillsier/shared';
import { MMKV } from 'react-native-mmkv';
import '@/lib/i18n'; // Initialize i18n
import '@/lib/performance'; // Initialize performance optimizations
import '../global.css';

const storage = new MMKV();

setTokenStorage({
  getAccessToken: () => storage.getString('accessToken') || null,
  setAccessToken: (token: string) => storage.set('accessToken', token),
  getRefreshToken: () => storage.getString('refreshToken') || null,
  setRefreshToken: (token: string) => storage.set('refreshToken', token),
  clearTokens: () => {
    storage.delete('accessToken');
    storage.delete('refreshToken');
  },
});

const queryClient = new QueryClient({
  defaultOptions: {
    queries: {
      staleTime: 60 * 1000,
      retry: 3,
    },
  },
});

export default function RootLayout() {
  return (
    <QueryClientProvider client={queryClient}>
      <Stack screenOptions={{ headerShown: false }}>
        <Stack.Screen name="index" />
        <Stack.Screen name="(auth)" />
        <Stack.Screen name="(tabs)" />
      </Stack>
    </QueryClientProvider>
  );
}
