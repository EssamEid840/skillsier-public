// packages/shared/src/constants/app.ts
// Application-wide constants

export const APP_NAME = 'Skillsier';
export const APP_DESCRIPTION = 'Connect with talented freelancers and clients worldwide';
export const APP_VERSION = '1.0.0';

// API Configuration
export const API_TIMEOUT = 30000; // 30 seconds
export const API_RETRY_ATTEMPTS = 3;

// Pagination
export const DEFAULT_PAGE_SIZE = 20;
export const MAX_PAGE_SIZE = 100;

// File Upload
export const MAX_FILE_SIZE_MB = 10;
export const MAX_IMAGE_SIZE_MB = 5;
export const ALLOWED_IMAGE_TYPES = ['image/jpeg', 'image/png', 'image/webp'];
export const ALLOWED_DOCUMENT_TYPES = ['application/pdf', 'application/msword', 'application/vnd.openxmlformats-officedocument.wordprocessingml.document'];

// User Limits
export const MAX_SKILLS = 20;
export const MAX_WORK_EXPERIENCE = 10;
export const MAX_EDUCATION = 5;
export const MAX_CERTIFICATIONS = 15;
export const MAX_PORTFOLIO_ITEMS = 20;
export const MAX_PORTFOLIO_IMAGES = 5;

// Validation
export const MIN_PASSWORD_LENGTH = 8;
export const MAX_BIO_LENGTH = 5000;
export const MAX_TITLE_LENGTH = 100;
export const MAX_DESCRIPTION_LENGTH = 2000;

// Rating
export const MIN_RATING = 1;
export const MAX_RATING = 5;

// Currency
export const DEFAULT_CURRENCY = 'USD';
export const SUPPORTED_CURRENCIES = ['USD', 'EUR', 'GBP', 'CAD', 'AUD'];

// Languages
export const SUPPORTED_LANGUAGES = [
  { code: 'en', name: 'English', nativeName: 'English' },
  { code: 'ar', name: 'Arabic', nativeName: 'العربية', rtl: true },
  { code: 'es', name: 'Spanish', nativeName: 'Español' },
  { code: 'fr', name: 'French', nativeName: 'Français' },
  { code: 'de', name: 'German', nativeName: 'Deutsch' },
  { code: 'zh', name: 'Chinese', nativeName: '中文' },
  { code: 'hi', name: 'Hindi', nativeName: 'हिन्दी' },
  { code: 'tr', name: 'Turkish', nativeName: 'Türkçe' },
  { code: 'ru', name: 'Russian', nativeName: 'Русский' },
];

// Skill Levels
export const SKILL_LEVELS = [
  { value: 'beginner', label: 'Beginner', years: '0-2' },
  { value: 'intermediate', label: 'Intermediate', years: '2-5' },
  { value: 'expert', label: 'Expert', years: '5+' },
] as const;

// Language Proficiency
export const LANGUAGE_PROFICIENCY = [
  { value: 'basic', label: 'Basic' },
  { value: 'conversational', label: 'Conversational' },
  { value: 'fluent', label: 'Fluent' },
  { value: 'native', label: 'Native' },
] as const;

// Job Status
export const JOB_STATUS = {
  DRAFT: 'draft',
  OPEN: 'open',
  IN_PROGRESS: 'in_progress',
  COMPLETED: 'completed',
  CANCELLED: 'cancelled',
} as const;

// Proposal Status
export const PROPOSAL_STATUS = {
  PENDING: 'pending',
  ACCEPTED: 'accepted',
  REJECTED: 'rejected',
  WITHDRAWN: 'withdrawn',
} as const;

// User Availability
export const AVAILABILITY_STATUS = {
  AVAILABLE: 'available',
  BUSY: 'busy',
  UNAVAILABLE: 'unavailable',
} as const;

// Routes
export const ROUTES = {
  HOME: '/',
  LOGIN: '/login',
  REGISTER: '/register',
  DASHBOARD: '/dashboard',
  PROFILE: '/profile',
  JOBS: '/jobs',
  PROPOSALS: '/proposals',
  MESSAGES: '/messages',
  SETTINGS: '/settings',
} as const;