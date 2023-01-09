package repository

import (
	"context"

	"github.com/kod-source/docker-goa-next/app/model"
)

type RoomRepository interface {
	Create(ctx context.Context, name string, isGroup bool, userIDs []model.UserID, img *string) (*model.RoomUser, error)
	Delete(ctx context.Context, id model.RoomID) error
	Index(ctx context.Context, id model.UserID, nextID model.RoomID) ([]*model.IndexRoom, *int, error)
	// DMのグループを取得する
	GetNoneGroup(ctx context.Context, myID model.UserID, id model.UserID) (*model.Room, error)
	Show(ctx context.Context, id model.RoomID) (*model.RoomUser, error)
}
