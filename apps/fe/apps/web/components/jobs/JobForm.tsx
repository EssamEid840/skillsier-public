'use client';

import { Button, Card, Input } from '@skillsier/ui';
import { useCreateJob } from '@skillsier/hooks';
import { zodResolver } from '@hookform/resolvers/zod';
import { useRouter } from 'next/navigation';
import { useForm } from 'react-hook-form';
import { z } from 'zod';

const jobSchema = z.object({
  title: z.string().min(10, 'Title must be at least 10 characters'),
  description: z.string().min(50, 'Description must be at least 50 characters'),
  budget: z.number().min(1, 'Budget must be greater than 0').optional(),
  skills: z.string().optional(),
  duration: z.string().optional(),
});

type JobFormData = z.infer<typeof jobSchema>;

export function JobForm() {
  const router = useRouter();
  const { mutate: createJob, isPending } = useCreateJob();

  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<JobFormData>({
    resolver: zodResolver(jobSchema),
  });

  const onSubmit = (data: JobFormData) => {
    createJob(
      {
        title: data.title,
        description: data.description,
        budget: data.budget,
        skills: data.skills?.split(',').map((s) => s.trim()) || [],
        duration: data.duration,
      },
      {
        onSuccess: (job) => {
          router.push(`/(authenticated)/(dashboard)/jobs/${job.id}`);
        },
      }
    );
  };

  return (
    <Card className="p-6">
      <h2 className="text-2xl font-bold text-gray-900 mb-6">Post a New Job</h2>

      <form onSubmit={handleSubmit(onSubmit)} className="space-y-6">
        {/* Title */}
        <div>
          <label className="block text-sm font-medium text-gray-700 mb-2">
            Job Title <span className="text-red-500">*</span>
          </label>
          <Input
            {...register('title')}
            placeholder="e.g., Full-Stack Developer for SaaS Platform"
            error={errors.title?.message}
          />
        </div>

        {/* Description */}
        <div>
          <label className="block text-sm font-medium text-gray-700 mb-2">
            Description <span className="text-red-500">*</span>
          </label>
          <textarea
            {...register('description')}
            rows={8}
            placeholder="Describe the job requirements, responsibilities, and qualifications..."
            className="w-full rounded-md border-gray-300 shadow-sm focus:border-primary-500 focus:ring-primary-500"
          />
          {errors.description && (
            <p className="mt-1 text-sm text-red-500">{errors.description.message}</p>
          )}
        </div>

        {/* Budget */}
        <div>
          <label className="block text-sm font-medium text-gray-700 mb-2">
            Budget (USD)
          </label>
          <Input
            type="number"
            {...register('budget', { valueAsNumber: true })}
            placeholder="5000"
            error={errors.budget?.message}
          />
        </div>

        {/* Skills */}
        <div>
          <label className="block text-sm font-medium text-gray-700 mb-2">
            Required Skills (comma-separated)
          </label>
          <Input
            {...register('skills')}
            placeholder="React, Node.js, TypeScript, PostgreSQL"
          />
        </div>

        {/* Duration */}
        <div>
          <label className="block text-sm font-medium text-gray-700 mb-2">
            Project Duration
          </label>
          <select
            {...register('duration')}
            className="w-full rounded-md border-gray-300 shadow-sm focus:border-primary-500 focus:ring-primary-500"
          >
            <option value="">Select duration</option>
            <option value="less-than-1-month">Less than 1 month</option>
            <option value="1-3-months">1-3 months</option>
            <option value="3-6-months">3-6 months</option>
            <option value="more-than-6-months">More than 6 months</option>
          </select>
        </div>

        {/* Actions */}
        <div className="flex items-center gap-4 pt-4">
          <Button type="submit" disabled={isPending}>
            {isPending ? 'Posting...' : 'Post Job'}
          </Button>
          <Button
            type="button"
            variant="outline"
            onClick={() => router.back()}
          >
            Cancel
          </Button>
        </div>
      </form>
    </Card>
  );
}