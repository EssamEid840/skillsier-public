// packages/shared/src/features/user/types/user.types.ts
// User-related types for freelancing platform

export type UserType = 'freelancer' | 'client';

export interface User {
  id: string;
  keycloakId: string;
  email: string;
  firstName: string;
  lastName: string;
  userType: UserType;
  avatarUrl?: string;
  createdAt: Date;
  updatedAt: Date;
}

export interface FreelancerProfile {
  id: string;
  userId: string;
  title: string;
  bio: string;
  hourlyRate: number;
  availability: 'available' | 'busy' | 'unavailable';
  skills: UserSkill[];
  workExperience: WorkExperience[];
  education: Education[];
  certifications: Certification[];
  portfolio: PortfolioItem[];
  languages: Language[];
  completionRate: number;
  responseTime: number; // in hours
  totalEarnings: number;
  totalJobs: number;
  rating: number;
  reviewCount: number;
  createdAt: Date;
  updatedAt: Date;
}

export interface ClientProfile {
  id: string;
  userId: string;
  companyName?: string;
  companySize?: string;
  industry?: string;
  website?: string;
  bio?: string;
  totalSpent: number;
  totalJobs: number;
  rating: number;
  reviewCount: number;
  createdAt: Date;
  updatedAt: Date;
}

export interface UserSkill {
  id: string;
  freelancerProfileId: string;
  skillName: string;
  proficiencyLevel: 'beginner' | 'intermediate' | 'expert';
  yearsOfExperience: number;
  projectsCompleted: number;
  createdAt: Date;
  updatedAt: Date;
}

export interface WorkExperience {
  id: string;
  freelancerProfileId: string;
  title: string;
  company: string;
  location?: string;
  startDate: Date;
  endDate?: Date;
  current: boolean;
  description?: string;
  createdAt: Date;
  updatedAt: Date;
}

export interface Education {
  id: string;
  freelancerProfileId: string;
  institution: string;
  degree: string;
  fieldOfStudy: string;
  startDate: Date;
  endDate?: Date;
  current: boolean;
  description?: string;
  createdAt: Date;
  updatedAt: Date;
}

export interface Certification {
  id: string;
  freelancerProfileId: string;
  name: string;
  issuingOrganization: string;
  issueDate: Date;
  expirationDate?: Date;
  credentialId?: string;
  credentialUrl?: string;
  createdAt: Date;
  updatedAt: Date;
}

export interface PortfolioItem {
  id: string;
  freelancerProfileId: string;
  title: string;
  description: string;
  projectUrl?: string;
  imageUrls: string[];
  skills: string[];
  completionDate: Date;
  createdAt: Date;
  updatedAt: Date;
}

export interface Language {
  id: string;
  freelancerProfileId: string;
  languageName: string;
  proficiency: 'basic' | 'conversational' | 'fluent' | 'native';
  createdAt: Date;
  updatedAt: Date;
}

export interface UserStats {
  totalEarnings: number;
  totalJobs: number;
  activeJobs: number;
  completedJobs: number;
  rating: number;
  reviewCount: number;
  successRate: number;
  responseTime: number;
}

export interface Earnings {
  id: string;
  userId: string;
  amount: number;
  currency: string;
  jobId: string;
  status: 'pending' | 'completed' | 'withdrawn';
  createdAt: Date;
  updatedAt: Date;
}

export interface Review {
  id: string;
  fromUserId: string;
  toUserId: string;
  jobId: string;
  rating: number;
  comment: string;
  createdAt: Date;
  updatedAt: Date;
}