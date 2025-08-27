package types

// Holder returns a pointer to an Address representing the Holder field (optional).
func Holder(address Address) *Address {
	return &address
}
