package mock

import (
	"context"

	"github.com/kod-source/docker-goa-next/app/model"
	"github.com/kod-source/docker-goa-next/app/repository"
)

var _ repository.RoomRepository = (*MockRoomRepository)(nil)

type MockRoomRepository struct {
	CreateFunc func(ctx context.Context, name string, isGroup bool, userIDs []model.UserID) (*model.RoomUser, error)
	DeleteFunc func(ctx context.Context, id model.RoomID) error
	IndexFunc  func(ctx context.Context, id model.UserID, nextID model.RoomID) ([]*model.RoomUser, *int, error)
}

func (m *MockRoomRepository) Create(ctx context.Context, name string, isGroup bool, userIDs []model.UserID) (*model.RoomUser, error) {
	return m.CreateFunc(ctx, name, isGroup, userIDs)
}

func (m *MockRoomRepository) Delete(ctx context.Context, id model.RoomID) error {
	return m.DeleteFunc(ctx, id)
}

func (m *MockRoomRepository) Index(ctx context.Context, id model.UserID, nextID model.RoomID) ([]*model.RoomUser, *int, error) {
	return m.IndexFunc(ctx, id, nextID)
}
