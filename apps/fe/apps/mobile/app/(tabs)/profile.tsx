import { View, Text, ScrollView, TouchableOpacity, Image } from 'react-native';
import { SafeAreaView } from 'react-native-safe-area-context';
import { router } from 'expo-router';
import { useTranslation } from 'react-i18next';
import { useAuth, useLogout, useFreelancerProfile, useFreelancerSkills } from '@skillsier/shared';
import { Avatar, Card, Badge } from '@skillsier/ui';
import { LanguageSwitcher } from '@/components/LanguageSwitcher';
import {
  User,
  Edit,
  Briefcase,
  Award,
  DollarSign,
  Star,
  Settings,
  Bell,
  HelpCircle,
  LogOut,
  ChevronRight,
  Eye,
  CheckCircle,
} from 'lucide-react-native';

export default function ProfileScreen() {
  const { t } = useTranslation();
  const { user } = useAuth();
  const { mutate: logout } = useLogout();
  const { data: profile } = useFreelancerProfile();
  const { data: skills } = useFreelancerSkills();

  const isFreelancer = user?.userType === 'FREELANCER' || user?.userType === 'BOTH';

  const handleLogout = () => {
    logout(undefined, {
      onSuccess: () => {
        router.replace('/');
      },
    });
  };

  return (
    <SafeAreaView className="flex-1 bg-gray-50">
      <View className="px-6 py-4 bg-white border-b border-gray-200">
        <Text className="text-2xl font-bold text-gray-900">{t('profile.title')}</Text>
      </View>

      <ScrollView className="flex-1" showsVerticalScrollIndicator={false}>
        {/* Profile Header */}
        <View className="px-6 py-8 bg-white border-b border-gray-200">
          <View className="flex-row items-start gap-4">
            <View className="relative">
              <Avatar src={user?.avatar} alt={user?.username || 'User'} size="xl" />
              {profile?.availability === 'AVAILABLE' && (
                <View className="absolute bottom-0 right-0 h-4 w-4 bg-green-500 border-2 border-white rounded-full" />
              )}
            </View>
            
            <View className="flex-1">
              <Text className="text-xl font-bold text-gray-900">
                {user?.firstName} {user?.lastName}
              </Text>
              <Text className="text-base text-gray-700 mt-1">
                {profile?.professionalTitle || user?.title}
              </Text>
              <Text className="text-sm text-gray-600">@{user?.username}</Text>
              
              <View className="flex-row flex-wrap gap-2 mt-2">
                <Badge variant={user?.emailVerified ? 'success' : 'warning'} size="sm">
                  {user?.emailVerified ? '✓ Verified' : 'Unverified'}
                </Badge>
                {user?.identityVerified && (
                  <Badge variant="success" size="sm">✓ ID</Badge>
                )}
                <Badge size="sm">{user?.userType}</Badge>
              </View>
            </View>

            <TouchableOpacity onPress={() => router.push('/profile/edit')}>
              <View className="p-2 bg-gray-100 rounded-lg">
                <Edit color="#6b7280" size={20} />
              </View>
            </TouchableOpacity>
          </View>

          {profile?.overview && (
            <Text className="text-gray-700 mt-4 leading-relaxed" numberOfLines={3}>
              {profile.overview}
            </Text>
          )}

          {isFreelancer && profile && (
            <View className="flex-row items-center gap-4 mt-4">
              <View className="flex-row items-center gap-1">
                <DollarSign color="#10b981" size={18} />
                <Text className="font-semibold text-gray-900">${profile.hourlyRate}/hr</Text>
              </View>
              <View className="flex-row items-center gap-1">
                <Star color="#f59e0b" size={18} />
                <Text className="font-semibold text-gray-900">{profile.rating.toFixed(1)}</Text>
                <Text className="text-gray-600 text-sm">({profile.totalReviews})</Text>
              </View>
            </View>
          )}
        </View>

        {/* Stats Section - Freelancer Only */}
        {isFreelancer && profile && (
          <View className="px-6 py-6 bg-white border-b border-gray-200">
            <Text className="text-lg font-semibold text-gray-900 mb-4">Statistics</Text>
            <View className="flex-row flex-wrap -mx-2">
              <View className="w-1/2 px-2 mb-4">
                <Card padding="md">
                  <View className="flex-row items-center gap-2 mb-1">
                    <DollarSign color="#10b981" size={20} />
                    <Text className="text-2xl font-bold text-gray-900">
                      ${profile.totalEarnings.toLocaleString()}
                    </Text>
                  </View>
                  <Text className="text-sm text-gray-600">Total Earned</Text>
                </Card>
              </View>
              
              <View className="w-1/2 px-2 mb-4">
                <Card padding="md">
                  <View className="flex-row items-center gap-2 mb-1">
                    <CheckCircle color="#3b82f6" size={20} />
                    <Text className="text-2xl font-bold text-gray-900">{profile.completedJobs}</Text>
                  </View>
                  <Text className="text-sm text-gray-600">Jobs Done</Text>
                </Card>
              </View>
              
              <View className="w-1/2 px-2 mb-4">
                <Card padding="md">
                  <View className="flex-row items-center gap-2 mb-1">
                    <Star color="#f59e0b" size={20} />
                    <Text className="text-2xl font-bold text-gray-900">
                      {profile.rating.toFixed(1)}
                    </Text>
                  </View>
                  <Text className="text-sm text-gray-600">Rating</Text>
                </Card>
              </View>
              
              <View className="w-1/2 px-2 mb-4">
                <Card padding="md">
                  <View className="flex-row items-center gap-2 mb-1">
                    <Award color="#8b5cf6" size={20} />
                    <Text className="text-2xl font-bold text-gray-900">{profile.successRate}%</Text>
                  </View>
                  <Text className="text-sm text-gray-600">Success</Text>
                </Card>
              </View>
            </View>
          </View>
        )}

        {/* Skills Preview */}
        {isFreelancer && skills && skills.length > 0 && (
          <View className="px-6 py-6 bg-white border-b border-gray-200">
            <View className="flex-row items-center justify-between mb-4">
              <Text className="text-lg font-semibold text-gray-900">Top Skills</Text>
              <TouchableOpacity onPress={() => router.push('/profile/skills')}>
                <Text className="text-primary-600 font-medium">View All</Text>
              </TouchableOpacity>
            </View>
            <View className="flex-row flex-wrap gap-2">
              {skills.slice(0, 6).map((skill) => (
                <View key={skill.id} className="px-3 py-2 bg-primary-50 rounded-lg">
                  <Text className="text-primary-700 font-medium">{skill.name}</Text>
                </View>
              ))}
            </View>
          </View>
        )}

        {/* Profile Strength */}
        {isFreelancer && profile && (
          <View className="px-6 py-6 bg-white border-b border-gray-200">
            <View className="flex-row items-center justify-between mb-3">
              <Text className="text-lg font-semibold text-gray-900">Profile Strength</Text>
              <Text className="text-xl font-bold text-primary-600">{profile.profileStrength}%</Text>
            </View>
            <View className="w-full bg-gray-200 rounded-full h-2">
              <View
                className="bg-primary-600 h-2 rounded-full"
                style={{ width: `${profile.profileStrength}%` }}
              />
            </View>
            {profile.profileStrength < 100 && (
              <Text className="text-sm text-gray-600 mt-2">
                Complete your profile to get more job opportunities
              </Text>
            )}
          </View>
        )}

        {/* Menu Options */}
        <View className="px-6 py-6">
          <Card padding="none">
            <TouchableOpacity
              onPress={() => router.push('/profile/edit')}
              className="flex-row items-center justify-between p-4 border-b border-gray-200"
            >
              <View className="flex-row items-center">
                <User color="#6b7280" size={20} />
                <Text className="ml-3 text-gray-900">{t('profile.editProfile')}</Text>
              </View>
              <ChevronRight color="#9ca3af" size={20} />
            </TouchableOpacity>

            {isFreelancer && (
              <>
                <TouchableOpacity
                  onPress={() => router.push('/profile/portfolio')}
                  className="flex-row items-center justify-between p-4 border-b border-gray-200"
                >
                  <View className="flex-row items-center">
                    <Briefcase color="#6b7280" size={20} />
                    <Text className="ml-3 text-gray-900">Portfolio</Text>
                  </View>
                  <ChevronRight color="#9ca3af" size={20} />
                </TouchableOpacity>

                <TouchableOpacity
                  onPress={() => router.push('/profile/skills')}
                  className="flex-row items-center justify-between p-4 border-b border-gray-200"
                >
                  <View className="flex-row items-center">
                    <Award color="#6b7280" size={20} />
                    <Text className="ml-3 text-gray-900">Skills & Experience</Text>
                  </View>
                  <ChevronRight color="#9ca3af" size={20} />
                </TouchableOpacity>
              </>
            )}

            <TouchableOpacity
              onPress={() => router.push('/profile/settings')}
              className="flex-row items-center justify-between p-4 border-b border-gray-200"
            >
              <View className="flex-row items-center">
                <Settings color="#6b7280" size={20} />
                <Text className="ml-3 text-gray-900">{t('profile.settings')}</Text>
              </View>
              <ChevronRight color="#9ca3af" size={20} />
            </TouchableOpacity>

            <TouchableOpacity
              onPress={() => {}}
              className="flex-row items-center justify-between p-4 border-b border-gray-200"
            >
              <View className="flex-row items-center">
                <Bell color="#6b7280" size={20} />
                <Text className="ml-3 text-gray-900">{t('profile.notifications')}</Text>
              </View>
              <ChevronRight color="#9ca3af" size={20} />
            </TouchableOpacity>

            <TouchableOpacity
              onPress={() => {}}
              className="flex-row items-center justify-between p-4 border-b border-gray-200"
            >
              <View className="flex-row items-center">
                <HelpCircle color="#6b7280" size={20} />
                <Text className="ml-3 text-gray-900">{t('profile.helpSupport')}</Text>
              </View>
              <ChevronRight color="#9ca3af" size={20} />
            </TouchableOpacity>

            <LanguageSwitcher />
          </Card>

          {isFreelancer && (
            <TouchableOpacity className="mt-4">
              <Card padding="md">
                <View className="flex-row items-center justify-between">
                  <View className="flex-row items-center">
                    <Eye color="#6366f1" size={20} />
                    <Text className="ml-3 text-primary-600 font-medium">View Public Profile</Text>
                  </View>
                  <ChevronRight color="#6366f1" size={20} />
                </View>
              </Card>
            </TouchableOpacity>
          )}

          <TouchableOpacity
            onPress={handleLogout}
            className="flex-row items-center justify-center bg-red-50 rounded-lg p-4 mt-6"
          >
            <LogOut color="#dc2626" size={20} />
            <Text className="ml-2 text-red-600 font-medium">{t('profile.signOut')}</Text>
          </TouchableOpacity>
        </View>

        <View className="h-8" />
      </ScrollView>
    </SafeAreaView>
  );
}
