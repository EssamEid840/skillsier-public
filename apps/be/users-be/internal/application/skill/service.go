// internal/application/skill/service.go
package skill

import (
    "context"
    "fmt"
    "gorm.io/gorm"
    "users-be/internal/domain/skill"
    "users-be/internal/domain/outbox"
)

type Service struct {
    repo       skill.Repository
    outboxRepo outbox.Repository
    db         *gorm.DB
}

func NewService(repo skill.Repository, outboxRepo outbox.Repository, db *gorm.DB) *Service {
    return &Service{
        repo:       repo,
        outboxRepo: outboxRepo,
        db:         db,
    }
}

func (s *Service) AddSkill(ctx context.Context, dto AddSkillDTO) (*SkillDTO, error) {
    existing, _ := s.repo.FindByUserIDAndName(ctx, dto.UserID, dto.SkillName)
    if existing != nil {
        return nil, fmt.Errorf("skill already exists")
    }
    
    sk := &skill.Skill{
        UserID:            dto.UserID,
        SkillName:         dto.SkillName,
        Proficiency:       skill.Proficiency(dto.Proficiency),
        YearsOfExperience: dto.YearsOfExperience,
        LastUsedYear:      dto.LastUsedYear,
        IsPrimary:         dto.IsPrimary,
    }
    
    err := s.db.Transaction(func(tx *gorm.DB) error {
        if err := s.repo.Create(ctx, sk); err != nil {
            return err
        }
        
        event := &outbox.Event{
            AggregateID:   dto.UserID,
            AggregateType: "skill",
            EventType:     "skill.added",
            Payload:       fmt.Sprintf(`{"user_id":"%s","skill_name":"%s"}`, dto.UserID, dto.SkillName),
        }
        
        return s.outboxRepo.Create(ctx, event)
    })
    
    if err != nil {
        return nil, err
    }
    
    return ToSkillDTO(sk), nil
}

func (s *Service) UpdateSkill(ctx context.Context, id string, dto UpdateSkillDTO) (*SkillDTO, error) {
    sk, err := s.repo.FindByID(ctx, id)
    if err != nil {
        return nil, err
    }
    
    if dto.Proficiency != nil {
        sk.Proficiency = skill.Proficiency(*dto.Proficiency)
    }
    if dto.YearsOfExperience != nil {
        sk.YearsOfExperience = *dto.YearsOfExperience
    }
    if dto.LastUsedYear != nil {
        sk.LastUsedYear = *dto.LastUsedYear
    }
    if dto.IsPrimary != nil {
        sk.IsPrimary = *dto.IsPrimary
    }
    if dto.TestScore != nil {
        sk.TestScore = *dto.TestScore
    }
    
    if err := s.repo.Update(ctx, sk); err != nil {
        return nil, err
    }
    
    return ToSkillDTO(sk), nil
}

func (s *Service) GetUserSkills(ctx context.Context, userID string) ([]*SkillDTO, error) {
    skills, err := s.repo.FindByUserID(ctx, userID)
    if err != nil {
        return nil, err
    }
    
    dtos := make([]*SkillDTO, len(skills))
    for i, sk := range skills {
        dtos[i] = ToSkillDTO(sk)
    }
    
    return dtos, nil
}

func (s *Service) GetPrimarySkills(ctx context.Context, userID string) ([]*SkillDTO, error) {
    skills, err := s.repo.FindPrimarySkills(ctx, userID)
    if err != nil {
        return nil, err
    }
    
    dtos := make([]*SkillDTO, len(skills))
    for i, sk := range skills {
        dtos[i] = ToSkillDTO(sk)
    }
    
    return dtos, nil
}

func (s *Service) GetVerifiedSkills(ctx context.Context, userID string) ([]*SkillDTO, error) {
    skills, err := s.repo.FindVerifiedSkills(ctx, userID)
    if err != nil {
        return nil, err
    }
    
    dtos := make([]*SkillDTO, len(skills))
    for i, sk := range skills {
        dtos[i] = ToSkillDTO(sk)
    }
    
    return dtos, nil
}

func (s *Service) DeleteSkill(ctx context.Context, id string) error {
    return s.repo.Delete(ctx, id)
}

func (s *Service) ReorderSkills(ctx context.Context, userID string, dto ReorderSkillsDTO) error {
    return s.repo.UpdateDisplayOrder(ctx, userID, dto.SkillIDs)
}

func (s *Service) EndorseSkill(ctx context.Context, id string) error {
    return s.repo.IncrementEndorsements(ctx, id)
}

func (s *Service) VerifySkill(ctx context.Context, id, verifiedBy string) error {
    sk, err := s.repo.FindByID(ctx, id)
    if err != nil {
        return err
    }
    
    sk.IsVerified = true
    sk.VerifiedBy = verifiedBy
    now := time.Now()
    sk.VerifiedAt = &now
    
    return s.repo.Update(ctx, sk)
}

func (s *Service) GetTopSkills(ctx context.Context, limit int) ([]*SkillDTO, error) {
    skills, err := s.repo.GetTopSkills(ctx, limit)
    if err != nil {
        return nil, err
    }
    
    dtos := make([]*SkillDTO, len(skills))
    for i, sk := range skills {
        dtos[i] = ToSkillDTO(sk)
    }
    
    return dtos, nil
}

func (s *Service) GetSkillStatistics(ctx context.Context, skillName string) (map[string]interface{}, error) {
    return s.repo.GetSkillStats(ctx, skillName)
}

func ToSkillDTO(s *skill.Skill) *SkillDTO {
    if s == nil {
        return nil
    }
    
    return &SkillDTO{
        ID:                s.ID,
        UserID:            s.UserID,
        SkillName:         s.SkillName,
        Proficiency:       string(s.Proficiency),
        YearsOfExperience: s.YearsOfExperience,
        LastUsedYear:      s.LastUsedYear,
        IsPrimary:         s.IsPrimary,
        IsVerified:        s.IsVerified,
        TestScore:         s.TestScore,
        EndorsementCount:  s.EndorsementCount,
        ProjectCount:      s.ProjectCount,
        DisplayOrder:      s.DisplayOrder,
        SkillScore:        s.CalculateScore(),
        CreatedAt:         s.CreatedAt,
        UpdatedAt:         s.UpdatedAt,
    }
}
