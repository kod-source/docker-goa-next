package interactor

import (
	"context"

	"github.com/google/wire"
	"github.com/kod-source/docker-goa-next/app/model"
	myerrors "github.com/kod-source/docker-goa-next/app/my_errors"
	"github.com/kod-source/docker-goa-next/app/repository"
	"github.com/kod-source/docker-goa-next/app/usecase"
)

var _ usecase.LikeUsecase = (*likeInteractor)(nil)

var LikeInteractorSet = wire.NewSet(
	NewLikeInteractor,
	wire.Bind(new(usecase.LikeUsecase), new(*likeInteractor)),
)

type likeInteractor struct {
	lr repository.LikeRepository
}

func NewLikeInteractor(lr repository.LikeRepository) *likeInteractor {
	return &likeInteractor{lr: lr}
}

func (l *likeInteractor) Create(ctx context.Context, userID, postID int) (*model.Like, error) {
	if userID == 0 || postID == 0 {
		return nil, myerrors.BadRequestIntError
	}
	like, err := l.lr.Create(ctx, userID, postID)
	if err != nil {
		return nil, err
	}

	return like, nil
}

func (l *likeInteractor) Delete(ctx context.Context, userID, postID int) error {
	if userID == 0 || postID == 0 {
		return myerrors.BadRequestIntError
	}
	err := l.lr.Delete(ctx, userID, postID)
	if err != nil {
		return err
	}
	return nil
}

func (l *likeInteractor) GetPostIDs(ctx context.Context, userID int) ([]int, error) {
	postIDs, err := l.lr.GetPostIDs(ctx, userID)
	if err != nil {
		return nil, err
	}
	return postIDs, nil
}
