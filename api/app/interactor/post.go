package interactor

import (
	"context"
	"database/sql"
	"errors"

	"github.com/google/wire"
	"github.com/kod-source/docker-goa-next/app/model"
	myerrors "github.com/kod-source/docker-goa-next/app/my_errors"
	"github.com/kod-source/docker-goa-next/app/repository"
	"github.com/kod-source/docker-goa-next/app/usecase"
)

var _ usecase.PostUseCase = (*postInteractor)(nil)

var PostInteractorSet = wire.NewSet(
	NewPostInteractor,
	wire.Bind(new(usecase.PostUseCase), new(*postInteractor)),
)

type postInteractor struct {
	pr repository.PostRepository
}

func NewPostInteractor(pr repository.PostRepository) *postInteractor {
	return &postInteractor{pr: pr}
}

func (p *postInteractor) CreatePost(ctx context.Context, userID int, title string, img *string) (*model.IndexPost, error) {
	if len(title) == 0 {
		return nil, myerrors.ErrEmptyString
	}
	indexPost, err := p.pr.CreatePost(ctx, userID, title, img)
	if err != nil {
		return nil, err
	}
	return indexPost, nil
}

func (p *postInteractor) ShowAll(ctx context.Context, nextID int) ([]*model.IndexPostWithCountLike, *int, error) {
	indexPostsWithCountLike, nID, err := p.pr.ShowAll(ctx, nextID)
	if err != nil {
		return nil, nil, err
	}

	return indexPostsWithCountLike, nID, nil
}

func (p *postInteractor) Delete(ctx context.Context, id int) error {
	err := p.pr.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (p *postInteractor) Update(ctx context.Context, id int, title string, img *string) (*model.IndexPost, error) {
	indexPosts, err := p.pr.Update(ctx, id, title, img)
	if err != nil {
		return nil, err
	}

	return indexPosts, nil
}

func (p *postInteractor) Show(ctx context.Context, id int) (*model.ShowPost, error) {
	showPost, err := p.pr.Show(ctx, id)
	if err != nil {
		return nil, err
	}
	return showPost, nil
}

func (p *postInteractor) ShowMyLike(ctx context.Context, userID, nextID int) ([]*model.IndexPostWithCountLike, *int, error) {
	ips, nID, err := p.pr.ShowMyLike(ctx, userID, nextID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil, nil
		}
		return nil, nil, err
	}

	return ips, nID, nil
}

// ShowPostMy 指定したUserIDが投稿したものを取得する
func (p *postInteractor) ShowPostMy(ctx context.Context, userID, nextID int) ([]*model.IndexPostWithCountLike, *int, error) {
	ips, nID, err := p.pr.ShowPostMy(ctx, userID, nextID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil, nil
		}
		return nil, nil, err
	}

	return ips, nID, nil
}

func (p *postInteractor) ShowPostMedia(ctx context.Context, userID, nextID int) ([]*model.IndexPostWithCountLike, *int, error) {
	ips, nID, err := p.pr.ShowPostMedia(ctx, userID, nextID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil, nil
		}
		return nil, nil, err
	}
	return ips, nID, nil
}
