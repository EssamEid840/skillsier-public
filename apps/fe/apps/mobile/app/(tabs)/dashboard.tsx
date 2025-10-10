// apps/mobile/app/(tabs)/dashboard.tsx
import { View, Text, ScrollView, TouchableOpacity, RefreshControl } from 'react-native';
import { useState } from 'react';
import { useRouter } from 'expo-router';

export default function DashboardScreen() {
  const router = useRouter();
  const [refreshing, setRefreshing] = useState(false);

  const onRefresh = async () => {
    setRefreshing(true);
    // TODO: Refresh data
    setTimeout(() => setRefreshing(false), 1000);
  };

  return (
    <ScrollView
      className="flex-1 bg-gray-50"
      refreshControl={
        <RefreshControl refreshing={refreshing} onRefresh={onRefresh} />
      }
    >
      {/* Header */}
      <View className="bg-blue-600 pt-12 pb-8 px-6">
        <Text className="text-white text-2xl font-bold mb-2">Dashboard</Text>
        <Text className="text-blue-100">Welcome back!</Text>
      </View>

      {/* Stats Cards */}
      <View className="px-4 -mt-6">
        <View className="flex-row flex-wrap justify-between">
          <View className="bg-white rounded-xl p-4 mb-4 w-[48%] shadow-sm">
            <Text className="text-gray-600 text-sm mb-1">Active Jobs</Text>
            <Text className="text-2xl font-bold text-gray-900">12</Text>
            <Text className="text-green-600 text-xs mt-1">+3 this week</Text>
          </View>

          <View className="bg-white rounded-xl p-4 mb-4 w-[48%] shadow-sm">
            <Text className="text-gray-600 text-sm mb-1">Proposals</Text>
            <Text className="text-2xl font-bold text-gray-900">8</Text>
            <Text className="text-blue-600 text-xs mt-1">2 pending</Text>
          </View>

          <View className="bg-white rounded-xl p-4 mb-4 w-[48%] shadow-sm">
            <Text className="text-gray-600 text-sm mb-1">Earnings</Text>
            <Text className="text-2xl font-bold text-gray-900">$4.2k</Text>
            <Text className="text-gray-500 text-xs mt-1">This month</Text>
          </View>

          <View className="bg-white rounded-xl p-4 mb-4 w-[48%] shadow-sm">
            <Text className="text-gray-600 text-sm mb-1">Rating</Text>
            <Text className="text-2xl font-bold text-gray-900">4.9</Text>
            <Text className="text-yellow-500 text-xs mt-1">★★★★★</Text>
          </View>
        </View>
      </View>

      {/* Recent Activity */}
      <View className="px-4 mt-4">
        <Text className="text-lg font-bold text-gray-900 mb-3">Recent Activity</Text>
        
        <View className="bg-white rounded-xl p-4 mb-3 shadow-sm">
          <Text className="font-semibold text-gray-900 mb-1">New job posted</Text>
          <Text className="text-gray-600 text-sm mb-2">
            Build a mobile app for e-commerce
          </Text>
          <Text className="text-gray-400 text-xs">2 hours ago</Text>
        </View>

        <View className="bg-white rounded-xl p-4 mb-3 shadow-sm">
          <Text className="font-semibold text-gray-900 mb-1">Proposal accepted</Text>
          <Text className="text-gray-600 text-sm mb-2">
            Your proposal for React Native app was accepted
          </Text>
          <Text className="text-gray-400 text-xs">5 hours ago</Text>
        </View>

        <View className="bg-white rounded-xl p-4 mb-3 shadow-sm">
          <Text className="font-semibold text-gray-900 mb-1">Payment received</Text>
          <Text className="text-gray-600 text-sm mb-2">
            $500 from Website Redesign project
          </Text>
          <Text className="text-gray-400 text-xs">1 day ago</Text>
        </View>
      </View>

      {/* Quick Actions */}
      <View className="px-4 mt-4 mb-8">
        <Text className="text-lg font-bold text-gray-900 mb-3">Quick Actions</Text>
        
        <TouchableOpacity
          onPress={() => router.push('/(tabs)/courses')}
          className="bg-blue-600 rounded-xl p-4 mb-3 active:bg-blue-700"
        >
          <Text className="text-white font-semibold text-center">
            Browse New Jobs
          </Text>
        </TouchableOpacity>

        <TouchableOpacity
          className="bg-white border border-gray-300 rounded-xl p-4 active:bg-gray-50"
        >
          <Text className="text-gray-700 font-semibold text-center">
            View My Proposals
          </Text>
        </TouchableOpacity>
      </View>
    </ScrollView>
  );
}