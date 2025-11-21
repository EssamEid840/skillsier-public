import { useAuth } from '@skillsier/auth';
import { Button, Input } from '@skillsier/ui';
import { useRouter } from 'expo-router';
import { useState } from 'react';
import {
  ActivityIndicator,
  KeyboardAvoidingView,
  Platform,
  ScrollView,
  Text,
  View,
} from 'react-native';

export default function LoginScreen() {
  const router = useRouter();
  const { login, isLoading } = useAuth();
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [error, setError] = useState('');

  const handleLogin = async () => {
    try {
      setError('');
      await login(email, password);
      router.replace('/(authenticated)/(tabs)');
    } catch (err) {
      setError('Invalid email or password');
    }
  };

  return (
    <KeyboardAvoidingView
      behavior={Platform.OS === 'ios' ? 'padding' : 'height'}
      className="flex-1"
    >
      <ScrollView
        contentContainerClassName="flex-1 justify-center px-6"
        keyboardShouldPersistTaps="handled"
      >
        <View className="mb-8">
          <Text className="text-4xl font-bold text-primary-500 mb-2">
            Skillsier
          </Text>
          <Text className="text-lg text-gray-600">
            Sign in to your account
          </Text>
        </View>

        <View className="space-y-4">
          <Input
            placeholder="Email"
            value={email}
            onChangeText={setEmail}
            keyboardType="email-address"
            autoCapitalize="none"
            autoComplete="email"
          />

          <Input
            placeholder="Password"
            value={password}
            onChangeText={setPassword}
            secureTextEntry
            autoComplete="password"
          />

          {error ? (
            <Text className="text-red-500 text-sm">{error}</Text>
          ) : null}

          <Button
            onPress={handleLogin}
            disabled={isLoading || !email || !password}
            className="mt-4"
          >
            {isLoading ? (
              <ActivityIndicator color="#fff" />
            ) : (
              <Text className="text-white font-semibold">Sign In</Text>
            )}
          </Button>
        </View>

        <View className="mt-8">
          <Text className="text-sm text-gray-500 text-center">
            Dev Accounts:
          </Text>
          <Text className="text-xs text-gray-400 text-center mt-2">
            admin@skillsier.dev / admin123
          </Text>
          <Text className="text-xs text-gray-400 text-center">
            client@skillsier.dev / client123
          </Text>
          <Text className="text-xs text-gray-400 text-center">
            freelancer@skillsier.dev / freelancer123
          </Text>
        </View>
      </ScrollView>
    </KeyboardAvoidingView>
  );
}