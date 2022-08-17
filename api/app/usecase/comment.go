package usecase

import (
	"context"

	"github.com/kod-source/docker-goa-next/app/interactor"
	"github.com/kod-source/docker-goa-next/app/model"
)

type CommentUsecase interface {
	Create(ctx context.Context, text string, img *string) (*model.Comment, error)
}

type commentUsecase struct {
	ci interactor.CommentInteractor
}

func NewcommentUsecase(ci interactor.CommentInteractor) CommentUsecase {
	return &commentUsecase{ci: ci}
}

func (c *commentUsecase) Create(ctx context.Context, text string, img *string) (*model.Comment, error) {
	comment, err := c.ci.Create(ctx, text, img)
	if err != nil {
		return nil, err
	}
	return comment, nil
}
