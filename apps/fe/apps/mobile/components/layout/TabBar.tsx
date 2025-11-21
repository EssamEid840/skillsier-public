import type { BottomTabBarProps } from '@react-navigation/bottom-tabs';
import { Pressable, Text, View } from 'react-native';

export function TabBar({ state, descriptors, navigation }: BottomTabBarProps) {
  return (
    <View className="flex-row bg-white border-t border-gray-200">
      {state.routes.map((route, index) => {
        const { options } = descriptors[route.key];
        const label = options.title || route.name;
        const isFocused = state.index === index;

        const onPress = () => {
          const event = navigation.emit({
            type: 'tabPress',
            target: route.key,
            canPreventDefault: true,
          });

          if (!isFocused && !event.defaultPrevented) {
            navigation.navigate(route.name);
          }
        };

        return (
          <Pressable
            key={route.key}
            onPress={onPress}
            className={`flex-1 items-center justify-center py-3 ${
              isFocused ? 'bg-primary-50' : ''
            }`}
          >
            <Text
              className={`text-sm font-medium ${
                isFocused ? 'text-primary-500' : 'text-gray-600'
              }`}
            >
              {label}
            </Text>
          </Pressable>
        );
      })}
    </View>
  );
}