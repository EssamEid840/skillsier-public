import { View, Text, TouchableOpacity, ScrollView } from 'react-native';
import { DrawerContentScrollView, DrawerContentComponentProps } from '@react-navigation/drawer';
import { useTranslation } from 'react-i18next';
import { useAuth } from '@skillsier/shared';
import { Badge } from '@skillsier/ui';
import {
  Home,
  Briefcase,
  Award,
  User,
  Settings,
  HelpCircle,
  LogOut,
  ChevronRight,
} from 'lucide-react-native';

interface NavItem {
  label: string;
  icon: React.ComponentType<any>;
  route: string;
  badge?: string;
}

export function DrawerContent(props: DrawerContentComponentProps) {
  const { t } = useTranslation();
  const { user, logout } = useAuth();

  const navItems: NavItem[] = [
    { label: t('dashboard.title'), icon: Home, route: '/(tabs)/dashboard' },
    { label: t('jobs.title'), icon: Briefcase, route: '/(tabs)/courses' },
    { label: t('skills.title'), icon: Award, route: '/(tabs)/skills' },
    { label: t('profile.title'), icon: User, route: '/(tabs)/profile' },
  ];

  const secondaryItems: NavItem[] = [
    { label: t('profile.settings'), icon: Settings, route: '/settings' },
    { label: t('profile.helpSupport'), icon: HelpCircle, route: '/help' },
  ];

  const navigateTo = (route: string) => {
    props.navigation.navigate(route as never);
  };

  return (
    <DrawerContentScrollView {...props} className="flex-1 bg-white">
      {/* Header */}
      <View className="px-6 py-8 bg-gradient-to-br from-primary-50 to-purple-50">
        <View className="flex-row items-center mb-4">
          <View className="h-16 w-16 rounded-full bg-gradient-to-br from-primary-600 to-purple-600 items-center justify-center">
            <Text className="text-2xl font-bold text-white">
              {user?.firstName?.[0]}{user?.lastName?.[0]}
            </Text>
          </View>
          <View className="ml-4 flex-1">
            <Text className="text-lg font-bold text-gray-900">
              {user?.firstName} {user?.lastName}
            </Text>
            <Text className="text-sm text-gray-600">{user?.email}</Text>
          </View>
        </View>
        
        <View className="flex-row gap-2">
          {user?.emailVerified && (
            <Badge variant="success" size="sm">
              ✓ Verified
            </Badge>
          )}
          <Badge variant="default" size="sm">
            {user?.userType}
          </Badge>
        </View>
      </View>

      {/* Main Navigation */}
      <View className="px-3 py-4">
        <Text className="px-3 text-xs font-semibold text-gray-500 uppercase tracking-wider mb-2">
          {t('common.navigation')}
        </Text>
        {navItems.map((item, index) => {
          const isActive = props.state.routes[props.state.index].name === item.route;
          return (
            <TouchableOpacity
              key={index}
              onPress={() => navigateTo(item.route)}
              className={`flex-row items-center justify-between px-3 py-3 rounded-lg mb-1 ${
                isActive ? 'bg-primary-50' : 'bg-transparent'
              }`}
            >
              <View className="flex-row items-center flex-1">
                <item.icon
                  color={isActive ? '#6366f1' : '#6b7280'}
                  size={22}
                  strokeWidth={isActive ? 2.5 : 2}
                />
                <Text
                  className={`ml-3 text-base ${
                    isActive
                      ? 'text-primary-700 font-semibold'
                      : 'text-gray-700'
                  }`}
                >
                  {item.label}
                </Text>
              </View>
              {item.badge && (
                <Badge variant="error" size="sm">
                  {item.badge}
                </Badge>
              )}
            </TouchableOpacity>
          );
        })}
      </View>

      {/* Secondary Navigation */}
      <View className="px-3 py-4 border-t border-gray-200">
        {secondaryItems.map((item, index) => (
          <TouchableOpacity
            key={index}
            onPress={() => navigateTo(item.route)}
            className="flex-row items-center justify-between px-3 py-3 rounded-lg mb-1"
          >
            <View className="flex-row items-center flex-1">
              <item.icon color="#6b7280" size={22} />
              <Text className="ml-3 text-base text-gray-700">{item.label}</Text>
            </View>
            <ChevronRight color="#9ca3af" size={18} />
          </TouchableOpacity>
        ))}
      </View>

      {/* Logout */}
      <View className="px-3 py-4 border-t border-gray-200">
        <TouchableOpacity
          onPress={() => {
            props.navigation.closeDrawer();
            logout();
          }}
          className="flex-row items-center px-3 py-3 rounded-lg"
        >
          <LogOut color="#ef4444" size={22} />
          <Text className="ml-3 text-base text-red-600 font-medium">
            {t('auth.logout')}
          </Text>
        </TouchableOpacity>
      </View>

      {/* Footer */}
      <View className="px-6 py-4">
        <Text className="text-xs text-gray-500 text-center">
          Skillsier v1.0.0
        </Text>
      </View>
    </DrawerContentScrollView>
  );
}