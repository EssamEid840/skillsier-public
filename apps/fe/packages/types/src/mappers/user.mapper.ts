import type { User, UserProfile } from '../domains/user';
import type { UserDTO, UserProfileDTO } from '../models/user.model';
import {
  UserRole,
  AccountStatus,
  KYCStatus,
} from '../enums/user.enum';

export const mapUserDTOToDomain = (dto: UserDTO): User => ({
  id: dto.user_id,
  email: dto.email,
  firstName: dto.first_name,
  lastName: dto.last_name,
  role: dto.role as UserRole,
  accountStatus: dto.account_status as AccountStatus,
  emailVerified: dto.email_verified,
  phoneVerified: dto.phone_verified,
  kycStatus: dto.kyc_status as KYCStatus,
  profilePicture: dto.profile_picture,
  createdAt: new Date(dto.created_at),
  updatedAt: new Date(dto.updated_at),
});

export const mapUserProfileDTOToDomain = (dto: UserProfileDTO): UserProfile => ({
  userId: dto.user_id,
  bio: dto.bio,
  title: dto.title,
  hourlyRate: dto.hourly_rate,
  skills: dto.skills,
  languages: dto.languages,
  location: dto.location,
  portfolio: dto.portfolio,
  experience: dto.experience?.map(exp => ({
    title: exp.title,
    company: exp.company,
    startDate: new Date(exp.start_date),
    endDate: exp.end_date ? new Date(exp.end_date) : undefined,
    description: exp.description,
  })),
  education: dto.education?.map(edu => ({
    degree: edu.degree,
    institution: edu.institution,
    startDate: new Date(edu.start_date),
    endDate: edu.end_date ? new Date(edu.end_date) : undefined,
  })),
});