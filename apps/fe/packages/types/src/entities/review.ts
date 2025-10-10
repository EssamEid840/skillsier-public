// packages/types/src/entities/review.ts
// Review entity types

export interface Review {
  id: string;
  contractId: string;
  jobId: string;
  fromUserId: string;
  toUserId: string;
  rating: number;
  comment: string;
  skills: ReviewSkill[];
  isPublic: boolean;
  response?: string;
  createdAt: Date;
  updatedAt: Date;
}

export interface ReviewSkill {
  skill: string;
  rating: number;
}

export interface ReviewDetails extends Review {
  contract: {
    id: string;
    title: string;
  };
  job: {
    id: string;
    title: string;
  };
  fromUser: {
    id: string;
    name: string;
    avatarUrl?: string;
    userType: 'freelancer' | 'client';
  };
  toUser: {
    id: string;
    name: string;
    avatarUrl?: string;
    userType: 'freelancer' | 'client';
  };
}

export interface CreateReviewRequest {
  contractId: string;
  rating: number;
  comment: string;
  skills?: ReviewSkill[];
  isPublic: boolean;
}

export interface RespondToReviewRequest {
  reviewId: string;
  response: string;
}

export interface ReviewStats {
  averageRating: number;
  totalReviews: number;
  ratingDistribution: {
    5: number;
    4: number;
    3: number;
    2: number;
    1: number;
  };
  skillRatings: {
    skill: string;
    averageRating: number;
  }[];
}