package datastore

import "github.com/google/wire"

// Set ...
var Set = wire.NewSet(
	UserDatastoreSet,
	PostDatastoreSet,
	LikeDatastoreSet,
	JWTDatastoreSet,
	CommentDatastoreSet,
)
