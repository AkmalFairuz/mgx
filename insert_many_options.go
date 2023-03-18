package mgx

import "go.mongodb.org/mongo-driver/mongo/options"

// InsertManyOptions is a wrapper around mongo options.
type InsertManyOptions struct {
	o *options.InsertManyOptions
}

// opt returns the underlying mongo option.
func (o *InsertManyOptions) opt() *options.InsertManyOptions {
	return o.o
}

// BypassDocumentValidation specifies whether to bypass document validation during the insert operation.
func (o *InsertManyOptions) BypassDocumentValidation(v ...bool) *InsertManyOptions {
	x := true
	if len(v) > 0 {
		x = v[0]
	}
	o.o.SetBypassDocumentValidation(x)
	return o
}

// Ordered specifies whether the operations in the bulk write should be executed in order.
func (o *InsertManyOptions) Ordered(v ...bool) *InsertManyOptions {
	x := true
	if len(v) > 0 {
		x = v[0]
	}
	o.o.SetOrdered(x)
	return o
}

// Comment specifies a string or document that will be included in server logs, profiling logs, and currentOp queries to help trace
// the operation.  The default value is nil, which means that no comment will be included in the logs.
func (o *InsertManyOptions) Comment(v string) *InsertManyOptions {
	o.o.SetComment(v)
	return o
}

// NewInsertManyOptions creates a new InsertManyOptions.
func NewInsertManyOptions() *InsertManyOptions {
	return &InsertManyOptions{o: options.InsertMany()}
}
