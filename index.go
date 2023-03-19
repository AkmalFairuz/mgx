package mgx

import "go.mongodb.org/mongo-driver/mongo"

// IndexModel is a model for an index to create.
type IndexModel struct {
	// Keys is a document that contains the index keys and their corresponding sort order.
	Keys any

	Options *IndexOptions
}

// mongoIndex returns the underlying mongo index model.
func (m IndexModel) mongoIndex() mongo.IndexModel {
	return mongo.IndexModel{
		Keys:    m.Keys,
		Options: m.Options.opt(),
	}
}

// NewIndexModel creates a new IndexModel.
func NewIndexModel(keys any, options *IndexOptions) *IndexModel {
	return &IndexModel{
		Keys:    keys,
		Options: options,
	}
}
