package swagger

import (
	"io"
	"mime/multipart"
)

// File represents an uploaded file.
type File struct {
	Data   multipart.File
	Header *multipart.FileHeader
}

// OperationHandlerFunc an adapter for a function to the OperationHandler interface
type OperationHandlerFunc func(interface{}) (interface{}, error)

// Handle implements the operation handler interface
func (s OperationHandlerFunc) Handle(data interface{}) (interface{}, error) {
	return s(data)
}

// OperationHandler a handler for a swagger operation
type OperationHandler interface {
	Handle(interface{}) (interface{}, error)
}

// ConsumerFunc represents a function that can be used as a consumer
type ConsumerFunc func(io.Reader, interface{}) error

// Consume consumes the reader into the data parameter
func (fn ConsumerFunc) Consume(reader io.Reader, data interface{}) error {
	return fn(reader, data)
}

// Consumer implementations know how to bind the values on the provided interface to
// data provided by the request body
type Consumer interface {
	// Consume performs the binding of request values
	Consume(io.Reader, interface{}) error
}

// ProducerFunc represents a function that can be used as a producer
type ProducerFunc func(io.Writer, interface{}) error

// Produce produces the response for the provided data
func (f ProducerFunc) Produce(writer io.Writer, data interface{}) error {
	return f(writer, data)
}

// Producer implementations know how to turn the provided interface into a valid
// HTTP response
type Producer interface {
	// Produce writes to the http response
	Produce(io.Writer, interface{}) error
}

// AuthenticatorFunc turns a function into an authenticator
type AuthenticatorFunc func(interface{}) (bool, interface{}, error)

// Authenticate authenticates the request with the provided data
func (f AuthenticatorFunc) Authenticate(params interface{}) (bool, interface{}, error) {
	return f(params)
}

// Authenticator represents an authentication strategy
// implementations of Authenticator know how to authenticate the
// request data and translate that into a valid principal object or an error
type Authenticator interface {
	Authenticate(interface{}) (bool, interface{}, error)
}
