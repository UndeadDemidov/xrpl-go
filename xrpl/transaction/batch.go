package transaction

import (
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
)

const (
	// Batch transaction flags
	tfAllOrNothing uint32 = 0x00010000
	tfOnlyOne      uint32 = 0x00020000
	tfUntilFailure uint32 = 0x00040000
	tfIndependent  uint32 = 0x00080000
)

// Batch represents a Batch transaction that can execute multiple transactions atomically.
//
// Example:
//
// ```json
//
//	{
//	    "TransactionType": "Batch",
//	    "Account": "rN7n7otQDd6FczFgLdSqtcsAUxDkw6fzRH",
//	    "Fee": "100",
//	    "Flags": 65536,
//	    "Sequence": 1,
//	    "RawTransactions": [
//	        {
//	            "RawTransaction": {
//	                "TransactionType": "Payment",
//	                "Account": "rN7n7otQDd6FczFgLdSqtcsAUxDkw6fzRH",
//	                "Amount": "1000000",
//	                "Destination": "rPT1Sjq2YGrBMTttX4GZHjKu9dyfzbpAYe",
//	                "Flags": 1073741824,
//	                "Fee": "0",
//	                "SigningPubKey": ""
//	            }
//	        }
//	    ]
//	}
//
// ```
type Batch struct {
	BaseTx
	// Array of transactions to be executed as part of this batch.
	RawTransactions []types.RawTransaction `json:"RawTransactions"`
	// Optional array of batch signers for multi-signing the batch.
	BatchSigners []types.BatchSigner `json:"BatchSigners,omitempty"`
}

// TxType returns the type of the transaction (Batch).
func (*Batch) TxType() TxType {
	return BatchTx
}

// **********************************
// Batch Flags
// **********************************

// SetAllOrNothingFlag sets the AllOrNothing flag.
//
// AllOrNothing: Execute all transactions in the batch or none at all.
// If any transaction fails, the entire batch fails.
func (b *Batch) SetAllOrNothingFlag() {
	b.Flags |= tfAllOrNothing
}

// SetOnlyOneFlag sets the OnlyOne flag.
//
// OnlyOne: Execute only the first transaction that succeeds.
// Stop execution after the first successful transaction.
func (b *Batch) SetOnlyOneFlag() {
	b.Flags |= tfOnlyOne
}

// SetUntilFailureFlag sets the UntilFailure flag.
//
// UntilFailure: Execute transactions until one fails.
// Stop execution at the first failed transaction.
func (b *Batch) SetUntilFailureFlag() {
	b.Flags |= tfUntilFailure
}

// SetIndependentFlag sets the Independent flag.
//
// Independent: Execute all transactions independently.
// The failure of one transaction does not affect others.
func (b *Batch) SetIndependentFlag() {
	b.Flags |= tfIndependent
}

// Flatten returns the flattened map of the Batch transaction.
func (b *Batch) Flatten() FlatTransaction {
	flattenedTx := b.BaseTx.Flatten()

	flattenedTx["TransactionType"] = b.TxType().String()

	rawTxs := make([]map[string]any, len(b.RawTransactions))
	for i, rtw := range b.RawTransactions {
		rawTxs[i] = rtw.Flatten()
	}
	flattenedTx["RawTransactions"] = rawTxs

	if len(b.BatchSigners) > 0 {
		signers := make([]map[string]any, len(b.BatchSigners))
		for i, bs := range b.BatchSigners {
			signers[i] = bs.Flatten()
		}
		flattenedTx["BatchSigners"] = signers
	}

	return flattenedTx
}

// Validate validates the Batch transaction.
func (b *Batch) Validate() (bool, error) {
	_, err := b.BaseTx.Validate()
	if err != nil {
		return false, err
	}

	if len(b.RawTransactions) == 0 {
		return false, ErrBatchRawTransactionsEmpty
	}

	// Validate each RawTransaction
	for _, rawTx := range b.RawTransactions {
		if valid, err := rawTx.Validate(); !valid {
			return false, err
		}
	}

	for _, batchSigner := range b.BatchSigners {
		if err := batchSigner.Validate(); err != nil {
			return false, err
		}
	}

	return true, nil
}
