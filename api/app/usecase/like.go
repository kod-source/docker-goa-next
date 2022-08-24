package usecase

import (
	"context"

	"github.com/kod-source/docker-goa-next/app/interactor"
	"github.com/kod-source/docker-goa-next/app/model"
	myerrors "github.com/kod-source/docker-goa-next/app/my_errors"
)

type LikeUsecase interface {
	Create(ctx context.Context, userID, postID int) (*model.Like, error)
	Delete(ctx context.Context, userID, postID int) error
	GetPostIDs(ctx context.Context, userID int) ([]int, error)
}

type likeUsecase struct {
	li interactor.LikeInteractor
}

func NewLikeUsecase(li interactor.LikeInteractor) LikeUsecase {
	return &likeUsecase{li: li}
}

func (l *likeUsecase) Create(ctx context.Context, userID, postID int) (*model.Like, error) {
	if userID == 0 || postID == 0 {
		return nil, myerrors.BadRequestIntError
	}
	like, err := l.li.Create(ctx, userID, postID)
	if err != nil {
		return nil, err
	}

	return like, nil
}

func (l *likeUsecase) Delete(ctx context.Context, userID, postID int) error {
	if userID == 0 || postID == 0 {
		return myerrors.BadRequestIntError
	}
	err := l.li.Delete(ctx, userID, postID)
	if err != nil {
		return err
	}
	return nil
}

func (l *likeUsecase) GetPostIDs(ctx context.Context, userID int) ([]int, error) {
	postIDs, err := l.li.GetPostIDs(ctx, userID)
	if err != nil {
		return nil, err
	}
	return postIDs, nil
}
