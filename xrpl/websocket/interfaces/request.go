// revive:disable:var-naming
// Package interfaces defines common interfaces for XRPL WebSocket.
package interfaces

// Request represents a generic XRPL WebSocket request.
// It is implemented by all request types that can be sent to the XRPL server.
type Request interface {
	Method() string
	Validate() error
	APIVersion() int
}
