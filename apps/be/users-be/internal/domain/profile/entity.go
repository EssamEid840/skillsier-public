// internal/domain/profile/entity.go
package profile

import (
    "time"
    "gorm.io/gorm"
)

type Profile struct {
    // Identifiers
    ID                   string         `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
    UserID               string         `gorm:"uniqueIndex;not null;type:uuid"`
    
    // Basic Profile Info
    Title                string         `gorm:"type:varchar(200)"` // Professional title
    Bio                  string         `gorm:"type:text"`
    Tagline              string         `gorm:"type:varchar(150)"` // Short catchy line
    Overview             string         `gorm:"type:text"` // Detailed overview
    
    // Location & Contact
    Location             string         `gorm:"type:varchar(200)"`
    City                 string         `gorm:"type:varchar(100)"`
    State                string         `gorm:"type:varchar(100)"`
    Country              string         `gorm:"type:varchar(2)"` // ISO code
    PostalCode           string         `gorm:"type:varchar(20)"`
    Address              string         `gorm:"type:text"`
    
    // Links & Media
    ProfilePictureURL    string         `gorm:"type:varchar(500)"`
    CoverImageURL        string         `gorm:"type:varchar(500)"`
    VideoIntroURL        string         `gorm:"type:varchar(500)"`
    VideoIntroThumbnail  string         `gorm:"type:varchar(500)"`
    
    // Social Media Links
    WebsiteURL           string         `gorm:"type:varchar(500)"`
    LinkedInURL          string         `gorm:"type:varchar(500)"`
    GithubURL            string         `gorm:"type:varchar(500)"`
    TwitterURL           string         `gorm:"type:varchar(500)"`
    FacebookURL          string         `gorm:"type:varchar(500)"`
    InstagramURL         string         `gorm:"type:varchar(500)"`
    BehanceURL           string         `gorm:"type:varchar(500)"`
    DribbbleURL          string         `gorm:"type:varchar(500)"`
    
    // Professional Details
    YearsOfExperience    int            `gorm:"default:0"`
    EducationLevel       string         `gorm:"type:varchar(50)"` // high_school, bachelors, masters, phd
    Industry             string         `gorm:"type:varchar(100)"`
    Specialization       string         `gorm:"type:varchar(200)"`
    
    // Work Preferences (for freelancers)
    AvailabilityStatus   string         `gorm:"type:varchar(30);default:'available'"` // available, busy, not_available
    HoursPerWeek         int            `gorm:"default:0"`
    PreferredProjectSize string         `gorm:"type:varchar(30)"` // small, medium, large, any
    PreferredProjectLength string       `gorm:"type:varchar(30)"` // short, medium, long, any
    WorkingHoursStart    string         `gorm:"type:varchar(10)"` // "09:00"
    WorkingHoursEnd      string         `gorm:"type:varchar(10)"` // "17:00"
    
    // Rates & Budget (for freelancers)
    HourlyRate           float64        `gorm:"type:decimal(10,2);default:0"`
    MinimumBudget        float64        `gorm:"type:decimal(10,2);default:0"`
    MaximumBudget        float64        `gorm:"type:decimal(10,2);default:0"`
    Currency             string         `gorm:"type:varchar(3);default:'USD'"`
    
    // Hiring Budget (for clients)
    TypicalBudgetRange   string         `gorm:"type:varchar(50)"` // <$100, $100-$500, $500-$1k, etc.
    TotalBudgetSpent     float64        `gorm:"type:decimal(15,2);default:0"`
    AverageProjectBudget float64        `gorm:"type:decimal(10,2);default:0"`
    
    // Profile Metrics
    ProfileViews         int            `gorm:"default:0"`
    ProfileViewsThisWeek int            `gorm:"default:0"`
    ProfileViewsThisMonth int           `gorm:"default:0"`
    SearchAppearances    int            `gorm:"default:0"`
    ContactClicks        int            `gorm:"default:0"`
    HireClicks           int            `gorm:"default:0"`
    
    // Completion Tracking
    CompletionPercentage int            `gorm:"default:0;check:completion_percentage >= 0 AND completion_percentage <= 100"`
    MissingFields        string         `gorm:"type:jsonb"` // Array of missing field names
    LastCompletedSection string         `gorm:"type:varchar(50)"`
    
    // SEO & Searchability
    MetaTitle            string         `gorm:"type:varchar(200)"`
    MetaDescription      string         `gorm:"type:varchar(500)"`
    Keywords             string         `gorm:"type:text"` // Comma-separated
    SearchTags           string         `gorm:"type:jsonb"` // JSON array of tags
    
    // Profile Quality Score (internal)
    QualityScore         float64        `gorm:"type:decimal(5,2);default:0"` // 0-100
    LastQualityCheck     *time.Time
    QualityIssues        string         `gorm:"type:jsonb"` // JSON array of issues
    
    // Visibility & Privacy
    IsPublic             bool           `gorm:"default:true"`
    ShowEmail            bool           `gorm:"default:false"`
    ShowPhone            bool           `gorm:"default:false"`
    ShowLocation         bool           `gorm:"default:true"`
    ShowRates            bool           `gorm:"default:true"`
    SearchableProfile    bool           `gorm:"default:true"`
    
    // Verification Badges
    VideoIntroVerified   bool           `gorm:"default:false"`
    PortfolioVerified    bool           `gorm:"default:false"`
    ReferencesVerified   bool           `gorm:"default:false"`
    
    // Timestamps
    CreatedAt            time.Time
    UpdatedAt            time.Time
    DeletedAt            gorm.DeletedAt `gorm:"index"`
    LastViewedAt         *time.Time
    LastEditedAt         *time.Time
}

func (Profile) TableName() string {
    return "profiles"
}

// Business methods
func (p *Profile) IsComplete() bool {
    return p.CompletionPercentage >= 80
}

func (p *Profile) NeedsImprovement() bool {
    return p.QualityScore < 70
}

func (p *Profile) IsAvailableForWork() bool {
    return p.AvailabilityStatus == "available" && p.IsPublic
}