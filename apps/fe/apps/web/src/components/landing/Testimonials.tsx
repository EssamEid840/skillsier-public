import { Avatar, Card } from '@skillsier/ui';
import { Star } from 'lucide-react';

const testimonials = [
  {
    name: 'Sarah Johnson',
    role: 'VP of Learning & Development',
    company: 'TechCorp Inc.',
    avatar: '/images/avatars/sarah.jpg',
    content: 'Skillsier transformed our L&D program. We saw a 45% increase in skill acquisition and employee engagement skyrocketed.',
    rating: 5,
  },
  {
    name: 'Michael Chen',
    role: 'Chief Technology Officer',
    company: 'Innovation Labs',
    avatar: '/images/avatars/michael.jpg',
    content: 'The AI-powered learning paths are incredible. Our developers are now 3x more productive thanks to targeted upskilling.',
    rating: 5,
  },
  {
    name: 'Emily Rodriguez',
    role: 'HR Director',
    company: 'Global Solutions',
    avatar: '/images/avatars/emily.jpg',
    content: 'Best investment we made this year. The analytics dashboard gives us real insights into our workforce development.',
    rating: 5,
  },
];

export function Testimonials() {
  return (
    <section className="py-24 bg-gradient-to-br from-gray-50 to-primary-50">
      <div className="container mx-auto px-4 sm:px-6 lg:px-8">
        <div className="text-center max-w-3xl mx-auto mb-16">
          <h2 className="text-4xl font-bold text-gray-900 sm:text-5xl">
            Trusted by industry leaders
          </h2>
          <p className="mt-4 text-lg text-gray-600">
            See what our customers have to say about their experience
          </p>
        </div>
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8">
          {testimonials.map((testimonial, index) => (
            <Card key={index} padding="lg" shadow="md">
              <div className="flex gap-1 mb-4">
                {[...Array(testimonial.rating)].map((_, i) => (
                  <Star key={i} className="h-5 w-5 fill-yellow-400 text-yellow-400" />
                ))}
              </div>
              <p className="text-gray-700 mb-6">&quot;{testimonial.content}&quot;</p>
              <div className="flex items-center gap-3">
                <Avatar src={testimonial.avatar} alt={testimonial.name} size="md" />
                <div>
                  <p className="font-semibold text-gray-900">{testimonial.name}</p>
                  <p className="text-sm text-gray-600">{testimonial.role}</p>
                  <p className="text-sm text-gray-500">{testimonial.company}</p>
                </div>
              </div>
            </Card>
          ))}
        </div>
      </div>
    </section>
  );
}
