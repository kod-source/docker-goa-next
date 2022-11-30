package usecase

import (
	"context"

	"github.com/kod-source/docker-goa-next/app/model"
)

type CommentUsecase interface {
	Create(ctx context.Context, postID, userID int, text string, img *string) (*model.CommentWithUser, error)
	ShowByPostID(ctx context.Context, postID int) ([]*model.CommentWithUser, error)
	Update(ctx context.Context, id int, text string, img *string) (*model.Comment, error)
	Delete(ctx context.Context, id int) error
}
