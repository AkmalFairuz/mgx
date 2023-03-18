package mgx

import (
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

// CreateIndexesOptions is a wrapper around mongo options.
type CreateIndexesOptions struct {
	o *options.CreateIndexesOptions
}

// opt returns the underlying mongo option.
func (o *CreateIndexesOptions) opt() *options.CreateIndexesOptions {
	return o.o
}

// MaxTime specifies the maximum amount of time to allow the query to run.
func (o *CreateIndexesOptions) MaxTime(v time.Duration) *CreateIndexesOptions {
	o.o.SetMaxTime(v)
	return o
}

// CommitQuorumInt specifies the number of data-bearing members of a replica set, including the primary, that must complete the index builds successfully before the primary marks the indexes as ready.
func (o *CreateIndexesOptions) CommitQuorumInt(v int32) *CreateIndexesOptions {
	o.o.SetCommitQuorumInt(v)
	return o
}

// CommitQuorumString specifies the number of data-bearing members of a replica set, including the primary, that must complete the index builds successfully before the primary marks the indexes as ready.
func (o *CreateIndexesOptions) CommitQuorumString(v string) *CreateIndexesOptions {
	o.o.SetCommitQuorumString(v)
	return o
}

// CommitQuorumMajority specifies the number of data-bearing members of a replica set, including the primary, that must complete the index builds successfully before the primary marks the indexes as ready.
func (o *CreateIndexesOptions) CommitQuorumMajority() *CreateIndexesOptions {
	o.o.SetCommitQuorumMajority()
	return o
}

// CommitQuorumVotingMembers specifies the number of data-bearing members of a replica set, including the primary, that must complete the index builds successfully before the primary marks the indexes as ready.
func (o *CreateIndexesOptions) CommitQuorumVotingMembers() *CreateIndexesOptions {
	o.o.SetCommitQuorumVotingMembers()
	return o
}

// NewCreateIndexesOptions creates a new CreateIndexesOptions.
func NewCreateIndexesOptions() *CreateIndexesOptions {
	return &CreateIndexesOptions{o: options.CreateIndexes()}
}
