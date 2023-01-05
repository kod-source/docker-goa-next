package mock

import (
	"context"

	"github.com/kod-source/docker-goa-next/app/model"
	"github.com/kod-source/docker-goa-next/app/usecase"
)

var _ usecase.RoomUseCase = (*MockRoomUsecase)(nil)

type MockRoomUsecase struct {
	CreateFunc func(ctx context.Context, name string, isGroup bool, userIDs []model.UserID) (*model.RoomUser, error)
	IndexFunc  func(ctx context.Context, id model.UserID, nextID model.RoomID) ([]*model.IndexRoom, *int, error)
	ExistsFunc func(ctx context.Context, myID model.UserID, id model.UserID) (*model.Room, error)
	ShowFunc   func(ctx context.Context, id model.RoomID, myID model.UserID) (*model.RoomUser, error)
}

func (m *MockRoomUsecase) Create(ctx context.Context, name string, isGroup bool, userIDs []model.UserID) (*model.RoomUser, error) {
	return m.CreateFunc(ctx, name, isGroup, userIDs)
}

func (m *MockRoomUsecase) Index(ctx context.Context, id model.UserID, nextID model.RoomID) ([]*model.IndexRoom, *int, error) {
	return m.IndexFunc(ctx, id, nextID)
}

func (m *MockRoomUsecase) Exists(ctx context.Context, myID model.UserID, id model.UserID) (*model.Room, error) {
	return m.ExistsFunc(ctx, myID, id)
}

func (m *MockRoomUsecase) Show(ctx context.Context, id model.RoomID, myID model.UserID) (*model.RoomUser, error) {
	return m.ShowFunc(ctx, id, myID)
}
