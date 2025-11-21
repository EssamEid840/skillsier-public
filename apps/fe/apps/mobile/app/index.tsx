import { useAuth } from '@skillsier/auth';
import { Redirect } from 'expo-router';
import { ActivityIndicator, View } from 'react-native';

export default function Index() {
  const { isAuthenticated, isLoading } = useAuth();

  if (isLoading) {
    return (
      <View className="flex-1 items-center justify-center bg-white">
        <ActivityIndicator size="large" color="#E60023" />
      </View>
    );
  }

  if (isAuthenticated) {
    return <Redirect href="/(authenticated)/(tabs)" />;
  }

  return <Redirect href="/(auth)/login" />;
}