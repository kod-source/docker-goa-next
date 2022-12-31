package repository

import (
	"context"

	"github.com/kod-source/docker-goa-next/app/model"
)

type RoomRepository interface {
	Create(ctx context.Context, name string, isGroup bool, userIDs []model.UserID) (*model.RoomUser, error)
	Delete(ctx context.Context, id model.RoomID) error
	Index(ctx context.Context, id model.UserID, nextID model.RoomID) ([]*model.IndexRoom, *int, error)
}
