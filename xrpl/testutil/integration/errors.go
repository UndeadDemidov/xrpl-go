package integration

import "errors"

var (
	// ErrFailedToFundWallet is returned when funding a wallet fails after exceeding retry limit.
	ErrFailedToFundWallet = errors.New("failed to fund wallet")
)
