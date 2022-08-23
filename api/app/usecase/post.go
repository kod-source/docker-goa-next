package usecase

import (
	"context"
	"fmt"

	"github.com/kod-source/docker-goa-next/app/interactor"
	"github.com/kod-source/docker-goa-next/app/model"
	myerrors "github.com/kod-source/docker-goa-next/app/my_errors"
)

type PostUseCase interface {
	CreatePost(ctx context.Context, userID int, title string, img *string) (*model.IndexPost, error)
	ShowAll(ctx context.Context, nextID int) ([]*model.IndexPostWithCountLike, *string, error)
	Delete(ctx context.Context, id int) error
	Update(ctx context.Context, id int, title string, img *string) (*model.IndexPost, error)
	Show(ctx context.Context, id int) (*model.ShowPost, error)
}

type postUseCase struct {
	pi interactor.PostInteractor
}

func NewPostUseCase(pi interactor.PostInteractor) PostUseCase {
	return &postUseCase{pi: pi}
}

func (p *postUseCase) CreatePost(ctx context.Context, userID int, title string, img *string) (*model.IndexPost, error) {
	if len(title) == 0 {
		return nil, myerrors.EmptyStringError
	}
	indexPost, err := p.pi.CreatePost(ctx, userID, title, img)
	if err != nil {
		return nil, err
	}
	return indexPost, nil
}

func (p *postUseCase) ShowAll(ctx context.Context, nextID int) ([]*model.IndexPostWithCountLike, *string, error) {
	indexPostsWithCountLike, nextToken, err := p.pi.ShowAll(ctx, nextID)
	if err != nil {
		fmt.Println("エラー")
		fmt.Println(err)
		return nil, nil, err
	}

	return indexPostsWithCountLike, nextToken, nil
}

func (p *postUseCase) Delete(ctx context.Context, id int) error {
	err := p.pi.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (p *postUseCase) Update(ctx context.Context, id int, title string, img *string) (*model.IndexPost, error) {
	indexPosts, err := p.pi.Update(ctx, id, title, img)
	if err != nil {
		return nil, err
	}

	return indexPosts, nil
}

func (p *postUseCase) Show(ctx context.Context, id int) (*model.ShowPost, error) {
	showPost, err := p.pi.Show(ctx, id)
	if err != nil {
		return nil, err
	}
	return showPost, nil
}
