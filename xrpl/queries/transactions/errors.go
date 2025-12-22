package transactions

import "errors"

var (
	// ErrNoTxBlob is returned when no TxBlob is defined in the SubmitRequest.
	ErrNoTxBlob = errors.New("no TxBlob defined")
)
