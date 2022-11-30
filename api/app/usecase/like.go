package usecase

import (
	"context"

	"github.com/google/wire"
	"github.com/kod-source/docker-goa-next/app/datastore"
	"github.com/kod-source/docker-goa-next/app/model"
	myerrors "github.com/kod-source/docker-goa-next/app/my_errors"
)

var _ LikeUsecase = (*likeUsecase)(nil)

var LikeUseCaseSet = wire.NewSet(
	NewLikeUsecase,
	wire.Bind(new(LikeUsecase), new(*likeUsecase)),
)

type LikeUsecase interface {
	Create(ctx context.Context, userID, postID int) (*model.Like, error)
	Delete(ctx context.Context, userID, postID int) error
	GetPostIDs(ctx context.Context, userID int) ([]int, error)
}

type likeUsecase struct {
	ld datastore.LikeDatastore
}

func NewLikeUsecase(ld datastore.LikeDatastore) LikeUsecase {
	return &likeUsecase{ld: ld}
}

func (l *likeUsecase) Create(ctx context.Context, userID, postID int) (*model.Like, error) {
	if userID == 0 || postID == 0 {
		return nil, myerrors.BadRequestIntError
	}
	like, err := l.ld.Create(ctx, userID, postID)
	if err != nil {
		return nil, err
	}

	return like, nil
}

func (l *likeUsecase) Delete(ctx context.Context, userID, postID int) error {
	if userID == 0 || postID == 0 {
		return myerrors.BadRequestIntError
	}
	err := l.ld.Delete(ctx, userID, postID)
	if err != nil {
		return err
	}
	return nil
}

func (l *likeUsecase) GetPostIDs(ctx context.Context, userID int) ([]int, error) {
	postIDs, err := l.ld.GetPostIDs(ctx, userID)
	if err != nil {
		return nil, err
	}
	return postIDs, nil
}
