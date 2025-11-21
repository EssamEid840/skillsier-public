import Link from 'next/link';
import { Button } from '@skillsier/ui';

export default function Home() {
  return (
    <main className="min-h-screen">
      <nav className="border-b">
        <div className="container mx-auto flex h-16 items-center justify-between px-4">
          <h1 className="text-2xl font-bold text-primary">Skillsier</h1>
          <div className="flex gap-4">
            <Link href="/login">
              <Button variant="ghost">Login</Button>
            </Link>
            <Link href="/register">
              <Button>Sign Up</Button>
            </Link>
          </div>
        </div>
      </nav>

      <section className="container mx-auto px-4 py-20 text-center">
        <h2 className="mb-4 text-5xl font-bold">
          Find the Perfect Freelancer
        </h2>
        <p className="mb-8 text-xl text-secondary-600">
          Connect with top talent from around the world
        </p>
        <div className="flex justify-center gap-4">
          <Link href="/register">
            <Button size="lg">Get Started</Button>
          </Link>
          <Link href="/jobs">
            <Button variant="outline" size="lg">
              Browse Jobs
            </Button>
          </Link>
        </div>
      </section>

      <section className="bg-secondary-50 py-20">
        <div className="container mx-auto px-4">
          <h3 className="mb-12 text-center text-3xl font-bold">
            How It Works
          </h3>
          <div className="grid gap-8 md:grid-cols-3">
            <div className="text-center">
              <div className="mb-4 text-4xl">📝</div>
              <h4 className="mb-2 text-xl font-semibold">Post a Job</h4>
              <p className="text-secondary-600">
                Describe your project and get proposals
              </p>
            </div>
            <div className="text-center">
              <div className="mb-4 text-4xl">🔍</div>
              <h4 className="mb-2 text-xl font-semibold">Find Talent</h4>
              <p className="text-secondary-600">
                Review proposals and hire the best fit
              </p>
            </div>
            <div className="text-center">
              <div className="mb-4 text-4xl">✅</div>
              <h4 className="mb-2 text-xl font-semibold">Get Work Done</h4>
              <p className="text-secondary-600">
                Collaborate and complete your project
              </p>
            </div>
          </div>
        </div>
      </section>
    </main>
  );
}