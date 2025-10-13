// internal/domain/certification/errors.go
package certification

import "errors"

var (
    ErrCertificationNotFound  = errors.New("certification not found")
    ErrAlreadyVerified        = errors.New("certification already verified")
    ErrExpired                = errors.New("certification has expired")
    ErrInvalidCredential      = errors.New("invalid credential")
    ErrDuplicateCertification = errors.New("certification already exists")
)