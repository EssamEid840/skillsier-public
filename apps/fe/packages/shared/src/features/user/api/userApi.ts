import apiClient from '../../../lib/api/client';
import { API_ENDPOINTS } from '../../../constants/api';
import type {
  User,
  FreelancerProfile,
  ClientProfile,
  UpdateProfileRequest,
  UpdateFreelancerProfileRequest,
  UpdateClientProfileRequest,
  AddSkillRequest,
  UpdateSkillRequest,
  AddWorkExperienceRequest,
  UpdateWorkExperienceRequest,
  AddEducationRequest,
  AddCertificationRequest,
  AddPortfolioItemRequest,
  ChangePasswordRequest,
  FreelancerSkill,
  WorkExperience,
  Education,
  Certification,
  PortfolioItem,
} from '@skillsier/types';

export const userApi = {
  // ========== Basic Profile ==========
  getProfile: async (): Promise<User> => {
    return apiClient.get(API_ENDPOINTS.USERS.PROFILE);
  },

  updateProfile: async (data: UpdateProfileRequest): Promise<User> => {
    return apiClient.patch(API_ENDPOINTS.USERS.UPDATE_PROFILE, data);
  },

  uploadAvatar: async (file: File | Blob): Promise<{ avatarUrl: string }> => {
    const formData = new FormData();
    formData.append('avatar', file);
    return apiClient.post(API_ENDPOINTS.USERS.UPDATE_AVATAR, formData, {
      headers: { 'Content-Type': 'multipart/form-data' },
    });
  },

  deleteAvatar: async (): Promise<void> => {
    return apiClient.delete(API_ENDPOINTS.USERS.DELETE_AVATAR);
  },

  changePassword: async (data: ChangePasswordRequest): Promise<void> => {
    return apiClient.post(API_ENDPOINTS.USERS.CHANGE_PASSWORD, data);
  },

  deleteAccount: async (password: string): Promise<void> => {
    return apiClient.delete(API_ENDPOINTS.USERS.DELETE_ACCOUNT, { data: { password } });
  },

  // ========== Freelancer Profile ==========
  getFreelancerProfile: async (): Promise<FreelancerProfile> => {
    return apiClient.get(API_ENDPOINTS.USERS.FREELANCER_PROFILE);
  },

  updateFreelancerProfile: async (data: UpdateFreelancerProfileRequest): Promise<FreelancerProfile> => {
    return apiClient.patch(API_ENDPOINTS.USERS.UPDATE_FREELANCER_PROFILE, data);
  },

  // ========== Client Profile ==========
  getClientProfile: async (): Promise<ClientProfile> => {
    return apiClient.get(API_ENDPOINTS.USERS.CLIENT_PROFILE);
  },

  updateClientProfile: async (data: UpdateClientProfileRequest): Promise<ClientProfile> => {
    return apiClient.patch(API_ENDPOINTS.USERS.UPDATE_CLIENT_PROFILE, data);
  },

  // ========== Skills ==========
  getSkills: async (): Promise<FreelancerSkill[]> => {
    return apiClient.get(API_ENDPOINTS.USERS.SKILLS);
  },

  addSkill: async (data: AddSkillRequest): Promise<FreelancerSkill> => {
    return apiClient.post(API_ENDPOINTS.USERS.ADD_SKILL, data);
  },

  updateSkill: async (skillId: string, data: UpdateSkillRequest): Promise<FreelancerSkill> => {
    return apiClient.patch(API_ENDPOINTS.USERS.UPDATE_SKILL(skillId), data);
  },

  deleteSkill: async (skillId: string): Promise<void> => {
    return apiClient.delete(API_ENDPOINTS.USERS.DELETE_SKILL(skillId));
  },

  // ========== Work Experience ==========
  getWorkExperience: async (): Promise<WorkExperience[]> => {
    return apiClient.get(API_ENDPOINTS.USERS.WORK_EXPERIENCE);
  },

  addWorkExperience: async (data: AddWorkExperienceRequest): Promise<WorkExperience> => {
    return apiClient.post(API_ENDPOINTS.USERS.ADD_WORK_EXPERIENCE, data);
  },

  updateWorkExperience: async (id: string, data: UpdateWorkExperienceRequest): Promise<WorkExperience> => {
    return apiClient.patch(API_ENDPOINTS.USERS.UPDATE_WORK_EXPERIENCE(id), data);
  },

  deleteWorkExperience: async (id: string): Promise<void> => {
    return apiClient.delete(API_ENDPOINTS.USERS.DELETE_WORK_EXPERIENCE(id));
  },

  // ========== Education ==========
  getEducation: async (): Promise<Education[]> => {
    return apiClient.get(API_ENDPOINTS.USERS.EDUCATION);
  },

  addEducation: async (data: AddEducationRequest): Promise<Education> => {
    return apiClient.post(API_ENDPOINTS.USERS.ADD_EDUCATION, data);
  },

  updateEducation: async (id: string, data: Partial<AddEducationRequest>): Promise<Education> => {
    return apiClient.patch(API_ENDPOINTS.USERS.UPDATE_EDUCATION(id), data);
  },

  deleteEducation: async (id: string): Promise<void> => {
    return apiClient.delete(API_ENDPOINTS.USERS.DELETE_EDUCATION(id));
  },

  // ========== Certifications ==========
  getCertifications: async (): Promise<Certification[]> => {
    return apiClient.get(API_ENDPOINTS.USERS.CERTIFICATIONS);
  },

  addCertification: async (data: AddCertificationRequest): Promise<Certification> => {
    return apiClient.post(API_ENDPOINTS.USERS.ADD_CERTIFICATION, data);
  },

  updateCertification: async (id: string, data: Partial<AddCertificationRequest>): Promise<Certification> => {
    return apiClient.patch(API_ENDPOINTS.USERS.UPDATE_CERTIFICATION(id), data);
  },

  deleteCertification: async (id: string): Promise<void> => {
    return apiClient.delete(API_ENDPOINTS.USERS.DELETE_CERTIFICATION(id));
  },

  // ========== Portfolio ==========
  getPortfolio: async (): Promise<PortfolioItem[]> => {
    return apiClient.get(API_ENDPOINTS.USERS.PORTFOLIO);
  },

  addPortfolioItem: async (data: AddPortfolioItemRequest): Promise<PortfolioItem> => {
    return apiClient.post(API_ENDPOINTS.USERS.ADD_PORTFOLIO_ITEM, data);
  },

  updatePortfolioItem: async (id: string, data: Partial<AddPortfolioItemRequest>): Promise<PortfolioItem> => {
    return apiClient.patch(API_ENDPOINTS.USERS.UPDATE_PORTFOLIO_ITEM(id), data);
  },

  deletePortfolioItem: async (id: string): Promise<void> => {
    return apiClient.delete(API_ENDPOINTS.USERS.DELETE_PORTFOLIO_ITEM(id));
  },

  uploadPortfolioImage: async (itemId: string, file: File | Blob): Promise<{ imageUrl: string }> => {
    const formData = new FormData();
    formData.append('image', file);
    return apiClient.post(API_ENDPOINTS.USERS.UPLOAD_PORTFOLIO_IMAGE(itemId), formData, {
      headers: { 'Content-Type': 'multipart/form-data' },
    });
  },

  // ========== Stats ==========
  getStats: async () => {
    return apiClient.get(API_ENDPOINTS.USERS.STATS);
  },

  getEarnings: async (params?: { from?: string; to?: string }) => {
    return apiClient.get(API_ENDPOINTS.USERS.EARNINGS, { params });
  },

  getReviews: async (params?: { page?: number; limit?: number }) => {
    return apiClient.get(API_ENDPOINTS.USERS.REVIEWS, { params });
  },
};