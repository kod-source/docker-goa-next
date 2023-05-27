package usecase

import (
	"context"

	"github.com/kod-source/docker-goa-next/app/model"
)

//go:generate mockgen -source=./user_room.go -package=mock_usecase -destination=./mock/user_room_usecase.go

type UserRoomUseCase interface {
	// InviteRoom ルームに招待する
	InviteRoom(ctx context.Context, roomID model.RoomID, userID model.UserID) (*model.UserRoom, error)
	// Delete ルームから除外する
	Delete(ctx context.Context, id model.UserRoomID) error
}
