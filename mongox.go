package mgx

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Database ...
type Database struct {
	client *mongo.Client
	db     *mongo.Database
	ctx    context.Context
}

// New creates a new MongoX instance.
func New(uri, database string, ctx ...context.Context) (*Database, error) {
	// connect to mongodb with uri

	c := context.Background()
	if len(ctx) > 0 {
		c = ctx[0]
	}

	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}
	err = client.Connect(nil)
	if err != nil {
		return nil, err
	}
	return &Database{
		client: client,
		ctx:    c,
		db:     client.Database(database),
	}, nil
}

// Collection returns a collection with the given name.
func (db *Database) Collection(v string, ctx ...context.Context) *Collection {
	c := context.TODO()
	if len(ctx) > 0 {
		c = ctx[0]
	}
	return &Collection{
		collection: db.db.Collection(v),
		ctx:        c,
	}
}

// CreateCollection creates a new collection with the given name.
func (db *Database) CreateCollection(v string, ctx ...context.Context) *CreateCollectionResult {
	result := &CreateCollectionResult{}
	result.Err = db.db.CreateCollection(db.ctx, v)
	result.Name = v
	result.Collection = db.Collection(v, ctx...)
	return result
}

// DropCollection drops the collection with the given name.
func (db *Database) DropCollection(v string) *DropCollectionResult {
	result := &DropCollectionResult{}
	result.Err = db.db.Collection(v).Drop(db.ctx)
	result.Name = v
	return result
}

// ListCollections lists all collections in the database.
func (db *Database) ListCollections(result, filter any, opts ...*ListCollectionsOptions) *ListCollectionsResult {
	var mongoOpts []*options.ListCollectionsOptions
	for _, opt := range opts {
		mongoOpts = append(mongoOpts, opt.o)
	}
	r, err := db.db.ListCollections(db.ctx, filter, mongoOpts...)
	if err != nil {
		return &ListCollectionsResult{Err: err}
	}
	if err := r.All(db.ctx, result); err != nil {
		return &ListCollectionsResult{Err: err}
	}
	return &ListCollectionsResult{}
}

// ListCollectionNames lists all collection names in the database.
func (db *Database) ListCollectionNames(filter any, opts ...*ListCollectionsOptions) *ListCollectionsResult {
	var mongoOpts []*options.ListCollectionsOptions
	for _, opt := range opts {
		mongoOpts = append(mongoOpts, opt.o)
	}
	r, err := db.db.ListCollectionNames(db.ctx, filter, mongoOpts...)
	if err != nil {
		return &ListCollectionsResult{Err: err}
	}
	return &ListCollectionsResult{Names: r}
}

// IsCollectionExists checks if a collection with the given name exists.
func (db *Database) IsCollectionExists(v string) bool {
	r, err := db.db.ListCollectionNames(db.ctx, bson.M{"name": v})
	if err != nil {
		return false
	}
	return len(r) > 0
}

// Close closes the connection to the database.
func (db *Database) Close(ctx ...context.Context) error {
	c := db.ctx
	if len(ctx) > 0 {
		c = ctx[0]
	}
	db.db = nil
	return db.client.Disconnect(c)
}
