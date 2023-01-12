package mock

import (
	"context"

	"github.com/kod-source/docker-goa-next/app/model"
	"github.com/kod-source/docker-goa-next/app/repository"
)

var _ repository.ThreadRepository = (*MockThreadRepository)(nil)

type MockThreadRepository struct {
	CreateFunc func(ctx context.Context, text string, roomID model.RoomID, userID model.UserID, img *string) (*model.ThreadUser, error)
}

func (m *MockThreadRepository) Create(ctx context.Context, text string, roomID model.RoomID, userID model.UserID, img *string) (*model.ThreadUser, error) {
	return m.CreateFunc(ctx, text, roomID, userID, img)
}
