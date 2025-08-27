package types

import "errors"

var (
	// batch

	// ErrBatchSignableInvalid is returned when the batch signable is invalid.
	ErrBatchSignableInvalid = errors.New("batch signable is invalid")

	// fields

	// ErrFlagsFieldIsNotAnUint32 is returned when the flags field is not an uint32.
	ErrFlagsFieldIsNotAnUint32 = errors.New("flags field is not an uint32")
	// ErrRawTransactionsFieldIsNotAnArray is returned when the raw transactions field is not an array.
	ErrRawTransactionsFieldIsNotAnArray = errors.New("raw transactions field is not an array")
	// ErrRawTransactionFieldIsNotAnObject is returned when the raw transaction field is not an object.
	ErrRawTransactionFieldIsNotAnObject = errors.New("raw transaction field is not an object")

	// transaction id

	// ErrFailedToGetTxIDFromRawTransaction is returned when getting txID from raw transaction fails.
	ErrFailedToGetTxIDFromRawTransaction = errors.New("failed to get txID from raw transaction")
)
