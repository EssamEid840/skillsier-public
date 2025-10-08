export interface User {
  id: string;
  keycloakId: string; // Keycloak UUID
  email: string;
  username: string;
  firstName: string;
  lastName: string;
  avatar?: string;
  phoneNumber?: string;
  bio?: string;
  title?: string; // Professional title (e.g., "Full Stack Developer")
  dateOfBirth?: string;
  gender?: 'MALE' | 'FEMALE' | 'OTHER' | 'PREFER_NOT_TO_SAY';
  country?: string;
  city?: string;
  timezone?: string;
  language: string;
  languages: string[]; // Spoken languages
  hourlyRate?: number; // For freelancers
  currency: string;
  availability: 'AVAILABLE' | 'BUSY' | 'NOT_AVAILABLE';
  userType: UserType;
  role: UserRole;
  emailVerified: boolean;
  phoneVerified: boolean;
  identityVerified: boolean;
  paymentVerified: boolean;
  status: UserStatus;
  createdAt: string;
  updatedAt: string;
  lastLogin?: string;
  preferences: UserPreferences;
}

export enum UserType {
  FREELANCER = 'FREELANCER',
  CLIENT = 'CLIENT',
  BOTH = 'BOTH', // Can be both freelancer and client
}

export enum UserRole {
  USER = 'USER',
  ADMIN = 'ADMIN',
  MODERATOR = 'MODERATOR',
}

export enum UserStatus {
  ACTIVE = 'ACTIVE',
  INACTIVE = 'INACTIVE',
  SUSPENDED = 'SUSPENDED',
  PENDING_VERIFICATION = 'PENDING_VERIFICATION',
  BANNED = 'BANNED',
}

export interface UserPreferences {
  theme: 'light' | 'dark' | 'system';
  language: string;
  notifications: NotificationSettings;
  privacy: PrivacySettings;
}

export interface NotificationSettings {
  email: boolean;
  push: boolean;
  sms: boolean;
  newJobAlerts: boolean;
  proposalUpdates: boolean;
  messageNotifications: boolean;
  paymentNotifications: boolean;
  reviewNotifications: boolean;
  marketingEmails: boolean;
}

export interface PrivacySettings {
  profileVisibility: 'PUBLIC' | 'PRIVATE' | 'CLIENTS_ONLY';
  showEmail: boolean;
  showPhoneNumber: boolean;
  showHourlyRate: boolean;
  showEarnings: boolean;
  allowDirectContact: boolean;
}

// Freelancer-specific profile
export interface FreelancerProfile extends User {
  // Professional Info
  professionalTitle: string;
  overview: string; // Detailed bio
  expertise: string[];
  skills: FreelancerSkill[];
  
  // Experience
  experience: WorkExperience[];
  education: Education[];
  certifications: Certification[];
  
  // Portfolio
  portfolio: PortfolioItem[];
  
  // Stats
  totalEarnings: number;
  completedJobs: number;
  ongoingJobs: number;
  successRate: number;
  rating: number;
  totalReviews: number;
  responseTime: number; // in minutes
  
  // Visibility
  profileViews: number;
  profileStrength: number; // 0-100
  
  // Settings
  hourlyRate: number;
  minimumProjectBudget?: number;
  availability: 'AVAILABLE' | 'BUSY' | 'NOT_AVAILABLE';
  preferredJobTypes: JobType[];
}

// Client-specific profile
export interface ClientProfile extends User {
  // Company Info
  companyName?: string;
  companySize?: string;
  industry?: string;
  website?: string;
  
  // Stats
  totalSpent: number;
  postedJobs: number;
  activeJobs: number;
  hiredFreelancers: number;
  rating: number;
  totalReviews: number;
  
  // Payment
  paymentMethodVerified: boolean;
  preferredPaymentMethod?: string;
}

export interface FreelancerSkill {
  id: string;
  skillId: string;
  name: string;
  category: string;
  level: 'BEGINNER' | 'INTERMEDIATE' | 'ADVANCED' | 'EXPERT';
  yearsOfExperience: number;
  endorsements: number;
  addedAt: string;
}

export interface WorkExperience {
  id: string;
  title: string;
  company: string;
  location?: string;
  startDate: string;
  endDate?: string;
  isCurrent: boolean;
  description: string;
  skills: string[];
}

export interface Education {
  id: string;
  degree: string;
  institution: string;
  fieldOfStudy: string;
  startDate: string;
  endDate?: string;
  description?: string;
}

export interface Certification {
  id: string;
  name: string;
  issuer: string;
  issueDate: string;
  expiryDate?: string;
  credentialId?: string;
  credentialUrl?: string;
}

export interface PortfolioItem {
  id: string;
  title: string;
  description: string;
  images: string[];
  projectUrl?: string;
  skills: string[];
  completedAt: string;
  featured: boolean;
}

export enum JobType {
  FIXED_PRICE = 'FIXED_PRICE',
  HOURLY = 'HOURLY',
  CONTRACT = 'CONTRACT',
}