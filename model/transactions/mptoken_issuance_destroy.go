package transactions

type MPTokenIssuanceDestroy struct {
	BaseTx
	MPTokenIssuanceID string `json:",omitempty"`
}

func (*MPTokenIssuanceDestroy) TxType() TxType {
	return MPTokenIssuanceDestroyTx
}
