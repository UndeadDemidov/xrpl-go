package websocket

import "errors"

// Static errors
var (
	// ErrMissingTxSignatureOrSigningPubKey is returned when a transaction lacks both TxSignature and SigningPubKey.
	ErrMissingTxSignatureOrSigningPubKey = errors.New("transaction must have a TxSignature or SigningPubKey set")
	// ErrMissingLastLedgerSequenceInTransaction is returned when LastLedgerSequence is missing from a transaction.
	ErrMissingLastLedgerSequenceInTransaction = errors.New("missing LastLedgerSequence in transaction")
	// ErrMissingWallet is returned when a wallet is required but not provided for an unsigned transaction.
	ErrMissingWallet = errors.New("wallet must be provided when submitting an unsigned transaction")

	// ErrRawTransactionsFieldIsNotAnArray is returned when the RawTransactions field is not an array type.
	ErrRawTransactionsFieldIsNotAnArray = errors.New("RawTransactions field is not an array")
	// ErrRawTransactionFieldIsNotAnObject is returned when the RawTransaction field is not an object type.
	ErrRawTransactionFieldIsNotAnObject = errors.New("RawTransaction field is not an object")

	// ErrSigningPubKeyFieldMustBeEmpty is returned when the signingPubKey field should be empty but isn't.
	ErrSigningPubKeyFieldMustBeEmpty = errors.New("signingPubKey field must be empty")
	// ErrTxnSignatureFieldMustBeEmpty is returned when the txnSignature field should be empty but isn't.
	ErrTxnSignatureFieldMustBeEmpty = errors.New("txnSignature field must be empty")
	// ErrSignersFieldMustBeEmpty is returned when the signers field should be empty but isn't.
	ErrSignersFieldMustBeEmpty = errors.New("signers field must be empty")
	// ErrAccountFieldIsNotAString is returned when the account field is not a string type.
	ErrAccountFieldIsNotAString = errors.New("account field is not a string")
)

// Dynamic errors

// ClientError represents a dynamic error with a custom error message string.
type ClientError struct {
	ErrorString string
}

// Error returns the error message string for ClientError.
func (e *ClientError) Error() string {
	return e.ErrorString
}
