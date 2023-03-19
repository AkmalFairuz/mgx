package mgx

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Collection is a wrapper around mongo.Collection.
// It provides more simple methods for common operations.
type Collection struct {
	collection *mongo.Collection
	ctx        context.Context
}

// Get returns the first document that matches the filter.
func (c *Collection) Get(filter, result any, opts ...*GetOptions) *SingleResult {
	var mongoOpts []*options.FindOneOptions
	for _, opt := range opts {
		mongoOpts = append(mongoOpts, opt.opt())
	}

	r := c.collection.FindOne(c.ctx, filter, mongoOpts...)
	err := r.Decode(result)
	if err != nil {
		return &SingleResult{Err: err}
	}
	return &SingleResult{Err: r.Err()}
}

// FindOne returns the first document that matches the filter.
func (c *Collection) FindOne(filter, result any, opt ...*GetOptions) *SingleResult {
	return c.Get(filter, result, opt...)
}

// Find returns all documents that match the filter.
func (c *Collection) Find(filter, result any, opts ...*FindOptions) *FindResult {
	var mongoOpts []*options.FindOptions
	for _, opt := range opts {
		mongoOpts = append(mongoOpts, opt.opt())
	}

	cursor, err := c.collection.Find(c.ctx, filter, mongoOpts...)
	if err != nil {
		return &FindResult{Err: err}
	}
	if err := cursor.All(nil, result); err != nil {
		return &FindResult{Err: err}
	}
	if err := cursor.Close(c.ctx); err != nil {
		return &FindResult{Err: fmt.Errorf("error closing cursor: %w", err)}
	}
	return &FindResult{Err: cursor.Err()}
}

// FindOneAndUpdate returns the first document that matches the filter.
func (c *Collection) FindOneAndUpdate(filter, update, result any, opts ...*FindOneAndUpdateOptions) *SingleResult {
	var mongoOpts []*options.FindOneAndUpdateOptions
	for _, opt := range opts {
		mongoOpts = append(mongoOpts, opt.opt())
	}

	r := c.collection.FindOneAndUpdate(c.ctx, filter, update, mongoOpts...)
	err := r.Decode(result)
	if err != nil {
		return &SingleResult{Err: err}
	}
	return &SingleResult{Err: r.Err()}
}

// FindOneAndDelete returns the first document that matches the filter.
func (c *Collection) FindOneAndDelete(filter, result any, opt ...*FindOneAndDeleteOptions) *SingleResult {
	var mongoOpts []*options.FindOneAndDeleteOptions
	for _, opt := range opt {
		mongoOpts = append(mongoOpts, opt.opt())
	}

	r := c.collection.FindOneAndDelete(c.ctx, filter, mongoOpts...)
	err := r.Decode(result)
	if err != nil {
		return &SingleResult{Err: err}
	}
	return &SingleResult{Err: r.Err()}
}

// FindOneAndReplace returns the first document that matches the filter.
func (c *Collection) FindOneAndReplace(filter, replacement, result any, opt ...*FindOneAndReplaceOptions) *SingleResult {
	var mongoOpts []*options.FindOneAndReplaceOptions
	for _, opt := range opt {
		mongoOpts = append(mongoOpts, opt.opt())
	}

	r := c.collection.FindOneAndReplace(c.ctx, filter, replacement, mongoOpts...)
	err := r.Decode(result)
	if err != nil {
		return &SingleResult{Err: err}
	}
	return &SingleResult{Err: r.Err()}
}

// Insert is an alias for InsertOne.
func (c *Collection) Insert(doc any, opts ...*InsertOneOptions) *InsertResult {
	return c.InsertOne(doc, opts...)
}

// InsertOne inserts a single document into the collection.
func (c *Collection) InsertOne(doc any, opts ...*InsertOneOptions) *InsertResult {
	var mongoOpts []*options.InsertOneOptions
	for _, opt := range opts {
		mongoOpts = append(mongoOpts, opt.opt())
	}

	result, err := c.collection.InsertOne(c.ctx, doc, mongoOpts...)
	if err != nil {
		return &InsertResult{Err: err}
	}
	return &InsertResult{InsertedID: result.InsertedID}
}

// InsertMany inserts multiple documents into the collection.
func (c *Collection) InsertMany(docs []any, opts ...*InsertManyOptions) *InsertManyResult {
	var mongoOpts []*options.InsertManyOptions
	for _, opt := range opts {
		mongoOpts = append(mongoOpts, opt.opt())
	}

	result, err := c.collection.InsertMany(c.ctx, docs, mongoOpts...)
	if err != nil {
		return &InsertManyResult{Err: err}
	}
	return &InsertManyResult{InsertedIDs: result.InsertedIDs}
}

// Update is an alias for UpdateOne.
func (c *Collection) Update(filter, update any, opts ...*UpdateOptions) *UpdateResult {
	return c.UpdateOne(filter, update, opts...)
}

// UpdateOne updates a single document in the collection.
func (c *Collection) UpdateOne(filter, update any, opts ...*UpdateOptions) *UpdateResult {
	var mongoOpts []*options.UpdateOptions
	for _, opt := range opts {
		mongoOpts = append(mongoOpts, opt.opt())
	}

	result, err := c.collection.UpdateOne(c.ctx, filter, update, mongoOpts...)
	if err != nil {
		return &UpdateResult{Err: err}
	}
	return &UpdateResult{
		MatchedCount:  result.MatchedCount,
		ModifiedCount: result.ModifiedCount,
		UpsertedCount: result.UpsertedCount,
		UpsertedID:    result.UpsertedID,
	}
}

// UpdateMany updates multiple documents in the collection.
func (c *Collection) UpdateMany(filter, update any, opts ...*options.UpdateOptions) *UpdateResult {
	var mongoOpts []*options.UpdateOptions
	for _, opt := range opts {
		mongoOpts = append(mongoOpts, opt)
	}

	result, err := c.collection.UpdateMany(c.ctx, filter, update, mongoOpts...)
	if err != nil {
		return &UpdateResult{Err: err}
	}
	return &UpdateResult{
		MatchedCount:  result.MatchedCount,
		ModifiedCount: result.ModifiedCount,
		UpsertedCount: result.UpsertedCount,
		UpsertedID:    result.UpsertedID,
	}
}

// Delete is an alias for DeleteOne.
func (c *Collection) Delete(filter any, opts ...*DeleteOptions) *DeleteResult {
	return c.DeleteOne(filter, opts...)
}

// DeleteOne deletes a single document from the collection.
func (c *Collection) DeleteOne(filter any, opts ...*DeleteOptions) *DeleteResult {
	var mongoOpts []*options.DeleteOptions
	for _, opt := range opts {
		mongoOpts = append(mongoOpts, opt.opt())
	}

	result, err := c.collection.DeleteOne(c.ctx, filter, mongoOpts...)
	if err != nil {
		return &DeleteResult{Err: err}
	}
	return &DeleteResult{DeletedCount: result.DeletedCount}
}

// DeleteMany deletes multiple documents from the collection.
func (c *Collection) DeleteMany(filter any, opts ...*options.DeleteOptions) *DeleteResult {
	var mongoOpts []*options.DeleteOptions
	for _, opt := range opts {
		mongoOpts = append(mongoOpts, opt)
	}

	result, err := c.collection.DeleteMany(c.ctx, filter, mongoOpts...)
	if err != nil {
		return &DeleteResult{Err: err}
	}
	return &DeleteResult{DeletedCount: result.DeletedCount}
}

// ReplaceOne replaces a single document in the collection.
func (c *Collection) ReplaceOne(filter, replacement any, opts ...*ReplaceOptions) *UpdateResult {
	var mongoOpts []*options.ReplaceOptions
	for _, opt := range opts {
		mongoOpts = append(mongoOpts, opt.opt())
	}

	result, err := c.collection.ReplaceOne(c.ctx, filter, replacement, mongoOpts...)
	if err != nil {
		return &UpdateResult{Err: err}
	}
	return &UpdateResult{
		MatchedCount:  result.MatchedCount,
		ModifiedCount: result.ModifiedCount,
		UpsertedCount: result.UpsertedCount,
		UpsertedID:    result.UpsertedID,
	}
}

// Count returns the number of documents in the collection.
func (c *Collection) Count(filter any, opts ...*CountOptions) *CountResult {
	var mongoOpts []*options.CountOptions
	for _, opt := range opts {
		mongoOpts = append(mongoOpts, opt.opt())
	}

	result, err := c.collection.CountDocuments(c.ctx, filter, mongoOpts...)
	if err != nil {
		return &CountResult{Err: err}
	}
	return &CountResult{Count: result}
}

// Distinct returns a list of distinct values for the given key across a single collection.
func (c *Collection) Distinct(key string, filter any, opts ...*DistinctOptions) *DistinctResult {
	var mongoOpts []*options.DistinctOptions
	for _, opt := range opts {
		mongoOpts = append(mongoOpts, opt.opt())
	}

	result, err := c.collection.Distinct(c.ctx, key, filter, mongoOpts...)
	if err != nil {
		return &DistinctResult{Err: err}
	}
	return &DistinctResult{Result: result}
}

// Aggregate performs an aggregation framework pipeline against the collection.
func (c *Collection) Aggregate(result, pipeline []any, opts ...*AggregateOptions) *AggregateResult {
	var mongoOpts []*options.AggregateOptions
	for _, opt := range opts {
		mongoOpts = append(mongoOpts, opt.opt())
	}

	cursor, err := c.collection.Aggregate(c.ctx, pipeline, mongoOpts...)
	if err := cursor.All(c.ctx, result); err != nil {
		return &AggregateResult{Err: err}
	}
	if err := cursor.Close(c.ctx); err != nil {
		return &AggregateResult{Err: fmt.Errorf("error closing cursor: %w", err)}
	}
	return &AggregateResult{Err: err}
}

// Watch creates a change stream cursor for a collection.
func (c *Collection) Watch(pipeline []any, opts ...*WatchOptions) *WatchResult {
	var mongoOpts []*options.ChangeStreamOptions
	for _, opt := range opts {
		mongoOpts = append(mongoOpts, opt.opt())
	}

	changeStream, err := c.collection.Watch(c.ctx, pipeline, mongoOpts...)
	if err != nil {
		return &WatchResult{Err: err}
	}
	return &WatchResult{changeStream: changeStream}
}

// CreateIndex creates a single index on the collection.
func (c *Collection) CreateIndex(model IndexModel, opts ...*CreateIndexesOptions) *CreateIndexResult {
	var mongoOpts []*options.CreateIndexesOptions
	for _, opt := range opts {
		mongoOpts = append(mongoOpts, opt.opt())
	}

	result, err := c.collection.Indexes().CreateOne(c.ctx, model.mongoIndex(), mongoOpts...)
	if err != nil {
		return &CreateIndexResult{Err: err}
	}
	return &CreateIndexResult{Name: result}
}

// CreateIndexes creates multiple indexes on the collection.
func (c *Collection) CreateIndexes(models []IndexModel, opts ...*CreateIndexesOptions) *CreateIndexesResult {
	var mongoOpts []*options.CreateIndexesOptions
	for _, opt := range opts {
		mongoOpts = append(mongoOpts, opt.opt())
	}

	var mongoIndexModels []mongo.IndexModel
	for _, model := range models {
		mongoIndexModels = append(mongoIndexModels, model.mongoIndex())
	}

	result, err := c.collection.Indexes().CreateMany(c.ctx, mongoIndexModels, mongoOpts...)
	if err != nil {
		return &CreateIndexesResult{Err: err}
	}
	return &CreateIndexesResult{Names: result}
}

// ListIndexes lists the indexes on the collection.
func (c *Collection) ListIndexes(result []IndexModel, opts ...*ListIndexesOptions) *ListIndexesResult {
	var mongoOpts []*options.ListIndexesOptions
	for _, opt := range opts {
		mongoOpts = append(mongoOpts, opt.opt())
	}

	cursor, err := c.collection.Indexes().List(c.ctx, mongoOpts...)
	if err := cursor.All(c.ctx, result); err != nil {
		return &ListIndexesResult{Err: err}
	}
	if err := cursor.Close(c.ctx); err != nil {
		return &ListIndexesResult{Err: fmt.Errorf("error closing cursor: %w", err)}
	}
	return &ListIndexesResult{Err: err}
}

// DropIndex drops a single index from the collection.
func (c *Collection) DropIndex(name string, opts ...*DropIndexesOptions) *DropIndexResult {
	var mongoOpts []*options.DropIndexesOptions
	for _, opt := range opts {
		mongoOpts = append(mongoOpts, opt.opt())
	}

	result, err := c.collection.Indexes().DropOne(c.ctx, name, mongoOpts...)
	if err != nil {
		return &DropIndexResult{Err: err}
	}
	return &DropIndexResult{Name: string(result)}
}

// DropIndexes drops all indexes from the collection.
func (c *Collection) DropIndexes(opts ...*DropIndexesOptions) *DropIndexesResult {
	var mongoOpts []*options.DropIndexesOptions
	for _, opt := range opts {
		mongoOpts = append(mongoOpts, opt.opt())
	}

	result, err := c.collection.Indexes().DropAll(c.ctx, mongoOpts...)
	if err != nil {
		return &DropIndexesResult{Err: err}
	}
	var names []string
	for _, name := range result {
		names = append(names, string(name))
	}
	return &DropIndexesResult{
		Names: names,
	}
}

// WithContext sets the context for the collection.
// This context will be used for all operations on the collection.
func (c *Collection) WithContext(ctx context.Context) *Collection {
	c.ctx = ctx
	return c
}
