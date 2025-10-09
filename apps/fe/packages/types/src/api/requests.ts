// Authentication Requests
export interface LoginRequest {
  username: string;
  password: string;
}

export interface RegisterRequest {
  firstName: string;
  lastName: string;
  email: string;
  password: string;
  userType?: 'FREELANCER' | 'CLIENT' | 'BOTH';
}

export interface RefreshTokenRequest {
  refreshToken: string;
}

export interface ChangePasswordRequest {
  currentPassword: string;
  newPassword: string;
}

// User Profile Requests
export interface UpdateProfileRequest {
  firstName?: string;
  lastName?: string;
  phone?: string;
  location?: string;
  bio?: string;
  timezone?: string;
}

export interface UpdatePreferencesRequest {
  language?: string;
  emailNotifications?: boolean;
  smsNotifications?: boolean;
  marketingEmails?: boolean;
}

export interface UpdateSocialLinksRequest {
  linkedin?: string;
  github?: string;
  twitter?: string;
  website?: string;
  portfolio?: string;
}

// Freelancer Profile Requests
export interface UpdateFreelancerProfileRequest {
  professionalTitle?: string;
  overview?: string;
  hourlyRate?: number;
  availability?: 'AVAILABLE' | 'BUSY' | 'NOT_AVAILABLE';
  preferredJobTypes?: ('FIXED_PRICE' | 'HOURLY' | 'CONTRACT')[];
}

// Client Profile Requests
export interface UpdateClientProfileRequest {
  companyName?: string;
  companySize?: string;
  industry?: string;
  website?: string;
}

// Skills Requests
export interface AddSkillRequest {
  name: string;
  category: string;
  level: 'BEGINNER' | 'INTERMEDIATE' | 'ADVANCED' | 'EXPERT';
  yearsOfExperience: number;
}

export interface UpdateSkillRequest {
  level?: 'BEGINNER' | 'INTERMEDIATE' | 'ADVANCED' | 'EXPERT';
  yearsOfExperience?: number;
}

// Work Experience Requests
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

// Education Requests
export interface AddEducationRequest {
  degree: string;
  institution: string;
  fieldOfStudy: string;
  startDate: string;
  endDate?: string;
  description?: string;
}

export interface UpdateEducationRequest {
  degree?: string;
  institution?: string;
  fieldOfStudy?: string;
  startDate?: string;
  endDate?: string;
  description?: string;
}

// Certification Requests
export interface AddCertificationRequest {
  name: string;
  issuer: string;
  issueDate: string;
  expiryDate?: string;
  credentialId?: string;
  credentialUrl?: string;
}

export interface UpdateCertificationRequest {
  name?: string;
  issuer?: string;
  issueDate?: string;
  expiryDate?: string;
  credentialId?: string;
  credentialUrl?: string;
}

// Portfolio Requests
export interface AddPortfolioItemRequest {
  title: string;
  description: string;
  projectUrl?: string;
  skills: string[];
  completedAt: string;
  featured?: boolean;
}

export interface UpdatePortfolioItemRequest {
  title?: string;
  description?: string;
  projectUrl?: string;
  skills?: string[];
  completedAt?: string;
  featured?: boolean;
}