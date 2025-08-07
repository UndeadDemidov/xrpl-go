//revive:disable var-naming
package types

import (
	"github.com/Peersyst/xrpl-go/xrpl/wallet"
)

// SubmitOptions configures transaction submission options over WebSocket, including autofill, wallet and fail-hard.
type SubmitOptions struct {
	Autofill bool
	Wallet   *wallet.Wallet
	FailHard bool
}

// SubmitBatchOptions configures options for submitting a batch of transactions over WebSocket, including number of signers.
type SubmitBatchOptions struct {
	Autofill bool
	Wallet   *wallet.Wallet
	FailHard bool
	NSigners uint64
}
