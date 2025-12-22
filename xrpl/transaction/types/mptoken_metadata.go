//revive:disable:var-naming
package types

// MPTokenMetadata returns a pointer to a string containing metadata for an MPToken (optional).
func MPTokenMetadata(value string) *string {
	return &value
}
