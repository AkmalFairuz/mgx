package mgx

import (
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

// CountOptions is a wrapper around mongo options.
type CountOptions struct {
	o *options.CountOptions
}

// opt returns the underlying mongo option.
func (o *CountOptions) opt() *options.CountOptions {
	return o.o
}

// Collation specifies a collation.
func (o *CountOptions) Collation(v *options.Collation) *CountOptions {
	o.o.SetCollation(v)
	return o
}

// Hint specifies the index to use.
func (o *CountOptions) Hint(v any) *CountOptions {
	o.o.SetHint(v)
	return o
}

// Limit specifies the maximum number of documents to count.
func (o *CountOptions) Limit(v int64) *CountOptions {
	o.o.SetLimit(v)
	return o
}

// MaxTime specifies the maximum amount of time to allow the query to run.
func (o *CountOptions) MaxTime(v time.Duration) *CountOptions {
	o.o.SetMaxTime(v)
	return o
}

// Skip specifies the number of documents to skip before returning.
func (o *CountOptions) Skip(v int64) *CountOptions {
	o.o.SetSkip(v)
	return o
}

// Comment specifies a string or document that will be included in server logs, profiling logs, and currentOp queries to help trace
// the operation.  The default value is nil, which means that no comment will be included in the logs.
func (o *CountOptions) Comment(v string) *CountOptions {
	o.o.SetComment(v)
	return o
}

// NewCountOptions creates a new CountOptions.
func NewCountOptions() *CountOptions {
	return &CountOptions{o: options.Count()}
}
