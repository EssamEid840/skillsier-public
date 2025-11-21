import { useJob } from '@skillsier/hooks';
import { Button, Card } from '@skillsier/ui';
import { useLocalSearchParams, useRouter } from 'expo-router';
import { ActivityIndicator, ScrollView, Text, View } from 'react-native';

export default function JobDetailScreen() {
  const router = useRouter();
  const { id } = useLocalSearchParams<{ id: string }>();
  const { data: job, isLoading } = useJob(id);

  if (isLoading) {
    return (
      <View className="flex-1 items-center justify-center bg-white">
        <ActivityIndicator size="large" color="#E60023" />
      </View>
    );
  }

  if (!job) {
    return (
      <View className="flex-1 items-center justify-center bg-white">
        <Text className="text-gray-500">Job not found</Text>
      </View>
    );
  }

  return (
    <ScrollView className="flex-1 bg-gray-50">
      <View className="p-4 space-y-4">
        <Card className="p-6">
          <Text className="text-2xl font-bold text-gray-900 mb-2">
            {job.title}
          </Text>
          <Text className="text-gray-600 mb-4">{job.description}</Text>

          <View className="space-y-2">
            <View className="flex-row justify-between">
              <Text className="text-gray-600">Budget</Text>
              <Text className="font-semibold text-gray-900">
                ${job.budget?.toLocaleString() || 'N/A'}
              </Text>
            </View>
            <View className="flex-row justify-between">
              <Text className="text-gray-600">Status</Text>
              <Text className="font-semibold text-primary-500">
                {job.status}
              </Text>
            </View>
          </View>
        </Card>

        <Button onPress={() => router.back()}>
          <Text className="text-white font-semibold">Back to Jobs</Text>
        </Button>
      </View>
    </ScrollView>
  );
}