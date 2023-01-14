package mock

import (
	"context"

	"github.com/kod-source/docker-goa-next/app/model"
	"github.com/kod-source/docker-goa-next/app/repository"
)

var _ repository.RoomRepository = (*MockRoomRepository)(nil)

type MockRoomRepository struct {
	CreateFunc       func(ctx context.Context, name string, isGroup bool, userIDs []model.UserID, img *string) (*model.RoomUser, error)
	DeleteFunc       func(ctx context.Context, id model.RoomID) error
	IndexFunc        func(ctx context.Context, id model.UserID, nextID model.RoomID) ([]*model.IndexRoom, *int, error)
	GetNoneGroupFunc func(ctx context.Context, myID model.UserID, id model.UserID) (*model.Room, error)
	ShowFunc         func(ctx context.Context, id model.RoomID) (*model.RoomUser, error)
}

func (m *MockRoomRepository) Create(ctx context.Context, name string, isGroup bool, userIDs []model.UserID, img *string) (*model.RoomUser, error) {
	return m.CreateFunc(ctx, name, isGroup, userIDs, img)
}

func (m *MockRoomRepository) Delete(ctx context.Context, id model.RoomID) error {
	return m.DeleteFunc(ctx, id)
}

func (m *MockRoomRepository) Index(ctx context.Context, id model.UserID, nextID model.RoomID) ([]*model.IndexRoom, *int, error) {
	return m.IndexFunc(ctx, id, nextID)
}

func (m *MockRoomRepository) GetNoneGroup(ctx context.Context, myID model.UserID, id model.UserID) (*model.Room, error) {
	return m.GetNoneGroupFunc(ctx, myID, id)
}

func (m *MockRoomRepository) Show(ctx context.Context, id model.RoomID) (*model.RoomUser, error) {
	return m.ShowFunc(ctx, id)
}
