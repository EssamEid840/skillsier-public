import { Brain, Target, TrendingUp, Users, Shield, Zap } from 'lucide-react';
import { Card } from '@skillsier/ui';

const features = [
  {
    icon: Brain,
    title: 'AI-Powered Learning Paths',
    description: 'Personalized curriculum adapts to each learner\'s pace and style for maximum retention.',
    color: 'text-primary-600',
    bgColor: 'bg-primary-50',
  },
  {
    icon: Target,
    title: 'Skills Gap Analysis',
    description: 'Identify organizational skill gaps and create targeted training programs instantly.',
    color: 'text-purple-600',
    bgColor: 'bg-purple-50',
  },
  {
    icon: TrendingUp,
    title: 'Real-Time Analytics',
    description: 'Track progress, engagement, and ROI with comprehensive dashboards and reports.',
    color: 'text-green-600',
    bgColor: 'bg-green-50',
  },
  {
    icon: Users,
    title: 'Team Collaboration',
    description: 'Foster peer learning with discussion forums, group projects, and mentorship programs.',
    color: 'text-blue-600',
    bgColor: 'bg-blue-50',
  },
  {
    icon: Shield,
    title: 'Enterprise Security',
    description: 'Bank-grade encryption, SSO, and compliance with GDPR, CCPA, and SOC 2 standards.',
    color: 'text-red-600',
    bgColor: 'bg-red-50',
  },
  {
    icon: Zap,
    title: 'Lightning Fast',
    description: 'Optimized performance with offline mode and mobile apps for learning anywhere.',
    color: 'text-yellow-600',
    bgColor: 'bg-yellow-50',
  },
];

export function Features() {
  return (
    <section className="py-24 bg-white">
      <div className="container mx-auto px-4 sm:px-6 lg:px-8">
        <div className="text-center max-w-3xl mx-auto mb-16">
          <h2 className="text-4xl font-bold text-gray-900 sm:text-5xl">
            Everything you need to upskill your team
          </h2>
          <p className="mt-4 text-lg text-gray-600">
            Powerful features designed for enterprise-scale learning and development
          </p>
        </div>
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8">
          {features.map((feature, index) => (
            <Card key={index} padding="lg" hoverable className="group">
              <div className={`inline-flex p-3 rounded-lg ${feature.bgColor} mb-4`}>
                <feature.icon className={`h-6 w-6 ${feature.color}`} />
              </div>
              <h3 className="text-xl font-semibold text-gray-900 mb-2">{feature.title}</h3>
              <p className="text-gray-600">{feature.description}</p>
            </Card>
          ))}
        </div>
      </div>
    </section>
  );
}
