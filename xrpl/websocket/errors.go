package websocket

import "errors"

// Static errors
var (
	// ErrMissingTxSignatureOrSigningPubKey is returned when a transaction lacks both TxSignature and SigningPubKey.
	ErrMissingTxSignatureOrSigningPubKey = errors.New("transaction must include either TxSignature or SigningPubKey")

	// ErrMissingLastLedgerSequenceInTransaction is returned when LastLedgerSequence is missing from a transaction.
	ErrMissingLastLedgerSequenceInTransaction = errors.New("missing LastLedgerSequence in transaction")

	// ErrMissingWallet is returned when a wallet is required but not provided for an unsigned transaction.
	ErrMissingWallet = errors.New("wallet must be provided when submitting an unsigned transaction")

	// ErrRawTransactionsFieldIsNotAnArray is returned when the RawTransactions field is not an array type.
	ErrRawTransactionsFieldIsNotAnArray = errors.New("field RawTransactions must be an array")

	// ErrRawTransactionFieldIsNotAnObject is returned when the RawTransaction field is not an object type.
	ErrRawTransactionFieldIsNotAnObject = errors.New("field RawTransaction must be an object")

	// ErrSigningPubKeyFieldMustBeEmpty is returned when the SigningPubKey field should be empty but isn't.
	ErrSigningPubKeyFieldMustBeEmpty = errors.New("field SigningPubKey must be empty")

	// ErrTxnSignatureFieldMustBeEmpty is returned when the TxnSignature field should be empty but isn't.
	ErrTxnSignatureFieldMustBeEmpty = errors.New("field TxnSignature must be empty")

	// ErrSignersFieldMustBeEmpty is returned when the Signers field should be empty but isn't.
	ErrSignersFieldMustBeEmpty = errors.New("field Signers must be empty")

	// ErrAccountFieldIsNotAString is returned when the Account field is not a string type.
	ErrAccountFieldIsNotAString = errors.New("field Account must be a string")

	// client

	// ErrIncorrectID indicates that a response contains an incorrect request ID.
	ErrIncorrectID = errors.New("incorrect id")
	// ErrNotConnectedToServer indicates that the client is not connected to a WebSocket server.
	ErrNotConnectedToServer = errors.New("not connected to server")
	// ErrRequestTimedOut indicates that a request to the server timed out.
	ErrRequestTimedOut = errors.New("request timed out")
	// ErrCannotFundWalletWithoutClassicAddress is returned when attempting to fund a wallet without a classic address.
	ErrCannotFundWalletWithoutClassicAddress = errors.New("fund wallet: cannot fund a wallet without a classic address")
	// ErrSignerDataIsEmpty is returned when signer data is empty or missing.
	ErrSignerDataIsEmpty = errors.New("signer data is empty")
	// ErrTransactionNotFound is returned when a transaction cannot be found.
	ErrTransactionNotFound = errors.New("transaction not found")
	// ErrMissingAccountInTransaction is returned when the Account field is missing from a transaction.
	ErrMissingAccountInTransaction = errors.New("missing Account in transaction")
	// ErrCouldNotGetBaseFeeXrp is returned when BaseFeeXrp cannot be retrieved from ServerInfo.
	ErrCouldNotGetBaseFeeXrp = errors.New("getFeeXrp: could not get BaseFeeXrp from ServerInfo")
	// ErrAccountCannotBeDeleted is returned when an account cannot be deleted due to associated objects.
	ErrAccountCannotBeDeleted = errors.New("account cannot be deleted; there are Escrows, PayChannels, RippleStates, or Checks associated with the account")
	// ErrAmountAndDeliverMaxMustBeIdentical is returned when Amount and DeliverMax fields are not identical.
	ErrAmountAndDeliverMaxMustBeIdentical = errors.New("payment transaction: Amount and DeliverMax fields must be identical when both are provided")
	// ErrTransactionTypeMissing is returned when the transaction type is missing from a transaction.
	ErrTransactionTypeMissing = errors.New("transaction type is missing in transaction")
	// ErrCouldNotFetchOwnerReserve is returned when the owner reserve fee cannot be fetched.
	ErrCouldNotFetchOwnerReserve = errors.New("could not fetch Owner Reserve")
	// ErrRawTransactionsFieldMissing is returned when the RawTransactions field is missing from a Batch transaction.
	ErrRawTransactionsFieldMissing = errors.New("RawTransactions field missing from Batch transaction")
	// ErrRawTransactionFieldMissing is returned when the RawTransaction field is missing from a wrapper.
	ErrRawTransactionFieldMissing = errors.New("RawTransaction field missing from wrapper")
	// ErrFeeFieldMissing is returned when the fee field is missing after calculation.
	ErrFeeFieldMissing = errors.New("fee field missing after calculation")
	// ErrInvalidFulfillmentLength is returned when the fulfillment length is invalid.
	ErrInvalidFulfillmentLength = errors.New("invalid fulfillment length")
	// ErrTagMustEqualAddressTag is returned when a tag must equal the address tag.
	ErrTagMustEqualAddressTag = errors.New("tag, if present, must be equal to the tag of the address")
	// ErrFailedToParseFee is returned when fee parsing fails.
	ErrFailedToParseFee = errors.New("failed to parse fee")
	// ErrUnknownStreamType is returned when an unknown stream type is encountered.
	ErrUnknownStreamType = errors.New("unknown stream type")
	// ErrMaxReconnectionAttemptsReached is returned when maximum reconnection attempts are reached.
	ErrMaxReconnectionAttemptsReached = errors.New("max reconnection attempts reached")

	// connection

	// ErrNotConnected is returned when attempting to perform operations on a connection that is not established.
	ErrNotConnected = errors.New("connection is not connected")
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
