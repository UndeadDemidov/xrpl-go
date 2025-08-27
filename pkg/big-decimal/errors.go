package bigdecimal

import (
	"errors"
)

var (
	// ErrInvalidCharacter is returned when a string contains disallowed characters.
	ErrInvalidCharacter = errors.New("value contains invalid characters")
	// ErrInvalidZeroValue indicates the value string represents zero or is invalid zero.
	ErrInvalidZeroValue = errors.New("value cannot be zero")
)
