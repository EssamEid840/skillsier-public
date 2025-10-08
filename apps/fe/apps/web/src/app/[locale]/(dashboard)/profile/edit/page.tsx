'use client';

import { useState } from 'react';
import { useTranslations } from 'next-intl';
import { useRouter } from 'next/navigation';
import { useAuth } from '@skillsier/shared';
import { 
  useUpdateProfile, 
  useUpdateFreelancerProfile,
  useUploadAvatar,
  useDeleteAvatar 
} from '@skillsier/shared';
import { Card, Button, Input, Avatar } from '@skillsier/ui';
import { Upload, X, Loader } from 'lucide-react';

export default function EditProfilePage() {
  const t = useTranslations('profile');
  const router = useRouter();
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

  const updateProfile = useUpdateProfile();
  const updateFreelancerProfile = useUpdateFreelancerProfile();
  const uploadAvatar = useUploadAvatar();
  const deleteAvatar = useDeleteAvatar();

  const [avatarFile, setAvatarFile] = useState<File | null>(null);
  const [avatarPreview, setAvatarPreview] = useState<string | null>(user?.avatar || null);

  const handleAvatarChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const file = e.target.files?.[0];
    if (file) {
      setAvatarFile(file);
      const reader = new FileReader();
      reader.onloadend = () => {
        setAvatarPreview(reader.result as string);
      };
      reader.readAsDataURL(file);
    }
  };

  const handleRemoveAvatar = async () => {
    try {
      await deleteAvatar.mutateAsync();
      setAvatarFile(null);
      setAvatarPreview(null);
    } catch (error) {
      console.error('Failed to delete avatar:', error);
    }
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();

    try {
      // Upload avatar if changed
      if (avatarFile) {
        await uploadAvatar.mutateAsync(avatarFile);
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

      router.push('/profile');
    } catch (error) {
      console.error('Failed to update profile:', error);
    }
  };

  const isLoading = updateProfile.isPending || updateFreelancerProfile.isPending || uploadAvatar.isPending;

  return (
    <div className="max-w-4xl mx-auto space-y-6">
      <div className="flex items-center justify-between">
        <h1 className="text-3xl font-bold text-gray-900">Edit Profile</h1>
        <Button variant="ghost" onClick={() => router.back()}>
          Cancel
        </Button>
      </div>

      <form onSubmit={handleSubmit} className="space-y-6">
        {/* Avatar Section */}
        <Card padding="lg">
          <h2 className="text-lg font-semibold text-gray-900 mb-4">Profile Photo</h2>
          <div className="flex items-center gap-6">
            <Avatar 
              src={avatarPreview || undefined} 
              alt={user?.username || 'User'} 
              size="xl"
              className="h-24 w-24"
            />
            <div className="flex-1">
              <div className="flex gap-3">
                <label className="cursor-pointer">
                  <input
                    type="file"
                    accept="image/*"
                    onChange={handleAvatarChange}
                    className="hidden"
                  />
                  <div className="flex items-center gap-2 px-4 py-2 bg-primary-600 text-white rounded-lg hover:bg-primary-700 transition-colors">
                    <Upload className="h-4 w-4" />
                    Upload Photo
                  </div>
                </label>
                {avatarPreview && (
                  <button
                    type="button"
                    onClick={handleRemoveAvatar}
                    className="flex items-center gap-2 px-4 py-2 bg-red-50 text-red-600 rounded-lg hover:bg-red-100 transition-colors"
                  >
                    <X className="h-4 w-4" />
                    Remove
                  </button>
                )}
              </div>
              <p className="text-sm text-gray-600 mt-2">
                JPG, PNG or GIF. Max size 5MB. Recommended: 400x400px
              </p>
            </div>
          </div>
        </Card>

        {/* Basic Information */}
        <Card padding="lg">
          <h2 className="text-lg font-semibold text-gray-900 mb-4">Basic Information</h2>
          <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
            <Input
              label="First Name"
              value={formData.firstName}
              onChange={(e) => setFormData({ ...formData, firstName: e.target.value })}
              required
            />
            <Input
              label="Last Name"
              value={formData.lastName}
              onChange={(e) => setFormData({ ...formData, lastName: e.target.value })}
              required
            />
            <Input
              label="Professional Title"
              value={formData.title}
              onChange={(e) => setFormData({ ...formData, title: e.target.value })}
              placeholder="e.g., Full Stack Developer"
              className="md:col-span-2"
            />
            <Input
              label="Phone Number"
              type="tel"
              value={formData.phoneNumber}
              onChange={(e) => setFormData({ ...formData, phoneNumber: e.target.value })}
            />
            <Input
              label="City"
              value={formData.city}
              onChange={(e) => setFormData({ ...formData, city: e.target.value })}
            />
            <Input
              label="Country"
              value={formData.country}
              onChange={(e) => setFormData({ ...formData, country: e.target.value })}
              className="md:col-span-2"
            />
          </div>
          
          <div className="mt-4">
            <label className="block text-sm font-medium text-gray-700 mb-2">
              Bio
            </label>
            <textarea
              value={formData.bio}
              onChange={(e) => setFormData({ ...formData, bio: e.target.value })}
              rows={4}
              className="w-full px-4 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-primary-500"
              placeholder="Tell us about yourself..."
            />
          </div>
        </Card>

        {/* Freelancer Specific */}
        {(user?.userType === 'FREELANCER' || user?.userType === 'BOTH') && (
          <Card padding="lg">
            <h2 className="text-lg font-semibold text-gray-900 mb-4">Freelancer Settings</h2>
            <div className="space-y-4">
              <Input
                label="Professional Title"
                value={formData.professionalTitle}
                onChange={(e) => setFormData({ ...formData, professionalTitle: e.target.value })}
                placeholder="e.g., Senior React Developer"
              />
              
              <div>
                <label className="block text-sm font-medium text-gray-700 mb-2">
                  Professional Overview
                </label>
                <textarea
                  value={formData.overview}
                  onChange={(e) => setFormData({ ...formData, overview: e.target.value })}
                  rows={6}
                  className="w-full px-4 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-primary-500"
                  placeholder="Describe your expertise, experience, and what makes you unique..."
                />
              </div>

              <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
                <Input
                  label="Hourly Rate (USD)"
                  type="number"
                  value={formData.hourlyRate}
                  onChange={(e) => setFormData({ ...formData, hourlyRate: Number(e.target.value) })}
                  min="0"
                  step="5"
                />
                
                <div>
                  <label className="block text-sm font-medium text-gray-700 mb-2">
                    Availability
                  </label>
                  <select
                    value={formData.availability}
                    onChange={(e) => setFormData({ ...formData, availability: e.target.value })}
                    className="w-full px-4 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-primary-500"
                  >
                    <option value="AVAILABLE">Available</option>
                    <option value="BUSY">Busy</option>
                    <option value="NOT_AVAILABLE">Not Available</option>
                  </select>
                </div>
              </div>
            </div>
          </Card>
        )}

        {/* Submit Buttons */}
        <div className="flex justify-end gap-3">
          <Button type="button" variant="outline" onClick={() => router.back()}>
            Cancel
          </Button>
          <Button type="submit" loading={isLoading}>
            {isLoading ? (
              <>
                <Loader className="h-4 w-4 animate-spin mr-2" />
                Saving...
              </>
            ) : (
              'Save Changes'
            )}
          </Button>
        </div>
      </form>
    </div>
  );
}