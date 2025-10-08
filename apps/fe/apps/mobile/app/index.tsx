import { useEffect } from 'react';
import { router } from 'expo-router';
import { View, Text, Image, ActivityIndicator } from 'react-native';
import { useAuth } from '@skillsier/shared';
import { Button } from '@skillsier/ui';
import { SafeAreaView } from 'react-native-safe-area-context';

export default function LandingScreen() {
  const { isAuthenticated, isLoading } = useAuth();

  useEffect(() => {
    if (!isLoading && isAuthenticated) {
      router.replace('/(tabs)/dashboard');
    }
  }, [isAuthenticated, isLoading]);

  if (isLoading) {
    return (
      <View className="flex-1 items-center justify-center bg-white">
        <ActivityIndicator size="large" color="#6366f1" />
      </View>
    );
  }

  return (
    <SafeAreaView className="flex-1 bg-gradient-to-br from-primary-50 to-purple-50">
      <View className="flex-1 px-6 justify-center">
        <View className="items-center mb-12">
          <View className="h-20 w-20 rounded-2xl bg-gradient-to-br from-primary-600 to-purple-600 mb-6" />
          <Text className="text-4xl font-bold text-gray-900 text-center">
            Welcome to Skillsier
          </Text>
          <Text className="text-lg text-gray-600 text-center mt-4 px-4">
            Transform your career with enterprise-grade learning
          </Text>
        </View>

        <View className="space-y-4">
          <Button
            onPress={() => router.push('/(auth)/register')}
            size="lg"
            fullWidth
          >
            Get Started Free
          </Button>
          <Button
            onPress={() => router.push('/(auth)/login')}
            variant="outline"
            size="lg"
            fullWidth
          >
            Sign In
          </Button>
        </View>

        <View className="mt-12">
          <View className="flex-row justify-around px-8">
            <View className="items-center">
              <Text className="text-3xl font-bold text-primary-600">500K+</Text>
              <Text className="text-sm text-gray-600 mt-1">Learners</Text>
            </View>
            <View className="items-center">
              <Text className="text-3xl font-bold text-primary-600">10K+</Text>
              <Text className="text-sm text-gray-600 mt-1">Courses</Text>
            </View>
            <View className="items-center">
              <Text className="text-3xl font-bold text-primary-600">98%</Text>
              <Text className="text-sm text-gray-600 mt-1">Success</Text>
            </View>
          </View>
        </View>
      </View>
    </SafeAreaView>
  );
}
