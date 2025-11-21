import { Card } from '@skillsier/ui';
import type { Job } from '@skillsier/types';
import { useRouter } from 'expo-router';
import { Pressable, Text, View } from 'react-native';

interface JobCardProps {
  job: Job;
}

export default function JobCard({ job }: JobCardProps) {
  const router = useRouter();

  return (
    <Pressable onPress={() => router.push(`/(authenticated)/(tabs)/jobs/${job.id}`)}>
      <Card className="p-4 active:opacity-70">
        <View className="flex-row items-start justify-between mb-2">
          <Text className="flex-1 text-lg font-semibold text-gray-900 mr-2">
            {job.title}
          </Text>
          <View className="bg-primary-100 px-2 py-1 rounded-full">
            <Text className="text-xs font-medium text-primary-700">
              {job.status}
            </Text>
          </View>
        </View>

        <Text className="text-gray-600 mb-3" numberOfLines={3}>
          {job.description}
        </Text>

        {/* Skills */}
        {job.skills && job.skills.length > 0 && (
          <View className="flex-row flex-wrap gap-2 mb-3">
            {job.skills.slice(0, 3).map((skill, index) => (
              <View key={index} className="bg-gray-100 px-2 py-1 rounded">
                <Text className="text-xs text-gray-700">{skill}</Text>
              </View>
            ))}
            {job.skills.length > 3 && (
              <View className="bg-gray-100 px-2 py-1 rounded">
                <Text className="text-xs text-gray-700">
                  +{job.skills.length - 3} more
                </Text>
              </View>
            )}
          </View>
        )}

        {/* Bottom Row */}
        <View className="flex-row items-center justify-between">
          <View className="flex-row items-center gap-3">
            {job.budget && (
              <Text className="text-base font-bold text-primary-500">
                ${job.budget.toLocaleString()}
              </Text>
            )}
            {job.duration && (
              <Text className="text-xs text-gray-500">{job.duration}</Text>
            )}
          </View>
          {job.proposalCount !== undefined && (
            <Text className="text-xs text-gray-500">
              {job.proposalCount} proposals
            </Text>
          )}
        </View>

        {/* Posted Date */}
        {job.createdAt && (
          <Text className="text-xs text-gray-400 mt-2">
            Posted {new Date(job.createdAt).toLocaleDateString()}
          </Text>
        )}
      </Card>
    </Pressable>
  );
}