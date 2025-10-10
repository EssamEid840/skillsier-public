import { useState } from 'react';
import { View, Text, ScrollView, TouchableOpacity, TextInput, RefreshControl } from 'react-native';
import { SafeAreaView } from 'react-native-safe-area-context';
import { useTranslation } from 'react-i18next';
import { Card } from '@skillsier/ui';
import { Search, Filter, Briefcase, DollarSign, Clock, MapPin } from 'lucide-react-native';

// Mock data - replace with actual API call
const mockJobs = [
  {
    id: '1',
    title: 'Full Stack Developer Needed',
    description: 'Looking for an experienced developer to build a modern web application...',
    budget: 5000,
    budgetType: 'FIXED_PRICE',
    duration: '3 months',
    skills: ['React', 'Node.js', 'PostgreSQL'],
    clientName: 'Tech Startup Inc',
    location: 'Remote',
    postedAt: '2 hours ago',
    proposals: 15,
  },
  {
    id: '2',
    title: 'Mobile App UI/UX Design',
    description: 'Need a talented designer to create modern, user-friendly mobile app designs...',
    budget: 80,
    budgetType: 'HOURLY',
    duration: '2 weeks',
    skills: ['Figma', 'UI/UX', 'Mobile Design'],
    clientName: 'Digital Agency',
    location: 'Remote',
    postedAt: '5 hours ago',
    proposals: 8,
  },
  {
    id: '3',
    title: 'Content Writer for Tech Blog',
    description: 'Looking for a skilled content writer with technical background...',
    budget: 50,
    budgetType: 'HOURLY',
    duration: 'Ongoing',
    skills: ['Content Writing', 'Technical Writing', 'SEO'],
    clientName: 'Tech Media Co',
    location: 'Remote',
    postedAt: '1 day ago',
    proposals: 23,
  },
];

export default function JobsScreen() {
  const { t } = useTranslation();
  const [searchQuery, setSearchQuery] = useState('');
  const [refreshing, setRefreshing] = useState(false);

  const onRefresh = () => {
    setRefreshing(true);
    // Simulate API call
    setTimeout(() => setRefreshing(false), 1500);
  };

  const filteredJobs = mockJobs.filter(job =>
    job.title.toLowerCase().includes(searchQuery.toLowerCase()) ||
    job.description.toLowerCase().includes(searchQuery.toLowerCase())
  );

  return (
    <SafeAreaView className="flex-1 bg-gray-50">
      {/* Header */}
      <View className="bg-white border-b border-gray-200 px-6 py-4">
        <Text className="text-2xl font-bold text-gray-900 mb-4">
          {t('jobs.findWork')}
        </Text>
        
        {/* Search Bar */}
        <View className="flex-row gap-3">
          <View className="flex-1 flex-row items-center bg-gray-100 rounded-lg px-4 py-3">
            <Search color="#9ca3af" size={20} />
            <TextInput
              className="flex-1 ml-2 text-gray-900"
              placeholder={t('common.search')}
              value={searchQuery}
              onChangeText={setSearchQuery}
              placeholderTextColor="#9ca3af"
            />
          </View>
          <TouchableOpacity className="bg-primary-600 rounded-lg px-4 py-3 justify-center">
            <Filter color="#fff" size={20} />
          </TouchableOpacity>
        </View>
      </View>

      {/* Jobs List */}
      <ScrollView
        className="flex-1 px-6 py-4"
        refreshControl={
          <RefreshControl refreshing={refreshing} onRefresh={onRefresh} />
        }
      >
        <Text className="text-sm text-gray-600 mb-4">
          {filteredJobs.length} jobs found
        </Text>

        <View className="space-y-4">
          {filteredJobs.map((job) => (
            <TouchableOpacity key={job.id}>
              <Card className="p-6">
                {/* Job Title */}
                <Text className="text-lg font-semibold text-gray-900 mb-2">
                  {job.title}
                </Text>

                {/* Client Info */}
                <View className="flex-row items-center mb-3">
                  <Text className="text-sm text-gray-600">{job.clientName}</Text>
                  <View className="h-1 w-1 rounded-full bg-gray-400 mx-2" />
                  <Text className="text-sm text-gray-500">{job.postedAt}</Text>
                </View>

                {/* Description */}
                <Text className="text-sm text-gray-700 mb-4" numberOfLines={2}>
                  {job.description}
                </Text>

                {/* Budget & Duration */}
                <View className="flex-row items-center gap-4 mb-4">
                  <View className="flex-row items-center">
                    <DollarSign color="#10b981" size={18} />
                    <Text className="text-sm font-semibold text-gray-900 ml-1">
                      {job.budgetType === 'HOURLY' 
                        ? `$${job.budget}/hr`
                        : `$${job.budget}`
                      }
                    </Text>
                  </View>
                  <View className="flex-row items-center">
                    <Clock color="#6b7280" size={18} />
                    <Text className="text-sm text-gray-600 ml-1">
                      {job.duration}
                    </Text>
                  </View>
                  <View className="flex-row items-center">
                    <MapPin color="#6b7280" size={18} />
                    <Text className="text-sm text-gray-600 ml-1">
                      {job.location}
                    </Text>
                  </View>
                </View>

                {/* Skills */}
                <View className="flex-row flex-wrap gap-2 mb-3">
                  {job.skills.map((skill, index) => (
                    <View key={index} className="bg-primary-50 px-3 py-1 rounded-full">
                      <Text className="text-xs text-primary-700">{skill}</Text>
                    </View>
                  ))}
                </View>

                {/* Proposals */}
                <View className="flex-row items-center pt-3 border-t border-gray-200">
                  <Briefcase color="#9ca3af" size={16} />
                  <Text className="text-sm text-gray-600 ml-1">
                    {job.proposals} proposals
                  </Text>
                </View>
              </Card>
            </TouchableOpacity>
          ))}
        </View>
      </ScrollView>
    </SafeAreaView>
  );
}