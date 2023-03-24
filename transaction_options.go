package mgx

import (
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readconcern"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
	"time"
)

// TransactionOptions is a wrapper around mongo transaction options.
type TransactionOptions struct {
	o *options.TransactionOptions
}

func (o *TransactionOptions) opt() *options.TransactionOptions {
	return o.o
}

// ReadConcern specifies the read concern for the transaction.
func (o *TransactionOptions) ReadConcern(v *readconcern.ReadConcern) *TransactionOptions {
	o.o.SetReadConcern(v)
	return o
}

// ReadPref specifies the read preference for the transaction.
func (o *TransactionOptions) ReadPref(v *readpref.ReadPref) *TransactionOptions {
	o.o.SetReadPreference(v)
	return o
}

// WriteConcern specifies the write concern for the transaction.
func (o *TransactionOptions) WriteConcern(v *writeconcern.WriteConcern) *TransactionOptions {
	o.o.SetWriteConcern(v)
	return o
}

// MaxCommitTime specifies the maximum amount of time that a CommitTransaction operation can run on the server.
func (o *TransactionOptions) MaxCommitTime(v *time.Duration) *TransactionOptions {
	o.o.SetMaxCommitTime(v)
	return o
}

// NewTransactionOptions creates a new TransactionOptions.
func NewTransactionOptions() *TransactionOptions {
	return &TransactionOptions{}
}
