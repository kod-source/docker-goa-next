package mock

import (
	"context"

	"github.com/kod-source/docker-goa-next/app/model"
	"github.com/kod-source/docker-goa-next/app/usecase"
)

var _ usecase.ContentUsecase = (*MockContentUsecase)(nil)

type MockContentUsecase struct {
	DeleteFunc func(ctx context.Context, myID model.UserID, contentID model.ContentID) error
}

func (m *MockContentUsecase) Delete(ctx context.Context, myID model.UserID, contentID model.ContentID) error {
	return m.DeleteFunc(ctx, myID, contentID)
}
