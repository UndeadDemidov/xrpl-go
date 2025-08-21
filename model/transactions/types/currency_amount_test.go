package types

import (
	"encoding/hex"
	"encoding/json"
	"testing"
)

func TestUnmarshalCurrencyAmount(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    CurrencyAmount
		wantErr bool
	}{
		{
			name:    "XRPCurrencyAmount as string",
			input:   `"1000000"`,
			want:    XRPCurrencyAmount(1000000),
			wantErr: false,
		},
		{
			name:    "XRPCurrencyAmount zero",
			input:   `"0"`,
			want:    XRPCurrencyAmount(0),
			wantErr: false,
		},
		{
			name:    "IssuedCurrencyAmount with all fields",
			input:   `{"issuer":"rajgkBmMxmz161r8bWYH7CQAFZP5bA9oSG","currency":"USD","value":"100.50"}`,
			want:    IssuedCurrencyAmount{Issuer: "rajgkBmMxmz161r8bWYH7CQAFZP5bA9oSG", Currency: "USD", Value: "100.50"},
			wantErr: false,
		},
		{
			name:    "IssuedCurrencyAmount without issuer",
			input:   `{"currency":"EUR","value":"250.00"}`,
			want:    IssuedCurrencyAmount{Currency: "EUR", Value: "250.00"},
			wantErr: false,
		},
		{
			name:    "IssuedCurrencyAmount with empty issuer",
			input:   `{"issuer":"","currency":"GBP","value":"75.25"}`,
			want:    IssuedCurrencyAmount{Issuer: "", Currency: "GBP", Value: "75.25"},
			wantErr: false,
		},
		{
			name:    "MPTCurrencyAmount with all fields",
			input:   `{"mp_issuance_id":"000004C463C52827307480341125DA0577DEFC38405B0E3E","value":"500.00"}`,
			want:    MPTCurrencyAmount{MPTIssuanceID: "000004C463C52827307480341125DA0577DEFC38405B0E3E", Value: "500.00"},
			wantErr: false,
		},
		{
			name:    "Empty input",
			input:   ``,
			want:    nil,
			wantErr: false,
		},
		{
			name:    "Invalid JSON object",
			input:   `{"invalid":}`,
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Invalid XRP amount string",
			input:   `"not_a_number"`,
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Invalid XRP amount number",
			input:   `-1000`,
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := UnmarshalCurrencyAmount([]byte(tt.input))

			if (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalCurrencyAmount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.wantErr {
				return
			}

			if tt.input == "" && got == nil {
				// Empty input should return nil
				return
			}

			if got == nil {
				t.Errorf("UnmarshalCurrencyAmount() returned nil, expected %v", tt.want)
				return
			}

			// Check the kind
			if got.Kind() != tt.want.Kind() {
				t.Errorf("UnmarshalCurrencyAmount() kind = %v, want %v", got.Kind(), tt.want.Kind())
			}

			// Check specific type and values
			switch want := tt.want.(type) {
			case XRPCurrencyAmount:
				if gotXRPCurrencyAmount, ok := got.(XRPCurrencyAmount); !ok {
					t.Errorf("UnmarshalCurrencyAmount() returned wrong type, expected XRPCurrencyAmount")
				} else if gotXRPCurrencyAmount != want {
					t.Errorf("UnmarshalCurrencyAmount() = %v, want %v", gotXRPCurrencyAmount, want)
				}
			case IssuedCurrencyAmount:
				if gotIssuedCurrencyAmount, ok := got.(IssuedCurrencyAmount); !ok {
					t.Errorf("UnmarshalCurrencyAmount() returned wrong type, expected IssuedCurrencyAmount")
				} else if gotIssuedCurrencyAmount.Issuer != want.Issuer ||
					gotIssuedCurrencyAmount.Currency != want.Currency ||
					gotIssuedCurrencyAmount.Value != want.Value {
					t.Errorf("UnmarshalCurrencyAmount() = %v, want %v", gotIssuedCurrencyAmount, want)
				}
			case MPTCurrencyAmount:
				if gotMPTCurrencyAmount, ok := got.(MPTCurrencyAmount); !ok {
					t.Errorf("UnmarshalCurrencyAmount() returned wrong type, expected MPTCurrencyAmount")
				} else if gotMPTCurrencyAmount.MPTIssuanceID != want.MPTIssuanceID ||
					gotMPTCurrencyAmount.Value != want.Value {
					t.Errorf("UnmarshalCurrencyAmount() = %v, want %v", gotMPTCurrencyAmount, want)
				}
			}
		})
	}
}

func TestUnmarshalCurrencyAmount_JSONRoundTrip(t *testing.T) {
	// Test that we can marshal and unmarshal each type correctly
	testCases := []struct {
		name string
		data CurrencyAmount
	}{
		{
			name: "XRPCurrencyAmount",
			data: XRPCurrencyAmount(1000000),
		},
		{
			name: "IssuedCurrencyAmount",
			data: IssuedCurrencyAmount{
				Issuer:   "rHb9CJAWyB4rj91VRWn96DkukG4bwdtyTh",
				Currency: "USD",
				Value:    "100.50",
			},
		},
		{
			name: "MPTCurrencyAmount",
			data: MPTCurrencyAmount{
				MPTIssuanceID: "000004C463C52827307480341125DA0577DEFC38405B0E3E",
				Value:         "500.00",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Marshal to JSON
			jsonData, err := json.Marshal(tc.data)
			if err != nil {
				t.Fatalf("Failed to marshal %s: %v", tc.name, err)
			}

			// Unmarshal back
			unmarshaled, err := UnmarshalCurrencyAmount(jsonData)
			if err != nil {
				t.Fatalf("Failed to unmarshal %s: %v", tc.name, err)
			}

			// Check that we got the same type back
			if unmarshaled.Kind() != tc.data.Kind() {
				t.Errorf("Kind mismatch: got %v, want %v", unmarshaled.Kind(), tc.data.Kind())
			}

			// Check that the data is equivalent
			if !jsonEqual(tc.data, unmarshaled) {
				t.Errorf("Data mismatch: got %v, want %v", unmarshaled, tc.data)
			}
		})
	}
}

func TestMPTCurrencyAmount_Methods(t *testing.T) {
	// Test with the specific values provided by the user
	mpt := MPTCurrencyAmount{
		MPTIssuanceID: "005342F3D0045E31BC8E02B4A85A5FE56B33AA25C3745405",
		Value:         "100.00",
	}

	t.Run("Validate", func(t *testing.T) {
		err := mpt.Validate()
		if err != nil {
			t.Errorf("Validate() failed: %v", err)
		}
	})

	t.Run("Kind", func(t *testing.T) {
		kind := mpt.Kind()
		if kind != MPT {
			t.Errorf("Kind() = %v, want %v", kind, MPT)
		}
	})

	t.Run("SequenceNumber", func(t *testing.T) {
		sequence, err := mpt.SequenceNumber()
		if err != nil {
			t.Errorf("SequenceNumber() failed: %v", err)
		}

		expectedSequence := uint32(5456627)
		if sequence != expectedSequence {
			t.Errorf("SequenceNumber() = %v, want %v", sequence, expectedSequence)
		}
	})

	t.Run("IssuerAccountID", func(t *testing.T) {
		issuerID, err := mpt.IssuerAccountID()
		if err != nil {
			t.Errorf("IssuerAccountID() failed: %v", err)
		}

		// The issuer account ID should be the last 20 bytes (40 hex chars) of the MPTIssuanceID
		// Removing the first 4 bytes (8 hex chars) which represent the sequence number
		expectedIssuerHex := "D0045E31BC8E02B4A85A5FE56B33AA25C3745405"
		expectedIssuerBytes, _ := hex.DecodeString(expectedIssuerHex)

		if len(issuerID) != len(expectedIssuerBytes) {
			t.Errorf("IssuerAccountID() length = %v, want %v", len(issuerID), len(expectedIssuerBytes))
		}

		for i, b := range issuerID {
			if b != expectedIssuerBytes[i] {
				t.Errorf("IssuerAccountID()[%d] = %02x, want %02x", i, b, expectedIssuerBytes[i])
			}
		}
	})

	t.Run("JSON marshaling", func(t *testing.T) {
		jsonData, err := json.Marshal(mpt)
		if err != nil {
			t.Errorf("JSON marshaling failed: %v", err)
		}

		expectedJSON := `{"mp_issuance_id":"005342F3D0045E31BC8E02B4A85A5FE56B33AA25C3745405","value":"100.00"}`
		if string(jsonData) != expectedJSON {
			t.Errorf("JSON marshaling = %s, want %s", string(jsonData), expectedJSON)
		}
	})
}

// jsonEqual compares two CurrencyAmount values by marshaling them to JSON and comparing the strings
func jsonEqual(a, b CurrencyAmount) bool {
	jsonA, errA := json.Marshal(a)
	jsonB, errB := json.Marshal(b)

	if errA != nil || errB != nil {
		return false
	}

	return string(jsonA) == string(jsonB)
}
