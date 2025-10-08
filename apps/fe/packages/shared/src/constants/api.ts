export const API_ENDPOINTS = {
  AUTH: {
    LOGIN: '/auth/login',
    REGISTER: '/auth/register',
    LOGOUT: '/auth/logout',
    REFRESH: '/auth/refresh',
    ME: '/auth/me',
    VERIFY_EMAIL: '/auth/verify-email',
    RESEND_VERIFICATION: '/auth/resend-verification',
    FORGOT_PASSWORD: '/auth/forgot-password',
    RESET_PASSWORD: '/auth/reset-password',
  },
  
  USERS: {
    // Basic Profile
    PROFILE: '/users/profile',
    UPDATE_PROFILE: '/users/profile',
    UPDATE_AVATAR: '/users/profile/avatar',
    DELETE_AVATAR: '/users/profile/avatar',
    UPDATE_PREFERENCES: '/users/profile/preferences',
    CHANGE_PASSWORD: '/users/profile/password',
    DELETE_ACCOUNT: '/users/profile',
    
    // Freelancer Profile
    FREELANCER_PROFILE: '/users/freelancer/profile',
    UPDATE_FREELANCER_PROFILE: '/users/freelancer/profile',
    
    // Client Profile
    CLIENT_PROFILE: '/users/client/profile',
    UPDATE_CLIENT_PROFILE: '/users/client/profile',
    
    // Skills
    SKILLS: '/users/profile/skills',
    ADD_SKILL: '/users/profile/skills',
    UPDATE_SKILL: (skillId: string) => `/users/profile/skills/${skillId}`,
    DELETE_SKILL: (skillId: string) => `/users/profile/skills/${skillId}`,
    
    // Work Experience
    WORK_EXPERIENCE: '/users/profile/experience',
    ADD_WORK_EXPERIENCE: '/users/profile/experience',
    UPDATE_WORK_EXPERIENCE: (id: string) => `/users/profile/experience/${id}`,
    DELETE_WORK_EXPERIENCE: (id: string) => `/users/profile/experience/${id}`,
    
    // Education
    EDUCATION: '/users/profile/education',
    ADD_EDUCATION: '/users/profile/education',
    UPDATE_EDUCATION: (id: string) => `/users/profile/education/${id}`,
    DELETE_EDUCATION: (id: string) => `/users/profile/education/${id}`,
    
    // Certifications
    CERTIFICATIONS: '/users/profile/certifications',
    ADD_CERTIFICATION: '/users/profile/certifications',
    UPDATE_CERTIFICATION: (id: string) => `/users/profile/certifications/${id}`,
    DELETE_CERTIFICATION: (id: string) => `/users/profile/certifications/${id}`,
    
    // Portfolio
    PORTFOLIO: '/users/profile/portfolio',
    ADD_PORTFOLIO_ITEM: '/users/profile/portfolio',
    UPDATE_PORTFOLIO_ITEM: (id: string) => `/users/profile/portfolio/${id}`,
    DELETE_PORTFOLIO_ITEM: (id: string) => `/users/profile/portfolio/${id}`,
    UPLOAD_PORTFOLIO_IMAGE: (id: string) => `/users/profile/portfolio/${id}/images`,
    
    // Stats & Analytics
    STATS: '/users/profile/stats',
    EARNINGS: '/users/profile/earnings',
    REVIEWS: '/users/profile/reviews',
    
    // Verification
    VERIFY_IDENTITY: '/users/profile/verify/identity',
    VERIFY_PAYMENT: '/users/profile/verify/payment',
  },
  
  JOBS: {
    LIST: '/jobs',
    CREATE: '/jobs',
    DETAILS: (id: string) => `/jobs/${id}`,
    UPDATE: (id: string) => `/jobs/${id}`,
    DELETE: (id: string) => `/jobs/${id}`,
    MY_JOBS: '/jobs/my-jobs',
  },
  
  PROPOSALS: {
    LIST: '/proposals',
    CREATE: '/proposals',
    DETAILS: (id: string) => `/proposals/${id}`,
    UPDATE: (id: string) => `/proposals/${id}`,
    WITHDRAW: (id: string) => `/proposals/${id}/withdraw`,
    MY_PROPOSALS: '/proposals/my-proposals',
  },
  
  CONTRACTS: {
    LIST: '/contracts',
    DETAILS: (id: string) => `/contracts/${id}`,
    ACCEPT: (id: string) => `/contracts/${id}/accept`,
    DECLINE: (id: string) => `/contracts/${id}/decline`,
    COMPLETE: (id: string) => `/contracts/${id}/complete`,
    MY_CONTRACTS: '/contracts/my-contracts',
  },
  
  REVIEWS: {
    CREATE: '/reviews',
    MY_REVIEWS: '/reviews/my-reviews',
    GIVEN_REVIEWS: '/reviews/given',
    RECEIVED_REVIEWS: '/reviews/received',
  },
} as const;
