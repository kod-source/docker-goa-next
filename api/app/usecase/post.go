package usecase

import (
	"context"

	"github.com/kod-source/docker-goa-next/app/interactor"
	"github.com/kod-source/docker-goa-next/app/model"
	myerrors "github.com/kod-source/docker-goa-next/app/my_errors"
)

type PostUseCase interface {
	CreatePost(ctx context.Context, userID int, title string, img *string) (*model.Post, error)
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
