package system

import "errors"

// ErrNotImplemented declares error for method that isn't implemented
var ErrNotImplemented = errors.New("This method is not implemented")

// Operator defines reload, maintenance and shutdown interface
type Operator interface {
	Shutdown() error
}

// Handling implements simplest Operator interface
type Handling struct{}

// Shutdown operation implementation
func (h Handling) Shutdown() error {

	return ErrNotImplemented
}
