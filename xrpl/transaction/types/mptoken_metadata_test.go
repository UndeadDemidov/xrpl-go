package types

import (
	"encoding/json"
	"testing"
)

func TestMPTokenMetadataFromBlob(t *testing.T) {
	tests := []struct {
		name        string
		blob        string
		expected    *MPTokenMetadata
		expectError bool
		errorMsg    string
	}{
		{
			name: "valid metadata with all fields",
			blob: "7b227469636b6572223a22425443222c226e616d65223a22426974636f696e222c2264657363223a2241206469676974616c2063757272656e6379222c2269636f6e223a2268747470733a2f2f6578616d706c652e636f6d2f626974636f696e2e706e67222c2261737365745f636c617373223a2263727970746f63757272656e6379222c2261737365745f737562636c617373223a22626974636f696e222c226973737565725f6e616d65223a225361746f736869204e616b616d6f746f222c2275726c73223a5b7b2275726c223a2268747470733a2f2f626974636f696e2e6f7267222c2274797065223a22746578742f68746d6c222c227469746c65223a22426974636f696e204f6666696369616c227d5d2c226164646974696f6e616c5f696e666f223a7b2276657273696f6e223a22312e30227d7d",
			expected: &MPTokenMetadata{
				Ticker:        "BTC",
				Name:          "Bitcoin",
				Desc:          "A digital currency",
				Icon:          "https://example.com/bitcoin.png",
				AssetClass:    "cryptocurrency",
				AssetSubclass: "bitcoin",
				IssuerName:    "Satoshi Nakamoto",
				Urls: []MPTokenMetadataUrl{
					{
						Url:   "https://bitcoin.org",
						Type:  "text/html",
						Title: "Bitcoin Official",
					},
				},
				AdditionalInfo: json.RawMessage(`{"version":"1.0"}`),
			},
			expectError: false,
		},
		{
			name: "valid metadata with minimal fields",
			blob: "7b227469636b6572223a22455448222c226e616d65223a22457468657265756d227d",
			expected: &MPTokenMetadata{
				Ticker: "ETH",
				Name:   "Ethereum",
			},
			expectError: false,
		},
		{
			name:        "valid metadata with empty fields",
			blob:        "7b7d",
			expected:    &MPTokenMetadata{},
			expectError: false,
		},
		{
			name: "valid metadata with urls array",
			blob: "7b227469636b6572223a2255534454222c226e616d65223a2255534420546574686572222c2275726c73223a5b7b2275726c223a2268747470733a2f2f6578616d706c652e636f6d222c2274797065223a226170706c69636174696f6e2f6a736f6e227d2c7b2275726c223a2268747470733a2f2f646f63732e6578616d706c652e636f6d222c227469746c65223a22446f63756d656e746174696f6e227d5d7d",
			expected: &MPTokenMetadata{
				Ticker: "USDT",
				Name:   "USD Tether",
				Urls: []MPTokenMetadataUrl{
					{
						Url:  "https://example.com",
						Type: "application/json",
					},
					{
						Url:   "https://docs.example.com",
						Title: "Documentation",
					},
				},
			},
			expectError: false,
		},
		{
			name:        "invalid hex string",
			blob:        "invalid-hex-string",
			expected:    nil,
			expectError: true,
			errorMsg:    "decode from blob in hex",
		},
		{
			name:        "empty hex string",
			blob:        "",
			expected:    nil,
			expectError: true,
			errorMsg:    "metadata is not in XLS-0089d schema",
		},
		{
			name:        "invalid JSON in hex",
			blob:        "7b227469636b6572223a22425443222c226e616d65223a22426974636f696e222c2264657363223a2241206469676974616c2063757272656e6379222c2269636f6e223a2268747470733a2f2f6578616d706c652e636f6d2f626974636f696e2e706e67222c2261737365745f636c617373223a2263727970746f63757272656e6379222c2261737365745f737562636c617373223a22626974636f696e222c226973737565725f6e616d65223a225361746f736869204e616b616d6f746f222c2275726c73223a5b7b2275726c223a2268747470733a2f2f626974636f696e2e6f7267222c2274797065223a22746578742f68746d6c222c227469746c65223a22426974636f696e204f6666696369616c227d5d2c226164646974696f6e616c5f696e666f223a7b2276657273696f6e223a22312e30227d", // missing closing brace
			expected:    nil,
			expectError: true,
			errorMsg:    "metadata is not in XLS-0089d schema",
		},
		{
			name:        "hex string with odd length",
			blob:        "7b227469636b6572223a22425443222c226e616d65223a22426974636f696e222c2264657363223a2241206469676974616c2063757272656e6379222c2269636f6e223a2268747470733a2f2f6578616d706c652e636f6d2f626974636f696e2e706e67222c2261737365745f636c617373223a2263727970746f63757272656e6379222c2261737365745f737562636c617373223a22626974636f696e222c226973737565725f6e616d65223a225361746f736869204e616b616d6f746f222c2275726c73223a5b7b2275726c223a2268747470733a2f2f626974636f696e2e6f7267222c2274797065223a22746578742f68746d6c222c227469746c65223a22426974636f696e204f6666696369616c227d5d2c226164646974696f6e616c5f696e666f223a7b2276657273696f6e223a22312e30227d7", // odd length
			expected:    nil,
			expectError: true,
			errorMsg:    "decode from blob in hex",
		},
		{
			name:        "non-hex characters",
			blob:        "7g227469636b6572223a22425443222c226e616d65223a22426974636f696e227d",
			expected:    nil,
			expectError: true,
			errorMsg:    "decode from blob in hex",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := MPTokenMetadataFromBlob(tt.blob)

			if tt.expectError {
				if err == nil {
					t.Errorf("expected error but got none")
					return
				}
				if tt.errorMsg != "" && !contains(err.Error(), tt.errorMsg) {
					t.Errorf("expected error message to contain %q, got %q", tt.errorMsg, err.Error())
				}
				return
			}

			if err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}

			if result == nil {
				t.Errorf("expected result but got nil")
				return
			}

			// Compare the result with expected
			if !compareMPTokenMetadata(result, tt.expected) {
				t.Errorf("expected %+v, got %+v", tt.expected, result)
			}
		})
	}
}

// Helper function to check if a string contains a substring
func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(substr) == 0 ||
		(len(s) > len(substr) && (s[:len(substr)] == substr || s[len(s)-len(substr):] == substr ||
			contains(s[1:], substr))))
}

// Helper function to compare MPTokenMetadata structs
func compareMPTokenMetadata(a, b *MPTokenMetadata) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}

	if a.Ticker != b.Ticker ||
		a.Name != b.Name ||
		a.Desc != b.Desc ||
		a.Icon != b.Icon ||
		a.AssetClass != b.AssetClass ||
		a.AssetSubclass != b.AssetSubclass ||
		a.IssuerName != b.IssuerName {
		return false
	}

	// Compare URLs
	if len(a.Urls) != len(b.Urls) {
		return false
	}
	for i, url := range a.Urls {
		if url.Url != b.Urls[i].Url ||
			url.Type != b.Urls[i].Type ||
			url.Title != b.Urls[i].Title {
			return false
		}
	}

	// Compare AdditionalInfo
	if string(a.AdditionalInfo) != string(b.AdditionalInfo) {
		return false
	}

	return true
}

func TestMPTokenMetadata_Blob(t *testing.T) {
	tests := []struct {
		name        string
		metadata    MPTokenMetadata
		expected    string
		expectError bool
		errorMsg    string
	}{
		{
			name: "valid metadata with all fields",
			metadata: MPTokenMetadata{
				Ticker:        "BTC",
				Name:          "Bitcoin",
				Desc:          "A digital currency",
				Icon:          "https://example.com/bitcoin.png",
				AssetClass:    "cryptocurrency",
				AssetSubclass: "bitcoin",
				IssuerName:    "Satoshi Nakamoto",
				Urls: []MPTokenMetadataUrl{
					{
						Url:   "https://bitcoin.org",
						Type:  "text/html",
						Title: "Bitcoin Official",
					},
				},
				AdditionalInfo: json.RawMessage(`{"version":"1.0"}`),
			},
			expected:    "7b227469636b6572223a22425443222c226e616d65223a22426974636f696e222c2264657363223a2241206469676974616c2063757272656e6379222c2269636f6e223a2268747470733a2f2f6578616d706c652e636f6d2f626974636f696e2e706e67222c2261737365745f636c617373223a2263727970746f63757272656e6379222c2261737365745f737562636c617373223a22626974636f696e222c226973737565725f6e616d65223a225361746f736869204e616b616d6f746f222c2275726c73223a5b7b2275726c223a2268747470733a2f2f626974636f696e2e6f7267222c2274797065223a22746578742f68746d6c222c227469746c65223a22426974636f696e204f6666696369616c227d5d2c226164646974696f6e616c5f696e666f223a7b2276657273696f6e223a22312e30227d7d",
			expectError: false,
		},
		{
			name: "valid metadata with minimal fields",
			metadata: MPTokenMetadata{
				Ticker: "ETH",
				Name:   "Ethereum",
			},
			expected:    "7b227469636b6572223a22455448222c226e616d65223a22457468657265756d227d",
			expectError: false,
		},
		{
			name:        "empty metadata",
			metadata:    MPTokenMetadata{},
			expected:    "7b7d",
			expectError: false,
		},
		{
			name: "metadata with only ticker",
			metadata: MPTokenMetadata{
				Ticker: "USDT",
			},
			expected:    "7b227469636b6572223a2255534454227d",
			expectError: false,
		},
		{
			name: "metadata with only name",
			metadata: MPTokenMetadata{
				Name: "USD Tether",
			},
			expected:    "7b226e616d65223a2255534420546574686572227d",
			expectError: false,
		},
		{
			name: "metadata with urls array",
			metadata: MPTokenMetadata{
				Ticker: "USDT",
				Name:   "USD Tether",
				Urls: []MPTokenMetadataUrl{
					{
						Url:  "https://example.com",
						Type: "application/json",
					},
					{
						Url:   "https://docs.example.com",
						Title: "Documentation",
					},
				},
			},
			expected:    "7b227469636b6572223a2255534454222c226e616d65223a2255534420546574686572222c2275726c73223a5b7b2275726c223a2268747470733a2f2f6578616d706c652e636f6d222c2274797065223a226170706c69636174696f6e2f6a736f6e227d2c7b2275726c223a2268747470733a2f2f646f63732e6578616d706c652e636f6d222c227469746c65223a22446f63756d656e746174696f6e227d5d7d",
			expectError: false,
		},
		{
			name: "metadata with additional info",
			metadata: MPTokenMetadata{
				Ticker:         "CUSTOM",
				Name:           "Custom Token",
				AdditionalInfo: json.RawMessage(`{"custom_field":"value","number":123}`),
			},
			expected:    "7b227469636b6572223a22435553544f4d222c226e616d65223a22437573746f6d20546f6b656e222c226164646974696f6e616c5f696e666f223a7b22637573746f6d5f6669656c64223a2276616c7565222c226e756d626572223a3132337d7d",
			expectError: false,
		},
		{
			name: "metadata with special characters",
			metadata: MPTokenMetadata{
				Ticker: "TØKËN",
				Name:   "Tøkën with spéciál chârs",
				Desc:   "Description with émojis 🚀 and symbols & < > \" '",
			},
			expected:    "", // We'll validate this dynamically since encoding can vary
			expectError: false,
		},
		{
			name: "metadata with empty urls array",
			metadata: MPTokenMetadata{
				Ticker: "EMPTY",
				Name:   "Empty URLs",
				Urls:   []MPTokenMetadataUrl{},
			},
			expected:    "7b227469636b6572223a22454d505459222c226e616d65223a22456d7074792055524c73227d",
			expectError: false,
		},
		{
			name: "metadata with empty additional info",
			metadata: MPTokenMetadata{
				Ticker:         "EMPTY_INFO",
				Name:           "Empty Additional Info",
				AdditionalInfo: json.RawMessage(`{}`),
			},
			expected:    "7b227469636b6572223a22454d5054595f494e464f222c226e616d65223a22456d707479204164646974696f6e616c20496e666f222c226164646974696f6e616c5f696e666f223a7b7d7d",
			expectError: false,
		},
		{
			name: "metadata exceeding max size",
			metadata: MPTokenMetadata{
				Ticker: "LARGE",
				Name:   "Large Token",
				Desc:   generateLargeString(2000), // Generate a string larger than MPTokenMetadataMaxSize
			},
			expected:    "",
			expectError: true,
			errorMsg:    "blob is too large",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := tt.metadata.Blob()

			if tt.expectError {
				if err == nil {
					t.Errorf("expected error but got none")
					return
				}
				if tt.errorMsg != "" && !contains(err.Error(), tt.errorMsg) {
					t.Errorf("expected error message to contain %q, got %q", tt.errorMsg, err.Error())
				}
				return
			}

			if err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}

			// For special characters test, we validate dynamically since encoding can vary
			if tt.name == "metadata with special characters" {
				// Just verify it's not empty and can be decoded
				if result == "" {
					t.Errorf("expected non-empty result for special characters test")
				}
			} else if result != tt.expected {
				t.Errorf("expected %q, got %q", tt.expected, result)
			}

			// Verify that the blob can be decoded back to the original metadata
			decoded, err := MPTokenMetadataFromBlob(result)
			if err != nil {
				t.Errorf("failed to decode generated blob: %v", err)
				return
			}

			if !compareMPTokenMetadata(decoded, &tt.metadata) {
				t.Errorf("decoded metadata doesn't match original. Original: %+v, Decoded: %+v", tt.metadata, decoded)
			}
		})
	}
}

// Helper function to generate a large string for testing size limits
func generateLargeString(size int) string {
	result := make([]byte, size)
	for i := range result {
		result[i] = 'A' + byte(i%26)
	}
	return string(result)
}
