package mgx

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

// WatchOptions is a wrapper around mongo options.
type WatchOptions struct {
	o *options.ChangeStreamOptions
}

// opt returns the underlying mongo option.
func (o *WatchOptions) opt() *options.ChangeStreamOptions {
	return o.o
}

// BatchSize sets the number of documents to return in each batch.
func (o *WatchOptions) BatchSize(v int32) *WatchOptions {
	o.o.SetBatchSize(v)
	return o
}

// Collation specifies a collation.
func (o *WatchOptions) Collation(v Collation) *WatchOptions {
	o.o.SetCollation(*v.Collation)
	return o
}

// Comment adds a comment to the query.
func (o *WatchOptions) Comment(v string) *WatchOptions {
	o.o.SetComment(v)
	return o
}

// FullDocument specifies how a change stream should return the modified document.
type FullDocument string

const (
	// FullDocumentDefault does not include a document copy.
	FullDocumentDefault FullDocument = "default"
	// FullDocumentOff is the same as sending no value for fullDocumentBeforeChange.
	FullDocumentOff FullDocument = "off"
	// FullDocumentRequired is the same as WhenAvailable but raises a server-side error if the post-image is not available.
	FullDocumentRequired FullDocument = "required"
	// FullDocumentUpdateLookup includes a delta describing the changes to the document and a copy of the entire document that
	// was changed.
	FullDocumentUpdateLookup FullDocument = "updateLookup"
	// FullDocumentWhenAvailable includes a post-image of the the modified document for replace and update change events
	// if the post-image for this event is available.
	FullDocumentWhenAvailable FullDocument = "whenAvailable"
)

// FullDocument specifies whether to return the full document or just the changes to the document.
func (o *WatchOptions) FullDocument(v FullDocument) *WatchOptions {
	o.o.SetFullDocument(options.FullDocument(v))
	return o
}

// MaxAwaitTime specifies the maximum amount of time for the server to wait on new documents to satisfy a tailable cursor
// query. This only applies to tailable cursor queries issued against a replica set. The default value is nil, which means that
// there is no maximum await time.
func (o *WatchOptions) MaxAwaitTime(v time.Duration) *WatchOptions {
	o.o.SetMaxAwaitTime(v)
	return o
}

// FullDocumentBeforeChange specifies whether to return the full document before the change occurred.
func (o *WatchOptions) FullDocumentBeforeChange(v FullDocument) *WatchOptions {
	o.o.SetFullDocumentBeforeChange(options.FullDocument(v))
	return o
}

// ResumeAfter specifies the logical starting point for the new change stream.
func (o *WatchOptions) ResumeAfter(v any) *WatchOptions {
	o.o.SetResumeAfter(v)
	return o
}

// ShowExpandedEvents specifies whether to include the postImage field in update notifications.
func (o *WatchOptions) ShowExpandedEvents(v ...bool) *WatchOptions {
	x := true
	if len(v) > 0 {
		x = v[0]
	}
	o.o.SetShowExpandedEvents(x)
	return o
}

// StartAtOperationTime specifies a logical starting point for the new change stream.
func (o *WatchOptions) StartAtOperationTime(v *Timestamp) *WatchOptions {
	o.o.SetStartAtOperationTime(&primitive.Timestamp{I: v.I, T: v.T})
	return o
}

// StartAfter specifies the logical starting point for the new change stream.
func (o *WatchOptions) StartAfter(v any) *WatchOptions {
	o.o.SetStartAfter(v)
	return o
}

// Custom ...
func (o *WatchOptions) Custom(v M) *WatchOptions {
	o.o.SetCustom(bson.M(v))
	return o
}

// CustomPipeline ...
func (o *WatchOptions) CustomPipeline(v M) *WatchOptions {
	o.o.SetCustomPipeline(bson.M(v))
	return o
}

// NewWatchOptions creates a new WatchOptions.
func NewWatchOptions() *WatchOptions {
	return &WatchOptions{o: options.ChangeStream()}
}
