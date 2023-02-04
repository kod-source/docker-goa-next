package repository

import (
	"context"

	"github.com/kod-source/docker-goa-next/app/model"
)

type ContentRepository interface {
	Delete(ctx context.Context, myID model.UserID, contentID model.ContentID) error
	Create(ctx context.Context, text string, threadID model.ThreadID, myID model.UserID, img *string) (*model.ContentUser, error)
	GetByThread(ctx context.Context, threadID model.ThreadID) ([]*model.ContentUser, error)
}
