package datastore

import (
	"context"
	"database/sql"

	"github.com/google/wire"
	"github.com/kod-source/docker-goa-next/app/model"
	"github.com/kod-source/docker-goa-next/app/repository"
)

var _ repository.ThreadRepository = (*threadDatastore)(nil)

// ThreadRepositorySet ...
var ThreadRepositorySet = wire.NewSet(
	NewThreadRepository,
	wire.Bind(new(repository.ThreadRepository), new(*threadDatastore)),
)

type threadDatastore struct {
	db *sql.DB
	tr repository.TimeRepository
}

func NewThreadRepository(db *sql.DB, tr repository.TimeRepository) *threadDatastore {
	return &threadDatastore{db: db, tr: tr}
}

func (td *threadDatastore) Create(ctx context.Context, text string, roomID model.RoomID, userID model.UserID, img *string) (*model.ThreadUser, error) {
	return nil, nil
}
