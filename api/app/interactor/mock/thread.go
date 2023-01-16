package mock

import (
	"context"

	"github.com/kod-source/docker-goa-next/app/model"
	"github.com/kod-source/docker-goa-next/app/usecase"
)

var _ usecase.ThreadUsecase = (*MockThreadUsecase)(nil)

type MockThreadUsecase struct {
	CreateFunc func(ctx context.Context, text string, roomID model.RoomID, userID model.UserID, img *string) (*model.ThreadUser, error)
}

func (m *MockThreadUsecase) Create(ctx context.Context, text string, roomID model.RoomID, userID model.UserID, img *string) (*model.ThreadUser, error) {
	return m.CreateFunc(ctx, text, roomID, userID, img)
}
