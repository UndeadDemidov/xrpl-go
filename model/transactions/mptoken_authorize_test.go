package transactions

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/CreatureDev/xrpl-go/model/transactions/types"
	"github.com/CreatureDev/xrpl-go/test"
)

func TestMPTokenAuthorizeTx(t *testing.T) {
	s := MPTokenAuthorize{
		BaseTx: BaseTx{
			Account:            "rBqb89MRQJnMPq8wTwEbtz4kvxrEDfcYvt",
			TransactionType:    MPTokenAuthorizeTx,
			Fee:                types.XRPCurrencyAmount(12),
			Sequence:           8,
			LastLedgerSequence: 7108682,
		},
		MPTokenIssuanceID: "MPT1234567890ABCDEF",
		Holder:            "rN7n7otQDd6FczFgLdSqtcsAUxDkw6fzRH",
		Flags:             types.SetFlag(types.TfMPTUnauthorize),
	}

	j := `{
	"Account": "rBqb89MRQJnMPq8wTwEbtz4kvxrEDfcYvt",
	"TransactionType": "MPTokenAuthorize",
	"Fee": "12",
	"Sequence": 8,
	"LastLedgerSequence": 7108682,
	"MPTokenIssuanceID": "MPT1234567890ABCDEF",
	"Holder": "rN7n7otQDd6FczFgLdSqtcsAUxDkw6fzRH",
	"Flags": 1
}`
	if err := test.SerializeAndDeserialize(t, s, j); err != nil {
		t.Error(err)
	}

	tx, err := UnmarshalTx(json.RawMessage(j))
	if err != nil {
		t.Errorf("UnmarshalTx error: %s", err.Error())
	}
	if !reflect.DeepEqual(tx, &s) {
		t.Error("UnmarshalTx result differs from expected")
	}
}
