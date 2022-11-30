package repository

import "github.com/google/wire"

// Set ...
var Set = wire.NewSet(
	TimeRepositorySet,
)
