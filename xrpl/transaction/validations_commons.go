package transaction

import (
	"fmt"

	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
)

// ValidateRequiredField validates that a required field exists in the transaction and passes the validity check.
func ValidateRequiredField(tx FlatTransaction, field string, checkValidity func(interface{}) bool) error {
	// Check if the field is present in the transaction map.
	if _, ok := tx[field]; !ok {
		return fmt.Errorf("%w: %s", ErrFieldMissing, field)
	}

	// Check if the field is valid.
	if !checkValidity(tx[field]) {
		transactionType, _ := tx["TransactionType"].(string)
		return fmt.Errorf("%w: %s invalid field %s", ErrInvalidField, transactionType, field)
	}

	return nil
}

// ValidateOptionalField validates an optional field in the transaction map.
func ValidateOptionalField(tx FlatTransaction, paramName string, checkValidity func(interface{}) bool) error {
	// Check if the field is present in the transaction map.
	if value, ok := tx[paramName]; ok {
		// Check if the field is valid.
		if !checkValidity(value) {
			transactionType, _ := tx["TransactionType"].(string)
			return fmt.Errorf("%w: %s invalid field %s", ErrInvalidField, transactionType, paramName)
		}
	}

	return nil
}

// validateMemos validates the Memos field in the transaction map.
func validateMemos(memoWrapper []types.MemoWrapper) error {
	// loop through each memo and validate it
	for _, memo := range memoWrapper {
		isMemo, err := IsMemo(memo.Memo)
		if !isMemo {
			return err
		}
	}

	return nil
}

// validateSigners validates the Signers field in the transaction map.
func validateSigners(signers []types.Signer) error {
	// loop through each signer and validate it
	for _, signer := range signers {
		isSigner, err := IsSigner(signer.SignerData)
		if !isSigner {
			return err
		}
	}

	return nil
}
