// internal/domain/user/statistics.go
package user

import "time"

// UserStatistics contains comprehensive user statistics
type UserStatistics struct {
	// Total counts by status
	TotalUsers     int64 `json:"total_users"`
	ActiveUsers    int64 `json:"active_users"`
	PendingUsers   int64 `json:"pending_users"`
	InactiveUsers  int64 `json:"inactive_users"`
	SuspendedUsers int64 `json:"suspended_users"`
	BannedUsers    int64 `json:"banned_users"`
	DeletedUsers   int64 `json:"deleted_users"`
	RestrictedUsers int64 `json:"restricted_users"`
	
	// Counts by user type
	TotalFreelancers int64 `json:"total_freelancers"`
	TotalClients     int64 `json:"total_clients"`
	TotalBoth        int64 `json:"total_both"`
	TotalStaff       int64 `json:"total_staff"`
	
	// Verification statistics
	VerifiedUsers          int64   `json:"verified_users"`
	UnverifiedUsers        int64   `json:"unverified_users"`
	EmailVerifiedUsers     int64   `json:"email_verified_users"`
	PhoneVerifiedUsers     int64   `json:"phone_verified_users"`
	IdentityVerifiedUsers  int64   `json:"identity_verified_users"`
	VerificationRate       float64 `json:"verification_rate"`
	
	// Activity statistics
	OnlineUsers        int64   `json:"online_users"`
	ActiveToday        int64   `json:"active_today"`
	ActiveThisWeek     int64   `json:"active_this_week"`
	ActiveThisMonth    int64   `json:"active_this_month"`
	
	// Quality metrics
	AverageRating           float64 `json:"average_rating"`
	AverageCompleteness     float64 `json:"average_completeness"`
	AverageSuccessRate      float64 `json:"average_success_rate"`
	AverageCompletedJobs    float64 `json:"average_completed_jobs"`
	
	// Badge statistics
	TopRatedCount       int64 `json:"top_rated_count"`
	RisingTalentCount   int64 `json:"rising_talent_count"`
	ExpertVettedCount   int64 `json:"expert_vetted_count"`
	FeaturedCount       int64 `json:"featured_count"`
	
	// Growth statistics
	UsersCreatedToday     int64   `json:"users_created_today"`
	UsersCreatedYesterday int64   `json:"users_created_yesterday"`
	UsersCreatedThisWeek  int64   `json:"users_created_this_week"`
	UsersCreatedLastWeek  int64   `json:"users_created_last_week"`
	UsersCreatedThisMonth int64   `json:"users_created_this_month"`
	UsersCreatedLastMonth int64   `json:"users_created_last_month"`
	GrowthRateDaily       float64 `json:"growth_rate_daily"`
	GrowthRateWeekly      float64 `json:"growth_rate_weekly"`
	GrowthRateMonthly     float64 `json:"growth_rate_monthly"`
	
	// Moderation statistics
	UsersWithWarnings   int64 `json:"users_with_warnings"`
	FlaggedUsers        int64 `json:"flagged_users"`
	TotalWarnings       int64 `json:"total_warnings"`
	TotalSuspensions    int64 `json:"total_suspensions"`
	TotalBans           int64 `json:"total_bans"`
	
	// Referral statistics
	UsersWithReferrals  int64   `json:"users_with_referrals"`
	TotalReferrals      int64   `json:"total_referrals"`
	AverageReferrals    float64 `json:"average_referrals"`
	
	// Financial statistics (cached)
	TotalEarningsAll    float64 `json:"total_earnings_all"`
	TotalSpendingAll    float64 `json:"total_spending_all"`
	AverageEarnings     float64 `json:"average_earnings"`
	AverageSpending     float64 `json:"average_spending"`
	
	// Timestamp
	GeneratedAt time.Time `json:"generated_at"`
}

// UserGrowthStats contains user growth statistics over time
type UserGrowthStats struct {
	Period            string    `json:"period"`
	StartDate         time.Time `json:"start_date"`
	EndDate           time.Time `json:"end_date"`
	
	TotalUsers        int64     `json:"total_users"`
	NewUsers          int64     `json:"new_users"`
	DeletedUsers      int64     `json:"deleted_users"`
	NetGrowth         int64     `json:"net_growth"`
	
	GrowthRate        float64   `json:"growth_rate"`
	ChurnRate         float64   `json:"churn_rate"`
	
	ActiveUsers       int64     `json:"active_users"`
	RetentionRate     float64   `json:"retention_rate"`
	
	NewFreelancers    int64     `json:"new_freelancers"`
	NewClients        int64     `json:"new_clients"`
	
	VerifiedUsers     int64     `json:"verified_users"`
	VerificationRate  float64   `json:"verification_rate"`
}

// CountryStats contains user statistics by country
type CountryStats struct {
	Country        string  `json:"country"`
	CountryCode    string  `json:"country_code"`
	UserCount      int64   `json:"user_count"`
	Percentage     float64 `json:"percentage"`
	FreelancerCount int64  `json:"freelancer_count"`
	ClientCount    int64   `json:"client_count"`
	AverageRating  float64 `json:"average_rating"`
}

// CityStats contains user statistics by city
type CityStats struct {
	City           string  `json:"city"`
	Country        string  `json:"country"`
	CountryCode    string  `json:"country_code"`
	UserCount      int64   `json:"user_count"`
	Percentage     float64 `json:"percentage"`
	FreelancerCount int64  `json:"freelancer_count"`
	ClientCount    int64   `json:"client_count"`
}

// TimezoneStats contains user statistics by timezone
type TimezoneStats struct {
	Timezone    string  `json:"timezone"`
	UserCount   int64   `json:"user_count"`
	Percentage  float64 `json:"percentage"`
	OnlineCount int64   `json:"online_count"`
}

// RatingDistribution contains distribution of user ratings
type RatingDistribution struct {
	Rating0to1 int64   `json:"rating_0_to_1"`
	Rating1to2 int64   `json:"rating_1_to_2"`
	Rating2to3 int64   `json:"rating_2_to_3"`
	Rating3to4 int64   `json:"rating_3_to_4"`
	Rating4to5 int64   `json:"rating_4_to_5"`
	NoRating   int64   `json:"no_rating"`
	Average    float64 `json:"average"`
	Median     float64 `json:"median"`
}

// ProfileCompletenessDistribution contains distribution of profile completeness
type ProfileCompletenessDistribution struct {
	Completeness0to20   int64   `json:"completeness_0_to_20"`
	Completeness20to40  int64   `json:"completeness_20_to_40"`
	Completeness40to60  int64   `json:"completeness_40_to_60"`
	Completeness60to80  int64   `json:"completeness_60_to_80"`
	Completeness80to100 int64   `json:"completeness_80_to_100"`
	Average             float64 `json:"average"`
	Median              float64 `json:"median"`
}

// UserActivityStats contains user activity statistics
type UserActivityStats struct {
	TotalLogins         int64     `json:"total_logins"`
	UniqueLogins        int64     `json:"unique_logins"`
	AverageLoginsPerUser float64  `json:"average_logins_per_user"`
	
	OnlineNow           int64     `json:"online_now"`
	ActiveLast24Hours   int64     `json:"active_last_24_hours"`
	ActiveLast7Days     int64     `json:"active_last_7_days"`
	ActiveLast30Days    int64     `json:"active_last_30_days"`
	
	InactiveLast30Days  int64     `json:"inactive_last_30_days"`
	InactiveLast90Days  int64     `json:"inactive_last_90_days"`
	InactiveLast180Days int64     `json:"inactive_last_180_days"`
	
	AverageSessionDuration float64 `json:"average_session_duration"`
	GeneratedAt            time.Time `json:"generated_at"`
}

// UserTypeDistribution contains distribution by user type
type UserTypeDistribution struct {
	Freelancers int64   `json:"freelancers"`
	Clients     int64   `json:"clients"`
	Both        int64   `json:"both"`
	Staff       int64   `json:"staff"`
	Total       int64   `json:"total"`
	
	FreelancerPercentage float64 `json:"freelancer_percentage"`
	ClientPercentage     float64 `json:"client_percentage"`
	BothPercentage       float64 `json:"both_percentage"`
	StaffPercentage      float64 `json:"staff_percentage"`
}

// StatusDistribution contains distribution by account status
type StatusDistribution struct {
	Active      int64   `json:"active"`
	Pending     int64   `json:"pending"`
	Inactive    int64   `json:"inactive"`
	Suspended   int64   `json:"suspended"`
	Banned      int64   `json:"banned"`
	Deleted     int64   `json:"deleted"`
	Restricted  int64   `json:"restricted"`
	Total       int64   `json:"total"`
	
	ActivePercentage     float64 `json:"active_percentage"`
	PendingPercentage    float64 `json:"pending_percentage"`
	SuspendedPercentage  float64 `json:"suspended_percentage"`
	BannedPercentage     float64 `json:"banned_percentage"`
}

// BadgeDistribution contains distribution of badges
type BadgeDistribution struct {
	BadgeType  BadgeType `json:"badge_type"`
	UserCount  int64     `json:"user_count"`
	Percentage float64   `json:"percentage"`
}

// CohortAnalysis contains cohort analysis data
type CohortAnalysis struct {
	CohortMonth        string    `json:"cohort_month"`
	InitialUsers       int64     `json:"initial_users"`
	RetainedMonth1     int64     `json:"retained_month_1"`
	RetainedMonth3     int64     `json:"retained_month_3"`
	RetainedMonth6     int64     `json:"retained_month_6"`
	RetainedMonth12    int64     `json:"retained_month_12"`
	RetentionRate      float64   `json:"retention_rate"`
	ChurnRate          float64   `json:"churn_rate"`
}

// UserSegment represents a user segment for analytics
type UserSegment struct {
	SegmentName   string                 `json:"segment_name"`
	Description   string                 `json:"description"`
	UserCount     int64                  `json:"user_count"`
	Criteria      map[string]interface{} `json:"criteria"`
	CreatedAt     time.Time              `json:"created_at"`
}

// TimeSeriesDataPoint represents a single data point in time series
type TimeSeriesDataPoint struct {
	Timestamp time.Time              `json:"timestamp"`
	Value     float64                `json:"value"`
	Metrics   map[string]interface{} `json:"metrics,omitempty"`
}

// TimeSeriesData represents time series data
type TimeSeriesData struct {
	Metric     string                 `json:"metric"`
	Interval   string                 `json:"interval"`
	StartDate  time.Time              `json:"start_date"`
	EndDate    time.Time              `json:"end_date"`
	DataPoints []TimeSeriesDataPoint  `json:"data_points"`
}