package mgx

import "go.mongodb.org/mongo-driver/mongo/options"

// ReplaceOptions is a wrapper around mongo options.
type ReplaceOptions struct {
	o *options.ReplaceOptions
}

// opt returns the underlying mongo option.
func (o *ReplaceOptions) opt() *options.ReplaceOptions {
	return o.o
}

// Collation specifies a collation.
func (o *ReplaceOptions) Collation(v *options.Collation) *ReplaceOptions {
	o.o.SetCollation(v)
	return o
}

// Hint specifies the index to use.
func (o *ReplaceOptions) Hint(v any) *ReplaceOptions {
	o.o.SetHint(v)
	return o
}

// Comment specifies a string or document that will be included in server logs, profiling logs, and currentOp queries to help trace
// the operation.  The default value is nil, which means that no comment will be included in the logs.
func (o *ReplaceOptions) Comment(v string) *ReplaceOptions {
	o.o.SetComment(v)
	return o
}

// Let specifies variables accessible from the replace expression.
func (o *ReplaceOptions) Let(v any) *ReplaceOptions {
	o.o.SetLet(v)
	return o
}

// Upsert specifies whether to insert a new document if no documents match the filter.
func (o *ReplaceOptions) Upsert(v ...bool) *ReplaceOptions {
	x := true
	if len(v) > 0 {
		x = v[0]
	}
	o.o.SetUpsert(x)
	return o
}

// BypassDocumentValidation specifies whether to bypass document validation during the replace operation.
func (o *ReplaceOptions) BypassDocumentValidation(v ...bool) *ReplaceOptions {
	x := true
	if len(v) > 0 {
		x = v[0]
	}
	o.o.SetBypassDocumentValidation(x)
	return o
}

// NewReplaceOptions creates a new ReplaceOptions.
func NewReplaceOptions() *ReplaceOptions {
	return &ReplaceOptions{o: options.Replace()}
}
