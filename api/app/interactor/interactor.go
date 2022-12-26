package interactor

import "github.com/google/wire"

var Set = wire.NewSet(
	CommentInteractorSet,
	LikeInteractorSet,
	PostInteractorSet,
	UserInteractorSet,
	RoomInteractorSet,
)
