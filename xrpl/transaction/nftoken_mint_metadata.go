package transaction

import (
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
)

// NFTokenMintMetadata contains metadata for an NFTokenMint transaction, including the newly minted token ID and optional offer ID.
type NFTokenMintMetadata struct {
	TxObjMeta
	// rippled 1.11.0 or later
	NFTokenID *types.NFTokenID `json:"nftoken_id,omitempty"`
	// if Amount is present
	OfferID *types.Hash256 `json:"offer_id,omitempty"`
}

// TxMeta implements the TxMeta interface for NFTokenMintMetadata.
func (NFTokenMintMetadata) TxMeta() {}
