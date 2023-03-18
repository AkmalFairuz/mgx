package mgx

import (
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

// DistinctOptions is a wrapper around mongo options.
type DistinctOptions struct {
	o *options.DistinctOptions
}

// opt returns the underlying mongo option.
func (o *DistinctOptions) opt() *options.DistinctOptions {
	return o.o
}

// Collation specifies a collation.
func (o *DistinctOptions) Collation(v *Collation) *DistinctOptions {
	o.o.SetCollation(v.Collation)
	return o
}

// Comment specifies a string or document that will be included in server logs, profiling logs, and currentOp queries to help trace
// the operation.  The default value is nil, which means that no comment will be included in the logs.
func (o *DistinctOptions) Comment(v string) *DistinctOptions {
	o.o.SetComment(v)
	return o
}

// MaxTime specifies the maximum amount of time to allow the query to run.
func (o *DistinctOptions) MaxTime(v time.Duration) *DistinctOptions {
	o.o.SetMaxTime(v)
	return o
}

// NewDistinctOptions creates a new DistinctOptions.
func NewDistinctOptions() *DistinctOptions {
	return &DistinctOptions{o: options.Distinct()}
}
