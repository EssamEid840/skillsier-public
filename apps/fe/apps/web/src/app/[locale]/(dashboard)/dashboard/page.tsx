'use client';

import { useAuth } from '@skillsier/shared';
import { Card, Badge } from '@skillsier/ui';
import { TrendingUp, Clock, Award, Target } from 'lucide-react';

export default function DashboardPage() {
  const { user } = useAuth();

  const stats = [
    {
      name: 'Courses Enrolled',
      value: '12',
      change: '+2 this month',
      icon: TrendingUp,
      color: 'text-green-600',
      bgColor: 'bg-green-50',
    },
    {
      name: 'Hours Learned',
      value: '47',
      change: '+8 this week',
      icon: Clock,
      color: 'text-blue-600',
      bgColor: 'bg-blue-50',
    },
    {
      name: 'Skills Acquired',
      value: '23',
      change: '+5 this month',
      icon: Award,
      color: 'text-purple-600',
      bgColor: 'bg-purple-50',
    },
    {
      name: 'Completion Rate',
      value: '87%',
      change: '+3% improvement',
      icon: Target,
      color: 'text-orange-600',
      bgColor: 'bg-orange-50',
    },
  ];

  return (
    <div className="space-y-8">
      <div>
        <h1 className="text-3xl font-bold text-gray-900">
          Welcome back, {user?.firstName || user?.username}!
        </h1>
        <p className="mt-2 text-gray-600">
          Here&apos;s what&apos;s happening with your learning journey today.
        </p>
      </div>

      <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
        {stats.map((stat) => (
          <Card key={stat.name} padding="lg" hoverable>
            <div className="flex items-start justify-between">
              <div>
                <p className="text-sm text-gray-600">{stat.name}</p>
                <p className="mt-2 text-3xl font-bold text-gray-900">{stat.value}</p>
                <p className="mt-1 text-sm text-green-600">{stat.change}</p>
              </div>
              <div className={`p-3 rounded-lg ${stat.bgColor}`}>
                <stat.icon className={`h-6 w-6 ${stat.color}`} />
              </div>
            </div>
          </Card>
        ))}
      </div>

      <div className="grid grid-cols-1 lg:grid-cols-3 gap-6">
        <Card padding="lg" className="lg:col-span-2">
          <h2 className="text-xl font-semibold text-gray-900 mb-4">Continue Learning</h2>
          <div className="space-y-4">
            {[1, 2, 3].map((i) => (
              <div key={i} className="flex gap-4 p-4 rounded-lg hover:bg-gray-50 cursor-pointer">
                <div className="w-24 h-16 bg-gradient-to-br from-primary-400 to-purple-600 rounded-lg flex-shrink-0" />
                <div className="flex-1">
                  <h3 className="font-semibold text-gray-900">Advanced React Patterns</h3>
                  <p className="text-sm text-gray-600 mt-1">67% complete • 3 hours left</p>
                  <div className="mt-2 w-full bg-gray-200 rounded-full h-2">
                    <div className="bg-primary-600 h-2 rounded-full" style={{ width: '67%' }} />
                  </div>
                </div>
              </div>
            ))}
          </div>
        </Card>

        <Card padding="lg">
          <h2 className="text-xl font-semibold text-gray-900 mb-4">Recommended</h2>
          <div className="space-y-4">
            {[1, 2, 3].map((i) => (
              <div key={i} className="group cursor-pointer">
                <div className="w-full h-32 bg-gradient-to-br from-blue-400 to-cyan-600 rounded-lg mb-2" />
                <h3 className="font-medium text-gray-900 group-hover:text-primary-600">
                  TypeScript Masterclass
                </h3>
                <div className="flex items-center gap-2 mt-1">
                  <Badge size="sm">Beginner</Badge>
                  <span className="text-sm text-gray-600">8 hours</span>
                </div>
              </div>
            ))}
          </div>
        </Card>
      </div>
    </div>
  );
} className="h-10 w-10 rounded-lg bg-gradient-to-br from-primary-600 to-purple-600" />
          <span className="text-2xl font-bold text-gray-900">Skillsier</span>
        </Link>
        <h1 className="text-2xl font-bold text-gray-900">Welcome back</h1>
        <p className="mt-2 text-gray-600">Sign in to your account to continue</p>
      </div>

      {error && (
        <div className="mb-6 p-4 bg-red-50 border border-red-200 rounded-lg flex items-start gap-3">
          <AlertCircle className="h-5 w-5 text-red-600 flex-shrink-0 mt-0.5" />
          <div>
            <p className="text-sm font-medium text-red-800">Authentication failed</p>
            <p className="text-sm text-red-700 mt-1">
              {error.message || 'Invalid username or password'}
            </p>
          </div>
        </div>
      )}

      <form onSubmit={handleSubmit} className="space-y-6">
        <Input
          label="Email or Username"
          type="text"
          value={formData.username}
          onChange={(e) => setFormData({ ...formData, username: e.target.value })}
          placeholder="Enter your email or username"
          required
          leftIcon={<Mail className="h-5 w-5 text-gray-400" />}
          autoComplete="username"
          autoFocus
        />

        <Input
          label="Password"
          type="password"
          value={formData.password}
          onChange={(e) => setFormData({ ...formData, password: e.target.value })}
          placeholder="Enter your password"
          required
          leftIcon={<Lock className="h-5 w-5 text-gray-400" />}
          autoComplete="current-password"
        />

        <div className="flex items-center justify-between">
          <label className="flex items-center gap-2">
            <input
              type="checkbox"
              className="h-4 w-4 rounded border-gray-300 text-primary-600 focus:ring-primary-500"
            />
            <span className="text-sm text-gray-700">Remember me</span>
          </label>
          <Link
            href="/forgot-password"
            className="text-sm font-medium text-primary-600 hover:text-primary-700"
          >
            Forgot password?
          </Link>
        </div>

        <Button type="submit" fullWidth loading={isPending}>
          Sign In
        </Button>
      </form>

      <div className="mt-6">
        <div className="relative">
          <div className="absolute inset-0 flex items-center">
            <div className="w-full border-t border-gray-300" />
          </div>
          <div className="relative flex justify-center text-sm">
            <span className="px-2 bg-white text-gray-500">Or continue with</span>
          </div>
        </div>

        <div className="mt-6 grid grid-cols-2 gap-3">
          <Button variant="outline" type="button">
            <svg className="h-5 w-5 mr-2" viewBox="0 0 24 24">
              <path
                fill="currentColor"
                d="M22.56 12.25c0-.78-.07-1.53-.2-2.25H12v4.26h5.92c-.26 1.37-1.04 2.53-2.21 3.31v2.77h3.57c2.08-1.92 3.28-4.74 3.28-8.09z"
              />
              <path
                fill="currentColor"
                d="M12 23c2.97 0 5.46-.98 7.28-2.66l-3.57-2.77c-.98.66-2.23 1.06-3.71 1.06-2.86 0-5.29-1.93-6.16-4.53H2.18v2.84C3.99 20.53 7.7 23 12 23z"
              />
              <path
                fill="currentColor"
                d="M5.84 14.09c-.22-.66-.35-1.36-.35-2.09s.13-1.43.35-2.09V7.07H2.18C1.43 8.55 1 10.22 1 12s.43 3.45 1.18 4.93l2.85-2.22.81-.62z"
              />
              <path
                fill="currentColor"
                d="M12 5.38c1.62 0 3.06.56 4.21 1.64l3.15-3.15C17.45 2.09 14.97 1 12 1 7.7 1 3.99 3.47 2.18 7.07l3.66 2.84c.87-2.6 3.3-4.53 6.16-4.53z"
              />
            </svg>
            Google
          </Button>
          <Button variant="outline" type="button">
            <svg className="h-5 w-5 mr-2" fill="currentColor" viewBox="0 0 24 24">
              <path d="M12 0C5.37 0 0 5.37 0 12c0 5.31 3.435 9.795 8.205 11.385.6.105.825-.255.825-.57 0-.285-.015-1.23-.015-2.235-3.015.555-3.795-.735-4.035-1.41-.135-.345-.72-1.41-1.23-1.695-.42-.225-1.02-.78-.015-.795.945-.015 1.62.87 1.845 1.23 1.08 1.815 2.805 1.305 3.495.99.105-.78.42-1.305.765-1.605-2.67-.3-5.46-1.335-5.46-5.925 0-1.305.465-2.385 1.23-3.225-.12-.3-.54-1.53.12-3.18 0 0 1.005-.315 3.3 1.23.96-.27 1.98-.405 3-.405s2.04.135 3 .405c2.295-1.56 3.3-1.23 3.3-1.23.66 1.65.24 2.88.12 3.18.765.84 1.23 1.905 1.23 3.225 0 4.605-2.805 5.625-5.475 5.925.435.375.81 1.095.81 2.22 0 1.605-.015 2.895-.015 3.3 0 .315.225.69.825.57A12.02 12.02 0 0024 12c0-6.63-5.37-12-12-12z" />
            </svg>
            GitHub
          </Button>
        </div>
      </div>

      <p className="mt-8 text-center text-sm text-gray-600">
        Don&apos;t have an account?{' '}
        <Link href="/register" className="font-medium text-primary-600 hover:text-primary-700">
          Sign up for free
        </Link>
      </p>
    </Card>
  );
}
