package usecase

import (
	"context"

	"github.com/kod-source/docker-goa-next/app/model"
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
