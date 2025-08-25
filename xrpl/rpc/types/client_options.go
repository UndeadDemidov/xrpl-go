// Package types contains data structures for RPC client configuration and options.
//
//revive:disable:var-naming
package types

import (
	"github.com/Peersyst/xrpl-go/xrpl/wallet"
)

// SubmitOptions specifies options for submitting a single transaction via RPC.
type SubmitOptions struct {
	Autofill bool
	Wallet   *wallet.Wallet
	FailHard bool
}

// SubmitBatchOptions specifies options for submitting a batch of transactions via RPC.
type SubmitBatchOptions struct {
	Autofill bool
	Wallet   *wallet.Wallet
	FailHard bool
	NSigners uint64
}
