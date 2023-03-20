package mgx

import "go.mongodb.org/mongo-driver/mongo"

// IsErrNoDocuments returns true if the error is a mongo.ErrNoDocuments.
func IsErrNoDocuments(err error) bool {
	return mongo.ErrNoDocuments == err
}

// IsErrNilDocument ...
func IsErrNilDocument(err error) bool {
	return mongo.ErrNilDocument == err
}

// IsErrNilCursor ...
func IsErrNilCursor(err error) bool {
	return mongo.ErrNilCursor == err
}

// IsErrEmptySlice ...
func IsErrEmptySlice(err error) bool {
	return mongo.ErrEmptySlice == err
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
