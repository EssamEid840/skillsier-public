// internal/application/education/dto.go
package education

import "time"

type EducationDTO struct {
    ID              string     `json:"id"`
    UserID          string     `json:"user_id"`
    School          string     `json:"school"`
    SchoolLogo      string     `json:"school_logo"`
    Degree          string     `json:"degree"`
    DegreeType      string     `json:"degree_type"`
    Field           string     `json:"field"`
    Grade           string     `json:"grade"`
    GPA             float64    `json:"gpa"`
    MaxGPA          float64    `json:"max_gpa"`
    GraduationYear  int        `json:"graduation_year"`
    IsCurrent       bool       `json:"is_current"`
    Description     string     `json:"description"`
    Activities      string     `json:"activities"`
    Location        string     `json:"location"`
    IsVerified      bool       `json:"is_verified"`
    CertificateURL  string     `json:"certificate_url"`
    DisplayOrder    int        `json:"display_order"`
    CreatedAt       time.Time  `json:"created_at"`
    UpdatedAt       time.Time  `json:"updated_at"`
}

type AddEducationDTO struct {
    UserID         string  `json:"user_id" binding:"required"`
    School         string  `json:"school" binding:"required,min=2,max=200"`
    Degree         string  `json:"degree" binding:"required,min=2,max=200"`
    DegreeType     string  `json:"degree_type" binding:"required,oneof=associate bachelor master doctorate certificate"`
    Field          string  `json:"field" binding:"required,min=2,max=200"`
    Grade          string  `json:"grade"`
    GPA            float64 `json:"gpa" binding:"omitempty,min=0,max=10"`
    MaxGPA         float64 `json:"max_gpa" binding:"omitempty,min=0,max=10"`
    GraduationYear int     `json:"graduation_year" binding:"required,min=1950,max=2100"`
    IsCurrent      bool    `json:"is_current"`
    Description    string  `json:"description"`
    Activities     string  `json:"activities"`
    Location       string  `json:"location"`
}

type UpdateEducationDTO struct {
    School         *string  `json:"school" binding:"omitempty,min=2,max=200"`
    Degree         *string  `json:"degree" binding:"omitempty,min=2,max=200"`
    DegreeType     *string  `json:"degree_type" binding:"omitempty,oneof=associate bachelor master doctorate certificate"`
    Field          *string  `json:"field" binding:"omitempty,min=2,max=200"`
    Grade          *string  `json:"grade"`
    GPA            *float64 `json:"gpa" binding:"omitempty,min=0,max=10"`
    MaxGPA         *float64 `json:"max_gpa" binding:"omitempty,min=0,max=10"`
    GraduationYear *int     `json:"graduation_year" binding:"omitempty,min=1950,max=2100"`
    IsCurrent      *bool    `json:"is_current"`
    Description    *string  `json:"description"`
    Activities     *string  `json:"activities"`
    Location       *string  `json:"location"`
    CertificateURL *string  `json:"certificate_url"`
}

type ReorderEducationsDTO struct {
    EducationIDs []string `json:"education_ids" binding:"required,min=1"`
}