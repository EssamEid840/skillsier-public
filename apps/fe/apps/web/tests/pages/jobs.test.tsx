import { JobCard } from '@/components/jobs/JobCard';
import { JobStatus } from '@skillsier/types';
import { render, screen } from '@testing-library/react';
import { describe, expect, it, vi } from 'vitest';

// Mock next/link
vi.mock('next/link', () => ({
  default: ({ children, href }: any) => <a href={href}>{children}</a>,
}));

const mockJob = {
  id: '1',
  title: 'Full-Stack Developer',
  description: 'Build a SaaS platform with Next.js and Node.js',
  budget: 5000,
  status: JobStatus.ACTIVE,
  proposalCount: 5,
  createdAt: new Date('2024-01-15').toISOString(),
  updatedAt: new Date('2024-01-15').toISOString(),
};

describe('JobCard', () => {
  it('renders job title', () => {
    render(<JobCard job={mockJob} />);
    expect(screen.getByText('Full-Stack Developer')).toBeInTheDocument();
  });

  it('renders job description', () => {
    render(<JobCard job={mockJob} />);
    expect(
      screen.getByText('Build a SaaS platform with Next.js and Node.js')
    ).toBeInTheDocument();
  });

  it('renders job budget', () => {
    render(<JobCard job={mockJob} />);
    expect(screen.getByText('$5,000')).toBeInTheDocument();
  });

  it('renders job status', () => {
    render(<JobCard job={mockJob} />);
    expect(screen.getByText('ACTIVE')).toBeInTheDocument();
  });

  it('renders proposal count', () => {
    render(<JobCard job={mockJob} />);
    expect(screen.getByText('5 proposals')).toBeInTheDocument();
  });

  it('renders posted date', () => {
    render(<JobCard job={mockJob} />);
    expect(screen.getByText(/Posted/)).toBeInTheDocument();
  });

  it('links to job detail page', () => {
    render(<JobCard job={mockJob} />);
    const link = screen.getByRole('link');
    expect(link).toHaveAttribute(
      'href',
      '/(authenticated)/(dashboard)/jobs/1'
    );
  });
});