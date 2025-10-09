// Export all user hooks
export * from './hooks';

// Export user API
export * from './api/userApi';

// Export user types (re-export from types package)
export type {
  User,
  FreelancerProfile,
  ClientProfile,
  FreelancerSkill,
  WorkExperience,
  Education,
  Certification,
  PortfolioItem,
  UpdateProfileRequest,
  UpdateFreelancerProfileRequest,
  UpdateClientProfileRequest,
  AddSkillRequest,
  UpdateSkillRequest,
  AddWorkExperienceRequest,
  UpdateWorkExperienceRequest,
  AddEducationRequest,
  UpdateEducationRequest,
  AddCertificationRequest,
  UpdateCertificationRequest,
  AddPortfolioItemRequest,
  UpdatePortfolioItemRequest,
  ChangePasswordRequest,
  UpdatePreferencesRequest,
  UpdateSocialLinksRequest,
} from '@skillsier/types';