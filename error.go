package backoff

// Retrier determines if an error exhibits the behavior of an error that is safe
// to retry. Useful if you don't want some types of errors to be retried (ie.
// something that causes an HTTP 400 is likely to be successful on a subsequent
// request if the request is unchanged).
type Retrier interface {
	Retry() bool
}

// NonRetryableError wraps an err to add a Retry() method that conforms to
// Retrier interface which then prevents a backoff runner from retrying.
type NonRetryableError struct {
	err error
}

// Unwrap returns the underlying wrapped error.
func (e NonRetryableError) Unwrap() error {
	return e.err
}

// Error returns the error as a string.
func (e NonRetryableError) Error() string {
	return e.err.Error()
}

// Retry tells the backoff not to retry this error.
func (e NonRetryableError) Retry() bool {
	return false
}

// NewNonRetryable initializes a non retryable error.
func NewNonRetryable(err error) error {
	if err == nil {
		return nil
	}
	return &NonRetryableError{err: err}
}
