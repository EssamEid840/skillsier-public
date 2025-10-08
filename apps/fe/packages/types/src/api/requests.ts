export interface RegisterRequest {
  email: string;
  username: string;
  password: string;
  firstName: string;
  lastName: string;
  userType: 'FREELANCER' | 'CLIENT' | 'BOTH';
  country: string;
  agreeToTerms: boolean;
}

export interface UpdateProfileRequest {
  firstName?: string;
  lastName?: string;
  bio?: string;
  title?: string;
  avatar?: string;
  phoneNumber?: string;
  dateOfBirth?: string;
  gender?: string;
  country?: string;
  city?: string;
  timezone?: string;
  language?: string;
  languages?: string[];
}

export interface UpdateFreelancerProfileRequest {
  professionalTitle?: string;
  overview?: string;
  expertise?: string[];
  hourlyRate?: number;
  minimumProjectBudget?: number;
  availability?: string;
  preferredJobTypes?: string[];
}

export interface UpdateClientProfileRequest {
  companyName?: string;
  companySize?: string;
  industry?: string;
  website?: string;
}

export interface AddSkillRequest {
  skillId: string;
  level: 'BEGINNER' | 'INTERMEDIATE' | 'ADVANCED' | 'EXPERT';
  yearsOfExperience: number;
}

export interface UpdateSkillRequest {
  level?: 'BEGINNER' | 'INTERMEDIATE' | 'ADVANCED' | 'EXPERT';
  yearsOfExperience?: number;
}

export interface AddWorkExperienceRequest {
  title: string;
  company: string;
  location?: string;
  startDate: string;
  endDate?: string;
  isCurrent: boolean;
  description: string;
  skills?: string[];
}

export interface UpdateWorkExperienceRequest {
  title?: string;
  company?: string;
  location?: string;
  startDate?: string;
  endDate?: string;
  isCurrent?: boolean;
  description?: string;
  skills?: string[];
}

export interface AddEducationRequest {
  degree: string;
  institution: string;
  fieldOfStudy: string;
  startDate: string;
  endDate?: string;
  description?: string;
}

export interface AddCertificationRequest {
  name: string;
  issuer: string;
  issueDate: string;
  expiryDate?: string;
  credentialId?: string;
  credentialUrl?: string;
}

export interface AddPortfolioItemRequest {
  title: string;
  description: string;
  projectUrl?: string;
  skills: string[];
  completedAt: string;
}

export interface ChangePasswordRequest {
  currentPassword: string;
  newPassword: string;
  confirmPassword: string;
}
