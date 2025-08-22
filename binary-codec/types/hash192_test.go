package types

import (
	"testing"

	"github.com/CreatureDev/xrpl-go/binary-codec/serdes"
	"github.com/CreatureDev/xrpl-go/model/transactions/types"
	"github.com/stretchr/testify/assert"
)

func TestHash192_FromJson(t *testing.T) {
	h := NewHash192()

	// Test with valid hex string
	validHex := "1234567890ABCDEF1234567890ABCDEF1234567890ABCDEF"
	result, err := h.FromJson(validHex)
	assert.NoError(t, err)
	assert.Equal(t, 24, len(result))

	// Test with Hash192 type
	hash192 := types.Hash192(validHex)
	result, err = h.FromJson(hash192)
	assert.NoError(t, err)
	assert.Equal(t, 24, len(result))

	// Test with invalid hex string
	invalidHex := "invalid"
	_, err = h.FromJson(invalidHex)
	assert.Error(t, err)

	// Test with wrong length hex string
	wrongLengthHex := "1234567890ABCDEF1234567890ABCDEF1234567890ABCDEF1234567890ABCDEF"
	_, err = h.FromJson(wrongLengthHex)
	assert.Error(t, err)

	// Test with invalid type
	_, err = h.FromJson(123)
	assert.Error(t, err)
}

func TestHash192_ToJson(t *testing.T) {
	h := NewHash192()

	// Create a binary parser with test data
	testData := []byte{0x12, 0x34, 0x56, 0x78, 0x90, 0xAB, 0xCD, 0xEF,
		0x12, 0x34, 0x56, 0x78, 0x90, 0xAB, 0xCD, 0xEF,
		0x12, 0x34, 0x56, 0x78, 0x90, 0xAB, 0xCD, 0xEF}

	parser := serdes.NewBinaryParser(testData)

	result, err := h.ToJson(parser)
	assert.NoError(t, err)
	assert.Equal(t, "1234567890ABCDEF1234567890ABCDEF1234567890ABCDEF", result)
}

func TestHash192_GetLength(t *testing.T) {
	h := NewHash192()
	assert.Equal(t, 24, h.getLength())
}

func TestNewHash192(t *testing.T) {
	h := NewHash192()
	assert.NotNil(t, h)
	assert.IsType(t, &Hash192{}, h)
}
