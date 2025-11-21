import type {
  UserRole,
  AccountStatus,
  VerificationStatus,
  KYCStatus,
} from '../enums/user.enum';

export interface User {
  id: string;
  email: string;
  firstName: string;
  lastName: string;
  role: UserRole;
  accountStatus: AccountStatus;
  emailVerified: boolean;
  phoneVerified: boolean;
  kycStatus: KYCStatus;
  profilePicture?: string;
  createdAt: Date;
  updatedAt: Date;
}

export interface UserProfile {
  userId: string;
  bio?: string;
  title?: string;
  hourlyRate?: number;
  skills: string[];
  languages: string[];
  location?: {
    country: string;
    city: string;
    timezone: string;
  };
  portfolio?: {
    url: string;
    title: string;
    description?: string;
  }[];
  experience?: {
    title: string;
    company: string;
    startDate: Date;
    endDate?: Date;
    description?: string;
  }[];
  education?: {
    degree: string;
    institution: string;
    startDate: Date;
    endDate?: Date;
  }[];
}

export interface UserSettings {
  userId: string;
  theme: 'light' | 'dark' | 'auto';
  language: string;
  timezone: string;
  emailNotifications: boolean;
  pushNotifications: boolean;
  smsNotifications: boolean;
}