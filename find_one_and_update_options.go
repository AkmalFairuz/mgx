package mgx

import (
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

// FindOneAndUpdateOptions is a wrapper around mongo options.
type FindOneAndUpdateOptions struct {
	o *options.FindOneAndUpdateOptions
}

// opt returns the underlying mongo option.
func (o *FindOneAndUpdateOptions) opt() *options.FindOneAndUpdateOptions {
	return o.o
}

type ArrayFilters struct {
	options.ArrayFilters
}

func (o *FindOneAndUpdateOptions) ArrayFilters(v ArrayFilters) *FindOneAndUpdateOptions {
	o.o.SetArrayFilters(v.ArrayFilters)
	return o
}

func (o *FindOneAndUpdateOptions) BypassDocumentValidation(v ...bool) *FindOneAndUpdateOptions {
	x := false
	if len(v) > 0 {
		x = v[0]
	}
	o.o.SetBypassDocumentValidation(x)
	return o
}

func (o *FindOneAndUpdateOptions) Collation(v Collation) *FindOneAndUpdateOptions {
	o.o.SetCollation(v.Collation)
	return o
}

func (o *FindOneAndUpdateOptions) Comment(v any) *FindOneAndUpdateOptions {
	o.o.SetComment(v)
	return o
}

func (o *FindOneAndUpdateOptions) MaxTime(v time.Duration) *FindOneAndUpdateOptions {
	o.o.SetMaxTime(v)
	return o
}

func (o *FindOneAndUpdateOptions) Projection(v any) *FindOneAndUpdateOptions {
	o.o.SetProjection(v)
	return o
}

type ReturnDocument int32

func (o *FindOneAndUpdateOptions) ReturnDocument(v ReturnDocument) *FindOneAndUpdateOptions {
	o.o.SetReturnDocument(options.ReturnDocument(v))
	return o
}

func (o *FindOneAndUpdateOptions) Sort(v any) *FindOneAndUpdateOptions {
	o.o.SetSort(v)
	return o
}

func (o *FindOneAndUpdateOptions) Upsert(v ...bool) *FindOneAndUpdateOptions {
	x := false
	if len(v) > 0 {
		x = v[0]
	}
	o.o.SetUpsert(x)
	return o
}

func (o *FindOneAndUpdateOptions) Hint(v any) *FindOneAndUpdateOptions {
	o.o.SetHint(v)
	return o
}

func (o *FindOneAndUpdateOptions) Let(v any) *FindOneAndUpdateOptions {
	o.o.SetLet(v)
	return o
}

// NewFindOneAndUpdateOptions creates a new FindOneAndUpdateOptions.
func NewFindOneAndUpdateOptions() *FindOneAndUpdateOptions {
	return &FindOneAndUpdateOptions{o: options.FindOneAndUpdate()}
}
