package mock

import (
	"context"

	"github.com/kod-source/docker-goa-next/app/model"
	"github.com/kod-source/docker-goa-next/app/repository"
)

var _ repository.ContentRepository = (*MockContentRepository)(nil)

type MockContentRepository struct {
	DeleteFunc func(ctx context.Context, myID model.UserID, contentID model.ContentID) error
}

func (m *MockContentRepository) Delete(ctx context.Context, myID model.UserID, contentID model.ContentID) error {
	return m.DeleteFunc(ctx, myID, contentID)
}
