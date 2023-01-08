package mock

import (
	"context"

	"github.com/kod-source/docker-goa-next/app/model"
	"github.com/kod-source/docker-goa-next/app/repository"
)

var _ (repository.UserRoomRepository) = (*MockUserRoomRepository)(nil)

type MockUserRoomRepository struct {
	CreateFunc func(ctx context.Context, roomID model.RoomID, userID model.UserID) (*model.UserRoom, error)
	DeleteFunc func(ctx context.Context, id model.UserRoomID) error
}

func (m *MockUserRoomRepository) Create(ctx context.Context, roomID model.RoomID, userID model.UserID) (*model.UserRoom, error) {
	return m.CreateFunc(ctx, roomID, userID)
}

func (m *MockUserRoomRepository) Delete(ctx context.Context, id model.UserRoomID) error {
	return m.DeleteFunc(ctx, id)
}
