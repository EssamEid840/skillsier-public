'use client';

import { useState } from 'react';
import { useTranslations } from 'next-intl';
import { usePortfolio, useAddPortfolioItem, useUploadPortfolioImage } from '@skillsier/shared';
import { Card, Button, Input } from '@skillsier/ui';
import { Plus, ExternalLink, Upload, X } from 'lucide-react';

export default function PortfolioPage() {
  const t = useTranslations('profile');
  const { data: portfolio, isLoading } = usePortfolio();
  const addPortfolioItem = useAddPortfolioItem();
  const uploadImage = useUploadPortfolioImage();

  const [showForm, setShowForm] = useState(false);
  const [formData, setFormData] = useState({
    title: '',
    description: '',
    projectUrl: '',
    skills: [] as string[],
    completedAt: new Date().toISOString().split('T')[0],
  });
  const [skillInput, setSkillInput] = useState('');

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    try {
      await addPortfolioItem.mutateAsync(formData);
      setShowForm(false);
      setFormData({
        title: '',
        description: '',
        projectUrl: '',
        skills: [],
        completedAt: new Date().toISOString().split('T')[0],
      });
    } catch (error) {
      console.error('Failed to add portfolio item:', error);
    }
  };

  const handleAddSkill = () => {
    if (skillInput.trim() && !formData.skills.includes(skillInput.trim())) {
      setFormData({
        ...formData,
        skills: [...formData.skills, skillInput.trim()],
      });
      setSkillInput('');
    }
  };

  const handleRemoveSkill = (skill: string) => {
    setFormData({
      ...formData,
      skills: formData.skills.filter(s => s !== skill),
    });
  };

  if (isLoading) {
    return (
      <div className="flex items-center justify-center h-96">
        <div className="h-8 w-8 animate-spin rounded-full border-4 border-primary-600 border-t-transparent" />
      </div>
    );
  }

  return (
    <div className="space-y-6">
      <div className="flex items-center justify-between">
        <div>
          <h1 className="text-3xl font-bold text-gray-900">{t('portfolio')}</h1>
          <p className="text-gray-600 mt-1">Showcase your best work</p>
        </div>
        <Button onClick={() => setShowForm(true)}>
          <Plus className="h-5 w-5 mr-2" />
          Add Project
        </Button>
      </div>

      {showForm && (
        <Card padding="lg">
          <h2 className="text-xl font-semibold text-gray-900 mb-4">Add Portfolio Item</h2>
          <form onSubmit={handleSubmit} className="space-y-4">
            <Input
              label="Project Title"
              value={formData.title}
              onChange={(e) => setFormData({ ...formData, title: e.target.value })}
              placeholder="My Awesome Project"
              required
            />

            <div>
              <label className="block text-sm font-medium text-gray-700 mb-2">
                Description
              </label>
              <textarea
                value={formData.description}
                onChange={(e) => setFormData({ ...formData, description: e.target.value })}
                rows={4}
                className="w-full px-4 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-primary-500"
                placeholder="Describe your project..."
                required
              />
            </div>

            <Input
              label="Project URL (Optional)"
              value={formData.projectUrl}
              onChange={(e) => setFormData({ ...formData, projectUrl: e.target.value })}
              placeholder="https://example.com"
              type="url"
            />

            <div>
              <label className="block text-sm font-medium text-gray-700 mb-2">
                Skills Used
              </label>
              <div className="flex gap-2 mb-2">
                <Input
                  value={skillInput}
                  onChange={(e) => setSkillInput(e.target.value)}
                  placeholder="Add a skill"
                  onKeyPress={(e) => {
                    if (e.key === 'Enter') {
                      e.preventDefault();
                      handleAddSkill();
                    }
                  }}
                />
                <Button type="button" onClick={handleAddSkill}>
                  Add
                </Button>
              </div>
              <div className="flex flex-wrap gap-2">
                {formData.skills.map((skill) => (
                  <span
                    key={skill}
                    className="inline-flex items-center gap-1 px-3 py-1 bg-primary-50 text-primary-700 rounded-full text-sm"
                  >
                    {skill}
                    <button
                      type="button"
                      onClick={() => handleRemoveSkill(skill)}
                      className="hover:text-primary-900"
                    >
                      <X className="h-3 w-3" />
                    </button>
                  </span>
                ))}
              </div>
            </div>

            <Input
              label="Completion Date"
              type="date"
              value={formData.completedAt}
              onChange={(e) => setFormData({ ...formData, completedAt: e.target.value })}
              required
            />

            <div className="flex gap-3">
              <Button type="submit" loading={addPortfolioItem.isPending}>
                Add Project
              </Button>
              <Button
                type="button"
                variant="outline"
                onClick={() => setShowForm(false)}
              >
                Cancel
              </Button>
            </div>
          </form>
        </Card>
      )}

      <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
        {portfolio?.map((item) => (
          <Card key={item.id} padding="lg" hoverable>
            <div className="relative aspect-video bg-gray-100 rounded-lg overflow-hidden mb-4">
              {item.images[0] ? (
                <img
                  src={item.images[0]}
                  alt={item.title}
                  className="w-full h-full object-cover"
                />
              ) : (
                <div className="w-full h-full flex items-center justify-center text-gray-400">
                  No Image
                </div>
              )}
            </div>

            <h3 className="font-semibold text-gray-900 mb-2">{item.title}</h3>
            <p className="text-sm text-gray-600 mb-3 line-clamp-2">{item.description}</p>

            {item.projectUrl && (
              <a
                href={item.projectUrl}
                target="_blank"
                rel="noopener noreferrer"
                className="inline-flex items-center gap-1 text-sm text-primary-600 hover:text-primary-700 mb-3"
              >
                View Project <ExternalLink className="h-3 w-3" />
              </a>
            )}

            <div className="flex flex-wrap gap-1">
              {item.skills.slice(0, 3).map((skill, idx) => (
                <span
                  key={idx}
                  className="px-2 py-0.5 bg-gray-100 text-gray-600 text-xs rounded"
                >
                  {skill}
                </span>
              ))}
              {item.skills.length > 3 && (
                <span className="px-2 py-0.5 bg-gray-100 text-gray-600 text-xs rounded">
                  +{item.skills.length - 3}
                </span>
              )}
            </div>
          </Card>
        ))}
      </div>

      {portfolio?.length === 0 && !showForm && (
        <Card padding="lg" className="text-center py-12">
          <p className="text-gray-600 mb-4">No portfolio items yet</p>
          <Button onClick={() => setShowForm(true)}>
            <Plus className="h-5 w-5 mr-2" />
            Add Your First Project
          </Button>
        </Card>
      )}
    </div>
  );
}