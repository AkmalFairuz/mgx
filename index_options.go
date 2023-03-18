package mgx

import "go.mongodb.org/mongo-driver/mongo/options"

// IndexOptions is a wrapper around mongo options.
type IndexOptions struct {
	o *options.IndexOptions
}

// opt returns the underlying mongo option.
func (o *IndexOptions) opt() *options.IndexOptions {
	return o.o
}

// Background specifies whether the index should be built in the background.
func (o *IndexOptions) Background(v ...bool) *IndexOptions {
	x := true
	if len(v) > 0 {
		x = v[0]
	}
	o.o.SetBackground(x)
	return o
}

// Collation specifies a collation.
func (o *IndexOptions) Collation(v *options.Collation) *IndexOptions {
	o.o.SetCollation(v)
	return o
}

// ExpireAfterSeconds specifies the number of seconds after which documents in the collection should expire.
func (o *IndexOptions) ExpireAfterSeconds(v int32) *IndexOptions {
	o.o.SetExpireAfterSeconds(v)
	return o
}

// Name specifies the name of the index.
func (o *IndexOptions) Name(v string) *IndexOptions {
	o.o.SetName(v)
	return o
}

// Sparse specifies whether the index only references documents with the specified field.
func (o *IndexOptions) Sparse(v ...bool) *IndexOptions {
	x := true
	if len(v) > 0 {
		x = v[0]
	}
	o.o.SetSparse(x)
	return o
}

// StorageEngine specifies storage engine options for the index.
func (o *IndexOptions) StorageEngine(v any) *IndexOptions {
	o.o.SetStorageEngine(v)
	return o
}

// Unique specifies whether the index should enforce a unique constraint on the indexed field.
func (o *IndexOptions) Unique(v ...bool) *IndexOptions {
	x := true
	if len(v) > 0 {
		x = v[0]
	}
	o.o.SetUnique(x)
	return o
}

// Version specifies the version number of the index.
func (o *IndexOptions) Version(v int32) *IndexOptions {
	o.o.SetVersion(v)
	return o
}

// Weights specifies the weight for each field in the index.
func (o *IndexOptions) Weights(v any) *IndexOptions {
	o.o.SetWeights(v)
	return o
}

// DefaultLanguage specifies the default language for the index.
func (o *IndexOptions) DefaultLanguage(v string) *IndexOptions {
	o.o.SetDefaultLanguage(v)
	return o
}

// LanguageOverride specifies the field in the document that contains the override language for the index.
func (o *IndexOptions) LanguageOverride(v string) *IndexOptions {
	o.o.SetLanguageOverride(v)
	return o
}

// TextVersion specifies the version of the text index.
func (o *IndexOptions) TextVersion(v int32) *IndexOptions {
	o.o.SetTextVersion(v)
	return o
}

// SphereVersion specifies the version of the 2dsphere index.
func (o *IndexOptions) SphereVersion(v int32) *IndexOptions {
	o.o.SetSphereVersion(v)
	return o
}

// Bits specifies the number of precision of the stored geohash value of the 2d index.
func (o *IndexOptions) Bits(v int32) *IndexOptions {
	o.o.SetBits(v)
	return o
}

// Max specifies the upper inclusive boundary for the longitude and latitude values for the 2d index.
func (o *IndexOptions) Max(v float64) *IndexOptions {
	o.o.SetMax(v)
	return o
}

// Min specifies the lower inclusive boundary for the longitude and latitude values for the 2d index.
func (o *IndexOptions) Min(v float64) *IndexOptions {
	o.o.SetMin(v)
	return o
}

// BucketSize specifies the number of units within which to group the location values for the geoHaystack index.
func (o *IndexOptions) BucketSize(v int32) *IndexOptions {
	o.o.SetBucketSize(v)
	return o
}

// PartialFilterExpression specifies a filter document that determines which documents in the collection should be included in the
// index.
func (o *IndexOptions) PartialFilterExpression(v any) *IndexOptions {
	o.o.SetPartialFilterExpression(v)
	return o
}

// Hidden specifies whether the index should be hidden.
func (o *IndexOptions) Hidden(v ...bool) *IndexOptions {
	x := true
	if len(v) > 0 {
		x = v[0]
	}
	o.o.SetHidden(x)
	return o
}

// WildcardProjection specifies a document that contains a field and its associated wildcard projection.
func (o *IndexOptions) WildcardProjection(v any) *IndexOptions {
	o.o.SetWildcardProjection(v)
	return o
}

// NewIndexOptions creates a new IndexOptions.
func NewIndexOptions() *IndexOptions {
	return &IndexOptions{o: options.Index()}
}
