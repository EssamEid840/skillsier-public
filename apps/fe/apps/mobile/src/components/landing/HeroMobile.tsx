import { View, Text } from 'react-native';
import { useTranslation } from 'react-i18next';
import { Button } from '@skillsier/ui';
import { router } from 'expo-router';
import { ArrowRight } from 'lucide-react-native';

export function HeroMobile() {
  const { t } = useTranslation();

  return (
    <View className="px-6 py-12">
      <View className="items-center">
        <View className="h-20 w-20 rounded-2xl bg-gradient-to-br from-primary-600 to-purple-600 mb-6" />
        
        <Text className="text-4xl font-bold text-gray-900 text-center">
          {t('landing.hero.title')}
        </Text>
        
        <Text className="text-4xl font-bold text-center bg-gradient-to-r from-primary-600 to-purple-600 bg-clip-text text-transparent mt-2">
          {t('landing.hero.titleHighlight')}
        </Text>
        
        <Text className="text-lg text-gray-600 text-center mt-4 px-4">
          {t('landing.hero.subtitle')}
        </Text>
      </View>

      <View className="mt-10 space-y-4">
        <Button
          onPress={() => router.push('/(auth)/register')}
          size="lg"
          fullWidth
        >
          <View className="flex-row items-center">
            <Text className="text-white font-semibold">
              {t('landing.hero.getStarted')}
            </Text>
            <ArrowRight color="#fff" size={20} style={{ marginLeft: 8 }} />
          </View>
        </Button>
        
        <Button
          onPress={() => router.push('/(auth)/login')}
          variant="outline"
          size="lg"
          fullWidth
        >
          <Text className="text-primary-600 font-semibold">
            {t('auth.login')}
          </Text>
        </Button>
      </View>

      <View className="mt-12">
        <View className="flex-row justify-around px-4">
          <View className="items-center">
            <Text className="text-3xl font-bold text-primary-600">500K+</Text>
            <Text className="text-sm text-gray-600 mt-1">
              {t('landing.hero.stats.learners')}
            </Text>
          </View>
          <View className="items-center">
            <Text className="text-3xl font-bold text-primary-600">2K+</Text>
            <Text className="text-sm text-gray-600 mt-1">
              {t('landing.hero.stats.clients')}
            </Text>
          </View>
          <View className="items-center">
            <Text className="text-3xl font-bold text-primary-600">98%</Text>
            <Text className="text-sm text-gray-600 mt-1">
              {t('landing.hero.stats.satisfaction')}
            </Text>
          </View>
        </View>
      </View>
    </View>
  );
}