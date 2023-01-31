package usecase

import (
	"context"

	"github.com/kod-source/docker-goa-next/app/model"
)

type ContentUsecase interface {
	// Delete スレッドの返信を削除する
	Delete(ctx context.Context, myID model.UserID, contentID model.ContentID) error
	// Create コンテントの返信を作成する
	Create(ctx context.Context, text string, threadID model.ThreadID, myID model.UserID, img *string) (*model.ContentUser, error)
}
