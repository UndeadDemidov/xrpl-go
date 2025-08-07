// revive:disable:var-naming
package types

import "errors"

var (
	// ErrNoAccountID is returned when no account ID is specified.
	ErrNoAccountID = errors.New("no account ID specified")
)
