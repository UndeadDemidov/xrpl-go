package transaction

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/Peersyst/xrpl-go/xrpl/testutil"
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
	"github.com/stretchr/testify/assert"
)

func TestPaymentChannelClaim_TxType(t *testing.T) {
	tx := &PaymentChannelClaim{}
	assert.Equal(t, PaymentChannelClaimTx, tx.TxType())
}

func TestPaymentChannelClaimFlags(t *testing.T) {
	tests := []struct {
		name     string
		setter   func(*PaymentChannelClaim)
		expected uint32
	}{
		{
			name: "pass - SetRenewFlag",
			setter: func(p *PaymentChannelClaim) {
				p.SetRenewFlag()
			},
			expected: tfRenew,
		},
		{
			name: "pass - SetCloseFlag",
			setter: func(p *PaymentChannelClaim) {
				p.SetCloseFlag()
			},
			expected: tfClose,
		},
		{
			name: "pass - SetRenewFlag and SetCloseFlag",
			setter: func(p *PaymentChannelClaim) {
				p.SetRenewFlag()
				p.SetCloseFlag()
			},
			expected: tfRenew | tfClose,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PaymentChannelClaim{}
			tt.setter(p)
			assert.Equal(t, tt.expected, p.Flags)
		})
	}
}

func TestPaymentChannelClaim_Flatten(t *testing.T) {
	tests := []struct {
		name     string
		claim    PaymentChannelClaim
		expected string
	}{
		{
			name: "pass - PaymentChannelClaim with Channel",
			claim: PaymentChannelClaim{
				BaseTx: BaseTx{
					Account:         "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
					TransactionType: PaymentChannelClaimTx,
				},
				Channel: types.Hash256("ABC123"),
			},
			expected: `{
				"Account": "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
				"TransactionType": "PaymentChannelClaim",
				"Channel": "ABC123"
			}`,
		},
		{
			name: "pass - PaymentChannelClaim with Balance and Amount as XRPCurrencyAmount",
			claim: PaymentChannelClaim{
				BaseTx: BaseTx{
					Account:         "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
					TransactionType: PaymentChannelClaimTx,
				},
				Balance: types.XRPCurrencyAmount(1000),
				Amount:  types.XRPCurrencyAmount(2000),
			},
			expected: `{
				"Account": "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
				"TransactionType": "PaymentChannelClaim",
				"Balance": "1000",
				"Amount": "2000"
			}`,
		},
		{
			name: "pass - PaymentChannelClaim with Balance and Amount as MPTCurrencyAmount",
			claim: PaymentChannelClaim{
				BaseTx: BaseTx{
					Account:         "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
					TransactionType: PaymentChannelClaimTx,
				},
				Balance: types.MPTCurrencyAmount{
					MPTIssuanceID: "1234567890ABCDEF1234567890ABCDEF1234567890ABCDEF",
					Value:         "1000",
				},
				Amount: types.MPTCurrencyAmount{
					MPTIssuanceID: "1234567890ABCDEF1234567890ABCDEF1234567890ABCDEF",
					Value:         "2000",
				},
			},
			expected: `{
				"Account": "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
				"TransactionType": "PaymentChannelClaim",
				"Balance": {
					"mpt_issuance_id": "1234567890ABCDEF1234567890ABCDEF1234567890ABCDEF",
					"value": "1000"
				},
				"Amount": {
					"mpt_issuance_id": "1234567890ABCDEF1234567890ABCDEF1234567890ABCDEF",
					"value": "2000"
				}
			}`,
		},
		{
			name: "pass - PaymentChannelClaim with Balance and Amount as IssuedCurrencyAmount",
			claim: PaymentChannelClaim{
				BaseTx: BaseTx{
					Account:         "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
					TransactionType: PaymentChannelClaimTx,
				},
				Balance: types.IssuedCurrencyAmount{
					Issuer:   "rEXAMPLE123456789ABCDEFGHJKLMNPQRSTUVWXYZ",
					Currency: "USD",
					Value:    "1000",
				},
				Amount: types.IssuedCurrencyAmount{
					Issuer:   "rEXAMPLE123456789ABCDEFGHJKLMNPQRSTUVWXYZ",
					Currency: "USD",
					Value:    "2000",
				},
			},
			expected: `{
				"Account": "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
				"TransactionType": "PaymentChannelClaim",
				"Balance": {
					"issuer": "rEXAMPLE123456789ABCDEFGHJKLMNPQRSTUVWXYZ",
					"currency": "USD",
					"value": "1000"
				},
				"Amount": {
					"issuer": "rEXAMPLE123456789ABCDEFGHJKLMNPQRSTUVWXYZ",
					"currency": "USD",
					"value": "2000"
				}
			}`,
		},
		{
			name: "pass - PaymentChannelClaim with Signature and PublicKey",
			claim: PaymentChannelClaim{
				BaseTx: BaseTx{
					Account:         "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
					TransactionType: PaymentChannelClaimTx,
				},
				Signature: "ABCDEF",
				PublicKey: "123456",
			},
			expected: `{
				"Account": "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
				"TransactionType": "PaymentChannelClaim",
				"Signature": "ABCDEF",
				"PublicKey": "123456"
			}`,
		},
		{
			name: "pass - PaymentChannelClaim with all fields",
			claim: PaymentChannelClaim{
				BaseTx: BaseTx{
					Account:         "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
					TransactionType: PaymentChannelClaimTx,
				},
				Channel:       types.Hash256("ABC123"),
				Balance:       types.XRPCurrencyAmount(1000),
				Amount:        types.XRPCurrencyAmount(2000),
				Signature:     "ABCDEF",
				PublicKey:     "123456",
				CredentialIDs: types.CredentialIDs{"1234567890abcdef"},
			},
			expected: `{
				"Account": "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
				"TransactionType": "PaymentChannelClaim",
				"Channel": "ABC123",
				"Balance": "1000",
				"Amount": "2000",
				"Signature": "ABCDEF",
				"PublicKey": "123456",
				"CredentialIDs": ["1234567890abcdef"]
			}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := testutil.CompareFlattenAndExpected(tt.claim.Flatten(), []byte(tt.expected))
			assert.NoError(t, err)
		})
	}
}

func TestPaymentChannelClaim_Validate(t *testing.T) {
	tests := []struct {
		name        string
		claim       PaymentChannelClaim
		wantValid   bool
		wantErr     bool
		expectedErr error
	}{
		{
			name: "pass - all fields valid",
			claim: PaymentChannelClaim{
				BaseTx: BaseTx{
					Account:         "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
					TransactionType: PaymentChannelClaimTx,
				},
				Balance:       types.XRPCurrencyAmount(1000),
				Channel:       "ABC123ABC123ABC123ABC123ABC123ABC123ABC123ABC123ABC123ABC123ABC1",
				Signature:     "ABCDEF",
				PublicKey:     "ABC123ABC123ABC123ABC123ABC123ABC123ABC123ABC123ABC123ABC123ABC1",
				CredentialIDs: types.CredentialIDs{"1234567890abcdef"},
			},
			wantValid: true,
			wantErr:   false,
		},
		{
			name: "fail - missing Account in BaseTx",
			claim: PaymentChannelClaim{
				BaseTx: BaseTx{
					TransactionType: PaymentChannelClaimTx,
				},
				Balance:   types.XRPCurrencyAmount(1000),
				Channel:   "ABC123ABC123ABC123ABC123ABC123ABC123ABC123ABC123ABC123ABC123ABC1",
				Signature: "ABCDEF",
				PublicKey: "ABC123ABC123ABC123ABC123ABC123ABC123ABC123ABC123ABC123ABC123ABC1",
			},
			wantValid:   false,
			wantErr:     true,
			expectedErr: ErrInvalidAccount,
		},
		{
			name: "fail - empty Channel",
			claim: PaymentChannelClaim{
				BaseTx: BaseTx{
					Account:         "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
					TransactionType: PaymentChannelClaimTx,
				},
			},
			wantValid:   false,
			wantErr:     true,
			expectedErr: ErrInvalidChannel,
		},
		{
			name: "fail - invalid Signature",
			claim: PaymentChannelClaim{
				BaseTx: BaseTx{
					Account:         "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
					TransactionType: PaymentChannelClaimTx,
				},
				Channel:   "ABC123ABC123ABC123ABC123ABC123ABC123ABC123ABC123ABC123ABC123ABC1",
				Signature: "INVALID_SIGNATURE",
			},
			wantValid:   false,
			wantErr:     true,
			expectedErr: ErrInvalidSignature,
		},
		{
			name: "pass - no Signature",
			claim: PaymentChannelClaim{
				BaseTx: BaseTx{
					Account:         "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
					TransactionType: PaymentChannelClaimTx,
				},
				Channel: "ABC123ABC123ABC123ABC123ABC123ABC123ABC123ABC123ABC123ABC123ABC1",
			},
			wantValid: true,
			wantErr:   false,
		},
		{
			name: "fail - invalid PublicKey",
			claim: PaymentChannelClaim{
				BaseTx: BaseTx{
					Account:         "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
					TransactionType: PaymentChannelClaimTx,
				},
				Channel:   "ABC123ABC123ABC123ABC123ABC123ABC123ABC123ABC123ABC123ABC123ABC1",
				PublicKey: "INVALID",
			},
			wantValid:   false,
			wantErr:     true,
			expectedErr: ErrInvalidHexPublicKey,
		},
		{
			name: "fail - invalid CredentialIDs",
			claim: PaymentChannelClaim{
				BaseTx: BaseTx{
					Account:         "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
					TransactionType: PaymentChannelClaimTx,
				},
				Channel:       "ABC123ABC123ABC123ABC123ABC123ABC123ABC123ABC123ABC123ABC123ABC1",
				CredentialIDs: types.CredentialIDs{"invalid"},
			},
			wantValid:   false,
			wantErr:     true,
			expectedErr: ErrInvalidCredentialIDs,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			valid, err := tt.claim.Validate()
			assert.Equal(t, tt.wantValid, valid)
			assert.Equal(t, tt.wantErr, err != nil)
			if err != nil && err != tt.expectedErr {
				t.Errorf("Validate() error = %v, expectedErr %v", err, tt.expectedErr)
			}
		})
	}
}

func TestPaymentChannelClaim_Unmarshal(t *testing.T) {
	tests := []struct {
		name                 string
		jsonData             string
		expectUnmarshalError bool
	}{
		{
			name: "pass - full PaymentChannelClaim with IssuedCurrencyAmount",
			jsonData: `{
				"TransactionType": "PaymentChannelClaim",
				"Account": "rEXAMPLE123456789ABCDEFGHJKLMNPQRSTUVWXYZ",
				"Channel": "ABC123ABC123ABC123ABC123ABC123ABC123ABC123ABC123ABC123ABC123ABC1",
				"Amount": {
					"issuer": "rEXAMPLE123456789ABCDEFGHJKLMNPQRSTUVWXYZ",
					"currency": "USD",
					"value": "1000000"
				},
				"Balance": {
					"issuer": "rEXAMPLE123456789ABCDEFGHJKLMNPQRSTUVWXYZ",
					"currency": "USD",
					"value": "900000"
				},
				"CredentialIDs": ["1234567890abcdef", "1234567890abcdefg"],
				"Fee": "10",
				"Sequence": 1,
				"Flags": 2147483648,
				"CancelAfter": 695123456,
				"FinishAfter": 695000000,
				"Condition": "A0258020C4F71E9B01F5A78023E932ABF6B2C1F020986E6C9E55678FFBAE67A2F5B474680103080000000000000000000000000000000000000000000000000000000000000000",
				"DestinationTag": 12345,
				"SourceTag": 54321,
				"OwnerNode": "0000000000000000",
				"PreviousTxnID": "C4F71E9B01F5A78023E932ABF6B2C1F020986E6C9E55678FFBAE67A2F5B47468",
				"LastLedgerSequence": 12345678,
				"NetworkID": 1024,
				"Memos": [
					{
					"Memo": {
						"MemoType": "657363726F77",
						"MemoData": "457363726F77206372656174656420666F72207061796D656E74"
					}
					}
				],
				"Signers": [
					{
					"Signer": {
						"Account": "rSIGNER123456789ABCDEFGHJKLMNPQRSTUVWXYZ",
						"SigningPubKey": "ED5F93AB1234567890ABCDEF1234567890ABCDEF1234567890ABCDEF12345678",
						"TxnSignature": "3045022100D7F67A81F343...B87D"
					}
					}
				],
				"SigningPubKey": "ED5F93AB1234567890ABCDEF1234567890ABCDEF1234567890ABCDEF12345678",
				"TxnSignature": "3045022100D7F67A81F343...B87D"
			}`,
			expectUnmarshalError: false,
		},
		{
			name: "pass - full PaymentChannelClaim with MPTCurrencyAmount",
			jsonData: `{
				"TransactionType": "PaymentChannelClaim",
				"Account": "rEXAMPLE123456789ABCDEFGHJKLMNPQRSTUVWXYZ",
				"Channel": "ABC123ABC123ABC123ABC123ABC123ABC123ABC123ABC123ABC123ABC123ABC1",
				"Amount": {
					"mpt_issuance_id": "1234567890ABCDEF1234567890ABCDEF1234567890ABCDEF",
					"value": "1000000"
				},
				"Balance": {
					"mpt_issuance_id": "1234567890ABCDEF1234567890ABCDEF1234567890ABCDEF",
					"value": "900000"
				},
				"CredentialIDs": ["1234567890abcdef", "1234567890abcdefg"],
				"Fee": "10",
				"Sequence": 1,
				"Flags": 2147483648,
				"CancelAfter": 695123456,
				"FinishAfter": 695000000,
				"Condition": "A0258020C4F71E9B01F5A78023E932ABF6B2C1F020986E6C9E55678FFBAE67A2F5B474680103080000000000000000000000000000000000000000000000000000000000000000",
				"DestinationTag": 12345,
				"SourceTag": 54321,
				"OwnerNode": "0000000000000000",
				"PreviousTxnID": "C4F71E9B01F5A78023E932ABF6B2C1F020986E6C9E55678FFBAE67A2F5B47468",
				"LastLedgerSequence": 12345678,
				"NetworkID": 1024,
				"Memos": [
					{
					"Memo": {
						"MemoType": "657363726F77",
						"MemoData": "457363726F77206372656174656420666F72207061796D656E74"
					}
					}
				],
				"Signers": [
					{
					"Signer": {
						"Account": "rSIGNER123456789ABCDEFGHJKLMNPQRSTUVWXYZ",
						"SigningPubKey": "ED5F93AB1234567890ABCDEF1234567890ABCDEF1234567890ABCDEF12345678",
						"TxnSignature": "3045022100D7F67A81F343...B87D"
					}
					}
				],
				"SigningPubKey": "ED5F93AB1234567890ABCDEF1234567890ABCDEF1234567890ABCDEF12345678",
				"TxnSignature": "3045022100D7F67A81F343...B87D"
			}`,
			expectUnmarshalError: false,
		},
		{
			name: "pass - full PaymentChannelClaim with XRPCurrencyAmount",
			jsonData: `{
				"TransactionType": "PaymentChannelClaim",
				"Account": "rEXAMPLE123456789ABCDEFGHJKLMNPQRSTUVWXYZ",
				"Channel": "ABC123ABC123ABC123ABC123ABC123ABC123ABC123ABC123ABC123ABC123ABC1",
				"Amount": "1000000",
				"Balance": "30000",
				"CredentialIDs": ["1234567890abcdef", "1234567890abcdefg"],
				"Fee": "10",
				"Sequence": 1,
				"Flags": 2147483648,
				"CancelAfter": 695123456,
				"FinishAfter": 695000000,
				"Condition": "A0258020C4F71E9B01F5A78023E932ABF6B2C1F020986E6C9E55678FFBAE67A2F5B474680103080000000000000000000000000000000000000000000000000000000000000000",
				"DestinationTag": 12345,
				"SourceTag": 54321,
				"OwnerNode": "0000000000000000",
				"PreviousTxnID": "C4F71E9B01F5A78023E932ABF6B2C1F020986E6C9E55678FFBAE67A2F5B47468",
				"LastLedgerSequence": 12345678,
				"NetworkID": 1024,
				"Memos": [
					{
					"Memo": {
						"MemoType": "657363726F77",
						"MemoData": "457363726F77206372656174656420666F72207061796D656E74"
					}
					}
				],
				"Signers": [
					{
					"Signer": {
						"Account": "rSIGNER123456789ABCDEFGHJKLMNPQRSTUVWXYZ",
						"SigningPubKey": "ED5F93AB1234567890ABCDEF1234567890ABCDEF1234567890ABCDEF12345678",
						"TxnSignature": "3045022100D7F67A81F343...B87D"
					}
					}
				],
				"SigningPubKey": "ED5F93AB1234567890ABCDEF1234567890ABCDEF1234567890ABCDEF12345678",
				"TxnSignature": "3045022100D7F67A81F343...B87D"
			}`,
			expectUnmarshalError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var paymentChannelCreate PaymentChannelClaim
			err := json.Unmarshal([]byte(tt.jsonData), &paymentChannelCreate)
			fmt.Println(paymentChannelCreate.TransactionType)
			if (err != nil) != tt.expectUnmarshalError {
				t.Errorf("Unmarshal() error = %v, expectUnmarshalError %v", err, tt.expectUnmarshalError)
				return
			}

		})
	}
}
