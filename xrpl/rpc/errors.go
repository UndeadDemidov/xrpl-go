package rpc

import "errors"

var (
	// transaction

	// ErrMissingTxSignatureOrSigningPubKey is returned when a transaction lacks both TxSignature and SigningPubKey.
	ErrMissingTxSignatureOrSigningPubKey = errors.New("transaction must include either TxSignature or SigningPubKey")
	// ErrSignerDataIsEmpty is returned when signer data is empty or missing.
	ErrSignerDataIsEmpty = errors.New("signer data must not be empty")
	// ErrMissingLastLedgerSequenceInTransaction is returned when LastLedgerSequence is missing from a transaction.
	ErrMissingLastLedgerSequenceInTransaction = errors.New("missing LastLedgerSequence in transaction")
	// ErrMissingWallet is returned when a wallet is required but not provided for an unsigned transaction.
	ErrMissingWallet = errors.New("wallet must be provided when submitting an unsigned transaction")
	// ErrMissingAccountInTransaction is returned when the Account field is missing from a transaction.
	ErrMissingAccountInTransaction = errors.New("missing Account in transaction")
	// ErrTransactionTypeMissing is returned when the transaction type is missing from a transaction.
	ErrTransactionTypeMissing = errors.New("transaction type is missing in transaction")
	// ErrTransactionNotFound is returned when a transaction cannot be found.
	ErrTransactionNotFound = errors.New("transaction not found")
	// ErrInvalidFulfillmentLength is returned when the fulfillment length is invalid.
	ErrInvalidFulfillmentLength = errors.New("invalid fulfillment length")
	// ErrMismatchedTag is returned when a transaction tag field does not match the expected value.
	ErrMismatchedTag = errors.New("transaction tag mismatch")

	// fields

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
	// ErrRawTransactionsFieldMissing is returned when the RawTransactions field is missing from a Batch transaction.
	ErrRawTransactionsFieldMissing = errors.New("RawTransactions field missing from Batch transaction")
	// ErrRawTransactionFieldMissing is returned when the RawTransaction field is missing from a wrapper.
	ErrRawTransactionFieldMissing = errors.New("RawTransaction field missing from wrapper")
	// ErrFeeFieldMissing is returned when the fee field is missing after calculation.
	ErrFeeFieldMissing = errors.New("fee field missing after calculation")

	// wallet

	// ErrCannotFundWalletWithoutClassicAddress is returned when attempting to fund a wallet without a classic address.
	ErrCannotFundWalletWithoutClassicAddress = errors.New("cannot fund wallet without a classic address")

	// fees

	// ErrCouldNotGetBaseFeeXrp is returned when BaseFeeXrp cannot be retrieved from ServerInfo.
	ErrCouldNotGetBaseFeeXrp = errors.New("get fee xrp: could not get BaseFeeXrp from ServerInfo")
	// ErrCouldNotFetchOwnerReserve is returned when the owner reserve fee cannot be fetched.
	ErrCouldNotFetchOwnerReserve = errors.New("could not fetch Owner Reserve")
	// ErrFailedToParseFee is returned when fee parsing fails.
	ErrFailedToParseFee = errors.New("failed to parse fee")

	// account

	// ErrAccountCannotBeDeleted is returned when an account cannot be deleted due to associated objects.
	ErrAccountCannotBeDeleted = errors.New("account cannot be deleted; there are Escrows, PayChannels, RippleStates, or Checks associated with the account")

	// payment

	// ErrAmountAndDeliverMaxMustBeIdentical is returned when Amount and DeliverMax fields are not identical.
	ErrAmountAndDeliverMaxMustBeIdentical = errors.New("payment transaction: Amount and DeliverMax fields must be identical when both are provided")

	// tags

	// ErrTagMustEqualAddressTag is returned when a tag must equal the address tag.
	ErrTagMustEqualAddressTag = errors.New("tag, if present, must be equal to the tag of the address")

	// json rpc

	// ErrFailedToMarshalJSONRPCRequest is returned when JSON-RPC request marshaling fails.
	ErrFailedToMarshalJSONRPCRequest = errors.New("failed to marshal JSON-RPC request")

	// config

	// ErrEmptyURL is returned when the provided URL is empty (no port or IP specified).
	ErrEmptyURL = errors.New("empty port and IP provided")
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
