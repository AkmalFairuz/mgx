package mgx

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
)

type InsertResult struct {
	// Err is the error that occurred during the operation, if any.
	Err error
	// InsertedID is the _id field of the inserted document, or nil if no document was inserted.
	InsertedID any
}

type InsertManyResult struct {
	// Err is the error that occurred during the operation, if any.
	Err error
	// InsertedIDs is the _id fields of the inserted documents, or nil if no documents were inserted.
	InsertedID []any
}

type UpdateResult struct {
	// Err is the error that occurred during the operation, if any.
	Err error
	// MatchedCount is the number of documents matched by the filter.
	MatchedCount int64
	// ModifiedCount is the number of documents modified by the operation.
	ModifiedCount int64
	// UpsertedCount is the number of documents upserted by the operation.
	UpsertedCount int64
	// UpsertedID is the _id of the upserted document, or nil if no document was upserted.
	UpsertedID any
}

type DeleteResult struct {
	// Err is the error that occurred during the operation, if any.
	Err error
	// DeletedCount is the number of documents deleted by the operation.
	DeletedCount int64
}

type SingleResult struct {
	// Err is the error that occurred during the operation, if any.
	Err error
}

type FindResult struct {
	// Err is the error that occurred during the operation, if any.
	Err error
}

type CountResult struct {
	// Err is the error that occurred during the operation, if any.
	Err error
	// Count is the number of documents returned by the operation.
	Count int64
}

type DistinctResult struct {
	// Err is the error that occurred during the operation, if any.
	Err error
	// Result is the list of distinct values returned by the operation.
	Result []any
}

type AggregateResult struct {
	// Err is the error that occurred during the operation, if any.
	Err error
}

type WatchResult struct {
	// Err is the error that occurred during the operation, if any.
	Err error

	changeStream *mongo.ChangeStream
	ctx          context.Context
	end          bool
}

func (r *WatchResult) Get(v any) error {
	if r.end {
		return fmt.Errorf("watch result is empty")
	}
	if err := r.changeStream.Decode(v); err != nil {
		return err
	}
	if !r.changeStream.TryNext(r.ctx) {
		r.end = true
	}
	return nil
}

type CreateIndexResult struct {
	// Err is the error that occurred during the operation, if any.
	Err error
	// Name is the name of the created index.
	Name string
}

type CreateIndexesResult struct {
	// Err is the error that occurred during the operation, if any.
	Err error
	// Names is the names of the created indexes.
	Names []string
}

type DropIndexResult struct {
	// Err is the error that occurred during the operation, if any.
	Err error
	// Name is the name of the dropped index.
	Name string
}

type DropIndexesResult struct {
	// Err is the error that occurred during the operation, if any.
	Err error
	// Names is the names of the dropped indexes.
	Names []string
}

type ListIndexesResult struct {
	// Err is the error that occurred during the operation, if any.
	Err error
}

type CreateCollectionResult struct {
	// Err is the error that occurred during the operation, if any.
	Err error
	// Name is the name of the created collection.
	Name string
	// Collection is the created collection.
	Collection *Collection
}

type DropCollectionResult struct {
	// Err is the error that occurred during the operation, if any.
	Err error
	// Name is the name of the dropped collection.
	Name string
}

type ListCollectionsResult struct {
	// Err is the error that occurred during the operation, if any.
	Err error

	// Names is the names of the collections returned by the operation.
	Names []string
}
