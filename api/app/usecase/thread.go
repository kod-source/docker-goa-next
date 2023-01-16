package usecase

import (
	"context"

	"github.com/kod-source/docker-goa-next/app/model"
)

type ThreadUsecase interface {
	// Create スレッドの作成
	Create(ctx context.Context, text string, roomID model.RoomID, userID model.UserID, img *string) (*model.ThreadUser, error)
}
