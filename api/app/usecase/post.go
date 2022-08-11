package usecase

import (
	"context"

	"github.com/kod-source/docker-goa-next/app/interactor"
	"github.com/kod-source/docker-goa-next/app/model"
	myerrors "github.com/kod-source/docker-goa-next/app/my_errors"
)

type PostUseCase interface {
	CreatePost(ctx context.Context, userID int, title string, img *string) (*model.Post, error)
	ShowAll(ctx context.Context) ([]*model.IndexPost, error)
	Delete(ctx context.Context, id int) error
}

type postUseCase struct {
	pi interactor.PostInteractor
}

func NewPostUseCase(pi interactor.PostInteractor) PostUseCase {
	return postUseCase{pi: pi}
}

func (p postUseCase) CreatePost(ctx context.Context, userID int, title string, img *string) (*model.Post, error) {
	if len(title) == 0 {
		return nil, myerrors.EmptyStringError
	}
	post, err := p.pi.CreatePost(ctx, userID, title, img)
	if err != nil {
		return nil, err
	}
	return post, nil
}

func (p postUseCase) ShowAll(ctx context.Context) ([]*model.IndexPost, error) {
	indexPosts, err := p.pi.ShowAll(ctx)
	if err != nil {
		return nil, err
	}

	return indexPosts, nil
}

func (p postUseCase) Delete(ctx context.Context, id int) error {
	err := p.pi.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
