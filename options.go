package mgx

// Options is a wrapper around mongo options.
type Options[T any] interface {
	// opt returns the underlying mongo option.
	opt() T
}
