package mgx

import (
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

// FindOneAndReplaceOptions is a wrapper around mongo options.
type FindOneAndReplaceOptions struct {
	o *options.FindOneAndReplaceOptions
}

// opt returns the underlying mongo option.
func (o *FindOneAndReplaceOptions) opt() *options.FindOneAndReplaceOptions {
	return o.o
}

func (o *FindOneAndReplaceOptions) Collation(v Collation) *FindOneAndReplaceOptions {
	o.o.SetCollation(v.Collation)
	return o
}

func (o *FindOneAndReplaceOptions) Comment(v any) *FindOneAndReplaceOptions {
	o.o.SetComment(v)
	return o
}

func (o *FindOneAndReplaceOptions) MaxTime(v time.Duration) *FindOneAndReplaceOptions {
	o.o.SetMaxTime(v)
	return o
}

func (o *FindOneAndReplaceOptions) Projection(v any) *FindOneAndReplaceOptions {
	o.o.SetProjection(v)
	return o
}

func (o *FindOneAndReplaceOptions) Sort(v any) *FindOneAndReplaceOptions {
	o.o.SetSort(v)
	return o
}

func (o *FindOneAndReplaceOptions) Hint(v any) *FindOneAndReplaceOptions {
	o.o.SetHint(v)
	return o
}

func (o *FindOneAndReplaceOptions) Let(v any) *FindOneAndReplaceOptions {
	o.o.SetLet(v)
	return o
}

func (o *FindOneAndReplaceOptions) Upsert(v ...bool) *FindOneAndReplaceOptions {
	x := false
	if len(v) > 0 {
		x = v[0]
	}
	o.o.SetUpsert(x)
	return o
}

func (o *FindOneAndReplaceOptions) BypassDocumentValidation(v ...bool) *FindOneAndReplaceOptions {
	x := false
	if len(v) > 0 {
		x = v[0]
	}
	o.o.SetBypassDocumentValidation(x)
	return o
}

func (o *FindOneAndReplaceOptions) ReturnDocument(v ReturnDocument) *FindOneAndReplaceOptions {
	o.o.SetReturnDocument(options.ReturnDocument(v))
	return o
}

// NewFindOneAndReplaceOptions creates a new FindOneAndReplaceOptions.
func NewFindOneAndReplaceOptions() *FindOneAndReplaceOptions {
	return &FindOneAndReplaceOptions{o: options.FindOneAndReplace()}
}
