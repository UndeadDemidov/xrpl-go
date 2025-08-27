package ledger

import (
	"errors"
)

var (
	// ledger object

	// ErrUnrecognizedLedgerObjectType is returned when an unrecognized ledger object type is encountered.
	ErrUnrecognizedLedgerObjectType = errors.New("unrecognized LedgerObject type")
	// ErrUnsupportedLedgerObjectType is returned when an unsupported ledger object type is encountered.
	ErrUnsupportedLedgerObjectType = errors.New("unsupported ledger object type")

	// oracle

	// ErrPriceDataScale is returned when the scale is greater than the maximum allowed.
	ErrPriceDataScale = errors.New("invalid price data scale")
	// ErrPriceDataAssetPriceAndScale is returned when the asset price and scale are not set together.
	ErrPriceDataAssetPriceAndScale = errors.New("asset price and scale must be set together")
	// ErrPriceDataBaseAsset is returned when the base asset is required but not set.
	ErrPriceDataBaseAsset = errors.New("base asset is required")
	// ErrPriceDataQuoteAsset is returned when the quote asset is required but not set.
	ErrPriceDataQuoteAsset = errors.New("quote asset is required")
)
