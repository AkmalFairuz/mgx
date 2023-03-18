package mgx

import (
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

// DropIndexesOptions is a wrapper around mongo options.
type DropIndexesOptions struct {
	o *options.DropIndexesOptions
}

// opt returns the underlying mongo option.
func (o *DropIndexesOptions) opt() *options.DropIndexesOptions {
	return o.o
}

// MaxTime specifies the maximum amount of time to allow the query to run.
func (o *DropIndexesOptions) MaxTime(v time.Duration) *DropIndexesOptions {
	o.o.SetMaxTime(v)
	return o
}

// NewDropIndexesOptions creates a new DropIndexesOptions.
func NewDropIndexesOptions() *DropIndexesOptions {
	return &DropIndexesOptions{o: options.DropIndexes()}
}
