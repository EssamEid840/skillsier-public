'use client';

import { useState } from 'react';
import { useTranslations } from 'next-intl';
import { useFreelancerSkills, useAddSkill, useUpdateSkill, useDeleteSkill } from '@skillsier/shared';
import { Card, Button, Input, Badge } from '@skillsier/ui';
import { Plus, Edit, Trash2, X } from 'lucide-react';

export default function SkillsPage() {
  const t = useTranslations('skills');
  const { data: skills, isLoading } = useFreelancerSkills();
  const addSkill = useAddSkill();
  const updateSkill = useUpdateSkill();
  const deleteSkill = useDeleteSkill();

  const [showForm, setShowForm] = useState(false);
  const [editingSkill, setEditingSkill] = useState<string | null>(null);
  const [formData, setFormData] = useState({
    skillId: '',
    name: '',
    level: 'INTERMEDIATE' as 'BEGINNER' | 'INTERMEDIATE' | 'ADVANCED' | 'EXPERT',
    yearsOfExperience: 1,
  });

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    try {
      if (editingSkill) {
        await updateSkill.mutateAsync({
          skillId: editingSkill,
          data: {
            level: formData.level,
            yearsOfExperience: formData.yearsOfExperience,
          },
        });
      } else {
        await addSkill.mutateAsync(formData);
      }
      resetForm();
    } catch (error) {
      console.error('Failed to save skill:', error);
    }
  };

  const resetForm = () => {
    setShowForm(false);
    setEditingSkill(null);
    setFormData({
      skillId: '',
      name: '',
      level: 'INTERMEDIATE',
      yearsOfExperience: 1,
    });
  };

  const handleEdit = (skill: any) => {
    setEditingSkill(skill.id);
    setFormData({
      skillId: skill.skillId,
      name: skill.name,
      level: skill.level,
      yearsOfExperience: skill.yearsOfExperience,
    });
    setShowForm(true);
  };

  const handleDelete = async (skillId: string) => {
    if (confirm('Are you sure you want to delete this skill?')) {
      await deleteSkill.mutateAsync(skillId);
    }
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
          <h1 className="text-3xl font-bold text-gray-900">{t('title')}</h1>
          <p className="text-gray-600 mt-1">Manage your professional skills</p>
        </div>
        <Button onClick={() => setShowForm(true)}>
          <Plus className="h-5 w-5 mr-2" />
          {t('addNew')}
        </Button>
      </div>

      {showForm && (
        <Card padding="lg">
          <div className="flex items-center justify-between mb-4">
            <h2 className="text-xl font-semibold text-gray-900">
              {editingSkill ? 'Edit Skill' : t('addNew')}
            </h2>
            <button onClick={resetForm}>
              <X className="h-5 w-5 text-gray-500" />
            </button>
          </div>

          <form onSubmit={handleSubmit} className="space-y-4">
            {!editingSkill && (
              <Input
                label={t('skillName')}
                value={formData.name}
                onChange={(e) => setFormData({ ...formData, name: e.target.value })}
                placeholder="e.g., React, Python, UI/UX Design"
                required
              />
            )}

            <div>
              <label className="block text-sm font-medium text-gray-700 mb-2">
                {t('level')}
              </label>
              <div className="grid grid-cols-4 gap-2">
                {['BEGINNER', 'INTERMEDIATE', 'ADVANCED', 'EXPERT'].map((level) => (
                  <button
                    key={level}
                    type="button"
                    onClick={() => setFormData({ ...formData, level: level as any })}
                    className={`px-4 py-2 rounded-lg border-2 transition-colors ${
                      formData.level === level
                        ? 'border-primary-600 bg-primary-50 text-primary-700'
                        : 'border-gray-300 hover:border-gray-400'
                    }`}
                  >
                    {t(`levels.${level.toLowerCase()}`)}
                  </button>
                ))}
              </div>
            </div>

            <Input
              label={t('yearsOfExperience')}
              type="number"
              min="0"
              max="50"
              value={formData.yearsOfExperience}
              onChange={(e) =>
                setFormData({ ...formData, yearsOfExperience: Number(e.target.value) })
              }
              required
            />

            <div className="flex gap-3">
              <Button
                type="submit"
                loading={addSkill.isPending || updateSkill.isPending}
              >
                {editingSkill ? 'Update Skill' : 'Add Skill'}
              </Button>
              <Button type="button" variant="outline" onClick={resetForm}>
                Cancel
              </Button>
            </div>
          </form>
        </Card>
      )}

      <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
        {skills?.map((skill) => (
          <Card key={skill.id} padding="lg" hoverable>
            <div className="flex items-start justify-between mb-3">
              <div className="flex-1">
                <h3 className="font-semibold text-gray-900 text-lg">{skill.name}</h3>
                <p className="text-sm text-gray-600 mt-1">{skill.category}</p>
              </div>
              <Badge variant="info" size="sm">
                {t(`levels.${skill.level.toLowerCase()}`)}
              </Badge>
            </div>

            <div className="space-y-2 mb-4">
              <div className="flex justify-between text-sm">
                <span className="text-gray-600">Experience:</span>
                <span className="font-medium text-gray-900">
                  {skill.yearsOfExperience} {skill.yearsOfExperience === 1 ? 'year' : 'years'}
                </span>
              </div>
              {skill.endorsements > 0 && (
                <div className="flex justify-between text-sm">
                  <span className="text-gray-600">Endorsements:</span>
                  <span className="font-medium text-primary-600">
                    {skill.endorsements}
                  </span>
                </div>
              )}
            </div>

            <div className="flex gap-2">
              <Button
                variant="outline"
                size="sm"
                fullWidth
                onClick={() => handleEdit(skill)}
              >
                <Edit className="h-4 w-4 mr-1" />
                Edit
              </Button>
              <Button
                variant="outline"
                size="sm"
                fullWidth
                onClick={() => handleDelete(skill.id)}
              >
                <Trash2 className="h-4 w-4 mr-1" />
                Delete
              </Button>
            </div>
          </Card>
        ))}
      </div>

      {skills?.length === 0 && !showForm && (
        <Card padding="lg" className="text-center py-12">
          <p className="text-gray-600 mb-4">No skills added yet</p>
          <Button onClick={() => setShowForm(true)}>
            <Plus className="h-5 w-5 mr-2" />
            Add Your First Skill
          </Button>
        </Card>
      )}
    </div>
  );
}