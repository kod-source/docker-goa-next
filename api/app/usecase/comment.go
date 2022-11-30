package usecase

import (
	"context"

	"github.com/google/wire"
	"github.com/kod-source/docker-goa-next/app/datastore"
	"github.com/kod-source/docker-goa-next/app/model"
	myerrors "github.com/kod-source/docker-goa-next/app/my_errors"
)

var _ CommentUsecase = (*commentUsecase)(nil)

var CommentUseCaseSet = wire.NewSet(
	NewCommentUsecase,
	wire.Bind(new(CommentUsecase), new(*commentUsecase)),
)

type CommentUsecase interface {
	Create(ctx context.Context, postID, userID int, text string, img *string) (*model.CommentWithUser, error)
	ShowByPostID(ctx context.Context, postID int) ([]*model.CommentWithUser, error)
	Update(ctx context.Context, id int, text string, img *string) (*model.Comment, error)
	Delete(ctx context.Context, id int) error
}

type commentUsecase struct {
	cd datastore.CommentDatastore
}

func NewCommentUsecase(cd datastore.CommentDatastore) CommentUsecase {
	return &commentUsecase{cd: cd}
}

func (c *commentUsecase) Create(ctx context.Context, postID, userID int, text string, img *string) (*model.CommentWithUser, error) {
	if len(text) == 0 {
		return nil, myerrors.BadRequestStingError
	}
	cu, err := c.cd.Create(ctx, postID, userID, text, img)
	if err != nil {
		return nil, err
	}
	return cu, nil
}

func (c *commentUsecase) ShowByPostID(ctx context.Context, postID int) ([]*model.CommentWithUser, error) {
	cus, err := c.cd.ShowByPostID(ctx, postID)
	if err != nil {
		return nil, err
	}
	return cus, nil
}

func (c *commentUsecase) Update(ctx context.Context, id int, text string, img *string) (*model.Comment, error) {
	if len(text) == 0 {
		return nil, myerrors.BadRequestStingError
	}
	comment, err := c.cd.Update(ctx, id, text, img)
	if err != nil {
		return nil, err
	}
	return comment, nil
}

func (c *commentUsecase) Delete(ctx context.Context, id int) error {
	err := c.cd.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
