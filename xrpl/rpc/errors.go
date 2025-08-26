package rpc

import "errors"

// Static errors
var (
	// ErrIncorrectID is returned when the request ID doesn't match the response ID.
	ErrIncorrectID = errors.New("request ID does not match response ID")

	// ErrMissingTxSignatureOrSigningPubKey is returned when a transaction lacks both TxSignature and SigningPubKey.
	ErrMissingTxSignatureOrSigningPubKey = errors.New("transaction must include either TxSignature or SigningPubKey")

	// ErrSignerDataIsEmpty is returned when signer data is empty or missing.
	ErrSignerDataIsEmpty = errors.New("signer data must not be empty")

	// ErrCannotFundWalletWithoutClassicAddress is returned when attempting to fund a wallet without a classic address.
	ErrCannotFundWalletWithoutClassicAddress = errors.New("cannot fund wallet without a classic address")

	// ErrMissingLastLedgerSequenceInTransaction is returned when LastLedgerSequence is missing from a transaction.
	ErrMissingLastLedgerSequenceInTransaction = errors.New("missing LastLedgerSequence in transaction")

	// ErrMissingWallet is returned when a wallet is required but not provided for an unsigned transaction.
	ErrMissingWallet = errors.New("wallet must be provided when submitting an unsigned transaction")

	// ErrRawTransactionsFieldIsNotAnArray is returned when the RawTransactions field is not an array type.
	ErrRawTransactionsFieldIsNotAnArray = errors.New("field RawTransactions must be an array")

	// ErrRawTransactionFieldIsNotAnObject is returned when the RawTransaction field is not an object type.
	ErrRawTransactionFieldIsNotAnObject = errors.New("field RawTransaction must be an object")

	// ErrSigningPubKeyFieldMustBeEmpty is returned when the signingPubKey field should be empty but isn't.
	ErrSigningPubKeyFieldMustBeEmpty = errors.New("field SigningPubKey must be empty")

	// ErrTxnSignatureFieldMustBeEmpty is returned when the txnSignature field should be empty but isn't.
	ErrTxnSignatureFieldMustBeEmpty = errors.New("field TxnSignature must be empty")

	// ErrSignersFieldMustBeEmpty is returned when the signers field should be empty but isn't.
	ErrSignersFieldMustBeEmpty = errors.New("field Signers must be empty")

	// ErrAccountFieldIsNotAString is returned when the account field is not a string type.
	ErrAccountFieldIsNotAString = errors.New("field Account must be a string")
)

// Dynamic errors

// ClientError represents a dynamic error with a custom error message string from the RPC client.
type ClientError struct {
	ErrorString string
}

// Error returns the error message string for ClientError.
func (e *ClientError) Error() string {
	return e.ErrorString
}
