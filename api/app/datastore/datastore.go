package datastore

import "github.com/google/wire"

const (
	// データを取得する上限数
	LIMIT = 20
)

// Set ...
var Set = wire.NewSet(
	UserDatastoreSet,
	PostDatastoreSet,
	LikeDatastoreSet,
	CommentDatastoreSet,
	RoomDatastoreSet,
)
