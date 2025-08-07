package transaction

var _ Tx = (*FlatTransaction)(nil)

// FlatTransaction is a flattened transaction represented as a map from field names to interface{} values.
// It satisfies the Tx interface for generic transaction handling.
type FlatTransaction map[string]interface{}

// TxType returns the transaction type of the flattened transaction.
func (f FlatTransaction) TxType() TxType {
	txType, ok := f["TransactionType"].(string)
	if !ok {
		return TxType("")
	}
	return TxType(txType)
}
