package interactor

import (
	"context"
	"database/sql"

	"github.com/kod-source/docker-goa-next/app/model"
	"github.com/kod-source/docker-goa-next/app/repository"
)

type CommentInteractor interface {
	Create(ctx context.Context, postID int, text string, img *string) (*model.Comment, error)
}

type commentInteractor struct {
	db *sql.DB
	tr repository.TimeRepository
}

func NewCommentInteractor(db *sql.DB, tr repository.TimeRepository) CommentInteractor {
	return &commentInteractor{db: db, tr: tr}
}

func (c *commentInteractor) Create(ctx context.Context, postID int, text string, img *string) (*model.Comment, error) {
	return nil, nil
}
