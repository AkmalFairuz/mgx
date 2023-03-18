package mgx

import (
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

// FindOptions is a wrapper around mongo options.
type FindOptions struct {
	o *options.FindOptions
}

// opt returns the underlying mongo option.
func (o *FindOptions) opt() *options.FindOptions {
	return o.o
}

// Limit sets the limit for the query.
func (o *FindOptions) Limit(v int64) *FindOptions {
	o.o.SetLimit(v)
	return o
}

// Skip sets the number of documents to skip before returning.
func (o *FindOptions) Skip(v int64) *FindOptions {
	o.o.SetSkip(v)
	return o
}

// Sort sets the order in which to return matching documents.
func (o *FindOptions) Sort(v any) *FindOptions {
	o.o.SetSort(v)
	return o
}

// Projection sets the fields to return for all matching documents.
func (o *FindOptions) Projection(v any) *FindOptions {
	o.o.SetProjection(v)
	return o
}

// BatchSize sets the number of documents to return in each batch.
func (o *FindOptions) BatchSize(v int32) *FindOptions {
	o.o.SetBatchSize(v)
	return o
}

// MaxTimeMS sets the maximum amount of time to allow the query to run.
func (o *FindOptions) MaxTimeMS(v time.Duration) *FindOptions {
	o.o.SetMaxTime(v)
	return o
}

// Comment adds a comment to the query.
func (o *FindOptions) Comment(v string) *FindOptions {
	o.o.SetComment(v)
	return o
}

// Hint sets the index to use.
func (o *FindOptions) Hint(v any) *FindOptions {
	o.o.SetHint(v)
	return o
}

// MaxAwaitTime sets the maximum amount of time for the server to wait on new documents to satisfy a tailable cursor
// query. This only applies when CursorType is set to CursorTailableAwait.
func (o *FindOptions) MaxAwaitTime(v time.Duration) *FindOptions {
	o.o.SetMaxAwaitTime(v)
	return o
}

// AllowDiskUse enables writing to temporary files.
func (o *FindOptions) AllowDiskUse(v ...bool) *FindOptions {
	x := true
	if len(v) > 0 {
		x = v[0]
	}
	o.o.SetAllowDiskUse(x)
	return o
}

type Collation struct {
	*options.Collation
}

// Collation specifies a collation.
func (o *FindOptions) Collation(v *Collation) *FindOptions {
	o.o.SetCollation(v.Collation)
	return o
}

type CursorType int8

const (
	// CursorNonTailable specifies that a cursor should close after retrieving the last data.
	CursorNonTailable CursorType = iota
	// CursorTailable specifies that a cursor should not close when the last data is retrieved and can be resumed later.
	CursorTailable
	// CursorTailableAwait specifies that a cursor should not close when the last data is retrieved and
	// that it should block for a certain amount of time for new data before returning no data.
	CursorTailableAwait
)

// CursorType sets the type of cursor to use.
func (o *FindOptions) CursorType(v CursorType) *FindOptions {
	o.o.SetCursorType(options.CursorType(v))
	return o
}

// NoCursorTimeout specifies that a cursor should not time out after a period of inactivity.
func (o *FindOptions) NoCursorTimeout(v ...bool) *FindOptions {
	x := true
	if len(v) > 0 {
		x = v[0]
	}
	o.o.SetNoCursorTimeout(x)
	return o
}

// AllowPartialResults specifies that a cursor should not error when some shards are unavailable.
func (o *FindOptions) AllowPartialResults(v ...bool) *FindOptions {
	x := true
	if len(v) > 0 {
		x = v[0]
	}
	o.o.SetAllowPartialResults(x)
	return o
}

// Max specifies the exclusive upper bound for a specific index.
func (o *FindOptions) Max(v any) *FindOptions {
	o.o.SetMax(v)
	return o
}

// Min specifies the inclusive lower bound for a specific index.
func (o *FindOptions) Min(v any) *FindOptions {
	o.o.SetMin(v)
	return o
}

// ReturnKey specifies that a cursor should only return the index field or fields for the results.
func (o *FindOptions) ReturnKey(v ...bool) *FindOptions {
	x := true
	if len(v) > 0 {
		x = v[0]
	}
	o.o.SetReturnKey(x)
	return o
}

// ShowRecordID specifies that a cursor should return the record identifier for each document.
func (o *FindOptions) ShowRecordID(v ...bool) *FindOptions {
	x := true
	if len(v) > 0 {
		x = v[0]
	}
	o.o.SetShowRecordID(x)
	return o
}

// Let specifies variables to use in the query.
func (o *FindOptions) Let(v any) *FindOptions {
	o.o.SetLet(v)
	return o
}

// MaxTime sets the maximum amount of time to allow the query to run.
func (o *FindOptions) MaxTime(v time.Duration) *FindOptions {
	o.o.SetMaxTime(v)
	return o
}

// OpLogReplay specifies that a cursor should use the oplog for the query.
func (o *FindOptions) OpLogReplay(v ...bool) *FindOptions {
	x := true
	if len(v) > 0 {
		x = v[0]
	}
	o.o.SetOplogReplay(x)
	return o
}

// Snapshot specifies that a cursor should use a snapshot for the query.
func (o *FindOptions) Snapshot(v ...bool) *FindOptions {
	x := true
	if len(v) > 0 {
		x = v[0]
	}
	o.o.SetSnapshot(x)
	return o
}

// NewFindOptions creates a new FindOptions.
func NewFindOptions() *FindOptions {
	return &FindOptions{options.Find()}
}
