// internal/domain/education/errors.go
package education

import "errors"

var (
    ErrEducationNotFound      = errors.New("education not found")
    ErrInvalidYear            = errors.New("invalid graduation year")
    ErrInvalidGPA             = errors.New("invalid GPA")
    ErrDuplicateEducation     = errors.New("education already exists")
)
