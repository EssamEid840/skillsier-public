import { useState } from 'react';
import { View, Text, ScrollView, TouchableOpacity, Image, Alert, TextInput } from 'react-native';
import { SafeAreaView } from 'react-native-safe-area-context';
import { router } from 'expo-router';
import { useTranslation } from 'react-i18next';
import * as ImagePicker from 'expo-image-picker';
import { useAuth } from '@skillsier/shared';
import { 
  useUpdateProfile, 
  useUpdateFreelancerProfile,
  useUploadAvatar,
  useDeleteAvatar 
} from '@skillsier/shared';
import { Button, Input, Avatar } from '@skillsier/ui';
import { ArrowLeft, Camera, X } from 'lucide-react-native';

export default function EditProfileScreen() {
  const { t } = useTranslation();
  const { user } = useAuth();
  
  const [formData, setFormData] = useState({
    firstName: user?.firstName || '',
    lastName: user?.lastName || '',
    title: user?.title || '',
    bio: user?.bio || '',
    phoneNumber: user?.phoneNumber || '',
    city: user?.city || '',
    country: user?.country || '',
    // Freelancer specific
    professionalTitle: '',
    overview: '',
    hourlyRate: 0,
    availability: 'AVAILABLE',
  });

  const [avatarUri, setAvatarUri] = useState<string | null>(user?.avatar || null);
  const [avatarFile, setAvatarFile] = useState<any>(null);

  const updateProfile = useUpdateProfile();
  const updateFreelancerProfile = useUpdateFreelancerProfile();
  const uploadAvatar = useUploadAvatar();
  const deleteAvatar = useDeleteAvatar();

  const pickImage = async () => {
    const { status } = await ImagePicker.requestMediaLibraryPermissionsAsync();
    
    if (status !== 'granted') {
      Alert.alert('Permission needed', 'We need camera roll permissions to upload photos');
      return;
    }

    const result = await ImagePicker.launchImageLibraryAsync({
      mediaTypes: ImagePicker.MediaTypeOptions.Images,
      allowsEditing: true,
      aspect: [1, 1],
      quality: 0.8,
    });

    if (!result.canceled) {
      setAvatarUri(result.assets[0].uri);
      setAvatarFile(result.assets[0]);
    }
  };

  const handleRemoveAvatar = async () => {
    try {
      await deleteAvatar.mutateAsync();
      setAvatarUri(null);
      setAvatarFile(null);
    } catch (error) {
      Alert.alert('Error', 'Failed to remove avatar');
    }
  };

  const handleSave = async () => {
    try {
      // Upload avatar if changed
      if (avatarFile) {
        const formData = new FormData();
        formData.append('avatar', {
          uri: avatarFile.uri,
          type: 'image/jpeg',
          name: 'avatar.jpg',
        } as any);
        await uploadAvatar.mutateAsync(formData as any);
      }

      // Update basic profile
      await updateProfile.mutateAsync({
        firstName: formData.firstName,
        lastName: formData.lastName,
        title: formData.title,
        bio: formData.bio,
        phoneNumber: formData.phoneNumber,
        city: formData.city,
        country: formData.country,
      });

      // Update freelancer profile if applicable
      if (user?.userType === 'FREELANCER' || user?.userType === 'BOTH') {
        await updateFreelancerProfile.mutateAsync({
          professionalTitle: formData.professionalTitle,
          overview: formData.overview,
          hourlyRate: formData.hourlyRate,
          availability: formData.availability as any,
        });
      }

      Alert.alert('Success', 'Profile updated successfully');
      router.back();
    } catch (error) {
      Alert.alert('Error', 'Failed to update profile');
    }
  };

  const isLoading = updateProfile.isPending || updateFreelancerProfile.isPending || uploadAvatar.isPending;

  return (
    <SafeAreaView className="flex-1 bg-white">
      <View className="flex-row items-center justify-between px-6 py-4 border-b border-gray-200">
        <TouchableOpacity onPress={() => router.back()}>
          <ArrowLeft color="#000" size={24} />
        </TouchableOpacity>
        <Text className="text-xl font-bold text-gray-900">Edit Profile</Text>
        <View className="w-6" />
      </View>

      <ScrollView className="flex-1 px-6 py-6" showsVerticalScrollIndicator={false}>
        {/* Avatar Section */}
        <View className="items-center mb-8">
          <View className="relative">
            <Avatar 
              src={avatarUri || undefined} 
              alt={user?.username || 'User'} 
              size="xl"
              className="h-24 w-24"
            />
            <TouchableOpacity
              onPress={pickImage}
              className="absolute bottom-0 right-0 p-2 bg-primary-600 rounded-full"
            >
              <Camera color="#fff" size={16} />
            </TouchableOpacity>
          </View>
          {avatarUri && (
            <TouchableOpacity onPress={handleRemoveAvatar} className="mt-2">
              <Text className="text-red-600 text-sm">Remove Photo</Text>
            </TouchableOpacity>
          )}
        </View>

        {/* Basic Information */}
        <View className="space-y-4 mb-6">
          <Text className="text-lg font-semibold text-gray-900">Basic Information</Text>
          
          <Input
            label="First Name"
            value={formData.firstName}
            onChangeText={(text) => setFormData({ ...formData, firstName: text })}
          />
          
          <Input
            label="Last Name"
            value={formData.lastName}
            onChangeText={(text) => setFormData({ ...formData, lastName: text })}
          />
          
          <Input
            label="Professional Title"
            value={formData.title}
            onChangeText={(text) => setFormData({ ...formData, title: text })}
            placeholder="e.g., Full Stack Developer"
          />
          
          <View>
            <Text className="text-sm font-medium text-gray-700 mb-2">Bio</Text>
            <TextInput
              value={formData.bio}
              onChangeText={(text) => setFormData({ ...formData, bio: text })}
              multiline
              numberOfLines={4}
              className="w-full px-4 py-3 border border-gray-300 rounded-lg text-gray-900"
              placeholder="Tell us about yourself..."
            />
          </View>
          
          <Input
            label="Phone Number"
            value={formData.phoneNumber}
            onChangeText={(text) => setFormData({ ...formData, phoneNumber: text })}
            keyboardType="phone-pad"
          />
          
          <Input
            label="City"
            value={formData.city}
            onChangeText={(text) => setFormData({ ...formData, city: text })}
          />
          
          <Input
            label="Country"
            value={formData.country}
            onChangeText={(text) => setFormData({ ...formData, country: text })}
          />
        </View>

        {/* Freelancer Specific */}
        {(user?.userType === 'FREELANCER' || user?.userType === 'BOTH') && (
          <View className="space-y-4 mb-6">
            <Text className="text-lg font-semibold text-gray-900">Freelancer Settings</Text>
            
            <Input
              label="Professional Title"
              value={formData.professionalTitle}
              onChangeText={(text) => setFormData({ ...formData, professionalTitle: text })}
              placeholder="e.g., Senior React Developer"
            />
            
            <View>
              <Text className="text-sm font-medium text-gray-700 mb-2">
                Professional Overview
              </Text>
              <TextInput
                value={formData.overview}
                onChangeText={(text) => setFormData({ ...formData, overview: text })}
                multiline
                numberOfLines={6}
                className="w-full px-4 py-3 border border-gray-300 rounded-lg text-gray-900"
                placeholder="Describe your expertise and experience..."
              />
            </View>
            
            <Input
              label="Hourly Rate (USD)"
              value={String(formData.hourlyRate)}
              onChangeText={(text) => setFormData({ ...formData, hourlyRate: Number(text) })}
              keyboardType="numeric"
            />
            
            <View>
              <Text className="text-sm font-medium text-gray-700 mb-2">Availability</Text>
              <View className="flex-row gap-2">
                {['AVAILABLE', 'BUSY', 'NOT_AVAILABLE'].map((status) => (
                  <TouchableOpacity
                    key={status}
                    onPress={() => setFormData({ ...formData, availability: status })}
                    className={`flex-1 px-4 py-3 rounded-lg border ${
                      formData.availability === status
                        ? 'bg-primary-50 border-primary-600'
                        : 'bg-white border-gray-300'
                    }`}
                  >
                    <Text
                      className={`text-center text-sm ${
                        formData.availability === status ? 'text-primary-600 font-medium' : 'text-gray-700'
                      }`}
                    >
                      {status === 'AVAILABLE' ? 'Available' : status === 'BUSY' ? 'Busy' : 'Unavailable'}
                    </Text>
                  </TouchableOpacity>
                ))}
              </View>
            </View>
          </View>
        )}

        <Button onPress={handleSave} loading={isLoading} fullWidth size="lg">
          Save Changes
        </Button>

        <View className="h-8" />
      </ScrollView>
    </SafeAreaView>
  );
}