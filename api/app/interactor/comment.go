package interactor

import (
	"context"

	"github.com/google/wire"
	"github.com/kod-source/docker-goa-next/app/model"
	myerrors "github.com/kod-source/docker-goa-next/app/my_errors"
	"github.com/kod-source/docker-goa-next/app/repository"
	"github.com/kod-source/docker-goa-next/app/usecase"
)

var _ usecase.CommentUsecase = (*commentInteractor)(nil)

var CommentInteractorSet = wire.NewSet(
	NewCommentInteractor,
	wire.Bind(new(usecase.CommentUsecase), new(*commentInteractor)),
)

type commentInteractor struct {
	cr repository.CommentRepository
}

func NewCommentInteractor(cr repository.CommentRepository) *commentInteractor {
	return &commentInteractor{cr: cr}
}

func (c *commentInteractor) Create(ctx context.Context, postID, userID int, text string, img *string) (*model.CommentWithUser, error) {
	if len(text) == 0 {
		return nil, myerrors.ErrBadRequestSting
	}
	cu, err := c.cr.Create(ctx, postID, userID, text, img)
	if err != nil {
		return nil, err
	}
	return cu, nil
}

func (c *commentInteractor) ShowByPostID(ctx context.Context, postID int) ([]*model.CommentWithUser, error) {
	cus, err := c.cr.ShowByPostID(ctx, postID)
	if err != nil {
		return nil, err
	}
	return cus, nil
}

func (c *commentInteractor) Update(ctx context.Context, id int, text string, img *string) (*model.Comment, error) {
	if len(text) == 0 {
		return nil, myerrors.ErrBadRequestSting
	}
	comment, err := c.cr.Update(ctx, id, text, img)
	if err != nil {
		return nil, err
	}
	return comment, nil
}

func (c *commentInteractor) Delete(ctx context.Context, id int) error {
	err := c.cr.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
