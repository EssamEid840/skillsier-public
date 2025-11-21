import { Button, Input } from '@skillsier/ui';
import { useRouter } from 'expo-router';
import { useState } from 'react';
import {
  KeyboardAvoidingView,
  Platform,
  ScrollView,
  Text,
  View,
} from 'react-native';

export default function CreateJobScreen() {
  const router = useRouter();
  const [title, setTitle] = useState('');
  const [description, setDescription] = useState('');

  const handleSubmit = () => {
    // TODO: Implement job creation
    router.back();
  };

  return (
    <KeyboardAvoidingView
      behavior={Platform.OS === 'ios' ? 'padding' : 'height'}
      className="flex-1"
    >
      <ScrollView className="flex-1 bg-white">
        <View className="p-4 space-y-4">
          <View>
            <Text className="text-sm font-medium text-gray-700 mb-2">
              Job Title
            </Text>
            <Input
              placeholder="e.g., Full-Stack Developer"
              value={title}
              onChangeText={setTitle}
            />
          </View>

          <View>
            <Text className="text-sm font-medium text-gray-700 mb-2">
              Description
            </Text>
            <Input
              placeholder="Describe the job requirements..."
              value={description}
              onChangeText={setDescription}
              multiline
              numberOfLines={6}
              textAlignVertical="top"
            />
          </View>

          <Button
            onPress={handleSubmit}
            disabled={!title || !description}
            className="mt-4"
          >
            <Text className="text-white font-semibold">Post Job</Text>
          </Button>
        </View>
      </ScrollView>
    </KeyboardAvoidingView>
  );
}