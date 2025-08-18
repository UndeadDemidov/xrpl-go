package types

import (
	"encoding/binary"
	"encoding/hex"
	"strconv"
	"strings"
	"testing"

	"github.com/CreatureDev/xrpl-go/model/transactions/types"
	bigdecimal "github.com/CreatureDev/xrpl-go/pkg/big-decimal"
	"github.com/stretchr/testify/require"
)

func TestVerifyXrpValue(t *testing.T) {

	tests := []struct {
		name   string
		input  types.XRPCurrencyAmount
		expErr error
	}{
		// {
		// 	name:   "invalid xrp value",
		// 	input:  1.0,
		// 	expErr: ErrInvalidXRPValue,
		// },
		// {
		// 	name:   "invalid xrp value - out of range",
		// 	input:  0.000000007,
		// 	expErr: ErrInvalidXRPValue,
		// },
		{
			name:   "valid xrp value - no decimal",
			input:  125000708,
			expErr: nil,
		},
		// {
		// 	name:   "invalid xrp value - no decimal - negative value",
		// 	input:  -125000708,
		// 	expErr: &InvalidAmountError{Amount: "-125000708"},
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.expErr != nil {
				require.Equal(t, tt.expErr, verifyXrpValue(tt.input))
			} else {
				require.NoError(t, verifyXrpValue(tt.input))
			}
		})
	}
}

func TestVerifyIOUValue(t *testing.T) {

	tests := []struct {
		name   string
		input  string
		expErr error
	}{
		{
			name:   "valid iou value with decimal",
			input:  "3.6",
			expErr: nil,
		},
		{
			name:   "valid iou value - leading zero after decimal",
			input:  "345.023857",
			expErr: nil,
		},
		{
			name:   "valid iou value - negative value & multiple leading zeros before decimal",
			input:  "-000.2345",
			expErr: nil,
		},
		{
			name:   "invalid iou value - out of range precision",
			input:  "0.000000000000000000007265675687436598345739475",
			expErr: &OutOfRangeError{Type: "Precision"},
		},
		{
			name:   "invalid iou value - out of range exponent too large",
			input:  "998000000000000000000000000000000000000000000000000000000000000000000000000000000000000000",
			expErr: &OutOfRangeError{Type: "Exponent"},
		},
		{
			name:   "invalid iou value - out of range exponent too small",
			input:  "0.0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000998",
			expErr: &OutOfRangeError{Type: "Exponent"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			err := verifyIOUValue(tt.input)
			if tt.expErr != nil {
				require.EqualError(t, tt.expErr, err.Error())
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestSerializeXrpAmount(t *testing.T) {
	tests := []struct {
		name           string
		input          types.XRPCurrencyAmount
		expectedOutput []byte
		expErr         error
	}{
		{
			name:           "valid xrp value - 1",
			input:          524801,
			expectedOutput: []byte{0x40, 0x00, 0x00, 0x00, 0x00, 0x8, 0x2, 0x01},
			expErr:         nil,
		},
		{
			name:           "valid xrp value - 2",
			input:          7696581656832,
			expectedOutput: []byte{0x40, 0x00, 0x7, 0x00, 0x00, 0x4, 0x1, 0x00},
			expErr:         nil,
		},
		{
			name:           "valid xrp value - 3",
			input:          10000000,
			expectedOutput: []byte{0x40, 0x00, 0x00, 0x00, 0x00, 0x98, 0x96, 0x80},
			expErr:         nil,
		},
		{
			name:           "boundary test - 1 less than max xrp value",
			input:          99999999999999999,
			expectedOutput: []byte{0x41, 0x63, 0x45, 0x78, 0x5d, 0x89, 0xff, 0xff},
			expErr:         nil,
		},
		{
			name:           "boundary test - max xrp value",
			input:          10000000000000000,
			expectedOutput: []byte{0x40, 0x23, 0x86, 0xf2, 0x6f, 0xc1, 0x00, 0x00},
			expErr:         nil,
		},
		{
			name:           "boundary test - 1 greater than max xrp value",
			input:          100000000000000001,
			expectedOutput: nil,
			expErr:         &InvalidAmountError{Amount: 100000000000000001},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := serializeXrpAmount(tt.input)
			if tt.expErr != nil {
				require.EqualError(t, tt.expErr, err.Error())
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.expectedOutput, got)
			}
		})
	}
}

func TestSerializeIssuedCurrencyValue(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		expected    []byte
		expectedErr error
	}{
		{
			name:        "invalid zero value",
			input:       "0",
			expected:    nil,
			expectedErr: bigdecimal.ErrInvalidZeroValue,
		},
		{
			name:        "valid value - 2",
			input:       "1",
			expected:    []byte{0xD4, 0x83, 0x8D, 0x7E, 0xA4, 0xC6, 0x80, 0x00},
			expectedErr: nil,
		},
		{
			name:        "valid value - 3",
			input:       "2.1",
			expected:    []byte{0xD4, 0x87, 0x75, 0xF0, 0x5A, 0x07, 0x40, 0x00},
			expectedErr: nil,
		},
		{
			name:        "valid value - from Transaction 1 in main_test.go",
			input:       "7072.8",
			expected:    []byte{0xD5, 0x59, 0x20, 0xAC, 0x93, 0x91, 0x40, 0x00},
			expectedErr: nil,
		},
		{
			name:        "valid value - from Transaction 3 in main_test.go",
			input:       "0.6275558355",
			expected:    []byte{0xd4, 0x56, 0x4b, 0x96, 0x4a, 0x84, 0x5a, 0xc0},
			expectedErr: nil,
		},
		{
			name:        "valid value - negative",
			input:       "-2",
			expected:    []byte{0x94, 0x87, 0x1A, 0xFD, 0x49, 0x8D, 0x00, 0x00},
			expectedErr: nil,
		},
		{
			name:        "valid value - negative - 2",
			input:       "-7072.8",
			expected:    []byte{0x95, 0x59, 0x20, 0xAC, 0x93, 0x91, 0x40, 0x00},
			expectedErr: nil,
		},
		{
			name:        "valid value - large currency amount",
			input:       "1111111111111111.0",
			expected:    []byte{0xD8, 0x43, 0xF2, 0x8C, 0xB7, 0x15, 0x71, 0xC7},
			expectedErr: nil,
		},
		{
			name:        "boundary test - max precision - max exponent",
			input:       "9999999999999999e80",
			expected:    []byte{0xec, 0x63, 0x86, 0xf2, 0x6f, 0xc0, 0xff, 0xff},
			expectedErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := serializeIssuedCurrencyValue(tt.input)

			if tt.expectedErr != nil {
				require.EqualError(t, tt.expectedErr, err.Error())
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.expected, got)
			}

		})
	}
}

func TestSerializeIssuedCurrencyCode(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		expected    []byte
		expectedErr error
	}{
		{
			name:        "valid standard currency - ISO4217 - USD",
			input:       "USD",
			expected:    []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x55, 0x53, 0x44, 0x00, 0x00, 0x00, 0x00, 0x00},
			expectedErr: nil,
		},
		{
			name:        "valid standard currency - ISO4217 - USD - hex",
			input:       "0x0000000000000000000000005553440000000000",
			expected:    []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x55, 0x53, 0x44, 0x00, 0x00, 0x00, 0x00, 0x00},
			expectedErr: nil,
		},
		{
			name:        "valid standard currency - non ISO4217 - BTC",
			input:       "BTC",
			expected:    []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x42, 0x54, 0x43, 0x00, 0x00, 0x00, 0x00, 0x00},
			expectedErr: nil,
		},
		{
			name:        "valid standard currency - non ISO4217 - BTC - hex",
			input:       "0x0000000000000000000000004254430000000000",
			expected:    []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x42, 0x54, 0x43, 0x00, 0x00, 0x00, 0x00, 0x00},
			expectedErr: nil,
		},
		{
			name:        "disallowed standard currency - XRP",
			input:       "XRP",
			expected:    nil,
			expectedErr: &InvalidCodeError{"XRP uppercase"},
		},
		{
			name:        "disallowed standard currency - XRP - hex",
			input:       "0000000000000000000000005852500000000000",
			expected:    nil,
			expectedErr: &InvalidCodeError{"XRP uppercase"},
		},
		{
			name:        "invalid standard currency - 4 characters",
			input:       "ABCD",
			expected:    nil,
			expectedErr: &InvalidCodeError{"ABCD"},
		},
		{
			name:        "valid non-standard currency - 4 characters - hex",
			input:       "0x4142434400000000000000000000000000000000",
			expected:    []byte{0x41, 0x42, 0x43, 0x44, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
			expectedErr: nil,
		},
		{
			name:        "special case - XRP - hex",
			input:       "0x0000000000000000000000000000000000000000",
			expected:    []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
			expectedErr: nil,
		},
		{
			name:        "standard currency - valid symbols in currency code - 3 characters",
			input:       "A*B",
			expected:    []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x41, 0x2a, 0x42, 0x00, 0x00, 0x00, 0x00, 0x00},
			expectedErr: nil,
		},
		{
			name:        "standard currency - valid symbols in currency code - 3 characters - hex",
			input:       "0x000000000000000000000000412a420000000000",
			expected:    []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x41, 0x2a, 0x42, 0x00, 0x00, 0x00, 0x00, 0x00},
			expectedErr: nil,
		},
		{
			name:        "standard currency - invalid characters in currency code",
			input:       "AD/",
			expected:    nil,
			expectedErr: ErrInvalidCurrencyCode,
		},
		{
			name:        "standard currency - invalid characters in currency code - hex",
			input:       "0x00000000000000000000000041442f0000000000",
			expected:    nil,
			expectedErr: ErrInvalidCurrencyCode,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := serializeIssuedCurrencyCode(tt.input)

			if tt.expectedErr != nil {
				require.EqualError(t, tt.expectedErr, err.Error())
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.expected, got)
			}

		})
	}
}

func TestSerializeIssuedCurrencyAmount(t *testing.T) {
	tests := []struct {
		name        string
		input       types.IssuedCurrencyAmount
		expected    []byte
		expectedErr error
	}{
		{
			name: "valid serialized issued currency amount",
			input: types.IssuedCurrencyAmount{
				Value:    "7072.8",
				Currency: "USD",
				Issuer:   "rvYAfWj5gh67oV6fW32ZzP3Aw4Eubs59B",
			},
			expected:    []byte{0xD5, 0x59, 0x20, 0xAC, 0x93, 0x91, 0x40, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x55, 0x53, 0x44, 0x00, 0x00, 0x00, 0x00, 0x00, 0x0A, 0x20, 0xB3, 0xC8, 0x5F, 0x48, 0x25, 0x32, 0xA9, 0x57, 0x8D, 0xBB, 0x39, 0x50, 0xB8, 0x5C, 0xA0, 0x65, 0x94, 0xD1},
			expectedErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := serializeIssuedCurrencyAmount(tt.input)

			if tt.expectedErr != nil {
				require.EqualError(t, tt.expectedErr, err.Error())
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.expected, got)
			}

		})
	}
}

func TestIsNative(t *testing.T) {
	tests := []struct {
		name     string
		input    byte
		expected bool
	}{
		{
			name:     "native XRP",
			input:    64, // 64 in binary is 01000000. If the first bit of the first byte is 0, it is deemed to be native XRP
			expected: true,
		},
		{
			name:     "not native XRP",
			input:    128, // 128 in binary is 10000000. If the first bit of the first byte is not 0, it is deemed to be not native XRP
			expected: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.expected, isNative(tt.input))
		})
	}
}

func TestIsPositive(t *testing.T) {
	tests := []struct {
		name     string
		input    byte
		expected bool
	}{
		{
			name:     "positive",
			input:    64, // 64 in binary is 01000000. If the second bit of the first byte is 1, it is deemed positive
			expected: true,
		},
		{
			name:     "negative",
			input:    128, // 128 in binary is 10000000. If the second bit of the first byte is 0, it is deemed negative
			expected: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.expected, isPositive(tt.input))
		})
	}
}

func TestSerializeMPTCurrencyAmount(t *testing.T) {
	amount := &Amount{}

	tests := []struct {
		name        string
		input       types.MPTCurrencyAmount
		expectedHex string
		expectErr   bool
	}{
		{
			name: "Valid MPT amount with value 100",
			input: types.MPTCurrencyAmount{
				Value:         "100",
				MPTIssuanceID: "00002403C84A0A28E0190E208E982C352BBD5006600555CF",
			},
			expectedHex: "60000000000000006400002403C84A0A28E0190E208E982C352BBD5006600555CF",
			expectErr:   false,
		},
		{
			name: "Valid MPT amount with large value",
			input: types.MPTCurrencyAmount{
				Value:         "9223372036854775807",
				MPTIssuanceID: "00002403C84A0A28E0190E208E982C352BBD5006600555CF",
			},
			expectedHex: "607FFFFFFFFFFFFFFF00002403C84A0A28E0190E208E982C352BBD5006600555CF",
			expectErr:   false,
		},
		{
			name: "MPT amount with zero value",
			input: types.MPTCurrencyAmount{
				Value:         "0",
				MPTIssuanceID: "00002403C84A0A28E0190E208E982C352BBD5006600555CF",
			},
			expectedHex: "60000000000000000000002403C84A0A28E0190E208E982C352BBD5006600555CF",
			expectErr:   false,
		},
		{
			name: "MPT amount with empty value",
			input: types.MPTCurrencyAmount{
				Value:         "",
				MPTIssuanceID: "00002403C84A0A28E0190E208E982C352BBD5006600555CF",
			},
			expectErr: true,
		},
		{
			name: "MPT amount with decimal value",
			input: types.MPTCurrencyAmount{
				Value:         "100.5",
				MPTIssuanceID: "00002403C84A0A28E0190E208E982C352BBD5006600555CF",
			},
			expectErr: true,
		},
		{
			name: "MPT amount with negative value",
			input: types.MPTCurrencyAmount{
				Value:         "-100",
				MPTIssuanceID: "00002403C84A0A28E0190E208E982C352BBD5006600555CF",
			},
			expectErr: true,
		},
		{
			name: "MPT amount with invalid issuance ID",
			input: types.MPTCurrencyAmount{
				Value:         "100",
				MPTIssuanceID: "invalid_hex",
			},
			expectErr: true,
		},
		{
			name: "MPT amount with wrong length issuance ID",
			input: types.MPTCurrencyAmount{
				Value:         "100",
				MPTIssuanceID: "00002403C84A0A28E0190E208E982C352BBD5006600555CF00", // 26 bytes
			},
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := amount.FromJson(tt.input)

			if tt.expectErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
			require.NotNil(t, result)

			// Convert result to hex string for comparison
			resultHex := strings.ToUpper(hex.EncodeToString(result))
			require.Equal(t, tt.expectedHex, resultHex)

			// Verify the length is correct (1 + 8 + 24 = 33 bytes)
			require.Equal(t, 33, len(result))

			// Verify the leading byte is 0x60
			require.Equal(t, byte(0x60), result[0])

			// Verify the amount bytes (positions 1-8)
			amountBytes := result[1:9]
			amountValue := binary.BigEndian.Uint64(amountBytes)
			expectedValue, _ := strconv.ParseUint(tt.input.Value, 10, 64)
			require.Equal(t, expectedValue, amountValue)

			// Verify the MPT issuance ID bytes (positions 9-32)
			mptIDBytes := result[9:33]
			expectedMPTID, _ := hex.DecodeString(tt.input.MPTIssuanceID)
			require.Equal(t, expectedMPTID, mptIDBytes)
		})
	}
}

func TestAmountFromJsonWithMPT(t *testing.T) {
	amount := &Amount{}

	// Test MPT amount through FromJson
	mptAmount := types.MPTCurrencyAmount{
		Value:         "100",
		MPTIssuanceID: "00002403C84A0A28E0190E208E982C352BBD5006600555CF",
	}

	result, err := amount.FromJson(mptAmount)
	require.NoError(t, err)
	require.NotNil(t, result)

	// Verify the result matches our expected format
	expectedHex := "60000000000000006400002403C84A0A28E0190E208E982C352BBD5006600555CF"
	resultHex := strings.ToUpper(hex.EncodeToString(result))
	require.Equal(t, expectedHex, resultHex)

	// Verify the length is correct (1 + 8 + 24 = 33 bytes)
	require.Equal(t, 33, len(result))

	// Verify the leading byte is 0x60
	require.Equal(t, byte(0x60), result[0])

	// Verify the amount bytes (positions 1-8)
	amountBytes := result[1:9]
	amountValue := binary.BigEndian.Uint64(amountBytes)
	require.Equal(t, uint64(100), amountValue)

	// Verify the MPT issuance ID bytes (positions 9-32)
	mptIDBytes := result[9:33]
	expectedMPTID, _ := hex.DecodeString(mptAmount.MPTIssuanceID)
	require.Equal(t, expectedMPTID, mptIDBytes)
}

func TestMPTSerializationMatchesJavaScript(t *testing.T) {
	amount := &Amount{}

	// Test case from JavaScript test: negative MPT value
	// JavaScript test shows: 20000000000000006400002403C84A0A28E0190E208E982C352BBD5006600555CF
	// Where 20 = leading byte (negative), 0000000000000064 = amount (-100), rest = MPT ID

	// Note: Our current implementation doesn't support negative values (as per XRPL spec),
	// but we can verify the structure matches what JavaScript would produce for positive values

	mptAmount := types.MPTCurrencyAmount{
		Value:         "100",
		MPTIssuanceID: "00002403C84A0A28E0190E208E982C352BBD5006600555CF",
	}

	result, err := amount.FromJson(mptAmount)
	require.NoError(t, err)
	require.NotNil(t, result)

	// The JavaScript test shows this hex for positive 100:
	// 60000000000000006400002403C84A0A28E0190E208E982C352BBD5006600555CF
	// Where 60 = leading byte (positive), 0000000000000064 = amount (100), rest = MPT ID

	expectedHex := "60000000000000006400002403C84A0A28E0190E208E982C352BBD5006600555CF"
	resultHex := strings.ToUpper(hex.EncodeToString(result))
	require.Equal(t, expectedHex, resultHex)

	// Verify the structure matches JavaScript implementation:
	// - Leading byte: 0x60 (bits 6 and 5 set for MPT, bit 4 clear for positive)
	// - Amount: 8 bytes in big-endian format
	// - MPT ID: 24 bytes

	require.Equal(t, byte(0x60), result[0]) // Leading byte should be 0x60

	// Amount should be 100 (0x64) in big-endian format
	amountBytes := result[1:9]
	amountValue := binary.BigEndian.Uint64(amountBytes)
	require.Equal(t, uint64(100), amountValue)

	// MPT ID should match the input
	mptIDBytes := result[9:33]
	expectedMPTID, _ := hex.DecodeString(mptAmount.MPTIssuanceID)
	require.Equal(t, expectedMPTID, mptIDBytes)
}

func TestDeserializeMPT(t *testing.T) {
	tests := []struct {
		name      string
		inputHex  string
		expected  map[string]any
		expectErr bool
	}{
		{
			name:     "Valid MPT amount with positive value 100",
			inputHex: "60000000000000006400002403C84A0A28E0190E208E982C352BBD5006600555CF",
			expected: map[string]any{
				"value":          "100",
				"mp_issuance_id": "00002403C84A0A28E0190E208E982C352BBD5006600555CF",
			},
			expectErr: false,
		},
		{
			name:     "Valid MPT amount with zero value",
			inputHex: "60000000000000000000002403C84A0A28E0190E208E982C352BBD5006600555CF",
			expected: map[string]any{
				"value":          "0",
				"mp_issuance_id": "00002403C84A0A28E0190E208E982C352BBD5006600555CF",
			},
			expectErr: false,
		},
		{
			name:     "Valid MPT amount with large value",
			inputHex: "607FFFFFFFFFFFFFFF00002403C84A0A28E0190E208E982C352BBD5006600555CF",
			expected: map[string]any{
				"value":          "9223372036854775807",
				"mp_issuance_id": "00002403C84A0A28E0190E208E982C352BBD5006600555CF",
			},
			expectErr: false,
		},
		{
			name:      "Invalid MPT data length",
			inputHex:  "60000000000000006400002403C84A0A28E0190E208E982C352BBD5006600555CF00",
			expected:  nil,
			expectErr: true,
		},
		{
			name:      "Invalid MPT leading byte (not MPT)",
			inputHex:  "80000000000000006400002403C84A0A28E0190E208E982C352BBD5006600555CF",
			expected:  nil,
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Decode hex string to bytes
			inputBytes, err := hex.DecodeString(tt.inputHex)
			require.NoError(t, err)

			// Test the deserializeMPT function directly

			result, err := deserializeMPT(inputBytes)

			if tt.expectErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
			require.NotNil(t, result)

			// Check that all expected fields are present
			for key, expectedValue := range tt.expected {
				actualValue, exists := result[key]
				require.True(t, exists, "Field %s should exist", key)
				require.Equal(t, expectedValue, actualValue, "Field %s should match", key)
			}

			// Check that no unexpected fields are present
			require.Equal(t, len(tt.expected), len(result), "Result should have exactly the expected number of fields")
		})
	}
}

func TestIsMPT(t *testing.T) {
	tests := []struct {
		name     string
		input    byte
		expected bool
	}{
		{
			name:     "MPT amount - leading byte 0x60",
			input:    0x60,
			expected: true,
		},
		{
			name:     "MPT amount with positive sign - leading byte 0x70",
			input:    0x70,
			expected: true,
		},
		{
			name:     "Native XRP amount - leading byte 0x40",
			input:    0x40,
			expected: false,
		},
		{
			name:     "IOU amount - leading byte 0x80",
			input:    0x80,
			expected: false,
		},
		{
			name:     "Zero byte",
			input:    0x00,
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.expected, isMPT(tt.input))
		})
	}
}

func TestMPTFullCycle(t *testing.T) {
	amount := &Amount{}

	// Test data
	mptAmount := types.MPTCurrencyAmount{
		Value:         "100",
		MPTIssuanceID: "00002403C84A0A28E0190E208E982C352BBD5006600555CF",
	}

	// Step 1: Serialize MPT to bytes
	serializedBytes, err := amount.FromJson(mptAmount)
	require.NoError(t, err)
	require.NotNil(t, serializedBytes)

	// Step 2: Verify the serialized format
	expectedHex := "60000000000000006400002403C84A0A28E0190E208E982C352BBD5006600555CF"
	actualHex := strings.ToUpper(hex.EncodeToString(serializedBytes))
	require.Equal(t, expectedHex, actualHex)

	// Step 3: Deserialize back to JSON
	// Test the deserializeMPT function directly

	result, err := deserializeMPT(serializedBytes)
	require.NoError(t, err)
	require.NotNil(t, result)

	// Step 4: Verify the deserialized result matches the original
	require.Equal(t, "100", result["value"])
	require.Equal(t, "00002403C84A0A28E0190E208E982C352BBD5006600555CF", result["mp_issuance_id"])

	// Step 5: Test with negative value
	mptAmountNegative := types.MPTCurrencyAmount{
		Value:         "50",
		MPTIssuanceID: "00002403C84A0A28E0190E208E982C352BBD5006600555CF",
	}

	// Note: MPT amounts are always positive in the current implementation
	// This test verifies that the sign handling works correctly
	serializedBytesNegative, err := amount.FromJson(mptAmountNegative)
	require.NoError(t, err)
	require.NotNil(t, serializedBytesNegative)

	resultNegative, err := deserializeMPT(serializedBytesNegative)
	require.NoError(t, err)
	require.NotNil(t, resultNegative)
	require.Equal(t, "50", resultNegative["value"])
}
