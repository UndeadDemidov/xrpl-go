package ledger

import (
	"testing"

	"github.com/CreatureDev/xrpl-go/model/transactions/types"
	"github.com/CreatureDev/xrpl-go/test"
)

func TestMPToken(t *testing.T) {
	var s LedgerObject = &MPToken{
		MPTAmount: types.MPTCurrencyAmount{
			MPTIssuanceID: "MPT1234567890ABCDEF",
			Value:         "1000",
		},
		Flags:             types.SetFlag(types.TfMPTCanTrade | types.TfMPTCanTransfer),
		MPTokenIssuanceID: "MPT1234567890ABCDEF",
		LedgerEntryType:   MPTokenEntry,
		LockedAmount:      "100",
		OwnerNode:         "0000000000000000",
		PreviousTxnID:     "F0AB71E777B2DA54B86231E19B82554EF1F8211F92ECA473121C655BFC5329BF",
		PreviousTxnLgrSeq: 14524914,
		Index:             "F0AB71E777B2DA54B86231E19B82554EF1F8211F92ECA473121C655BFC5329BF",
	}

	j := `{
	"MPTAmount": {
		"mp_issuance_id": "MPT1234567890ABCDEF",
		"value": "1000"
	},
	"Flags": 48,
	"mpt_issuance_id": "MPT1234567890ABCDEF",
	"LedgerEntryType": "MPToken",
	"LockedAmount": "100",
	"OwnerNode": "0000000000000000",
	"PreviousTxnID": "F0AB71E777B2DA54B86231E19B82554EF1F8211F92ECA473121C655BFC5329BF",
	"PreviousTxnLgrSeq": 14524914,
	"index": "F0AB71E777B2DA54B86231E19B82554EF1F8211F92ECA473121C655BFC5329BF"
}`

	if err := test.SerializeAndDeserialize(t, s, j); err != nil {
		t.Error(err)
	}
}

func TestMPToken_EntryType(t *testing.T) {
	mpt := &MPToken{}
	if mpt.EntryType() != MPTokenEntry {
		t.Errorf("Expected EntryType to be %s, got %s", MPTokenEntry, mpt.EntryType())
	}
}
