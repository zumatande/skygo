package errs

import (
	"errors"
)

// ErrPanic is error response after recovery from panic
var ErrPanic = errors.New("error: system panic")

// ErrInternal represents an unhandled error in the business domain
// e.g. a network partition, or a DB error
type ErrInternal struct {
	Underlying error
}

func (eint ErrInternal) Error() string {
	return eint.Underlying.Error()
}

// ErrRequest indicates an error in the service request
// e.g. missing params
type ErrRequest struct {
	Underlying error
}

func (ereq ErrRequest) Error() string {
	return ereq.Underlying.Error()
}

// ErrDomain represents an error in the business domain
// e.g. a zero balance account requesting a debit
type ErrDomain struct {
	Underlying error
}

func (edom ErrDomain) Error() string {
	return edom.Underlying.Error()
}
