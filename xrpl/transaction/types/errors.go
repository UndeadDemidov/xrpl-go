package types

import "errors"

var (
	// xchain bridge

	// ErrInvalidIssuingChainDoorAddress is returned when the issuing chain door address is invalid.
	ErrInvalidIssuingChainDoorAddress = errors.New("xchain bridge: invalid issuing chain door address")
	// ErrInvalidIssuingChainIssueAddress is returned when the issuing chain issue address is invalid.
	ErrInvalidIssuingChainIssueAddress = errors.New("xchain bridge: invalid issuing chain issue address")
	// ErrInvalidLockingChainDoorAddress is returned when the locking chain door address is invalid.
	ErrInvalidLockingChainDoorAddress = errors.New("xchain bridge: invalid locking chain door address")
	// ErrInvalidLockingChainIssueAddress is returned when the locking chain issue address is invalid.
	ErrInvalidLockingChainIssueAddress = errors.New("xchain bridge: invalid locking chain issue address")

	// raw tx

	// ErrBatchRawTransactionMissing is returned when the RawTransaction field is missing from an array element.
	ErrBatchRawTransactionMissing = errors.New("batch RawTransaction field is missing")
	// ErrBatchRawTransactionFieldNotObject is returned when the RawTransaction field is not an object.
	ErrBatchRawTransactionFieldNotObject = errors.New("batch RawTransaction field is not an object")
	// ErrBatchNestedTransaction is returned when trying to include a Batch transaction within another Batch.
	ErrBatchNestedTransaction = errors.New("batch cannot contain nested Batch transactions")
	// ErrBatchMissingInnerFlag is returned when an inner transaction lacks the TfInnerBatchTxn flag.
	ErrBatchMissingInnerFlag = errors.New("batch RawTransaction must contain the TfInnerBatchTxn flag")
	// ErrBatchInnerTransactionInvalid is returned when an inner transaction fails its own validation.
	ErrBatchInnerTransactionInvalid = errors.New("batch inner transaction validation failed")

	// permission

	// ErrInvalidPermissionValue is returned when PermissionValue is empty or undefined.
	ErrInvalidPermissionValue = errors.New("permission value cannot be empty or undefined")

	// batch signer

	// ErrBatchSignerAccountMissing is returned when a BatchSigner lacks the required Account field.
	ErrBatchSignerAccountMissing = errors.New("batch BatchSigner Account is missing")
	// ErrBatchSignerSigningPubKeyMissing is returned when a BatchSigner lacks the required SigningPubKey field.
	ErrBatchSignerSigningPubKeyMissing = errors.New("batch BatchSigner SigningPubKey is missing")
	// ErrBatchSignerInvalidTxnSignature is returned when a BatchSigner has an invalid TxnSignature field.
	ErrBatchSignerInvalidTxnSignature = errors.New("batch BatchSigner TxnSignature is invalid")

	// credential

	// ErrInvalidCredentialType is returned when the credential type is invalid; it must be a hexadecimal string between 1 and 64 bytes.
	ErrInvalidCredentialType = errors.New("invalid credential type, must be a hexadecimal string between 1 and 64 bytes")
	// ErrInvalidCredentialIssuer is returned when the credential Issuer field is missing.
	ErrInvalidCredentialIssuer = errors.New("credential type: missing field Issuer")

	// ErrEmptyCredentials is returned when the credential list is empty.
	ErrEmptyCredentials = errors.New("credentials list cannot be empty")
	// ErrInvalidCredentialCount is returned when the credential list size is out of allowed range.
	ErrInvalidCredentialCount = errors.New("accepted credentials list must contain at least one and no more than the maximum allowed number of items")
	// ErrDuplicateCredentials is returned when duplicate credentials are present in the list.
	ErrDuplicateCredentials = errors.New("credentials list cannot contain duplicate elements")
)
