package mgx

import "go.mongodb.org/mongo-driver/mongo/options"

// UpdateOptions is a wrapper around mongo options.
type UpdateOptions struct {
	o *options.UpdateOptions
}

// opt returns the underlying mongo option.
func (o *UpdateOptions) opt() *options.UpdateOptions {
	return o.o
}

// Upsert specifies whether to insert a new document if no documents match the filter.
func (o *UpdateOptions) Upsert(v ...bool) *UpdateOptions {
	x := true
	if len(v) > 0 {
		x = v[0]
	}
	o.o.SetUpsert(x)
	return o
}

// BypassDocumentValidation specifies whether to bypass document validation during the update operation.
func (o *UpdateOptions) BypassDocumentValidation(v ...bool) *UpdateOptions {
	x := true
	if len(v) > 0 {
		x = v[0]
	}
	o.o.SetBypassDocumentValidation(x)
	return o
}

// ArrayFilters specifies an array of filters specifying to which array elements an update should apply.
func (o *UpdateOptions) ArrayFilters(v ArrayFilters) *UpdateOptions {
	o.o.SetArrayFilters(v.ArrayFilters)
	return o
}

// Collation specifies a collation.
func (o *UpdateOptions) Collation(v Collation) *UpdateOptions {
	o.o.SetCollation(v.Collation)
	return o
}

// Comment specifies a string or document that will be included in server logs, profiling logs, and currentOp queries to help trace
// the operation.  The default value is nil, which means that no comment will be included in the logs.
func (o *UpdateOptions) Comment(v string) *UpdateOptions {
	o.o.SetComment(v)
	return o
}

// Hint specifies the index to use.  The default value is nil, which means that no hint will be sent.
func (o *UpdateOptions) Hint(v any) *UpdateOptions {
	o.o.SetHint(v)
	return o
}

// Let specifies variables accessible from the update expression.
func (o *UpdateOptions) Let(v any) *UpdateOptions {
	o.o.SetLet(v)
	return o
}

// NewUpdateOptions creates a new UpdateOptions.
func NewUpdateOptions() *UpdateOptions {
	return &UpdateOptions{o: options.Update()}
}
