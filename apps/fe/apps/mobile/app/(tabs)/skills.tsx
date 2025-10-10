// apps/mobile/app/(tabs)/skills.tsx
import { View, Text, ScrollView, TouchableOpacity } from 'react-native';
import { useState } from 'react';

export default function SkillsScreen() {
  const [skills] = useState([
    { id: 1, name: 'React Native', level: 90, projects: 15 },
    { id: 2, name: 'TypeScript', level: 85, projects: 20 },
    { id: 3, name: 'Node.js', level: 80, projects: 12 },
    { id: 4, name: 'UI/UX Design', level: 75, projects: 8 },
    { id: 5, name: 'GraphQL', level: 70, projects: 6 },
  ]);

  const getLevelColor = (level: number) => {
    if (level >= 80) return 'bg-green-500';
    if (level >= 60) return 'bg-blue-500';
    return 'bg-yellow-500';
  };

  const getLevelText = (level: number) => {
    if (level >= 80) return 'Expert';
    if (level >= 60) return 'Intermediate';
    return 'Beginner';
  };

  return (
    <View className="flex-1 bg-gray-50">
      {/* Header */}
      <View className="bg-blue-600 pt-12 pb-6 px-6">
        <Text className="text-white text-2xl font-bold mb-2">My Skills</Text>
        <Text className="text-blue-100">Manage your expertise</Text>
      </View>

      <ScrollView className="flex-1 px-4 pt-4">
        {/* Add Skill Button */}
        <TouchableOpacity className="bg-blue-600 rounded-xl p-4 mb-4 active:bg-blue-700">
          <Text className="text-white text-center font-semibold">
            + Add New Skill
          </Text>
        </TouchableOpacity>

        {/* Skills List */}
        {skills.map((skill) => (
          <View
            key={skill.id}
            className="bg-white rounded-xl p-4 mb-3 shadow-sm"
          >
            <View className="flex-row justify-between items-start mb-3">
              <View className="flex-1">
                <Text className="text-lg font-bold text-gray-900 mb-1">
                  {skill.name}
                </Text>
                <Text className="text-gray-600 text-sm">
                  {skill.projects} projects completed
                </Text>
              </View>
              <View className="bg-green-50 rounded-full px-3 py-1">
                <Text className="text-green-600 text-xs font-medium">
                  {getLevelText(skill.level)}
                </Text>
              </View>
            </View>

            {/* Progress Bar */}
            <View className="mb-3">
              <View className="flex-row justify-between mb-1">
                <Text className="text-gray-600 text-sm">Proficiency</Text>
                <Text className="text-gray-900 text-sm font-medium">
                  {skill.level}%
                </Text>
              </View>
              <View className="h-2 bg-gray-200 rounded-full overflow-hidden">
                <View
                  className={`h-full ${getLevelColor(skill.level)}`}
                  style={{ width: `${skill.level}%` }}
                />
              </View>
            </View>

            {/* Actions */}
            <View className="flex-row space-x-2">
              <TouchableOpacity className="flex-1 bg-blue-50 rounded-lg py-2 active:bg-blue-100">
                <Text className="text-blue-600 text-center font-medium text-sm">
                  Edit
                </Text>
              </TouchableOpacity>
              <TouchableOpacity className="flex-1 bg-red-50 rounded-lg py-2 active:bg-red-100">
                <Text className="text-red-600 text-center font-medium text-sm">
                  Remove
                </Text>
              </TouchableOpacity>
            </View>
          </View>
        ))}

        {/* Skill Statistics */}
        <View className="bg-white rounded-xl p-4 mb-6 shadow-sm">
          <Text className="text-lg font-bold text-gray-900 mb-4">
            Skill Statistics
          </Text>
          
          <View className="flex-row justify-between mb-3">
            <Text className="text-gray-600">Total Skills</Text>
            <Text className="text-gray-900 font-bold">{skills.length}</Text>
          </View>

          <View className="flex-row justify-between mb-3">
            <Text className="text-gray-600">Expert Level</Text>
            <Text className="text-gray-900 font-bold">
              {skills.filter(s => s.level >= 80).length}
            </Text>
          </View>

          <View className="flex-row justify-between mb-3">
            <Text className="text-gray-600">Total Projects</Text>
            <Text className="text-gray-900 font-bold">
              {skills.reduce((sum, s) => sum + s.projects, 0)}
            </Text>
          </View>

          <View className="flex-row justify-between">
            <Text className="text-gray-600">Avg. Proficiency</Text>
            <Text className="text-gray-900 font-bold">
              {Math.round(skills.reduce((sum, s) => sum + s.level, 0) / skills.length)}%
            </Text>
          </View>
        </View>
      </ScrollView>
    </View>
  );
}