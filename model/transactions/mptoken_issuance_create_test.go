package transactions

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/CreatureDev/xrpl-go/model/transactions/types"
	"github.com/CreatureDev/xrpl-go/test"
)

func TestMPTokenIssuanceCreateTx(t *testing.T) {
	s := MPTokenIssuanceCreate{
		BaseTx: BaseTx{
			Account:            "rBqb89MRQJnMPq8wTwEbtz4kvxrEDfcYvt",
			TransactionType:    MPTokenIssuanceCreateTx,
			Fee:                types.XRPCurrencyAmount(12),
			Sequence:           8,
			LastLedgerSequence: 7108682,
		},
		AssetScale:      6,
		MaximumAmount:   "1000000",
		TransferFee:     1000,
		MPTokenMetadata: "7B227469636B6572223A2254455354222C226E616D65223A225465737420546F6B656E222C2261737365745F636C617373223A2264656669222C2261737365745F737562636C617373223A22737461626C65636F696E227D",
		Flags:           types.SetFlag(types.TfMPTCanTrade | types.TfMPTCanTransfer | types.TfMPTCanEscrow),
	}

	j := `{
	"Account": "rBqb89MRQJnMPq8wTwEbtz4kvxrEDfcYvt",
	"TransactionType": "MPTokenIssuanceCreate",
	"Fee": "12",
	"Sequence": 8,
	"LastLedgerSequence": 7108682,
	"AssetScale": 6,
	"MaximumAmount": "1000000",
	"TransferFee": 1000,
	"MPTokenMetadata": "7B227469636B6572223A2254455354222C226E616D65223A225465737420546F6B656E222C2261737365745F636C617373223A2264656669222C2261737365745F737562636C617373223A22737461626C65636F696E227D",
	"Flags": 56
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
