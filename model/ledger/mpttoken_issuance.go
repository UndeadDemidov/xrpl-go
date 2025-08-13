package ledger

import "github.com/CreatureDev/xrpl-go/model/transactions/types"

type MPTokenIssuance struct {
	LedgerEntryType   LedgerEntryType `json:",omitempty"`
	Flags             *types.Flag     `json:",omitempty"`
	Issuer            types.Address   `json:",omitempty"`
	AssetScale        uint8           `json:",omitempty"`
	MaximumAmount     string          `json:",omitempty"`
	OutstandingAmount string          `json:",omitempty"`
	TransferFee       uint16          `json:",omitempty"`
	MPTokenMetadata   string          `json:",omitempty"`
	OwnerNode         string          `json:",omitempty"`
	LockedAmount      string          `json:",omitempty"`
	PreviousTxnID     types.Hash256   `json:",omitempty"`
	PreviousTxnLgrSeq uint32          `json:",omitempty"`
	Index             types.Hash256   `json:"index,omitempty"`
}

func (*MPTokenIssuance) EntryType() LedgerEntryType {
	return MPTokenIssuanceEntry
}
