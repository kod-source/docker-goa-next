package repository

import (
	"context"

	"github.com/kod-source/docker-goa-next/app/model"
)

type UserRoomRepository interface {
	Create(ctx context.Context, roomID model.RoomID, userID model.UserID) (*model.UserRoom, error)
}
