package mock

import (
	"context"

	"github.com/kod-source/docker-goa-next/app/model"
	"github.com/kod-source/docker-goa-next/app/usecase"
)

var _ usecase.RoomUseCase = (*MockRoomUsecase)(nil)

type MockRoomUsecase struct {
	CreateFunc func(ctx context.Context, name string, isGroup bool, userIDs []model.UserID) (*model.RoomUser, error)
}

func (m *MockRoomUsecase) Create(ctx context.Context, name string, isGroup bool, userIDs []model.UserID) (*model.RoomUser, error) {
	return m.CreateFunc(ctx, name, isGroup, userIDs)
}
