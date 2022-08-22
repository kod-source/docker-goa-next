package interactor

import (
	"context"
	"database/sql"

	"github.com/kod-source/docker-goa-next/app/model"
)

type LikeInteractor interface {
	Create(ctx context.Context, userID, postID int) (*model.Like, error)
}

type likeInteractor struct {
	db *sql.DB
}

func NewLikeInteractor(db *sql.DB) LikeInteractor {
	return &likeInteractor{db: db}
}

func (l *likeInteractor) Create(ctx context.Context, userID, postID int) (*model.Like, error) {
	return nil, nil
}
