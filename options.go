package mgx

import "go.mongodb.org/mongo-driver/mongo/options"

// Options is a wrapper around mongo options.
type Options[T any] interface {
	// opt returns the underlying mongo option.
	opt() T
}

// FindOptions is a wrapper around mongo options.
type FindOptions struct {
	*options.FindOptions
}

// opt returns the underlying mongo option.
func (o *FindOptions) opt() *options.FindOptions {
	return o.FindOptions
}

// GetOptions is a wrapper around mongo options.
type GetOptions struct {
	*options.FindOneOptions
}

// opt returns the underlying mongo option.
func (o *GetOptions) opt() *options.FindOneOptions {
	return o.FindOneOptions
}

// FindOneAndUpdateOptions is a wrapper around mongo options.
type FindOneAndUpdateOptions struct {
	*options.FindOneAndUpdateOptions
}

// opt returns the underlying mongo option.
func (o *FindOneAndUpdateOptions) opt() *options.FindOneAndUpdateOptions {
	return o.FindOneAndUpdateOptions
}

// FindOneAndDeleteOptions is a wrapper around mongo options.
type FindOneAndDeleteOptions struct {
	*options.FindOneAndDeleteOptions
}

// opt returns the underlying mongo option.
func (o *FindOneAndDeleteOptions) opt() *options.FindOneAndDeleteOptions {
	return o.FindOneAndDeleteOptions
}

// FindOneAndReplaceOptions is a wrapper around mongo options.
type FindOneAndReplaceOptions struct {
	*options.FindOneAndReplaceOptions
}

// opt returns the underlying mongo option.
func (o *FindOneAndReplaceOptions) opt() *options.FindOneAndReplaceOptions {
	return o.FindOneAndReplaceOptions
}

// InsertOneOptions is a wrapper around mongo options.
type InsertOneOptions struct {
	*options.InsertOneOptions
}

// opt returns the underlying mongo option.
func (o *InsertOneOptions) opt() *options.InsertOneOptions {
	return o.InsertOneOptions
}

// InsertManyOptions is a wrapper around mongo options.
type InsertManyOptions struct {
	*options.InsertManyOptions
}

// opt returns the underlying mongo option.
func (o *InsertManyOptions) opt() *options.InsertManyOptions {
	return o.InsertManyOptions
}

// UpdateOptions is a wrapper around mongo options.
type UpdateOptions struct {
	*options.UpdateOptions
}

// opt returns the underlying mongo option.
func (o *UpdateOptions) opt() *options.UpdateOptions {
	return o.UpdateOptions
}

// DeleteOptions is a wrapper around mongo options.
type DeleteOptions struct {
	*options.DeleteOptions
}

// opt returns the underlying mongo option.
func (o *DeleteOptions) opt() *options.DeleteOptions {
	return o.DeleteOptions
}

// ReplaceOptions is a wrapper around mongo options.
type ReplaceOptions struct {
	*options.ReplaceOptions
}

// opt returns the underlying mongo option.
func (o *ReplaceOptions) opt() *options.ReplaceOptions {
	return o.ReplaceOptions
}

// CountOptions is a wrapper around mongo options.
type CountOptions struct {
	*options.CountOptions
}

// opt returns the underlying mongo option.
func (o *CountOptions) opt() *options.CountOptions {
	return o.CountOptions
}

// DistinctOptions is a wrapper around mongo options.
type DistinctOptions struct {
	*options.DistinctOptions
}

// opt returns the underlying mongo option.
func (o *DistinctOptions) opt() *options.DistinctOptions {
	return o.DistinctOptions
}

// AggregateOptions is a wrapper around mongo options.
type AggregateOptions struct {
	*options.AggregateOptions
}

// opt returns the underlying mongo option.
func (o *AggregateOptions) opt() *options.AggregateOptions {
	return o.AggregateOptions
}

// WatchOptions is a wrapper around mongo options.
type WatchOptions struct {
	*options.ChangeStreamOptions
}

// opt returns the underlying mongo option.
func (o *WatchOptions) opt() *options.ChangeStreamOptions {
	return o.ChangeStreamOptions
}

// IndexOptions is a wrapper around mongo options.
type IndexOptions struct {
	*options.IndexOptions
}

// opt returns the underlying mongo option.
func (o *IndexOptions) opt() *options.IndexOptions {
	return o.IndexOptions
}

// CreateIndexesOptions is a wrapper around mongo options.
type CreateIndexesOptions struct {
	*options.CreateIndexesOptions
}

// opt returns the underlying mongo option.
func (o *CreateIndexesOptions) opt() *options.CreateIndexesOptions {
	return o.CreateIndexesOptions
}

// ListIndexesOptions is a wrapper around mongo options.
type ListIndexesOptions struct {
	*options.ListIndexesOptions
}

// opt returns the underlying mongo option.
func (o *ListIndexesOptions) opt() *options.ListIndexesOptions {
	return o.ListIndexesOptions
}

// DropIndexesOptions is a wrapper around mongo options.
type DropIndexesOptions struct {
	*options.DropIndexesOptions
}

// opt returns the underlying mongo option.
func (o *DropIndexesOptions) opt() *options.DropIndexesOptions {
	return o.DropIndexesOptions
}
