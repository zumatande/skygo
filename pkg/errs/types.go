package errs

// ErrInternal represents an unhandled error in the business domain
// e.g. a network partition, or a DB error
type ErrInternal struct {
	error
}

// ErrRequest indicates an error in the service request
// e.g. missing params
type ErrRequest struct {
	error
}

// ErrDomain represents an error in the business domain
// e.g. a zero balance account requesting a debit
type ErrDomain struct {
	error
}
