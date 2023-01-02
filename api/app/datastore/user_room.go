package datastore

import (
	"context"

	"github.com/google/wire"
	"github.com/kod-source/docker-goa-next/app/model"
	"github.com/kod-source/docker-goa-next/app/repository"
)

var _ (repository.UserRoomRepository) = (*userRoomDatastore)(nil)

var UserRoomRepositorySet = wire.NewSet(
	NewUserRoomRepository,
	wire.Bind(new(repository.UserRoomRepository), new(*userRoomDatastore)),
)

type userRoomDatastore struct {
	tr repository.TimeRepository
}

func NewUserRoomRepository(tr repository.TimeRepository) *userRoomDatastore {
	return &userRoomDatastore{tr: tr}
}

func (urd *userRoomDatastore) Create(ctx context.Context, roomID model.RoomID, userID model.UserID) (*model.UserRoom, error) {
	return nil, nil
}
