import React, { useState } from 'react';
import { View, Text, TouchableOpacity, Modal, StyleSheet } from 'react-native';
import { useTranslation } from 'react-i18next';
import { Globe, Check } from 'lucide-react-native';
import { setLanguage } from '@/lib/i18n';

const languages = [
  { code: 'en', name: 'English', nativeName: 'English', flag: '🇺🇸' },
  { code: 'ar', name: 'Arabic', nativeName: 'العربية', flag: '🇸🇦' },
];

export function LanguageSwitcher() {
  const { i18n } = useTranslation();
  const [modalVisible, setModalVisible] = useState(false);

  const handleLanguageChange = async (languageCode: string) => {
    await setLanguage(languageCode);
    setModalVisible(false);
  };

  const currentLanguage = languages.find((lang) => lang.code === i18n.language);

  return (
    <>
      <TouchableOpacity
        onPress={() => setModalVisible(true)}
        className="flex-row items-center p-4 border-b border-gray-200"
      >
        <Globe color="#6b7280" size={20} />
        <Text className="ml-3 flex-1 text-gray-900">{currentLanguage?.nativeName}</Text>
        <Text className="text-2xl">{currentLanguage?.flag}</Text>
      </TouchableOpacity>

      <Modal
        animationType="slide"
        transparent={true}
        visible={modalVisible}
        onRequestClose={() => setModalVisible(false)}
      >
        <View className="flex-1 justify-end bg-black/50">
          <View className="bg-white rounded-t-3xl p-6">
            <Text className="text-xl font-bold text-gray-900 mb-4">
              Select Language
            </Text>
            {languages.map((language) => (
              <TouchableOpacity
                key={language.code}
                onPress={() => handleLanguageChange(language.code)}
                className="flex-row items-center py-4 border-b border-gray-100"
              >
                <Text className="text-2xl mr-3">{language.flag}</Text>
                <View className="flex-1">
                  <Text className="text-base font-medium text-gray-900">
                    {language.nativeName}
                  </Text>
                  <Text className="text-sm text-gray-600">{language.name}</Text>
                </View>
                {i18n.language === language.code && (
                  <Check color="#6366f1" size={20} />
                )}
              </TouchableOpacity>
            ))}
            <TouchableOpacity
              onPress={() => setModalVisible(false)}
              className="mt-4 bg-gray-100 rounded-lg p-4"
            >
              <Text className="text-center text-gray-900 font-medium">Close</Text>
            </TouchableOpacity>
          </View>
        </View>
      </Modal>
    </>
  );
}
