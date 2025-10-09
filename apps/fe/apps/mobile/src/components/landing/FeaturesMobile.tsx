import { View, Text, ScrollView } from 'react-native';
import { useTranslation } from 'react-i18next';
import { Card } from '@skillsier/ui';
import {
  Briefcase,
  DollarSign,
  Globe,
  Award,
  Clock,
  Shield,
} from 'lucide-react-native';

const features = [
  {
    icon: Briefcase,
    titleKey: 'freelancing.findWork',
    descriptionKey: 'Find thousands of jobs from clients worldwide',
    color: '#6366f1',
    bgColor: '#eef2ff',
  },
  {
    icon: DollarSign,
    titleKey: 'freelancing.earnings',
    descriptionKey: 'Competitive rates and secure payment processing',
    color: '#10b981',
    bgColor: '#d1fae5',
  },
  {
    icon: Globe,
    titleKey: 'Global reach',
    descriptionKey: 'Connect with clients from over 180 countries',
    color: '#8b5cf6',
    bgColor: '#f3e8ff',
  },
  {
    icon: Award,
    titleKey: 'Build reputation',
    descriptionKey: 'Earn badges and showcase your portfolio',
    color: '#f59e0b',
    bgColor: '#fef3c7',
  },
  {
    icon: Clock,
    titleKey: 'Flexible schedule',
    descriptionKey: 'Work when you want, from wherever you want',
    color: '#3b82f6',
    bgColor: '#dbeafe',
  },
  {
    icon: Shield,
    titleKey: 'Secure platform',
    descriptionKey: 'Protected payments and dispute resolution',
    color: '#ef4444',
    bgColor: '#fee2e2',
  },
];

export function FeaturesMobile() {
  const { t } = useTranslation();

  return (
    <View className="px-6 py-12 bg-gray-50">
      <Text className="text-3xl font-bold text-gray-900 text-center mb-4">
        Why Choose Skillsier
      </Text>
      <Text className="text-lg text-gray-600 text-center mb-8 px-4">
        Everything you need to succeed as a freelancer
      </Text>

      <View className="space-y-4">
        {features.map((feature, index) => (
          <Card key={index} className="p-6">
            <View className="flex-row items-start">
              <View
                className="p-3 rounded-xl"
                style={{ backgroundColor: feature.bgColor }}
              >
                <feature.icon color={feature.color} size={24} />
              </View>
              <View className="flex-1 ml-4">
                <Text className="text-lg font-semibold text-gray-900 mb-1">
                  {t(feature.titleKey)}
                </Text>
                <Text className="text-sm text-gray-600">
                  {feature.descriptionKey}
                </Text>
              </View>
            </View>
          </Card>
        ))}
      </View>
    </View>
  );
}