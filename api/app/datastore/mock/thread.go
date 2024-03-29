package mock

import (
	"context"

	"github.com/kod-source/docker-goa-next/app/model"
	"github.com/kod-source/docker-goa-next/app/repository"
)

var _ repository.ThreadRepository = (*MockThreadRepository)(nil)

type MockThreadRepository struct {
	CreateFunc           func(ctx context.Context, text string, roomID model.RoomID, userID model.UserID, img *string) (*model.ThreadUser, error)
	DeleteFunc           func(ctx context.Context, myID model.UserID, threadID model.ThreadID) error
	GetThreadsByRoomFunc func(ctx context.Context, roomID model.RoomID, nextID model.ThreadID) ([]*model.IndexThread, *int, error)
}

func (m *MockThreadRepository) Create(ctx context.Context, text string, roomID model.RoomID, userID model.UserID, img *string) (*model.ThreadUser, error) {
	return m.CreateFunc(ctx, text, roomID, userID, img)
}

func (m *MockThreadRepository) Delete(ctx context.Context, myID model.UserID, threadID model.ThreadID) error {
	return m.DeleteFunc(ctx, myID, threadID)
}

func (m *MockThreadRepository) GetThreadsByRoom(ctx context.Context, roomID model.RoomID, nextID model.ThreadID) ([]*model.IndexThread, *int, error) {
	return m.GetThreadsByRoomFunc(ctx, roomID, nextID)
}
