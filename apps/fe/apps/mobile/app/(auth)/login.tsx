import { useState } from 'react';
import { View, Text, ScrollView, KeyboardAvoidingView, Platform, TouchableOpacity } from 'react-native';
import { router } from 'expo-router';
import { SafeAreaView } from 'react-native-safe-area-context';
import { useTranslation } from 'react-i18next';
import { Button, Input } from '@skillsier/ui';
import { useLogin } from '@skillsier/shared';
import { Mail, Lock, AlertCircle, ArrowLeft } from 'lucide-react-native';

export default function LoginScreen() {
  const { t } = useTranslation();
  const { mutate: login, isPending, error } = useLogin();
  const [formData, setFormData] = useState({
    username: '',
    password: '',
  });

  const handleLogin = () => {
    login(formData, {
      onSuccess: () => {
        router.replace('/(tabs)/dashboard');
      },
    });
  };

  return (
    <SafeAreaView className="flex-1 bg-white">
      <KeyboardAvoidingView
        behavior={Platform.OS === 'ios' ? 'padding' : 'height'}
        className="flex-1"
      >
        <ScrollView className="flex-1 px-6" showsVerticalScrollIndicator={false}>
          <TouchableOpacity onPress={() => router.back()} className="py-4">
            <ArrowLeft color="#6b7280" size={24} />
          </TouchableOpacity>

          <View className="items-center mb-8">
            <View className="h-16 w-16 rounded-2xl bg-gradient-to-br from-primary-600 to-purple-600 mb-4" />
            <Text className="text-3xl font-bold text-gray-900">{t('auth.welcomeBack')}</Text>
            <Text className="text-base text-gray-600 mt-2">{t('auth.signInToContinue')}</Text>
          </View>

          {error && (
            <View className="mb-6 p-4 bg-red-50 border border-red-200 rounded-lg flex-row items-start">
              <AlertCircle color="#dc2626" size={20} />
              <View className="ml-3 flex-1">
                <Text className="text-sm font-medium text-red-800">
                  {t('errors.authentication')}
                </Text>
                <Text className="text-sm text-red-700 mt-1">
                  {error.message || t('errors.authentication')}
                </Text>
              </View>
            </View>
          )}

          <View className="space-y-6 mb-6">
            <Input
              label={t('auth.email')}
              value={formData.username}
              onChangeText={(text) => setFormData({ ...formData, username: text })}
              placeholder={t('auth.email')}
              autoCapitalize="none"
              autoComplete="username"
              leftIcon={<Mail color="#9ca3af" size={20} />}
            />

            <Input
              label={t('auth.password')}
              type="password"
              value={formData.password}
              onChangeText={(text) => setFormData({ ...formData, password: text })}
              placeholder={t('auth.password')}
              autoComplete="current-password"
              leftIcon={<Lock color="#9ca3af" size={20} />}
            />
          </View>

          <TouchableOpacity className="mb-6">
            <Text className="text-sm font-medium text-primary-600 text-right">
              {t('auth.forgotPassword')}
            </Text>
          </TouchableOpacity>

          <Button onPress={handleLogin} loading={isPending} fullWidth size="lg">
            {t('auth.login')}
          </Button>

          <View className="flex-row justify-center items-center mt-8">
            <Text className="text-gray-600">{t('auth.dontHaveAccount')} </Text>
            <TouchableOpacity onPress={() => router.push('/(auth)/register')}>
              <Text className="font-medium text-primary-600">{t('auth.register')}</Text>
            </TouchableOpacity>
          </View>
        </ScrollView>
      </KeyboardAvoidingView>
    </SafeAreaView>
  );
}
