import { View, Text } from 'react-native';
import { Link, Stack } from 'expo-router';
import { SafeAreaView } from 'react-native-safe-area-context';
import { useTranslation } from 'react-i18next';
import { Button } from '@skillsier/ui';
import { Home } from 'lucide-react-native';

export default function NotFoundScreen() {
  const { t } = useTranslation();

  return (
    <>
      <Stack.Screen options={{ title: 'Oops!' }} />
      <SafeAreaView className="flex-1 bg-white">
        <View className="flex-1 items-center justify-center px-6">
          <Text className="text-8xl font-bold text-primary-600 mb-4">404</Text>
          <Text className="text-2xl font-bold text-gray-900 text-center mb-2">
            {t('errors.notFound')}
          </Text>
          <Text className="text-gray-600 text-center mb-8 px-4">
            The page you're looking for doesn't exist or has been moved.
          </Text>
          
          <Link href="/(tabs)/dashboard" asChild>
            <Button size="lg">
              <View className="flex-row items-center gap-2">
                <Home color="#fff" size={20} />
                <Text className="text-white font-semibold">
                  {t('common.goHome')}
                </Text>
              </View>
            </Button>
          </Link>
        </View>
      </SafeAreaView>
    </>
  );
}