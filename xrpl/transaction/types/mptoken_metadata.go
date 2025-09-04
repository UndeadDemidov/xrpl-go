package types

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
)

const (
	// MPTokenMetadataMaxSize defines the maximum allowed size in bytes for
	// multi-purpose token metadata when serialized to JSON.
	// https://github.com/XRPLF/XRPL-Standards/tree/master/XLS-0089d-multi-purpose-token-metadata-schema
	MPTokenMetadataMaxSize = 1024
)

// MPTokenMetadataUrl represents a URL reference within multi-purpose token metadata.
// It follows the XLS-0089d standard for multi-purpose token metadata schema.
type MPTokenMetadataUrl struct {
	Url   string `json:"url,omitempty"`   // The URL address
	Type  string `json:"type,omitempty"`  // The MIME type of the resource
	Title string `json:"title,omitempty"` // Human-readable title for the URL
}

// MPTokenMetadata represents the metadata structure for multi-purpose tokens
// following the XLS-0089d standard. This structure defines the schema for
// token metadata that can be attached to XRPL tokens.
type MPTokenMetadata struct {
	Ticker         string               `json:"ticker,omitempty"`          // Short symbol for the token
	Name           string               `json:"name,omitempty"`            // Full name of the token
	Desc           string               `json:"desc,omitempty"`            // Description of the token
	Icon           string               `json:"icon,omitempty"`            // URL to token icon image
	AssetClass     string               `json:"asset_class,omitempty"`     // Primary classification of the asset
	AssetSubclass  string               `json:"asset_subclass,omitempty"`  // Secondary classification of the asset
	IssuerName     string               `json:"issuer_name,omitempty"`     // Name of the token issuer
	Urls           []MPTokenMetadataUrl `json:"urls,omitempty"`            // Additional URL references
	AdditionalInfo json.RawMessage      `json:"additional_info,omitempty"` // Custom additional metadata
}

// MPTokenMetadataFromBlob parses a hex-encoded blob string into MPTokenMetadata.
// The blob should contain JSON data that conforms to the XLS-0089d standard
// for multi-purpose token metadata schema.
//
// Returns an error if the blob is not valid hex or if the JSON doesn't conform
// to the expected schema.
func MPTokenMetadataFromBlob(blob string) (*MPTokenMetadata, error) {
	b, err := hex.DecodeString(blob)
	if err != nil {
		return nil, fmt.Errorf("decode from blob in hex: %w", err)
	}
	m := MPTokenMetadata{}

	err = json.Unmarshal(b, &m)
	if err != nil {
		return nil, fmt.Errorf("metadata is not in XLS-0089d schema: %w", err)
	}
	return &m, nil
}

// Blob serializes the MPTokenMetadata to a hex-encoded string.
// The metadata is first marshaled to JSON and then encoded as a hex string.
//
// Returns an error if the JSON marshaling fails or if the resulting
// blob exceeds the maximum allowed size (MPTokenMetadataMaxSize).
func (m MPTokenMetadata) Blob() (string, error) {
	json, err := json.Marshal(m)
	if err != nil {
		return "", fmt.Errorf("marshal to json for blob: %w", err)
	}

	if len(json) > MPTokenMetadataMaxSize {
		return "", fmt.Errorf("blob is too large: %d", len(json))
	}

	return hex.EncodeToString(json), nil
}
