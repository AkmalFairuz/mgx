package mgx

import "go.mongodb.org/mongo-driver/mongo"

// IsErrNotFound returns true if the error is a mongo.ErrNoDocuments.
func IsErrNotFound(err error) bool {
	return err == mongo.ErrNoDocuments
}

// IsErrDuplicateKey ...
func IsErrDuplicateKey(err error) bool {
	return mongo.IsDuplicateKeyError(err)
}

// IsErrTimeout ...
func IsErrTimeout(err error) bool {
	return mongo.IsTimeout(err)
}

// IsErrNetwork ...
func IsErrNetwork(err error) bool {
	return mongo.IsNetworkError(err)
}
