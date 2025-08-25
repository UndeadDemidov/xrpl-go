package transaction

import (
	"errors"
	"fmt"
)

// Common flags for AMM transactions (Deposit and Withdraw).
const (
	// Perform a double-asset withdrawal/deposit and receive the specified amount of LP Tokens.
	tfLPToken uint32 = 65536
	// Perform a single-asset withdrawal/deposit with a specified amount of the asset to deposit.
	tfSingleAsset uint32 = 524288
	// Perform a double-asset withdrawal/deposit with specified amounts of both assets.
	tfTwoAsset uint32 = 1048576
	// Perform a single-asset withdrawal/deposit and receive the specified amount of LP Tokens.
	tfOneAssetLPToken uint32 = 2097152
	// Perform a single-asset withdrawal/deposit with a specified effective price.
	tfLimitLPToken uint32 = 4194304

	// AmmMaxTradingFee is the maximum trading fee; a value of 1000 corresponds to a 1% fee.
	AmmMaxTradingFee = 1000
)

var (
	// ErrAMMTradingFeeTooHigh is returned when the specified trading fee exceeds AmmMaxTradingFee.
	ErrAMMTradingFeeTooHigh = fmt.Errorf("trading fee is too high, max value is %d", AmmMaxTradingFee)
	// ErrAMMMustSetAmountWithAmount2 is returned when Amount2 is set without Amount.
	ErrAMMMustSetAmountWithAmount2 = errors.New("must set Amount with Amount2")
	// ErrAMMMustSetAmountWithEPrice is returned when EPrice is set without Amount.
	ErrAMMMustSetAmountWithEPrice = errors.New("must set Amount with EPrice")
)
