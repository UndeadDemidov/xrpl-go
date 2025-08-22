package ledger

import (
	"encoding/json"
	"testing"

	"github.com/CreatureDev/xrpl-go/model/transactions/types"
	"github.com/CreatureDev/xrpl-go/test"
)

func TestMPTokenIssuance(t *testing.T) {
	var s LedgerObject = &MPTokenIssuance{
		LedgerEntryType:   MPTokenIssuanceEntry,
		Flags:             types.SetFlag(types.TfMPTCanTrade | types.TfMPTCanTransfer | types.TfMPTCanEscrow),
		Issuer:            "rBqb89MRQJnMPq8wTwEbtz4kvxrEDfcYvt",
		AssetScale:        6,
		MaximumAmount:     "1000000",
		OutstandingAmount: "500000",
		TransferFee:       1000,
		MPTokenMetadata:   "7B227469636B6572223A2254455354222C226E616D65223A225465737420546F6B656E222C2261737365745F636C617373223A2264656669222C2261737365745F737562636C617373223A22737461626C65636F696E227D",
		OwnerNode:         "0000000000000000",
		LockedAmount:      "10000",
		PreviousTxnID:     "F0AB71E777B2DA54B86231E19B82554EF1F8211F92ECA473121C655BFC5329BF",
		PreviousTxnLgrSeq: 14524914,
		Index:             "F0AB71E777B2DA54B86231E19B82554EF1F8211F92ECA473121C655BFC5329BF",
	}

	j := `{
	"LedgerEntryType": "MPTokenIssuance",
	"Flags": 56,
	"Issuer": "rBqb89MRQJnMPq8wTwEbtz4kvxrEDfcYvt",
	"AssetScale": 6,
	"MaximumAmount": "1000000",
	"OutstandingAmount": "500000",
	"TransferFee": 1000,
	"MPTokenMetadata": "7B227469636B6572223A2254455354222C226E616D65223A225465737420546F6B656E222C2261737365745F636C617373223A2264656669222C2261737365745F737562636C617373223A22737461626C65636F696E227D",
	"OwnerNode": "0000000000000000",
	"LockedAmount": "10000",
	"PreviousTxnID": "F0AB71E777B2DA54B86231E19B82554EF1F8211F92ECA473121C655BFC5329BF",
	"PreviousTxnLgrSeq": 14524914,
	"index": "F0AB71E777B2DA54B86231E19B82554EF1F8211F92ECA473121C655BFC5329BF"
}`

	if err := test.SerializeAndDeserialize(t, s, j); err != nil {
		t.Error(err)
	}
}

func TestMPTokenIssuance_EntryType(t *testing.T) {
	mpti := &MPTokenIssuance{}
	if mpti.EntryType() != MPTokenIssuanceEntry {
		t.Errorf("Expected EntryType to be %s, got %s", MPTokenIssuanceEntry, mpti.EntryType())
	}
}

func TestMPTokenMetadata(t *testing.T) {
	metadata := MPTokenMetadata{
		Ticker:        "TEST",
		Name:          "Test Token",
		Desc:          "A test token for testing purposes",
		Icon:          "https://example.com/icon.png",
		AssetClass:    "defi",
		AssetSubclass: "stablecoin",
		IssuerName:    "Test Issuer",
		Urls: []MPTokenMetadataUrl{
			{
				Url:   "https://example.com",
				Type:  "website",
				Title: "Official Website",
			},
		},
		AdditionalInfo: json.RawMessage(`{"custom_field": "custom_value"}`),
	}

	// Test GetBlob
	blob, err := metadata.GetBlob()
	if err != nil {
		t.Errorf("GetBlob failed: %v", err)
	}

	// Test NewMPTokenMetadataFromBlob
	recoveredMetadata, err := NewMPTokenMetadataFromBlob(blob)
	if err != nil {
		t.Errorf("NewMPTokenMetadataFromBlob failed: %v", err)
	}

	if recoveredMetadata.Ticker != metadata.Ticker {
		t.Errorf("Expected Ticker %s, got %s", metadata.Ticker, recoveredMetadata.Ticker)
	}

	if recoveredMetadata.Name != metadata.Name {
		t.Errorf("Expected Name %s, got %s", metadata.Name, recoveredMetadata.Name)
	}

	if recoveredMetadata.AssetClass != metadata.AssetClass {
		t.Errorf("Expected AssetClass %s, got %s", metadata.AssetClass, recoveredMetadata.AssetClass)
	}

	if recoveredMetadata.AssetSubclass != metadata.AssetSubclass {
		t.Errorf("Expected AssetSubclass %s, got %s", metadata.AssetSubclass, recoveredMetadata.AssetSubclass)
	}
}

func TestMPTokenMetadata_Validate(t *testing.T) {
	tests := []struct {
		name          string
		assetClass    string
		assetSubclass string
		shouldPass    bool
	}{
		{"valid defi stablecoin", "defi", "stablecoin", true},
		{"valid rwa real_estate", "rwa", "real_estate", true},
		{"valid gaming equity", "gaming", "equity", true},
		{"valid memes other", "memes", "other", true},
		{"valid wrapped treasury", "wrapped", "treasury", true},
		{"invalid asset class", "invalid", "stablecoin", false},
		{"invalid asset subclass", "defi", "invalid", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			metadata := MPTokenMetadata{
				AssetClass:    tt.assetClass,
				AssetSubclass: tt.assetSubclass,
			}

			err := metadata.Validate()
			if tt.shouldPass && err == nil {
				// Test should pass
				return
			}
			if !tt.shouldPass && err != nil {
				// Test should fail
				return
			}

			if tt.shouldPass {
				t.Errorf("Expected validation to pass, but got error: %v", err)
			} else {
				t.Errorf("Expected validation to fail, but got no error")
			}
		})
	}
}
