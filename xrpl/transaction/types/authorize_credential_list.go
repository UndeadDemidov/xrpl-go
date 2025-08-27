package types

import "errors"

var (
	// ErrEmptyCredentials is returned when the credential list is empty.
	ErrEmptyCredentials = errors.New("credentials list cannot be empty")
	// ErrInvalidCredentialCount is returned when the credential list size is out of allowed range.
	ErrInvalidCredentialCount = errors.New("accepted credentials list must contain at least one and no more than the maximum allowed number of items")
	// ErrDuplicateCredentials is returned when duplicate credentials are present in the list.
	ErrDuplicateCredentials = errors.New("credentials list cannot contain duplicate elements")
)

// AuthorizeCredentialList represents a list of AuthorizeCredential entries with validation and flattening.
type AuthorizeCredentialList []AuthorizeCredential

// Validate checks that the list is non-empty, within allowed size, has no duplicates, and each credential is valid.
func (ac *AuthorizeCredentialList) Validate() error {
	if len(*ac) == 0 {
		return ErrEmptyCredentials
	}
	if len(*ac) > MaxAcceptedCredentials {
		return ErrInvalidCredentialCount
	}
	seen := make(map[string]bool)
	for _, cred := range *ac {
		key := cred.Credential.Issuer.String() + cred.Credential.CredentialType.String()
		if seen[key] {
			return ErrDuplicateCredentials
		}
		seen[key] = true

		if err := cred.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Flatten returns a slice of maps representing each AuthorizeCredential in JSON-like format.
func (ac *AuthorizeCredentialList) Flatten() []map[string]interface{} {
	acs := make([]map[string]interface{}, len(*ac))
	for i, c := range *ac {
		acs[i] = c.Flatten()
	}
	return acs
}
