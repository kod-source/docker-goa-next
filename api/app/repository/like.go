package repository

import (
	"context"
	"database/sql"

	"github.com/kod-source/docker-goa-next/app/model"
)

//go:generate mockgen -source=./like.go -package=mock_repository -destination=./mock/like_repository.go

type LikeRepository interface {
	Create(ctx context.Context, tx *sql.Tx, userID, postID int) (*model.Like, error)
	Delete(ctx context.Context, tx *sql.Tx, userID, postID int) error
	GetPostIDs(ctx context.Context, tx *sql.Tx, userID int) ([]int, error)
}
