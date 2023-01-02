package usecase

import (
	"context"

	"github.com/kod-source/docker-goa-next/app/model"
)

type UserRoomUseCase interface {
	// InviteRoom ルームに招待する
	InviteRoom(ctx context.Context, roomID model.RoomID, userID model.UserID) (*model.UserRoom, error)
}
