package mgx

import "go.mongodb.org/mongo-driver/mongo/options"

// InsertOneOptions is a wrapper around mongo options.
type InsertOneOptions struct {
	o *options.InsertOneOptions
}

// opt returns the underlying mongo option.
func (o *InsertOneOptions) opt() *options.InsertOneOptions {
	return o.o
}

// BypassDocumentValidation specifies whether to bypass document validation during the insert operation.
func (o *InsertOneOptions) BypassDocumentValidation(v ...bool) *InsertOneOptions {
	x := true
	if len(v) > 0 {
		x = v[0]
	}
	o.o.SetBypassDocumentValidation(x)
	return o
}

// Comment specifies a string or document that will be included in server logs, profiling logs, and currentOp queries to help trace
// the operation.  The default value is nil, which means that no comment will be included in the logs.
func (o *InsertOneOptions) Comment(v string) *InsertOneOptions {
	o.o.SetComment(v)
	return o
}

// NewInsertOneOptions creates a new InsertOneOptions.
func NewInsertOneOptions() *InsertOneOptions {
	return &InsertOneOptions{o: options.InsertOne()}
}
