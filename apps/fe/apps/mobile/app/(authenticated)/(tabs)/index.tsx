import { useAuth } from '@skillsier/auth';
import { Button, Card } from '@skillsier/ui';
import { useRouter } from 'expo-router';
import { ScrollView, Text, View } from 'react-native';

export default function DashboardScreen() {
  const router = useRouter();
  const { user, logout } = useAuth();

  const handleLogout = async () => {
    await logout();
    router.replace('/(auth)/login');
  };

  return (
    <ScrollView className="flex-1 bg-gray-50">
      <View className="p-4 space-y-4">
        <Card className="p-6">
          <Text className="text-2xl font-bold text-gray-900 mb-2">
            Welcome back!
          </Text>
          <Text className="text-gray-600">
            {user?.email || 'Guest User'}
          </Text>
        </Card>

        <Card className="p-6">
          <Text className="text-lg font-semibold text-gray-900 mb-4">
            Quick Stats
          </Text>
          <View className="space-y-3">
            <View className="flex-row justify-between">
              <Text className="text-gray-600">Active Jobs</Text>
              <Text className="font-semibold text-gray-900">0</Text>
            </View>
            <View className="flex-row justify-between">
              <Text className="text-gray-600">Proposals</Text>
              <Text className="font-semibold text-gray-900">0</Text>
            </View>
            <View className="flex-row justify-between">
              <Text className="text-gray-600">Contracts</Text>
              <Text className="font-semibold text-gray-900">0</Text>
            </View>
          </View>
        </Card>

        <Button onPress={handleLogout} variant="outline">
          <Text className="text-primary-500 font-semibold">Sign Out</Text>
        </Button>
      </View>
    </ScrollView>
  );
}