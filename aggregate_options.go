package mgx

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

// AggregateOptions is a wrapper around mongo options.
type AggregateOptions struct {
	o *options.AggregateOptions
}

// opt returns the underlying mongo option.
func (o *AggregateOptions) opt() *options.AggregateOptions {
	return o.o
}

// AllowDiskUse specifies whether to allow disk use.
func (o *AggregateOptions) AllowDiskUse(v ...bool) *AggregateOptions {
	x := true
	if len(v) > 0 {
		x = v[0]
	}
	o.o.SetAllowDiskUse(x)
	return o
}

// BatchSize specifies the number of documents to return per batch.
func (o *AggregateOptions) BatchSize(v int32) *AggregateOptions {
	o.o.SetBatchSize(v)
	return o
}

// BypassDocumentValidation specifies whether to bypass document validation during the aggregate operation.
func (o *AggregateOptions) BypassDocumentValidation(v ...bool) *AggregateOptions {
	x := true
	if len(v) > 0 {
		x = v[0]
	}
	o.o.SetBypassDocumentValidation(x)
	return o
}

// Collation specifies a collation.
func (o *AggregateOptions) Collation(v *Collation) *AggregateOptions {
	o.o.SetCollation(v.Collation)
	return o
}

// Comment specifies a string or document that will be included in server logs, profiling logs, and currentOp queries to help trace
// the operation.  The default value is nil, which means that no comment will be included in the logs.
func (o *AggregateOptions) Comment(v string) *AggregateOptions {
	o.o.SetComment(v)
	return o
}

// MaxTime specifies the maximum amount of time to allow the query to run.
func (o *AggregateOptions) MaxTime(v time.Duration) *AggregateOptions {
	o.o.SetMaxTime(v)
	return o
}

// MaxAwaitTime specifies the maximum amount of time for the server to wait on new documents to satisfy a tailable cursor
// query. This only applies to tailable cursor queries issued against a replica set. The default value is nil, which means that
// there is no maximum await time.
func (o *AggregateOptions) MaxAwaitTime(v time.Duration) *AggregateOptions {
	o.o.SetMaxAwaitTime(v)
	return o
}

// Hint specifies the index to use.
func (o *AggregateOptions) Hint(v any) *AggregateOptions {
	o.o.SetHint(v)
	return o
}

// Let specifies variables that can be accessed in the pipeline.
func (o *AggregateOptions) Let(v any) *AggregateOptions {
	o.o.SetLet(v)
	return o
}

// Custom specifies options to be added to aggregate expression.
func (o *AggregateOptions) Custom(v M) *AggregateOptions {
	o.o.SetCustom(bson.M(v))
	return o
}

// NewAggregateOptions creates a new AggregateOptions.
func NewAggregateOptions() *AggregateOptions {
	return &AggregateOptions{o: options.Aggregate()}
}
