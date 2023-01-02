package mock

import (
	"context"

	"github.com/kod-source/docker-goa-next/app/model"
	"github.com/kod-source/docker-goa-next/app/usecase"
)

var _ usecase.UserRoomUseCase = (*MockUserRoomUsecase)(nil)

type MockUserRoomUsecase struct {
	InviteRoomFunc func(ctx context.Context, roomID model.RoomID, userID model.UserID) (model.UserRoom, error)
}

func (m *MockUserRoomUsecase) InviteRoom(ctx context.Context, roomID model.RoomID, userID model.UserID) (*model.UserRoom, error) {
	return m.InviteRoom(ctx, roomID, userID)
}
