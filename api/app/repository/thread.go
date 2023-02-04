package repository

import (
	"context"

	"github.com/kod-source/docker-goa-next/app/model"
)

type ThreadRepository interface {
	// Create ...
	Create(ctx context.Context, text string, roomID model.RoomID, userID model.UserID, img *string) (*model.ThreadUser, error)
	// Delete ...
	Delete(ctx context.Context, myID model.UserID, threadID model.ThreadID) error
	// GetThreadsByRoom ...
	GetThreadsByRoom(ctx context.Context, roomID model.RoomID, nextID model.ThreadID) ([]*model.IndexThread, *int, error)
}
