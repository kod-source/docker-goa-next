package datastore

import (
	"context"
	"database/sql"

	"github.com/google/wire"
	"github.com/kod-source/docker-goa-next/app/model"
	"github.com/kod-source/docker-goa-next/app/repository"
)

var _ repository.RoomRepository = (*roomDatastore)(nil)

// RoomDatastoreSet ...
var RoomDatastoreSet = wire.NewSet(
	NewRoomDatastore,
	wire.Bind(new(repository.RoomRepository), new(*roomDatastore)),
)

type roomDatastore struct {
	db *sql.DB
	tr repository.TimeRepository
}

func NewRoomDatastore(db *sql.DB, tr repository.TimeRepository) *roomDatastore {
	return &roomDatastore{db: db, tr: tr}
}

// Create ルームの作成 DB処理
func (rd *roomDatastore) Create(ctx context.Context, name string, isGroup bool, userIDs []model.UserID) (*model.RoomUser, error) {
	return nil, nil
}
