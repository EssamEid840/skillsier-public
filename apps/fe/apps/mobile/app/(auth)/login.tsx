// apps/mobile/app/(auth)/login.tsx
// Mobile login screen with Google SSO

import { useState } from 'react';
import {
  View,
  Text,
  TextInput,
  TouchableOpacity,
  KeyboardAvoidingView,
  Platform,
  Alert,
  ActivityIndicator,
} from 'react-native';
import { useRouter } from 'expo-router';
import { loginWithProvider } from '@/lib/keycloak-mobile';

export default function LoginScreen() {
  const router = useRouter();
  const [isLoading, setIsLoading] = useState(false);
  const [formData, setFormData] = useState({
    email: '',
    password: '',
  });

  const handleGoogleLogin = async () => {
    setIsLoading(true);
    try {
      await loginWithProvider('google');
      router.replace('/(tabs)/dashboard');
    } catch (error) {
      Alert.alert(
        'Login Failed',
        error instanceof Error ? error.message : 'Please try again'
      );
    } finally {
      setIsLoading(false);
    }
  };

  const handleEmailLogin = async () => {
    if (!formData.email || !formData.password) {
      Alert.alert('Error', 'Please enter email and password');
      return;
    }

    setIsLoading(true);
    try {
      await loginWithProvider('local');
      router.replace('/(tabs)/dashboard');
    } catch (error) {
      Alert.alert(
        'Login Failed',
        error instanceof Error ? error.message : 'Please try again'
      );
    } finally {
      setIsLoading(false);
    }
  };

  return (
    <KeyboardAvoidingView
      behavior={Platform.OS === 'ios' ? 'padding' : 'height'}
      className="flex-1 bg-white"
    >
      <View className="flex-1 justify-center px-6">
        {/* Header */}
        <View className="mb-8">
          <Text className="text-3xl font-bold text-gray-900 mb-2">
            Welcome Back
          </Text>
          <Text className="text-gray-600">
            Sign in to continue to Skillsier
          </Text>
        </View>

        {/* Google Sign In Button */}
        <TouchableOpacity
          onPress={handleGoogleLogin}
          disabled={isLoading}
          className="flex-row items-center justify-center bg-white border-2 border-gray-300 rounded-xl py-4 mb-6 active:bg-gray-50"
        >
          {isLoading ? (
            <ActivityIndicator color="#3B82F6" />
          ) : (
            <>
              <View className="w-6 h-6 mr-3">
                <Text>🔷</Text>
              </View>
              <Text className="text-gray-700 font-semibold text-base">
                Continue with Google
              </Text>
            </>
          )}
        </TouchableOpacity>

        {/* Divider */}
        <View className="flex-row items-center mb-6">
          <View className="flex-1 h-px bg-gray-300" />
          <Text className="mx-4 text-gray-500">or</Text>
          <View className="flex-1 h-px bg-gray-300" />
        </View>

        {/* Email Input */}
        <View className="mb-4">
          <Text className="text-sm font-medium text-gray-700 mb-2">
            Email
          </Text>
          <TextInput
            value={formData.email}
            onChangeText={(text) =>
              setFormData((prev) => ({ ...prev, email: text }))
            }
            placeholder="you@example.com"
            keyboardType="email-address"
            autoCapitalize="none"
            autoComplete="email"
            className="bg-gray-50 border border-gray-300 rounded-xl px-4 py-3 text-gray-900"
          />
        </View>

        {/* Password Input */}
        <View className="mb-6">
          <Text className="text-sm font-medium text-gray-700 mb-2">
            Password
          </Text>
          <TextInput
            value={formData.password}
            onChangeText={(text) =>
              setFormData((prev) => ({ ...prev, password: text }))
            }
            placeholder="••••••••"
            secureTextEntry
            autoComplete="password"
            className="bg-gray-50 border border-gray-300 rounded-xl px-4 py-3 text-gray-900"
          />
        </View>

        {/* Sign In Button */}
        <TouchableOpacity
          onPress={handleEmailLogin}
          disabled={isLoading}
          className="bg-blue-600 rounded-xl py-4 mb-4 active:bg-blue-700"
        >
          <Text className="text-white text-center font-semibold text-base">
            {isLoading ? 'Signing in...' : 'Sign In'}
          </Text>
        </TouchableOpacity>

        {/* Sign Up Link */}
        <TouchableOpacity
          onPress={() => router.push('/(auth)/register')}
          className="py-2"
        >
          <Text className="text-center text-gray-600">
            Don't have an account?{' '}
            <Text className="text-blue-600 font-semibold">Sign Up</Text>
          </Text>
        </TouchableOpacity>
      </View>
    </KeyboardAvoidingView>
  );
}