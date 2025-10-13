// internal/domain/profile/repository.go
package profile

import "context"

type Repository interface {
    // Create
    Create(ctx context.Context, profile *Profile) error
    CreateBatch(ctx context.Context, profiles []*Profile) error
    
    // Read
    FindByID(ctx context.Context, id string) (*Profile, error)
    FindByUserID(ctx context.Context, userID string) (*Profile, error)
    FindByUserIDs(ctx context.Context, userIDs []string) ([]*Profile, error)
    List(ctx context.Context, filter ListFilter) ([]*Profile, int64, error)
    Search(ctx context.Context, query string, filter ListFilter) ([]*Profile, int64, error)
    
    // Read - Business queries
    FindAvailableForWork(ctx context.Context, limit int) ([]*Profile, error)
    FindByLocation(ctx context.Context, country, city string) ([]*Profile, error)
    FindByRateRange(ctx context.Context, minRate, maxRate float64) ([]*Profile, error)
    FindFeatured(ctx context.Context, limit int) ([]*Profile, error)
    FindRecentlyUpdated(ctx context.Context, limit int) ([]*Profile, error)
    FindIncomplete(ctx context.Context) ([]*Profile, error)
    FindWithLowQuality(ctx context.Context, threshold float64) ([]*Profile, error)
    
    // Update
    Update(ctx context.Context, profile *Profile) error
    UpdateCompletionPercentage(ctx context.Context, userID string, percentage int) error
    UpdateQualityScore(ctx context.Context, userID string, score float64) error
    IncrementViews(ctx context.Context, userID string) error
    UpdateAvailabilityStatus(ctx context.Context, userID, status string) error
    
    // Delete
    Delete(ctx context.Context, id string) error
    SoftDelete(ctx context.Context, id string) error
    
    // Analytics
    CountByCountry(ctx context.Context, country string) (int64, error)
    GetAverageQualityScore(ctx context.Context) (float64, error)
    GetAverageCompletionRate(ctx context.Context) (float64, error)
}

type ListFilter struct {
    Page                int
    PageSize            int
    SortBy              string
    SortOrder           string
    Country             *string
    City                *string
    MinRate             *float64
    MaxRate             *float64
    AvailabilityStatus  *string
    MinYearsExperience  *int
    Industry            *string
    IsPublic            *bool
    MinQualityScore     *float64
    MinCompletion       *int
}
