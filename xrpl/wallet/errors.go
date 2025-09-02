package wallet

import "errors"

var (
	// address

	// ErrAddressTagNotZero is returned when the address tag is not zero.
	ErrAddressTagNotZero = errors.New("address tag is not zero")

	// batch

	// ErrBatchAccountNotFound is returned when the batch account is not found in the transaction.
	ErrBatchAccountNotFound = errors.New("batch account not found in transaction")
	// ErrTransactionMustBeBatch is returned when the transaction is not a batch transaction.
	ErrTransactionMustBeBatch = errors.New("transaction must be a batch transaction")
	// ErrNoTransactionsProvided is returned when no transactions are provided.
	ErrNoTransactionsProvided = errors.New("no transactions provided")
	// ErrTxMustIncludeBatchSigner is returned when the transaction does not include a batch signer.
	ErrTxMustIncludeBatchSigner = errors.New("transaction must include a batch signer")
	// ErrTransactionAlreadySigned is returned when the transaction has already been signed.
	ErrTransactionAlreadySigned = errors.New("transaction has already been signed")
	// ErrBatchSignableNotEqual is returned when the batch signable is not equal.
	ErrBatchSignableNotEqual = errors.New("batch signable is not equal")
)
