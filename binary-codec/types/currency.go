//revive:disable:var-naming
package types

import (
	"bytes"
	"encoding/hex"
	"errors"

	"github.com/Peersyst/xrpl-go/binary-codec/types/interfaces"
)

var (
	// ErrMissingCurrencyLengthOption is returned when no length option is
	// provided to Currency.ToJSON.
	ErrMissingCurrencyLengthOption = errors.New("missing length option for Currency.ToJSON")
)

// Currency handles encoding and decoding of currency values in the binary codec.
type Currency struct{}

// FromJSON parses a JSON value into its binary currency representation.
func (c *Currency) FromJSON(json any) ([]byte, error) {
	if str, ok := json.(string); ok {
		return c.fromString(str)
	}
	return nil, ErrInvalidCurrency
}

// ToJSON serializes a binary currency value into a JSON-compatible format.
// It requires a length option specifying the byte length to read.
func (c *Currency) ToJSON(p interfaces.BinaryParser, opts ...int) (any, error) {
	if len(opts) == 0 {
		return nil, ErrMissingCurrencyLengthOption
	}

	currencyBytes, err := p.ReadBytes(opts[0])
	if err != nil {
		return nil, err
	}

	if bytes.Equal(currencyBytes, XRPBytes) {
		return "XRP", nil
	}

	// Check if bytes has exactly 3 non-zero bytes at positions 12-14
	nonZeroCount := 0
	var currencyStr string
	for i := 0; i < len(currencyBytes); i++ {
		if currencyBytes[i] != 0 {
			if i >= 12 && i <= 14 {
				nonZeroCount++
				currencyStr += string(currencyBytes[i])
			} else {
				nonZeroCount = 0
				break
			}
		}
	}

	if nonZeroCount == 3 {
		return currencyStr, nil
	}

	return hex.EncodeToString(currencyBytes), nil
}

func (c *Currency) fromString(str string) ([]byte, error) {
	if len(str) == 3 {
		var bytes [20]byte
		if str != "XRP" {
			isoBytes := []byte(str)
			copy(bytes[12:], isoBytes)
		}
		return bytes[:], nil
	}

	bytes, err := hex.DecodeString(str)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}
