import { Tabs } from 'expo-router';
import { useTranslation } from 'react-i18next';
import { Home, Briefcase, Award, User } from 'lucide-react-native';
import { TabBar } from '../../src/components/navigation/TabBar';

export default function TabsLayout() {
  const { t } = useTranslation();

  return (
    <Tabs
      tabBar={(props) => <TabBar {...props} />}
      screenOptions={{
        headerShown: false,
      }}
    >
      <Tabs.Screen
        name="dashboard"
        options={{
          title: t('dashboard.title'),
          tabBarLabel: t('dashboard.title'),
        }}
      />
      <Tabs.Screen
        name="courses"
        options={{
          title: t('jobs.title'),
          tabBarLabel: t('jobs.title'),
        }}
      />
      <Tabs.Screen
        name="skills"
        options={{
          title: t('skills.title'),
          tabBarLabel: t('skills.title'),
        }}
      />
      <Tabs.Screen
        name="profile"
        options={{
          title: t('profile.title'),
          tabBarLabel: t('profile.title'),
        }}
      />
    </Tabs>
  );
}