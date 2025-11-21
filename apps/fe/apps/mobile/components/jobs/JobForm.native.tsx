import { Button, Input } from '@skillsier/ui';
import { useCreateJob } from '@skillsier/hooks';
import { JobStatus } from '@skillsier/types';
import { useRouter } from 'expo-router';
import { useState } from 'react';
import {
  Alert,
  KeyboardAvoidingView,
  Platform,
  Pressable,
  ScrollView,
  Text,
  TextInput,
  View,
} from 'react-native';

export default function JobForm() {
  const router = useRouter();
  const { mutate: createJob, isPending } = useCreateJob();

  const [title, setTitle] = useState('');
  const [description, setDescription] = useState('');
  const [budget, setBudget] = useState('');
  const [skills, setSkills] = useState('');
  const [duration, setDuration] = useState('');
  const [showDurationPicker, setShowDurationPicker] = useState(false);

  const durations = [
    { value: 'less-than-1-month', label: 'Less than 1 month' },
    { value: '1-3-months', label: '1-3 months' },
    { value: '3-6-months', label: '3-6 months' },
    { value: 'more-than-6-months', label: 'More than 6 months' },
  ];

  const handleSubmit = () => {
    if (!title.trim()) {
      Alert.alert('Error', 'Please enter a job title');
      return;
    }
    if (title.length < 10) {
      Alert.alert('Error', 'Title must be at least 10 characters');
      return;
    }
    if (!description.trim()) {
      Alert.alert('Error', 'Please enter a job description');
      return;
    }
    if (description.length < 50) {
      Alert.alert('Error', 'Description must be at least 50 characters');
      return;
    }

    createJob(
      {
        title: title.trim(),
        description: description.trim(),
        budget: budget ? Number(budget) : undefined,
        skills: skills
          .split(',')
          .map((s) => s.trim())
          .filter((s) => s),
        duration: duration || undefined,
        status: JobStatus.DRAFT,
      },
      {
        onSuccess: (job) => {
          Alert.alert('Success', 'Job posted successfully', [
            {
              text: 'OK',
              onPress: () => router.push(`/(authenticated)/(tabs)/jobs/${job.id}`),
            },
          ]);
        },
        onError: () => {
          Alert.alert('Error', 'Failed to post job. Please try again.');
        },
      }
    );
  };

  return (
    <KeyboardAvoidingView
      behavior={Platform.OS === 'ios' ? 'padding' : 'height'}
      className="flex-1"
      keyboardVerticalOffset={100}
    >
      <ScrollView className="flex-1 bg-white" keyboardShouldPersistTaps="handled">
        <View className="p-4 space-y-6">
          {/* Title */}
          <View>
            <Text className="text-sm font-medium text-gray-700 mb-2">
              Job Title <Text className="text-red-500">*</Text>
            </Text>
            <TextInput
              placeholder="e.g., Full-Stack Developer for SaaS Platform"
              value={title}
              onChangeText={setTitle}
              className="px-4 py-3 bg-gray-50 rounded-lg text-gray-900"
              multiline
            />
            <Text className="text-xs text-gray-500 mt-1">
              {title.length}/200 characters
            </Text>
          </View>

          {/* Description */}
          <View>
            <Text className="text-sm font-medium text-gray-700 mb-2">
              Description <Text className="text-red-500">*</Text>
            </Text>
            <TextInput
              placeholder="Describe the job requirements, responsibilities, and qualifications..."
              value={description}
              onChangeText={setDescription}
              multiline
              numberOfLines={8}
              textAlignVertical="top"
              className="px-4 py-3 bg-gray-50 rounded-lg text-gray-900 min-h-[160px]"
            />
            <Text className="text-xs text-gray-500 mt-1">
              {description.length} characters (min 50)
            </Text>
          </View>

          {/* Budget */}
          <View>
            <Text className="text-sm font-medium text-gray-700 mb-2">
              Budget (USD)
            </Text>
            <TextInput
              placeholder="5000"
              value={budget}
              onChangeText={setBudget}
              keyboardType="numeric"
              className="px-4 py-3 bg-gray-50 rounded-lg text-gray-900"
            />
          </View>

          {/* Skills */}
          <View>
            <Text className="text-sm font-medium text-gray-700 mb-2">
              Required Skills
            </Text>
            <TextInput
              placeholder="React, Node.js, TypeScript, PostgreSQL"
              value={skills}
              onChangeText={setSkills}
              className="px-4 py-3 bg-gray-50 rounded-lg text-gray-900"
            />
            <Text className="text-xs text-gray-500 mt-1">
              Separate skills with commas
            </Text>
          </View>

          {/* Duration Picker */}
          <View>
            <Text className="text-sm font-medium text-gray-700 mb-2">
              Project Duration
            </Text>
            <Pressable
              onPress={() => setShowDurationPicker(!showDurationPicker)}
              className="px-4 py-3 bg-gray-50 rounded-lg"
            >
              <Text className={duration ? 'text-gray-900' : 'text-gray-400'}>
                {duration
                  ? durations.find((d) => d.value === duration)?.label
                  : 'Select duration'}
              </Text>
            </Pressable>

            {showDurationPicker && (
              <View className="mt-2 bg-white border border-gray-200 rounded-lg">
                {durations.map((d) => (
                  <Pressable
                    key={d.value}
                    onPress={() => {
                      setDuration(d.value);
                      setShowDurationPicker(false);
                    }}
                    className="px-4 py-3 border-b border-gray-100 active:bg-gray-50"
                  >
                    <Text
                      className={
                        duration === d.value
                          ? 'text-primary-500 font-medium'
                          : 'text-gray-900'
                      }
                    >
                      {d.label}
                    </Text>
                  </Pressable>
                ))}
              </View>
            )}
          </View>

          {/* Submit Button */}
          <View className="pt-4 space-y-3">
            <Button onPress={handleSubmit} disabled={isPending}>
              <Text className="text-white font-semibold">
                {isPending ? 'Posting...' : 'Post Job'}
              </Text>
            </Button>
            <Button
              onPress={() => router.back()}
              variant="outline"
              disabled={isPending}
            >
              <Text className="text-gray-700 font-semibold">Cancel</Text>
            </Button>
          </View>
        </View>
      </ScrollView>
    </KeyboardAvoidingView>
  );
}