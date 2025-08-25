package hash

import "errors"

var (
	// ErrNonSignedTransaction indicates that a transaction lacks the required signature fields.
	ErrNonSignedTransaction = errors.New("transaction must have at least one of TxnSignature, Signers, or SigningPubKey")
)
