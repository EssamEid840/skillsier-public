// internal/application/education/service.go
package education

import (
    "context"
    "fmt"
    "time"
    "gorm.io/gorm"
    "users-be/internal/domain/education"
    "users-be/internal/domain/outbox"
)

type Service struct {
    repo       education.Repository
    outboxRepo outbox.Repository
    db         *gorm.DB
}

func NewService(repo education.Repository, outboxRepo outbox.Repository, db *gorm.DB) *Service {
    return &Service{
        repo:       repo,
        outboxRepo: outboxRepo,
        db:         db,
    }
}

func (s *Service) AddEducation(ctx context.Context, dto AddEducationDTO) (*EducationDTO, error) {
    if dto.GPA > dto.MaxGPA {
        return nil, education.ErrInvalidGPA
    }
    
    edu := &education.Education{
        UserID:         dto.UserID,
        School:         dto.School,
        Degree:         dto.Degree,
        DegreeType:     dto.DegreeType,
        Field:          dto.Field,
        Grade:          dto.Grade,
        GPA:            dto.GPA,
        MaxGPA:         dto.MaxGPA,
        GraduationYear: dto.GraduationYear,
        IsCurrent:      dto.IsCurrent,
        Description:    dto.Description,
        Activities:     dto.Activities,
        Location:       dto.Location,
    }
    
    err := s.db.Transaction(func(tx *gorm.DB) error {
        if err := s.repo.Create(ctx, edu); err != nil {
            return err
        }
        
        event := &outbox.Event{
            AggregateID:   dto.UserID,
            AggregateType: "education",
            EventType:     "education.added",
            Payload:       fmt.Sprintf(`{"user_id":"%s","school":"%s","degree":"%s"}`, dto.UserID, dto.School, dto.Degree),
        }
        
        return s.outboxRepo.Create(ctx, event)
    })
    
    if err != nil {
        return nil, err
    }
    
    return ToEducationDTO(edu), nil
}

func (s *Service) UpdateEducation(ctx context.Context, id string, dto UpdateEducationDTO) (*EducationDTO, error) {
    edu, err := s.repo.FindByID(ctx, id)
    if err != nil {
        return nil, err
    }
    
    if dto.School != nil {
        edu.School = *dto.School
    }
    if dto.Degree != nil {
        edu.Degree = *dto.Degree
    }
    if dto.DegreeType != nil {
        edu.DegreeType = *dto.DegreeType
    }
    if dto.Field != nil {
        edu.Field = *dto.Field
    }
    if dto.Grade != nil {
        edu.Grade = *dto.Grade
    }
    if dto.GPA != nil {
        edu.GPA = *dto.GPA
    }
    if dto.MaxGPA != nil {
        edu.MaxGPA = *dto.MaxGPA
    }
    if dto.GraduationYear != nil {
        edu.GraduationYear = *dto.GraduationYear
    }
    if dto.IsCurrent != nil {
        edu.IsCurrent = *dto.IsCurrent
    }
    if dto.Description != nil {
        edu.Description = *dto.Description
    }
    if dto.Activities != nil {
        edu.Activities = *dto.Activities
    }
    if dto.Location != nil {
        edu.Location = *dto.Location
    }
    if dto.CertificateURL != nil {
        edu.CertificateURL = *dto.CertificateURL
    }
    
    if err := s.repo.Update(ctx, edu); err != nil {
        return nil, err
    }
    
    return ToEducationDTO(edu), nil
}

func (s *Service) GetUserEducations(ctx context.Context, userID string) ([]*EducationDTO, error) {
    educations, err := s.repo.FindByUserID(ctx, userID)
    if err != nil {
        return nil, err
    }
    
    dtos := make([]*EducationDTO, len(educations))
    for i, edu := range educations {
        dtos[i] = ToEducationDTO(edu)
    }
    
    return dtos, nil
}

func (s *Service) GetVerifiedEducations(ctx context.Context, userID string) ([]*EducationDTO, error) {
    educations, err := s.repo.FindVerified(ctx, userID)
    if err != nil {
        return nil, err
    }
    
    dtos := make([]*EducationDTO, len(educations))
    for i, edu := range educations {
        dtos[i] = ToEducationDTO(edu)
    }
    
    return dtos, nil
}

func (s *Service) GetCurrentEducations(ctx context.Context, userID string) ([]*EducationDTO, error) {
    educations, err := s.repo.FindCurrent(ctx, userID)
    if err != nil {
        return nil, err
    }
    
    dtos := make([]*EducationDTO, len(educations))
    for i, edu := range educations {
        dtos[i] = ToEducationDTO(edu)
    }
    
    return dtos, nil
}

func (s *Service) GetHighestDegree(ctx context.Context, userID string) (*EducationDTO, error) {
    edu, err := s.repo.GetHighestDegree(ctx, userID)
    if err != nil {
        return nil, err
    }
    return ToEducationDTO(edu), nil
}

func (s *Service) DeleteEducation(ctx context.Context, id string) error {
    return s.repo.Delete(ctx, id)
}

func (s *Service) ReorderEducations(ctx context.Context, userID string, dto ReorderEducationsDTO) error {
    return s.repo.UpdateDisplayOrder(ctx, userID, dto.EducationIDs)
}

func (s *Service) VerifyEducation(ctx context.Context, id, verifiedBy string) error {
    edu, err := s.repo.FindByID(ctx, id)
    if err != nil {
        return err
    }
    
    edu.IsVerified = true
    edu.VerifiedBy = verifiedBy
    now := time.Now()
    edu.VerifiedAt = &now
    
    return s.repo.Update(ctx, edu)
}

func ToEducationDTO(e *education.Education) *EducationDTO {
    if e == nil {
        return nil
    }
    
    return &EducationDTO{
        ID:             e.ID,
        UserID:         e.UserID,
        School:         e.School,
        SchoolLogo:     e.SchoolLogo,
        Degree:         e.Degree,
        DegreeType:     e.DegreeType,
        Field:          e.Field,
        Grade:          e.Grade,
        GPA:            e.GPA,
        MaxGPA:         e.MaxGPA,
        GraduationYear: e.GraduationYear,
        IsCurrent:      e.IsCurrent,
        Description:    e.Description,
        Activities:     e.Activities,
        Location:       e.Location,
        IsVerified:     e.IsVerified,
        CertificateURL: e.CertificateURL,
        DisplayOrder:   e.DisplayOrder,
        CreatedAt:      e.CreatedAt,
        UpdatedAt:      e.UpdatedAt,
    }
}
