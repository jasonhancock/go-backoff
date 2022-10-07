package backoff

// Retrier determines if an error exhibits the behavior of an error that is safe
// to retry. Useful if you don't want some types of errors to be retried (ie.
// something that causes an HTTP 400 is likely to be successful on a subsequent
// request if the request is unchanged).
type Retrier interface {
	Retry() bool
}
