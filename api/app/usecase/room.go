package usecase

import (
	"context"

	"github.com/kod-source/docker-goa-next/app/model"
)

type RoomUseCase interface {
	// Create ルームの作成 DM,グルーム両方に対応
	Create(ctx context.Context, name string, isGroup bool, userIDs []model.UserID) (*model.RoomUser, error)
	// Index 自分が入っているルームの一覧
	Index(ctx context.Context, id model.UserID, nextID model.RoomID) ([]*model.IndexRoom, *int, error)
}
