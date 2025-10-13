package certification

import "users-be/internal/domain/certification"

func ToResponseDTO(cert *certification.Certification) *CertificationResponseDTO {
	if cert == nil {
		return nil
	}
	return &CertificationResponseDTO{
		ID:            cert.ID,
		UserID:        cert.UserID,
		Name:          cert.Name,
		Issuer:        cert.Issuer,
		IssueDate:     cert.IssueDate,
		ExpiryDate:    cert.ExpiryDate,
		CredentialID:  cert.CredentialID,
		CredentialURL: cert.CredentialURL,
		IsExpired:     cert.IsExpired(),
		CreatedAt:     cert.CreatedAt,
		UpdatedAt:     cert.UpdatedAt,
	}
}

func ToListResponse(certs []*certification.Certification) *ListCertificationsResponseDTO {
	dtos := make([]*CertificationResponseDTO, len(certs))
	for i, cert := range certs {
		dtos[i] = ToResponseDTO(cert)
	}
	return &ListCertificationsResponseDTO{
		Certifications: dtos,
		Total:          len(certs),
	}
}