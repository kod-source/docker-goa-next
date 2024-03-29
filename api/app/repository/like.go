package repository

import (
	"context"

	"github.com/kod-source/docker-goa-next/app/model"
)

//go:generate mockgen -source=./like.go -package=mock_repository -destination=./mock/like_repository.go

type LikeRepository interface {
	Create(ctx context.Context, userID, postID int) (*model.Like, error)
	Delete(ctx context.Context, userID, postID int) error
	GetPostIDs(ctx context.Context, userID int) ([]int, error)
}
