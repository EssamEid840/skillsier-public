import { useQuery, useMutation, useQueryClient } from '@tanstack/react-query';
import { userApi } from '../api/userApi';
import { queryKeys } from '../../../lib/api/queryClient';
import type {
  UpdateFreelancerProfileRequest,
  AddSkillRequest,
  UpdateSkillRequest,
  AddWorkExperienceRequest,
  UpdateWorkExperienceRequest,
  AddEducationRequest,
  AddCertificationRequest,
  AddPortfolioItemRequest,
} from '@skillsier/types';

// Get Freelancer Profile
export const useFreelancerProfile = () => {
  return useQuery({
    queryKey: queryKeys.users.freelancerProfile,
    queryFn: userApi.getFreelancerProfile,
  });
};

// Update Freelancer Profile
export const useUpdateFreelancerProfile = () => {
  const queryClient = useQueryClient();
  return useMutation({
    mutationFn: (data: UpdateFreelancerProfileRequest) => userApi.updateFreelancerProfile(data),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: queryKeys.users.freelancerProfile });
      queryClient.invalidateQueries({ queryKey: queryKeys.users.profile });
    },
  });
};

// Skills Management
export const useFreelancerSkills = () => {
  return useQuery({
    queryKey: queryKeys.users.skills,
    queryFn: userApi.getSkills,
  });
};

export const useAddSkill = () => {
  const queryClient = useQueryClient();
  return useMutation({
    mutationFn: (data: AddSkillRequest) => userApi.addSkill(data),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: queryKeys.users.skills });
      queryClient.invalidateQueries({ queryKey: queryKeys.users.freelancerProfile });
    },
  });
};

export const useUpdateSkill = () => {
  const queryClient = useQueryClient();
  return useMutation({
    mutationFn: ({ skillId, data }: { skillId: string; data: UpdateSkillRequest }) =>
      userApi.updateSkill(skillId, data),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: queryKeys.users.skills });
    },
  });
};

export const useDeleteSkill = () => {
  const queryClient = useQueryClient();
  return useMutation({
    mutationFn: (skillId: string) => userApi.deleteSkill(skillId),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: queryKeys.users.skills });
      queryClient.invalidateQueries({ queryKey: queryKeys.users.freelancerProfile });
    },
  });
};

// Work Experience Management
export const useWorkExperience = () => {
  return useQuery({
    queryKey: queryKeys.users.workExperience,
    queryFn: userApi.getWorkExperience,
  });
};

export const useAddWorkExperience = () => {
  const queryClient = useQueryClient();
  return useMutation({
    mutationFn: (data: AddWorkExperienceRequest) => userApi.addWorkExperience(data),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: queryKeys.users.workExperience });
    },
  });
};

export const useUpdateWorkExperience = () => {
  const queryClient = useQueryClient();
  return useMutation({
    mutationFn: ({ id, data }: { id: string; data: UpdateWorkExperienceRequest }) =>
      userApi.updateWorkExperience(id, data),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: queryKeys.users.workExperience });
    },
  });
};

export const useDeleteWorkExperience = () => {
  const queryClient = useQueryClient();
  return useMutation({
    mutationFn: (id: string) => userApi.deleteWorkExperience(id),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: queryKeys.users.workExperience });
    },
  });
};

// Education Management
export const useEducation = () => {
  return useQuery({
    queryKey: queryKeys.users.education,
    queryFn: userApi.getEducation,
  });
};

export const useAddEducation = () => {
  const queryClient = useQueryClient();
  return useMutation({
    mutationFn: (data: AddEducationRequest) => userApi.addEducation(data),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: queryKeys.users.education });
    },
  });
};

// Certifications Management
export const useCertifications = () => {
  return useQuery({
    queryKey: queryKeys.users.certifications,
    queryFn: userApi.getCertifications,
  });
};

export const useAddCertification = () => {
  const queryClient = useQueryClient();
  return useMutation({
    mutationFn: (data: AddCertificationRequest) => userApi.addCertification(data),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: queryKeys.users.certifications });
    },
  });
};

// Portfolio Management
export const usePortfolio = () => {
  return useQuery({
    queryKey: queryKeys.users.portfolio,
    queryFn: userApi.getPortfolio,
  });
};

export const useAddPortfolioItem = () => {
  const queryClient = useQueryClient();
  return useMutation({
    mutationFn: (data: AddPortfolioItemRequest) => userApi.addPortfolioItem(data),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: queryKeys.users.portfolio });
    },
  });
};

export const useUploadPortfolioImage = () => {
  const queryClient = useQueryClient();
  return useMutation({
    mutationFn: ({ itemId, file }: { itemId: string; file: File | Blob }) =>
      userApi.uploadPortfolioImage(itemId, file),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: queryKeys.users.portfolio });
    },
  });
};