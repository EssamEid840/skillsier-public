import Link from 'next/link';
import { Button } from '@skillsier/ui';
import { ArrowRight } from 'lucide-react';

export function CTA() {
  return (
    <section className="py-24 bg-gradient-to-r from-primary-600 to-purple-600">
      <div className="container mx-auto px-4 sm:px-6 lg:px-8">
        <div className="text-center max-w-3xl mx-auto">
          <h2 className="text-4xl font-bold text-white sm:text-5xl">
            Ready to transform your workforce?
          </h2>
          <p className="mt-4 text-lg text-primary-100">
            Join thousands of companies already using Skillsier to build a future-ready team
          </p>
          <div className="mt-10 flex flex-col sm:flex-row gap-4 justify-center">
            <Link href="/register">
              <Button size="lg" variant="secondary" className="w-full sm:w-auto">
                Start Free Trial
                <ArrowRight className="ml-2 h-5 w-5" />
              </Button>
            </Link>
            <Link href="/demo">
              <Button size="lg" variant="outline" className="w-full sm:w-auto border-white text-white hover:bg-white hover:text-primary-600">
                Schedule Demo
              </Button>
            </Link>
          </div>
          <p className="mt-6 text-sm text-primary-100">
            No credit card required • 14-day free trial • Cancel anytime
          </p>
        </div>
      </div>
    </section>
  );
}