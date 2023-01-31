package mock

import (
	"context"

	"github.com/kod-source/docker-goa-next/app/model"
	"github.com/kod-source/docker-goa-next/app/usecase"
)

var _ usecase.ThreadUsecase = (*MockThreadUsecase)(nil)

type MockThreadUsecase struct {
	CreateFunc           func(ctx context.Context, text string, roomID model.RoomID, myID model.UserID, img *string) (*model.ThreadUser, error)
	DeleteFunc           func(ctx context.Context, myID model.UserID, threadID model.ThreadID) error
	GetThreadsByRoomFunc func(ctx context.Context, roomID model.RoomID, nextID model.ThreadID) ([]*model.IndexThread, *int, error)
}

func (m *MockThreadUsecase) Create(ctx context.Context, text string, roomID model.RoomID, myID model.UserID, img *string) (*model.ThreadUser, error) {
	return m.CreateFunc(ctx, text, roomID, myID, img)
}

func (m *MockThreadUsecase) Delete(ctx context.Context, myID model.UserID, threadID model.ThreadID) error {
	return m.DeleteFunc(ctx, myID, threadID)
}

func (m *MockThreadUsecase) GetThreadsByRoom(ctx context.Context, roomID model.RoomID, nextID model.ThreadID) ([]*model.IndexThread, *int, error) {
	return m.GetThreadsByRoomFunc(ctx, roomID, nextID)
}
