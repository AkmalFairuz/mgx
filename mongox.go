package mgx

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoX is a wrapper around mongo.Client.
type MongoX struct {
	client *mongo.Client
	db     *mongo.Database
}

// New creates a new MongoX instance.
func New(uri, database string) (*MongoX, error) {
	// connect to mongodb with uri
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}
	err = client.Connect(nil)
	if err != nil {
		return nil, err
	}
	return &MongoX{
		client: client,
		db:     client.Database(database),
	}, nil
}

// Collection returns a collection with the given name.
func (m *MongoX) Collection(v string, ctx ...context.Context) *Collection {
	var c context.Context
	if len(ctx) > 0 {
		c = ctx[0]
	}
	return &Collection{
		collection: m.db.Collection(v),
		ctx:        c,
	}
}

// CreateCollection creates a new collection with the given name.
func (m *MongoX) CreateCollection(v string, ctx ...context.Context) *CreateCollectionResult {
	result := &CreateCollectionResult{}
	result.Err = m.db.CreateCollection(nil, v)
	result.Name = v
	result.Collection = m.Collection(v, ctx...)
	return result
}

// DropCollection drops the collection with the given name.
func (m *MongoX) DropCollection(v string) *DropCollectionResult {
	result := &DropCollectionResult{}
	result.Err = m.db.Collection(v).Drop(nil)
	result.Name = v
	return result
}

// Close closes the connection to the database.
func (m *MongoX) Close() error {
	m.db = nil
	return m.client.Disconnect(nil)
}
