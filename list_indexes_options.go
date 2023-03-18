package mgx

import (
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

// ListIndexesOptions is a wrapper around mongo options.
type ListIndexesOptions struct {
	o *options.ListIndexesOptions
}

// opt returns the underlying mongo option.
func (o *ListIndexesOptions) opt() *options.ListIndexesOptions {
	return o.o
}

// BatchSize specifies the maximum number of documents to return in each batch of the returned cursor.
func (o *ListIndexesOptions) BatchSize(v int32) *ListIndexesOptions {
	o.o.SetBatchSize(v)
	return o
}

// MaxTime specifies the maximum amount of time to allow the query to run.
func (o *ListIndexesOptions) MaxTime(v time.Duration) *ListIndexesOptions {
	o.o.SetMaxTime(v)
	return o
}

// NewListIndexesOptions creates a new ListIndexesOptions.
func NewListIndexesOptions() *ListIndexesOptions {
	return &ListIndexesOptions{o: options.ListIndexes()}
}
