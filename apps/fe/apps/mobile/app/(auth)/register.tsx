import { useState } from 'react';
import { View, Text, ScrollView, KeyboardAvoidingView, Platform, TouchableOpacity } from 'react-native';
import { router } from 'expo-router';
import { SafeAreaView } from 'react-native-safe-area-context';
import { useTranslation } from 'react-i18next';
import { Button, Input } from '@skillsier/ui';
import { useRegister } from '@skillsier/shared';
import { Mail, Lock, User as UserIcon, AlertCircle, ArrowLeft, Eye, EyeOff } from 'lucide-react-native';

export default function RegisterScreen() {
  const { t } = useTranslation();
  const { mutate: register, isPending, error } = useRegister();
  const [showPassword, setShowPassword] = useState(false);
  const [showConfirmPassword, setShowConfirmPassword] = useState(false);
  const [formData, setFormData] = useState({
    firstName: '',
    lastName: '',
    email: '',
    password: '',
    confirmPassword: '',
  });
  const [formErrors, setFormErrors] = useState<Record<string, string>>({});

  const validateForm = () => {
    const errors: Record<string, string> = {};

    if (!formData.firstName.trim()) {
      errors.firstName = t('validation.firstNameRequired');
    }

    if (!formData.lastName.trim()) {
      errors.lastName = t('validation.lastNameRequired');
    }

    if (!formData.email.trim()) {
      errors.email = t('validation.emailRequired');
    } else if (!/^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(formData.email)) {
      errors.email = t('validation.emailInvalid');
    }

    if (!formData.password) {
      errors.password = t('validation.passwordRequired');
    } else if (formData.password.length < 8) {
      errors.password = t('validation.passwordTooShort');
    }

    if (formData.password !== formData.confirmPassword) {
      errors.confirmPassword = t('validation.passwordsDoNotMatch');
    }

    setFormErrors(errors);
    return Object.keys(errors).length === 0;
  };

  const handleRegister = () => {
    if (!validateForm()) return;

    register(
      {
        firstName: formData.firstName,
        lastName: formData.lastName,
        email: formData.email,
        password: formData.password,
      },
      {
        onSuccess: () => {
          router.replace('/(tabs)/dashboard');
        },
      }
    );
  };

  const getPasswordStrength = () => {
    const password = formData.password;
    if (password.length === 0) return { strength: 0, label: '', color: '' };
    if (password.length < 6) return { strength: 25, label: t('validation.weak'), color: 'bg-red-500' };
    if (password.length < 10) return { strength: 50, label: t('validation.medium'), color: 'bg-yellow-500' };
    if (password.length < 12) return { strength: 75, label: t('validation.strong'), color: 'bg-green-500' };
    return { strength: 100, label: t('validation.veryStrong'), color: 'bg-green-600' };
  };

  const passwordStrength = getPasswordStrength();

  return (
    <SafeAreaView className="flex-1 bg-white">
      <KeyboardAvoidingView
        behavior={Platform.OS === 'ios' ? 'padding' : 'height'}
        className="flex-1"
      >
        <ScrollView className="flex-1" keyboardShouldPersistTaps="handled">
          {/* Header */}
          <View className="px-6 pt-4 pb-8">
            <TouchableOpacity onPress={() => router.back()} className="mb-6">
              <ArrowLeft color="#6b7280" size={24} />
            </TouchableOpacity>

            <Text className="text-3xl font-bold text-gray-900">
              {t('auth.createAccount')}
            </Text>
            <Text className="text-gray-600 mt-2">
              {t('auth.signUpSubtitle')}
            </Text>
          </View>

          {/* Form */}
          <View className="px-6 space-y-4">
            {/* Error Message */}
            {error && (
              <View className="bg-red-50 border border-red-200 rounded-lg p-4 flex-row items-center">
                <AlertCircle color="#ef4444" size={20} />
                <Text className="text-red-700 text-sm ml-2 flex-1">
                  {error.message || t('errors.generic')}
                </Text>
              </View>
            )}

            {/* First Name */}
            <View>
              <Text className="text-sm font-medium text-gray-700 mb-2">
                {t('profile.firstName')}
              </Text>
              <View className="relative">
                <View className="absolute left-3 top-3 z-10">
                  <UserIcon color="#9ca3af" size={20} />
                </View>
                <Input
                  placeholder={t('profile.firstName')}
                  value={formData.firstName}
                  onChangeText={(text) => setFormData({ ...formData, firstName: text })}
                  className="pl-10"
                  autoCapitalize="words"
                />
              </View>
              {formErrors.firstName && (
                <Text className="text-red-600 text-xs mt-1">{formErrors.firstName}</Text>
              )}
            </View>

            {/* Last Name */}
            <View>
              <Text className="text-sm font-medium text-gray-700 mb-2">
                {t('profile.lastName')}
              </Text>
              <View className="relative">
                <View className="absolute left-3 top-3 z-10">
                  <UserIcon color="#9ca3af" size={20} />
                </View>
                <Input
                  placeholder={t('profile.lastName')}
                  value={formData.lastName}
                  onChangeText={(text) => setFormData({ ...formData, lastName: text })}
                  className="pl-10"
                  autoCapitalize="words"
                />
              </View>
              {formErrors.lastName && (
                <Text className="text-red-600 text-xs mt-1">{formErrors.lastName}</Text>
              )}
            </View>

            {/* Email */}
            <View>
              <Text className="text-sm font-medium text-gray-700 mb-2">
                {t('auth.email')}
              </Text>
              <View className="relative">
                <View className="absolute left-3 top-3 z-10">
                  <Mail color="#9ca3af" size={20} />
                </View>
                <Input
                  placeholder={t('auth.emailPlaceholder')}
                  value={formData.email}
                  onChangeText={(text) => setFormData({ ...formData, email: text })}
                  className="pl-10"
                  keyboardType="email-address"
                  autoCapitalize="none"
                />
              </View>
              {formErrors.email && (
                <Text className="text-red-600 text-xs mt-1">{formErrors.email}</Text>
              )}
            </View>

            {/* Password */}
            <View>
              <Text className="text-sm font-medium text-gray-700 mb-2">
                {t('auth.password')}
              </Text>
              <View className="relative">
                <View className="absolute left-3 top-3 z-10">
                  <Lock color="#9ca3af" size={20} />
                </View>
                <Input
                  placeholder={t('auth.passwordPlaceholder')}
                  value={formData.password}
                  onChangeText={(text) => setFormData({ ...formData, password: text })}
                  className="pl-10 pr-10"
                  secureTextEntry={!showPassword}
                />
                <TouchableOpacity
                  onPress={() => setShowPassword(!showPassword)}
                  className="absolute right-3 top-3"
                >
                  {showPassword ? (
                    <EyeOff color="#9ca3af" size={20} />
                  ) : (
                    <Eye color="#9ca3af" size={20} />
                  )}
                </TouchableOpacity>
              </View>
              {formErrors.password && (
                <Text className="text-red-600 text-xs mt-1">{formErrors.password}</Text>
              )}

              {/* Password Strength */}
              {formData.password.length > 0 && (
                <View className="mt-2">
                  <View className="h-2 bg-gray-200 rounded-full overflow-hidden">
                    <View
                      className={`h-full ${passwordStrength.color}`}
                      style={{ width: `${passwordStrength.strength}%` }}
                    />
                  </View>
                  <Text className="text-xs text-gray-600 mt-1">
                    {passwordStrength.label}
                  </Text>
                </View>
              )}
            </View>

            {/* Confirm Password */}
            <View>
              <Text className="text-sm font-medium text-gray-700 mb-2">
                {t('auth.confirmPassword')}
              </Text>
              <View className="relative">
                <View className="absolute left-3 top-3 z-10">
                  <Lock color="#9ca3af" size={20} />
                </View>
                <Input
                  placeholder={t('auth.confirmPassword')}
                  value={formData.confirmPassword}
                  onChangeText={(text) => setFormData({ ...formData, confirmPassword: text })}
                  className="pl-10 pr-10"
                  secureTextEntry={!showConfirmPassword}
                />
                <TouchableOpacity
                  onPress={() => setShowConfirmPassword(!showConfirmPassword)}
                  className="absolute right-3 top-3"
                >
                  {showConfirmPassword ? (
                    <EyeOff color="#9ca3af" size={20} />
                  ) : (
                    <Eye color="#9ca3af" size={20} />
                  )}
                </TouchableOpacity>
              </View>
              {formErrors.confirmPassword && (
                <Text className="text-red-600 text-xs mt-1">{formErrors.confirmPassword}</Text>
              )}
            </View>

            {/* Register Button */}
            <View className="pt-4">
              <Button
                onPress={handleRegister}
                disabled={isPending}
                size="lg"
                fullWidth
              >
                <Text className="text-white font-semibold">
                  {isPending ? t('common.loading') : t('auth.register')}
                </Text>
              </Button>
            </View>

            {/* Sign In Link */}
            <View className="flex-row justify-center py-4">
              <Text className="text-gray-600">{t('auth.alreadyHaveAccount')} </Text>
              <TouchableOpacity onPress={() => router.push('/(auth)/login')}>
                <Text className="text-primary-600 font-semibold">{t('auth.login')}</Text>
              </TouchableOpacity>
            </View>
          </View>
        </ScrollView>
      </KeyboardAvoidingView>
    </SafeAreaView>
  );
}