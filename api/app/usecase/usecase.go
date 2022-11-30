package usecase

import "github.com/google/wire"

// Set ...
var Set = wire.NewSet(
	UserUseCaseSet,
	PostUseCaseSet,
	LikeUseCaseSet,
	CommentUseCaseSet,
)
