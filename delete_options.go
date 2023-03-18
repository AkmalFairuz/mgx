package mgx

import "go.mongodb.org/mongo-driver/mongo/options"

// DeleteOptions is a wrapper around mongo options.
type DeleteOptions struct {
	o *options.DeleteOptions
}

// opt returns the underlying mongo option.
func (o *DeleteOptions) opt() *options.DeleteOptions {
	return o.o
}

// Collation specifies a collation.
func (o *DeleteOptions) Collation(v *options.Collation) *DeleteOptions {
	o.o.SetCollation(v)
	return o
}

// Hint specifies the index to use.
func (o *DeleteOptions) Hint(v any) *DeleteOptions {
	o.o.SetHint(v)
	return o
}

// Comment specifies a string or document that will be included in server logs, profiling logs, and currentOp queries to help trace
// the operation.  The default value is nil, which means that no comment will be included in the logs.
func (o *DeleteOptions) Comment(v string) *DeleteOptions {
	o.o.SetComment(v)
	return o
}

// Let specifies variables accessible from the delete expression.
func (o *DeleteOptions) Let(v any) *DeleteOptions {
	o.o.SetLet(v)
	return o
}

// NewDeleteOptions creates a new DeleteOptions.
func NewDeleteOptions() *DeleteOptions {
	return &DeleteOptions{o: options.Delete()}
}
