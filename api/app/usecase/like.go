package usecase

import (
	"context"

	"github.com/kod-source/docker-goa-next/app/model"
)

type LikeUsecase interface {
	Create(ctx context.Context, userID, postID int) (*model.Like, error)
	Delete(ctx context.Context, userID, postID int) error
	GetPostIDs(ctx context.Context, userID int) ([]int, error)
}
