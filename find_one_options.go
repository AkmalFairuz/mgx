package mgx

import (
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

// GetOptions is a wrapper around mongo options.
type GetOptions struct {
	o *options.FindOneOptions
}

// opt returns the underlying mongo option.
func (o *GetOptions) opt() *options.FindOneOptions {
	return o.o
}

// AllowPartialResults specifies whether to allow partial results if some shards are
func (o *GetOptions) AllowPartialResults(v ...bool) *GetOptions {
	x := true
	if len(v) > 0 {
		x = v[0]
	}
	o.o.SetAllowPartialResults(x)
	return o
}

// BatchSize specifies the maximum number of documents to return in the batch.
func (o *GetOptions) BatchSize(v int32) *GetOptions {
	o.o.SetBatchSize(v)
	return o
}

// Collation specifies a collation.
func (o *GetOptions) Collation(v *Collation) *GetOptions {
	o.o.SetCollation(v.Collation)
	return o
}

// Comment specifies a string to help trace the operation through the database profiler, currentOp and logs.
func (o *GetOptions) Comment(v string) *GetOptions {
	o.o.SetComment(v)
	return o
}

// Hint specifies the index to use.
func (o *GetOptions) Hint(v any) *GetOptions {
	o.o.SetHint(v)
	return o
}

// Max specifies the exclusive upper bound for a specific index.
func (o *GetOptions) Max(v any) *GetOptions {
	o.o.SetMax(v)
	return o
}

// MaxTime specifies the maximum amount of time to allow the query to run.
func (o *GetOptions) MaxTime(v time.Duration) *GetOptions {
	o.o.SetMaxTime(v)
	return o
}

// Min specifies the inclusive lower bound for a specific index.
func (o *GetOptions) Min(v any) *GetOptions {
	o.o.SetMin(v)
	return o
}

// Projection specifies the fields to return.
func (o *GetOptions) Projection(v any) *GetOptions {
	o.o.SetProjection(v)
	return o
}

// ReturnKey specifies whether to only return the index keys in the resulting documents.
func (o *GetOptions) ReturnKey(v ...bool) *GetOptions {
	x := true
	if len(v) > 0 {
		x = v[0]
	}
	o.o.SetReturnKey(x)
	return o
}

// ShowRecordID specifies whether to add a record identifier for each document.
func (o *GetOptions) ShowRecordID(v ...bool) *GetOptions {
	x := true
	if len(v) > 0 {
		x = v[0]
	}
	o.o.SetShowRecordID(x)
	return o
}

// Skip specifies the number of documents to skip before returning.
func (o *GetOptions) Skip(v int64) *GetOptions {
	o.o.SetSkip(v)
	return o
}

// Sort specifies the order in which to return results.
func (o *GetOptions) Sort(v any) *GetOptions {
	o.o.SetSort(v)
	return o
}

// Snapshot specifies whether to prevent the cursor from returning a document more than once because of an intervening write operation.
func (o *GetOptions) Snapshot(v ...bool) *GetOptions {
	x := true
	if len(v) > 0 {
		x = v[0]
	}
	o.o.SetSnapshot(x)
	return o
}

// NoCursorTimeout specifies whether to allow the cursor to timeout.
func (o *GetOptions) NoCursorTimeout(v ...bool) *GetOptions {
	x := true
	if len(v) > 0 {
		x = v[0]
	}
	o.o.SetNoCursorTimeout(x)
	return o
}

// OplogReplay specifies whether to allow the cursor to timeout.
func (o *GetOptions) OplogReplay(v ...bool) *GetOptions {
	x := true
	if len(v) > 0 {
		x = v[0]
	}
	o.o.SetOplogReplay(x)
	return o
}

// NewGetOptions creates a new GetOptions.
func NewGetOptions() *GetOptions {
	return &GetOptions{o: options.FindOne()}
}
