package xrpl

import "errors"

var (
	// ErrNoTxToMultisign is returned when no transaction blobs are provided to Multisign.
	ErrNoTxToMultisign = errors.New("no transaction to multisign")
)
