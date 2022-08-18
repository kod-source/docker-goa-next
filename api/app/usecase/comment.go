package usecase

import (
	"context"

	"github.com/kod-source/docker-goa-next/app/interactor"
	"github.com/kod-source/docker-goa-next/app/model"
)

type CommentUsecase interface {
	Create(ctx context.Context, postID int, text string, img *string) (*model.Comment, error)
	ShowByPostID(ctx context.Context, postID int) ([]*model.Comment, error)
	Update(ctx context.Context, id int, text, img string) (*model.Comment, error)
	Delete(ctx context.Context, id int) error
}

type commentUsecase struct {
	ci interactor.CommentInteractor
}

func NewcommentUsecase(ci interactor.CommentInteractor) CommentUsecase {
	return &commentUsecase{ci: ci}
}

func (c *commentUsecase) Create(ctx context.Context, postID int, text string, img *string) (*model.Comment, error) {
	comment, err := c.ci.Create(ctx, postID, text, img)
	if err != nil {
		return nil, err
	}
	return comment, nil
}

func (c *commentUsecase) ShowByPostID(ctx context.Context, postID int) ([]*model.Comment, error) {
	comments, err := c.ci.ShowByPostID(ctx, postID)
	if err != nil {
		return nil, err
	}
	return comments, nil
}

func (c *commentUsecase) Update(ctx context.Context, id int, text, img string) (*model.Comment, error) {
	comment, err := c.ci.Update(ctx, id, text, img)
	if err != nil {
		return nil, err
	}
	return comment, nil
}

func (c *commentUsecase) Delete(ctx context.Context, id int) error {
	err := c.ci.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
