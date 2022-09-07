package usecase

import (
	"context"

	"github.com/kod-source/docker-goa-next/app/interactor"
	"github.com/kod-source/docker-goa-next/app/model"
	myerrors "github.com/kod-source/docker-goa-next/app/my_errors"
)

type PostUseCase interface {
	CreatePost(ctx context.Context, userID int, title string, img *string) (*model.IndexPost, error)
	ShowAll(ctx context.Context, nextID int) ([]*model.IndexPostWithCountLike, *int, error)
	Delete(ctx context.Context, id int) error
	Update(ctx context.Context, id int, title string, img *string) (*model.IndexPost, error)
	Show(ctx context.Context, id int) (*model.ShowPost, error)
	ShowMyLike(ctx context.Context, userID, nextID int) ([]*model.IndexPostWithCountLike, *int, error)
	// ShowPostMy 指定したUserIDが投稿したものを取得する
	ShowPostMy(ctx context.Context, userID, nextID int) ([]*model.IndexPostWithCountLike, *int, error)
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

func (p *postUseCase) ShowAll(ctx context.Context, nextID int) ([]*model.IndexPostWithCountLike, *int, error) {
	indexPostsWithCountLike, ni, err := p.pi.ShowAll(ctx, nextID)
	if err != nil {
		return nil, nil, err
	}

	return indexPostsWithCountLike, ni, nil
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

func (p *postUseCase) ShowMyLike(ctx context.Context, userID, nextID int) ([]*model.IndexPostWithCountLike, *int, error) {
	ips, ni, err := p.pi.ShowMyLike(ctx, userID, nextID)
	if err != nil {
		return nil, nil, err
	}

	return ips, ni, nil
}

// ShowPostMy 指定したUserIDが投稿したものを取得する
func (p *postUseCase) ShowPostMy(ctx context.Context, userID, nextID int) ([]*model.IndexPostWithCountLike, *int, error) {
	ips, ni, err := p.pi.ShowPostMy(ctx, userID, nextID)
	if err != nil {
		return nil, nil, err
	}

	return ips, ni, nil
}
