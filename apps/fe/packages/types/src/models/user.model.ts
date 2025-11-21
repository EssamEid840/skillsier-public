export interface UserDTO {
  user_id: string;
  email: string;
  first_name: string;
  last_name: string;
  role: string;
  account_status: string;
  email_verified: boolean;
  phone_verified: boolean;
  kyc_status: string;
  profile_picture?: string;
  created_at: string;
  updated_at: string;
}

export interface UserProfileDTO {
  user_id: string;
  bio?: string;
  title?: string;
  hourly_rate?: number;
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
    start_date: string;
    end_date?: string;
    description?: string;
  }[];
  education?: {
    degree: string;
    institution: string;
    start_date: string;
    end_date?: string;
  }[];
}

export interface CreateUserRequest {
  email: string;
  password: string;
  first_name: string;
  last_name: string;
  role: string;
}

export interface UpdateUserRequest {
  first_name?: string;
  last_name?: string;
  profile_picture?: string;
}