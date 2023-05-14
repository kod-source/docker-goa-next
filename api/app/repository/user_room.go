package repository

import (
	"context"

	"github.com/kod-source/docker-goa-next/app/model"
)

//go:generate mockgen -source=./user_room.go -package=mock_repository -destination=./mock/user_room_repository.go

type UserRoomRepository interface {
	Create(ctx context.Context, roomID model.RoomID, userID model.UserID) (*model.UserRoom, error)
	Delete(ctx context.Context, id model.UserRoomID) error
}
