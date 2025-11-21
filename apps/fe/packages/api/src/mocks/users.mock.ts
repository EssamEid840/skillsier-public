import type { UserDTO, UserProfileDTO } from '@skillsier/types';

export const mockUsers: UserDTO[] = [
  {
    user_id: 'user-1',
    email: 'freelancer@skillsier.dev',
    first_name: 'John',
    last_name: 'Doe',
    role: 'FREELANCER',
    account_status: 'ACTIVE',
    email_verified: true,
    phone_verified: true,
    kyc_status: 'VERIFIED',
    profile_picture: 'https://api.dicebear.com/7.x/avataaars/svg?seed=John',
    created_at: '2024-01-15T10:00:00Z',
    updated_at: '2024-11-01T14:30:00Z',
  },
  {
    user_id: 'user-2',
    email: 'client@skillsier.dev',
    first_name: 'Jane',
    last_name: 'Smith',
    role: 'CLIENT',
    account_status: 'ACTIVE',
    email_verified: true,
    phone_verified: false,
    kyc_status: 'VERIFIED',
    profile_picture: 'https://api.dicebear.com/7.x/avataaars/svg?seed=Jane',
    created_at: '2024-02-20T08:00:00Z',
    updated_at: '2024-10-15T12:00:00Z',
  },
];

export const mockUserProfile: UserProfileDTO = {
  user_id: 'user-1',
  bio: 'Experienced full-stack developer with 8+ years of experience building scalable web applications.',
  title: 'Senior Full-Stack Developer',
  hourly_rate: 75,
  skills: ['React', 'Node.js', 'TypeScript', 'PostgreSQL', 'AWS'],
  languages: ['English', 'Spanish'],
  location: {
    country: 'United States',
    city: 'San Francisco',
    timezone: 'America/Los_Angeles',
  },
  portfolio: [
    {
      url: 'https://example.com/project1',
      title: 'E-commerce Platform',
      description: 'Built a full-featured e-commerce platform with React and Node.js',
    },
  ],
  experience: [
    {
      title: 'Senior Developer',
      company: 'Tech Corp',
      start_date: '2020-01-01',
      description: 'Led development of microservices architecture',
    },
  ],
  education: [
    {
      degree: 'B.S. Computer Science',
      institution: 'State University',
      start_date: '2012-09-01',
      end_date: '2016-05-01',
    },
  ],
};

export const getMockUser = (userId: string): UserDTO | undefined => {
  return mockUsers.find(user => user.user_id === userId);
};