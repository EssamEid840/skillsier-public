import { Button, Card } from '@skillsier/ui';
import { useJob } from '@skillsier/hooks';
import { useRouter } from 'expo-router';
import {
  ActivityIndicator,
  Pressable,
  ScrollView,
  Text,
  View,
} from 'react-native';

interface JobDetailsProps {
  jobId: string;
}

export default function JobDetails({ jobId }: JobDetailsProps) {
  const router = useRouter();
  const { data: job, isLoading } = useJob(jobId);

  if (isLoading) {
    return (
      <View className="flex-1 items-center justify-center bg-gray-50">
        <ActivityIndicator size="large" color="#E60023" />
      </View>
    );
  }

  if (!job) {
    return (
      <View className="flex-1 items-center justify-center bg-gray-50 p-6">
        <Text className="text-gray-500 text-center">Job not found</Text>
        <Button onPress={() => router.back()} className="mt-4">
          <Text className="text-white font-semibold">Go Back</Text>
        </Button>
      </View>
    );
  }

  return (
    <ScrollView className="flex-1 bg-gray-50">
      <View className="p-4 space-y-4">
        {/* Header Card */}
        <Card className="p-4">
          <View className="flex-row items-start justify-between mb-3">
            <View className="flex-1 mr-2">
              <Text className="text-2xl font-bold text-gray-900 mb-2">
                {job.title}
              </Text>
              <View className="flex-row items-center gap-2">
                <View className="bg-primary-100 px-3 py-1 rounded-full">
                  <Text className="text-xs font-medium text-primary-700">
                    {job.status}
                  </Text>
                </View>
                {job.createdAt && (
                  <Text className="text-xs text-gray-500">
                    {new Date(job.createdAt).toLocaleDateString()}
                  </Text>
                )}
              </View>
            </View>
          </View>

          <Text className="text-gray-700 leading-6">{job.description}</Text>
        </Card>

        {/* Quick Stats */}
        <View className="flex-row gap-3">
          <Card className="flex-1 p-4">
            <Text className="text-xs font-medium text-gray-500 mb-1">
              Budget
            </Text>
            <Text className="text-xl font-bold text-primary-500">
              ${job.budget?.toLocaleString() || 'N/A'}
            </Text>
          </Card>

          <Card className="flex-1 p-4">
            <Text className="text-xs font-medium text-gray-500 mb-1">
              Proposals
            </Text>
            <Text className="text-xl font-bold text-gray-900">
              {job.proposalCount || 0}
            </Text>
          </Card>
        </View>

        {/* Duration */}
        {job.duration && (
          <Card className="p-4">
            <Text className="text-xs font-medium text-gray-500 mb-2">
              Project Duration
            </Text>
            <Text className="text-base text-gray-900">{job.duration}</Text>
          </Card>
        )}

        {/* Skills */}
        {job.skills && job.skills.length > 0 && (
          <Card className="p-4">
            <Text className="text-sm font-semibold text-gray-900 mb-3">
              Required Skills
            </Text>
            <View className="flex-row flex-wrap gap-2">
              {job.skills.map((skill, index) => (
                <View key={index} className="bg-gray-100 px-3 py-2 rounded-full">
                  <Text className="text-sm text-gray-700">{skill}</Text>
                </View>
              ))}
            </View>
          </Card>
        )}

        {/* Actions */}
        <Card className="p-4 space-y-3">
          <Pressable
            onPress={() =>
              router.push(`/(authenticated)/(tabs)/jobs/${job.id}/proposals`)
            }
          >
            <View className="flex-row items-center justify-between py-3 border-b border-gray-100">
              <View className="flex-row items-center gap-3">
                <View className="w-10 h-10 bg-primary-100 rounded-full items-center justify-center">
                  <Text className="text-primary-500 text-lg">📄</Text>
                </View>
                <Text className="text-base font-medium text-gray-900">
                  View Proposals
                </Text>
              </View>
              <Text className="text-gray-400">›</Text>
            </View>
          </Pressable>

          <Pressable
            onPress={() =>
              router.push(`/(authenticated)/(tabs)/jobs/${job.id}/applicants`)
            }
          >
            <View className="flex-row items-center justify-between py-3 border-b border-gray-100">
              <View className="flex-row items-center gap-3">
                <View className="w-10 h-10 bg-primary-100 rounded-full items-center justify-center">
                  <Text className="text-primary-500 text-lg">👥</Text>
                </View>
                <Text className="text-base font-medium text-gray-900">
                  View Applicants
                </Text>
              </View>
              <Text className="text-gray-400">›</Text>
            </View>
          </Pressable>

          <Pressable
            onPress={() =>
              router.push(`/(authenticated)/(tabs)/jobs/${job.id}/analytics`)
            }
          >
            <View className="flex-row items-center justify-between py-3">
              <View className="flex-row items-center gap-3">
                <View className="w-10 h-10 bg-primary-100 rounded-full items-center justify-center">
                  <Text className="text-primary-500 text-lg">📊</Text>
                </View>
                <Text className="text-base font-medium text-gray-900">
                  View Analytics
                </Text>
              </View>
              <Text className="text-gray-400">›</Text>
            </View>
          </Pressable>
        </Card>

        {/* Edit/Delete Actions */}
        <View className="space-y-3 pb-6">
          <Button
            onPress={() =>
              router.push(`/(authenticated)/(tabs)/jobs/${job.id}/edit`)
            }
          >
            <Text className="text-white font-semibold">Edit Job</Text>
          </Button>
          <Button variant="outline" onPress={() => router.back()}>
            <Text className="text-gray-700 font-semibold">Back to Jobs</Text>
          </Button>
        </View>
      </View>
    </ScrollView>
  );
}