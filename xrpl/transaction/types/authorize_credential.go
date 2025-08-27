package types

import (
	"errors"
)

// MaxAcceptedCredentials is the maximum number of accepted credentials.
const MaxAcceptedCredentials int = 10

var (
	// Credential-specific errors.

	// ErrInvalidCredentialType is returned when the credential type is invalid; it must be a hexadecimal string between 1 and 64 bytes.
	ErrInvalidCredentialType = errors.New("invalid credential type, must be a hexadecimal string between 1 and 64 bytes")

	// ErrInvalidCredentialIssuer is returned when the credential Issuer field is missing.
	ErrInvalidCredentialIssuer = errors.New("credential type: missing field Issuer")
)

// AuthorizeCredential represents an accepted credential for PermissionedDomainSet transactions.
type AuthorizeCredential struct {
	Credential Credential
}

// Validate checks if the AuthorizeCredential is valid.
func (a AuthorizeCredential) Validate() error {
	if a.Credential.Issuer.String() == "" {
		return ErrInvalidCredentialIssuer
	}
	if !a.Credential.CredentialType.IsValid() {
		return ErrInvalidCredentialType
	}
	return nil
}

// Flatten returns a flattened map representation of the AuthorizeCredential.
func (a AuthorizeCredential) Flatten() map[string]interface{} {
	m := make(map[string]interface{})
	m["Credential"] = a.Credential.Flatten()
	return m
}
