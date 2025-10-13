// internal/application/profile/dto.go
package profile

import "time"

type ProfileDTO struct {
    ID                   string     `json:"id"`
    UserID               string     `json:"user_id"`
    Title                string     `json:"title"`
    Bio                  string     `json:"bio"`
    Tagline              string     `json:"tagline"`
    Overview             string     `json:"overview"`
    Location             string     `json:"location"`
    City                 string     `json:"city"`
    State                string     `json:"state"`
    Country              string     `json:"country"`
    ProfilePictureURL    string     `json:"profile_picture_url"`
    CoverImageURL        string     `json:"cover_image_url"`
    VideoIntroURL        string     `json:"video_intro_url"`
    WebsiteURL           string     `json:"website_url"`
    LinkedInURL          string     `json:"linkedin_url"`
    GithubURL            string     `json:"github_url"`
    TwitterURL           string     `json:"twitter_url"`
    YearsOfExperience    int        `json:"years_of_experience"`
    EducationLevel       string     `json:"education_level"`
    Industry             string     `json:"industry"`
    Specialization       string     `json:"specialization"`
    AvailabilityStatus   string     `json:"availability_status"`
    HoursPerWeek         int        `json:"hours_per_week"`
    HourlyRate           float64    `json:"hourly_rate"`
    MinimumBudget        float64    `json:"minimum_budget"`
    Currency             string     `json:"currency"`
    ProfileViews         int        `json:"profile_views"`
    CompletionPercentage int        `json:"completion_percentage"`
    QualityScore         float64    `json:"quality_score"`
    IsPublic             bool       `json:"is_public"`
    ShowRates            bool       `json:"show_rates"`
    CreatedAt            time.Time  `json:"created_at"`
    UpdatedAt            time.Time  `json:"updated_at"`
}

type CreateProfileDTO struct {
    UserID             string  `json:"user_id" binding:"required"`
    Title              string  `json:"title" binding:"required,min=10,max=200"`
    Bio                string  `json:"bio" binding:"required,min=50,max=5000"`
    Tagline            string  `json:"tagline" binding:"max=150"`
    Location           string  `json:"location"`
    City               string  `json:"city"`
    Country            string  `json:"country" binding:"required,len=2"`
    HourlyRate         float64 `json:"hourly_rate" binding:"min=0"`
    AvailabilityStatus string  `json:"availability_status" binding:"oneof=available busy not_available"`
}

type UpdateProfileDTO struct {
    Title                *string  `json:"title" binding:"omitempty,min=10,max=200"`
    Bio                  *string  `json:"bio" binding:"omitempty,min=50,max=5000"`
    Tagline              *string  `json:"tagline" binding:"omitempty,max=150"`
    Overview             *string  `json:"overview"`
    Location             *string  `json:"location"`
    City                 *string  `json:"city"`
    State                *string  `json:"state"`
    Country              *string  `json:"country" binding:"omitempty,len=2"`
    ProfilePictureURL    *string  `json:"profile_picture_url"`
    CoverImageURL        *string  `json:"cover_image_url"`
    VideoIntroURL        *string  `json:"video_intro_url"`
    WebsiteURL           *string  `json:"website_url"`
    LinkedInURL          *string  `json:"linkedin_url"`
    GithubURL            *string  `json:"github_url"`
    TwitterURL           *string  `json:"twitter_url"`
    YearsOfExperience    *int     `json:"years_of_experience"`
    EducationLevel       *string  `json:"education_level"`
    Industry             *string  `json:"industry"`
    Specialization       *string  `json:"specialization"`
    AvailabilityStatus   *string  `json:"availability_status"`
    HoursPerWeek         *int     `json:"hours_per_week"`
    HourlyRate           *float64 `json:"hourly_rate"`
    MinimumBudget        *float64 `json:"minimum_budget"`
    ShowRates            *bool    `json:"show_rates"`
    IsPublic             *bool    `json:"is_public"`
}
