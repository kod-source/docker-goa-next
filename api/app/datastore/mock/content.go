package mock

import (
	"context"

	"github.com/kod-source/docker-goa-next/app/model"
	"github.com/kod-source/docker-goa-next/app/repository"
)

var _ repository.ContentRepository = (*MockContentRepository)(nil)

type MockContentRepository struct {
	DeleteFunc func(ctx context.Context, myID model.UserID, contentID model.ContentID) error
	CreateFunc func(ctx context.Context, text string, threadID model.ThreadID, myID model.UserID, img *string) (*model.ContentUser, error)
	GetByThreadFunc func(ctx context.Context, threadID model.ThreadID) ([]*model.ContentUser, error)
}

func (m *MockContentRepository) Delete(ctx context.Context, myID model.UserID, contentID model.ContentID) error {
	return m.DeleteFunc(ctx, myID, contentID)
}

func (m *MockContentRepository) Create(ctx context.Context, text string, threadID model.ThreadID, myID model.UserID, img *string) (*model.ContentUser, error) {
	return m.CreateFunc(ctx, text, threadID, myID, img)
}

func (m *MockContentRepository) GetByThread(ctx context.Context, threadID model.ThreadID) ([]*model.ContentUser, error) {
	return m.GetByThreadFunc(ctx, threadID)
}
