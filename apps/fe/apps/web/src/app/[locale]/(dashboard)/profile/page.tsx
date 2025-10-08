'use client';

import { useTranslations } from 'next-intl';
import { useAuth } from '@skillsier/shared';
import { useFreelancerProfile, useFreelancerSkills, usePortfolio, useWorkExperience } from '@skillsier/shared';
import { Card, Avatar, Badge } from '@skillsier/ui';
import { 
  Mail, Phone, MapPin, Calendar, Briefcase, 
  Award, DollarSign, Star, TrendingUp, Edit, 
  Eye, Clock, CheckCircle, ExternalLink
} from 'lucide-react';
import Link from 'next/link';

export default function FreelancerProfilePage() {
  const t = useTranslations('profile');
  const { user } = useAuth();
  const { data: profile, isLoading } = useFreelancerProfile();
  const { data: skills } = useFreelancerSkills();
  const { data: portfolio } = usePortfolio();
  const { data: experience } = useWorkExperience();

  if (isLoading) {
    return (
      <div className="flex items-center justify-center h-96">
        <div className="h-8 w-8 animate-spin rounded-full border-4 border-primary-600 border-t-transparent" />
      </div>
    );
  }

  if (!profile) return null;

  const isFreelancer = user?.userType === 'FREELANCER' || user?.userType === 'BOTH';

  return (
    <div className="space-y-6">
      {/* Header Section */}
      <Card padding="lg">
        <div className="flex flex-col lg:flex-row gap-6 items-start">
          <div className="relative">
            <Avatar src={profile.avatar} alt={profile.username} size="xl" className="h-32 w-32" />
            {profile.availability === 'AVAILABLE' && (
              <div className="absolute bottom-2 right-2 h-6 w-6 bg-green-500 border-4 border-white rounded-full" />
            )}
          </div>
          
          <div className="flex-1 w-full">
            <div className="flex flex-col md:flex-row md:items-start md:justify-between gap-4">
              <div className="flex-1">
                <h1 className="text-3xl font-bold text-gray-900">
                  {profile.firstName} {profile.lastName}
                </h1>
                <p className="text-xl text-gray-700 mt-1">{profile.professionalTitle || profile.title}</p>
                <p className="text-gray-600 mt-1">@{profile.username}</p>
                
                <div className="flex flex-wrap gap-2 mt-3">
                  <Badge variant={profile.emailVerified ? 'success' : 'warning'}>
                    {profile.emailVerified ? '✓ Email Verified' : 'Email Not Verified'}
                  </Badge>
                  {profile.identityVerified && (
                    <Badge variant="success">✓ Identity Verified</Badge>
                  )}
                  {profile.paymentVerified && (
                    <Badge variant="success">✓ Payment Verified</Badge>
                  )}
                  <Badge>{profile.userType}</Badge>
                </div>
              </div>
              
              <div className="flex gap-2">
                <Link href="/profile/edit">
                  <button className="flex items-center gap-2 px-4 py-2 bg-gray-100 hover:bg-gray-200 rounded-lg transition-colors">
                    <Edit className="h-4 w-4" />
                    Edit Profile
                  </button>
                </Link>
              </div>
            </div>

            {profile.overview && (
              <p className="text-gray-700 mt-4 leading-relaxed">{profile.overview}</p>
            )}

            <div className="grid grid-cols-1 md:grid-cols-2 gap-4 mt-4">
              <div className="flex items-center gap-2 text-gray-600">
                <Mail className="h-4 w-4" />
                <span>{profile.email}</span>
              </div>
              {profile.phoneNumber && (
                <div className="flex items-center gap-2 text-gray-600">
                  <Phone className="h-4 w-4" />
                  <span>{profile.phoneNumber}</span>
                </div>
              )}
              <div className="flex items-center gap-2 text-gray-600">
                <MapPin className="h-4 w-4" />
                <span>{profile.city}, {profile.country}</span>
              </div>
              <div className="flex items-center gap-2 text-gray-600">
                <Calendar className="h-4 w-4" />
                <span>Member since {new Date(profile.createdAt).toLocaleDateString()}</span>
              </div>
            </div>

            {isFreelancer && (
              <div className="mt-4 flex items-center gap-6 text-gray-700">
                <div className="flex items-center gap-2">
                  <DollarSign className="h-5 w-5 text-green-600" />
                  <span className="font-semibold">${profile.hourlyRate}/hr</span>
                </div>
                <div className="flex items-center gap-2">
                  <Clock className="h-5 w-5 text-blue-600" />
                  <span>Responds in {profile.responseTime} min</span>
                </div>
              </div>
            )}
          </div>
        </div>
      </Card>

      {/* Stats Section - Freelancer Only */}
      {isFreelancer && (
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-4">
          <Card padding="lg">
            <div className="flex items-center gap-3">
              <div className="p-3 bg-green-50 rounded-lg">
                <DollarSign className="h-6 w-6 text-green-600" />
              </div>
              <div>
                <p className="text-2xl font-bold text-gray-900">${profile.totalEarnings.toLocaleString()}</p>
                <p className="text-sm text-gray-600">Total Earned</p>
              </div>
            </div>
          </Card>

          <Card padding="lg">
            <div className="flex items-center gap-3">
              <div className="p-3 bg-blue-50 rounded-lg">
                <CheckCircle className="h-6 w-6 text-blue-600" />
              </div>
              <div>
                <p className="text-2xl font-bold text-gray-900">{profile.completedJobs}</p>
                <p className="text-sm text-gray-600">Jobs Completed</p>
              </div>
            </div>
          </Card>

          <Card padding="lg">
            <div className="flex items-center gap-3">
              <div className="p-3 bg-yellow-50 rounded-lg">
                <Star className="h-6 w-6 text-yellow-600" />
              </div>
              <div>
                <p className="text-2xl font-bold text-gray-900">{profile.rating.toFixed(1)}</p>
                <p className="text-sm text-gray-600">{profile.totalReviews} Reviews</p>
              </div>
            </div>
          </Card>

          <Card padding="lg">
            <div className="flex items-center gap-3">
              <div className="p-3 bg-purple-50 rounded-lg">
                <TrendingUp className="h-6 w-6 text-purple-600" />
              </div>
              <div>
                <p className="text-2xl font-bold text-gray-900">{profile.successRate}%</p>
                <p className="text-sm text-gray-600">Success Rate</p>
              </div>
            </div>
          </Card>
        </div>
      )}

      {/* Skills Section */}
      {isFreelancer && skills && skills.length > 0 && (
        <Card padding="lg">
          <div className="flex items-center justify-between mb-4">
            <h2 className="text-xl font-semibold text-gray-900">Skills & Expertise</h2>
            <Link href="/profile/skills">
              <button className="text-primary-600 hover:text-primary-700 text-sm font-medium">
                Manage Skills
              </button>
            </Link>
          </div>
          
          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
            {skills.map((skill) => (
              <div key={skill.id} className="p-4 border border-gray-200 rounded-lg hover:border-primary-300 transition-colors">
                <div className="flex items-center justify-between mb-2">
                  <h3 className="font-medium text-gray-900">{skill.name}</h3>
                  <Badge size="sm" variant="info">{skill.level}</Badge>
                </div>
                <p className="text-sm text-gray-600 mb-2">{skill.category}</p>
                <div className="flex items-center justify-between text-sm">
                  <span className="text-gray-600">{skill.yearsOfExperience} years exp.</span>
                  <span className="text-primary-600">{skill.endorsements} endorsements</span>
                </div>
              </div>
            ))}
          </div>
        </Card>
      )}

      {/* Work Experience Section */}
      {experience && experience.length > 0 && (
        <Card padding="lg">
          <div className="flex items-center justify-between mb-4">
            <h2 className="text-xl font-semibold text-gray-900">Work Experience</h2>
            <Link href="/profile/experience">
              <button className="text-primary-600 hover:text-primary-700 text-sm font-medium">
                Manage Experience
              </button>
            </Link>
          </div>
          
          <div className="space-y-6">
            {experience.map((exp) => (
              <div key={exp.id} className="flex gap-4 pb-6 border-b border-gray-200 last:border-0">
                <div className="flex-shrink-0">
                  <div className="h-12 w-12 bg-gray-100 rounded-lg flex items-center justify-center">
                    <Briefcase className="h-6 w-6 text-gray-600" />
                  </div>
                </div>
                <div className="flex-1">
                  <h3 className="font-semibold text-gray-900">{exp.title}</h3>
                  <p className="text-gray-700">{exp.company}</p>
                  <p className="text-sm text-gray-600 mt-1">
                    {new Date(exp.startDate).toLocaleDateString()} - {exp.isCurrent ? 'Present' : new Date(exp.endDate!).toLocaleDateString()}
                    {exp.location && ` • ${exp.location}`}
                  </p>
                  {exp.description && (
                    <p className="text-gray-600 mt-2">{exp.description}</p>
                  )}
                  {exp.skills && exp.skills.length > 0 && (
                    <div className="flex flex-wrap gap-2 mt-3">
                      {exp.skills.map((skill, idx) => (
                        <span key={idx} className="px-2 py-1 bg-gray-100 text-gray-700 text-xs rounded">
                          {skill}
                        </span>
                      ))}
                    </div>
                  )}
                </div>
              </div>
            ))}
          </div>
        </Card>
      )}

      {/* Portfolio Section */}
      {portfolio && portfolio.length > 0 && (
        <Card padding="lg">
          <div className="flex items-center justify-between mb-4">
            <h2 className="text-xl font-semibold text-gray-900">Portfolio</h2>
            <Link href="/profile/portfolio">
              <button className="text-primary-600 hover:text-primary-700 text-sm font-medium">
                Manage Portfolio
              </button>
            </Link>
          </div>
          
          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
            {portfolio.map((item) => (
              <div key={item.id} className="group cursor-pointer">
                <div className="relative aspect-video bg-gray-100 rounded-lg overflow-hidden mb-3">
                  {item.images[0] ? (
                    <img 
                      src={item.images[0]} 
                      alt={item.title}
                      className="w-full h-full object-cover group-hover:scale-105 transition-transform duration-300"
                    />
                  ) : (
                    <div className="w-full h-full flex items-center justify-center">
                      <Award className="h-12 w-12 text-gray-400" />
                    </div>
                  )}
                  {item.featured && (
                    <div className="absolute top-2 right-2">
                      <Badge variant="warning">Featured</Badge>
                    </div>
                  )}
                </div>
                <h3 className="font-semibold text-gray-900 group-hover:text-primary-600 transition-colors">
                  {item.title}
                </h3>
                <p className="text-sm text-gray-600 mt-1 line-clamp-2">{item.description}</p>
                {item.projectUrl && (
                  <a 
                    href={item.projectUrl} 
                    target="_blank" 
                    rel="noopener noreferrer"
                    className="inline-flex items-center gap-1 text-sm text-primary-600 hover:text-primary-700 mt-2"
                  >
                    View Project <ExternalLink className="h-3 w-3" />
                  </a>
                )}
                <div className="flex flex-wrap gap-1 mt-2">
                  {item.skills.slice(0, 3).map((skill, idx) => (
                    <span key={idx} className="px-2 py-0.5 bg-gray-100 text-gray-600 text-xs rounded">
                      {skill}
                    </span>
                  ))}
                  {item.skills.length > 3 && (
                    <span className="px-2 py-0.5 bg-gray-100 text-gray-600 text-xs rounded">
                      +{item.skills.length - 3}
                    </span>
                  )}
                </div>
              </div>
            ))}
          </div>
        </Card>
      )}

      {/* Profile Strength */}
      {isFreelancer && (
        <Card padding="lg">
          <div className="flex items-center justify-between mb-4">
            <h2 className="text-xl font-semibold text-gray-900">Profile Strength</h2>
            <span className="text-2xl font-bold text-primary-600">{profile.profileStrength}%</span>
          </div>
          <div className="w-full bg-gray-200 rounded-full h-3 mb-4">
            <div 
              className="bg-primary-600 h-3 rounded-full transition-all duration-500"
              style={{ width: `${profile.profileStrength}%` }}
            />
          </div>
          <div className="space-y-2 text-sm">
            {profile.profileStrength < 100 && (
              <>
                {!profile.avatar && (
                  <p className="text-gray-600">• Add a profile photo (+10%)</p>
                )}
                {!profile.overview && (
                  <p className="text-gray-600">• Write a professional overview (+15%)</p>
                )}
                {(!skills || skills.length < 5) && (
                  <p className="text-gray-600">• Add at least 5 skills (+15%)</p>
                )}
                {(!portfolio || portfolio.length === 0) && (
                  <p className="text-gray-600">• Add portfolio items (+20%)</p>
                )}
                {(!experience || experience.length === 0) && (
                  <p className="text-gray-600">• Add work experience (+15%)</p>
                )}
                {!profile.identityVerified && (
                  <p className="text-gray-600">• Verify your identity (+10%)</p>
                )}
              </>
            )}
            {profile.profileStrength === 100 && (
              <p className="text-green-600 font-medium">✓ Your profile is complete!</p>
            )}
          </div>
        </Card>
      )}

      {/* View Public Profile */}
      <div className="flex justify-center">
        <Link href={`/freelancers/${profile.username}`} target="_blank">
          <button className="flex items-center gap-2 px-6 py-3 border-2 border-gray-300 rounded-lg hover:border-primary-600 hover:text-primary-600 transition-colors">
            <Eye className="h-5 w-5" />
            View Public Profile
          </button>
        </Link>
      </div>
    </div>
  );
}
