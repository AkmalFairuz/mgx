package mgx

import "go.mongodb.org/mongo-driver/mongo/options"

type ListCollectionsOptions struct {
	o *options.ListCollectionsOptions
}

// NameOnly specifies whether to limit the documents returned to only contain the name of the collection.
func (o *ListCollectionsOptions) NameOnly(v bool) *ListCollectionsOptions {
	o.o.SetNameOnly(v)
	return o
}

// BatchSize specifies the maximum number of documents to be included in each batch returned by the server.
func (o *ListCollectionsOptions) BatchSize(v int32) *ListCollectionsOptions {
	o.o.SetBatchSize(v)
	return o
}

// AuthorizedCollections specifies whether to limit the documents returned to only contain collections the user is authorized to use.
func (o *ListCollectionsOptions) AuthorizedCollections(v bool) *ListCollectionsOptions {
	o.o.SetAuthorizedCollections(v)
	return o
}

func NewListCollectionsOptions() *ListCollectionsOptions {
	return &ListCollectionsOptions{
		o: options.ListCollections(),
	}
}
