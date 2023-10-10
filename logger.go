package backoff

// Logger is the interface the Backoff will accept for logging messages.
type Logger interface {
	Warn(msg any, keyvals ...any)
}
