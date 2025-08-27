package faucet

import "errors"

var (
	// ErrMarshalPayload is returned when the faucet payload cannot be marshaled to JSON.
	ErrMarshalPayload = errors.New("failed to marshal faucet payload")

	// ErrCreateRequest is returned when the faucet request cannot be created.
	ErrCreateRequest = errors.New("failed to create faucet request")

	// ErrSendRequest is returned when the faucet request cannot be sent.
	ErrSendRequest = errors.New("failed to send POST request to faucet")

	// ErrUnexpectedStatusCode is returned when the faucet responds with a non-200 status code.
	ErrUnexpectedStatusCode = errors.New("unexpected faucet response status code")
)
