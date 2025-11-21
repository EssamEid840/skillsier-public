import JobCard from '@/components/jobs/JobCard.native';
import JobFilters from '@/components/jobs/JobFilters.native';
import { useJobs } from '@skillsier/hooks';
import { useJobsStore } from '@skillsier/stores';
import { useState } from 'react';
import {
  ActivityIndicator,
  FlatList,
  RefreshControl,
  Text,
  View,
} from 'react-native';

export default function JobsScreen() {
  const { filters } = useJobsStore();
  const { data: jobs, isLoading, refetch, isRefreshing } = useJobs(filters);
  const [showFilters, setShowFilters] = useState(false);

  if (isLoading) {
    return (
      <View className="flex-1 items-center justify-center bg-gray-50">
        <ActivityIndicator size="large" color="#E60023" />
      </View>
    );
  }

  return (
    <View className="flex-1 bg-gray-50">
      <JobFilters
        visible={showFilters}
        onClose={() => setShowFilters(false)}
      />

      <FlatList
        data={jobs}
        keyExtractor={(item) => item.id}
        renderItem={({ item }) => <JobCard job={item} />}
        contentContainerStyle={{ padding: 16 }}
        ItemSeparatorComponent={() => <View className="h-3" />}
        ListEmptyComponent={
          <View className="items-center justify-center py-12">
            <Text className="text-gray-500 text-center">
              No jobs found
            </Text>
          </View>
        }
        refreshControl={
          <RefreshControl
            refreshing={isRefreshing}
            onRefresh={refetch}
            tintColor="#E60023"
          />
        }
      />
    </View>
  );
}