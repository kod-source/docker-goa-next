package usecase

import (
	"context"

	"github.com/google/wire"
	"github.com/kod-source/docker-goa-next/app/model"
	myerrors "github.com/kod-source/docker-goa-next/app/my_errors"
	"github.com/kod-source/docker-goa-next/app/repository"
)

var _ PostUseCase = (*postUseCase)(nil)

var PostUseCaseSet = wire.NewSet(
	NewPostUseCase,
	wire.Bind(new(PostUseCase), new(*postUseCase)),
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
	// ShowPostMedia 指定したUserIDが画像投稿したものを取得する
	ShowPostMedia(ctx context.Context, userID, nextID int) ([]*model.IndexPostWithCountLike, *int, error)
}

type postUseCase struct {
	pr repository.PostRepository
}

func NewPostUseCase(pr repository.PostRepository) *postUseCase {
	return &postUseCase{pr: pr}
}

func (p *postUseCase) CreatePost(ctx context.Context, userID int, title string, img *string) (*model.IndexPost, error) {
	if len(title) == 0 {
		return nil, myerrors.EmptyStringError
	}
	indexPost, err := p.pr.CreatePost(ctx, userID, title, img)
	if err != nil {
		return nil, err
	}
	return indexPost, nil
}

func (p *postUseCase) ShowAll(ctx context.Context, nextID int) ([]*model.IndexPostWithCountLike, *int, error) {
	indexPostsWithCountLike, nID, err := p.pr.ShowAll(ctx, nextID)
	if err != nil {
		return nil, nil, err
	}

	return indexPostsWithCountLike, nID, nil
}

func (p *postUseCase) Delete(ctx context.Context, id int) error {
	err := p.pr.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (p *postUseCase) Update(ctx context.Context, id int, title string, img *string) (*model.IndexPost, error) {
	indexPosts, err := p.pr.Update(ctx, id, title, img)
	if err != nil {
		return nil, err
	}

	return indexPosts, nil
}

func (p *postUseCase) Show(ctx context.Context, id int) (*model.ShowPost, error) {
	showPost, err := p.pr.Show(ctx, id)
	if err != nil {
		return nil, err
	}
	return showPost, nil
}

func (p *postUseCase) ShowMyLike(ctx context.Context, userID, nextID int) ([]*model.IndexPostWithCountLike, *int, error) {
	ips, nID, err := p.pr.ShowMyLike(ctx, userID, nextID)
	if err != nil {
		return nil, nil, err
	}

	return ips, nID, nil
}

// ShowPostMy 指定したUserIDが投稿したものを取得する
func (p *postUseCase) ShowPostMy(ctx context.Context, userID, nextID int) ([]*model.IndexPostWithCountLike, *int, error) {
	ips, nID, err := p.pr.ShowPostMy(ctx, userID, nextID)
	if err != nil {
		return nil, nil, err
	}

	return ips, nID, nil
}

func (p *postUseCase) ShowPostMedia(ctx context.Context, userID, nextID int) ([]*model.IndexPostWithCountLike, *int, error) {
	ips, nID, err := p.pr.ShowPostMedia(ctx, userID, nextID)
	if err != nil {
		return nil, nil, err
	}
	return ips, nID, nil
}
