import { Button } from '@skillsier/ui';
import { useJobsStore } from '@skillsier/stores';
import { JobStatus } from '@skillsier/types';
import { useState } from 'react';
import {
  Modal,
  Pressable,
  ScrollView,
  Text,
  TextInput,
  View,
} from 'react-native';

interface JobFiltersProps {
  visible: boolean;
  onClose: () => void;
}

export default function JobFilters({ visible, onClose }: JobFiltersProps) {
  const { filters, setFilters, resetFilters } = useJobsStore();
  const [localFilters, setLocalFilters] = useState(filters);

  const statuses = [
    { value: JobStatus.DRAFT, label: 'Draft' },
    { value: JobStatus.ACTIVE, label: 'Active' },
    { value: JobStatus.CLOSED, label: 'Closed' },
    { value: JobStatus.CANCELLED, label: 'Cancelled' },
  ];

  const handleApply = () => {
    setFilters(localFilters);
    onClose();
  };

  const handleReset = () => {
    setLocalFilters({});
    resetFilters();
    onClose();
  };

  const toggleStatus = (status: JobStatus) => {
    const currentStatuses = localFilters.status || [];
    if (currentStatuses.includes(status)) {
      setLocalFilters({
        ...localFilters,
        status: currentStatuses.filter((s) => s !== status),
      });
    } else {
      setLocalFilters({
        ...localFilters,
        status: [...currentStatuses, status],
      });
    }
  };

  return (
    <Modal
      visible={visible}
      animationType="slide"
      presentationStyle="pageSheet"
      onRequestClose={onClose}
    >
      <View className="flex-1 bg-white">
        {/* Header */}
        <View className="flex-row items-center justify-between p-4 border-b border-gray-200">
          <Text className="text-xl font-semibold text-gray-900">Filters</Text>
          <Pressable onPress={onClose}>
            <Text className="text-base text-primary-500 font-medium">Done</Text>
          </Pressable>
        </View>

        <ScrollView className="flex-1 p-4">
          {/* Search */}
          <View className="mb-6">
            <Text className="text-sm font-medium text-gray-700 mb-2">
              Search
            </Text>
            <TextInput
              placeholder="Search jobs..."
              value={localFilters.search || ''}
              onChangeText={(text) =>
                setLocalFilters({ ...localFilters, search: text || undefined })
              }
              className="px-4 py-3 bg-gray-50 rounded-lg text-gray-900"
            />
          </View>

          {/* Status */}
          <View className="mb-6">
            <Text className="text-sm font-medium text-gray-700 mb-3">
              Status
            </Text>
            <View className="space-y-2">
              {statuses.map((status) => (
                <Pressable
                  key={status.value}
                  onPress={() => toggleStatus(status.value)}
                  className="flex-row items-center py-2"
                >
                  <View
                    className={`w-5 h-5 rounded border-2 mr-3 items-center justify-center ${
                      localFilters.status?.includes(status.value)
                        ? 'bg-primary-500 border-primary-500'
                        : 'border-gray-300'
                    }`}
                  >
                    {localFilters.status?.includes(status.value) && (
                      <Text className="text-white text-xs">✓</Text>
                    )}
                  </View>
                  <Text className="text-gray-900">{status.label}</Text>
                </Pressable>
              ))}
            </View>
          </View>

          {/* Budget Range */}
          <View className="mb-6">
            <Text className="text-sm font-medium text-gray-700 mb-3">
              Budget Range
            </Text>
            <View className="flex-row items-center gap-3">
              <View className="flex-1">
                <TextInput
                  placeholder="Min"
                  keyboardType="numeric"
                  value={localFilters.budgetMin?.toString() || ''}
                  onChangeText={(text) =>
                    setLocalFilters({
                      ...localFilters,
                      budgetMin: text ? Number(text) : undefined,
                    })
                  }
                  className="px-4 py-3 bg-gray-50 rounded-lg text-gray-900"
                />
              </View>
              <Text className="text-gray-500">-</Text>
              <View className="flex-1">
                <TextInput
                  placeholder="Max"
                  keyboardType="numeric"
                  value={localFilters.budgetMax?.toString() || ''}
                  onChangeText={(text) =>
                    setLocalFilters({
                      ...localFilters,
                      budgetMax: text ? Number(text) : undefined,
                    })
                  }
                  className="px-4 py-3 bg-gray-50 rounded-lg text-gray-900"
                />
              </View>
            </View>
          </View>

          {/* Sort By */}
          <View className="mb-6">
            <Text className="text-sm font-medium text-gray-700 mb-3">
              Sort By
            </Text>
            <View className="space-y-2">
              {[
                { value: 'recent', label: 'Most Recent' },
                { value: 'budget-high', label: 'Highest Budget' },
                { value: 'budget-low', label: 'Lowest Budget' },
                { value: 'proposals', label: 'Most Proposals' },
              ].map((option) => (
                <Pressable
                  key={option.value}
                  onPress={() =>
                    setLocalFilters({ ...localFilters, sortBy: option.value })
                  }
                  className="flex-row items-center py-2"
                >
                  <View
                    className={`w-5 h-5 rounded-full border-2 mr-3 ${
                      localFilters.sortBy === option.value
                        ? 'bg-primary-500 border-primary-500'
                        : 'border-gray-300'
                    }`}
                  >
                    {localFilters.sortBy === option.value && (
                      <View className="w-2 h-2 bg-white rounded-full self-center" />
                    )}
                  </View>
                  <Text className="text-gray-900">{option.label}</Text>
                </Pressable>
              ))}
            </View>
          </View>
        </ScrollView>

        {/* Footer Actions */}
        <View className="p-4 border-t border-gray-200 space-y-3">
          <Button onPress={handleApply}>
            <Text className="text-white font-semibold">Apply Filters</Text>
          </Button>
          <Button onPress={handleReset} variant="outline">
            <Text className="text-primary-500 font-semibold">Reset All</Text>
          </Button>
        </View>
      </View>
    </Modal>
  );
}