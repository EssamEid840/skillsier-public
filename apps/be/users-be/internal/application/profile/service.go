// internal/application/profile/service.go
package profile

import (
    "context"
    "fmt"
    "gorm.io/gorm"
    "users-be/internal/domain/profile"
    "users-be/internal/domain/outbox"
    "users-be/internal/infrastructure/cache/redis"
)

type Service struct {
    repo       profile.Repository
    outboxRepo outbox.Repository
    db         *gorm.DB
    cache      *redis.ProfileCache
}

func NewService(repo profile.Repository, outboxRepo outbox.Repository, db *gorm.DB, cache *redis.ProfileCache) *Service {
    return &Service{
        repo:       repo,
        outboxRepo: outboxRepo,
        db:         db,
        cache:      cache,
    }
}

// CREATE
func (s *Service) CreateProfile(ctx context.Context, dto CreateProfileDTO) (*ProfileDTO, error) {
    // Check if profile already exists
    existing, _ := s.repo.FindByUserID(ctx, dto.UserID)
    if existing != nil {
        return nil, profile.ErrProfileAlreadyExists
    }
    
    p := &profile.Profile{
        UserID:             dto.UserID,
        Title:              dto.Title,
        Bio:                dto.Bio,
        Tagline:            dto.Tagline,
        Location:           dto.Location,
        City:               dto.City,
        Country:            dto.Country,
        HourlyRate:         dto.HourlyRate,
        AvailabilityStatus: dto.AvailabilityStatus,
        IsPublic:           true,
        ShowRates:          true,
        SearchableProfile:  true,
    }
    
    // Calculate initial completion
    p.CompletionPercentage = s.calculateCompletion(p)
    
    err := s.db.Transaction(func(tx *gorm.DB) error {
        if err := s.repo.Create(ctx, p); err != nil {
            return err
        }
        
        event := &outbox.Event{
            AggregateID:   p.UserID,
            AggregateType: "profile",
            EventType:     "profile.created",
            Payload:       s.buildEventPayload(p, "created"),
        }
        
        return s.outboxRepo.Create(ctx, event)
    })
    
    if err != nil {
        return nil, err
    }
    
    _ = s.cache.Set(ctx, p)
    return ToProfileDTO(p), nil
}

// READ
func (s *Service) GetProfile(ctx context.Context, userID string) (*ProfileDTO, error) {
    // Try cache
    p, err := s.cache.Get(ctx, userID)
    if err == nil && p != nil {
        return ToProfileDTO(p), nil
    }
    
    p, err = s.repo.FindByUserID(ctx, userID)
    if err != nil {
        return nil, err
    }
    
    _ = s.cache.Set(ctx, p)
    return ToProfileDTO(p), nil
}

func (s *Service) GetProfileByID(ctx context.Context, id string) (*ProfileDTO, error) {
    p, err := s.repo.FindByID(ctx, id)
    if err != nil {
        return nil, err
    }
    return ToProfileDTO(p), nil
}

func (s *Service) ListProfiles(ctx context.Context, filter profile.ListFilter) ([]*ProfileDTO, int64, error) {
    profiles, total, err := s.repo.List(ctx, filter)
    if err != nil {
        return nil, 0, err
    }
    
    dtos := make([]*ProfileDTO, len(profiles))
    for i, p := range profiles {
        dtos[i] = ToProfileDTO(p)
    }
    
    return dtos, total, nil
}

func (s *Service) SearchProfiles(ctx context.Context, query string, filter profile.ListFilter) ([]*ProfileDTO, int64, error) {
    profiles, total, err := s.repo.Search(ctx, query, filter)
    if err != nil {
        return nil, 0, err
    }
    
    dtos := make([]*ProfileDTO, len(profiles))
    for i, p := range profiles {
        dtos[i] = ToProfileDTO(p)
    }
    
    return dtos, total, nil
}

func (s *Service) GetAvailableProfiles(ctx context.Context, limit int) ([]*ProfileDTO, error) {
    profiles, err := s.repo.FindAvailableForWork(ctx, limit)
    if err != nil {
        return nil, err
    }
    
    dtos := make([]*ProfileDTO, len(profiles))
    for i, p := range profiles {
        dtos[i] = ToProfileDTO(p)
    }
    
    return dtos, nil
}

func (s *Service) GetProfilesByLocation(ctx context.Context, country, city string) ([]*ProfileDTO, error) {
    profiles, err := s.repo.FindByLocation(ctx, country, city)
    if err != nil {
        return nil, err
    }
    
    dtos := make([]*ProfileDTO, len(profiles))
    for i, p := range profiles {
        dtos[i] = ToProfileDTO(p)
    }
    
    return dtos, nil
}

func (s *Service) GetProfilesByRateRange(ctx context.Context, minRate, maxRate float64) ([]*ProfileDTO, error) {
    profiles, err := s.repo.FindByRateRange(ctx, minRate, maxRate)
    if err != nil {
        return nil, err
    }
    
    dtos := make([]*ProfileDTO, len(profiles))
    for i, p := range profiles {
        dtos[i] = ToProfileDTO(p)
    }
    
    return dtos, nil
}

// UPDATE
func (s *Service) UpdateProfile(ctx context.Context, userID string, dto UpdateProfileDTO) (*ProfileDTO, error) {
    p, err := s.repo.FindByUserID(ctx, userID)
    if err != nil {
        return nil, err
    }
    
    updated := false
    
    if dto.Title != nil && *dto.Title != p.Title {
        p.Title = *dto.Title
        updated = true
    }
    
    if dto.Bio != nil && *dto.Bio != p.Bio {
        p.Bio = *dto.Bio
        updated = true
    }
    
    if dto.Tagline != nil && *dto.Tagline != p.Tagline {
        p.Tagline = *dto.Tagline
        updated = true
    }
    
    if dto.Overview != nil && *dto.Overview != p.Overview {
        p.Overview = *dto.Overview
        updated = true
    }
    
    if dto.Location != nil && *dto.Location != p.Location {
        p.Location = *dto.Location
        updated = true
    }
    
    if dto.City != nil && *dto.City != p.City {
        p.City = *dto.City
        updated = true
    }
    
    if dto.State != nil && *dto.State != p.State {
        p.State = *dto.State
        updated = true
    }
    
    if dto.Country != nil && *dto.Country != p.Country {
        p.Country = *dto.Country
        updated = true
    }
    
    if dto.ProfilePictureURL != nil && *dto.ProfilePictureURL != p.ProfilePictureURL {
        p.ProfilePictureURL = *dto.ProfilePictureURL
        updated = true
    }
    
    if dto.CoverImageURL != nil && *dto.CoverImageURL != p.CoverImageURL {
        p.CoverImageURL = *dto.CoverImageURL
        updated = true
    }
    
    if dto.VideoIntroURL != nil && *dto.VideoIntroURL != p.VideoIntroURL {
        p.VideoIntroURL = *dto.VideoIntroURL
        updated = true
    }
    
    if dto.WebsiteURL != nil && *dto.WebsiteURL != p.WebsiteURL {
        p.WebsiteURL = *dto.WebsiteURL
        updated = true
    }
    
    if dto.LinkedInURL != nil && *dto.LinkedInURL != p.LinkedInURL {
        p.LinkedInURL = *dto.LinkedInURL
        updated = true
    }
    
    if dto.GithubURL != nil && *dto.GithubURL != p.GithubURL {
        p.GithubURL = *dto.GithubURL
        updated = true
    }
    
    if dto.TwitterURL != nil && *dto.TwitterURL != p.TwitterURL {
        p.TwitterURL = *dto.TwitterURL
        updated = true
    }
    
    if dto.YearsOfExperience != nil && *dto.YearsOfExperience != p.YearsOfExperience {
        p.YearsOfExperience = *dto.YearsOfExperience
        updated = true
    }
    
    if dto.EducationLevel != nil && *dto.EducationLevel != p.EducationLevel {
        p.EducationLevel = *dto.EducationLevel
        updated = true
    }
    
    if dto.Industry != nil && *dto.Industry != p.Industry {
        p.Industry = *dto.Industry
        updated = true
    }
    
    if dto.Specialization != nil && *dto.Specialization != p.Specialization {
        p.Specialization = *dto.Specialization
        updated = true
    }
    
    if dto.AvailabilityStatus != nil && *dto.AvailabilityStatus != p.AvailabilityStatus {
        p.AvailabilityStatus = *dto.AvailabilityStatus
        updated = true
    }
    
    if dto.HoursPerWeek != nil && *dto.HoursPerWeek != p.HoursPerWeek {
        p.HoursPerWeek = *dto.HoursPerWeek
        updated = true
    }
    
    if dto.HourlyRate != nil && *dto.HourlyRate != p.HourlyRate {
        p.HourlyRate = *dto.HourlyRate
        updated = true
    }
    
    if dto.MinimumBudget != nil && *dto.MinimumBudget != p.MinimumBudget {
        p.MinimumBudget = *dto.MinimumBudget
        updated = true
    }
    
    if dto.ShowRates != nil && *dto.ShowRates != p.ShowRates {
        p.ShowRates = *dto.ShowRates
        updated = true
    }
    
    if dto.IsPublic != nil && *dto.IsPublic != p.IsPublic {
        p.IsPublic = *dto.IsPublic
        updated = true
    }
    
    if !updated {
        return ToProfileDTO(p), nil
    }
    
    // Recalculate completion
    p.CompletionPercentage = s.calculateCompletion(p)
    
    err = s.db.Transaction(func(tx *gorm.DB) error {
        if err := s.repo.Update(ctx, p); err != nil {
            return err
        }
        
        event := &outbox.Event{
            AggregateID:   p.UserID,
            AggregateType: "profile",
            EventType:     "profile.updated",
            Payload:       s.buildEventPayload(p, "updated"),
        }
        
        return s.outboxRepo.Create(ctx, event)
    })
    
    if err != nil {
        return nil, err
    }
    
    _ = s.cache.Invalidate(ctx, userID)
    
    return ToProfileDTO(p), nil
}

func (s *Service) UpdateAvailabilityStatus(ctx context.Context, userID, status string) error {
    if err := s.repo.UpdateAvailabilityStatus(ctx, userID, status); err != nil {
        return err
    }
    
    _ = s.cache.Invalidate(ctx, userID)
    return nil
}

func (s *Service) IncrementProfileViews(ctx context.Context, userID string) error {
    if err := s.repo.IncrementViews(ctx, userID); err != nil {
        return err
    }
    
    _ = s.cache.Invalidate(ctx, userID)
    return nil
}

// DELETE
func (s *Service) DeleteProfile(ctx context.Context, userID string) error {
    p, err := s.repo.FindByUserID(ctx, userID)
    if err != nil {
        return err
    }
    
    if err := s.repo.SoftDelete(ctx, p.ID); err != nil {
        return err
    }
    
    _ = s.cache.Invalidate(ctx, userID)
    return nil
}

// ANALYTICS
func (s *Service) GetProfileStatistics(ctx context.Context) (map[string]interface{}, error) {
    avgQuality, err := s.repo.GetAverageQualityScore(ctx)
    if err != nil {
        return nil, err
    }
    
    avgCompletion, err := s.repo.GetAverageCompletionRate(ctx)
    if err != nil {
        return nil, err
    }
    
    return map[string]interface{}{
        "average_quality_score":      avgQuality,
        "average_completion_rate":    avgCompletion,
    }, nil
}

func (s *Service) GetIncompleteProfiles(ctx context.Context) ([]*ProfileDTO, error) {
    profiles, err := s.repo.FindIncomplete(ctx)
    if err != nil {
        return nil, err
    }
    
    dtos := make([]*ProfileDTO, len(profiles))
    for i, p := range profiles {
        dtos[i] = ToProfileDTO(p)
    }
    
    return dtos, nil
}

func (s *Service) GetLowQualityProfiles(ctx context.Context, threshold float64) ([]*ProfileDTO, error) {
    profiles, err := s.repo.FindWithLowQuality(ctx, threshold)
    if err != nil {
        return nil, err
    }
    
    dtos := make([]*ProfileDTO, len(profiles))
    for i, p := range profiles {
        dtos[i] = ToProfileDTO(p)
    }
    
    return dtos, nil
}

// HELPERS
func (s *Service) calculateCompletion(p *profile.Profile) int {
    completion := 0
    totalFields := 20
    
    if p.Title != "" {
        completion += 5
    }
    if p.Bio != "" && len(p.Bio) >= 50 {
        completion += 10
    }
    if p.Tagline != "" {
        completion += 3
    }
    if p.Overview != "" {
        completion += 5
    }
    if p.ProfilePictureURL != "" {
        completion += 10
    }
    if p.CoverImageURL != "" {
        completion += 5
    }
    if p.VideoIntroURL != "" {
        completion += 10
    }
    if p.Location != "" {
        completion += 3
    }
    if p.Country != "" {
        completion += 2
    }
    if p.WebsiteURL != "" || p.LinkedInURL != "" {
        completion += 5
    }
    if p.YearsOfExperience > 0 {
        completion += 3
    }
    if p.EducationLevel != "" {
        completion += 3
    }
    if p.Industry != "" {
        completion += 3
    }
    if p.Specialization != "" {
        completion += 5
    }
    if p.HourlyRate > 0 {
        completion += 5
    }
    if p.HoursPerWeek > 0 {
        completion += 3
    }
    if p.AvailabilityStatus != "" {
        completion += 3
    }
    
    // Skills, experience, education, portfolio are checked separately
    // These would add additional 20-30% completion
    
    if completion > 100 {
        completion = 100
    }
    
    return completion
}

func (s *Service) buildEventPayload(p *profile.Profile, eventType string) string {
    return fmt.Sprintf(`{
        "user_id": "%s",
        "profile_id": "%s",
        "completion_percentage": %d,
        "quality_score": %.2f,
        "event_type": "%s"
    }`, p.UserID, p.ID, p.CompletionPercentage, p.QualityScore, eventType)
}

// MAPPER
func ToProfileDTO(p *profile.Profile) *ProfileDTO {
    if p == nil {
        return nil
    }
    
    return &ProfileDTO{
        ID:                   p.ID,
        UserID:               p.UserID,
        Title:                p.Title,
        Bio:                  p.Bio,
        Tagline:              p.Tagline,
        Overview:             p.Overview,
        Location:             p.Location,
        City:                 p.City,
        State:                p.State,
        Country:              p.Country,
        ProfilePictureURL:    p.ProfilePictureURL,
        CoverImageURL:        p.CoverImageURL,
        VideoIntroURL:        p.VideoIntroURL,
        WebsiteURL:           p.WebsiteURL,
        LinkedInURL:          p.LinkedInURL,
        GithubURL:            p.GithubURL,
        TwitterURL:           p.TwitterURL,
        YearsOfExperience:    p.YearsOfExperience,
        EducationLevel:       p.EducationLevel,
        Industry:             p.Industry,
        Specialization:       p.Specialization,
        AvailabilityStatus:   p.AvailabilityStatus,
        HoursPerWeek:         p.HoursPerWeek,
        HourlyRate:           p.HourlyRate,
        MinimumBudget:        p.MinimumBudget,
        Currency:             p.Currency,
        ProfileViews:         p.ProfileViews,
        CompletionPercentage: p.CompletionPercentage,
        QualityScore:         p.QualityScore,
        IsPublic:             p.IsPublic,
        ShowRates:            p.ShowRates,
        CreatedAt:            p.CreatedAt,
        UpdatedAt:            p.UpdatedAt,
    }
}
