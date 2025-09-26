package transaction

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/Peersyst/xrpl-go/xrpl/testutil"
	rippletime "github.com/Peersyst/xrpl-go/xrpl/time"
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
	"github.com/stretchr/testify/assert"
)

func TestPaymentChannelFund_TxType(t *testing.T) {
	tx := &PaymentChannelFund{}
	assert.Equal(t, PaymentChannelFundTx, tx.TxType())
}

func TestPaymentChannelFund_Flatten(t *testing.T) {
	tests := []struct {
		name     string
		tx       *PaymentChannelFund
		expected string
	}{
		{
			name: "pass - without Expiration",
			tx: &PaymentChannelFund{
				BaseTx: BaseTx{
					Account:         "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
					TransactionType: PaymentChannelFundTx,
				},
				Channel: "ABC123",
				Amount:  types.XRPCurrencyAmount(200000),
			},
			expected: `{
				"Account": "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
				"TransactionType": "PaymentChannelFund",
				"Channel": "ABC123",
				"Amount":  "200000"
			}`,
		},
		{
			name: "pass - with Expiration",
			tx: &PaymentChannelFund{
				BaseTx: BaseTx{
					Account:         "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
					TransactionType: PaymentChannelFundTx,
				},
				Channel:    "DEF456",
				Amount:     types.XRPCurrencyAmount(300000),
				Expiration: 543171558,
			},
			expected: `{
				"Account": "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
				"TransactionType": "PaymentChannelFund",
				"Channel": "DEF456",
				"Amount": "300000",
				"Expiration": 543171558
			}`,
		},
		{
			name: "pass - with Expiration and Amount as MPTCurrencyAmount",
			tx: &PaymentChannelFund{
				BaseTx: BaseTx{
					Account:         "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
					TransactionType: PaymentChannelFundTx,
				},
				Channel: "DEF456",
				Amount: types.MPTCurrencyAmount{
					MPTIssuanceID: "1234567890ABCDEF1234567890ABCDEF1234567890ABCDEF",
					Value:         "300000",
				},
				Expiration: 543171558,
			},
			expected: `{
				"Account": "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
				"TransactionType": "PaymentChannelFund",
				"Channel": "DEF456",
				"Amount": {
					"mpt_issuance_id": "1234567890ABCDEF1234567890ABCDEF1234567890ABCDEF",
					"value": "300000"
				},
				"Expiration": 543171558
			}`,
		},
		{
			name: "pass - with Expiration and Amount as IssuedCurrencyAmount",
			tx: &PaymentChannelFund{
				BaseTx: BaseTx{
					Account:         "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
					TransactionType: PaymentChannelFundTx,
				},
				Channel: "DEF456",
				Amount: types.IssuedCurrencyAmount{
					Issuer:   "rEXAMPLE123456789ABCDEFGHJKLMNPQRSTUVWXYZ",
					Currency: "USD",
					Value:    "300000",
				},
				Expiration: 543171558,
			},
			expected: `{
				"Account": "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
				"TransactionType": "PaymentChannelFund",
				"Channel": "DEF456",
				"Amount": {
					"issuer": "rEXAMPLE123456789ABCDEFGHJKLMNPQRSTUVWXYZ",
					"currency": "USD",
					"value": "300000"
				},
				"Expiration": 543171558
			}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := testutil.CompareFlattenAndExpected(tt.tx.Flatten(), []byte(tt.expected))
			if err != nil {
				t.Error(err)
			}
		})
	}
}

func TestPaymentChannelFund_Validate(t *testing.T) {
	tests := []struct {
		name             string
		tx               *PaymentChannelFund
		expirationSetter func(tx *PaymentChannelFund)
		wantValid        bool
		wantErr          bool
		expectedErr      error
	}{
		{
			name: "pass - valid Transaction",
			tx: &PaymentChannelFund{
				BaseTx: BaseTx{
					Account:         "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
					TransactionType: PaymentChannelFundTx,
				},
				Channel: "ABC123",
				Amount:  types.XRPCurrencyAmount(200000),
			},
			expirationSetter: func(tx *PaymentChannelFund) {
				tx.Expiration = uint32(time.Now().Unix()) + 5000
			},
			wantValid:   true,
			wantErr:     false,
			expectedErr: nil,
		},
		{
			name: "fail - invalid BaseTx, missing Account",
			tx: &PaymentChannelFund{
				BaseTx: BaseTx{
					TransactionType: PaymentChannelFundTx,
				},
				Channel: "ABC123",
				Amount:  types.XRPCurrencyAmount(200000),
			},
			expirationSetter: func(tx *PaymentChannelFund) {
				tx.Expiration = uint32(rippletime.UnixTimeToRippleTime(time.Now().Unix()) + 5000)
			},
			wantValid:   false,
			wantErr:     true,
			expectedErr: ErrInvalidAccount,
		},
		{
			name: "fail - invalid Expiration",
			tx: &PaymentChannelFund{
				BaseTx: BaseTx{
					Account:         "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
					TransactionType: PaymentChannelFundTx,
				},
				Channel:    "DEF456",
				Amount:     types.XRPCurrencyAmount(300000),
				Expiration: 1, // Invalid expiration time
			},
			wantValid:   false,
			wantErr:     true,
			expectedErr: ErrInvalidExpiration,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			valid, err := tt.tx.Validate()
			if tt.expirationSetter != nil {
				tt.expirationSetter(tt.tx)
			}

			assert.Equal(t, tt.wantValid, valid)
			if (err != nil) && err != tt.expectedErr {
				t.Errorf("Validate() got error message = %v, want error message %v", err, tt.expectedErr)
				return
			}
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func TestPaymentChannelFund_Unmarshal(t *testing.T) {
	tests := []struct {
		name                 string
		jsonData             string
		expectUnmarshalError bool
	}{
		{
			name: "pass - full PaymentChannelFund with MPTCurrencyAmount",
			jsonData: `{
				"TransactionType": "PaymentChannelFund",
				"Account": "rEXAMPLE123456789ABCDEFGHJKLMNPQRSTUVWXYZ",
				"Channel": "DEF456",
				"Amount": {
					"mpt_issuance_id": "1234567890ABCDEF1234567890ABCDEF1234567890ABCDEF",
					"value": "300000"
				},
				"Expiration": 543171558
			}`,
		},
		{
			name: "pass - full PaymentChannelFund with IssuedCurrencyAmount",
			jsonData: `{
				"TransactionType": "PaymentChannelFund",
				"Account": "rEXAMPLE123456789ABCDEFGHJKLMNPQRSTUVWXYZ",
				"Channel": "DEF456",
				"Amount": {
					"issuer": "rEXAMPLE123456789ABCDEFGHJKLMNPQRSTUVWXYZ",
					"currency": "USD",
					"value": "300000"
				},
				"Expiration": 543171558
			}`,
		},
		{
			name: "pass - full PaymentChannelFund with XRPCurrencyAmount",
			jsonData: `{
				"TransactionType": "PaymentChannelFund",
				"Account": "rEXAMPLE123456789ABCDEFGHJKLMNPQRSTUVWXYZ",
				"Channel": "DEF456",
				"Amount": "300000",
				"Expiration": 543171558
			}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var paymentChannelFund PaymentChannelFund
			err := json.Unmarshal([]byte(tt.jsonData), &paymentChannelFund)
			assert.Equal(t, tt.expectUnmarshalError, err != nil)
		})
	}
}
