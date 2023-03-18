package mgx

import (
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

// FindOneAndDeleteOptions is a wrapper around mongo options.
type FindOneAndDeleteOptions struct {
	o *options.FindOneAndDeleteOptions
}

// opt returns the underlying mongo option.
func (o *FindOneAndDeleteOptions) opt() *options.FindOneAndDeleteOptions {
	return o.o
}

func (o *FindOneAndDeleteOptions) Collation(v Collation) *FindOneAndDeleteOptions {
	o.o.SetCollation(v.Collation)
	return o
}

func (o *FindOneAndDeleteOptions) Comment(v any) *FindOneAndDeleteOptions {
	o.o.SetComment(v)
	return o
}

func (o *FindOneAndDeleteOptions) MaxTime(v time.Duration) *FindOneAndDeleteOptions {
	o.o.SetMaxTime(v)
	return o
}

func (o *FindOneAndDeleteOptions) Projection(v any) *FindOneAndDeleteOptions {
	o.o.SetProjection(v)
	return o
}

func (o *FindOneAndDeleteOptions) Sort(v any) *FindOneAndDeleteOptions {
	o.o.SetSort(v)
	return o
}

func (o *FindOneAndDeleteOptions) Hint(v any) *FindOneAndDeleteOptions {
	o.o.SetHint(v)
	return o
}

func (o *FindOneAndDeleteOptions) Let(v any) *FindOneAndDeleteOptions {
	o.o.SetLet(v)
	return o
}

// NewFindOneAndDeleteOptions creates a new FindOneAndDeleteOptions.
func NewFindOneAndDeleteOptions() *FindOneAndDeleteOptions {
	return &FindOneAndDeleteOptions{o: options.FindOneAndDelete()}
}
