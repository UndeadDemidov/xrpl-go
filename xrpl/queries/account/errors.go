package account

import "errors"

var (
	// ErrNoAccountID is returned when no account ID is specified in a request.
	ErrNoAccountID = errors.New("no account ID specified")
)
