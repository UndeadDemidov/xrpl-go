package crypto

import "errors"

var (
	// keypair

	// ErrValidatorKeypairDerivation is returned when a validator keypair is attempted to be derived
	ErrValidatorKeypairDerivation = errors.New("validator keypair derivation not supported")
	// ErrInvalidPrivateKey is returned when a private key is invalid
	ErrInvalidPrivateKey = errors.New("invalid private key")
	// ErrInvalidMessage is returned when a message is required but not provided
	ErrInvalidMessage = errors.New("message is required")
	// ErrValidatorNotSupported is returned when a validator keypair is used with the ED25519 algorithm.
	ErrValidatorNotSupported = errors.New("validator keypairs can not use Ed25519")

	// der

	// ErrInvalidHexString is returned when the hex string is invalid.
	ErrInvalidHexString = errors.New("invalid hex string")
	// ErrInvalidDERNotEnoughData is returned when the DER data is not enough.
	ErrInvalidDERNotEnoughData = errors.New("invalid DER: not enough data")
	// ErrInvalidDERIntegerTag is returned when the DER integer tag is invalid.
	ErrInvalidDERIntegerTag = errors.New("invalid DER: expected integer tag")
	// ErrInvalidDERSignature is returned when the DER signature is invalid.
	ErrInvalidDERSignature = errors.New("invalid signature: incorrect length")
	// ErrLeftoverBytes is returned when there are leftover bytes after parsing the DER signature.
	ErrLeftoverBytes = errors.New("invalid signature: left bytes after parsing")
)
